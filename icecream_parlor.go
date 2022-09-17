
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

func indexOf(e1 int32,e2 int32, data []int32){
//([]int32,[]int32) {
   var i int32
   indices1 := make([]int32,0)
   indices2 := make([]int32,0)

   for i=0; i< int32(len(data));i++{
     if data[i] == e1{
       indices1 = append(indices1,i+1)
     }else if data[i] == e2{
       indices2 = append(indices2,i+1)
   }
 }
 fmt.Print(indices1, indices2)
 /*
  indices := make(int32,0)
    indices = append(indices,indices1[0])
   for i:=0; i < len(indices1);i++{
    for j:=0; j < len(indices2); j++{
      if indices1[i] != indices2[j]{
        indices = append(indices,indices2[j])
        break
      }
    }
}

  fmt.Print(indices)
*/
   //return indices1, indices2
}


func icecreamParlor(m int32, arr []int32) []int32 {
  var pools int32 = 2
  c := Combinations(pools, arr)
  buy := make([]int32,2)
  //flavors := make([]int32,2)
  //fmt.Print(c)
  for i:=0; i < len(c[1]);i++{
    if c[1][i][0] + c[1][i][1] == m {
      buy[0],buy[1] = c[1][i][0] ,c[1][i][1]
      break
    }
  }
  //indices:= make([]int32,0)
  //indices =
  //append(indices,
    indexOf(buy[0], buy[1], arr)
  return arr
}


func main(){

  costs := []int32{2,2,4,3}
  var m int32 = 4

  fmt.Print(icecreamParlor(m, costs))
  return
  //fmt.Println(indexOf(3,costs))
}
