package main

// 1216. Valid Palindrome III
// Given a string s and an integer k, return true if s is a k-palindrome.
// A string is k-palindrome if it can be transformed into a palindrome by removing at most k characters from it.

// Example 1:
// Input: s = "abcdeca", k = 2
// Output: true
// Explanation: Remove 'b' and 'e' characters.

// Example 2:
// Input: s = "abbababa", k = 1
// Output: true

// Constraints:
//     1 <= s.length <= 1000
//     s consists of only lowercase English letters.
//     1 <= k <= s.length

import "fmt"

func isValidPalindrome(s string, k int) bool {
    n := len(s)
    dp := make([][]int, n + 1) // dp[i][j]表示，s的前i位及t的前j位中，一共有多少个相同字符
    for i,_ := range dp{
        dp[i] = make([]int, n+1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == s[n-1-(j-1)] {
                dp[i][j] = dp[i-1][j-1] + 1
            }else{
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return len(s) - dp[n][n] <= k // 当len(s)-dp[len(s)][len(s)] <= k时，说明通过移除最多k个元素，可以将s转变为一个回文字符串
}

func main() {
    // Example 1:
    // Input: s = "abcdeca", k = 2
    // Output: true
    // Explanation: Remove 'b' and 'e' characters.
    fmt.Println(isValidPalindrome("abcdeca", 2)) // true
    // Example 2:
    // Input: s = "abbababa", k = 1
    // Output: true
    fmt.Println(isValidPalindrome("abbababa", 1)) // true
}