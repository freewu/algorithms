package main

// 1745. Palindrome Partitioning IV
// Given a string s, return true if it is possible to split the string s into three non-empty palindromic substrings. 
// Otherwise, return false.​​​​​

// A string is said to be palindrome if it the same string when reversed.

// Example 1:
// Input: s = "abcbdd"
// Output: true
// Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.

// Example 2:
// Input: s = "bcbddxy"
// Output: false
// Explanation: s cannot be split into 3 palindromes.

// Constraints:
//     3 <= s.length <= 2000
//     s​​​​​​ consists only of lowercase English letters.

import "fmt"

func checkPartitioning(s string) bool {
    n := len(s)
    dp := make([][]bool, n)
    for i, _ := range dp {
        dp[i] = make([]bool, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            if s[j] == s[i] {
                if j + 1 <= i - 1 {
                    dp[j][i] = dp[j+1][i-1]
                } else {
                    dp[j][i] = true
                }
            } else {
                dp[j][i] = false
            }
        }
    }
    for i := 1; i < n - 1; i++ {
        for j := i; j < n - 1; j++ {
            if dp[0][i-1] && dp[i][j] && dp[j+1][n-1] {
                return true
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "abcbdd"
    // Output: true
    // Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.
    fmt.Println(checkPartitioning("abcbdd")) // true
    // Example 2:
    // Input: s = "bcbddxy"
    // Output: false
    // Explanation: s cannot be split into 3 palindromes.
    fmt.Println(checkPartitioning("bcbddxy")) // false
}