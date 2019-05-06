// Quick Sort in Golang
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    qsort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []float64 {

    slice := make([]float64, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Float64() - rand.Float64()
    }
    return slice
}

func qsort(a []float64) []float64 {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)

    a[pivot], a[right] = a[right], a[pivot]

    for i := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    qsort(a[:left])
    qsort(a[left+1:])

    return a
}
