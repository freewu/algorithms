package main

// 1961. Check If String Is a Prefix of Array
// Given a string s and an array of strings words, determine whether s is a prefix string of words.

// A string s is a prefix string of words if s can be made by concatenating the first k strings in words for some positive k no larger than words.length.

// Return true if s is a prefix string of words, or false otherwise.

// Example 1:
// Input: s = "iloveleetcode", words = ["i","love","leetcode","apples"]
// Output: true
// Explanation:
// s can be made by concatenating "i", "love", and "leetcode" together.

// Example 2:
// Input: s = "iloveleetcode", words = ["apples","i","love","leetcode"]
// Output: false
// Explanation:
// It is impossible to make s using a prefix of arr.

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 20
//     1 <= s.length <= 1000
//     words[i] and s consist of only lowercase English letters.

import "fmt"

func isPrefixString(s string, words []string) bool {
    index, n := 0, len(s)
    for _, word := range words {
        for i := 0; i < len(word); i++ {
            if index == n { return false }
            if word[i] != s[index] { return false }
            index++
        }
        if index == n { return true }
    }
    return false
}

func isPrefixString1(s string, words []string) bool {
    concatenated := ""
    for i := 0; i < len(words); i++ {
        concatenated += words[i]
        if concatenated == s { return true }
        if len(concatenated) > len(s) { return false }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "iloveleetcode", words = ["i","love","leetcode","apples"]
    // Output: true
    // Explanation:
    // s can be made by concatenating "i", "love", and "leetcode" together.
    fmt.Println(isPrefixString("iloveleetcode", []string{"i","love","leetcode","apples"})) // true
    // Example 2:
    // Input: s = "iloveleetcode", words = ["apples","i","love","leetcode"]
    // Output: false
    // Explanation:
    // It is impossible to make s using a prefix of arr.
    fmt.Println(isPrefixString("iloveleetcode", []string{"apples","i","love","leetcode"})) // false
    fmt.Println(isPrefixString("iloveleetcode", []string{"i","love","leet","codeapples"})) // false

    fmt.Println(isPrefixString1("iloveleetcode", []string{"i","love","leetcode","apples"})) // true
    fmt.Println(isPrefixString1("iloveleetcode", []string{"apples","i","love","leetcode"})) // false
    fmt.Println(isPrefixString1("iloveleetcode", []string{"i","love","leet","codeapples"})) // false
}