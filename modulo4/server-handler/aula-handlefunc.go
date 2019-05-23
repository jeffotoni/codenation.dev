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

	Ping2 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json := `{"status":"ok", "msg":"sucesso"}`
		json += "\n"
		w.Write([]byte(json))
	}

	http.HandleFunc('"/api/v1/ping1", Ping')

	http.Handle("/api/v1/ping2", http.HandlerFunc(Ping2))

	http.HandleFunc("/api/v1/ping3", Ping)
	http.HandleFunc("/api/v1/ping4", Ping)

	http.ListenAndServe(":8080", nil)
}
