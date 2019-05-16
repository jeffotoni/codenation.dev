package main

import (
	"fmt"
	"time"
)

func funcao2(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Println(s)
	}
}

func funcao(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(600 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go funcao("goroutine")
	go funcao2("threads são do mau")
	funcao("chamada direta")

}
