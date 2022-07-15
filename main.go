package main

import (
	"log"
	"tasklist/tasklist/cmd"
	"tasklist/tasklist/db"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
