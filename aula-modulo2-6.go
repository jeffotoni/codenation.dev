package main

import (
	"fmt"
)

func contador() func() int {
	var i int = 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	next := contador()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	// zera... começa
	next = contador()
	fmt.Println(next())
	fmt.Println(next())

	// var m = make(map[string]string)
	// m["GO0"] = "VIDA LONGA!"
	// m["GO1"] = "VIDA LONGA!"
	// m["GO2"] = "VIDA LONGA!"
	// m["GO3"] = "VIDA LONGA!"
	// m["GO4"] = "VIDA LONGA!"
	// m["GO5"] = "VIDA LONGA!"
	// m["JS"] = "VIDA LONGA!"

	// fmt.Println("ok", m)

}
