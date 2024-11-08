package main

// 1784. Check if Binary String Has at Most One Segment of Ones
// Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. 
// Otherwise, return false.

// Example 1:
// Input: s = "1001"
// Output: false
// Explanation: The ones do not form a contiguous segment.

// Example 2:
// Input: s = "110"
// Output: true

// Constraints:
//     1 <= s.length <= 100
//     s[i]​​​​ is either '0' or '1'.
//     s[0] is '1'.

import "fmt"
import "strings"

func checkOnesSegment(s string) bool {
    if s == "1"{  return true }
    arr := []int{}
    for i, v := range s {
        if v == '1' {
            arr = append(arr, i)
        }
    }
    sum, z, a, b := 0, 0, arr[0], arr[len(arr) - 1]
    for i := a; i <= b; i++{
        z += i
    }
    for _, v := range arr {
        sum += v
    }
    return z == sum
}

func checkOnesSegment1(s string) bool {
    return !strings.Contains(s, "01")
}


func checkOnesSegment2(s string) bool {
    for i := 1; i < len(s); i++ {
        if s[i - 1] == '0' && s[i] == '1' { return false }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "1001"
    // Output: false
    // Explanation: The ones do not form a contiguous segment.
    fmt.Println(checkOnesSegment("1001")) // false
    // Example 2:
    // Input: s = "110"
    // Output: true
    fmt.Println(checkOnesSegment("110")) // true

    fmt.Println(checkOnesSegment("1")) // true
    fmt.Println(checkOnesSegment("10")) // true

    fmt.Println(checkOnesSegment1("1001")) // false
    fmt.Println(checkOnesSegment1("110")) // true
    fmt.Println(checkOnesSegment1("1")) // true
    fmt.Println(checkOnesSegment1("10")) // true

    fmt.Println(checkOnesSegment2("1001")) // false
    fmt.Println(checkOnesSegment2("110")) // true
    fmt.Println(checkOnesSegment2("1")) // true
    fmt.Println(checkOnesSegment2("10")) // true
}