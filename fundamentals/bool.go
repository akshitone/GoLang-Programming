package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 10
	b := 20
	fmt.Println(a == b)
	println(runtime.GOOS)
	println(runtime.GOARCH)

	s := "Akshit"
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
		// fmt.Printf("%#U ", s[i])
	}
	fmt.Println()

	for i, v := range s {
		fmt.Printf("at index %d we have %b\n", i, v)
	}
}
