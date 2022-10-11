package main
import (
  "fmt"
  "math"
  //"os"
)


func multiply(a float64, b float64){
  fmt.Println(      math.Pow ( 2.0, math.Log2(a) + math.Log2(b)))
  //fmt.Println(math.Pow(math.Log2(a) + math.Log2(b), 2.0))
}

func factorial(num float64) float64{
   if num == 1 || num == 0{
      return 1
   }
   //fmt.Println(math.Log2(num), math.Log2(factorial(num-1)))
   return math.Pow(2.0, math.Log2(num) * math.Log2(factorial(num-1)))
}


func main(){
   fmt.Println(factorial(30.0))
   fmt.Println(factorial(40.0))
   fmt.Println(factorial(50.0))
   //multiply(20.0, 34.0)
}
