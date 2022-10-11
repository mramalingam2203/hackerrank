// https://www.hackerrank.com/domains/algorithms?filters%5Bstatus%5D%5B%5D=unsolved
package main

import (
      "fmt"
      "os"
)

var NumberToWord = map[int32]string{
      0: " o' clock",1:  "one",2:  "two",3:  "three",4:  "four",5:  "five",
      6:  "six",7:  "seven",8:  "eight",9:  "nine",10: "ten",11: "eleven",
      12: "twelve",13: "thirteen",14: "fourteen",15: "quarter",16: "sixteen",
      17: "seventeen",18: "eighteen",19: "nineteen",20: "twenty",30: "half",
      40: "forty",50: "fifty"}


func convert1to50(n int32) (w string) {
      if n < 20 {
              w = NumberToWord[n]
              return
      }
      r := n % 10
      if r == 0 {
              w = NumberToWord[n]
      } else {
              w = NumberToWord[n-r] + " " + NumberToWord[r]
      }
      return
}

func timeInWords(h int32, m int32) string {
  if h < 1 || h >  12{
    fmt.Println("Hours invalid")
    os.Exit(0)
  }

  if m < 0 || m >  60{
    fmt.Println("Minutes invalid")
    os.Exit(0)
  }
  var msg string
  switch {
    case m == 0:
      msg = convert1to50(h) + convert1to50(m)
    case m == 1:
      msg = convert1to50(m) + " minute past " + convert1to50(h)
    case m == 15 || m == 30:
      msg = convert1to50(m) + " past " + convert1to50(h)
    case m == 45:
      msg = convert1to50(15) + " to " + convert1to50(h)
    case m > 1 && m < 30:
      msg = convert1to50(m) + " minutes past " + convert1to50(h)
    case m > 30 || m < 45 && m != 45 :
      msg = convert1to50(60-m) + " minutes to " + convert1to50(h+1)
    }
    fmt.Print(msg)
    fmt.Println()
    return msg
}


func main() {
      fmt.Println(timeInWords(5,47))
}
