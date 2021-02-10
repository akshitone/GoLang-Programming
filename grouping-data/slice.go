package main

import "fmt"

func main() {
	sli := []int{}
	slice := []int{2, 3, 6, 7, 8, 10}

	slice = append(slice, 15, 20)
	sli = append(slice, 1, 2)

	for i, val := range slice {
		fmt.Println(i, "-->", val)
	}

	fmt.Println(slice[1:3])
	fmt.Println(sli)

	sli = append(sli[:3], sli[5:]...)
	fmt.Println(sli)

}
