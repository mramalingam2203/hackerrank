package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {
	var words int32 = 1
	for _,r:=range s {
   if unicode.IsUpper(r) {
			words += 1
   }
	 if words > 1e5{
		 fmt.Println("Sentence too long")
		 os.Exit(0)
	 }
}
	return words
}

func main(){
	sentence := "muthuAteAChickenAndTwoEggsForDinner"
	fmt.Println(camelcase(sentence))
}
