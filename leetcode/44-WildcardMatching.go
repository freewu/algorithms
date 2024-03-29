package main

// 44. Wildcard Matching
// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
//     '?' Matches any single character.
//     '*' Matches any sequence of characters (including the empty sequence).

// The matching should cover the entire input string (not partial).

// Example 1:
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".

// Example 2:
// Input: s = "aa", p = "*"
// Output: true
// Explanation: '*' matches any sequence.

// Example 3:
// Input: s = "cb", p = "?a"
// Output: false
// Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.

// Constraints:
//     0 <= s.length, p.length <= 2000
//     s contains only lowercase English letters.
//     p contains only lowercase English letters, '?' or '*'.

import "fmt"

func isMatch(s string, p string) bool {
    m, n := len(p) + 1, len(s) + 1
    res := make([][]bool, m)
    for i := 0; i < m; i++ {
        res[i] = make([]bool, n)
        if i == 0 {
            res[0][0] = true
            continue
        }
        res[i][0] = p[i-1] == '*' && res[i-1][0]
    }

    for r := 1; r < m; r++ {
        for c := 1; c < n; c++ {
            if p[r-1] == '*' {
                res[r][c] = res[r-1][c-1] || res[r][c-1] || res[r-1][c]
                continue
            } else if p[r-1] == '?' {
                res[r][c] = res[r-1][c-1]
                continue
            } else {
                res[r][c] = p[r-1] == s[c-1] && res[r-1][c-1]
            }
        }
    }
    return res[m-1][n-1]
}

func main() {
    // Explanation: "a" does not match the entire string "aa".
    fmt.Println(isMatch("aa","a")) // false
    // Explanation: '*' matches any sequence.
    fmt.Println(isMatch("aa","*")) // true
    // Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
    fmt.Println(isMatch("cb","?a")) // false
}