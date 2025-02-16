package main

// 3136. Valid Word
// A word is considered valid if:
//     1. It contains a minimum of 3 characters.
//     2. It contains only digits (0-9), and English letters (uppercase and lowercase).
//     3. It includes at least one vowel.
//     4. It includes at least one consonant.

// You are given a string word.

// Return true if word is valid, otherwise, return false.

// Notes:
//     1. 'a', 'e', 'i', 'o', 'u', and their uppercases are vowels.
//     2. A consonant is an English letter that is not a vowel.

// Example 1:
// Input: word = "234Adas"
// Output: true
// Explanation:
// This word satisfies the conditions.

// Example 2:
// Input: word = "b3"
// Output: false
// Explanation:
// The length of this word is fewer than 3, and does not have a vowel.

// Example 3:
// Input: word = "a3$e"
// Output: false
// Explanation:
// This word contains a '$' character and does not have a consonant.

// Constraints:
//     1 <= word.length <= 20
//     word consists of English uppercase and lowercase letters, digits, '@', '#', and '$'.

import "fmt"

func isValid(word string) bool {
    if len(word) < 3 {  return false } // 1. It contains a minimum of 3 characters.
    vowels, consonant := 0, 0
    isVowel := func(ch byte) bool {
        return 'a' == ch || 'e' == ch || 'i' == ch || 'o' == ch || 'u' == ch ||
               'A' == ch || 'E' == ch || 'I' == ch || 'O' == ch || 'U' == ch
    }
    for i := range word {
        if word[i] == '@' || word[i] == '#' || word[i] == '$' {  return false } // 2. It contains only digits (0-9), and English letters (uppercase and lowercase).
        if isVowel(word[i]) {
            vowels++
        } else if (word[i] >= 'a' && word[i] <='z' ) || (word[i] >= 'A' && word[i] <= 'Z') {
            consonant++
        }
    }
    // 3. It includes at least one vowel.
    // 4. It includes at least one consonant.
    return vowels > 0 && consonant > 0
}

func main() {
    // Example 1:
    // Input: word = "234Adas"
    // Output: true
    // Explanation:
    // This word satisfies the conditions.
    fmt.Println(isValid("234Adas")) // true
    // Example 2:
    // Input: word = "b3"
    // Output: false
    // Explanation:
    // The length of this word is fewer than 3, and does not have a vowel.
    fmt.Println(isValid("b3")) // false
    // Example 3:
    // Input: word = "a3$e"
    // Output: false
    // Explanation:
    // This word contains a '$' character and does not have a consonant.
    fmt.Println(isValid("a3$e")) // false

    fmt.Println(isValid("aaaa")) // false
    fmt.Println(isValid("bluefrog")) // true
    fmt.Println(isValid("leetcode")) // true
}