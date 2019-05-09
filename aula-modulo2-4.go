package main

import (
	//"encoding/json"
	"errors"
	"fmt"
	"log"
	//"strings"
	//"os" //fmt.Println(os.Getenv("HOME"))
)

var s struct {
	i, j int
	fl   float64
	f    func()
	// blank _
	_ int64
}

type s2 struct {
	i, j int
	fl   float64
	f    func()
	S    *[]string
	S2   struct{ n string }
	M    map[string]interface{} //INTERFACE VAZIA..
	s3   s3
}

// private
type s3 struct {
	name string // private
	age  int
}

// public

type Estado struct {
	//Nome string  `json:"nome"`
	//Nome string  `json:"-"`
	Nome string  `json:"nome"`
	Area float64 `json:"area"`
}

type I interface {
	Methodi1()
	Methodi2()
}

type estados []Estado

type Estado2 []struct {
	Nome string
	Area float64
}

func (e Estado) Method1() (string, error) {

	if len(e.Nome) <= 0 {
		return "", errors.New("Error Nome está vazio..")
	}

	if e.Area <= 0.0 {
		return "", errors.New("Error Area está vazia..")
	}

	s := `Primeiro metodo: ` + e.Nome + ` area: ` + fmt.Sprintf("%v", e.Area)
	return s, nil
}

func validateRequest() error {
	return errors.New("Error em validateRequest!")
}

func validateWrite() error {
	return errors.New("Error em validateWrite!")
}

func getUserFromRequest() (string, error) {
	return "golang", nil
}

func GetUserData() (string, error) {
	return "", errors.New("Error em GetUserData!")
}

