package main

// 1771. Maximize Palindrome Length From Subsequences
// You are given two strings, word1 and word2. 
// You want to construct a string in the following manner:
//     Choose some non-empty subsequence subsequence1 from word1.
//     Choose some non-empty subsequence subsequence2 from word2.
//     Concatenate the subsequences: subsequence1 + subsequence2, to make the string.

// Return the length of the longest palindrome that can be constructed in the described manner. 
// If no palindromes can be constructed, return 0.

// A subsequence of a string s is a string that can be made by deleting some (possibly none) characters from s 
// without changing the order of the remaining characters.

// A palindrome is a string that reads the same forward as well as backward.

// Example 1:
// Input: word1 = "cacb", word2 = "cbba"
// Output: 5
// Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba", which is a palindrome.

// Example 2:
// Input: word1 = "ab", word2 = "ab"
// Output: 3
// Explanation: Choose "ab" from word1 and "a" from word2 to make "aba", which is a palindrome.

// Example 3:
// Input: word1 = "aa", word2 = "bb"
// Output: 0
// Explanation: You cannot construct a palindrome from the described method, so return 0.

// Constraints:
//     1 <= word1.length, word2.length <= 1000
//     word1 and word2 consist of lowercase English letters.

import "fmt"

func longestPalindrome(word1 string, word2 string) int {
    s := word1 + word2
    res, n, m := 0, len(word1), len(s)
    dp := make([][]int, m)
    for i := 0 ;i < m; i++ {
        dp[i] = make([]int ,m)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := m - 1; i >= 0; i-- {
        dp[i][i] = 1
        for j := i + 1; j < m; j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i + 1][j - 1] + 2
                if i < n && j >= n && res < dp[i][j] {
                    res = dp[i][j]
                }
            } else {
                dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word1 = "cacb", word2 = "cbba"
    // Output: 5
    // Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba", which is a palindrome.
    fmt.Println(longestPalindrome("cacb","cbba")) // 5
    // Example 2:
    // Input: word1 = "ab", word2 = "ab"
    // Output: 3
    // Explanation: Choose "ab" from word1 and "a" from word2 to make "aba", which is a palindrome.
    fmt.Println(longestPalindrome("ab","ab")) // 3
    // Example 3:
    // Input: word1 = "aa", word2 = "bb"
    // Output: 0
    // Explanation: You cannot construct a palindrome from the described method, so return 0.
    fmt.Println(longestPalindrome("aa","bb")) // 0
}