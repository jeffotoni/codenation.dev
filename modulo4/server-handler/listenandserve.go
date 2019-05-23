package main

import (
	"io"
	"log"
	"net/http"
)

// Func ListenAndServe

// O ListenAndServe escuta o endereço de rede TCP addr e, em seguida, chama o Servir com o manipulador
// para manipular solicitações em conexões de entrada.
// As conexões aceitas são configuradas para ativar keep-alives de TCP.
// O manipulador é normalmente nulo, caso em que o DefaultServeMux é usado.
// ListenAndServe sempre retorna um erro não nulo.
func HandlerPing(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "DevopsBH, Golang for Devops ping!\n")
}

func main() {

	go func() {
		log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "DevopsBH, Golang for Devops!\n")
		}))) //ok
	}()
	go func() {
		log.Fatal(http.ListenAndServe(":8081", http.HandlerFunc(HandlerPing))) //ok
	}()

	log.Fatal(http.ListenAndServe(":8082", nil)) // ok

}
