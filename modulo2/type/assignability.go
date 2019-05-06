package main

import "fmt"

func f(n [2]int) {
	fmt.Println(n)
}

type T [2]int

func main() {
	var v T
	v[0] = 1
	v[1] = 2
	f(v)
}
