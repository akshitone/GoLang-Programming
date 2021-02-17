package main

func main() {
	a := 10
	b := 30
	sum := func(a int, b int) int {
		return a + b
	}
	println(sum(a, b))

	newFun := func() {
		for i := 0; i < 3; i++ {
			println(i)
		}
	}

	newFun()
}
