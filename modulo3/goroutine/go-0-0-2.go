package main

import (
	"fmt"
	//"sync"
	"time"
)

func main() {

	// buffered channel with room for 10 strings
	sc := make(chan int, 1)

	go func() {
		//fmt.Println("func 1")
		// envia
		fmt.Println("envia: 1")
		sc <- 1
	}()
	go func() {
		//fmt.Println("func 2")
		// sc <- 2
		n := <-sc
		fmt.Println("recebe:", n)
	}()

	//<-sc
	//<-sc
	// defer close(sc)
	// go func() {
	// 	for c := range sc {
	// 		fmt.Println(c)
	// 	}
	// }()

	fmt.Println("feito")
	time.Sleep(time.Second * 2)

}
