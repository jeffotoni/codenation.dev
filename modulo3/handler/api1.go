package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    if os.Getenv("FEATURE_TOGGLE") == "TRUE" {
        fmt.Println(os.Getenv("FEATURE_TOGGLE"))
        fmt.Fprintf(w, "Exciting New Feature")
    } else {
        fmt.Println(os.Getenv("FEATURE_TOGGLE"))
        fmt.Fprintf(w, "existing boring feature")
    }
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}
