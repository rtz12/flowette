package commands

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/rtz12/flowette/database"
	"github.com/rtz12/flowette/helper"

	"github.com/codegangsta/cli"
)

func getStatusBool(status string) bool {
	var statusBool bool
	switch status {
	case "true":
		statusBool = true
	case "false":
		statusBool = false
	default:
		log.Fatal("Invalid value for `status`")
	}
	return statusBool
}

func Records() cli.Command {
	return cli.Command{
		Name:      "records",
		ShortName: "r",
		Usage:     "manages records for the date database",
		Flags: []cli.Flag{
			cli.BoolFlag{"list, l", "list all records"},
			cli.BoolFlag{"add, a", "add record to database"},
			cli.StringFlag{"date, d", "", "date for adding or filtering"},
			cli.StringFlag{"status, s", "", "status for adding or filtering, valid values are `true` or `false`"},
		},
		Action: func(c *cli.Context) {
			dateString := c.String("date")
			status := c.String("status")
			var date time.Time
			if dateString != "" {
				var err error
				date, err = time.Parse(helper.DateFormat, dateString)
				if err != nil {
					log.Fatal(fmt.Sprintf(""+
						"Could not parse date! "+
						"Make sure date is in the following format: %s", helper.DateFormat))
				}
			}
			switch {
			case c.Bool("add"):
				if c.Bool("list") {
					log.Fatal("Please choose only one action")
				}

				if date.IsZero() {
					log.Fatal("Date must be set for adding")
				}

				if status == "" {
					log.Fatal("Status must be set for adding")
				}

				statusBool := getStatusBool(status)

				dbPath := helper.GetDBPath(c.String("database"))
				db := database.Open(dbPath)
				db.AddDate(date, statusBool)
				db.Close()
				fmt.Println("Record added")
			default:
				if c.Bool("add") {
					log.Fatal("Please choose only one action")
				}

				filter := database.FilterStatus
				statusBool := false
				if status == "true" {
					statusBool = true
				} else if status != "false" {
					filter = 0
				}
				if !date.IsZero() {
					filter = filter | database.FilterDate
				}

				dbPath := helper.GetDBPath(c.String("database"))
				db := database.Open(dbPath)
				dates := db.GetDates(date, statusBool, filter)
				db.Close()

				tb := tabwriter.NewWriter(os.Stdout, 1, 8, 0, '\t', 0)
				_, err := tb.Write([]byte("Date\tStatus\n"))
				if err == nil {
					for _, e := range dates {
						formattedDate := e.Date.Format(helper.DateFormat)
						row := fmt.Sprintf("%s\t%t\n", formattedDate, e.Status)
						_, err = tb.Write([]byte(row))
						if err != nil {
							break
						}
					}
				}
				if err == nil {
					err = tb.Flush()
				}
				if err != nil {
					log.Fatal("Error writing to output stream")
				}
			}
		},
	}
}
