// https://www.hackerrank.com/challenges/sherlock-and-squares/problem?isFullScreen=true

package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	//floor(sqrt(b)) - ceil(sqrt(a)) + 1
	var float_a, float_b = float64(a), float64(b)
	return int32(math.Floor(math.Sqrt(float_b)) - math.Ceil(math.Sqrt(float_a)) + 1)

}

func main() {
	var a, b int32 = 1, 30
	fmt.Print(squares(a, b))
}
