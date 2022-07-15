package cmd

import (
	"log"
	"strconv"
	"tasklist/tasklist/db"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task",
	Long:  `delete task`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = db.DeleteTask(key)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
