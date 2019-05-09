package main

import (
	"fmt"
	"time"
)

func main() {

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
