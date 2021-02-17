package main

import "fmt"

func main() {
	num := 20
	foo(&num)
	fmt.Println(num)
}

func foo(num *int) {
	fmt.Println(*num)
	*num = 10
	fmt.Println(*num)
}
