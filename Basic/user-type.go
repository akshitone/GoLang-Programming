package main

import "fmt"

var fo int

type one8 int

var foo one8

func main() {
	fmt.Printf("%T\n", fo)
	fmt.Printf("%T\n", foo)
	// type conversion
	foo = one8(fo)

	//sprintf
	x := 10
	y := true
	s := fmt.Sprintf("%v\t%v\n", x, y)
	fmt.Println(s)
}
