// Struct Types
// Struct Type Tags Json

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var a struct {
	x, y int
	u    float32
	_    float32 // padding
	F    func()
}

type t struct {
	x, y int
	u    float32
	_    float32 // padding
	F    func()
	S    *[]string
	I    *[]int
	SS   struct{ n string }
	M    []map[string]interface{}
	// a    a // error no type
}

type estado struct {
	Nome string  `json:"nome"`
	Area float64 `json:"area"`
}

var E []estado

func main() {
	// --------------------------------------
	// noingo := 1 + "ok"
	a.x = 1
	a.y = 2
	a.u = 3.3
	a.F = func() { println("func struct..") }

	// a.F()
	fmt.Println(a)

	//var tt t
	//var tt = &t{}
	var tt = new(t)

	tt.x = 1
	tt.y = 2
	tt.u = 1.3
	// tt._ = 2.3 error
	// tt.a.x = 1 // error
	tt.S = &[]string{"Occam", "Limbo", "Modula", "Go"}
	tt.I = &[]int{1, 2, 3, 4, 5}

	// not type
	tt.SS.n = "jef"
	// var ttss tt.SS error

	fmt.Println(tt)
	fmt.Println(tt.S)
	fmt.Println(tt.I)

	// adicionado varios
	E = append(E, estado{"Acre", 1323.22})
	E = append(E, estado{"Minas Gerais", 2023.22})
	fmt.Println(E)

	var E1 = estado{"Amapa", 322.3}
	var E2 = estado{}

	// adicionado varios
	copyx(&E1, &E2)
	fmt.Println(E2)

	//
	var eb = make([]estado, len(E))
	copy(eb, E)

	fmt.Println("eb:: ", eb, " len and cap: ", len(eb), " :: ", cap(eb))

	////////// Decode
	// "Parents":["Arthur","Francisco"]
	jsonData := `[{"Name":"Jef","Age":36},{"Name":"Go","Age":10}]`
	var m []map[string]interface{}

	// uma linha decoder
	json.NewDecoder(strings.NewReader(jsonData)).Decode(&m)
	fmt.Println(&m)
	tt.M = m
	fmt.Println(tt.M)

	jsonByte := []byte(`[{"Name":"Occan","Year":1983},{"Name":"Go","Year":2009}]`)
	var mm []map[string]interface{}

	//////// outra forma de fazer decode.. Unmarshal
	if err := json.Unmarshal(jsonByte, &mm); err != nil {
		fmt.Println("json error: ", err)
		return
	}

	fmt.Println(&mm)
	fmt.Println(mm[0]["Name"])
	fmt.Println(mm[0]["Year"])

	// ---------STRUCT Unmarshal -----------------------------

	jsonEstado := []byte(`[
		{"Nome":"Acre","Area":345.66},
		{"Nome":"Amapá","Area":3445.26}
		]`)

	var ej []estado
	err := json.Unmarshal(jsonEstado, &ej)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ej)
	////////////////////////

	// [
	// 	{"name":"jeff"},
	// 	{"name":"Go"}
	// ]

	// type Dinamic []struct {
	// 	Name string `json:"name"`
	// }

	// [
	// 	{"name":"jeff", "lang":[{"name":"Go", "name":"Occam", "name":"Go"}]},
	// 	{"name":"Ramir","lang":[{"name":"C++", "name":"Limbo", "name":"C"}]}
	// ]

	// type Dinamic []struct {
	// 	Name string `json:"name"`
	// 	Lang []struct {
	// 		Name string `json:"name"`
	// 	} `json:"lang,omitempty"`
	// }
}

func copyx(orig *estado, dst *estado) {

	*dst = *orig
}

// String Types
// Slice Types
// Pointer
// Map Types
// Map Literals Continued
// Blank Identifier
// Identifiers nil
// Interface as type
