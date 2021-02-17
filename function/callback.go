package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(numbers...))

	evenNums := even(sum, numbers...)
	fmt.Println(evenNums)
}

func sum(numbers ...int) int {
	sumNum := 0
	for _, num := range numbers {
		sumNum += num
	}
	return sumNum
}

func even(funEven func(numbS ...int) int, numbers ...int) int {
	var evenNums []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		}
	}
	return funEven(evenNums...)
}
