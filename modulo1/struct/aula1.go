// aula struct

package main

import (
	"fmt"
)

// type SpaUi interface {
// 	Sum()
// 	String()
// 	Auth()
// }

type S = string

type Login struct { // minusculo private
	Id    int    `json:"id"`    // maiculo Public
	Email string `json:"email"` // minusculo private
}

type Pessoal struct { // minusculo private
	Id   int    `json:"id"`   // maiculo Public
	Nome string `json:"nome"` // minusculo private
	L    *Login
}

func main() {
	var s S
	s = "curso codenation..."
	var l = &Login{1003, `jeff.otoni@email.com` + s}
	fmt.Println("struct:: ", l)

	// var l = &login{}
	// l := &login{}

}
