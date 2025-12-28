package main

// 3735. Lexicographically Smallest String After Reverse II
// You are given a string s of length n consisting of lowercase English letters.

// You must perform exactly one operation by choosing any integer k such that 1 <= k <= n and either:
//     1. reverse the first k characters of s, or
//     2. reverse the last k characters of s.

// Return the lexicographically smallest string that can be obtained after exactly one such operation.

// Example 1:
// Input: s = "dcab"
// Output: "acdb"
// Explanation:
// Choose k = 3, reverse the first 3 characters.
// Reverse "dca" to "acd", resulting string s = "acdb", which is the lexicographically smallest string achievable.

// Example 2:
// Input: s = "abba"
// Output: "aabb"
// Explanation:
// Choose k = 3, reverse the last 3 characters.
// Reverse "bba" to "abb", so the resulting string is "aabb", which is the lexicographically smallest string achievable.

// Example 3:
// Input: s = "zxy"
// Output: "xzy"
// Explanation:
// Choose k = 2, reverse the first 2 characters.
// Reverse "zx" to "xz", so the resulting string is "xzy", which is the lexicographically smallest string achievable.
 
// Constraints:
//     1 <= n == s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"
import "slices"

// 超出时间限制 992 / 999 个通过的测试用例
func lexSmallest(s string) string {
    n := len(s)
    if n == 0 { return s }
    res := s // 初始化为原字符串
    // 考虑反转前k个字符的情况
    for k := 1; k <= n; k++ {
        // 反转前k个字符
        b := []byte(s)
        for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
            b[i], b[j] = b[j], b[i]
        }
        candidate := string(b)
        if candidate < res {
            res = candidate
        }
    }
    // 考虑反转后k个字符的情况
    for k := 1; k <= n; k++ {
        // 反转后k个字符
        b := []byte(s)
        start := n - k
        for i, j := start, n-1; i < j; i, j = i+1, j-1 {
            b[i], b[j] = b[j], b[i]
        }
        candidate := string(b)
        if candidate < res {
            res = candidate
        }
    }
    return res 
}

// 超出时间限制 992 / 999
func lexSmallest1(s string) string {
    n := len(s)
    res := s // k = 1 时，操作不改变 s
    for k := 2; k <= n; k++ {
        t := []byte(s[:k])
        slices.Reverse(t)
        res = min(res, string(t)+s[k:])

        t = []byte(s[n-k:])
        slices.Reverse(t)
        res = min(res, s[:n-k]+string(t))
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "dcab"
    // Output: "acdb"
    // Explanation:
    // Choose k = 3, reverse the first 3 characters.
    // Reverse "dca" to "acd", resulting string s = "acdb", which is the lexicographically smallest string achievable.
    fmt.Println(lexSmallest("dcab")) // "acdb"
    // Example 2:
    // Input: s = "abba"
    // Output: "aabb"
    // Explanation:
    // Choose k = 3, reverse the last 3 characters.
    // Reverse "bba" to "abb", so the resulting string is "aabb", which is the lexicographically smallest string achievable.
    fmt.Println(lexSmallest("abba")) // "aabb"
    // Example 3:
    // Input: s = "zxy"
    // Output: "xzy"
    // Explanation:
    // Choose k = 2, reverse the first 2 characters.
    // Reverse "zx" to "xz", so the resulting string is "xzy", which is the lexicographically smallest string achievable.  
    fmt.Println(lexSmallest("zxy")) // "xzy"

    fmt.Println(lexSmallest("bluefrog")) // "bgorfeul"
    fmt.Println(lexSmallest("leetcode")) // "cteelode"

    fmt.Println(lexSmallest1("dcab")) // "acdb"
    fmt.Println(lexSmallest1("abba")) // "aabb"
    fmt.Println(lexSmallest1("zxy")) // "xzy"
    fmt.Println(lexSmallest1("bluefrog")) // "bgorfeul"
    fmt.Println(lexSmallest1("leetcode")) // "cteelode"
}

