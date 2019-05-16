package main

import (
	"fmt"
	//"os"
	"sync"
	//"time"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("goroutine..")
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}

	wg.Wait()

	// Aguarde até que a fila de execução
	// global se esgote.
	time.Sleep(5 * time.Second)
}

func work(wg *sync.WaitGroup) {

	//time.Sleep(time.Second)
	var counter int
	for i := 0; i < 100000; i++ {
		counter++
	}
	//fmt.Println(counter)
	wg.Done()
}
