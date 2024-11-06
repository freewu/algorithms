package main

// 1763. Longest Nice Substring
// A string s is nice if, for every letter of the alphabet that s contains, it appears both in uppercase and lowercase. 
// For example, "abABB" is nice because 'A' and 'a' appear, and 'B' and 'b' appear. 
// However, "abA" is not because 'b' appears, but 'B' does not.

// Given a string s, return the longest substring of s that is nice. 
// If there are multiple, return the substring of the earliest occurrence. 
// If there are none, return an empty string.

// Example 1:
// Input: s = "YazaAay"
// Output: "aAa"
// Explanation: "aAa" is a nice string because 'A/a' is the only letter of the alphabet in s, and both 'A' and 'a' appear.
// "aAa" is the longest nice substring.

// Example 2:
// Input: s = "Bb"
// Output: "Bb"
// Explanation: "Bb" is a nice string because both 'B' and 'b' appear. The whole string is a substring.

// Example 3:
// Input: s = "c"
// Output: ""
// Explanation: There are no nice substrings.

// Constraints:
//     1 <= s.length <= 100
//     s consists of uppercase and lowercase English letters.

import "fmt"

func longestNiceSubstring(s string) string {
    res, n := []string{}, len(s)
    helper := func(s string) bool {
        res := make(map[rune]bool)
        for _, v := range s { res[v] = true }
        for _, v := range s {
            if v >= 'A' && v <= 'Z' {
                if v, ok := res[v + 32]; !ok || v == false { return false }
            } else {
                if v, ok := res[v - 32]; !ok || v == false { return false }
            }
        }
        return true
    }
    for i := 0; i < n - 1; i++ {
        for j := i + 2; j <= n; j++ {
            if helper(s[i:j]) { res = append(res, s[i:j]) }
        }
    }
    l := 0
    if len(res) == 0 { return "" }
    if len(res) == 1 { return res[0] }
    for _, v := range res { if len(v) > l { l = len(v) } }
    for _, v := range res { if len(v) == l { return v } }
    return ""
}

func longestNiceSubstring1(s string) string {
    res, n := "", len(s)
    for i := 0; i < n; i++ {
        x, y := 0, 0
        for j := i;j < n; j++ {
            //if unicode.IsLower(rune(s[j])) {
            if s[j] >= 'a' && s[j] <= 'z' {
                x |= (1 << (s[j] - 'a'))
            } else {
                y |= (1 << (s[j] - 'A'))
            }
            if x == y && (j -i + 1) > len(res) {
                res = s[i:j + 1]
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "YazaAay"
    // Output: "aAa"
    // Explanation: "aAa" is a nice string because 'A/a' is the only letter of the alphabet in s, and both 'A' and 'a' appear.
    // "aAa" is the longest nice substring.
    fmt.Println(longestNiceSubstring("YazaAay")) // "aAa"
    // Example 2:
    // Input: s = "Bb"
    // Output: "Bb"
    // Explanation: "Bb" is a nice string because both 'B' and 'b' appear. The whole string is a substring.
    fmt.Println(longestNiceSubstring("Bb")) // "Bb"
    // Example 3:
    // Input: s = "c"
    // Output: ""
    // Explanation: There are no nice substrings.
    fmt.Println(longestNiceSubstring("c")) // ""

    fmt.Println(longestNiceSubstring("abcd")) // ""
    fmt.Println(longestNiceSubstring("aAbBcCdD")) // "aAbBcCdD"

    fmt.Println(longestNiceSubstring1("YazaAay")) // "aAa"
    fmt.Println(longestNiceSubstring1("Bb")) // "Bb"
    fmt.Println(longestNiceSubstring1("c")) // ""
    fmt.Println(longestNiceSubstring1("abcd")) // ""
    fmt.Println(longestNiceSubstring1("aAbBcCdD")) // "aAbBcCdD"
}