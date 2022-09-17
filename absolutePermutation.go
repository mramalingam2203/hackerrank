package main

import
(
  "fmt"
  "os"
)


//function to swap the variables
func swap(a *int32,b *int32){
    var temp int32
    temp = *a
    *a = *b
    *b = temp
}

// function to find Absolute value of an inteer
func Abs(a int32) int32{
  if a < 0{
    return -1*a
  }else{
    return a
  }
}

//function to print the array
func printarray(arr[] int32){
  //  if isAbsolutePermut(arr,2) == true{
    for i:=0; i<len(arr); i++{
        fmt.Print(arr[i])
    }
    fmt.Println()
//  }
}

func isAbsolutePermut(arr[] int32, k int32)bool{
  var i int32
  var temp bool = true
  for i = 0; i < int32(len(arr));i++{
    if Abs(arr[i]-i-1) == k{
      temp = temp && true
      }else{
        temp = temp && false
      }
  }
  return temp
}

func permutation(arr []int32, start int32, end int32){
    if start == end{
      printarray(arr)
    }
    var i int32
    for i =start; i<=end; i++{
      swap(&arr[i], &arr[start])
      permutation(arr, start+1, end);
      swap(&arr[i], &arr[start])
    }
}

func absolutePermutation(n int32, k int32){
    if n< 1 || n > 1e5{
      fmt.Println("n out of range!")
      os.Exit(0)
    }

    if k < 0 || k >n{
      fmt.Println("Unreassonable k!")
      os.Exit(0)
    }

    array := []int32{}

    for i:=0; i < int(n); i++{
      array = append(array,int32(i+1))
    }

    // Write your code here
    permutation(array, 0, n-1)

}


func main(){
  //array := []int32{1,2,3,4}
  //permutation(array, 0, int32(len(array)-1))
  //pointerArith()
  absolutePermutation(4,2)
}
