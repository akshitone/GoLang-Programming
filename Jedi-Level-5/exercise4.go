package main

import "fmt"

func main() {
	heyy := struct {
		message string
		to      string
	}{
		"hello", "akshit",
	}
	fmt.Println(heyy)
}
