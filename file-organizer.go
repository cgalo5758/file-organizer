package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "file-organizer",
		Usage: "Organize files according to a given specification",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
