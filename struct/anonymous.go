package main

import "fmt"

func main() {
	viral := struct {
		name     string
		username string
		age      int
	}{
		name:     "viral",
		username: "vshastri",
		age:      23,
	}
	fmt.Println(viral)

}
