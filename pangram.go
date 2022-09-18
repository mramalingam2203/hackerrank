package main

import
(
  "fmt"
  "os"
  "strings"
  "sort"
  "unicode"
)


func removeDuplicateInt(intSlice []int32) []int32 {
    allKeys := make(map[int32]bool)
    list := []int32{}
    for _, item := range intSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func pangram(s string) string {
  s = strings.ReplaceAll(s, " ", "")

  if len(s) < 0 || len(s) >1000{
    fmt.Println("string invalid")
    os.Exit(0)
  }

  if !IsLetter(s) {
    fmt.Println("String corrupt")
    os.Exit(0)
  }

  //  fmt.Println(r)
    rr := []int32(s)
    sort.Slice(rr, func(i, j int) bool {
        return int(rr[i]) < int(rr[j])
    })

    rr = removeDuplicateInt(rr)
    fmt.Print(rr)

    if len(rr) >= 26{
      return "pangram"
    }else{
      return "not pangram"
    }

}

// A-->B : 97-->122
//
func main(){
  my_string := "YaCoVaGAaXxrzUvZ ZaHyacbUZCZUbZxzAYb YefxAwV yyABAdVatYazC TuyYYedxSf aA XBvSg EvYcl badDxvaZXWCyCUAaZvJcyc YdVbDYAdObgc FeCyxpdXxffubDbGAbFBxnzzT WzZ WcBZAaYCgYzseZb PYXbswxchtYIedhyaXtvzVxZSwWBLxxEaAaYAfGzybZzQo AH tCBcszyXZaAgwzYB QdVZBvwzAYbwwAVysxCRdTTT bXzxtWwyXZebEBYNBaDCLbZbwsEAB YTFBAcD bybU axAZAhhay ZkWydfxyGAeAYaZlabazUZssGTBcCXBr dWs XzyZAEzAAZyclC bCGzPfXcCccAFyvazX ZzYAB zAbsuCZADkeWwUuAbaZ zWCtYzgZZBzXXD c VsrbEaG aYYFZJBUlW iXqZxxswaWTJvb Y xuwebj CF zyZYZVdYYdaRyZ bTatyzYZw wVfaZEZauyZ A yo afJeCBAyVDXWBbAxYBzYyiNuWxBbexEcbeVaAqYz XAjawBzEqDzaafz bTaUaAzWYBxXWZzbazbUwYhqCdV DdzyWafztJajczbt CtYVVfzBgdtvXEGyBxy bdZz C xtczZ SZVfW ZCAx aZYDaa cCyZuVUEBGZ ACbawWXdxxLXa EawAgzOABYFbzTf TszVYaDc fACydzZYdAazSaaygBAYbGzdz yyYBYa Vasay xAz A AVWeYXY aTdYCcZIVxxHcWazZyaWaeYZybYUVZ Bu zXwgACWwzXzBaCwAVddb YB aYBBoGXUAcBZPbzVUGtX DeVduXZGXXtOBwagbXZAcDZIDZvTzA yyuUb AZ"
  result := pangram(my_string)
  fmt.Print(result)
  /*
  s1 := "eidbaooo"

    runeSlice := []rune(s1)

    fmt.Println(string(runeSlice))

    sort.Slice(runeSlice, func(i, j int) bool {
        return runeSlice[i] < runeSlice[j]
    })

    fmt.Println(string(runeSlice))
    */
}
