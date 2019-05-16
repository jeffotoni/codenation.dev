package main

import (
	"fmt"
	//"runtime"
)

func main() {
	fmt.Println("This will happen first")

	go func() {
		fmt.Println("Func one")
	}()

	go func() {
		fmt.Println("Func two")
	}()

	fmt.Println("Func 3")
	// fmt.Scanln()
	// select {}
	// for {
	// 	runtime.Gosched()
	// }
	// for {
	// }

	fmt.Println("done")
}
