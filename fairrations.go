package main
import(
  "fmt"
  "strconv"
  "os"
)


func fairRations(B []int32) string {
  var s,i int32 // fnd total sum of array // BUG:
  var res int32 = 0

  if B[0] <2 || B[0] > 1000{
    fmt.Println("Too many subjects")
    os.Exit(0)
  }

  for i=1; i < B[0]; i++{
    if B[i] < 1 || B[i] > 10{
      fmt.Println("Cannot give too few or too many loaves")
      os.Exit(0)
    }
  }


  if s%2 ==1{
    return "NO"
  }else {
    for i=0; i<int32(len(B)); i++{
      if B[i]%2 == 1{
        B[i+1] +=1
        res +=2
      }
    }

  }
  return strconv.Itoa(int(res))
}





func main(){

  Bread := []int32{4,5,6,7}
  fmt.Println(fairRations(Bread))
}
