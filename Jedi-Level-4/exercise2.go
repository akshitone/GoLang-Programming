package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for index, value := range slice {
		fmt.Println(index, value)
	}
	fmt.Printf("%T\n", slice)
}
