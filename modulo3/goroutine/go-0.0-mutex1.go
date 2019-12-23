package main

import "fmt"

// go run -race main.go
func main() {
	var x int

	for i := 0; i < 10000; i++ {
		go func() {
			x++
		}()
	}

	//x++
	fmt.Println(x)
}
