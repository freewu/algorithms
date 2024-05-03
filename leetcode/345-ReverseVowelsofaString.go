package main

// 345. Reverse Vowels of a String
// Given a string s, reverse only all the vowels in the string and return it.
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

// Example 1:
// Input: s = "hello"
// Output: "holle"

// Example 2:
// Input: s = "leetcode"
// Output: "leotcede"

// Constraints:
//     1 <= s.length <= 3 * 10^5
//     s consist of printable ASCII characters.

import "fmt"

func reverseVowels(s string) string {
    if len(s) <= 1 {
        return s
    }
    isVowel := func(ch byte) bool {
        return 'a' == ch || 'e' == ch || 'i' == ch || 'o' == ch || 'u' == ch ||
               'A' == ch || 'E' == ch || 'I' == ch || 'O' == ch || 'U' == ch
    }
    t, left, right :=  []byte(s), 0, len(s) - 1
    for left < right {
        if !isVowel(t[left]) { // find the vowel from begin
            left++
            continue
        }
        if !isVowel(t[right]) { // find the vowel from end
            right--
            continue
        }
        t[left], t[right] = t[right], t[left]
        left++
        right--
    }
    return string(t)
}

func main() {
    // Example 1:
    // Input: s = "hello"
    // Output: "holle"
    fmt.Println(reverseVowels("hello")) // holle
    // Example 2:
    // Input: s = "leetcode"
    // Output: "leotcede"
    fmt.Println(reverseVowels("leetcode")) // leotcede

    fmt.Println(reverseVowels("aA"))    // Aa
    fmt.Println(reverseVowels("abc"))   // abc
    fmt.Println(reverseVowels("ooee"))  // ee
    fmt.Println(reverseVowels("hello")) // holle
}
