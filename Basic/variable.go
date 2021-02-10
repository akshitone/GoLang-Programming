package main

import "fmt"

var y = 100
var z int
var str string = `"Do what makes you happy"
	-Akshit Mithaiwala`

func main() {
	x := 10
	println(x)

	x = 20
	println(x)

	name := "Akshit Mithaiwala"
	fmt.Println(name)

	fmt.Println(34 + y)
	y += 20
	foo()
	fmt.Println(z)
	fmt.Println(str)
}

func foo() {
	fmt.Println(12 + y)
}
