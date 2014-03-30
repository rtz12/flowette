package commands

import (
	"log"
	"strconv"

	"github.com/rtz12/flowette/database"
	"github.com/rtz12/flowette/helper"
	"github.com/rtz12/flowette/server"

	"github.com/codegangsta/cli"
)

func Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "starts server",
		Flags: []cli.Flag{
			cli.StringFlag{"port, p", "3000", "port to listen on"},
			cli.StringFlag{"host, n", "", "host to listen on"},
		},
		Action: func(c *cli.Context) {
			port, err := strconv.Atoi(c.String("port"))
			if err != nil {
				log.Fatal("Port has to be a number")
			}
			if port < 1024 || port > 49151 {
				log.Fatal("You can only use ports between 1024 and 49151")
			}
			log.Println("Open database...")
			db := database.Open(helper.GetDBPath(c.String("database")))
			log.Println("Create server...")
			s := server.New(db)
			log.Println("Listen...")
			s.Serve(c.String("host"), port)
		},
	}
}
