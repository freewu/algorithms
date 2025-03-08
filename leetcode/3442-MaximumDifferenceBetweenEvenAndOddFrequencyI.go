package main

// 3442. Maximum Difference Between Even and Odd Frequency I
// You are given a string s consisting of lowercase English letters. 
// Your task is to find the maximum difference between the frequency of two characters in the string such that:
//     1. One of the characters has an even frequency in the string.
//     2. The other character has an odd frequency in the string.

// Return the maximum difference, calculated as the frequency of the character with an odd frequency minus the frequency of the character with an even frequency.

// Example 1:
// Input: s = "aaaaabbc"
// Output: 3
// Explanation:
// The character 'a' has an odd frequency of 5, and 'b' has an even frequency of 2.
// The maximum difference is 5 - 2 = 3.

// Example 2:
// Input: s = "abcabcab"
// Output: 1
// Explanation:
// The character 'a' has an odd frequency of 3, and 'c' has an even frequency of 2.
// The maximum difference is 3 - 2 = 1.

// Constraints:
//     3 <= s.length <= 100
//     s consists only of lowercase English letters.
//     s contains at least one character with an odd frequency and one with an even frequency.

import "fmt"

func maxDifference(s string) int {
    mp := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        mp[s[i]]++
    }
    evenMin, evenMax, oddMin, oddMax := 101, 0, 101, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range mp {
        if v % 2 == 0 {
            evenMax, evenMin = max(evenMax, v), min(evenMin, v)
        } else {
            oddMax, oddMin = max(oddMax, v),  min(oddMin, v)
        }
    }
    return oddMax - evenMin
}

func main() {
    // Example 1:
    // Input: s = "aaaaabbc"
    // Output: 3
    // Explanation:
    // The character 'a' has an odd frequency of 5, and 'b' has an even frequency of 2.
    // The maximum difference is 5 - 2 = 3.
    fmt.Println(maxDifference("aaaaabbc")) // 3
    // Example 2:
    // Input: s = "abcabcab"
    // Output: 1
    // Explanation:
    // The character 'a' has an odd frequency of 3, and 'c' has an even frequency of 2.
    // The maximum difference is 3 - 2 = 1.
    fmt.Println(maxDifference("abcabcab")) // 1

    fmt.Println(maxDifference("bluefrog")) // -100
    fmt.Println(maxDifference("leetcode")) // -98
}