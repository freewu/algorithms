package main

// 2472. Maximum Number of Non-overlapping Palindrome Substrings
// You are given a string s and a positive integer k.

// Select a set of non-overlapping substrings from the string s that satisfy the following conditions:
//     1. The length of each substring is at least k.
//     2. Each substring is a palindrome.

// Return the maximum number of substrings in an optimal selection.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "abaccdbbd", k = 3
// Output: 2
// Explanation: We can select the substrings underlined in s = "abaccdbbd". Both "aba" and "dbbd" are palindromes and have a length of at least k = 3.
// It can be shown that we cannot find a selection with more than two valid substrings.

// Example 2:
// Input: s = "adbcda", k = 2
// Output: 0
// Explanation: There is no palindrome substring of length at least 2 in the string.

// Constraints:
//     1 <= k <= s.length <= 2000
//     s consists of lowercase English letters.

import "fmt"

func maxPalindromes(s string, k int) int {
    res, last, n := 0, -1, len(s)
    helper := func(l, r *int) {
        for *l >= 0 && *r < n && s[*l] == s[*r] && *l > last {
            if *r-*l+1 >= k {
                res++
                last = *r
                break // find the shortest palindrome
            } else {
                *l--
                *r++
            }
        }
    }
    for i := 0; i < n; i++ {
        l, r := i, i // odd length
        helper(&l, &r)
        l, r = i, i + 1 // even length
        helper(&l, &r)
    }
    return res
}

func maxPalindromes1(s string, k int) int {
    isPalindrome := func(s string) bool {
        for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 {
            if s[l] != s[r] {
                return false
            }
        }
        return true
    }
    res, n := 0, len(s)
    for i := 0; i <= n - k; {
        if j := i + k; isPalindrome(s[i:j]) {
            i = j
        } else if j := i + k + 1; j <= n && isPalindrome(s[i:j]) {
            i = j
        } else {
            i++
            continue
        }
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abaccdbbd", k = 3
    // Output: 2
    // Explanation: We can select the substrings underlined in s = "abaccdbbd". Both "aba" and "dbbd" are palindromes and have a length of at least k = 3.
    // It can be shown that we cannot find a selection with more than two valid substrings.
    fmt.Println(maxPalindromes("abaccdbbd", 3)) // 2
    // Example 2:
    // Input: s = "adbcda", k = 2
    // Output: 0
    // Explanation: There is no palindrome substring of length at least 2 in the string.
    fmt.Println(maxPalindromes("adbcda", 2)) // 0

    fmt.Println(maxPalindromes("bluefrog", 2)) // 0
    fmt.Println(maxPalindromes("leetcode", 2)) // 1

    fmt.Println(maxPalindromes1("abaccdbbd", 3)) // 2
    fmt.Println(maxPalindromes1("adbcda", 2)) // 0
    fmt.Println(maxPalindromes1("bluefrog", 2)) // 0
    fmt.Println(maxPalindromes1("leetcode", 2)) // 1
}