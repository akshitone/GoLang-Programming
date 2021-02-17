package main

func main() {
	a := 405
	b := &a
	println(a)
	println(b)

	*b = 34
	println(a)
}
