package main

// 2414. Length of the Longest Alphabetical Continuous Substring
// An alphabetical continuous string is a string consisting of consecutive letters in the alphabet. 
// In other words, it is any substring of the string "abcdefghijklmnopqrstuvwxyz".
//     For example, "abc" is an alphabetical continuous string, while "acb" and "za" are not.

// Given a string s consisting of lowercase letters only, 
// return the length of the longest alphabetical continuous substring.

// Example 1:
// Input: s = "abacaba"
// Output: 2
// Explanation: There are 4 distinct continuous substrings: "a", "b", "c" and "ab".
// "ab" is the longest continuous substring.

// Example 2:
// Input: s = "abcde"
// Output: 5
// Explanation: "abcde" is the longest continuous substring.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only English lowercase letters.

import "fmt"

// sliding window
func longestContinuousSubstring(s string) int {
    res, count := 0, 1
    isContinuous := func(a, b byte) bool { return b == a + 1 } // 是否连续字符
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(s); i++ {
        count = 1
        for i + 1 < len(s) && isContinuous(s[i], s[i+1]) {
            count++
            i++
        }
        res = max(res, count)
    }
    return res
}

func longestContinuousSubstring1(s string) int {
    res, count := 0, 1
    for i := 0; i < len(s)-1; i++ {
        if s[i] + 1 == s[i+1] { // 连续字符
            count++
            res = max(res, count)
        } else { // 断了，重新计数
            count = 1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(res, count)
}

func longestContinuousSubstring2(s string) int {
    res, i, n := 0, 0, len(s)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := 1; j < n; j++ {
        if s[j-1]+1 != s[j] {
            res  = max(res, j-i)
            i = j
        }
    }
    res = max(res, n - i)
    return res
}

func main() {
    // Example 1:
    // Input: s = "abacaba"
    // Output: 2
    // Explanation: There are 4 distinct continuous substrings: "a", "b", "c" and "ab".
    // "ab" is the longest continuous substring.
    fmt.Println(longestContinuousSubstring("abacaba")) // 2
    // Example 2:
    // Input: s = "abcde"
    // Output: 5
    // Explanation: "abcde" is the longest continuous substring.
    fmt.Println(longestContinuousSubstring("abcde")) // 5

    fmt.Println(longestContinuousSubstring1("abacaba")) // 2
    fmt.Println(longestContinuousSubstring1("abcde")) // 5

    fmt.Println(longestContinuousSubstring2("abacaba")) // 2
    fmt.Println(longestContinuousSubstring2("abcde")) // 5
}