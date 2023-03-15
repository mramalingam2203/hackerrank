// https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true&h_r=next-challenge&h_v=zen
package main

import (
	//"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func dayOfProgrammer(year int32) string {

	if !(year > 1700 && year < 2700) {
		fmt.Println("Year out of range")
		os.Exit(0)
	}

	no_of_days_in_a_month := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%4 == 0 {
		no_of_days_in_a_month[1] = 29
	}

	sum := 0
	var output string
	var zerodot string = "."
	for i := 0; i < 12; i++ {
		sum += no_of_days_in_a_month[i]

		if sum >= 256 {
			if i < 9 {
				zerodot += "0"
			}
			//fmt.Println(i+1, no_of_days_in_a_month[i]-(sum-256))
			output = strconv.Itoa(no_of_days_in_a_month[i]-(sum-256)) + zerodot +
				strconv.Itoa(i+1) + "." +
				strconv.Itoa(int(year))
			break
		}
	}
	return output
}

func main() {
	fmt.Print(dayOfProgrammer(2016))
}
