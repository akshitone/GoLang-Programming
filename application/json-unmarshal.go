package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
	Ltk   bool   `json:"Ltk"`
}

func main() {
	str := `[
		{"First":"akshit","Last":"mithaiwala","Age":22,"Ltk":true},
		{"First":"viral","Last":"shastri","Age":23,"Ltk":false}
		]`

	bytePeople := []byte(str)

	// fmt.Println(str) ---> json data into string format
	// fmt.Println(bytePeople) ---> byte data

	// people := []person{} ---> it's little hard to read

	var people []person

	// here, it's converted byte data of people into go data people using pointer address

	err := json.Unmarshal(bytePeople, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Data of all person:", people)

	for index, person := range people {
		fmt.Println("Person:", index+1)
		fmt.Println(person.First, person.Last, "|", person.Age, "|", person.Ltk)
	}
}
