package main

// 1347. Minimum Number of Steps to Make Two Strings Anagram
// You are given two strings of the same length s and t. 
// In one step you can choose any character of t and replace it with another character.

// Return the minimum number of steps to make t an anagram of s.

// An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

// Example 1:
// Input: s = "bab", t = "aba"
// Output: 1
// Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.

// Example 2:
// Input: s = "leetcode", t = "practice"
// Output: 5
// Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.

// Example 3:
// Input: s = "anagram", t = "mangaar"
// Output: 0
// Explanation: "anagram" and "mangaar" are anagrams. 

// Constraints:
//     1 <= s.length <= 5 * 10^4
//     s.length == t.length
//     s and t consist of lowercase English letters only.

import "fmt"

func minSteps(s string, t string) int {
    count := make([]int, 26)
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }
    steps := 0
    for _, v := range count {
        if v > 0 {
            steps += v
        }
    }
    return steps
}

func main() {
    // Example 1:
    // Input: s = "bab", t = "aba"
    // Output: 1
    // Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.
    fmt.Println(minSteps("bab", "aba")) // 1
    // Example 2:
    // Input: s = "leetcode", t = "practice"
    // Output: 5
    // Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.
    fmt.Println(minSteps("leetcode", "practice")) // 5
    // Example 3:
    // Input: s = "anagram", t = "mangaar"
    // Output: 0
    // Explanation: "anagram" and "mangaar" are anagrams. 
    fmt.Println(minSteps("anagram", "mangaar")) // 0
}