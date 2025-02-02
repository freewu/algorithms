package main

// 1638. Count Substrings That Differ by One Character
// Given two strings s and t, find the number of ways you can choose a non-empty substring of s 
// and replace a single character by a different character such that the resulting substring is a substring of t. 
// In other words, find the number of substrings in s that differ from some substring in t by exactly one character.

// For example, the underlined substrings in "computer" and "computation" only differ by the 'e'/'a', so this is a valid way.

// Return the number of substrings that satisfy the condition above.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "aba", t = "baba"
// Output: 6
// Explanation: The following are the pairs of substrings from s and t that differ by exactly 1 character:
// ("aba", "baba")
// ("aba", "baba")
// ("aba", "baba")
// ("aba", "baba")
// ("aba", "baba")
// ("aba", "baba")
// The underlined portions are the substrings that are chosen from s and t.

// ​​Example 2:
// Input: s = "ab", t = "bb"
// Output: 3
// Explanation: The following are the pairs of substrings from s and t that differ by 1 character:
// ("ab", "bb")
// ("ab", "bb")
// ("ab", "bb")
// ​​​​The underlined portions are the substrings that are chosen from s and t.

// Constraints:
//     1 <= s.length, t.length <= 100
//     s and t consist of lowercase English letters only.

import "fmt"

func countSubstrings(s string, t string) int {
    res := 0
    helper := func(s, t string, i, j int) int {
        res, pre, curr := 0, 0, 0
        for sn, tn := len(s), len(t); i < sn && j < tn; i, j = i + 1, j + 1 {
            curr++
            if s[i] != t[j] {
                pre = curr
                curr = 0
            }
            res += pre
        }
        return res
    }
    for i := 0; i < len(s); i++ {
        res += helper(s, t, i, 0)
    }
    for j := 1; j < len(t); j++ {
        res += helper(s, t, 0, j)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aba", t = "baba"
    // Output: 6
    // Explanation: The following are the pairs of substrings from s and t that differ by exactly 1 character:
    // ("aba", "baba")
    // ("aba", "baba")
    // ("aba", "baba")
    // ("aba", "baba")
    // ("aba", "baba")
    // ("aba", "baba")
    // The underlined portions are the substrings that are chosen from s and t.
    fmt.Println(countSubstrings("aba","baba")) // 6
    // ​​Example 2:
    // Input: s = "ab", t = "bb"
    // Output: 3
    // Explanation: The following are the pairs of substrings from s and t that differ by 1 character:
    // ("ab", "bb")
    // ("ab", "bb")
    // ("ab", "bb")
    // ​​​​The underlined portions are the substrings that are chosen from s and t.
    fmt.Println(countSubstrings("ab","bb")) // 3
}