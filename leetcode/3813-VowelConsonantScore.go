package main

// 3813. Vowel-Consonant Score
// You are given a string s consisting of lowercase English letters, spaces, and digits.

// Let v be the number of vowels in s and c be the number of consonants in s.

// A vowel is one of the letters 'a', 'e', 'i', 'o', or 'u', while any other letter in the English alphabet is considered a consonant.

// The score of the string s is defined as follows:
//     1. If c > 0, the score = floor(v / c) where floor denotes rounding down to the nearest integer.
//     2. Otherwise, the score = 0.

// Return an integer denoting the score of the string.

// Example 1:
// Input: s = "cooear"
// Output: 2
// Explanation:
// The string s = "cooear" contains v = 4 vowels ('o', 'o', 'e', 'a') and c = 2 consonants ('c', 'r').
// The score is floor(v / c) = floor(4 / 2) = 2.

// Example 2:
// Input: s = "axeyizou"
// Output: 1
// Explanation:
// The string s = "axeyizou" contains v = 5 vowels ('a', 'e', 'i', 'o', 'u') and c = 3 consonants ('x', 'y', 'z').
// The score is floor(v / c) = floor(5 / 3) = 1.

// Example 3:
// Input: s = "au 123"
// Output: 0
// Explanation:
// The string s = "au 123" contains no consonants (c = 0), so the score is 0.

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters, spaces and digits.

import "fmt"

func vowelConsonantScore(s string) int {
    isVowel := func(ch byte) bool { return 'a' == ch || 'e' == ch || 'i' == ch || 'o' == ch || 'u' == ch }
    c1, c2 := 0, 0
    for i := range s {
        if isVowel(s[i]) {
            c1++
        } else if s[i] >= 'a' && s[i] <= 'z' {
            c2++
        }
    }
    if c2 == 0 { return 0 }
    return c1 / c2
}

func main() {
    // Example 1:
    // Input: s = "cooear"
    // Output: 2
    // Explanation:
    // The string s = "cooear" contains v = 4 vowels ('o', 'o', 'e', 'a') and c = 2 consonants ('c', 'r').
    // The score is floor(v / c) = floor(4 / 2) = 2.
    fmt.Println(vowelConsonantScore("cooear")) // 2
    // Example 2:
    // Input: s = "axeyizou"
    // Output: 1
    // Explanation:
    // The string s = "axeyizou" contains v = 5 vowels ('a', 'e', 'i', 'o', 'u') and c = 3 consonants ('x', 'y', 'z').
    // The score is floor(v / c) = floor(5 / 3) = 1.
    fmt.Println(vowelConsonantScore("axeyizou")) // 1
    // Example 3:
    // Input: s = "au 123"
    // Output: 0
    // Explanation:
    // The string s = "au 123" contains no consonants (c = 0), so the score is 0.
    fmt.Println(vowelConsonantScore("au 123")) // 0

    fmt.Println(vowelConsonantScore("bluefrog")) // 0
    fmt.Println(vowelConsonantScore("leetcode")) // 1
}