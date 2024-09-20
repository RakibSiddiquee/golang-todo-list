package main

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"log"
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
	fmt.Println("Hello World")
}
