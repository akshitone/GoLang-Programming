package main

func main() {
	defer foo()
	bar()
}

func foo() {
	println("first")
}

func bar() {
	println("second")
}
