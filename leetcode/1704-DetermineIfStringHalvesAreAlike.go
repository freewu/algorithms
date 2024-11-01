package main

// 1704. Determine if String Halves Are Alike
// You are given a string s of even length. 
// Split this string into two halves of equal lengths, and let a be the first half and b be the second half.

// Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). 
// Notice that s contains uppercase and lowercase letters.

// Return true if a and b are alike. Otherwise, return false.

// Example 1:
// Input: s = "book"
// Output: true
// Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.

// Example 2:
// Input: s = "textbook"
// Output: false
// Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
// Notice that the vowel o is counted twice.

// Constraints:
//     2 <= s.length <= 1000
//     s.length is even.
//     s consists of uppercase and lowercase letters.

import "fmt"

func halvesAreAlike(s string) bool {
    isVowel := func(b byte) bool {
        return 'a' == b || 'e' == b || 'i' == b || 'o' == b || 'u' == b || 'A' == b || 'E' == b || 'I' == b || 'O' == b || 'U' == b
    }
    CountVowels := func(s string) int {
        res := 0
        for i := range s {
            if isVowel(s[i]) { res++ }
        }
        return res
    }
    return CountVowels(s[:len(s)/2]) == CountVowels(s[len(s)/2:])
}

func halvesAreAlike1(s string) bool {
    lcount, rcount := 0, 0
    for left, right := 0, len(s) - 1; left < right; {
        switch s[left] {
            case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U': lcount++
        }
        switch s[right] {
            case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U': rcount++
        }
        left++
        right--
    }
    return lcount == rcount
}

func main() {
    // Example 1:
    // Input: s = "book"
    // Output: true
    // Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.
    fmt.Println(halvesAreAlike("book")) // true
    // Example 2:
    // Input: s = "textbook"
    // Output: false
    // Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
    // Notice that the vowel o is counted twice.
    fmt.Println(halvesAreAlike("textbook")) // false

    fmt.Println(halvesAreAlike1("book")) // true
    fmt.Println(halvesAreAlike1("textbook")) // false
}