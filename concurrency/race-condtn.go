package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Routines:\t", runtime.NumGoroutine())

	counter := 0
	const goRoutine = 100

	var waitGroup sync.WaitGroup
	waitGroup.Add(goRoutine)

	var mutex sync.Mutex

	for i := 0; i < goRoutine; i++ {
		go func() {
			mutex.Lock()
			v := counter

			runtime.Gosched()

			v++
			counter = v

			mutex.Unlock()
			waitGroup.Done()
		}()
		fmt.Println("GO Routines:\t", runtime.NumGoroutine())
	}

	waitGroup.Wait()

	fmt.Println("Routines:\t", runtime.NumGoroutine())
	fmt.Println("Counter:\t", counter)
}
