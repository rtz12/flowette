package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type FilterMode int

const (
	FilterDate   FilterMode = 1 << iota
	FilterStatus FilterMode = 1 << iota
)

func Open(filename string) *Database {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal("Error opening Database")
	}
	return &Database{
		db: db,
	}
}

type Database struct {
	lock sync.Mutex
	db   *sql.DB
}

func (d *Database) Init() {
	sql := `
	CREATE TABLE "dates" (
		"date"  "INTEGER",
		"status" "INTEGER"
	);`
	if _, err := d.db.Exec(sql); err != nil {
		d.Close()
		log.Fatal("Error while initialising")
	}
}

func (d *Database) Close() {
	d.db.Close()
}

func getStatusInt(status bool) int8 {
	if status {
		return 1
	}
	return 0
}

func (d *Database) AddDate(date time.Time, status bool) {
	d.lock.Lock()
	defer d.lock.Unlock()

	sql := fmt.Sprintf(
		"INSERT INTO dates VALUES (%d, %d);",
		date.Unix(),
		getStatusInt(status))

	_, err := d.db.Exec(sql)

	if err != nil {
		log.Fatal("Error inserting date")
	}
}

func (d *Database) GetAllDates() []Date {
	return d.GetDates(time.Now(), false, 0)
}

func (d *Database) GetDates(date time.Time, status bool, filter FilterMode) []Date {
	d.lock.Lock()
	defer d.lock.Unlock()

	sql := "SELECT * FROM dates"
	if (filter & FilterDate) == FilterDate {
		sql += fmt.Sprintf(" WHERE (date = %d) AND", date.Unix())
	}
	if (filter & FilterStatus) == FilterStatus {
		sql += fmt.Sprintf(" WHERE (status = %d) AND", getStatusInt(status))
	}
	if filter > 0 {
		sql = sql[:len(sql)-4] + ";"
	}

	rows, err := d.db.Query(sql)
	if err != nil {
		log.Fatal("Error retreiving dates")
	}

	var dates []Date
	defer rows.Close()
	for rows.Next() {
		var dbDate int
		var dbStatus int
		rows.Scan(&dbDate, &dbStatus)
		dates = append(dates, dateFromDB(dbDate, dbStatus))
	}

	return dates
}
