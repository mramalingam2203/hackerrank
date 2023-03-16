package main

import "fmt"

/*
 * Complete the 'minimumDistances' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func minimumDistances(a []int32) int32 {
	n := len(a)
	minimum := n + 1
	index := make(map[int32]int32)
	_, _ = minimum, index
	for i := 0; i < n; i++{
		if a[i] in index{
			if i-index[a[i]] < minimum {
				fmt.Print("awe")
			}
		}
		

}

func main() {
	//var array = make([]int32, 6)
	array := []int32{7, 1, 3, 4, 1, 7}

	fmt.Print(minimumDistances(array))

}
