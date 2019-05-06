// Go in action
// @jeffotoni
// 2019-05-06

package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
    printf("Isto aqui é um code C: %s", s);
}
*/
import "C"

import "unsafe"

func callC() {
	cs := C.CString("Isto aqui é Go code\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func main() {
	callC()
}
