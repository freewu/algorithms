package main

// 2002. Maximum Product of the Length of Two Palindromic Subsequences
// Given a string s, find two disjoint palindromic subsequences of s such that the product of their lengths is maximized. 
// The two subsequences are disjoint if they do not both pick a character at the same index.

// Return the maximum possible product of the lengths of the two palindromic subsequences.

// A subsequence is a string that can be derived from another string by deleting some 
// or no characters without changing the order of the remaining characters. 
// A string is palindromic if it reads the same forward and backward.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/24/two-palindromic-subsequences.png" />
// Input: s = "leetcodecom"
// Output: 9
// Explanation: An optimal solution is to choose "ete" for the 1st subsequence and "cdc" for the 2nd subsequence.
// The product of their lengths is: 3 * 3 = 9.

// Example 2:
// Input: s = "bb"
// Output: 1
// Explanation: An optimal solution is to choose "b" (the first character) for the 1st subsequence and "b" (the second character) for the 2nd subsequence.
// The product of their lengths is: 1 * 1 = 1.

// Example 3:
// Input: s = "accbcaxxcxx"
// Output: 25
// Explanation: An optimal solution is to choose "accca" for the 1st subsequence and "xxcxx" for the 2nd subsequence.
// The product of their lengths is: 5 * 5 = 25.

// Constraints:
//     2 <= s.length <= 12
//     s consists of lowercase English letters only.

import "fmt"

func maxProduct(s string) int {
    m, p := make(map[int]int), make(map[string]bool)
    isPalindrome := func(s string) bool {
        for i := 0; i < len(s) / 2; i++ {
            if s[i] != s[len(s) - i - 1] { return false }
        }
        return true
    }
    var gen func(i int, t string, ti int) bool
    gen = func(i int, t string, ti int) bool {
        if i >= len(s) { return true }
        newT, newTi := t + s[i:i+1], ti | (1 << i)
        if p[newT] || isPalindrome(newT) {
            p[newT] = true
            m[newTi] = len(newT)
        }
        gen(i + 1, newT, newTi) // choose i
        gen(i + 1, t, ti) // not choose i
        return false
    }
    gen(0, "", 0)
    res := 0
    for k1, v1 := range m {
        for k2, v2 := range m {
            if k1 & k2 == 0 && v1 * v2 > res {
                res = v1 * v2
            }
        }
    }
    return res
}

func maxProduct1(s string) int {
    // 题目中要求不使用相同的下标，并且不改变原来的顺序即可
    // 思想：定义dfs(i)表示当前处理的字符选择下标，这样有三种选择 1.给做序列 2.给右序列 3.路边都不给
    n, mx := len(s), 0
    isPalindrome := func(s string) bool {
        for i := 0; i < len(s) / 2; i++ {
            if s[i] != s[len(s) - i - 1] { return false }
        }
        return true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i int, a, b string)
    dfs = func(i int, a, b string) {
        if i == n { // 如果已经分配完了，这个时候就得判断下乘积
            if isPalindrome(a) && isPalindrome(b) { // 两边都是回文就更新一下结果或
                mx = max(mx, len(a) * len(b))
            }
            return
        }
        dfs(i + 1, a + string(s[i]), b) // 1.交给左侧序列
        dfs(i + 1, a, b + string(s[i])) // 2.交给右侧序列
        dfs(i + 1, a, b) // 3.两边都不要
    }
    dfs(0,"","")
    return mx
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/24/two-palindromic-subsequences.png" />
    // Input: s = "leetcodecom"
    // Output: 9
    // Explanation: An optimal solution is to choose "ete" for the 1st subsequence and "cdc" for the 2nd subsequence.
    // The product of their lengths is: 3 * 3 = 9.
    fmt.Println(maxProduct("leetcodecom")) // 9
    // Example 2:
    // Input: s = "bb"
    // Output: 1
    // Explanation: An optimal solution is to choose "b" (the first character) for the 1st subsequence and "b" (the second character) for the 2nd subsequence.
    // The product of their lengths is: 1 * 1 = 1.
    fmt.Println(maxProduct("bb")) // 1
    // Example 3:
    // Input: s = "accbcaxxcxx"
    // Output: 25
    // Explanation: An optimal solution is to choose "accca" for the 1st subsequence and "xxcxx" for the 2nd subsequence.
    // The product of their lengths is: 5 * 5 = 25.
    fmt.Println(maxProduct("accbcaxxcxx")) // 25

    fmt.Println(maxProduct1("leetcodecom")) // 9
    fmt.Println(maxProduct1("bb")) // 1
    fmt.Println(maxProduct1("accbcaxxcxx")) // 25
}