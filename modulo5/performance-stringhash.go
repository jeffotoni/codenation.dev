package main

import (
    "fmt"
    "hash/fnv"
    "runtime"
    "strconv"
    "time"
)

const (
    numElements = 10000000
)

var foo = map[uint32]int{}

func timeGC() {
    t := time.Now()
    runtime.GC()
    fmt.Printf("gc took: %s\n", time.Since(t))
}

func hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}

func main() {
    for i := 0; i < numElements; i++ {
        foo[hash(strconv.Itoa(i))] = i
    }

    for {
        timeGC()
        time.Sleep(1 * time.Second)
    }

    fmt.Println(foo)
}
