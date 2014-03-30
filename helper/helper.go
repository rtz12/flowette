package helper

import (
	"log"
	"os"
	"os/user"
	"path"
)

const (
	DateFormat = "2006-01-02"
)

func GetDBPath(pathFromArgs string) string {
	if pathFromArgs != "" {
		return pathFromArgs
	}

	filePath := "flowette.db"
	if _, err := os.Stat(filePath); err == nil {
		return filePath
	}

	if usr, err := user.Current(); err == nil {
		filePath = path.Join(usr.HomeDir, "."+filePath)
		if _, err := os.Stat(filePath); err == nil {
			return filePath
		}
	}

	log.Fatal("Could not locate flowette.db")
	return ""
}
