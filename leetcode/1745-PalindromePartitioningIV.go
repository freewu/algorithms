package main

// 1745. Palindrome Partitioning IV
// Given a string s, return true if it is possible to split the string s into three non-empty palindromic substrings. 
// Otherwise, return false.​​​​​

// A string is said to be palindrome if it the same string when reversed.

// Example 1:
// Input: s = "abcbdd"
// Output: true
// Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.

// Example 2:
// Input: s = "bcbddxy"
// Output: false
// Explanation: s cannot be split into 3 palindromes.

// Constraints:
//     3 <= s.length <= 2000
//     s​​​​​​ consists only of lowercase English letters.

import "fmt"

func checkPartitioning(s string) bool {
    n := len(s)
    dp := make([][]bool, n)
    for i, _ := range dp {
        dp[i] = make([]bool, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            if s[j] == s[i] {
                if j + 1 <= i - 1 {
                    dp[j][i] = dp[j+1][i-1]
                } else {
                    dp[j][i] = true
                }
            } else {
                dp[j][i] = false
            }
        }
    }
    for i := 1; i < n - 1; i++ {
        for j := i; j < n - 1; j++ {
            if dp[0][i-1] && dp[i][j] && dp[j+1][n-1] {
                return true
            }
        }
    }
    return false
}

func checkPartitioning1(s string) bool {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    calcD1 := func(s string) []int {
        l, r, n := 0, -1, len(s)
        d1 := make([]int, n)
        for i := 0; i < n; i++ {
            k := 1
            if r >= i {
                k = min(d1[l+r-i], r-i+1)
            }
            for i-k >= 0 && i+k < n && s[i-k] == s[i+k] {
                k++
            }
            d1[i] = k
            k--
            if i+k > r {
                l, r = i-k, i+k
            }
        }
        return d1
    }
    calcD2 := func(s string) []int {
        l, r, n := 0, -1, len(s)
        d2 := make([]int, n)
        for i := 0; i < n; i++ {
            k := 0
            if r >= i {
                k = min(d2[l+r-i+1], r-i+1)
            }
            for i-k-1 >= 0 && i+k < n && s[i-k-1] == s[i+k] {
                k++
            }
            d2[i] = k
            k--
            if i+k > r {
                l = i - k - 1
                r = i + k
            }
        }
        return d2
    }
    manacher := func(s string) ([]int, []int) { return calcD1(s), calcD2(s) }
    n := len(s)
    d1, d2 := manacher(s)
    isLoop := func(a, b int) bool {
        mid := (a + b) / 2
        if (b - a + 1) % 2 == 0 {
            return d2[mid + 1] * 2 >= b - a + 1
        } else {
            return d1[mid] * 2 - 1 >= b - a + 1
        }
    }
    for i := 0; i < n; i++ {
        if isLoop(0, i) {
            for j := i + 1; j < n - 1; j++ {
                if isLoop(i + 1, j) && isLoop(j + 1, n - 1) {
                    return true
                }
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "abcbdd"
    // Output: true
    // Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.
    fmt.Println(checkPartitioning("abcbdd")) // true
    // Example 2:
    // Input: s = "bcbddxy"
    // Output: false
    // Explanation: s cannot be split into 3 palindromes.
    fmt.Println(checkPartitioning("bcbddxy")) // false

    fmt.Println(checkPartitioning("bluefrog")) // false
    fmt.Println(checkPartitioning("leetcode")) // false

    fmt.Println(checkPartitioning1("abcbdd")) // true
    fmt.Println(checkPartitioning1("bcbddxy")) // false
    fmt.Println(checkPartitioning1("bluefrog")) // false
    fmt.Println(checkPartitioning1("leetcode")) // false
}