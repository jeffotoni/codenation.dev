package main

import "fmt"

func main() {
	//Todas as perguntas são referentes ao arquivo data.csv
}

//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q1() (int, error) {
	return 0, fmt.Errorf("Not implemented")
}

//Quantos clubes (coluna `club`) diferentes existem no arquivo?
func q2() (int, error) {
	return 0, fmt.Errorf("Not implemented")
}

//Liste o primeiro nome dos 20 primeiros jogadores de acordo com a coluna `full_name`.
func q3() ([]string, error) {
	return []string{}, fmt.Errorf("Not implemented")
}

//Quem são os top 10 jogadores que ganham mais dinheiro (utilize as colunas `full_name` e `eur_wage`)?
func q4() ([]string, error) {
	return []string{}, fmt.Errorf("Not implemented")
}

//Quem são os 10 jogadores mais velhos (use como critério de desempate o campo `eur_wage`)?
func q5() ([]string, error) {
	return []string{}, fmt.Errorf("Not implemented")
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
	idades := make(map[int]int)
	return idades, fmt.Errorf("Not implemented")
}
