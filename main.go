package main

import (
	"os"

	"github.com/rtz12/flowette/commands"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "flowette"
	app.Usage = "backend server for flostatus.neocities.org"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{"database, db", "", "location of the database file"},
	}
	app.Commands = []cli.Command{
		commands.New(),
		commands.Records(),
		commands.Serve(),
	}
	app.Run(os.Args)
}
