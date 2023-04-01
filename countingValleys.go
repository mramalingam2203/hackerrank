package main

import (
    //  "bufio"
    "fmt"
    //  "io"
    // "os"
    //  "path/filepath"
    //  "strconv"
    //  "strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func IsValid(s string) bool {
    for _, r := range s {
        if r != 'U' || r != 'D' {
            return false
        }
    }
    return true
}

func CheckPath(s string) bool {
    var count_up, count_down int32
    for _, r := range s {
        if r == 'U' {
            count_up += 1
        } else {
            count_down += 1
        }
    }
    if !(count_up == count_down) {
        return false
    } else {
        return true
    }
}

func countingValleys(steps int32, path string) int32 {

    runes := []rune(path)

    sum := 0
    count := -1

    for i := 0; i < len(runes); i++ {
        if runes[i] == 'D' {
            sum -= 1
        } else {
            sum += 1
        }
        if sum == 0 {
            count += 1
        }
    }
    fmt.Println(count)
    return int32(count)

}

func main() {
    countingValleys(8, "UDDDUDUU")

}
