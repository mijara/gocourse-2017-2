package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 1

	for sum < 1000 { // ... What!?
		sum += sum
	}

	// This is a while loop!

	fmt.Println(sum)

	// just for the LoLs
	for {
		fmt.Println("yes")
		time.Sleep(500 * time.Millisecond)
	}
}
