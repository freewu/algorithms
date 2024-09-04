package main

// 1016. Binary String With Substrings Representing 1 To N
// Given a binary string s and a positive integer n, 
// return true if the binary representation of all the integers in the range [1, n] are substrings of s, 
// or false otherwise.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "0110", n = 3
// Output: true

// Example 2:
// Input: s = "0110", n = 4
// Output: false

// Constraints:
//     1 <= s.length <= 1000
//     s[i] is either '0' or '1'.
//     1 <= n <= 10^9

import "fmt"
import "strings"

func queryString(s string, n int) bool {
    for i := 1; i <= n; i++ {
        if !strings.Contains(s, fmt.Sprintf("%b", i)) {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "0110", n = 3
    // Output: true
    fmt.Println(queryString("0110", 3)) // true  00 | 01 | 10 | 11 
    // Example 2:
    // Input: s = "0110", n = 4
    // Output: false
    fmt.Println(queryString("0110", 4)) // false 00 | 01 | 10 | 11 | 101
}