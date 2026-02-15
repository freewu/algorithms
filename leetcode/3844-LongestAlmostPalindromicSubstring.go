package main

// 3844. Longest Almost-Palindromic Substring
// You are given a string s consisting of lowercase English letters.

// A substring is almost-palindromic if it becomes a palindrome after removing exactly one character from it.

// Return an integer denoting the length of the longest almost-palindromic substring in s.

// Example 1:
// Input: s = "abca"
// Output: 4
// Explanation:
// Choose the substring "abca".
// Remove "abca".
// The string becomes "aba", which is a palindrome.
// Therefore, "abca" is almost-palindromic.

// Example 2:
// Input: s = "abba"
// Output: 4
// Explanation:
// Choose the substring "abba".
// Remove "abba".
// The string becomes "aba", which is a palindrome.
// Therefore, "abba" is almost-palindromic.

// Example 3:
// Input: s = "zzabba"
// Output: 5
// Explanation:
// Choose the substring "zzabba".
// Remove "zabba".
// The string becomes "abba", which is a palindrome.
// Therefore, "zabba" is almost-palindromic.

// Constraints:
//     2 <= s.length <= 2500
//     s consists of only lowercase English letters.

import "fmt"

// 中心扩展法
func almostPalindromic(s string) int {
    res, n := 0, len(s)
    expand := func(l, r int) {
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
        }
        res = max(res, r - l - 1) // [l+1, r-1] 是回文串
    }
    for i := range 2*n - 1 {
        l, r := i/2, (i+1)/2
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
        }
        expand(l-1, r) // 删除 s[l]，继续扩展
        expand(l, r+1) // 删除 s[r]，继续扩展
        if res >= n { // 优化：提前返回答案
            return n
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abca"
    // Output: 4
    // Explanation:
    // Choose the substring "abca".
    // Remove "abca".
    // The string becomes "aba", which is a palindrome.
    // Therefore, "abca" is almost-palindromic.
    fmt.Println(almostPalindromic("abca")) // 4
    // Example 2:
    // Input: s = "abba"
    // Output: 4
    // Explanation:
    // Choose the substring "abba".
    // Remove "abba".
    // The string becomes "aba", which is a palindrome.
    // Therefore, "abba" is almost-palindromic.
    fmt.Println(almostPalindromic("abba")) // 4
    // Example 3:
    // Input: s = "zzabba"
    // Output: 5
    // Explanation:
    // Choose the substring "zzabba".
    // Remove "zabba".
    // The string becomes "abba", which is a palindrome.
    // Therefore, "zabba" is almost-palindromic.
    fmt.Println(almostPalindromic("zzabba")) // 5

    fmt.Println(almostPalindromic("bluefrog")) // 5
    fmt.Println(almostPalindromic("leetcode")) // 5
    fmt.Println(almostPalindromic("freewu")) // 5
}