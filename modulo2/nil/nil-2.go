package main

import "fmt"
import "os"

func Normal() error {
	return nil
}

func Bizarro() *os.PathError {
	return nil
}

func main() {

	fmt.Println(string(65))
	err := Normal()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // true

	err = Bizarro()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
}
