package main

import (
	"log"
	"net/http"
	"os"
	"github.com/dhyanio/boilerplate/handlers"

)

func main() {
	l := log.New(os.Stdout, "dhyanio-api", log.LstdFlags)
	hh := handlers.Newhello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}