package main

import (
	"fmt"
	//"sync"
	//"time"
)

func main() {

	// unbuffered channel of ints
	ic := make(chan int)

	go func() {
		fmt.Println("func 1")
		ic <- 1
	}()
	go func() {
		fmt.Println("func 2")
		ic <- 2
	}()

	// <-ic
	// <-ic
	close(ic)

	for c := range ic {
		fmt.Println(c)
	}

	fmt.Println("feito")

}
