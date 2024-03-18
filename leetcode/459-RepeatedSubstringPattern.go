package main

// 459. Repeated Substring Pattern  
// Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

// Example 1:
// Input: s = "abab"
// Output: true
// Explanation: It is the substring "ab" twice.

// Example 2:
// Input: s = "aba"
// Output: false

// Example 3:
// Input: s = "abcabcabcabc"
// Output: true
// Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
 
// Constraints:
//     1 <= s.length <= 10^4
//     s consists of lowercase English letters.

import "fmt"

import "strings"

func repeatedSubstringPattern(s string) bool {
    l := len(s)
    for i := 1; i <= l/2; i++ {
        if l%i == 0 {
            substring := s[:i]
            var builder strings.Builder
            for j := 0; j < l/i; j++ {
                builder.WriteString(substring)
            }
            if builder.String() == s {
                return true
            }
        }
    }
    return false
}

// best solution
func repeatedSubstringPattern1(s string) bool {
    n := len(s)
    next := make([]int, n)
    for i := range next {
        next[i] = -1
    }
    j := -1
    for i := 1; i < n; i ++ {
        for j >= 0 && s[i] != s[j+1] {
            j = next[j]
        }
        if s[i] == s[j+1] {
            j++
        } 
        next[i] = j
    }
    if next[n-1] != -1 && n % (n - 1 - next[n-1]) == 0 {
        return true
    }
    return false
}

func main() {
    // Explanation: It is the substring "ab" twice.
    fmt.Println(repeatedSubstringPattern("abab")) // true
    fmt.Println(repeatedSubstringPattern("aba")) // false
    // Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
    fmt.Println(repeatedSubstringPattern("abcabcabcabc")) // true
    fmt.Println(repeatedSubstringPattern("abcabcabc")) // true


    // Explanation: It is the substring "ab" twice.
    fmt.Println(repeatedSubstringPattern1("abab")) // true
    fmt.Println(repeatedSubstringPattern1("aba")) // false
    // Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
    fmt.Println(repeatedSubstringPattern1("abcabcabcabc")) // true
    fmt.Println(repeatedSubstringPattern1("abcabcabc")) // true
}