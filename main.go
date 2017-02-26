package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var greeting string

func main() {
	cmd := &cobra.Command{}
	cmd.Use = "hello [name]"
	cmd.Flags().StringVar(&greeting, "greeting", "Hello", "")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("name argument is required")
		}

		fmt.Printf("NArg: %d\n", len(args))
		fmt.Printf("%s, %s\n", greeting, args[0])

		return nil
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
