package main

import "fmt"

func main() {
	n := false
	switch n {
	case false:
		fmt.Println("It's false")
	case (2 == 4):
		fmt.Println("It's false too")
	case false, (2 == 2):
		fmt.Println("It's true")
		fallthrough
	case (4 == 3):
		fmt.Println("It's true too")
	default:
		fmt.Println("Ending...")
	}
}
