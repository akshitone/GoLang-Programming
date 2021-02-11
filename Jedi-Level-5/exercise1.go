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

	fmt.Println(akshit.firstName)
	for index, iceCream := range akshit.favoriteIceCream {
		fmt.Println(index, iceCream)
	}
}
