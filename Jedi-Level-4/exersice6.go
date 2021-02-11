package main

import "fmt"

func main() {
	state := make([]string, 50, 100)
	state = []string{"al", "ny", "la"}
	fmt.Println(state)

	for i := 0; i < len(state); i++ {
		fmt.Println(i, state[i])
	}
}
