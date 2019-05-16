package main

import "fmt"

type Promise struct {
	Result chan string
	Error  chan error
}

// channel
func main() {

	// var c chan int

	c1 := make(chan *Promise)
	c2 := make(chan string, 1)
	c3 := make(chan struct{})
	c4 := make(chan int, 1)
	c5 := make(map[string](chan int))

	c5["channel"] = make(chan int, 3)
	c5["channel"] <- 200
	c5["channel"] <- 100
	c5["channel"] <- 50

	// fechando channel
	defer close(c5["channel"])
	fmt.Println("c51: ", <-c5["channel"])
	fmt.Println("c52: ", <-c5["channel"])
	fmt.Println("c53: ", <-c5["channel"])

	c4 <- 200

	fmt.Println("c1: ", c1)
	fmt.Println("c2: ", c2)
	fmt.Println("c3: ", c3)
	fmt.Println("c4: ", <-c4)
	fmt.Println("c5: ", c5)

}
