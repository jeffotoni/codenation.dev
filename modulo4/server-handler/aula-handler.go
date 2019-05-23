package main

import (
	"fmt"
	"net/http"
	"strings"
)

func Ping(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method: ", r.Method)

	if strings.ToUpper(r.Method) == "POST" {

		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error permitido somente POST...\n"))

}

func main() {

	http.Handle("/api/v1/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method: ", r.Method)
		if strings.ToUpper(r.Method) == "POST" {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error permitido somente POST...\n"))
	}))
	http.ListenAndServe(":8080", nil)
}
