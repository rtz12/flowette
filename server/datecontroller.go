package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rtz12/flowette/database"
)

type DateController struct {
	db *database.Database
}

func (c *DateController) Control() {
	log.Println("Register DateController...")
	http.HandleFunc("/dates/", c.datesHandler)
}

func (c *DateController) datesHandler(w http.ResponseWriter, r *http.Request) {
	dates := c.db.GetAllDates()
	data, err := json.Marshal(dates)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(data)
}
