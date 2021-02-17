package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{4, 7, 2, 9, 5, 8, 1, 3, 6}
	strings := []string{"viral", "rajan", "akshit", "tushar", "suru", "wanda vision"}

	fmt.Println(numbers, sort.IntsAreSorted(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers, sort.IntsAreSorted(numbers))

	fmt.Println(strings, sort.StringsAreSorted(strings))

	sort.Strings(strings)
	fmt.Println(strings, sort.StringsAreSorted(strings))

}
