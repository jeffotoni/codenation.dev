package main

import "fmt"

// channel
func main() {

	// receber
	// enviar
	var c = make(chan int, 2)

	c <- 110 // recebendo
	c <- 111 // recebendo

	// c <- 112 // recebendo
	// c <- 113 // recebendo
	//close(c)
	// enviando
	//fmt.Println(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c)
}
