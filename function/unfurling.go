package main

import "fmt"

func main() {
	numbers := []int{4, 5, 6, 7, 8}
	names := []string{"akshit", "viral"}
	userName(names...)
	res := sum(numbers...)
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

func userName(names ...string) {
	fmt.Println(names)
}
