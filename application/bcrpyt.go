package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `akshitone@1999`
	confirmPassword := `akshitone1999`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Println(hashedPassword)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(confirmPassword))
	if err != nil {
		fmt.Println("mismatch!")
		return
	}
	fmt.Println("logged in!")
}
