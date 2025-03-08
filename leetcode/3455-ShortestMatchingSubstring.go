package main

// 3455. Shortest Matching Substring
// You are given a string s and a pattern string p, where p contains exactly two '*' characters.

// The '*' in p matches any sequence of zero or more characters.

// Return the length of the shortest substring in s that matches p. If there is no such substring, return -1.

// Note: The empty substring is considered valid.

// Example 1:
// Input: s = "abaacbaecebce", p = "ba*c*ce"
// Output: 8
// Explanation:
// The shortest matching substring of p in s is "baecebce".

// Example 2:
// Input: s = "baccbaadbc", p = "cc*baa*adb"
// Output: -1
// Explanation:
// There is no matching substring in s.

// Example 3:
// Input: s = "a", p = "**"
// Output: 0
// Explanation:
// The empty substring is the shortest matching substring.

// Example 4:
// Input: s = "madlogic", p = "*adlogi*"
// Output: 6
// Explanation:
// The shortest matching substring of p in s is "adlogi".

// Constraints:
//     1 <= s.length <= 10^5
//     2 <= p.length <= 10^5
//     s contains only lowercase English letters.
//     p contains only lowercase English letters and exactly two '*'.

import "fmt"
import "strings"

func shortestMatchingSubstring(s string, p string) int {
    parts := strings.FieldsFunc(p, func(c rune) bool { 
        return c == '*' 
    })
    if len(parts) == 0 { 
        return 0 
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := 1 << 31
    for i := 0; i < len(s); i++ {
        first, si, j := -1, i, 0
        for ; j < len(parts); j++ {
            index := strings.Index(s[si:], parts[j])
            if index == -1 {
                if res == 1 << 31 {
                    return -1
                }
                return res
            }
            if first == -1 {
                first = si + index
            }
            si = si + index + len(parts[j])
        }
        if j == len(parts) {
            res = min(res, si - first)
        }
        i = first
    }
    if res == 1 << 31 {
        return -1
    }
    return res
}

func shortestMatchingSubstring1(s string, p string) int {
    p = strings.Replace(p, "**", "*", -1)
    p = strings.Trim(p, "*")
    if p == "" {
        return 0
    }
    next := func(s string) []int {
        res := make([]int, len(s))
        for i := 1; i < len(s); i++ {
            pre := res[i-1]
            for ; pre != 0 && s[i] != s[pre]; pre = res[pre-1] {
            }
            if s[i] == s[pre] {
                res[i] = pre + 1
            }
        }
        return res
    }
    parts := strings.Split(p, "*")
    f := make([][]int, len(parts))
    for pi := range parts {
        pNext := next(parts[pi])
        pre := 0
        for i := range s {
            for pre != 0 && s[i] != parts[pi][pre] {
                pre = pNext[pre-1]
            }
            if s[i] == parts[pi][pre] {
                pre++
            }
            if pre == len(pNext) {
                pre = pNext[pre-1]
                f[pi] = append(f[pi], i)
            }
        }
    }
    res := -1
    ps := make([]int, len(parts))
    for ; ps[0] < len(f[0]); ps[0]++ {
        for j := 1; j < len(ps); j++ {
            for ; ps[j] < len(f[j]) && f[j][ps[j]]-len(parts[j]) < f[j-1][ps[j-1]]; ps[j]++ {
            }
            if ps[j] == len(f[j]) {
                return res
            }
        }
        cur := f[len(f)-1][ps[len(ps)-1]] - f[0][ps[0]] + len(parts[0])
        if res == -1 || res > cur {
            res = cur
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abaacbaecebce", p = "ba*c*ce"
    // Output: 8
    // Explanation:
    // The shortest matching substring of p in s is "baecebce".
    fmt.Println(shortestMatchingSubstring("abaacbaecebce", "ba*c*ce")) // 8
    // Example 2:
    // Input: s = "baccbaadbc", p = "cc*baa*adb"
    // Output: -1
    // Explanation:
    // There is no matching substring in s.
    fmt.Println(shortestMatchingSubstring("baccbaadbc", "cc*baa*adb")) // -1
    // Example 3:
    // Input: s = "a", p = "**"
    // Output: 0
    // Explanation:
    // The empty substring is the shortest matching substring.
    fmt.Println(shortestMatchingSubstring("a", "**")) // 0
    // Example 4:
    // Input: s = "madlogic", p = "*adlogi*"
    // Output: 6
    // Explanation:
    // The shortest matching substring of p in s is "adlogi".
    fmt.Println(shortestMatchingSubstring("madlogic", "*adlogi*")) // 6

    fmt.Println(shortestMatchingSubstring1("abaacbaecebce", "ba*c*ce")) // 8
    fmt.Println(shortestMatchingSubstring1("baccbaadbc", "cc*baa*adb")) // -1
    fmt.Println(shortestMatchingSubstring1("a", "**")) // 0
    fmt.Println(shortestMatchingSubstring1("madlogic", "*adlogi*")) // 6
}