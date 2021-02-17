package main

import "fmt"

func main() {
	number := foo()
	fmt.Println(number)
	rollno, name := bar()
	fmt.Println(rollno, name)
}

func foo() int {
	return 450
}

func bar() (int, string) {
	return 235, "akshit"
}
