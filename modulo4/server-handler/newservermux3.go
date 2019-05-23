// Type ServeMux Handle

// func (mux *ServeMux) Handle(pattern string, handler Handler)

// mux := http.NewServeMux()
// mux.Handle("/v1/api/ping", http.HandlerFunc(Ping))

package main

import (
	"log"
	"net/http"
	"time"
)

const (
	addr = ":8080"
)

func main() {

	mux := http.NewServeMux()

	// our function
	pingHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("\nDevops BH for Golang mux Handle()!"))
	}

	// handlerFunc
	mux.Handle("/v1/api/ping", http.HandlerFunc(pingHandler)) // ok

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Millisecond * 600,
		WriteTimeout: time.Millisecond * 400,
		// IdleTimeout:    1000 * time.Millisecond,
		//MaxHeaderBytes: cfg.MaxHeaderBytes,
	}

	log.Printf("\nServer run 8080\n")
	// Listen
	log.Fatal(srv.ListenAndServe())
}
