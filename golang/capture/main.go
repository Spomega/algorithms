package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i) // Capturing the iteration variable `i`
		}(i)
	}

	// Give the goroutines time to finish
	time.Sleep(1 * time.Second)

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(2)                                     // Set to 2
	fmt.Println("Updated GOMAXPROCS:", runtime.GOMAXPROCS(0)) // Print updated

}
