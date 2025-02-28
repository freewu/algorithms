package main

// 3460. Longest Common Prefix After at Most One Removal
// You are given two strings s and t.

// Return the length of the longest common prefix between s and t after removing at most one character from s.

// Note: s can be left without any removal.

// Example 1:
// Input: s = "madxa", t = "madam"
// Output: 4
// Explanation:
// Removing s[3] from s results in "mada", which has a longest common prefix of length 4 with t.

// Example 2:
// Input: s = "leetcode", t = "eetcode"
// Output: 7
// Explanation:
// Removing s[0] from s results in "eetcode", which matches t.

// Example 3:
// Input: s = "one", t = "one"
// Output: 3
// Explanation:
// No removal is needed.

// Example 4:
// Input: s = "a", t = "b"
// Output: 0
// Explanation:
// s and t cannot have a common prefix.

// Constraints:
//     1 <= s.length <= 10^5
//     1 <= t.length <= 10^5
//     s and t contain only lowercase English letters.

import "fmt"

func longestCommonPrefix(s string, t string) int {
    res, i, j, k := 0, 0, 0, 1
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++
            j++
            res++
        } else if k > 0 {
            i++
            k--
        } else {
            break
        }
    }
    return res
}

func longestCommonPrefix1(s string, t string) int {
    check := func(s string, t string, i int, j int) int {
        res := 0
        for i < len(s) && j < len(t) {
            if s[i] != t[j] { break }
            i, j = i + 1, j + 1
            res++
        }
        return res
    }
    i := 0
    for i < len(s) && i < len(t) {
        if s[i] == t[i] {
            i++
        } else {
            return i + check(s, t, i + 1, i)
        }
    }
    return i
}

func main() {
    // Example 1:
    // Input: s = "madxa", t = "madam"
    // Output: 4
    // Explanation:
    // Removing s[3] from s results in "mada", which has a longest common prefix of length 4 with t.
    fmt.Println(longestCommonPrefix("madxa", "madam")) // 4
    // Example 2:
    // Input: s = "leetcode", t = "eetcode"
    // Output: 7
    // Explanation:
    // Removing s[0] from s results in "eetcode", which matches t.
    fmt.Println(longestCommonPrefix("leetcode", "eetcode")) // 7
    // Example 3:
    // Input: s = "one", t = "one"
    // Output: 3
    // Explanation:
    // No removal is needed.
    fmt.Println(longestCommonPrefix("one", "one")) // 3
    // Example 4:
    // Input: s = "a", t = "b"
    // Output: 0
    // Explanation:
    // s and t cannot have a common prefix.
    fmt.Println(longestCommonPrefix("a", "b")) // 0

    fmt.Println(longestCommonPrefix("bluefrog", "leetcode")) // 1

    fmt.Println(longestCommonPrefix1("madxa", "madam")) // 4
    fmt.Println(longestCommonPrefix1("leetcode", "eetcode")) // 7
    fmt.Println(longestCommonPrefix1("one", "one")) // 3
    fmt.Println(longestCommonPrefix1("a", "b")) // 0
    fmt.Println(longestCommonPrefix1("bluefrog", "leetcode")) // 1
}