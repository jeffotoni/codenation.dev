// Go in action
// @jeffotoni
// 2019-05-06

package main

import (

	//"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
	"sort"
)

type Estado struct {
	Nome     string  `json:"nome"`
	Extensao float64 `json:"extensao"`
}

func main() {
	fmt.Println(os10maioresEstadosDoBrasil())
}

func openJson(src string) (e []Estado, err error) {

	// Open our jsonFile
	//jsonFile, err := os.Open(src)
	byteValue, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("error: ", err)
		return e, err
	}
	//defer jsonFile.Close()
	//byteValue, err := ioutil.ReadAll(jsonFile)
	//var ee Estados // []struct
	err = json.Unmarshal(byteValue, &e)
	if err != nil {
		fmt.Println(err)
		return e, err
	}
	//e = ee.Estados
	//fmt.Println(e)
	//os.Exit(0)
	return e, err
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
