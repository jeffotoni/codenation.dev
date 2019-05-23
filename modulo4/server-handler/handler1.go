package main

import (
	"fmt"
	"log"
	"net/http"
)

//  Type Handler
// type Handler interface {
//         ServeHTTP(ResponseWriter, *Request)
// }

type pingHandler struct{}

func (h pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevopsBH for Golang simple %s\n", r.URL.Path)
}

// curl -i -Xget localhost:8080/v1/api/ping
func main() {
	log.Printf("\nServer run 8080\n")

	// ListenAndServe inicia um servidor HTTP com um determinado endereço e manipulador.
	// O manipulador geralmente é nulo, o que significa usar DefaultServeMux.
	// Handle e HandleFunc adicionam manipuladores ao DefaultServeMux
	err := http.ListenAndServe(":8080", pingHandler{})
	log.Fatal(err)
}
