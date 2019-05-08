package main

import "fmt"

type Queue struct {
	elements chan interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		elements: make(chan interface{}, size),
	}
}

func (queue *Queue) Push(element interface{}) {
	select {
	case queue.elements <- element:
	default:
		fmt.Println("Queue full")
	}
}

func (queue *Queue) Pop() interface{} {
	select {
	case e := <-queue.elements:
		return e
	default:
		fmt.Println("Queue empty")
	}
	return nil
}

func main() {
	q := NewQueue(128)

	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Printf("Pop %d\n", q.Pop())
	fmt.Printf("Pop %d\n", q.Pop())
	fmt.Printf("Pop %d\n", q.Pop())
	fmt.Printf("Pop %d\n", q.Pop())
}
