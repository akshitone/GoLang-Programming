package main

import "fmt"

type person struct {
	name     string
	username string
	age      int
}

type fbi struct {
	person
	ltk bool
}

func main() {
	fbiAkshit := fbi{
		ltk: true,
		person: person{
			name:     "akshit",
			username: "akshitone",
			age:      22,
		},
	}
	viral := person{
		name:     "viral",
		username: "vshastri",
		age:      23,
	}
	fmt.Println(fbiAkshit)
	fmt.Println(fbiAkshit.age)
	fmt.Println(fbiAkshit.ltk)
	fmt.Println(fbiAkshit.person)
	fmt.Println(viral)

}
