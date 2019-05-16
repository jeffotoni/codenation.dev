package main

import "fmt"

// unidirecional - somente write
// send to receive-only type <-chan int
func f1(c <-chan int) {
	fmt.Println("Write, Desempilhar: ", <-c)
}

// unidirecional - somente read
// receive from send-only type chan<- int
func f2(c chan<- int) {
	fmt.Println("Read, Empilhar 300")
	c <- 300
}

// ch <- v    // Send v to channel ch. /read
// v := <-ch  // Receive from ch, and  /write
func main() {

	// channel bidirecional
	c1 := make(chan int)
	go f2(c1)
	f1(c1)

	// unidirecional - somente write
	// send to receive-only type <-chan int
	rc1 := make(<-chan int, 1)
	fmt.Println("rc1: ", rc1)

	// unidirecional - somente read
	// receive from send-only type chan<- int
	sc1 := make(chan<- int, 1)
	sc1 <- 1000
	fmt.Println("sc1: ", sc1)

}
