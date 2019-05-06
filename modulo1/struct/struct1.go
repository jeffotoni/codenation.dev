// Go in action
// @jeffotoni
// 2019-05-06

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{10, 201}
	v.X = 4
	fmt.Println(v)
}
