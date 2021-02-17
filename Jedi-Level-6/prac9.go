package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	callBack := func(slice []int) int {
		if len(slice) == 0 {
			return 0
		}
		if len(slice) == 1 {
			return slice[0]
		}
		return slice[0] + slice[len(slice)-1]

	}
	fmt.Println(takeCallBack(callBack, numbers))
}

func takeCallBack(callBack func(slice []int) int, slice []int) int {
	number := callBack(slice)
	number++
	return number

}
