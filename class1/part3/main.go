// A variable Hello, world!

package main

import "fmt"

func main() {
	var name = "Marcelo"

	// Can be written as:
	// var name string = "Marcelo"	// long syntax, explicit type.
	// name := "Marcelo"			// short syntax, inferred type

	// What happens if I do this?
	// var dontUseMe = 1

	fmt.Printf("Hello, %s\n", name)
}
