package cmd

import (
	"taskManager/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add command is used to add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("There is some problem adding the task")
			return
		}
		fmt.Println("Added", task, "to the list.")
		fmt.Println("Hello")
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
