package main

import (
	"fmt"
)

// can be const or var in this case.
const (
	Bool    bool    = true
	Int     int     = 1   // int8, int16, int32, int64
	Uint    uint    = 0   // uint, uint8, uint16, uint32, uint64, uintptr
	Rune    rune    = 0   // alias for int32, represents a Unicode code point
	Float32 float32 = 0.1 // float64, complex64, complex128
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Bool, Bool)
	fmt.Printf("Type: %T Value: %v\n", Int, Int)
	fmt.Printf("Type: %T Value: %v\n", Uint, Uint)
	fmt.Printf("Type: %T Value: %v\n", Rune, Rune)
	fmt.Printf("Type: %T Value: %v\n", Float32, Float32)

	// Type Conversions
	var a float32 = 1.2
	// var b int = a // Error!
	var b int = int(a) // Ok!
	// var b = int(a) // Better!

	fmt.Println(b)
}
