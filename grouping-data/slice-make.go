package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 10, 11)
	fmt.Println("slice:", slice)
	fmt.Println("length:", len(slice))
	fmt.Println("capacity:", cap(slice))

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}

	slice = append(slice, 10, 20, 40, 50)

	fmt.Println("slice:", slice)
	fmt.Println("length:", len(slice))
	fmt.Println("capacity:", cap(slice))
}
