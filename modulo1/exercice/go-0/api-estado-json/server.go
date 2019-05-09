// Go in action
// @jeffotoni
// 2019-05-06

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Estados struct {
	Nome     string  `json:"nome"`
	Extensao float64 `json:"extensao"`
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

func handlerEstado(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	jsons, err := json.MarshalIndent(GetEstados(), "", "\t")
	if err != nil {
		//log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(jsons))
	//json.NewEncoder(w).Encode(GetEstados())
}

func main() {
	fmt.Println("Run Server port:8080!")

	http.HandleFunc("/api/estado/extensao", handlerEstado)

	http.ListenAndServe(":8080", nil)
}
