package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pravuX/tri/todo"
)

var (
	newName     string
	newPriority int
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <task_id>",
	Short: "update task name or priority",
	Long:  `Modifies name or priority of selected task`,
	Args:  cobra.ExactArgs(1),
	Run:   updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}
	if i > 0 && i <= len(items) {
		// update text
		if newName != "" {
			fmt.Printf("%q name updated to %q\n", items[i-1].Text, newName)
			items[i-1].Text = newName
		}
		// update priority
		if newPriority != 0 {
			items[i-1].SetPriority(newPriority)
			fmt.Printf("%q priority updated to %v\n", items[i-1].Text, newPriority)
		}
		if err := todo.SaveItems(viper.GetString("datafile"), items); err != nil {
			fmt.Errorf("%v", err)
		}
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&newName,
		"name", "n", "", "new name of task")
	updateCmd.Flags().IntVarP(&newPriority,
		"priority", "p", 0, "priority: 1, 2, 3")
}
