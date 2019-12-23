package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	fmt.Println(hash("HelloWorld"))
	fmt.Println(hash("HelloWorld."))
}
