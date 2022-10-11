// https://www.hackerrank.com/challenges/reduced-string/problem?isFullScreen=true

package main

import (
  "fmt"
  "os"
  //"sort"
)


func removeDuplicateInt(intSlice []rune) []rune {
    allKeys := make(map[int]bool)
    list := []rune{}
    for _, item := range intSlice {
        if _, value := allKeys[int(item)]; value {
            allKeys[int(item)] = true
            list = append(list,item)
        }
    }
    return list
}


func appendCategory(d []rune) []rune {
	check := make(map[rune]int)
	res := make([]rune,0)
	for _, val := range d {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res,letter)
	}

	return res
}



func superReducedString(s string) string {

  if len(s) < 1 || len(s) >100{
    fmt.Println("string invalid")
    os.Exit(0)
  }
  r := []rune(s)
  //fmt.Print(r)
  //fmt.Print(removeDuplicateInt(r))
  fmt.Println(appendCategory(r))
  return s
}

func main(){
  s := "aaabccddd"
  superReducedString(s)

}