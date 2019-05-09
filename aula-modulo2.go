package main

import (
	"fmt"
)

// public
const PATH_API = "localhost:8080"
const PATHAPI = "localhost:8081"
const PathApi = "localhost:8082"

func Hello() {
	fmt.Println(PATH_API)
}

// public
func P() {
	// var breakx string = "ola"
	fmt.Println("P")
}

// private
func m(s string) {
	// var breakx string = "ola"
	fmt.Println("M: " + s)
}

func main() {

	var S string = "Curso Go!!"

	m(S)
	P()
	Hello()
}
