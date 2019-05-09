package main

import (
	//"fmt"
	//"log"
	"sort"
)

type Estados struct {
	N string
	E float64
}

const (
	QuantMaior = 10
)

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// more
	More(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type ByExt []Estados

func (a ByExt) Len() int           { return len(a) }
func (a ByExt) Less(i, j int) bool { return a[i].E < a[j].E }
func (a ByExt) More(i, j int) bool { return a[i].E > a[j].E }
func (a ByExt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	// e, err := os10maioresEstadosDoBrasil()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// for _, v := range e {
	// 	println(v)
	// }
}

func os10maioresEstadosDoBrasil() ([]string, error) {

	var data []string

	//estado := []struct {
	// 	Nome     string
	// 	Extensao float64
	// }

	estado := []Estados{
		{"Acre", 164123.040},
		{"Alagoas", 27778.506},
		{"Amapá", 142828.521},
		{"Amazonas", 1559159.148},
		{"Bahia Bahia", 564733.177},
		{"Distrito Federal", 5779.999},
		{"Espírito Santo", 46095.583},
		{"Maranhão", 331937.450},
		{"Mato Grosso", 357145.532},
		{"Mato Grosso do Sul", 357145.532},
		{"Minas Gerais", 586522.122},
	}

	sort.Sort(ByExt(estado))

	// sort.SliceStable(estado, func(i, j int) bool {
	// 	return estado[i].E > estado[j].E
	// })

	for i := 0; i < QuantMaior; i++ {
		//data[i] = estado[i].N
		data = append(data, estado[i].N)
	}

	return data, nil
}