func main() {

	// e := &Estado{"", 123.33}
	// s, err := e.Method1().Metho2().Method3()
	// if err != nil {
	// 	log.Println("err method1:: " + err.Error())
	// 	return
	// }

	var erro error
	defer func() {
		if erro != nil {
			log.Println("Error ocorreu ", erro)
		}
	}()

	erro = validateRequest()
	if erro != nil {
		return
	}

	//var user string
	_, erro = getUserFromRequest()
	if erro != nil {
		return
	}
	// fmt.Println(user)

	_, erro = GetUserData()
	if erro != nil {
		return
	}

	// err = validateWrite()
	// if err != nil {
	//     return
	// }

	// var j int
	// for i := 0; i < 5; i++ {
	// 	j = i + 2
	// 	fmt.Println(i, "::", j)
	// }
	// fmt.Println(s)

	// Em Go, nil pode representar valores zero dos seguintes tipos de tipos:
	// There must be sufficient information for
	// compiler to deduce the type of a nil value.
	// _ = (*struct{})(nil)
	// _ = []int(nil)
	// _ = map[int]bool(nil)
	// _ = chan string(nil)
	// _ = (func())(nil)
	// _ = interface{}(nil)

	// fmt.Println("testando nil", (*struct{})(nil) == (interface{})(nil))
	// fmt.Println("testando nil", (*struct{})(nil))

	// type ApiLogin struct {
	// 	Name  string `json:"name"`
	// 	Cpf   string `json:"cpf"`
	// 	Login string `json:"login"`
	// 	Email string `json:"email"`

	// 	// anonymous
	// 	And1 struct {
	// 		City string
	// 	}
	// }

	// var s = &ApiLogin{
	// 	Name:  "John",
	// 	Cpf:   "xx.xxx.xxx-33",
	// 	Login: "john.dev",
	// 	Email: "john@gm.com",
	// 	And1:  struct{ City string }{City: "BH"},
	// }

	// fmt.Println(s)

	// var ss = []Estado{
	// 	{"Minas Gerais", 3451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// 	{"Acre", 451.33},
	// }

	// for i, v := range ss {
	// 	fmt.Println(i, ":::", v)
	// }

	// ss = append(ss, Estado{"Acre", 34.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})
	// ss = append(ss, Estado{"Minas Gerais", 3451.33})

	//fmt.Println(ss)

	// var s = []string{"Go", "Elixir", "Clojure", "Occan"}
	// s = append(s, "C", "C++", "Assembly")
	// fmt.Println(s)

	// //ARRAY estatico
	// var array = [5]int{1, 22, 3, 4, 5}
	// array[2] = 1000
	// fmt.Println(array)
	// fmt.Println(array[0])
	// fmt.Println(array[2])

	// DINAMICO
	// var slice = []int{2, 1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10}
	// fmt.Println(slice)

	// var slice = []int{2, 3, 5, 7, 11}
	// fmt.Println(slice)
	// fmt.Println(slice[2:4])

	// DINAMICO
	// var a = make([]int, 3, 4)
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3
	// //a[3] = 4
	// fmt.Println(a, "::", len(a), cap(a))

	//DINAMICO
	// var a = new([3]int)
	// a[0] = 12
	// a[1] = 2
	// fmt.Println(a)

	// // DINAMICO
	// board := [][]string{
	// 	{"_", "_", "_"},
	// 	{"_", "_", "_"},
	// 	{"_", "_", "_"},
	// }

	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {

	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	// var nil string
	// nil = "golang nil"
	// e := &Estado2{
	// 	{Nome: "Acre", Area: 234.34},
	// 	{Nome: "Minas Gerais", Area: 1234.34},
	// 	{Nome: "Amapa", Area: 334.23},
	// }
	// jsonS := []byte(`[{"nome":"go","area":2009},{"nome":"Occan","area":1983},{"nome":"C","area":1974},{"nome":"Ruby","year":1995}]`)
	// // var i interface{}
	// // var mi []map[string]interface{}
	// e := &estados{}
	// err := json.Unmarshal(jsonS, &e)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(e)

	// for _, v := range *e {
	// 	fmt.Println(v.Nome)
	// 	fmt.Println(v.Area)
	// 	//fmt.Println(v)
	// }

	// b, err := json.Marshal(e)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(b))

	// jsonbyte, err := json.Marshal(jsonS)
	// if err != nil {
	// 	//log.Println
	// 	// panic
	// 	// log.Fatal
	// 	fmt.Println(err)
	// 	return
	// }

	//fmt.Println(e)
	//fmt.Println(jsonbyte)
	//fmt.Println(jsonS)
	// fmt.Println(string(jsonS))

	// noingo := "what.. " + "ok"
	// fmt.Println(noingo)

	//s := &s3{}
	//s.name = "@jeffotoni"
	//(*s).name = "@golang"
	//fmt.Println(s.name)

	// var a int = 20
	// incp := &a
	// *incp = 10
	// *incp++
	// *incp++
	// *incp++
	// *incp++
	// *incp++
	// *incp++
	// *incp++
	// fmt.Println(*incp)

	// type Tp *int64

	// var n1 int64 = 1
	// var n2 int64 = 2

	// var p1 Tp = &n1
	// var p2 *int64 = &n2

	// //
	// p1 = p2

	// fmt.Println(&p1)
	// fmt.Println(*p1)
	// fmt.Println(*p2)

	// s := s2{
	// 	i:  12,
	// 	j:  39,
	// 	s3: s3{name: "golang", age: 10},
	// }
	// fmt.Println(s)
	// s3_1 := &s2{}
	// var s3_2 s2
	// var s3_3 = &s2{}
	// var s3_4 = new(s2)

	//s3_5 := &s2{i: 23, j: 200}
	//fmt.Println(s3_5)

	// s.i = 10
	// s.j = 100
	// s.fl = 45.93
	// s.f = func() { fmt.Println("func anonima...") }
	// fmt.Println(s)

	// s.f()
	// r1, _, err := func2()
	// if err != nil {
	// 	//log.Fatal()
	// 	// panic()
	// 	fmt.Println(err)
	// }
	// _, _, err := func2()
	// if err != nil {
	// 	//log.Fatal()
	// 	// panic()
	// 	fmt.Println(err)
	// }
	// fmt.Println(r1, "::", r2)
}

func func2() (int, string, error) {

	// Identifiers nil
	return 1, "string aqui..", nil
}
