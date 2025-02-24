package main

// 3407. Substring Matching Pattern
// You are given a string s and a pattern string p, where p contains exactly one '*' character.

// The '*' in p can be replaced with any sequence of zero or more characters.

// Return true if p can be made a substring of s, and false otherwise.

// Example 1:
// Input: s = "leetcode", p = "ee*e"
// Output: true
// Explanation:
// By replacing the '*' with "tcod", the substring "eetcode" matches the pattern.

// Example 2:
// Input: s = "car", p = "c*v"
// Output: false
// Explanation:
// There is no substring matching the pattern.

// Example 3:
// Input: s = "luck", p = "u*"
// Output: true
// Explanation:
// The substrings "u", "uc", and "uck" match the pattern.

// Constraints:
//     1 <= s.length <= 50
//     1 <= p.length <= 50 
//     s contains only lowercase English letters.
//     p contains only lowercase English letters and exactly one '*'

import "fmt"
import "strings"

func hasMatch(s string, p string) bool {
    parts := strings.Split(p, "*")
    a, b := parts[0], parts[1]
    for i := 0; i <= len(s) - len(a); i++ {
        cur := s[i:i + len(a)]
        if cur == a {
            return strings.Contains(s[i + len(a):], b)
        }
    }
    return false
}

func hasMatch1(s string, p string) bool {
    ps := strings.Split(p, "*")
    i := strings.Index(s, ps[0])
    if i == -1 { return false }
    if ps[1] == "" { return true }
    ss := s[i+len(ps[0]):]
    return strings.Index(ss, ps[1]) > -1
}

func main() {
    // Example 1:
    // Input: s = "leetcode", p = "ee*e"
    // Output: true
    // Explanation:
    // By replacing the '*' with "tcod", the substring "eetcode" matches the pattern.
    fmt.Println(hasMatch("leetcode", "ee*e")) // true
    // Example 2:
    // Input: s = "car", p = "c*v"
    // Output: false
    // Explanation:
    // There is no substring matching the pattern.
    fmt.Println(hasMatch("car", "c*v")) // false
    // Example 3:
    // Input: s = "luck", p = "u*"
    // Output: true
    // Explanation:
    // The substrings "u", "uc", and "uck" match the pattern.
    fmt.Println(hasMatch("luck", "u*")) // true

    fmt.Println(hasMatch("bluefrog", "u*")) // true
    fmt.Println(hasMatch("pep", "q*")) // false

    fmt.Println(hasMatch1("leetcode", "ee*e")) // true
    fmt.Println(hasMatch1("car", "c*v")) // false
    fmt.Println(hasMatch1("luck", "u*")) // true
    fmt.Println(hasMatch1("bluefrog", "u*")) // true
    fmt.Println(hasMatch1("pep", "q*")) // false
}