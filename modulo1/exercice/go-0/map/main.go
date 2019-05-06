package main

import (
	//"fmt"
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

func GetEstados() map[float64]string {

	return map[float64]string{
		152.5:  "Acre",
		357.1:  "Mato Grosso do Sul",
		340.0:  "Goiás",
		224.2:  "Roraima",
		332.0:  "Maranhão",
		277.6:  "Tocantins",
		251.5:  "Piauí",
		248.2:  "São Paulo",
		148.8:  "Ceará",
		199.3:  "Paraná",
		237.5:  "Rondônia",
		1570.7: "Amazonas",
		5.8:    "Distrito Federal",
		142.8:  "Amapá",
		281.7:  "Rio Grande do Sul",
		98.3:   "Pernambuco",
		95.3:   "Santa Catarina",
		56.4:   "Paraíba",
		52.8:   "Rio Grande do Norte",
		46.0:   "Espírito Santo",
		43.6:   "Rio de Janeiro",
		564.6:  "Bahia",
		587.5:  "Minas Gerais",
		1247.6: "Pará",
		27.7:   "Alagoas",
		21.9:   "Sergipe",
		903.3:  "Mato Grosso",
	}
}

func os10maioresEstadosDoBrasil() ([]string, error) {

	var data []string
	m := GetEstados()

	//keys := make([]string, 0, len(m))
	keys := make([]float64, 0, len(m))
	for size := range m {
		keys = append(keys, size)
	}
	sort.Float64s(keys)
	for i := 0; i < 10; i++ {
		nameEstado := m[keys[26-i]]
		data = append(data, nameEstado)
		//fmt.Println(nameEstado)
	}

	return data, nil
}
