package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(500)
	}
	fmt.Println(arr)
}
