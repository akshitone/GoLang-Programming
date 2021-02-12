package main

import "fmt"

func main() {
	a := 10
	b := 20
	sum := func(a int, b int) {
		fmt.Println(a + b)
	}
	sum(a, b)
}
