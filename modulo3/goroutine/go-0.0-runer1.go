package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		fmt.Println("\nfim!")
		fmt.Print("\033[?25h")
		os.Exit(0)
	}()

	fmt.Print("\033[?25l")
	timer := time.Tick(time.Duration(300) * time.Millisecond)

	s := []rune(`◐◓◑◒`)
	i := 0
	for {
		<-timer
		fmt.Print("\r")
		fmt.Print(string(s[i]))
		i++
		if i == len(s) {
			i = 0
		}
	}
}
