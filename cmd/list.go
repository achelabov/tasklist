/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"tasklist/tasklist/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "short list",
	Long:  `long list`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadTasklist()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete")
		}
		for _, i := range tasks {
			fmt.Printf("[%d] %s\n", i.Key, i.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
