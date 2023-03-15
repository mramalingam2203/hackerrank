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
 * Complete the 'matrixRotation' function below.
 *
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY matrix
 *  2. INTEGER r
 */

func matrixRotation(matrix [][]int32, r int32) {
	var ring_m, ring_n, rings, x, i, j, p, q int32
	var arr [][]int32
	ring_m = int32(len(matrix))
	ring_n = int32(len(matrix[0]))
	fmt.Println(ring_m, ring_n)
	if ring_m < ring_n {
		rings = ring_m
	} else {
		rings = ring_n
	}
	fmt.Println(rings)

	for x = 0; x < (rings / 2); x++ {
		temp_r := r % (((ring_m - (x * 2)) + (ring_n - (x * 2)) - 2) * 2)
		fmt.Println(temp_r)
		for temp_r > 0 {
			i, j = x, x
			p = arr[i][j]
			for i < ring_m {
				q = arr[i][j]
				arr[i][j] = p
				p = q
				i++
			}
			i--
			j++

			for j < ring_n {
				q = arr[i][j]
				arr[i][j] = p
				p = q
				j++
			}
			j--
			i--

			for i >= x {
				q = arr[i][j]
				arr[i][j] = p
				p = q
				i--
			}
			i++
			j--
			for j >= x {
				q = arr[i][j]
				arr[i][j] = p
				p = q
				j--
			}
		}

		temp_r--
	}
	ring_m--
	ring_n--
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	rTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	r := int32(rTemp)

	var matrix [][]int32
	for i := 0; i < int(m); i++ {
		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(n) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	matrixRotation(matrix, r)
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

/*
func rotateRight(nums []int, k int) {
	k %= len(nums)
	new_array := make([]int, len(nums))
	copy(new_array[:k], nums[len(nums)-k:])
	copy(new_array[k:], nums[:len(nums)-k])
	copy(nums, new_array)
}



func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("The entered array is:", arr)
	rotateRight(arr, 2)
	fmt.Println("The array obtained by rotating it to right by 7 positions is:", arr)
}
*/
