package main

import (
	"fmt"
	//"sync"
	"time"
)

func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}

func main() {

	wait := Publish("Publica isto em 2 segundos.", 3*time.Second)

	<-wait // Aguarde enquanto o texto esta sendo publicado
}
