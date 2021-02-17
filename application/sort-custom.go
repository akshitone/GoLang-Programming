package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byAge []person
type byName []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a byName) Len() int {
	return len(a)
}

func (a byName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func main() {
	akshit := person{"akshit", 21}
	rajan := person{"rajan", 23}
	viral := person{"viral", 22}
	suru := person{"suru", 19}

	people := []person{akshit, rajan, viral, suru}
	fmt.Println(people)

	sort.Sort(byAge(people))
	fmt.Println("Sorted by age:", people)

	sort.Sort(byName(people))
	fmt.Println("Sorted by name:", people)
}
