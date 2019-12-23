package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	var b strings.Builder
	var wait sync.WaitGroup
	for i := 0; i < 20000; i++ {
		wait.Add(1)

		go func() {
			b.WriteString("1\n")
			fmt.Println(b.String())
			time.Sleep(time.Second * 1)
			wait.Done()
		}()
	}

	for i := 0; i < 10000; i++ {
		wait.Add(1)
		go func() {
			b.WriteString("2\n")
			fmt.Println(b.String())
			time.Sleep(time.Second * 1)
			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(len(b.String()))
	//fmt.Println(b.String())
}
