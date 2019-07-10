package main

import (
	"time"
)

func main() {

	var c1 = make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 10)
		c1 <- "Hello "
	}()

	go func() {
		time.Sleep(time.Millisecond * 12)
		c1 <- "World!"
	}()

	if 1 != 2 {
		println("ok")
	}

	print(<-c1)
	print(<-c1)
}
