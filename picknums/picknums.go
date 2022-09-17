package main

import (
  "fmt"
  "strconv"
  "os"
)

// Abs returns the absolute value of x.
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func pickingNumbers(a []int32) int32{

    if len(a) < 2 || len(a) > 100{
        fmt.Println("Arrays out of bound")
        os.Exit(0)
    }

    for i:=0; i < len(a); i++{
        if a[i] < 0 || a[i] > 100{
            fmt.Println("Values out of range")
            os.Exit(0)
        }
    }


  var count int32 = 0
  var i,j int32
  var temp string

  for i = 0; i < int32(len(a)); i++ {
      for j = 0; j < int32(len(a)); j++ {
          if Abs(a[i]-a[j]) == 1{
            if temp != strconv.Itoa(int(a[i])) + "," + strconv.Itoa(int(a[j])) {
                count +=  1
                fmt.Println(temp)
                temp = strconv.Itoa(int(a[i]))+","+strconv.Itoa(int(a[j]))
            }
          //  b = append(b,{{a[i],a[j]}}...)
          }

        }
      }
    count /= 2
    if count < 2{
        return 2
    }else{
        return count
    }
}



func main(){
  a := []int32{2,4,6,8,10,13,14}
  //a := []int32{4,6,5,3,3,1}

  pairsset := pickingNumbers(a)
  fmt.Println(pairsset)
  /*
  var array = []int32{1 ,2 ,2 ,3 ,1 ,2}
  fmt.Println(create_pairs(array))
  */
}
