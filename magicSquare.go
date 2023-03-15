// https://www.hackerrank.com/challenges/magic-square-forming/problem?isFullScreen=true

package main

import (
	"fmt"
	"os"
	"sort"
)

func formingMagicSquare(arr [][]int32) {
	//fmt.Print(arr)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !(arr[i][j] > 1 || arr[i][j] < 9) {
				os.Exit(0)
				fmt.Println("Numbers out of range")
				break
			}
		}
	}

	var s int32 = 0
	var tmp int32
	for j := 0; j < 3; j++ {
		s += arr[0][j]
	}

	// Checking if each row sum is same

	for i := 1; i <= 2; i++ {
		tmp = 0
		for j := 0; j < 3; j++ {
			tmp += arr[i][j]
		}

		if tmp != s {
			fmt.Println("Rows not tally")
			fmt.Println("Not a magic square")
		}
	}

	// Checking if each column sum is same
	for j := 0; j < 3; j++ {
		tmp = 0
		for i := 0; i < 3; i++ {
			tmp += arr[i][j]
		}
		if tmp != s {
			fmt.Println("Columns not tally")
			fmt.Println("Not a magic square")
		}
	}

	// Checking if diagonal 1 sum is same
	tmp = 0
	for i := 0; i < 3; i++ {
		tmp += arr[i][i]
	}
	if tmp != s {
		fmt.Println("Diagonal 1 not tally")
		fmt.Println("Not a magic square")
	}

	// Checking if diagonal 2 sum is same
	tmp = 0
	for i := 0; i < 3; i++ {
		tmp += arr[2-i][i]
	}
	if tmp != s {
		fmt.Println("Diagonal not tally")
		fmt.Println("Not a magic square")
	}

	fmt.Println("Magic square yes!")
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	a := [][]int{{2, 7, 6},
		{9, 5, 1},
		{5, 3, 8}}

	b := []int{}
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b = append(b, a[i][j])
			k++
		}
	}
	fmt.Print(b)
	//formingMagicSquare(a)
	for i := 1; NextPermutation(sort.IntSlice(b)); i++ {
		fmt.Println(i, b)
	}

}
