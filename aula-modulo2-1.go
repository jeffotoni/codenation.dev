package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan int
	var cc = make(chan int)
	var ccc = make(chan int, 10)
	fmt.Println(c)
	fmt.Println(cc)
	fmt.Println(ccc)

	//
	go func1()
	go func2()
	func3()

	time.Sleep(time.Second * 7)
}

func func1() {
	time.Sleep(time.Second * 3)
	fmt.Println("func 1")
}

func func2() {
	time.Sleep(time.Second * 2)
	fmt.Println("func 2")

}

func func3() {
	time.Sleep(time.Second * 1)
	fmt.Println("func 3")
}
