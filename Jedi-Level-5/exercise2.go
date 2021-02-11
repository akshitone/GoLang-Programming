package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	akshit := person{"akshit", "mithaiwala", []string{"chocolate", "strawbarry"}}
	viral := person{"viral", "shastri", []string{"vanilla", "american nuts"}}
	fmt.Println(akshit)
	fmt.Println(viral)

	userMap := map[string]person{
		akshit.firstName: akshit,
		viral.firstName:  viral,
	}
	fmt.Println(userMap)

	for user, userData := range userMap {
		fmt.Println(user, ":", userData.firstName, userData.lastName)
		for index, iceCream := range userData.favoriteIceCream {
			fmt.Println(index, iceCream)
		}
	}
}
