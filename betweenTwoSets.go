//https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true
/* Find all divisors of second Array
Filter first array with divisor array by factoring and count
*/
package main

import (
	//"bufio"
	"fmt"
	"os"
	//"strconv"
	//"strings"
)

/*
var exists = struct{}{}

type set struct {
    m map[int32]struct{}
}

func NewSet() *set {
    s := &set{}
    s.m = make(map[int32]struct{})
    return s
}

func (s *set) Add(value int32) {
    s.m[value] = exists
}

func (s *set) Remove(value int32) {
    delete(s.m, value)
}

func (s *set) Contains(value int32) bool {
    _, c := s.m[value]
    return c
}
*/

func gcd(a int32, b int32) int32 {
	//Function to calculate gcd of two numbers
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func getTotalX(a []int32, b []int32) int32 {

	var i, j int32

	// Error Handling
	for i = 0; i < int32(len(a)); i++ {
		if a[i] < 1 || a[i] > 100 {
			fmt.Println("Array A out of range")
			os.Exit(0)
		}

	}

	for i = 0; i < int32(len(b)); i++ {
		if b[i] < 1 || b[i] > 100 {
			fmt.Println("Array A out of range")
			os.Exit(0)
		}

	}

	if len(a) < 1 || len(a) > 10 || len(b) < 1 || len(b) > 10 {
		fmt.Println("Numbers out of range")
		os.Exit(0)
	}

	// Variable to find the gcd  of N numbers
	g := a[0]
	var N int32 = int32(len(a))

	// Set to store all the common
	//divisors := NewSet()
	divisors := make([]int32, 0)

	// Finding GCD of the given
	// N numbers
	for i = 1; i < N; i++ {
		g = gcd(a[i], g)
	}

	// Finding divisors of the
	// HCF of n numbers
	for i = 1; i*i <= g; i++ {
		if g%i == 0 {
			divisors = append(divisors, i)
			if g/i != i {
				divisors = append(divisors, g/i)
			}
		}
	}
	var count int32 = int32(len(divisors))
	fmt.Println(divisors)

	for i = 0; i < int32(len(divisors)); i++ {

		for j = 0; j < int32(len(b)); j++ {
			if divisors[i]%b[j] != 0 {
				fmt.Println(divisors[i], b[j], divisors[i]%b[j])
				count--
				break
			}
		}

	}
	fmt.Println(count)
	return count
}

func main() {
	a := []int32{2, 6}
	b := []int32{24, 36}
	fmt.Println(getTotalX(b, a))

	//fmt.Println(getTotalX(a,b))

}
