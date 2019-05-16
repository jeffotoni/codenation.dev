package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "service 2"
}

func main() {
	fmt.Println("start", time.Since(start))

	c1 := make(chan string)
	c2 := make(chan string)

	go service1(c1)
	go service2(c2)

	select {
	case res := <-c1:
		fmt.Println("Responda do service 1", res, time.Since(start))
	case res := <-c2:
		fmt.Println("Responda do service 2", res, time.Since(start))
	}

	fmt.Println("fim:: ", time.Since(start))
}
