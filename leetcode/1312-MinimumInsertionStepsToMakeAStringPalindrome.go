package main

// 1312. Minimum Insertion Steps to Make a String Palindrome
// Given a string s. In one step you can insert any character at any index of the string.
// Return the minimum number of steps to make s palindrome.
// A Palindrome String is one that reads the same backward as well as forward.

// Example 1:
// Input: s = "zzazz"
// Output: 0
// Explanation: The string "zzazz" is already palindrome we do not need any insertions.

// Example 2:
// Input: s = "mbadm"
// Output: 2
// Explanation: String can be "mbdadbm" or "mdbabdm".

// Example 3:
// Input: s = "leetcode"
// Output: 5
// Explanation: Inserting 5 characters the string becomes "leetcodocteel".
 
// Constraints:
//     1 <= s.length <= 500
//     s consists of lowercase English letters

import "fmt"

func minInsertions(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]
            } else { // 如果不等于需要调整1步，取最少的步骤 + 1
                dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1 
            }
        }
    }
    return dp[0][n-1]
}

func main() {
    // Explanation: The string "zzazz" is already palindrome we do not need any insertions.
    fmt.Println(minInsertions("zzazz")) // 0
    // Explanation: String can be "mbdadbm" or "mdbabdm".
    fmt.Println(minInsertions("mbadm")) // 2
    // Explanation: Inserting 5 characters the string becomes "leetcodocteel".
    // Explanation: String can be "mbdadbm" or "mdbabdm".
    fmt.Println(minInsertions("leetcode")) // 5
}