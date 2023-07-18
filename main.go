package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/hibiki-horimi/go-todo-api/cmd"
)

func main() {
	c := &cobra.Command{Use: "Select CLI Mode"}
	cmd.Register(c)
	err := c.Execute()
	if err != nil {
		log.Printf("An error occurred: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
