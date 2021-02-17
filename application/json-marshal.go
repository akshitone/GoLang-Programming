package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
	Ltk   bool
}

func main() {
	akshit := person{"akshit", "mithaiwala", 22, true}
	viral := person{"viral", "shastri", 23, false}

	people := []person{akshit, viral}
	fmt.Println(people)

	bytePeople, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bytePeople))
}
