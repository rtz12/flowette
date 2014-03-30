package database

import (
	"time"
)

type Date struct {
	Date   time.Time `json:"date"`
	Status bool      `json:"status"`
}

func getStatusBool(status int) bool {
	if status == 1 {
		return true
	}
	return false
}

func dateFromDB(date int, status int) Date {
	return Date{
		time.Unix(int64(date), 0),
		getStatusBool(status)}
}
