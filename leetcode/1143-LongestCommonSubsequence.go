package main

// 1143. Longest Common Subsequence
// Given two strings text1 and text2, return the length of their longest common subsequence.
// If there is no common subsequence, return 0.

// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
//     For example, "ace" is a subsequence of "abcde".

// A common subsequence of two strings is a subsequence that is common to both strings.

// Example 1:
// Input: text1 = "abcde", text2 = "ace" 
// Output: 3  
// Explanation: The longest common subsequence is "ace" and its length is 3.

// Example 2:
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.

// Example 3:
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.

// Constraints:
//     1 <= text1.length, text2.length <= 1000
//     text1 and text2 consist of only lowercase English characters.

import "fmt"

// dp
func longestCommonSubsequence(text1 string, text2 string) int {
    l1, l2 := len(text1), len(text2)
    if l1 == 0 || l2 == 0 {
        return 0
    }
    dp := make([][]int, l1 + 1)
    for i := range dp {
        dp[i] = make([]int, l2 + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 循环 text1
    for i := 1; i < l1 + 1; i++ {
        // 循环 tex2
        for j := 1; j < l2 + 1; j++ {
            // 判断出现相同字符就累加1
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i][j-1], dp[i-1][j])
            }
            // fmt.Println(dp)
        }
    }
    return dp[l1][l2]
}

// best solution
func longestCommonSubsequence1(text1 string, text2 string) int {
    dp := make([][]int, len(text1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(text2) + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= len(text1); i++ {
        for j := 1; j <= len(text2); j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i][j-1], dp[i-1][j])
            }
        }
    }
    return dp[len(text1)][len(text2)]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde","ace")) // 3
	fmt.Println(longestCommonSubsequence("abcde","aec")) // 2
	fmt.Println(longestCommonSubsequence("ababccde","abc")) // 3
	fmt.Println(longestCommonSubsequence("abc","def")) // 0

	fmt.Println(longestCommonSubsequence1("abcde","ace")) // 3
	fmt.Println(longestCommonSubsequence1("abcde","aec")) // 2
	fmt.Println(longestCommonSubsequence1("ababccde","abc")) // 3
	fmt.Println(longestCommonSubsequence1("abc","def")) // 0
}