// A CLI for Hello's!

package main

// Note the abbreviation
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Errorf("No name given!"))
	}

	fmt.Printf("Hello, %s!\n", os.Args[1])
}
