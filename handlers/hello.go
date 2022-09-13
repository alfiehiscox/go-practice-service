package handlers

import (
    "net/http"
    "log"
    "fmt"
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
    h.l.Println("Hello Handler request incomming")
    fmt.Fprint(rw, "Hello, world!")
}

