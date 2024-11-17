package main

// 3084. Count Substrings Starting and Ending with Given Character
// You are given a string s and a character c. 
// Return the total number of substrings of s that start and end with c.

// Example 1:
// Input: s = "abada", c = "a"
// Output: 6
// Explanation: Substrings starting and ending with "a" are: "abada", "abada", "abada", "abada", "abada", "abada".

// Example 2:
// Input: s = "zzz", c = "z"
// Output: 6
// Explanation: There are a total of 6 substrings in s and all start and end with "z".

// Constraints:
//     1 <= s.length <= 10^5
//     s and c consist only of lowercase English letters.

import "fmt"

func countSubstrings(s string, c byte) int64 {
    res, prev := 0, 0
    for i := 0; i < len(s); i++{
        if s[i] == c {
            prev++
            res += prev
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: s = "abada", c = "a"
    // Output: 6
    // Explanation: Substrings starting and ending with "a" are: "abada", "abada", "abada", "abada", "abada", "abada".
    fmt.Println(countSubstrings("abada", 'a')) // 6
    // Example 2:
    // Input: s = "zzz", c = "z"
    // Output: 6
    // Explanation: There are a total of 6 substrings in s and all start and end with "z".
    fmt.Println(countSubstrings("zzz", 'z')) // 6
}