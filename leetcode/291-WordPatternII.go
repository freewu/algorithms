package main

// 291. Word Pattern II
// Given a pattern and a string s, return true if s matches the pattern.

// A string s matches a pattern if there is some bijective mapping of single characters to non-empty strings such 
// that if each character in pattern is replaced by the string it maps to, then the resulting string is s. 
// A bijective mapping means that no two characters map to the same string, 
// and no character maps to two different strings.

// Example 1:
// Input: pattern = "abab", s = "redblueredblue"
// Output: true
// Explanation: One possible mapping is as follows:
// 'a' -> "red"
// 'b' -> "blue"

// Example 2:
// Input: pattern = "aaaa", s = "asdasdasdasd"
// Output: true
// Explanation: One possible mapping is as follows:
// 'a' -> "asd"

// Example 3:
// Input: pattern = "aabb", s = "xyzabcxzyabc"
// Output: false
 
// Constraints:
//     1 <= pattern.length, s.length <= 20
//     pattern and s consist of only lowercase English letters.

import "fmt"

func wordPatternMatch(pattern string, s string) bool {
    m, n := len(pattern), len(s)
    pm, sm := map[byte]string{}, map[string]byte{}
    var backtrace func(pi, si int) bool
    backtrace = func(pi, si int) bool {
        if pi >= m && si < n { return false }
        if pi < m && si >= n { return false }
        if pi >= m && si >= n { return true }
        res := false
        for i := si; i < n; i++ {
            if v, ok := pm[pattern[pi]]; ok {
                if  v == s[si:i+1] && sm[s[si:i+1]] == pattern[pi] {
                    if backtrace(pi+1, i+1) { return true }
                }
                continue
            }
            if _, ok := sm[s[si:i+1]]; ok { continue }
            pm[pattern[pi]] = s[si:i+1]
            sm[s[si:i+1]] = pattern[pi]
            if backtrace(pi+1, i+1) { return true }
            delete(pm, pattern[pi])
            delete(sm, s[si:i+1])
        }
        return res
    }
    return backtrace(0, 0)
}

func main() {
    // Example 1:
    // Input: pattern = "abab", s = "redblueredblue"
    // Output: true
    // Explanation: One possible mapping is as follows:
    // 'a' -> "red"
    // 'b' -> "blue"
    fmt.Println(wordPatternMatch("abab","redblueredblue")) // true
    // Example 2:
    // Input: pattern = "aaaa", s = "asdasdasdasd"
    // Output: true
    // Explanation: One possible mapping is as follows:
    // 'a' -> "asd"
    fmt.Println(wordPatternMatch("aaaa","asdasdasdasd")) // true
    // Example 3:
    // Input: pattern = "aabb", s = "xyzabcxzyabc"
    // Output: false
    fmt.Println(wordPatternMatch("aabb","xyzabcxzyabc")) // false
}