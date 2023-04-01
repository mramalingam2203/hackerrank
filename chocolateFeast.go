// https://www.hackerrank.com/challenges/chocolate-feast/problem?isFullScreen=true

package main

import (
	"fmt"
	"os"
)

func chocolateFeast(n int32, c int32, m int32) int32 {

	if n < 2 || n > 100000 || c < 1 || c > n || m < 2 || m > n {
		fmt.Println("Mismatchng data size. Cannot proceed")
		os.Exit(0)

	}

	var w, res int32
	w = n / c
	res = n / c
	for w >= m {
		res += w / m
		w = (w % m) + w/m
	}
	return res
}

func main() {
	fmt.Println(chocolateFeast(15, 3, 2))
}
