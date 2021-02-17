package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sum := sumWithVariadic(numbers...)
	fmt.Println(sum)
	sum = sumWithInt(numbers)
	fmt.Println(sum)
}

func sumWithVariadic(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func sumWithInt(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
