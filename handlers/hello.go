package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// Constructor
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// HTTP Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Received Request to Hello Handler")
	fmt.Fprint(rw, "Hello, world!")
}
