package main

// 161. One Edit Distance
// Given two strings s and t, return true if they are both one edit distance apart, otherwise return false.
// A string s is said to be one distance apart from a string t if you can:
//     Insert exactly one character into s to get t.
//     Delete exactly one character from s to get t.
//     Replace exactly one character of s with a different character to get t.
 
// Example 1:
// Input: s = "ab", t = "acb"
// Output: true
// Explanation: We can insert 'c' into s to get t.

// Example 2:
// Input: s = "", t = ""
// Output: false
// Explanation: We cannot get t from s by only one step.
 
// Constraints:
//     0 <= s.length, t.length <= 10^4
//     s and t consist of lowercase letters, uppercase letters, and digits.

import "fmt"

func isOneEditDistance(s string, t string) bool {
    // 比较，遇到不匹配时分情况讨论，当前位置是增删改？改动一次后则后面的应该相同
    n, m := len(s), len(t)
    for i := 0; i < m || i < n; i++ {
        if i < n && i < m && s[i] == t[i] { // 不需要改动
            continue
        }
        if (i + 1 <= m && s[i:] == t[i+1:]) || // 增当前位
           (i + 1 <= n && s[i+1:] == t[i:]) || // 删当前位
           (i + 1 <= m && i + 1 <= n && s[i+1:] == t[i+1:]) { // 改当前位
                return true
        } else { // 增删改后仍不匹配，失败
            return false
        }
    }
    return false
}

func isOneEditDistance1(s string, t string) bool {
    check := func (s string, t string) bool {
        if len(s) != len(t) { return false; }
        for i := 0; i < len(s); i++ {
            if s[i] != t[i] { return false; }
        }
        return true
    }
    i, j := 0, 0
    for i < len(s) && j < len(t) {
        if s[i] == t[j] { i++; j++; continue; } // 字符一样则比较下一组
        if check(s[i+1:], t[i+1:]) || check(s[i:], t[i+1:]) || check(s[i+1:], t[i:]) {
            return true
        } else {
            return false
        }
    }
    if len(s) - i == 1 || len(t) - j == 1 { // 处理后面只差一个符不一致的情况
        return true
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "ab", t = "acb"
    // Output: true
    // Explanation: We can insert 'c' into s to get t.
    fmt.Println(isOneEditDistance("ab","acb")) // true
    // Example 2:
    // Input: s = "", t = ""
    // Output: false
    // Explanation: We cannot get t from s by only one step.
    fmt.Println(isOneEditDistance("","")) // false

    fmt.Println(isOneEditDistance1("ab","acb")) // true
    fmt.Println(isOneEditDistance1("","")) // false
}