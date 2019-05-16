package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	fmt.Println("goroutine..")
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 10 goroutines
		go work(&wg)
	}

	time.Sleep(5 * time.Second)
}

func work(wg *sync.WaitGroup) {
	var count int
	for i := 0; i < 100000; i++ {
		count++
	}
	wg.Done()
}
