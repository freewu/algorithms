package main

// 583. Delete Operation for Two Strings
// Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.
// In one step, you can delete exactly one character in either string.

// Example 1:
// Input: word1 = "sea", word2 = "eat"
// Output: 2
// Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".

// Example 2:
// Input: word1 = "leetcode", word2 = "etco"
// Output: 4

// Constraints:
//     1 <= word1.length, word2.length <= 500
//     word1 and word2 consist of only lowercase English letters.

import "fmt"

func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = 0
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m+1; i++ {
        for j := 0; j < n+1; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = i + j
            } else if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[m][n]
}

func minDistance1(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m + 1)
    for i:= range dp{
        dp[i] = make([]int, n + 1)
    }
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 1; j <= n; j++ {
        dp[0][j] = j
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := range word1 {
        for j := range word2 {
            if word1[i] == word2[j] {
                dp[i+1][j+1] = dp[i][j]
            } else {
                dp[i+1][j+1] = min(dp[i+1][j] + 1, dp[i][j+1] + 1)
            }
        }
    }
    return dp[m][n]
}

func main() {
    // Example 1:
    // Input: word1 = "sea", word2 = "eat"
    // Output: 2
    // Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
    fmt.Println(minDistance("sea", "eat")) // 2
    // Example 2:
    // Input: word1 = "leetcode", word2 = "etco"
    // Output: 4
    fmt.Println(minDistance("leetcode", "etco")) // 4

    fmt.Println(minDistance1("sea", "eat")) // 2
    fmt.Println(minDistance1("leetcode", "etco")) // 4
}