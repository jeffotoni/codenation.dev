package main

import "fmt"

func main() {

	// ch <- v    // Send v to channel ch. / READ
	// v := <-ch  // Receive from ch, and  / WRITE

	ch := make(chan int, 1)
	//ch <- 1
	//var ch chan int
	//ch <- 1
	//ch <- 2
	//ch <- 3
	//ch <- 4
	//ch <- 5
	//<-ch
	fmt.Println("oi")
	//fmt.Println(<-ch)
}
