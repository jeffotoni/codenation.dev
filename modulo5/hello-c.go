package main

/*
#include <stdio.h>
#include <stdlib.h>

void GoPrint(char* s) {
	printf("%s\n", s);
}

*/
import "C"

func main() {
	C.GoPrint(C.CString("Hello world\n"))
}
