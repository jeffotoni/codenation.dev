// Go in action
// @jeffotoni
// 2019-05-06

package main

import (

	//"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"strconv"
	//"os"
	"sort"
)

type Estado struct {
	Nome     string  `json:"nome"`
	Extensao float64 `json:"extensao"`
}

func main() {

	fmt.Println(os10maioresEstadosDoBrasil())
	//openJson("./dado.json")
}

func openJson(src string) (e []Estado, err error) {

	byteValue, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("error: ", err)
		return e, err
	}

	//var f interface{}
	var fmpas []map[string]interface{}
	err = json.Unmarshal(byteValue, &fmpas)
	if err != nil {
		fmt.Println(err)
		return e, err
	}

	//fmt.Println(fmpas)
	//fmt.Println(fmpas[0]["extensao"])

	for _, v := range fmpas {
		var el Estado
		//e.Nome = fmt.Sprintf("%s", v["nome"])
		el.Nome = v["nome"].(string)
		el.Extensao = v["extensao"].(float64)
		e = append(e, el)
	}

	//fmt.Println(estados)
	// ff := f.([]interface{})
	// mapx := ff[0].(interface{})
	// fmt.Println(mapx.)
	return
}

func sortStructDesc(e []Estado) {
	sort.Slice(e, func(i, j int) bool {
		return e[i].Extensao > e[j].Extensao
	})
}

func os10maioresEstadosDoBrasil() ([]string, error) {
	var data []string
	e, err := openJson("./dado.json")
	if err != nil {
		return data, err
	}

	//fmt.Println(e)
	//return data, nil

	sortStructDesc(e)
	for i := 0; i < 10; i++ {
		data = append(data, e[i].Nome)
	}
	return data, nil
}
