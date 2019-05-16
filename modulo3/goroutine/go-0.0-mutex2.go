package main

import "sync"

func main() {
	var x int
	var m sync.Mutex

	go func() {
		m.Lock()
		x++
		m.Unlock()
	}()

	m.Lock()
	x++
	println(x)
	m.Unlock()
}
