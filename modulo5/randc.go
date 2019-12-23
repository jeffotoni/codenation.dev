package main

import (
	"fmt"
)

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}

func main() {

	fmt.Println(Random())
}
