package main

// 3303. Find the Occurrence of First Almost Equal Substring
// You are given two strings s and pattern.

// A string x is called almost equal to y if you can change at most one character in x to make it identical to y.

// Return the smallest starting index of a substring in s that is almost equal to pattern. 
// If no such index exists, return -1.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "abcdefg", pattern = "bcdffg"
// Output: 1
// Explanation:
// The substring s[1..6] == "bcdefg" can be converted to "bcdffg" by changing s[4] to "f".

// Example 2:
// Input: s = "ababbababa", pattern = "bacaba"
// Output: 4
// Explanation:
// The substring s[4..9] == "bababa" can be converted to "bacaba" by changing s[6] to "c".

// Example 3:
// Input: s = "abcd", pattern = "dba"
// Output: -1

// Example 4:
// Input: s = "dde", pattern = "d"
// Output: 0

// Constraints:
//     1 <= pattern.length < s.length <= 10^5
//     s and pattern consist only of lowercase English letters.

// Follow-up: Could you solve the problem if at most k consecutive characters can be changed?

import "fmt"

func minStartingIndex(s string, pattern string) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    longestPrefix := func(s string, pattern string) []int {
        m, prefix := make([]int, len(pattern)), make([]int, len(pattern))
        m[0], prefix[0] = -1, -1
        curr := -1
        for i := 1; i < len(pattern); i++ {
            prefix[i] = -1
            for pattern[curr+1] != pattern[i] && curr != -1 {
                prefix[i-1-curr] = curr
                curr = m[curr]
            }
            if pattern[curr+1] == pattern[i] {
                curr++
            }
            m[i] = curr
        }
        for curr != -1 {
            prefix[len(pattern)-1-curr] = curr
            curr = m[curr]
        }
        prefix[0], curr = len(pattern) - 1, -1
        res := make([]int, len(s))
        curr = -1
        for j := 0; j < len(res); j++ {
            res[j] = -1
            for curr+1 == len(pattern) || (pattern[curr+1] != s[j] && curr != -1) {
                for i := 0; i < curr-m[curr]; i++ {
                    res[j-1-curr+i] = min(prefix[i], curr-i)
                }
                curr = m[curr]
            }
            if pattern[curr+1] == s[j] {
                curr++
            }
        }
        for curr != -1 {
            for i := 0; i < curr-m[curr]; i++ {
                res[len(s)-1-curr+i] = min(prefix[i], curr-i)
            }
            curr = m[curr]
        }
        return res
    }
    reverse := func(s string) string {
        arr := []byte(s)
        for i := 0; i < len(arr) / 2; i++ {
            arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
        }
        return string(arr)
    }
    arr1, arr2 := longestPrefix(s, pattern), longestPrefix(reverse(s), reverse(pattern))
    for i := 0; i < len(s) - len(pattern) + 1; i++ {
        if arr1[i] + arr2[len(s) - (i + len(pattern))] + 2 >= len(pattern) - 1 {
            return i
        }
    }
    return -1
}

func minStartingIndex1(s, pattern string) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    calc := func(s string) []int {
        n, left, right := len(s), 0, 0 // z-box 左右边界
        res := make([]int, n)
        for i := 1; i < n; i++ {
            if i <= right {
                res[i] = min(res[i - left], right - i + 1)
            }
            for i + res[i] < n && s[res[i]] == s[i + res[i]] {
                left, right = i, i + res[i]
                res[i]++
            }
        }
        return res
    }
    reverse := func(s string) string {
        arr := []byte(s)
        for i := 0; i < len(arr) / 2; i++ {
            arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
        }
        return string(arr)
    }
    prefix, suffix := calc(pattern + s), calc(reverse(pattern) + reverse(s))
    //slices.Reverse(sufZ) // 也可以不反转，下面写 sufZ[len(sufZ)-i]
    n := len(pattern)
    for i := n; i <= len(s); i++ {
        // if prefix[i] + suffix[i - 1] >= n - 1 {
        if prefix[i] + suffix[len(suffix) - i] >= n - 1 {
            return i - n
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: s = "abcdefg", pattern = "bcdffg"
    // Output: 1
    // Explanation:
    // The substring s[1..6] == "bcdefg" can be converted to "bcdffg" by changing s[4] to "f".
    fmt.Println(minStartingIndex("abcdefg", "bcdffg")) // 1
    // Example 2:
    // Input: s = "ababbababa", pattern = "bacaba"
    // Output: 4
    // Explanation:
    // The substring s[4..9] == "bababa" can be converted to "bacaba" by changing s[6] to "c".
    fmt.Println(minStartingIndex("ababbababa", "bacaba")) // 4
    // Example 3:
    // Input: s = "abcd", pattern = "dba"
    // Output: -1
    fmt.Println(minStartingIndex("abcd", "dba")) // -1
    // Example 4:
    // Input: s = "dde", pattern = "d"
    // Output: 0
    fmt.Println(minStartingIndex("dde", "d")) // 0

    fmt.Println(minStartingIndex("bluefrog", "abc")) // -1
    fmt.Println(minStartingIndex("leetcode", "abc")) // -1

    fmt.Println(minStartingIndex1("abcdefg", "bcdffg")) // 1
    fmt.Println(minStartingIndex1("ababbababa", "bacaba")) // 4
    fmt.Println(minStartingIndex1("abcd", "dba")) // -1
    fmt.Println(minStartingIndex1("dde", "d")) // 0
    fmt.Println(minStartingIndex1("bluefrog", "abc")) // -1
    fmt.Println(minStartingIndex1("leetcode", "abc")) // -1
}