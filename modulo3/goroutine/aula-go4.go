package main

import (
	"fmt"
	"sync"
	"time"
)

func f1() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 100)
		fmt.Println("f1_", i)
	}
}

func f2() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond * 200)
		fmt.Println("f2_", i)
	}
}

func main() {

	//c := make(chan int, 2)
	var wg sync.WaitGroup
	fmt.Println("start..")

	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(time.Microsecond * 200)
		//c <- 1
		fmt.Println("func1")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Microsecond * 100)
		fmt.Println("func2")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Microsecond * 50)
		fmt.Println("func3")
	}()

	// go func() {
	// 	wg.Done()
	// 	fmt.Println("func4")
	// }()

	wg.Wait()

	time.Sleep(time.Second * 5)
	fmt.Println("fim..")

	//fmt.Println("done 2..: ", <-c)
	// go f2()
	// fmt.Println("iniciar..")
	// go f1()

	// time.Sleep(time.Second)
	// // fmt.Scanln()
	// fmt.Println("done..")
	// // for {
	// // }
}
