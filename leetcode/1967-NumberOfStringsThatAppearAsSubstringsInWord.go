package main

// 1967. Number of Strings That Appear as Substrings in Word
// Given an array of strings patterns and a string word, 
// return the number of strings in patterns that exist as a substring in word.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: patterns = ["a","abc","bc","d"], word = "abc"
// Output: 3
// Explanation:
// - "a" appears as a substring in "abc".
// - "abc" appears as a substring in "abc".
// - "bc" appears as a substring in "abc".
// - "d" does not appear as a substring in "abc".
// 3 of the strings in patterns appear as a substring in word.

// Example 2:
// Input: patterns = ["a","b","c"], word = "aaaaabbbbb"
// Output: 2
// Explanation:
// - "a" appears as a substring in "aaaaabbbbb".
// - "b" appears as a substring in "aaaaabbbbb".
// - "c" does not appear as a substring in "aaaaabbbbb".
// 2 of the strings in patterns appear as a substring in word.

// Example 3:
// Input: patterns = ["a","a","a"], word = "ab"
// Output: 3
// Explanation: Each of the patterns appears as a substring in word "ab".

// Constraints:
//     1 <= patterns.length <= 100
//     1 <= patterns[i].length <= 100
//     1 <= word.length <= 100
//     patterns[i] and word consist of lowercase English letters.

import "fmt"
import "strings"

func numOfStrings(patterns []string, word string) int {
    res := 0
    for _, v := range patterns {
        if strings.Index(word, v) >= 0 {
            res++
        }
    }
    return res
}

func numOfStrings1(patterns []string, word string) int {
    res := 0
    for _, v := range patterns {
        if strings.Contains(word, v) {
            res++
        }
    }
    return res
}

func numOfStrings2(patterns []string, word string) int {
    res, n, mp := 0, len(word), make(map[string]bool)
    for i := 1; i <= n; i++ {
        for j := 0; j + i <= n; j++ {
            mp[word[j:j + i]] = true
        }
    }
    for _, p := range patterns { 
        if mp[p] { res++ } 
    }
    return res
}

func main() {
    // Example 1:
    // Input: patterns = ["a","abc","bc","d"], word = "abc"
    // Output: 3
    // Explanation:
    // - "a" appears as a substring in "abc".
    // - "abc" appears as a substring in "abc".
    // - "bc" appears as a substring in "abc".
    // - "d" does not appear as a substring in "abc".
    // 3 of the strings in patterns appear as a substring in word.
    fmt.Println(numOfStrings([]string{"a","abc","bc","d"}, "abc")) // 3
    // Example 2:
    // Input: patterns = ["a","b","c"], word = "aaaaabbbbb"
    // Output: 2
    // Explanation:
    // - "a" appears as a substring in "aaaaabbbbb".
    // - "b" appears as a substring in "aaaaabbbbb".
    // - "c" does not appear as a substring in "aaaaabbbbb".
    // 2 of the strings in patterns appear as a substring in word.
    fmt.Println(numOfStrings([]string{"a","b","c"}, "aaaaabbbbb")) // 2
    // Example 3:
    // Input: patterns = ["a","a","a"], word = "ab"
    // Output: 3
    // Explanation: Each of the patterns appears as a substring in word "ab".
    fmt.Println(numOfStrings([]string{"a","a","a"}, "ab")) // 3

    fmt.Println(numOfStrings1([]string{"a","abc","bc","d"}, "abc")) // 3
    fmt.Println(numOfStrings1([]string{"a","b","c"}, "aaaaabbbbb")) // 2
    fmt.Println(numOfStrings1([]string{"a","a","a"}, "ab")) // 3

    fmt.Println(numOfStrings2([]string{"a","abc","bc","d"}, "abc")) // 3
    fmt.Println(numOfStrings2([]string{"a","b","c"}, "aaaaabbbbb")) // 2
    fmt.Println(numOfStrings2([]string{"a","a","a"}, "ab")) // 3
}