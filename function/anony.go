package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello madafaka!")
	}()

	a := 10
	b := 20
	func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
}
