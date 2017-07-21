// A function-oriented Hello, world!
package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func prod(a, b int) (x int) {
	x = a * b
	return
}

func sumProd(a, b int) (int, int) {
	return a + b, a * b
}


func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(prod(2, 3))

	var x, y = sumProd(2, 3)
	fmt.Println(x, y)
}
