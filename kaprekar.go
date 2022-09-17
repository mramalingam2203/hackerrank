
package main

import "fmt"

func sumDigits(digiNum int32)int32{
    var  digiSum, digiReminder int32
    for digiSum = 0; digiNum > 0; digiNum = digiNum / 10 {
        digiReminder = digiNum % 10
        digiSum = digiSum + digiReminder
    }
    //fmt.Println(digiSum)
    return digiSum
}

func kaprekarNumbers(p int32, q int32){
  var i int32
  for i = p; i <= q; i++{
    if sumDigits(i*i) == i{
      fmt.Println("Kaprekar Number:", i)
    }
  }
}

func main() {
  kaprekarNumbers(1,100)
}
