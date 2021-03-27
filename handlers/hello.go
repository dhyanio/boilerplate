package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger

}
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye
}

func (g *Goodbye) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}