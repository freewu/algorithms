package main

// 680. Valid Palindrome II
// Given a string s, return true if the s can be palindrome after deleting at most one character from it.

// Example 1:
// Input: s = "aba"
// Output: true

// Example 2:
// Input: s = "abca"
// Output: true
// Explanation: You could delete the character 'c'.

// Example 3:
// Input: s = "abc"
// Output: false
 
// Constraints:
// 1 <= s.length <= 10^5
// s consists of lowercase English letters.

import "fmt"

// func validPalindrome(s string) bool {
//     odd, mp := 0, make(map[byte]int)
//     for i := 0; i < len(s); i++ {
//         mp[s[i]]++
//     }
//     for _, v := range mp {
//         if v % 2 == 1 {
//             odd++
//             if odd > 2 {
//                 return false
//             }
//         }
//     }
//     return true
// }

func validPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    validSubstring := func (s string) bool {
        i,j := 0, len(s) - 1
        for i < j {
            if s[i] == s[j] {
                i++
                j--
            } else {
                return false
            }
        }
        return true
    }
    for i < j {
        if s[i] == s[j] {
            i++
            j--
        } else {
            return validSubstring(s[i+1:j+1]) || validSubstring(s[i:j-1+1])
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "aba"
    // Output: true
    fmt.Println(validPalindrome("aba")) // true
    // Example 2:
    // Input: s = "abca"
    // Output: true
    // Explanation: You could delete the character 'c'.
    fmt.Println(validPalindrome("abca")) // true
    // Example 3:
    // Input: s = "abc"
    // Output: false
    fmt.Println(validPalindrome("abc")) // false

    fmt.Println(validPalindrome("tebbem")) // true
}