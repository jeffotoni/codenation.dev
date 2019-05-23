// Ao usar canais como parâmetros de função, você pode
// especifica se um canal serve apenas para enviar ou receber
// valores.

package main

// Esta função "ping" só aceita um canal para envio
// valores. Seria um erro em tempo de compilação tentar
// receber neste canal.
// unidirecional - somente read
// receive from send-only type chan<- string
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Esta função "pong" aceita um canal para receber e um para enviar
// unidirecional -  write e read
// send to receive-only type <-chan string   // write
// receive from send-only type chan<- string // read
func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {

	pings := make(chan string, 1) // channel bi-direcional
	pongs := make(chan string, 1) // channel bi-direcional

	go ping(pings, "Enviando mensagem @jeffotoni..")
	pong(pings, pongs)
	println(<-pongs)
}
