package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pravuX/tri/todo"
)

// addCmd represents the add command
var removeCmd = &cobra.Command{
	Use:     "remove <task_id>",
	Aliases: []string{"rm"},
	Short:   "remove a task",
	Long:    `Remove deletes item from the task list.`,
	Args:    cobra.ExactArgs(1),
	Run:     removeRun,
}

func removeRun(cmd *cobra.Command, args []string) {
	// Add to the existing list of tasks.
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}
	if i > 0 && i <= len(items) {
		fmt.Printf("%q %v\n", items[i-1].Text, "removed")
		newItems := append(items[:i-1], items[i:]...)
		if err := todo.SaveItems(viper.GetString("datafile"), newItems); err != nil {
			fmt.Errorf("%v", err)
		}
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
