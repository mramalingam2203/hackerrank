//https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem?isFullScreen=true

package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {

	var count = make(map[string]int)
	var words = strings.Split(s, "")
	//fmt.Println(words)
	for _, word := range words {
		count[word]++
	}

	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
		if v != a[0] {
			fmt.Println("Invalid string")
		}
	}
}

func main() {
	isValid("Hello")
}

/*
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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
*/
