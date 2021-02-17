package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (per person) speak() {
	fmt.Println(per.first, per.age)
}

func main() {
	akshit := person{"aks", "mit", 22}
	akshit.speak()
}
