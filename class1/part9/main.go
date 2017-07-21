package main

import "fmt"

func recursiveFact(x int) int {
	if x == 1 {
		return 1
	}

	return recursiveFact(x - 1) * x
}

func main() {
	fmt.Println(recursiveFact(4))
}
