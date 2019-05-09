package main

import "fmt"

func main() {
	_ = parseHTML("golang.html")
}

func parseHTML(page string) error {
	return fmt.Errorf("Not implemented")
}
