package main

import (
	"fmt"
	"time"
)

func Task(name string) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d - Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func SerialProcessing() {
	Task("A")
	Task("B")
}

// SimultaneousProcessing is thread 1
func SimultaneousProcessing() {

	// Task C is thread 2
	go Task("C")

	// Task D is thread 3
	go Task("D")

	// Task anonymous E is thread 4
	go func(name string) {
		for i := 1; i < 11; i++ {
			fmt.Printf("%d - Task anonymous %s is running\n", i, name)
			time.Sleep(1 * time.Second)
		}
	}("E")

	// temporary solution to arrest Simultaneous Processing
	time.Sleep(10 * time.Second)
}

func main() {
	fmt.Println("SerialProcessing start")
	fmt.Println("")

	SerialProcessing()

	fmt.Println("")
	fmt.Println("SerialProcessing finish")

	fmt.Println("")
	fmt.Println("SimultaneousProcessing start")
	fmt.Println("")

	SimultaneousProcessing()

	fmt.Println("")
	fmt.Println("SimultaneousProcessing finish")
}
