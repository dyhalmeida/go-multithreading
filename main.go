package main

import (
	"fmt"
	"sync"
	"time"
)

func Task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - Task %s is running\n", i+1, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Wait Proups Parts

// add number of operations
// report that an operation has finished
// wait until all operations have finished

func SimultaneousProcessing() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	// Task A is thread 2
	go Task("A", &waitGroup)

	// Task B is thread 3
	go Task("B", &waitGroup)

	// Task anonymous C is thread 4
	go func(name string, wg *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d - Task anonymous %s is running\n", i+1, name)
			time.Sleep(1 * time.Second)
			wg.Done()
		}
	}("C", &waitGroup)

	waitGroup.Wait()
}

// main is thread 1
func main() {
	fmt.Println("SimultaneousProcessing start")
	fmt.Println("")

	SimultaneousProcessing()

	fmt.Println("")
	fmt.Println("SimultaneousProcessing finish")
}
