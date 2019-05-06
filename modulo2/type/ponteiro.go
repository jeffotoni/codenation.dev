package main

import "fmt"

type T *int64

func main() {
	var n int64 = 1
	var m int64 = 2

	var p T = &n
	var q *int64 = &m

	p = q
	fmt.Println(*p)
}
