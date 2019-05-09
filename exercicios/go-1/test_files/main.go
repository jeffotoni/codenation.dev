package main

import "fmt"

func main() {
	//@todo read dir name from stdin
	_ = jsonify("./")
}

func jsonify(dir string) error {
	return fmt.Errorf("Not implemented")
}
