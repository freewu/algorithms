package main

// 1542. Find Longest Awesome Substring
// You are given a string s. An awesome substring is a non-empty substring of s such that we can make any number of swaps in order to make it a palindrome.
// Return the length of the maximum length awesome substring of s.

// Example 1:
// Input: s = "3242415"
// Output: 5
// Explanation: "24241" is the longest awesome substring, we can form the palindrome "24142" with some swaps.

// Example 2:
// Input: s = "12345678"
// Output: 1

// Example 3:
// Input: s = "213123"
// Output: 6
// Explanation: "213123" is the longest awesome substring, we can form the palindrome "231132" with some swaps.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists only of digits.

import "fmt"

func longestAwesome(s string) int {
    res, mask, dp, n := 0, 0, make([]int, 1024), len(s)
    for i,_ := range dp { // Arrays.fill(dp, s.length());
        dp[i] = n
    }
    dp[0] = -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        mask ^= 1 << (s[i] - '0')
        res = max(res, i - dp[mask])
        for j := 0; j <= 9; j++ {
            res = max(res, i - dp[mask ^ (1 << j)])
        }
        dp[mask] = min(dp[mask], i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "3242415"
    // Output: 5
    // Explanation: "24241" is the longest awesome substring, we can form the palindrome "24142" with some swaps.
    fmt.Println(longestAwesome("3242415")) // 5  "24241"
    // Example 2:
    // Input: s = "12345678"
    // Output: 1
    fmt.Println(longestAwesome("12345678")) // 1
    // Example 3:
    // Input: s = "213123"
    // Output: 6
    // Explanation: "213123" is the longest awesome substring, we can form the palindrome "231132" with some swaps.
    fmt.Println(longestAwesome("213123")) // 6 "213123"
}