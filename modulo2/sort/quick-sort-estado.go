package main

import (
	"fmt"
	"math/rand"
)

var e = map[float64]string{
	152.5: "Acre",
		357.1:"Mato Grosso do Sul",
		340.0:"Goiás"
		332.0:"Maranhão",
		277.6:"Tocantins",
		251.5:"Piauí",
		"São Paulo":           248.2,
		"Ceará":               148.8,
		"Roraima":             224.2,
		"Paraná":              199.3,
		"Rondônia":            237.5,
		"Amazonas":            1570.7,
		"Distrito Federal":    5.8,
		"Amapá":               142.8,
		"Rio Grande do Sul":   281.7,
		"Pernambuco":          98.3,
		"Santa Catarina":      95.3,
		"Paraíba":             56.4,
		"Rio Grande do Norte": 52.8,
		"Espírito Santo":      46.0,
		"Rio de Janeiro":      43.6,
		"Bahia":               564.6,
		"Minas Gerais":        587.5,
		"Pará":                1247.6,
		"Alagoas":             27.7,
		"Sergipe":             21.9,
		"Mato Grosso":         903.3,
}

var t = []float64{
	164123.040,
	27778.506,
	142828.521,
	1559159.148,
	564733.177,
	5779.999,
	46095.583,
	331937.450,
	903366.192,
	357145.532,
	22586522.122,
}

func main() {

	var nameEstado [10]string
	em := qsort(t)
	fmt.Println(em)
	for i := 0; i < 10; i++ {
		//nameEstado = append(nameEstado, e[em[i]])
		pos := 9 - i
		nameEstado[pos] = e[em[i]]
	}

	fmt.Println(nameEstado)
}

func qsort(a []float64) []float64 {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
