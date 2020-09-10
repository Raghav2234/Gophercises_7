package cmd

import (
	 "taskManager/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var indexes []int
		for _, arg := range args {
			index, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("There is some problem with conversion")
			}
			indexes = append(indexes, index)
		}
		for _, index := range indexes {
			err := db.DeleteTask(index)
			if err != nil {
				fmt.Println("Something went wrong: ", err.Error())
			}
		}
		// fmt.Println(indexes)
	},
}

func init() {
	RootCmd.AddCommand(DoCmd)
}
