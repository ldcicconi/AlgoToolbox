package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func mergeSortInts(list *[]uint64) []uint64 {
	midpoint := int(math.Floor(float64(len(*list)) / 2.0))
	if midpoint == 0 {
		return []uint64{(*list)[0]}
	}

	leftSide := (*list)[0:midpoint]
	// fmt.Println("left Side", leftSide)
	rightSide := (*list)[midpoint:]
	// fmt.Println("right Side", rightSide)
	// time.Sleep(time.Second)

	leftSorted, rightSorted := mergeSortInts(&leftSide), mergeSortInts(&rightSide)

	return mergeIntLists(leftSorted, rightSorted)
}

// Complexity: O(len(listA)+len(listB))
func mergeIntLists(listA, listB []uint64) []uint64 {
	var (
		returnList     []uint64
		aIndex, bIndex int
	)

	// Assumption: input lists are sorted lowest to highest
	for {

		// Check if we have gone through either of the lists entirely
		// if so, fill the return list with all of the values
		// from the other list
		if aIndex >= len(listA) {
			for index := bIndex; index < len(listB); index++ {
				returnList = append(returnList, listB[index])
			}
			break
		}
		if bIndex >= len(listB) {
			for index := aIndex; index < len(listA); index++ {
				returnList = append(returnList, listA[index])
			}
			break
		}

		// Merge step: look at the two elements, pick the lower one
		// like looking at tip of 2 queues and picking the smaller one
		if listA[aIndex] == listB[bIndex] || listA[aIndex] > listB[bIndex] {
			returnList = append(returnList, listB[bIndex])
			bIndex++
		} else if listA[aIndex] < listB[bIndex] {
			returnList = append(returnList, listA[aIndex])
			aIndex++
		}

	}
	return returnList
}

func parseIntListFromArgs() []uint64 {
	if len(os.Args) < 3 {
		fmt.Println("Must input at least two digits")
		panic("parseIntListFromArgs")
	}
	var vals []uint64
	for i := 1; i < len(os.Args); i++ {
		if val, err := strconv.ParseUint(os.Args[i], 10, 64); err != nil {
			fmt.Println("Must input exactly two digits")
			panic("parseIntListFromArgs")
		} else {
			vals = append(vals, val)
		}
	}
	return vals
}

func main() {
	// fmt.Println("testing merging lists")
	// listA := []uint64{0, 0, 0, 0, 4, 10, 45, 60, 61, 65}
	// listB := []uint64{1, 4, 6, 9, 55, 61, 109}
	// fmt.Println("merged", mergeIntLists(listA, listB))

	// fmt.Println("testing merge sort")
	// listC := []uint64{10, 4, 100, 150, 300, 201, 10001, 3, 5, 10, 10, 5, 101}
	// fmt.Println("merged", mergeSortInts(&listC))

	maxPairwiseProduct := func(values []uint64) uint64 {
		startTime := time.Now()
		sorted := mergeSortInts(&values)
		delta := time.Since(startTime)
		fmt.Printf("Sorted: %+v SortingDelta: %s\n", sorted, delta)
		return sorted[len(sorted)-1] * sorted[len(sorted)-2]
	}

	generateRandomInts := func(n uint64) (ret []uint64) {
		var i uint64
		for i < n {
			ret = append(ret, uint64(rand.Int63n(int64(n))))
			i++
		}
		return
	}

	rand.Seed(time.Now().UnixNano())
	values := generateRandomInts(1000)
	// values := parseIntListFromArgs()
	fmt.Println("Values:", values)
	startTime := time.Now()
	mpp := maxPairwiseProduct(values)
	delta := time.Since(startTime)
	fmt.Printf("MaxPairwiseProduct: %d Delta: %s\n", mpp, delta)

}
