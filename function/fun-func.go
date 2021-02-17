package main

import "fmt"

func main() {
	// anonymous function
	x := func() int {
		return 451
	}()
	fmt.Println(x)

	// function return function
	fmt.Println(foo()())

	getFoo := foo()
	fmt.Printf("%T\n", getFoo)

	fmt.Println(getFoo())

	funFunc := getFoo()
	fmt.Println(funFunc)
}

func foo() func() string {
	return func() string {
		return "451 string"
	}
}
