package main

// 730. Count Different Palindromic Subsequences
// Given a string s, return the number of different non-empty palindromic subsequences in s. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// A subsequence of a string is obtained by deleting zero or more characters from the string.
// A sequence is palindromic if it is equal to the sequence reversed.
// Two sequences a1, a2, ... and b1, b2, ... are different if there is some i for which ai != bi.

// Example 1:
// Input: s = "bccb"
// Output: 6
// Explanation: The 6 different non-empty palindromic subsequences are 'b', 'c', 'bb', 'cc', 'bcb', 'bccb'.
// Note that 'bcb' is counted only once, even though it occurs twice.

// Example 2:
// Input: s = "abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"
// Output: 104860361
// Explanation: There are 3104860382 different non-empty palindromic subsequences, which is 104860361 modulo 109 + 7.

// Constraints:
//     1 <= s.length <= 1000
//     s[i] is either 'a', 'b', 'c', or 'd'.

import "fmt"

func countPalindromicSubsequences(s string) int {
    n, mod := len(s), 1_000_000_007
    dp, pre, next := make([][]int, n), make([]int, n), make([]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 0; i <= n - 1; i++ {
        find := false
        for j := i + 1; j <= n -1; j++ {
            if s[i] == s[j] {
                next[i], pre[j] = j, i
                find = true
                break
            }
        }
        if !find {
            next[i] = n
        }
    }
    for i := range pre {
        if pre[i] == 0 && s[i] != s[0] {
            pre[i] = -1
        }
    }
    pre[0] = -1
    for i := 0; i <= n-1; i++ {
        dp[i][i] = 1
        for j := i - 1; j >= 0; j-- {
            if s[i] == s[j] {
                s, e := next[j], pre[i]
                if s > e {
                    dp[i][j] = (2 * dp[i-1][j+1] + 2) % mod
                } else if s == e {
                    dp[i][j] = (2 * dp[i-1][j+1] + 1) % mod
                } else {
                    dp[i][j] = (2 * dp[i-1][j+1] - dp[e-1][s+1] + mod) % mod
                }
            } else {
                dp[i][j] = (dp[i][j+1] + dp[i-1][j] - dp[i-1][j+1] + mod) % mod

            }
        }
    }
    return dp[n-1][0] 
} 

func main() {
    // Example 1:
    // Input: s = "bccb"
    // Output: 6
    // Explanation: The 6 different non-empty palindromic subsequences are 'b', 'c', 'bb', 'cc', 'bcb', 'bccb'.
    // Note that 'bcb' is counted only once, even though it occurs twice.
    fmt.Println(countPalindromicSubsequences("bccb")) // 6
    // Example 2:
    // Input: s = "abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"
    // Output: 104860361
    // Explanation: There are 3104860382 different non-empty palindromic subsequences, which is 104860361 modulo 109 + 7.
    fmt.Println(countPalindromicSubsequences("abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba")) // 104860361
}