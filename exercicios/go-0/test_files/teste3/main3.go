package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {

	e, err := os10maioresEstadosDoBrasil()
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range e {
		println(v)
	}
}

func os10maioresEstadosDoBrasil() ([]string, error) {

	var data []string

	// m := map[string]float64{
	// 	"Acre":               164123.040,
	// 	"Alagoas":            27778.506,
	// 	"Amapá":              142828.521,
	// 	"Amazonas":           1559159.148,
	// 	"Bahia Bahia":        564733.177,
	// 	"Distrito Federal":   5779.999,
	// 	"Espírito Santo":     46095.583,
	// 	"Maranhão":           331937.450,
	// 	"Mato Grosso":        357145.532,
	// 	"Mato Grosso do Sul": 357145.532,
	// 	"Minas Gerais":       586522.122,
	// }

	m := map[float64]string{
		164123.040:   "Acre",
		27778.506:    "Alagoas",
		142828.521:   "Amapá",
		1559159.148:  "Amazonas",
		564733.177:   "Bahia Bahia",
		5779.999:     "Distrito Federal",
		46095.583:    "Espírito Santo",
		331937.450:   "Maranhão",
		903366.192:   "Mato Grosso",
		357145.532:   "Mato Grosso do Sul",
		22586522.122: "Minas Gerais",
	}

	//keys := make([]string, 0, len(m))
	keys := make([]float64, 0, len(m))
	//var i int
	for size := range m {
		//fmt.Println(k)
		//keys = append(keys, k)
		keys = append(keys, size)
		//keys[i] = size
		//i++
	}
	//sort.Strings(keys)
	sort.Float64sAreSorted(keys)

	//fmt.Println(keys)
	for i := 0; i < 10; i++ {
		fmt.Println(m[keys[i]])
	}

	// for x, k := range keys {
	// 	fmt.Println(x, k)
	// 	//fmt.Println(k, m[k])
	// 	//data = append(data, k)
	// }

	return data, nil
}
