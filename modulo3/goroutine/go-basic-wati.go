package main

import (
	"fmt"
	//"sync"
	//"time"
)

func main() {

	// unbuffered channel of ints
	ic := make(chan int)

	//var wg sync.WaitGroup
	//wg.Add(2)
	go func() {
		// Do work.
		//wg.Done()
		fmt.Println("func 1")
		ic <- 1
	}()
	go func() {
		// Do work.
		//wg.Done()
		fmt.Println("func 2")
		ic <- 2
	}()

	// wg.Wait()

	<-ic
	<-ic

	fmt.Println("feito")

	//time.Sleep(time.Second * 2)
	// for {
	// }
	// select {} error
}
