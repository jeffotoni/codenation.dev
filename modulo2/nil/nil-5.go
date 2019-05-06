package main

import (
	"fmt"
	"reflect"
)

var x = struct{ a, _, c int }{1, 2, 3}

var x1 = struct{ _ int }{1}

func main() {
	var x2 = struct{ _ int }{1}

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x1 == x2)
	var y = struct{ a, _, c int }{1, 2, 3}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(reflect.ValueOf(x).Field(1).Int())
	fmt.Println(reflect.ValueOf(y).Field(1).Int())
	fmt.Println(x == y)
}
