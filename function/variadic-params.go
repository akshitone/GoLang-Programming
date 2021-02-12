package main

import "fmt"

func main() {
	res := sum(4, 5, 6, 7, 8)
	fmt.Println(res)
}

func sum(lst ...int) int {
	fmt.Println(lst)
	sum := 0
	for i := 0; i < len(lst); i++ {
		sum += lst[i]
	}
	return sum
}
