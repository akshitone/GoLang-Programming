package main

import "fmt"

func main() {
	users := [][]string{}
	users = [][]string{
		[]string{"235", "akshit", "mithaiwala", "akshitone8@gmail.com"},
		[]string{"243", "viral", "shastri", "vshastri19@gmail.com"},
	}
	fmt.Println(users)

	for i, user := range users {
		fmt.Println(i, user)
		for j, value := range user {
			fmt.Println("\t", j, value)
		}
	}
}
