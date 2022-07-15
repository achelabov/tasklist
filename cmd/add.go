package cmd

import (
	"log"
	"strings"
	"tasklist/tasklist/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task",
	Long:  `add task`,
	Run: func(cmd *cobra.Command, args []string) {
		err := db.CreateTask(strings.Join(args, " "))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
