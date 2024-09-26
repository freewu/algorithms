package main

// 1397. Find All Good Strings
// Given the strings s1 and s2 of size n and the string evil, return the number of good strings.

// A good string has size n, it is alphabetically greater than or equal to s1, 
// it is alphabetically smaller than or equal to s2, and it does not contain the string evil as a substring. 
// Since the answer can be a huge number, return this modulo 10^9 + 7.

// Example 1:
// Input: n = 2, s1 = "aa", s2 = "da", evil = "b"
// Output: 51 
// Explanation: There are 25 good strings starting with 'a': "aa","ac","ad",...,"az". Then there are 25 good strings starting with 'c': "ca","cc","cd",...,"cz" and finally there is one good string starting with 'd': "da". 

// Example 2:
// Input: n = 8, s1 = "leetcode", s2 = "leetgoes", evil = "leet"
// Output: 0 
// Explanation: All strings greater than or equal to s1 and smaller than or equal to s2 start with the prefix "leet", therefore, there is not any good string.

// Example 3:
// Input: n = 2, s1 = "gx", s2 = "gz", evil = "x"
// Output: 2

// Constraints:
//     s1.length == n
//     s2.length == n
//     s1 <= s2
//     1 <= n <= 500
//     1 <= evil.length <= 50
//     All strings consist of lowercase English letters.

import "fmt"
import "strings"

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
    m, mod := len(evil), 1_000_000_007
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, m)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    next := make([]int, m)
    for i := 1; i < m; i++ {
        j := next[i-1]
        for j > 0 && evil[i] != evil[j] {
            j = next[j-1]
        }
        if evil[i] == evil[j] {
            j++
        }
        next[i] = j
    }
    var f func(int, int, bool, string) int
    f = func(i int, pos int, isLimit bool, s string) int {
        res := 0
        if pos == m { return 0}
        if i == n { return 1  }
        if !isLimit {
            p := &dp[i][pos]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        d, up := byte(97), byte(122)
        if isLimit {
            up = s[i]
        }
        for ; d <= up; d++ {
            nxt := pos
            for nxt > 0 && d != evil[nxt] {
                nxt = next[nxt-1]
            }
            // 此处要注意，当 nxt == 0 的时候，会存在 d != evil[nxt] 的情况
            // 若直接 nxt + 1 进入递归，是认为此时的两个字符一定是匹配上了，实际上可能并没有
            if nxt == 0 && d != evil[nxt] {
                nxt = -1
            }
            res += f(i+1, nxt+1, isLimit && d == up, s)
            res %= mod
        }
        return res
    }
    res := f(0, 0, true, s2) - f(0, 0, true, s1)
    if res < 0 { res += mod }
    if strings.Index(s1, evil) == -1 { res += 1 }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, s1 = "aa", s2 = "da", evil = "b"
    // Output: 51 
    // Explanation: There are 25 good strings starting with 'a': "aa","ac","ad",...,"az". Then there are 25 good strings starting with 'c': "ca","cc","cd",...,"cz" and finally there is one good string starting with 'd': "da". 
    fmt.Println(findGoodStrings(2, "aa", "da", "b")) // 51
    // Example 2:
    // Input: n = 8, s1 = "leetcode", s2 = "leetgoes", evil = "leet"
    // Output: 0 
    // Explanation: All strings greater than or equal to s1 and smaller than or equal to s2 start with the prefix "leet", therefore, there is not any good string.
    fmt.Println(findGoodStrings(8, "leetcode", "leetgoes", "leet")) // 0
    // Example 3:
    // Input: n = 2, s1 = "gx", s2 = "gz", evil = "x"
    // Output: 2
    fmt.Println(findGoodStrings(2, "gx", "gz", "x")) // 2
}