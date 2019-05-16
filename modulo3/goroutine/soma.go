package main

import "fmt"
import "time"

// Funcao que soma todos os valores de um slice e envia o resultado por um channel
func soma(s []int, channel chan int) {

	time.Sleep(time.Second * 1)
	fmt.Println("entrei")

	soma := 0

	// Laco que soma todos os valores
	for _, v := range s {
		soma += v
	}

	// Envia o resultado da soma para o channel
	channel <- soma
}

func main() {
	// Cria um slice com alguns valores
	slice := []int{7, 2, 8, 9, 4, 0}

	// Cria o canal
	channel := make(chan int)

	// Cria duas goroutines com a funcao soma

	// Passa como parametro a primeira metade do slice e o channel
	go soma(slice[:len(slice)/2], channel)
	// Passa como parametro a segunda metado do slice e o channel
	go soma(slice[len(slice)/2:], channel)

	// Recebe os resultados dos channels
	x, y := <-channel, <-channel

	// Mostra os valores resultantes das duas goroutines e
	// o resultado total que representa a soma de todo o slice
	fmt.Println(x, "+", y, "=", x+y)
}
