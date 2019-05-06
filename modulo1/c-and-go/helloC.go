// Go in action
// @jeffotoni
// 2019-05-06

package main

//#include<stdio.h>
//void inC() {
//    printf("Eu sou um Codigo C em Go!\n");
//}
import "C"

func main() {
	C.inC()
}
