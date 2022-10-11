// https://www.hackerrank.com/challenges/strong-password/problem?isFullScreen=true

package main

import (
    "fmt"
    "os"
    "strings"
  )


func minimumNumber(n int32, password string) int32 {

  if n < 1 || n > 100{
    fmt.Println("invalid password")
    os.Exit(0)
  }

  if len(password) < int(n){
    fmt.Println("Password too short")
    os.Exit(0)
  }

  numbers := "0123456789"
  lower_case := "abcdefghijklmnopqrstuvwxyz"
  upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  special_characters := "!@#$%^&*()-+"

  bigString := numbers + lower_case + upper_case + special_characters
  chars := []rune(password)
   for i := 0; i < len(chars); i++ {
       if !strings.Contains(bigString,string(chars[i])){
         fmt.Println("String Character invalid")
         os.Exit(0)
       }
   }
   char_numbers := []rune(numbers)
   char_lower_case := []rune(lower_case)
   char_upper_case := []rune(upper_case)
   char_special := []rune(special_characters)
   var count int32 = 0

   for i := 0; i < len(password); i++ {
     if strings.Contains(string(char_numbers[i]),password){
       fmt.Println("Contains numbers")
       count += 1
     }
   }

   for i := 0; i < len(password); i++ {
     if strings.Contains(string(char_lower_case[i]),password){
       fmt.Println("Contains lower case")
       count += 1
     }
   }

   for i := 0; i < len(password); i++ {
     if strings.Contains(string(char_upper_case[i]),password){
       fmt.Println("Contains Upper case")
       count += 1
     }
   }

   for i := 0; i < len(password); i++ {
     if strings.Contains(string(char_special[i]),password){
       fmt.Println("Contains special char")
       count += 1
     }
  }

   return count
}


func main(){
  in := "muthu1234"
  fmt.Println(minimumNumber(8,in))

}
