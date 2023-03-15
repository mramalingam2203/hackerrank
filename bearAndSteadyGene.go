//https://www.hackerrank.com/challenges/bear-and-steady-gene/problem?isFullScreen=true
/* A gene is represented as a string of length  (where  is divisible by ), composed of the letters , , , and . It is considered to be steady if each of the four letters occurs exactly  times. For example,  and  are both steady genes.
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'steadyGene' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING gene as parameter.
 */

func steadyGene(gene string) int32 {

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	if n < 4 || n > 500000 {
		fmt.Println("n out of range")
		os.Exit(0)
	}
	if n%4 != 0 {
		fmt.Println("n not a multiple of 4")
		os.Exit(0)
	}

	gene := readLine(reader)
	if !(strings.Contains(gene, "A") || strings.Contains(gene, "G") || strings.Contains(gene, "C") || strings.Contains(gene, "T")) == true {
		fmt.Println("bad character in gene")
		os.Exit(0)
	}

	result := steadyGene(gene)

	fmt.Fprintf(writer, "%d\n", result)

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
