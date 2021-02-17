package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {
	akshit := user{"akshit", 21}
	rajan := user{"rajan", 23}
	viral := user{"viral", 22}

	users := []user{akshit, rajan, viral}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(jsonData))
}
