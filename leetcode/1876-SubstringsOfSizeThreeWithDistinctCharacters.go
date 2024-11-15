package main

// 1876. Substrings of Size Three with Distinct Characters
// A string is good if there are no repeated characters.

// Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.

// Note that if there are multiple occurrences of the same substring, every occurrence should be counted.

// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: s = "xyzzaz"
// Output: 1
// Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz". 
// The only good substring of length 3 is "xyz".

// Example 2:
// Input: s = "aababcabc"
// Output: 4
// Explanation: There are 7 substrings of size 3: "aab", "aba", "bab", "abc", "bca", "cab", and "abc".
// The good substrings are "abc", "bca", "cab", and "abc".

// Constraints:
//     1 <= s.length <= 100
//     s​​​​​​ consists of lowercase English letters.

import "fmt"

func countGoodSubstrings(s string) int {
    res := 0
    for i := 2; i < len(s); i++ {
        if s[i] != s[i - 1] && s[i] != s[i - 2] && s[i - 1] != s[i - 2] { // 三个字符都不相等
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "xyzzaz"
    // Output: 1
    // Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz". 
    // The only good substring of length 3 is "xyz".
    fmt.Println(countGoodSubstrings("xyzzaz")) // 1
    // Example 2:
    // Input: s = "aababcabc"
    // Output: 4
    // Explanation: There are 7 substrings of size 3: "aab", "aba", "bab", "abc", "bca", "cab", and "abc".
    // The good substrings are "abc", "bca", "cab", and "abc".
    fmt.Println(countGoodSubstrings("aababcabc")) // 4
}