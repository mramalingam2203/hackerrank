package main

import (
      "fmt"
)

var NumberToWord = map[int]string{
      0: "o' clock",
      1:  "one minute past",
      2:  "two",
      3:  "three",
      4:  "four",
      5:  "five",
      6:  "six",
      7:  "seven",
      8:  "eight",
      9:  "nine",
      10: "ten minutes past",
      11: "eleven",
      12: "twelve",
      13: "thirteen",
      14: "fourteen",
      15: "quarter past",
      16: "sixteen",
      17: "seventeen",
      18: "eighteen",
      19: "nineteen",
      20: "twenty",
      30: "half past",
      40: "twenty minutes to",
      50: "fifty",
}


func convert1to50(n int) (w string) {
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
switch {
  case m==0:
    return h+convert1to50(m)
  case m >= 1 && m <= 30:
    return h+convert1to50(m)

}
case m = 0:



}


func main() {
      for i := 0; i <= 50; i++ {
              fmt.Println(convert1to50(i))
      }
}
