package main

import "fmt"

var a string

type f1 func()

//func f1() {
//	fmt.Println("loves")
//}

type Loves struct {
	nick string
}

type Tel struct {
	fone  string
	loves Loves
}

type User struct {
	Login string `json:login`
	Email string
	/**Tel   struct {
		fone  string
		loves struct {
			nick string
		}
	}*/

	T Tel

	F f1
}

func main() {

	fmt.Println("vim-go", a)

	var l Loves
	fmt.Printf("%#v%", l)

	var l1 = Loves{}
	fmt.Println(l1)

	var l2 = Loves{nick: "amanda"}
	fmt.Println(l2)

	var l3 = Loves{"amanda"}
	fmt.Println(l3)

	//l := Loves{}
	//l := &Loves{}
	//l := New(Loves)
	//println(l)

}
