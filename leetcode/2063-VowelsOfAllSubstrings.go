package main

// 2063. Vowels of All Substrings
// Given a string word, return the sum of the number of vowels ('a', 'e', 'i', 'o', and 'u') in every substring of word.

// A substring is a contiguous (non-empty) sequence of characters within a string.

// Note: Due to the large constraints, the answer may not fit in a signed 32-bit integer. 
// Please be careful during the calculations.

// Example 1:
// Input: word = "aba"
// Output: 6
// Explanation: 
// All possible substrings are: "a", "ab", "aba", "b", "ba", and "a".
// - "b" has 0 vowels in it
// - "a", "ab", "ba", and "a" have 1 vowel each
// - "aba" has 2 vowels in it
// Hence, the total sum of vowels = 0 + 1 + 1 + 1 + 1 + 2 = 6.

// Example 2:
// Input: word = "abc"
// Output: 3
// Explanation: 
// All possible substrings are: "a", "ab", "abc", "b", "bc", and "c".
// - "a", "ab", and "abc" have 1 vowel each
// - "b", "bc", and "c" have 0 vowels each
// Hence, the total sum of vowels = 1 + 1 + 1 + 0 + 0 + 0 = 3.

// Example 3:
// Input: word = "ltcd"
// Output: 0
// Explanation: There are no vowels in any substring of "ltcd".

// Constraints:
//     1 <= word.length <= 10^5
//     word consists of lowercase English letters.

import "fmt"

func countVowels(word string) int64 {
    mp := make(map[rune]bool)
    for _, v := range "aeiou" {
        mp[v] = true
    }
    res, n := 0, len(word)
    for i, v := range word {
        if mp[v] {
            res += (i + 1) * (n - i)
        }
    }
    return int64(res)
}

func countVowels1(word string) int64 {
    res, n := 0, len(word)
    isVowel := func(ch byte) bool {  return 'a' == ch || 'e' == ch || 'i' == ch || 'o' == ch || 'u' == ch }
    for i := range word {
        if isVowel(word[i]) {
            res += (i + 1) * (n - i)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: word = "aba"
    // Output: 6
    // Explanation: 
    // All possible substrings are: "a", "ab", "aba", "b", "ba", and "a".
    // - "b" has 0 vowels in it
    // - "a", "ab", "ba", and "a" have 1 vowel each
    // - "aba" has 2 vowels in it
    // Hence, the total sum of vowels = 0 + 1 + 1 + 1 + 1 + 2 = 6.
    fmt.Println(countVowels("aba")) // 6
    // Example 2:
    // Input: word = "abc"
    // Output: 3
    // Explanation: 
    // All possible substrings are: "a", "ab", "abc", "b", "bc", and "c".
    // - "a", "ab", and "abc" have 1 vowel each
    // - "b", "bc", and "c" have 0 vowels each
    // Hence, the total sum of vowels = 1 + 1 + 1 + 0 + 0 + 0 = 3.
    fmt.Println(countVowels("abc")) // 3
    // Example 3:
    // Input: word = "ltcd"
    // Output: 0
    // Explanation: There are no vowels in any substring of "ltcd".
    fmt.Println(countVowels("ltcd")) // 0

    fmt.Println(countVowels1("aba")) // 6
    fmt.Println(countVowels1("abc")) // 3
    fmt.Println(countVowels1("ltcd")) // 0
}