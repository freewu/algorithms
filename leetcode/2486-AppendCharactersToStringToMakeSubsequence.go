package main

// 2486. Append Characters to String to Make Subsequence
// You are given two strings s and t consisting of only lowercase English letters.
// Return the minimum number of characters that need to be appended to the end of s so that t becomes a subsequence of s.
// A subsequence is a string that can be derived from another string by deleting some or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "coaching", t = "coding"
// Output: 4
// Explanation: Append the characters "ding" to the end of s so that s = "coachingding".
// Now, t is a subsequence of s ("coachingding").
// It can be shown that appending any 3 characters to the end of s will never make t a subsequence.

// Example 2:
// Input: s = "abcde", t = "a"
// Output: 0
// Explanation: t is already a subsequence of s ("abcde").

// Example 3:
// Input: s = "z", t = "abcde"
// Output: 5
// Explanation: Append the characters "abcde" to the end of s so that s = "zabcde".
// Now, t is a subsequence of s ("zabcde").
// It can be shown that appending any 4 characters to the end of s will never make t a subsequence.
 
// Constraints:
//     1 <= s.length, t.length <= 10^5
//     s and t consist only of lowercase English letters.

import "fmt"

// 双指针
func appendCharacters(s string, t string) int {
    i1, i2 := 0, 0
    for i1 < len(s) && i2 < len(t) {
        if s[i1] == t[i2] {
            i2++
        }
        i1++
    }
    return len(t) - i2
}

func main() {
    // Example 1:
    // Input: s = "coaching", t = "coding"
    // Output: 4
    // Explanation: Append the characters "ding" to the end of s so that s = "coachingding".
    // Now, t is a subsequence of s ("coachingding").
    // It can be shown that appending any 3 characters to the end of s will never make t a subsequence.
    fmt.Println(appendCharacters("coaching","coding")) // 4
    // Example 2:
    // Input: s = "abcde", t = "a"
    // Output: 0
    // Explanation: t is already a subsequence of s ("abcde").
    fmt.Println(appendCharacters("abcde","a")) // 0
    // Example 3:
    // Input: s = "z", t = "abcde"
    // Output: 5
    // Explanation: Append the characters "abcde" to the end of s so that s = "zabcde".
    // Now, t is a subsequence of s ("zabcde").
    // It can be shown that appending any 4 characters to the end of s will never make t a subsequence.
    fmt.Println(appendCharacters("z","abcde")) // 5
}