package main

// 1682. Longest Palindromic Subsequence II
// A subsequence of a string s is considered a good palindromic subsequence if:
//     It is a subsequence of s.
//     It is a palindrome (has the same value if reversed).
//     It has an even length.
//     No two consecutive characters are equal, except the two middle ones.

// For example, if s = "abcabcabb", then "abba" is considered a good palindromic subsequence, 
// while "bcb" (not even length) and "bbbb" (has equal consecutive characters) are not.

// Given a string s, return the length of the longest good palindromic subsequence in s.

// Example 1:
// Input: s = "bbabab"
// Output: 4
// Explanation: The longest good palindromic subsequence of s is "baab".

// Example 2:
// Input: s = "dcbccacdb"
// Output: 4
// Explanation: The longest good palindromic subsequence of s is "dccd".

// Constraints:
//     1 <= s.length <= 250
//     s consists of lowercase English letters.

import "fmt"

func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp, disChar := make([][]int, n), make([][]byte, n) //一个额外数组存储[i,j]的回文串最外围的字符
    for i := 0; i < n; i++ {
        dp[i], disChar[i] = make([]int, n), make([]byte, n)
    }
    for i := n - 1; i >= 0; i-- { // 为使dp[i][j] = max(dp[i][j-1], dp[i-1][j])永远成立，我们从右下角开始遍历
        for j := i + 1; j < n; j++ {
            // 初始时，j-i==1时disChar[i+1][j-1] == 0肯定满足disChar[i+1][j-1] != s[i]的
            if s[i] == s[j] && disChar[i+1][j-1] != s[i] {
                dp[i][j] = dp[i+1][j-1] + 2 // 与内部回文串最外围的字符不同的字符可以继续转移
                disChar[i][j] = s[i]        // 存储下一次不能使用的字符
            } else {
                // 若不满足 s[i] == s[j] && disChar[i+1][j-1] != s[i] ，就从左或右去掉一个字符的dp转移过来
                if dp[i][j-1] > dp[i+1][j] {
                    dp[i][j] = dp[i][j-1]
                    disChar[i][j] = disChar[i][j-1]
                } else {
                    dp[i][j] = dp[i+1][j]
                    disChar[i][j] = disChar[i+1][j]
                }
            }
        }
    }
    return dp[0][n-1]
}

func main() {
    // Example 1:
    // Input: s = "bbabab"
    // Output: 4
    // Explanation: The longest good palindromic subsequence of s is "baab".
    fmt.Println(longestPalindromeSubseq("bbabab")) // 4
    // Example 2:
    // Input: s = "dcbccacdb"
    // Output: 4
    // Explanation: The longest good palindromic subsequence of s is "dccd".
    fmt.Println(longestPalindromeSubseq("dcbccacdb")) // 4
}