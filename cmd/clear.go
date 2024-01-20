package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pravuX/tri/todo"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear the task list",
	Long:  `Clear removes all items from the todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task list cleared")
		if err := todo.SaveItems(viper.GetString("datafile"), []todo.Item{}); err != nil {
			fmt.Errorf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
