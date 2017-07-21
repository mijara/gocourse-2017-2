package main

import "fmt"

func fact(x uint) uint {
	var r uint = 1

	for i := uint(1); i <= x; i++ {
		r *= i
	}

	return r
}

func main() {
	fmt.Println(fact(4))
}
