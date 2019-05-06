package main

import "fmt"

// type T int64
// type U W
// type W int64

// func main() {
// 	var n T = 1
// 	var m U = 2

// 	var p *T = &n
// 	var q *U = &m

// 	p = (*T)(q)

// 	fmt.Println(*p)
// }

type T *int64
type U *W
type W int64

func main() {
	var n int64 = 1
	var m W = 2
	var p T = &n
	var q U = &m
	p = T(q)
	fmt.Println(*p)
}
