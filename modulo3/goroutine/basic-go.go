package main

func main() {

	c1 := make(chan int)
	go func() { c1 <- 3 }()
	println(<-c1)

	c2 := make(chan int, 1)
	c2 <- 1
	println(<-c2)
}
