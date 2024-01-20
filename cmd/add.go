package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pravuX/tri/todo"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  `Add creates a new item in the task list.`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	// Add to the existing list of tasks.
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)
		items = append(items, item)
		fmt.Printf("%q added to task list with priority %v\n", arg, priority)
	}
	if err := todo.SaveItems(viper.GetString("datafile"), items); err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority,
		"priority", "p", 2, "priority: 1, 2, 3")
}
