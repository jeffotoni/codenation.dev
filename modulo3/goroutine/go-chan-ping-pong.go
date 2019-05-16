// Ao usar canais como parâmetros de função, você pode
// especifica se um canal serve apenas para enviar ou receber
// valores. Esta especificidade aumenta a segurança do tipo
// o programa.

package main

import "fmt"

// unidirecional - somente write
// send to receive-only type <-chan int
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// unidirecional - somente read
// receive from send-only type chan<- int
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Enviando mensagem @jeffotoni..")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
