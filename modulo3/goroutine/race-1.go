package main

import "fmt"

// go run -race race-1.go
func main() {
	// race()
	// sharingIsCaring()
}

func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		close(wait)
	}()
	n++ // conflicting access
	<-wait

	fmt.Println(n) // Output: <unspecified>
}

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another.
	n++
	fmt.Println(n) // Output: 2
}
