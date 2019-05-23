package main

import (
	"fmt"
	//"log"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//w.Write
	io.WriteString(w, "Curso golang porta 8080... \n")
}

func main() {

	//confserver.ListenAndServe()
	// http.ListenAndServe(":8080", nill)
	fmt.Println("Server Run:8080")

	// async
	go func() {
		http.ListenAndServe(":8080", http.HandlerFunc(Index))
	}()

	go func() {
		http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//w.Write
			io.WriteString(w, "Curso golang 8081... \n")
		}))
	}()

	http.ListenAndServe(":8082", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Write
		io.WriteString(w, "Curso golang 8082... \n")
	}))

	//http.ListenAndServe(":8082", nil)
	//http.ListenAndServe(":8083", nil)
}
