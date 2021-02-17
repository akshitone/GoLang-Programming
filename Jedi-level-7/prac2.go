package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	jsonString := `[
		{"Name":"akshit","Age":21},
		{"Name":"rajan","Age":23},
		{"Name":"viral","Age":22}
		]`

	var users []user

	err := json.Unmarshal([]byte(jsonString), &users)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(users)
}
