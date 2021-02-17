package main

func main() {
	a := 10
	b := 30
	sum := func(a int, b int) int {
		return a + b
	}(a, b)
	println(sum)
}
