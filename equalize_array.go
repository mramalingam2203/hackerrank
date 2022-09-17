package main
import (
	"os"
	"fmt"
)

func mostFrequent(arr []int32) int32 {
	if len(arr) < 1 || len(arr) > 100{
		fmt.Println("Too lengthy an array")
		os.Exit(0)
	}
	for i:=0; i < len(arr); i++{
		if arr[i] < 1|| arr[i] > 100{
			fmt.Println("Array Values out of range")
			os.Exit(0)
		}
	}


	// assuming no tie
	m := map[int32]int32{}
	var maxCnt int32
	var freq int32
	for _, a := range arr {
		m[a]++
		if m[a] > maxCnt {
			maxCnt = m[a]
			freq = a
		}
	}

	return int32(len(arr))-freq
  }

func main(){
	arr := []int32{1,2,1,2}
  fmt.Println(mostFrequent(arr))

}
