package main

import "fmt"

func main() {
	repeatAlpha()
	forCondtion()
	forJust()
	remainder()
}

func repeatAlpha() {
	for i := 65; i <= 67; i++ {
		fmt.Println(i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}

func forCondtion() {
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func forJust() {
	year := 1995
	for {
		if year == 1999 {
			fmt.Println("Born year:", year)
			break
		}
		year++
	}
}

func remainder() {
	for i := 10; i < 100; i++ {
		fmt.Println(i % 4)
	}
}
