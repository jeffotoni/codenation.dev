// Go in action
// @jeffotoni
// 2019-05-06

package main

/*
#include <stdio.h>
#include <stdlib.h>

int Sum(int a, int b) {
    return a+b;
}

int Mult(int a, int b) {
    return a*b;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Sum: ", C.Sum(2, 4))
	fmt.Println("Mult:", C.Mult(2, 4))
}
