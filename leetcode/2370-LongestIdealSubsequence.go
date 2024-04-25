package main 

// 2370. Longest Ideal Subsequence
// You are given a string s consisting of lowercase letters and an integer k. 
// We call a string t ideal if the following conditions are satisfied:
//     t is a subsequence of the string s.
//     The absolute difference in the alphabet order of every two adjacent letters in t is less than or equal to k.

// Return the length of the longest ideal string.

// A subsequence is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

// Note that the alphabet order is not cyclic. 
// For example, the absolute difference in the alphabet order of 'a' and 'z' is 25, not 1.

// Example 1:
// Input: s = "acfgbd", k = 2
// Output: 4
// Explanation: The longest ideal string is "acbd". The length of this string is 4, so 4 is returned.
// Note that "acfgbd" is not ideal because 'c' and 'f' have a difference of 3 in alphabet order.

// Example 2:
// Input: s = "abcd", k = 3
// Output: 4
// Explanation: The longest ideal string is "abcd". The length of this string is 4, so 4 is returned.
 
// Constraints:
//     1 <= s.length <= 10^5
//     0 <= k <= 25
//     s consists of lowercase English letters.

import "fmt"
import "slices"

func longestIdealString(s string, k int) int {
    abs := func(x int) int { if x < 0 { return -x }; return x }
    max := func (x, y int) int { if x > y { return x }; return y }
    dp := make([][26]int, len(s))
    dp[0][s[0]-'a'] = 1
    for i := 1; i < len(s); i++ {
        dp[i] = dp[i-1]
        for c := range dp[i-1] {
            if diff := int(s[i]-'a')-c; abs(diff) <= k {
                dp[i][s[i]-'a'] = max(dp[i][s[i]-'a'], dp[i-1][c] + 1)
            }
        }
    }
    res := 0
    for _, v := range dp[len(s) - 1] {
        if v > res { 
            res = v 
        }
    }
    return res
}

func longestIdealString1(s string, k int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp := [26]int{}
    for _, c := range s {
        c := int(c - 'a')
        dp[c] = 1 + slices.Max(dp[max(c - k, 0) : min(c + k + 1, 26)])
    }
    return slices.Max(dp[:])
}

func main() {
    // Example 1:
    // Input: s = "acfgbd", k = 2
    // Output: 4
    // Explanation: The longest ideal string is "acbd". The length of this string is 4, so 4 is returned.
    // Note that "acfgbd" is not ideal because 'c' and 'f' have a difference of 3 in alphabet order.
    fmt.Println(longestIdealString("acfgbd", 2)) // 4
    // Example 2:
    // Input: s = "abcd", k = 3
    // Output: 4
    // Explanation: The longest ideal string is "abcd". The length of this string is 4, so 4 is returned.
    fmt.Println(longestIdealString("abcd", 3)) // 4

    fmt.Println(longestIdealString1("acfgbd", 2)) // 4
    fmt.Println(longestIdealString1("abcd", 3)) // 4
}