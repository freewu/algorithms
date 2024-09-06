package main

// 1698. Number of Distinct Substrings in a String
// Given a string s, return the number of distinct substrings of s.
// A substring of a string is obtained by deleting any number of characters (possibly zero) from the front of the string and any number (possibly zero) from the back of the string.

// Example 1:
// Input: s = "aabbaba"
// Output: 21
// Explanation: The set of distinct strings is ["a","b","aa","bb","ab","ba","aab","abb","bab","bba","aba","aabb","abba","bbab","baba","aabba","abbab","bbaba","aabbab","abbaba","aabbaba"]

// Example 2:
// Input: s = "abcdefg"
// Output: 28

// Constraints:
//     1 <= s.length <= 500
//     s consists of lowercase English letters.

// Follow up: Can you solve this problem in O(n) time complexity?

import "fmt"
import "strings"

func countDistinct(s string) int {
    n := len(s)
    dp := make([]int, n)
    mark := make([]int, 26)  // 标记字母是否出现过
    dp[0] = 1
    mark[s[0]-'a'] = 1
    for i := 1; i < n; i++ {
        if mark[s[i]-'a'] == 0 {  // 第一次出现，则前面长度加当前字母 构成新的子串。
            dp[i] = dp[i-1] + i + 1
            mark[s[i]-'a'] = 1
            continue
        }
        diff := -1
        for j := i; j > -1; j-- {
            if !strings.Contains(s[:i], s[j:i+1]) {
                diff = j + 1
                break
            }
        }
        dp[i] = dp[i-1] + diff
    }
    return dp[len(dp)-1]
}

func main() {
    // Example 1:
    // Input: s = "aabbaba"
    // Output: 21
    // Explanation: The set of distinct strings is ["a","b","aa","bb","ab","ba","aab","abb","bab","bba","aba","aabb","abba","bbab","baba","aabba","abbab","bbaba","aabbab","abbaba","aabbaba"]
    fmt.Println(countDistinct("aabbaba")) // 21
    // Example 2:
    // Input: s = "abcdefg"
    // Output: 28
    fmt.Println(countDistinct("abcdefg")) // 28
}