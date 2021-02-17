package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type byName []user

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	akshit := user{"akshit", 21}
	rajan := user{"rajan", 23}
	viral := user{"viral", 22}

	users := []user{akshit, rajan, viral}
	fmt.Println(users)
	sort.Sort(byName(users))
	fmt.Println(users)
}
