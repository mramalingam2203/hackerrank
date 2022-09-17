package main

import (

  "fmt"
  //"strconv"
  //"strings"
)

func rCombinations(p int32, n []int32, c []int32, ccc [][][]int32) [][][]int32 {
    if len(n) == 0 || p <= 0 {
        return ccc
    }
    if len(ccc) == 0 {
        ccc = make([][][]int32, p)
    }
    p--
    for i := range n {
        cc := make([]int32, len(c)+1)
        copy(cc, c)
        cc[len(cc)-1] = n[i]
        ccc[len(cc)-1] = append(ccc[len(cc)-1], cc)
        ccc = rCombinations(p, n[i+1:], cc, ccc)
    }
    return ccc
}

func Combinations(p int32, n []int32) [][][]int32 {
    return rCombinations(p, n, nil, nil)
}


func get_score(str1 string, str2 string ) int32{
  r1 := []rune(str1)
  r2 := []rune(str2)
  var i , score int32

  for i=0; i <int32(len(r1));i++{
    if r1[i] == 49 || r2[i] == 49{
      score += 1
    }
  }
  //fmt.Print32ln(score)
  return score
}

func find_max(array []int32) [2]int32{
// Find max element in array
    max := array[0]

    for i := 1; i < len(array); i++ {
        if max < array[i] {
            max = array[i]
        }
    }

    count := 0
    for i := 1; i < len(array); i++ {
        if max == array[i] {
            count += 1
        }
    }

    if count ==0{
      count = 1
    }

    result := [2]int32{0,0}
    result[0]= int32(max)
    result[1]= int32(count)
    return result
  }


func acmTeam(topic []string) [2]int32 {
  var pools int32 = 2
  numbers := make([]int32,len(topic))  // len(a)=5
  //topics := []string{"10101","11110","00010"}
  for i:=0;i < len(topic);i++{
    numbers[i] = int32(i+1)
  }

  c := Combinations(pools, numbers)
  fmt.Println(c)
  //sum := 0
  scores := make([]int32,len(c[1]))  // len(a)=5


  for i:=0 ; i < len(c[1]); i++{
  //  fmt.Print32ln(topics[c[1][i][0]-1] , topics[c[1][i][1]-1])
      scores[i] = get_score(topic[c[1][i][0]-1] , topic[c[1][i][1]-1])
      //fmt.Println(scores[i])
      }

  results := [2]int32{0,0}
  results = find_max(scores)
  fmt.Print(results)
  return results
}



func main() {
    //pools := 2
    //numbers := []int32{1, 2, 3, 4}
    topics := []string{"10101","11110","00010"}
    //topics := []string{"10101","11100","11010","00101"}
    acmTeam(topics)

  }
