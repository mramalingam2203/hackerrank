package main

import(
  "fmt"
)

func generateIntPermutations(array []int, n int, result *[][]int) {
    if n == 1 {
        dst := make([]int, len(array))
        copy(dst, array[:])
        *result = append(*result, dst)
    } else {
        for i := 0; i < n; i++ {
            generateIntPermutations(array, n-1, result)
            if n%2 == 0 {
                // Golang allow us to do multiple assignments
                array[0], array[n-1] = array[n-1], array[0]
            } else {
                array[i], array[n-1] = array[n-1], array[i]
            }
        }
    }
}

func main(){

  numbers := []int{1,2,4,5,7,8,10}
  var result [][]int
  generateIntPermutations(numbers, 3, &result)
  fmt.Print(result)

}
