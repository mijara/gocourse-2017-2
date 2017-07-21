// A spicier Hello, world!
package main

import "fmt"

func main() {
	// zero valued variable
	var name string

	fmt.Print("Your name is?: ")

	// pointers!
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Hello, %s!\n", name)
}
