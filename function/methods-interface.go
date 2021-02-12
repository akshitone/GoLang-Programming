package main

import (
	"fmt"
)

type person struct {
	name     string
	username string
	age      int
}

type fbi struct {
	person
	ltk bool
}

func (fbiAgent fbi) speak() {
	fmt.Println("Hello I'm", fbiAgent.name, "- from FBI")
}
func (person person) speak() {
	fmt.Println("Hello I'm", person.name, "- from Person")
}

type human interface {
	speak()
}

func humanSpeak(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Person:", h)
	case fbi:
		fmt.Println("FBI:", h)
	}
}

func main() {
	fbiAkshit := fbi{
		ltk:    true,
		person: person{"akshit", "akshitone", 22},
	}
	fmt.Println(fbiAkshit)
	fbiAkshit.speak()

	viral := person{"viral", "vshastri", 243}
	fmt.Println(viral)
	viral.speak()

	humanSpeak(fbiAkshit)
	humanSpeak(viral)
}
