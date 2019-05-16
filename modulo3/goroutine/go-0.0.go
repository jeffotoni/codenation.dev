package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {

    messages := make(chan int)
    //done := make(chan bool)

    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        defer wg.Done()
        time.Sleep(time.Second * 3)
        messages <- 1
        //done <- true
    }()

    go func() {
        defer wg.Done()
        time.Sleep(time.Second * 2)
        messages <- 2
        //done <- true
    }()

    go func() {
        defer wg.Done()
        time.Sleep(time.Second * 1)
        messages <- 3
        //done <- true
    }()

    go func() {
        for i := range messages {
            fmt.Println(i)
            time.Sleep(time.Second * 1)
        }
    }()

    wg.Wait()
    //time.Sleep(time.Second * 5)

    // for i := 0; i < 3; i++ {
    //     <-done
    // }
}
