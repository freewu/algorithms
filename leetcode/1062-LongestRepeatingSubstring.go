package main

// 1062. Longest Repeating Substring
// Given a string s, return the length of the longest repeating substrings. 
// If no repeating substring exists, return 0.

// Example 1:
// Input: s = "abcd"
// Output: 0
// Explanation: There is no repeating substring.

// Example 2:
// Input: s = "abbaba"
// Output: 2
// Explanation: The longest repeating substrings are "ab" and "ba", each of which occurs twice.

// Example 3:
// Input: s = "aabcaabdaab"
// Output: 3
// Explanation: The longest repeating substring is "aab", which occurs 3 times.

// Constraints:
//     1 <= s.length <= 2000
//     s consists of lowercase English letters.

import "fmt"
import "sort"

func longestRepeatingSubstring(s string) int {
    res, n := 0, len(s)
    suffix := make([]string, n)
    for i := 1; i <= n; i++ { // 计算后缀数组
        suffix[i-1] = s[n-i:]
    }
    sort.Strings(suffix) // 排序 (前缀相同的字符串一定相邻)
    for i := 1; i < n; i++ {
        j := 0
        for j < len(suffix[i-1]) && j < len(suffix[i]) && suffix[i-1][j] == suffix[i][j] { // 计算相邻的字符串的最长公共前缀
            j++
        }
        if j > res {
            res = j
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcd"
    // Output: 0
    // Explanation: There is no repeating substring.
    fmt.Println(longestRepeatingSubstring("abcd")) // 0
    // Example 2:
    // Input: s = "abbaba"
    // Output: 2
    // Explanation: The longest repeating substrings are "ab" and "ba", each of which occurs twice.
    fmt.Println(longestRepeatingSubstring("abbaba")) // 2
    // Example 3:
    // Input: s = "aabcaabdaab"
    // Output: 3
    // Explanation: The longest repeating substring is "aab", which occurs 3 times.
    fmt.Println(longestRepeatingSubstring("aabcaabdaab")) // 3
}