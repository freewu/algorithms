package main

// 3541. Find Most Frequent Vowel and Consonant
// You are given a string s consisting of lowercase English letters ('a' to 'z').

// Your task is to:
//     1. Find the vowel (one of 'a', 'e', 'i', 'o', or 'u') with the maximum frequency.
//     2. Find the consonant (all other letters excluding vowels) with the maximum frequency.

// Return the sum of the two frequencies.

// Note: If multiple vowels or consonants have the same maximum frequency, you may choose any one of them. 
// If there are no vowels or no consonants in the string, consider their frequency as 0.

// The frequency of a letter x is the number of times it occurs in the string.

// Example 1:
// Input: s = "successes"
// Output: 6
// Explanation:
// The vowels are: 'u' (frequency 1), 'e' (frequency 2). The maximum frequency is 2.
// The consonants are: 's' (frequency 4), 'c' (frequency 2). The maximum frequency is 4.
// The output is 2 + 4 = 6.

// Example 2:
// Input: s = "aeiaeia"
// Output: 3
// Explanation:
// The vowels are: 'a' (frequency 3), 'e' ( frequency 2), 'i' (frequency 2). The maximum frequency is 3.
// There are no consonants in s. Hence, maximum consonant frequency = 0.
// The output is 3 + 0 = 3.

// Constraints:
//     1 <= s.length <= 100
//     s consists of lowercase English letters only.

import "fmt"

func maxFreqSum(s string) int {
    count := make([]int, 26)
    for _, v := range s { // 统计数量
        count[v - 'a']++
    }
    vowels, consonants := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    isVowel := func(c byte) bool { return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' }
    for i, v := range count {
        if isVowel(byte('a' + i)) {
            vowels = max(vowels, v)
        } else {
            consonants = max(consonants, v)
        }
    }
    return vowels + consonants;
}

func maxFreqSum1(s string) int {
    count := make([]int, 26)
    for _, v := range s { // 统计数量
        count[v - 'a']++
    }
    vowels, consonants := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    isVowel := func(c int) bool { return c == 0 || c == 4 || c == 8 || c == 14 || c == 20 }
    for i, v := range count {
        if isVowel(i) {
            vowels = max(vowels, v)
        } else {
            consonants = max(consonants, v)
        }
    }
    // fmt.Println("e:", 'e'- 'a' ); // 4
    // fmt.Println("i:", 'i'- 'a' ); // 8
    // fmt.Println("o:", 'o'- 'a' ); // 14
    // fmt.Println("u:", 'u'- 'a' ); // 20
    return vowels + consonants;
}

func main() {
    // Example 1:
    // Input: s = "successes"
    // Output: 6
    // Explanation:
    // The vowels are: 'u' (frequency 1), 'e' (frequency 2). The maximum frequency is 2.
    // The consonants are: 's' (frequency 4), 'c' (frequency 2). The maximum frequency is 4.
    // The output is 2 + 4 = 6.
    fmt.Println(maxFreqSum("successes")) // 6
    // Example 2:
    // Input: s = "aeiaeia"
    // Output: 3
    // Explanation:
    // The vowels are: 'a' (frequency 3), 'e' ( frequency 2), 'i' (frequency 2). The maximum frequency is 3.
    // There are no consonants in s. Hence, maximum consonant frequency = 0.
    // The output is 3 + 0 = 3.
    fmt.Println(maxFreqSum("aeiaeia")) // 3

    fmt.Println(maxFreqSum("abcdefghijklmnopqrstuvwxyz")) // 2

    fmt.Println(maxFreqSum1("successes")) // 6
    fmt.Println(maxFreqSum1("aeiaeia")) // 3
    fmt.Println(maxFreqSum1("abcdefghijklmnopqrstuvwxyz")) // 2
}