// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem?isFullScreen=true
// sorting technique for int32 array

package main

import (
  "fmt"
  "sort"
  //"os"
  //"math"
)



func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func rankify(input[] int32)int32{
   // Copy input array into newArray
   sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

   // Map to store the rank of  the array element
   var ranks = make(map[int32]int32)
   var element int32
   var rank int32 = 1

   Use(ranks, element, rank)

   for index := 0; index < len(input); index++{
 		  element = input[int32(index)]
 		  //fmt.Print(element," ")
 		// Update rank of element
 		if (ranks[element] == 0){
 			ranks[element] = rank
 			rank++;
 		}
 	}
  count := 0
  var soln int32
  for key, value := range ranks{
      count++
      Use(key)
      if count == len(ranks)-1{
        soln = value
      }
  }
  return soln
}


func remove(s []int32, i int32) []int32 {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func climbingLeaderboard(ranked []int32, player []int32)[] int32 {
  soln_array := []int32{}
  for i:=0; i< len(ranked); i++{
    player = append(player, ranked[i])
    fmt.Println(player)
    soln_array = append(soln_array,rankify(player))
    player = remove(player,int32(len(player)-1))
    fmt.Println(player)
  }
  //fmt.Print(soln_array)
  return soln_array
}


func main(){

  var player = []int32{100, 100, 50, 40, 40, 20, 10}
  var ranked = []int32{5,25, 50, 120}
  climbingLeaderboard(ranked,player)

}
