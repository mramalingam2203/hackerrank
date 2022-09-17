package main

import (
//    "bufio"
  "fmt"
  "math"
//  "math"
/*    "io"
    "os"*/
    "strconv"
//    "strings"

)

func chunkSlice(slice []string) []string {

  var chunks [][]string

  var chunkSize int = int(math.Sqrt(float64(len(slice))))

  for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
  fmt.Println(chunkSize)

  for i:=1; i < chunkSize-1;i++{
    for j:=1;j < chunkSize-1;j++{
      mid,_ := strconv.Atoi(chunks[i][j])
      west,_ := strconv.Atoi(chunks[i-1][j])
      east,_ := strconv.Atoi(chunks[i+1][j])
      south,_ := strconv.Atoi(chunks[i][j-1])
      north,_ := strconv.Atoi(chunks[i][j+1])
      if mid > west && mid > east && mid > north  && mid > south{
          fmt.Println(i,j)
          chunks[i][j] = "X"
        }
      }
    }

	return Flatten(chunks)
}

func Flatten[T any](lists [][]T) []T {
    var res []T
    for _, list := range lists {
        res = append(res, list...)
    }
    return res
}

func main(){
  //grid := []string{"1","1","1","2","1","9","1","2","1","8","9","2","1","2","3","4"}
  //grid := []string{"9","8","9","1","9","1","1","1","1"}
  grid := []string{"1","1","1","2","1","9","1","2","1","8","9","2","1","2","3","4"}
  chunkSlice(grid)
  //fmt.Print(result)


}
