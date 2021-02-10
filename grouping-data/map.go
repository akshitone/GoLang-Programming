package main

import "fmt"

func main() {
	user := map[string]string{
		"akshit": "akshitone",
		"viral":  "vshastri",
	}
	fmt.Println(user)
	fmt.Println(user["akshit"])
	fmt.Println(user["rajan"])

	value, out := user["viral"]
	fmt.Println(value)
	fmt.Println(out)

	if value, out := user["akshit"]; out {
		fmt.Println(value)
	}

	fmt.Println("----")

	user["rajan"] = "rjenfield"
	for key, value := range user {
		fmt.Println(key, value)
	}

	delete(user, "rajan")
	fmt.Println(user)
}
