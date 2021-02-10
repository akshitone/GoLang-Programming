package main

import "fmt"

const (
	pi float32 = 3.14
)

const (
	a = iota
	c = iota
	b = iota
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	//bitwise swifting
	fmt.Println(c << 5)
	fmt.Println(pi)
	fmt.Printf("%T\n", pi)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
