package main

import "fmt"

func main() {
	users := map[string]map[string]string{
		"akshit": {"id": "235", "username": "akshitone"},
		"viral":  {"id": "243", "username": "vshastri"},
	}

	fmt.Println(users)
	fmt.Println(users["akshit"])
	fmt.Println(users["akshit"]["username"])

	for user, userdata := range users {
		fmt.Println("\ndetails of", user)
		for field, data := range userdata {
			fmt.Println(field, ":", data)
		}
	}

	another := map[string][]string{
		"state": []string{"Guj", "Mah"},
		"city":  []string{"Sur", "Raj"},
	}
	fmt.Println(another)

	for key, state := range another {
		fmt.Println("\ndetails of", key)
		for index, city := range state {
			fmt.Println(index, city)
		}
	}

	delete(users["akshit"], "id")
	fmt.Println(users)
}
