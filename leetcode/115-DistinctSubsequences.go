package main

// 115. Distinct Subsequences
// Given two strings s and t, return the number of distinct subsequences of s which equals t.
// The test cases are generated so that the answer fits on a 32-bit signed integer.

// Example 1:
// Input: s = "rabbbit", t = "rabbit"
// Output: 3
// Explanation:
// As shown below, there are 3 ways you can generate "rabbit" from s.
//     rabb[b]it
//     rab[b]bit
//     ra[b]bbit

// Example 2:
// Input: s = "babgbag", t = "bag"
// Output: 5
// Explanation:
// As shown below, there are 5 ways you can generate "bag" from s.
//     ba[b]g[bag]
//     ba[bgba]g
//     b[abgb]ag
//     [ba]b[gb]ag
//     [babg]bag

// Constraints:
//     1 <= s.length, t.length <= 1000
//     s and t consist of English letters.

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
//（例如，“ACE” 是 “ABCDE” 的一个子序列，而 “AEC” 不是）题目数据保证答案符合 32 位带符号整数范围。

import "fmt"

// 压缩版 DP
func numDistinct(s string, t string) int {
    dp := make([]int, len(s)+1)
    for i, curT := range t {
        pre := 0
        for j, curS := range s {
            if i == 0 {
                pre = 1
            }
            newDP := dp[j+1]
            if curT == curS {
                dp[j+1] = dp[j] + pre
            } else {
                dp[j+1] = dp[j]
            }
            pre = newDP
        }
    }
    // fmt.Println(dp)
    return dp[len(s)]
}

// 普通 DP
func numDistinct1(s, t string) int {
    m, n := len(s), len(t)
    if m < n {
        return 0
    }
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        dp[i][n] = 1
    }
    //fmt.Println(dp)
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if s[i] == t[j] {
                dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
            } else {
                dp[i][j] = dp[i+1][j]
            }
        }
    }
    //fmt.Println(dp)
    return dp[0][0]
}

func numDistinct2(s string, t string) int {
    m, n := len(s), len(t)
    // dp[i][j]：前 i 个字符的 s 子串中，出现前 j 个字符的 t 子串的次数
    dp := make([][]int, m+1) // 二维dp数组
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i < m+1; i++ { // 遍历dp矩阵，填表
        for j := 0; j < n+1; j++ {
            if j == 0 { // base case
                dp[i][j] = 1
            } else if i == 0 { // base case
                dp[i][j] = 0
            } else { // 递推公式
                if s[i-1] == t[j-1] {
                    dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
                } else {
                    dp[i][j] = dp[i-1][j]
                }
            }
        }
    }
    //fmt.Println(dp)
    return dp[m][n] // 前 sLen 个字符的 s 串中，出现前 tLen 个字符的 t 串的次数
}

func main() {
    fmt.Println(numDistinct("babgbag","bag")) // 5
    fmt.Println(numDistinct("rabbbit","rabbit")) // 3


    fmt.Println(numDistinct1("babgbag","bag")) // 5
    fmt.Println(numDistinct1("rabbbit","rabbit")) // 3

    fmt.Println(numDistinct2("babgbag","bag")) // 5
    fmt.Println(numDistinct2("rabbbit","rabbit")) // 3
}