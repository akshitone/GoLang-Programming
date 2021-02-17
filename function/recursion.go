package main

import "fmt"

func main() {
	ans := factorial(5)
	fmt.Println(ans)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
