package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pravuX/tri/todo"
)

var doneOpt, allOpt, fieldsOpt bool

var showPriority int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list the tasks",
	Long:    `List displays all of the tasks.`,
	Run:     listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	if fieldsOpt {
		fmt.Fprintln(
			w,
			"task_id\tdone\tpriority\ttask",
		)
	}
	for _, item := range items {
		if showPriority > 0 && showPriority <= 3 {
			if item.Priority == showPriority {
				fmt.Fprintln(
					w,
					item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t",
				)
			}
		} else if allOpt || item.Done == doneOpt {
			fmt.Fprintln(
				w,
				item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t",
			)
		}
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "show completed tasks")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "show all tasks")
	listCmd.Flags().
		BoolVarP(&fieldsOpt, "fields", "f", false, "show field name describing each column")
	listCmd.Flags().IntVarP(&showPriority, "priority", "p", 0, "show tasks of specified priority")
}
