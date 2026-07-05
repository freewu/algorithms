package main

// 3983. Subsequence After One Replacement
// You are given two strings s and t consisting of lowercase English letters.

// You may choose at most one index in s and replace the character at that index with any lowercase English letter.

// Return true if it is possible to make s a subsequence of t; otherwise, return false.

// Example 1:
// Input: s = "cat", t = "chat"
// Output: true
// Explanation:
// Replace s[1] from 'a' to 'h'. The resulting string is "cht".
// "cht" is a subsequence of "chat" because we can match 'c', 'h', and 't' in order.

// Example 2:
// Input: s = "plane", t = "apple"
// Output: false
// Explanation:
// The characters 'p', 'l', and 'e' can be matched in t, but the remaining characters cannot be matched while preserving the required order.
// Even after replacing any one character in s, it is impossible to make s a subsequence of t.

// Constraints:
//     1 <= s.length, t.length <= 10^5
//     s and t consist only of lowercase English letters.

import "fmt"

func canMakeSubsequence(s string, t string) bool {
    n := len(s)
    j0 := 0 // 在不修改的情况下，s 的前缀 [0, j0-1] 是 t 的当前前缀的子序列
    j1 := 0 // 在改过一次的情况下，s 的前缀 [0, j1-1] 是 t 的当前前缀的子序列
    for _, ch := range t {
        // j1 普通匹配
        if s[j1] == byte(ch) {
            j1++
        }
        // 也可以修改 s[j0] 为 ch，强行匹配
        j1 = max(j1, j0+1)
        // j0 普通匹配
        if s[j0] == byte(ch) {
            j0++
        }
        if j0 == n || j1 == n { // s 是 t 的子序列
            return true
        }
    }
    return false
}

func canMakeSubsequence1(s string, t string) bool {
    var check func(s, t string, b bool) bool 
    check = func(s, t string, b bool) bool {
        i, j := 0, 0
        for i < len(s) {
            if len(t) - j < len(s) - i {
                return false
            }
            if s[i] == t[j] {
                i++
                j++
                continue
            }
            if check(s[i:], t[j+1:], b) {
                return true
            }
            if b && check(s[i+1:], t[j+1:], false) {
                return true
            }
            return false
        }
        return true
    }
    return check(s, t, true)
}

func main() {
    // Example 1:
    // Input: s = "cat", t = "chat"
    // Output: true
    // Explanation:
    // Replace s[1] from 'a' to 'h'. The resulting string is "cht".
    // "cht" is a subsequence of "chat" because we can match 'c', 'h', and 't' in order.
    fmt.Println(canMakeSubsequence("cat", "chat")) // true
    // Example 2:
    // Input: s = "plane", t = "apple"
    // Output: false
    // Explanation:
    // The characters 'p', 'l', and 'e' can be matched in t, but the remaining characters cannot be matched while preserving the required order.
    // Even after replacing any one character in s, it is impossible to make s a subsequence of t.
    fmt.Println(canMakeSubsequence("plane", "apple")) // false

    fmt.Println(canMakeSubsequence("bluefrog", "bluefrog")) // true
    fmt.Println(canMakeSubsequence("bluefrog", "leetcode")) // false
    fmt.Println(canMakeSubsequence("bluefrog", "freewu")) // false

    fmt.Println(canMakeSubsequence1("cat", "chat")) // true
    fmt.Println(canMakeSubsequence1("plane", "apple")) // false
    fmt.Println(canMakeSubsequence1("bluefrog", "bluefrog")) // true
    fmt.Println(canMakeSubsequence1("bluefrog", "leetcode")) // false
    fmt.Println(canMakeSubsequence1("bluefrog", "freewu")) // false
}