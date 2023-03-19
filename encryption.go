//https://www.hackerrank.com/challenges/encryption/problem?isFullScreen=true

package main

import (
	"bytes"
	"fmt"
	//"math"
	"os"
	"strings"
	"unicode"
)

func constraints(s string) {
	if len(s) < 1 || len(s) > 81 {
		fmt.Println("String length out of range")
		os.Exit(0)
	}
}

func encryption(s string) {
	s = stripSpaces(s)
	//var number float64 = math.Sqrt(float64(len(s)))
	s = insertNth(s, 8)
	fmt.Print(s)
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}

func main() {
	input := "if man was meant to stay on the ground god would have given us roots"
	constraints(input)
	encryption(input)

}
