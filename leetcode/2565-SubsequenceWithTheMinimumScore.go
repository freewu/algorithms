package main

// 2565. Subsequence With the Minimum Score
// You are given two strings s and t.

// You are allowed to remove any number of characters from the string t.

// The score of the string is 0 if no characters are removed from the string t, otherwise:
//     1. Let left be the minimum index among all removed characters.
//     2. Let right be the maximum index among all removed characters.

// Then the score of the string is right - left + 1.

// Return the minimum possible score to make t a subsequence of s.

// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters 
// without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// Example 1:
// Input: s = "abacaba", t = "bzaa"
// Output: 1
// Explanation: In this example, we remove the character "z" at index 1 (0-indexed).
// The string t becomes "baa" which is a subsequence of the string "abacaba" and the score is 1 - 1 + 1 = 1.
// It can be proven that 1 is the minimum score that we can achieve.

// Example 2:
// Input: s = "cde", t = "xyz"
// Output: 3
// Explanation: In this example, we remove characters "x", "y" and "z" at indices 0, 1, and 2 (0-indexed).
// The string t becomes "" which is a subsequence of the string "cde" and the score is 2 - 0 + 1 = 3.
// It can be proven that 3 is the minimum score that we can achieve.

// Constraints:
//     1 <= s.length, t.length <= 10^5
//     s and t consist of only lowercase English letters.

import "fmt"

func minimumScore(s string, t string) int {
    n, m := len(s), len(t)
    suffix := make([]int, n)
    j := m - 1
    for i := n - 1; i >= 0; i-- {
        if 0 <= j && s[i] == t[j] {
            j--
        }
        suffix[i] = j 
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := j + 1
    j = 0
    for i := range(s) {
        res = min(res, max(0, suffix[i] - j + 1))
        if j < m && s[i] == t[j] {
            j++
        }
    }
    return min(res, m - j)
}

func minimumScore1(s string, t string) int {
    n, m := len(s), len(t)
    count := make([]int, n + 1)
    for i, j := n - 1, m - 1; i >= 0; i-- {
        if j >= 0 && s[i] == t[j] {
            count[i] = count[i + 1] + 1
            j--
        } else {
            count[i] = count[i + 1]
        }
    }
    if count[0] >= m { return 0 }
    res := m - count[0]
    for i, j := 0, 0; j < m; i, j = i + 1, j + 1 {
        for i < n && s[i] != t[j] {
            i++
        }
        if i >= n { break  }
        if k := m - (count[i + 1] + j + 1); k < res {
            res = k
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abacaba", t = "bzaa"
    // Output: 1
    // Explanation: In this example, we remove the character "z" at index 1 (0-indexed).
    // The string t becomes "baa" which is a subsequence of the string "abacaba" and the score is 1 - 1 + 1 = 1.
    // It can be proven that 1 is the minimum score that we can achieve.
    fmt.Println(minimumScore("abacaba", "bzaa")) // 1
    // Example 2:
    // Input: s = "cde", t = "xyz"
    // Output: 3
    // Explanation: In this example, we remove characters "x", "y" and "z" at indices 0, 1, and 2 (0-indexed).
    // The string t becomes "" which is a subsequence of the string "cde" and the score is 2 - 0 + 1 = 3.
    // It can be proven that 3 is the minimum score that we can achieve.
    fmt.Println(minimumScore("cde", "xyz")) // 3

    fmt.Println(minimumScore("bluefrog", "leetcode")) // 6
    fmt.Println(minimumScore("leetcode", "bluefrog")) // 8

    fmt.Println(minimumScore1("abacaba", "bzaa")) // 1
    fmt.Println(minimumScore1("cde", "xyz")) // 3
    fmt.Println(minimumScore1("bluefrog", "leetcode")) // 6
    fmt.Println(minimumScore1("leetcode", "bluefrog")) // 8
}