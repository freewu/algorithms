package main

// 3329. Count Substrings With K-Frequency Characters II
// Given a string s and an integer k, 
// return the total number of substrings of s where at least one character appears at least k times.

// Example 1:
// Input: s = "abacb", k = 2
// Output: 4
// Explanation:
// The valid substrings are:
// "aba" (character 'a' appears 2 times).
// "abac" (character 'a' appears 2 times).
// "abacb" (character 'a' appears 2 times).
// "bacb" (character 'b' appears 2 times).

// Example 2:
// Input: s = "abcde", k = 1
// Output: 15
// Explanation:
// All substrings are valid because every character appears at least once.

// Constraints:
//     1 <= s.length <= 3 * 10^5
//     1 <= k <= s.length
//     s consists only of lowercase English letters.

import "fmt"

func numberOfSubstrings(s string, k int) int64 {
    res, i := int64(0), 0
    count := [26]int{}
    for _, v := range s {
        count[v - 'a']++
        for count[v - 'a'] >= k {
            count[s[i] - 'a']--
            i++
        }
        res += int64(i)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abacb", k = 2
    // Output: 4
    // Explanation:
    // The valid substrings are:
    // "aba" (character 'a' appears 2 times).
    // "abac" (character 'a' appears 2 times).
    // "abacb" (character 'a' appears 2 times).
    // "bacb" (character 'b' appears 2 times).
    fmt.Println(numberOfSubstrings("abacb", 2)) // 4
    // Example 2:
    // Input: s = "abcde", k = 1
    // Output: 15
    // Explanation:
    // All substrings are valid because every character appears at least once.
    fmt.Println(numberOfSubstrings("abcde", 1)) // 15 
}