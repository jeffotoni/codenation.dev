// Go in action
// @jeffotoni
// 2019-05-06

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	RunerTimer()
	// para as goroutines funcionarem
	// fica travado sempre aguardando
	c := make(chan bool)
	<-c
}

func RunerTimer() <-chan time.Time {

	timer := time.Tick(time.Duration(50) * time.Millisecond)

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fmt.Println("\nCancelando...")
		fmt.Print("\033[?25h")
		os.Exit(0)
	}()

	fmt.Print("\033[?25l")

	//s := []rune(`|/~\`)
	//s := []rune(`-=*=`)
	s := []rune(`◐◓◑◒`)
	i := 0

	go func() {
		for {
			<-timer
			fmt.Print("\r")
			fmt.Print(string(s[i]))
			i++

			if i == len(s) {
				i = 0
			}
		}
	}()

	return timer
}
