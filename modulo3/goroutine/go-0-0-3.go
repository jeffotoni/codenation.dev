package main

import "fmt"

func main() {

	ch := make(chan string)
	go func() {
		ch <- "goroutine here!"
		defer close(ch)
	}()

	//fmt.Println(<-ch) // Print "Hello!".
	//fmt.Println(<-ch) // Print the zero value "" without blocking.
	//fmt.Println(<-ch) // Once again print "".

	//v, ok := <-ch // v is "", ok is false.
	//fmt.Println("v:", v, " ok: ", ok)

	// Receive values from ch until closed.
	for v := range ch {
		fmt.Println(v) // Will not be executed.
	}
}
