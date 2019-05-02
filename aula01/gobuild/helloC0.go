package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("Ola: %s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from stdio in C!!!")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
