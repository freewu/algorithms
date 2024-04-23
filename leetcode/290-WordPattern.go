package main

// 290. Word Pattern
// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

// Example 1:
// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true

// Example 2:
// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false

// Example 3:
// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false
 
// Constraints:
//     1 <= pattern.length <= 300
//     pattern contains only lower-case English letters.
//     1 <= s.length <= 3000
//     s contains only lowercase English letters and spaces ' '.
//     s does not contain any leading or trailing spaces.
//     All the words in s are separated by a single space.

import "fmt"
import "strings"

func wordPattern(pattern string, s string) bool {
    arr := strings.Split(s, " ")
    if len(arr) != len(pattern) {
        return false
    }
    for i := 0; i < len(pattern) - 1; i++ {
        for j := i + 1; j < len(pattern); j++ {
            if pattern[i] == pattern[j] && arr[i] != arr[j] {
                return false
            } else if pattern[i] != pattern[j] && arr[i] == arr[j] {
                return false
            }
        }
    }
    return true
}

func main() {
    fmt.Println(wordPattern("abba","dog cat cat dog")) // true
    fmt.Println(wordPattern("abba","dog cat cat fish")) // false
    fmt.Println(wordPattern("aaaa","dog cat cat dog")) // false
    fmt.Println(wordPattern("abba","dog dog dog dog")) // false

    fmt.Println(wordPattern("aaa","dog dog dog dog")) // false
}