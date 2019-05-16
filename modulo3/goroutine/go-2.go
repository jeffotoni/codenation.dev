package main

import "fmt"
import "time"

func goroutine(c chan string) {
	fmt.Println("Eu sou um canal: " + <-c + "!")
}

func main() {

	fmt.Println("start goroutine!")
	c := make(chan string)
	go goroutine(c)
	c <- "jeffotoni"

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

	go func() {
		var i int
		for {
			time.Sleep(time.Second)
			fmt.Println("i: ", i)
			i++
		}
	}()

	go func() {
		var i int
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("j: ", i)
			i++
		}
	}()

	time.Sleep(time.Second * 5)
}
