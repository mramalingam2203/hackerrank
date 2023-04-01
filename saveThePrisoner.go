//https://www.hackerrank.com/challenges/save-the-prisoner/problem?isFullScreen=true

package main

import (
	"container/ring"
	"fmt"
)

/*
 * Complete the 'saveThePrisoner' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. INTEGER s
 */

func saveThePrisoner(n int32, m int32, s int32) int32 {

	numbers := []int{}

	for i := 0; i < int(n); i++ {
		numbers = append(numbers, i+1)
	}

	r := ring.New(len(numbers))

	// initialize ring
	for i := 0; i < r.Len(); i++ {
		r.Value = numbers[i]
		r = r.Next()
	}

	// Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
	// in the ring and returns that ring element. 	r must not be empty.
	//r = r.Move(5)

	r = r.Move(int(m - 2 + s))
	r.Do(func(x interface{}) {
	})
	fmt.Println(r.Value)

	var res int = r.Value.(int)
	return int32(res)

}

func main() {
	var n, m, s int32 = 3, 7, 3
	fmt.Println(saveThePrisoner(n, m, s))
}
