package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		fmt.Println("Heyy")
	}
	i := 1
	j := 3
	for i < j {
		fmt.Println(i)
		i++
	}
	fmt.Println("-----")
	x := 1
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)

	}
}
