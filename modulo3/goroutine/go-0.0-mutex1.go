package main

import "fmt"

// go run -race main.go
func main() {
	var x int

	go func() {
		x++
	}()

	x++
	fmt.Println(x)
}
