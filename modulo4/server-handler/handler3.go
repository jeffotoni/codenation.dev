package main

import (
    "io"
    "net/http"
)

func main() {
    // Create a basic http example to demonstrate example
    http.Handle("/", http.HandlerFunc(ExampleHandler))
    http.ListenAndServe(":8080", nil)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodGet {
        io.WriteString(w, "This is a get request")
    } else if r.Method == http.MethodPost {
        io.WriteString(w, "This is a post request")
    } else {
        io.WriteString(w, "This is a "+r.Method+" request")
    }
}
