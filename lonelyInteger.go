
package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {

	//slice := generateSlice(20)
  slice := []int32{1}
	quicksort(slice)
	//fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
  fmt.Println(lonelyinteger(slice))
}

func filter(a []int32, b []int32)int32{
  //c := []int32{}
  var c int32
  for i:=0; i < len(b); i++{
      if a[i] != b[i]{
        c  = a[i]
        break
      }else{
        c = a[len(a)-1]
      }
  }
  //fmt.Println(c)
  return c
}

func lonelyinteger(a []int32) int32 {
  var lonelyint int32 = 0
  var odd []int32
  var even []int32
  for i:=0; i< len(a); i++{
    if a[i] > 0{
    if i%2 == 0 {
        odd = append(odd, a[i])
    }else{
      even = append(even, a[i])
    }
  }

  }
  fmt.Print(odd)
  fmt.Print(even)

  lonelyint = filter(odd,even)
  return lonelyint
}

/*
// Generates a slice of size, size filled with random numbers
func generateSlice(size int32) []int32 {

	slice := make([]int32, size, size)
	rand.Seed(time.Now().UnixNano())
  var i int32
	for i= 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
*/

func quicksort(a []int32) []int32 {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)

    a[pivot], a[right] = a[right], a[pivot]

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    quicksort(a[:left])
    quicksort(a[left+1:])

    return a
}
