package main

import (
	"fmt"
)

func main() {

	// status and bool are two different types.
	type status bool

	// MyString and string are two different types.
	type MyString string

	// Id and uint64 are two different types.
	type Id uint64

	// real and float32 are two different types.
	type real float32 // float64

	type real2 float64 // float64

	/* Some type alias declarations */

	// boolean and bool denote the same type.
	type boolean = bool

	// Text and string denote the same type.
	type Text = string
	// U8, uint8 and byte denote the same type.
	type U8 = uint8

	// char, rune and int32 denote the same type.
	type char = rune

	hex := 0xF   // the hex form (starts with a "0x" or "0X")
	octal := 017 // the octal form (starts with a "0")
	decil := 15  // the decimal form (starts without a "0")

	fmt.Println(hex, octal, decil)
	fmt.Println('\u0061')
	fmt.Println('\U00000061')

}
