package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/rtz12/flowette/database"

	"github.com/codegangsta/cli"
)

func New() cli.Command {
	return cli.Command{
		Name:  "new",
		Usage: "creates a new database",
		Action: func(c *cli.Context) {
			filename := c.Args().First()
			if filename == "" {
				log.Fatal("No file specified!")
			}
			if _, err := os.Stat(filename); err == nil {
				log.Fatal("File already exist or path not accessible")
			}
			db := database.Open(filename)
			db.Init()
			db.Close()
			fmt.Println("Database successfully initialised!")
		},
	}
}
