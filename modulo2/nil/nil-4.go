package main

import "fmt"

func main() {

	fmt.Println(struct{ _ int }{1} == struct{ _ int }{2})
}
