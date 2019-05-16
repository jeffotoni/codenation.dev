package main

import "fmt"

func f(rc <-chan string) {
	fmt.Println("somente recebe o valor rc:  " + <-rc)
}

func f2(sc chan<- string) {
	sc <- "somente envia o valor sc.."
}

func main() {

	rc := make(<-chan int) // somente recebe
	sc := make(chan<- int) // somente envia

	fmt.Printf("Data type of roc is `%T`\n", rc)
	fmt.Printf("Data type of soc is `%T\n", sc)

	fmt.Println("inicia goroutine")
	c := make(chan string)
	c1 := make(chan string)

	go f(c)
	c <- "Codenation"

	go f2(c1)
	fmt.Println(<-c1)

	fmt.Println("stopped")
}
