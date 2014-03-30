package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/rtz12/flowette/database"
)

type Controller interface {
	Control()
}

func New(db *database.Database) Server {
	var s Server
	s.Init(db)
	return s
}

type Server struct {
	db          *database.Database
	controllers []Controller
}

func (s *Server) Init(db *database.Database) {
	s.db = db

	s.controllers = []Controller{
		Controller(&DateController{db}),
	}
	for _, c := range s.controllers {
		c.Control()
	}
}

func (s *Server) Serve(host string, port int) {
	p := strconv.Itoa(port)
	envPort := os.Getenv("PORT")
	if envPort != "" {
		p = envPort
	}
	envHost := os.Getenv("HOST")
	if envHost != "" {
		host = envHost
	}

	log.Printf("Server listening on %s:%d\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+p, nil))
}
