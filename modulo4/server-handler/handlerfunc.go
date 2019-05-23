package main

import (
	"fmt"
	"log"
	"net/http"
)

// type HandlerFunc func(ResponseWriter, *Request)
// handlerApiPing := http.HandlerFunc(Ping)
func main() {

	handlerfunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DevopsBH for Golang simple two %s\n", r.URL.Path)
		//w.Write([]byte("outra forma de escrever.."))
		//io.WriteString(w, "outra forma de escrever..")
	})

	log.Printf("\nServer run 8080\n")
	err := http.ListenAndServe(":8080", handlerfunc)
	log.Fatal(err)
}
