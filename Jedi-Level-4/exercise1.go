package main

import (
	"fmt"
)

func main() {

	arr := [5]int{3, 6, 7, 8, 4}
	fmt.Println(arr)

	for index, val := range arr {
		fmt.Println(index, val)
	}

	fmt.Printf("%T\n", arr)

}
