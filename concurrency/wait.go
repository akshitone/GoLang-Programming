package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())

	waitGroup.Add(1)
	go foo()
	bar()

	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())
	waitGroup.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	waitGroup.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
