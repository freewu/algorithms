package main

// 3713. Longest Balanced Substring I
// You are given a string s consisting of lowercase English letters.

// A substring of s is called balanced if all distinct characters in the substring appear the same number of times.

// Return the length of the longest balanced substring of s.

// Example 1:
// Input: s = "abbac"
// Output: 4
// Explanation:
// The longest balanced substring is "abba" because both distinct characters 'a' and 'b' each appear exactly 2 times.

// Example 2:
// Input: s = "zzabccy"
// Output: 4
// Explanation:
// The longest balanced substring is "zabc" because the distinct characters 'z', 'a', 'b', and 'c' each appear exactly 1 time.​​​​​​​

// Example 3:
// Input: s = "aba"
// Output: 2
// Explanation:
// ​​​​​​​One of the longest balanced substrings is "ab" because both distinct characters 'a' and 'b' each appear exactly 1 time. Another longest balanced substring is "ba".

// Constraints:
//     1 <= s.length <= 1000
//     s consists of lowercase English letters.

import "fmt"

func longestBalanced(s string) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; } 
    for i := range s {
        count := [26]int{}
        mx, kinds := 0, 0
        for j := i; j < len(s); j++ {
            b := s[j] - 'a'
            if count[b] == 0 {
                kinds++
            }
            count[b]++
            mx = max(mx, count[b])
            if mx * kinds == j-i+1 {
                res = max(res, j-i+1)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abbac"
    // Output: 4
    // Explanation:
    // The longest balanced substring is "abba" because both distinct characters 'a' and 'b' each appear exactly 2 times.
    fmt.Println(longestBalanced("abbac")) // 4
    // Example 2:
    // Input: s = "zzabccy"
    // Output: 4
    // Explanation:
    // The longest balanced substring is "zabc" because the distinct characters 'z', 'a', 'b', and 'c' each appear exactly 1 time.​​​​​​​
    fmt.Println(longestBalanced("zzabccy")) // 4
    // Example 3:
    // Input: s = "aba"
    // Output: 2
    // Explanation:
    // ​​​​​​​One of the longest balanced substrings is "ab" because both distinct characters 'a' and 'b' each appear exactly 1 time. Another longest balanced substring is "ba".
    fmt.Println(longestBalanced("aba")) // 2

    fmt.Println(longestBalanced("bluefrog")) // 8
    fmt.Println(longestBalanced("leetcode")) // 5
}