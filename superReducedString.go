// https://www.hackerrank.com/challenges/reduced-string/problem?isFullScreen=true

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func superReducedString(s string) string {

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] {
			fmt.Println(s[i+2])
		}
	}

	//fmt.Println(s)
	return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	a := readLine(reader)
	superReducedString(a)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
