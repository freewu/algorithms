package main

// 712. Minimum ASCII Delete Sum for Two Strings
// Given two strings s1 and s2, return the lowest ASCII sum of deleted characters to make two strings equal.

// Example 1:
// Input: s1 = "sea", s2 = "eat"
// Output: 231
// Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
// Deleting "t" from "eat" adds 116 to the sum.
// At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.

// Example 2:
// Input: s1 = "delete", s2 = "leet"
// Output: 403
// Explanation: Deleting "dee" from "delete" to turn the string into "let",
// adds 100[d] + 101[e] + 101[e] to the sum.
// Deleting "e" from "leet" adds 101[e] to the sum.
// At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
// If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
 
// Constraints:
//     1 <= s1.length, s2.length <= 1000
//     s1 and s2 consist of lowercase English letters.

import "fmt"

func minimumDeleteSum(s1 string, s2 string) int {
    l1, l2 := len(s1), len(s2)
    if l1 < l2 {
        return minimumDeleteSum(s2,s1)
    }
    // s2 with len=l2 is the shortest string
    dp := [2][]int{}
    dp[0], dp[1] = make([]int, l2 + 1), make([]int, l2 + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i <= l1; i++ {
        dp[0], dp[1] = dp[1], dp[0]
        for j := 0; j <= l2; j++ {
            if i == 0 && j == 0 { 
                continue 
            }
            dp[1][j] = 10_000_000_000
            if i > 0 { 
                dp[1][j] = min(dp[1][j], dp[0][j] + int(s1[i-1])) 
            }
            if j > 0 { 
                dp[1][j] = min(dp[1][j], dp[1][j-1] + int(s2[j-1])) 
            }
            if i > 0 && j > 0 && s1[i-1] == s2[j-1] { 
                dp[1][j] = min(dp[1][j], dp[0][j-1]) 
            }
        }
    }
    return dp[1][l2]
}

func minimumDeleteSum1(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := 0; i < m + 1; i++ {
        dp[i] = make([]int, n + 1)
    }
    for i := 1; i < m + 1; i++ {
        dp[i][0] = dp[i-1][0] + int(text1[i-1])
    }
    for j := 1; j < n + 1; j++ {
        dp[0][j] = dp[0][j-1] + int(text2[j-1])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i][j-1] + int(text2[j-1]), dp[i-1][j] + int(text1[i-1]))
            }
        }
    }
    return dp[m][n]
}


func main() {
    // Explanation: Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
    // Deleting "t" from "eat" adds 116 to the sum.
    // At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
    fmt.Println(minimumDeleteSum("sea","eat")) // 231
    // Explanation: Deleting "dee" from "delete" to turn the string into "let",
    // adds 100[d] + 101[e] + 101[e] to the sum.
    // Deleting "e" from "leet" adds 101[e] to the sum.
    // At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
    // If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417, which are higher.
    fmt.Println(minimumDeleteSum("delete","leet")) // 430

    fmt.Println(minimumDeleteSum("bluefrog","leetcode")) // 1051
    fmt.Println(minimumDeleteSum("leetcode","bluefrog")) // 1051
    fmt.Println(minimumDeleteSum("bluefrog","freewu")) // 1076
    fmt.Println(minimumDeleteSum("leetcode","freewu")) // 1087

    fmt.Println(minimumDeleteSum1("sea","eat")) // 231
    fmt.Println(minimumDeleteSum1("delete","leet")) // 430
    fmt.Println(minimumDeleteSum1("bluefrog","leetcode")) // 1051
    fmt.Println(minimumDeleteSum1("leetcode","bluefrog")) // 1051
    fmt.Println(minimumDeleteSum1("bluefrog","freewu")) // 1076
    fmt.Println(minimumDeleteSum("leetcode","freewu")) // 1087
}