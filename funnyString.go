package main

import (
	"fmt"
	"os"
)

func funnyString(s string) string {
	if len(s) < 2 || len(s) > 10000 {
		fmt.Println("Invalid string length")
		os.Exit(0)
	}
	r := []rune(s)
	rev_r := []rune(s)
	reverseArray(rev_r)
	temp := true
	for i := 1; i < len(r)-1; i++ {
		temp = temp && abs(r[i]-r[i-1]) == abs(rev_r[i]-rev_r[i-1])
	}
	if temp == true {
		return "Funny"
	} else {
		return "Not Funny"
	}

}

func abs(n int32) int32 {
	if n < 0 {
		return -1 * n
	} else if n >= 0 {
		return n
	}
	return n
}

func reverseArray(arr []rune) []rune {
	rev_arr := arr
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		rev_arr[i], rev_arr[j] = rev_arr[j], rev_arr[i]
	}

	return rev_arr
}

func main() {
	fmt.Println(funnyString("bcxz"))

}
