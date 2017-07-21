package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var seed, err = strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		panic(err)
	}

	var random = rand.New(rand.NewSource(seed))

	if x := random.Intn(10); x != 6 {
		fmt.Printf("Number is %d!\n", x)
	} else {
		fmt.Println("Random game-breaking bug! (Copyright Ubisoft)")
	}
}
