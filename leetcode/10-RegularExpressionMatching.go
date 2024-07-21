package main

// 10. Regular Expression Matching
// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
//     '.' Matches any single character.​​​​
//     '*' Matches zero or more of the preceding element.
//     The matching should cover the entire input string (not partial).
 
// Example 1:
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".

// Example 2:
// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

// Example 3:
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
 
// Constraints:
//     1 <= s.length <= 20
//     1 <= p.length <= 20
//     s contains only lowercase English letters.
//     p contains only lowercase English letters, '.', and '*'.
//     It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
// https://zhuanlan.zhihu.com/p/407952577

import "fmt"

// 递归
func isMatch(s string, p string) bool {
    if len(p) == 0 {
        return len(s) == 0
    }
    if len(p) == 1 {
        return (len(s) == 1) && (s[0] == p[0] || p[0] == '.')
    }
    if p[1] != '*' {
        if len(s) == 0 {
            return false
        } 
        return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
    }
    for len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
        if (isMatch(s, p[2:])) {
            return true
        }
        s = s[1:]
    }
    return isMatch(s, p[2:])
}

// best solution dp
func isMatch1(s string, p string) bool {
    n, m := len(s), len(p)
    dp := make([][]bool, n + 1)
    for i := 0; i <= n; i ++ {
        dp[i] = make([]bool, m + 1)
    }
    dp[0][0] = true
    match := func(i, j int) bool {
        if i == 0 || j == 0 {
            return false
        }
        return s[i - 1] == p[j - 1] || p[j - 1] == '.'
    }
    for i := 0; i <= n; i ++ {
        for j := 0; j <= m; j ++ {
            if match(i, j) {
                dp[i][j] = dp[i - 1][j - 1]
                continue
            }
            if j == 0 || p[j - 1] != '*' {
                continue
            }
            if match(i, j - 1) {
                dp[i][j] = dp[i][j - 1] || dp[i][j - 2] || dp[i - 1][j]
            } else {
                dp[i][j] = dp[i][j - 2]
            }
        }
    }
    return dp[n][m]
}

func isMatch2(s string, p string) bool {
    n, m := len(s), len(p)
    dp := make([][]bool, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]bool, m+1)
    }
    dp[0][0] = true
    for i := 0; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if p[j-1] == '*' {
                dp[i][j] = dp[i][j] || dp[i][j-2]
                if i >= 1 &&(s[i-1] == p[j-2] || p[j-2] == '.') {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            } else {
                if i >= 1 && (s[i-1] == p[j-1] || p[j-1] == '.') {
                    dp[i][j] = dp[i][j] || dp[i-1][j-1]
                }
            }
        }
    }
    return dp[len(s)][len(p)]
}

func main() {
    fmt.Println(isMatch("aa","a")) // false
    fmt.Println(isMatch("aa","a*")) // true
    fmt.Println(isMatch("ab",".*")) // true

    fmt.Println(isMatch1("aa","a")) // false
    fmt.Println(isMatch1("aa","a*")) // true
    fmt.Println(isMatch1("ab",".*")) // true

    fmt.Println(isMatch2("aa","a")) // false
    fmt.Println(isMatch2("aa","a*")) // true
    fmt.Println(isMatch2("ab",".*")) // true
}