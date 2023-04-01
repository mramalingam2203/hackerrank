// https://www.hackerrank.com/challenges/find-the-median/problem?isFullScreen=true

package main

import (
	"fmt"
	"os"
	"sort"
)

func findMedian(arr []int32) int32 {
	if len(arr)%2 == 0 {
		fmt.Println("even number of elements.cant do")
		os.Exit(0)
	}

	if len(arr) < 1 || len(arr) > 1000001 {
		fmt.Println("array out of size")
		os.Exit(0)
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] < -10000 || arr[i] > 10000 {
			fmt.Println("array out of range")
			os.Exit(0)
		}
	}

	arr_int := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arr_int[i] = int(arr[i])
	}
	sort.Ints(arr_int)
	// create slice of int
	// equate slice elements to arr
	// sort by median
	return int32(arr_int[len(arr_int)/2+1])

}

func main() {
	array := make([]int32){0,1,2,4,6,5,3}

	for i := 0; i < 5; i++ {
		array[i] = int32(i + 1)
	}
	fmt.Println(findMedian(array))
}
