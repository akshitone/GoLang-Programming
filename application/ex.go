package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (a person) String() string {
	return fmt.Sprintf("%s %d", a.Name, a.Age)
}

type byAge []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []person{{"akshit", 21}, {"rajan", 23}, {"viral", 22}, {"suru", 19}}
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
}
