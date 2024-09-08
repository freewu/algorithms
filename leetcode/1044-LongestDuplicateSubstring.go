package main

// 1044. Longest Duplicate Substring
// Given a string s, consider all duplicated substrings: (contiguous) substrings of s that occur 2 or more times. 
// The occurrences may overlap.

// Return any duplicated substring that has the longest possible length. 
// If s does not have a duplicated substring, the answer is "".

// Example 1:
// Input: s = "banana"
// Output: "ana"

// Example 2:
// Input: s = "abcd"
// Output: ""

// Constraints:
//     2 <= s.length <= 3 * 10^4
//     s consists of lowercase English letters.

import "fmt"
import "strings"

// // 二分法
// func longestDupSubstring(s string) string {
//     res, n := "", len(s)
//     l, r := 0, n - 1
//     for l <= r {
//         m := l + (r - l) / 2
//         d := ""
//         tmp := make(map[string]bool, n - m)
//         for i := 0; i + m <= n; i++ {
//             if _, ok := tmp[s[i:i + m]]; !ok {
//                 tmp[s[i:i + m]] = true
//             } else {
//                 d = s[i:i + m]
//                 break
//             }
//         }
//         if d == "" {
//             r = m - 1
//         } else {
//             l = m + 1
//             res = d
//         }
//     }
//     return res
// }

func longestDupSubstring(s string) string {
    res := ""
    for i := 0; i < len(s); i++ {
        if strings.Contains(s[:i], s[i - len(res):i + 1]) {
            res = s[i - len(res):i + 1]
        }
    }
    return res
}

func longestDupSubstring1(s string) string {
    n, P := len(s), 41
    index, l, r := 0, 0, n
    check := func(m int) bool {
        h, p := 0, 1
        for i := 0; i < m; i++ {
            c := int(s[i])
            h = h * P + c
            p *= P
        }
        cnt := map[int]bool{}
        cnt[h] = true
        for i := m; i < n; i++ {
            c := int(s[i])
            h = h * P - int(s[i-m]) * p + c
            if cnt[h] {
                index = i - m + 1
                return true
            }
            cnt[h] = true
        }
        return false
    }
    for l + 1 < r {
        m := (l + r) >> 1
        if check(m) {
            l = m
        } else {
            r = m
        }
    }
    return s[index : index + l]
}

func main() {
    // Example 1:
    // Input: s = "banana"
    // Output: "ana"
    fmt.Println(longestDupSubstring("banana")) // "ana"
    // Example 2:
    // Input: s = "abcd"
    // Output: ""
    fmt.Println(longestDupSubstring("abcd")) // ""

    fmt.Println(longestDupSubstring("aa")) // "a"

    fmt.Println(longestDupSubstring1("banana")) // "ana"
    fmt.Println(longestDupSubstring1("abcd")) // ""
    fmt.Println(longestDupSubstring1("aa")) // "a"
}