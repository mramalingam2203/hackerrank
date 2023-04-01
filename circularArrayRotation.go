package main

import (
	"fmt"
	"os"
)

/*
 * Complete the 'circularArrayRotation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER k
 *  3. INTEGER_ARRAY queries
 */

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	switch {
	case len(a) < 1 || len(a) > 1e5:
		fmt.Println("Array too long")
		os.Exit(0)
	case k < 1 || k > 1e5:
		fmt.Println("Too many rotations")
		os.Exit(0)
	case len(queries) < 1 || len(queries) > 500:
		fmt.Println("Too few or many queries")
		os.Exit(0)
	}

	for i := 0; i < len(a); i++ {
		if a[i] < 1 || a[i] > 1e5 {
			fmt.Println("Array values out of range")
			os.Exit(0)
		}
	}

	for i := 0; i < len(a); i++ {
		if queries[i] < 0 || queries[i] > int32(len(a)) {
			fmt.Println("Array values out of range")
			os.Exit(0)
		}
	}

	rotateR(&a, int(k))
	fmt.Print(a)

	s := make([]int32, len(queries))

	for i := 0; i < len(queries); i++ {
		s[i] = a[queries[i]]
	}

	return s
}

// circular rotation rightward
func rotateR(a *[]int32, i int) {
	x, b := (*a)[:(len(*a)-i)], (*a)[(len(*a)-i):]
	*a = append(b, x...)
}

func main() {
	var a = []int32{3, 4, 5}
	var k int32 = 1
	var q = []int32{1, 2}
	circularArrayRotation(a, k, q)
}
