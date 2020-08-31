package main

import (
	"fmt"
	"os"
	"strconv"
)

func sumOfNonNegativeIntegers(a, b uint64) uint64 {
	return a + b
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Must input exactly two digits")
		return
	}
	a, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Must input exactly two digits")
		return
	}
	b, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println("Must input exactly two digits")
		return
	}
	fmt.Println(sumOfNonNegativeIntegers(a, b))
}
