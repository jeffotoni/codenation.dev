// Go in Action
// @jeffotoni
// 06/05/2019

package main

import (
	"fmt"
	"sort"
)

type Estados struct {
	N string
	E float64
}

const (
	QuantMaior = 10
)

func main() {

	e, err := os10maioresEstadosDoBrasil()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range e {
		println(v)
	}
}

func GetEstados() []Estados {

	return []Estados{
		{"Acre", 152.5},
		{"Mato Grosso do Sul", 357.1},
		{"Goiás", 340.0},
		{"Roraima", 224.2},
		{"Maranhão", 332.0},
		{"Tocantins", 277.6},
		{"Piauí", 251.5},
		{"São Paulo", 248.2},
		{"Ceará", 148.8},
		{"Paraná", 199.3},
		{"Rondônia", 237.5},
		{"Amazonas", 1570.7},
		{"Distrito Federal", 5.8},
		{"Amapá", 142.8},
		{"Rio Grande do Sul", 281.7},
		{"Pernambuco", 98.3},
		{"Santa Catarina", 95.3},
		{"Paraíba", 56.4},
		{"Rio Grande do Norte", 52.8},
		{"Espírito Santo", 46.0},
		{"Rio de Janeiro", 43.6},
		{"Bahia", 564.6},
		{"Minas Gerais", 587.5},
		{"Pará", 1247.6},
		{"Alagoas", 27.7},
		{"Sergipe", 21.9},
		{"Mato Grosso", 903.3},
	}
}

func os10maioresEstadosDoBrasil() ([]string, error) {

	var data []string
	estado := GetEstados()
	sort.SliceStable(estado, func(i, j int) bool {
		return estado[i].E > estado[j].E
	})

	for i := 0; i < QuantMaior; i++ {
		//data[i] = estado[i].N
		data = append(data, estado[i].N)
	}

	return data, nil
}
