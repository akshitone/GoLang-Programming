package main

import "fmt"

func main() {
	name("Akshit Mithaiwala")

	result := anotherName("Harvey Specter")
	fmt.Println(result)

	name, married := nineNine("Jake Peralta", "Yes")
	fmt.Println(name, married)
}

func name(name string) {
	fmt.Println("Good morning,", name)
}

func anotherName(name string) string {
	return fmt.Sprint("This is contacting with ", name)
}

func nineNine(name string, married string) (string, bool) {
	if married == "Yes" {
		return name, true
	}
	return name, false
}
