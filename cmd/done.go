package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pravuX/tri/todo"
)

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		// mark as done
		if toggleDone {
			items[i-1].Done = !items[i-1].Done
			fmt.Printf("%q %v\n", items[i-1].Text, "toggled")
		} else {
			items[i-1].Done = true
			fmt.Printf("%q %v\n", items[i-1].Text, "marked done")
		}
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done <task_id>",
	Aliases: []string{"do"},
	Short:   "mark task as done",
	Long:    `Done is used to indicate a task being completed.`,
	Args:    cobra.ExactArgs(1),
	Run:     doneRun,
}

var toggleDone bool

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.Flags().BoolVarP(&toggleDone, "toggle", "t", false, "toggle done status of task")
}
