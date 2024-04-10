package main

// 516. Longest Palindromic Subsequence
// Given a string s, find the longest palindromic subsequence's length in s.
// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: s = "bbbab"
// Output: 4
// Explanation: One possible longest palindromic subsequence is "bbbb".

// Example 2:
// Input: s = "cbbd"
// Output: 2
// Explanation: One possible longest palindromic subsequence is "bb".

// Constraints:
//     1 <= s.length <= 1000
//     s consists only of lowercase English letters.

import "fmt"

func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := n - 1; i >= 0; i-- {
        dp[i] = make([]int, n)
        dp[i][i] = 1;
        for j := i+1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = 2 + dp[i + 1][j - 1]
            } else {
                dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
            }
        }
    }
    return dp[0][n-1]
}

func longestPalindromeSubseq1(s string) int {
    l := len(s)
    dp := make([]int, l)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := l - 1; i >= 0; i-- {
        dp[i] = 1
        pre := 0
        for j := i + 1; j < l; j++ {
            tmp := dp[j]
            if s[i] == s[j] {
                dp[j] = pre + 2
            } else {
                dp[j] = max(dp[j-1], dp[j])
            }
            pre = tmp
        }
    }
    return dp[l-1]
}

func main() {
    // Explanation: One possible longest palindromic subsequence is "bbbb".
    fmt.Println(longestPalindromeSubseq("bbbab")) // 4
    // Explanation: One possible longest palindromic subsequence is "bb".
    fmt.Println(longestPalindromeSubseq("cbbd")) // 2

    fmt.Println(longestPalindromeSubseq1("bbbab")) // 4
    fmt.Println(longestPalindromeSubseq1("cbbd")) // 2
}