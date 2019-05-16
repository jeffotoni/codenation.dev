package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	go func() { fmt.Println("oi sou uma goroutine1!") }()
	go func() { fmt.Println("oi sou uma goroutine2!") }()
	fmt.Println("Hello, playground")
	//time.Sleep(time.Second)
	<-done

}
