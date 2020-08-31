package main

import (
	"fmt"
	"os"
	"strconv"
)

func maxPairwiseProduct(values []uint64) uint64 { // O(n)
	var greatest, secondGreatest uint64
	for _, val := range values { // using an _ because index isn't needed
		if val > greatest {
			last := greatest
			greatest = val
			secondGreatest = last
		} else if val > secondGreatest {
			secondGreatest = val
		}
	}
	return greatest * secondGreatest
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Must input at least two digits")
		return
	}
	var vals []uint64
	for i := 1; i < len(os.Args); i++ {
		if val, err := strconv.ParseUint(os.Args[i], 10, 64); err != nil {
			fmt.Println("Must input exactly two digits")
			return
		} else {
			vals = append(vals, val)
		}
	}
	fmt.Println(maxPairwiseProduct(vals))
}
