package main

// 3083. Existence of a Substring in a String and Its Reverse
// Given a string s, find any substring of length 2 which is also present in the reverse of s.

// Return true if such a substring exists, and false otherwise.

// Example 1:
// Input: s = "leetcode"
// Output: true
// Explanation: Substring "ee" is of length 2 which is also present in reverse(s) == "edocteel".

// Example 2:
// Input: s = "abcba"
// Output: true
// Explanation: All of the substrings of length 2 "ab", "bc", "cb", "ba" are also present in reverse(s) == "abcba".

// Example 3:
// Input: s = "abcd"
// Output: false
// Explanation: There is no substring of length 2 in s, which is also present in the reverse of s.

// Constraints:
//     1 <= s.length <= 100
//     s consists only of lowercase English letters.

import "fmt"
import "strings"

func isSubstringPresent(s string) bool {
    reverse := ""
    for i := len(s) - 1; i >= 0; i-- {// 翻转
        reverse += string(s[i])
    }
    for i := 0; i < len(reverse) - 1; i++ {
        if strings.Contains(s, reverse[i:i+2]) { // 判断是否包含其中一部分
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "leetcode"
    // Output: true
    // Explanation: Substring "ee" is of length 2 which is also present in reverse(s) == "edocteel".
    fmt.Println(isSubstringPresent("leetcode")) // true
    // Example 2:
    // Input: s = "abcba"
    // Output: true
    // Explanation: All of the substrings of length 2 "ab", "bc", "cb", "ba" are also present in reverse(s) == "abcba".
    fmt.Println(isSubstringPresent("abcba")) // true
    // Example 3:
    // Input: s = "abcd"
    // Output: false
    // Explanation: There is no substring of length 2 in s, which is also present in the reverse of s.
    fmt.Println(isSubstringPresent("abcd")) // false
}