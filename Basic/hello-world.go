package main

import "fmt"

func main() {

	fmt.Println("Hello world!")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 && i > 90 {
			fmt.Println(i)
		}
	}

	end()
}

func foo() {
	fmt.Println("I'm foo!")
}
func end() {
	n, err := fmt.Println("Ending...", 42*2, true && false)
	fmt.Println(n, err)
}
