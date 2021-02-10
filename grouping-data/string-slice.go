package main

import "fmt"

func main() {
	akshit := []string{"235", "akshit", "mithaiwala"}
	viral := []string{"243", "viral", "shastri"}

	fmt.Println(akshit)
	fmt.Println(viral)

	student := [][]string{akshit, viral}
	fmt.Println(student)
	fmt.Println(student[0][1])
}
