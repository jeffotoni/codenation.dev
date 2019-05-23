package main

import (
    "log"
    "net/http"
    "time"
)

func main() {

    client := http.Client{
        Timeout: time.Duration(3 * time.Second),
    }

    body, err := client.Get("https://golangcode.com")
    if err != nil {
        log.Fatal(err)
    }

    log.Println(body)
}
