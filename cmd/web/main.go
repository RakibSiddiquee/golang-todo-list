package main

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"os"
)

type server struct {
	host string
	port string
	url  string
}
type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
}

func main() {
	// Init the server
	server := server{
		host: "localhost",
		port: "8091",
		url:  "http://localhost:8091",
	}

	// Init the application
	app := &application{
		server:  server,
		appName: "Golang Todo List",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stdout, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	// init jet template
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))
	}
}
