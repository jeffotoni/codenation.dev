package main

import "fmt"

func main() {

	// receber como enviar...
	// bidirecional e buffer
	c := make(chan string, 3)
	// c := make(chan string)
	c <- "ola 1"
	c <- "ola 2"
	c <- "ola 3"

	close(c)

	for ck := range c {
		fmt.Println("c: ", ck)
	}

	v, ok := <-c
	fmt.Println("v: ", v)
	fmt.Println("ok: ", ok)

	c1 := <-c
	c2 := <-c
	c3 := <-c

	fmt.Println(": ", c1)
	fmt.Println(": ", c2)
	fmt.Println(": ", c3)
}
