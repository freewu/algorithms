package main

// 面试题 01.09. String Rotation LCCI
// Given two strings, s1 and s2, 
// write code to check if s2 is a rotation of s1 (e.g.,"waterbottle" is a rotation of"erbottlewat"). 
// Can you use only one call to the method that checks if one word is a substring of another?

// Example 1:
// Input: s1 = "waterbottle", s2 = "erbottlewat"
// Output: True

// Example 2:
// Input: s1 = "aa", s2 = "aba"
// Output: False

// Note:
//     0 <= s1.length, s2.length <= 100000

import "fmt"
import "strings"

func isFlipedString(s1 string, s2 string) bool {
    return len(s1) == len(s2) && strings.Contains(s1 + s1, s2)
}

func isFlipedString1(s1 string, s2 string) bool {
    if len(s1) != len(s2) { return false }
    if s1 == s2 { return true }
    for i := 0; i < len(s1); i++ {
        s3 := s1[i:] + s1[0:i]
        if s3 == s2 {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s1 = "waterbottle", s2 = "erbottlewat"
    // Output: True
    fmt.Println(isFlipedString("waterbottle", "erbottlewat")) // true
    // Example 2:
    // Input: s1 = "aa", s2 = "aba"
    // Output: False
    fmt.Println(isFlipedString("aa", "aba")) // false

    fmt.Println(isFlipedString("bluefrog", "leetcode")) // false

    fmt.Println(isFlipedString1("waterbottle", "erbottlewat")) // true
    fmt.Println(isFlipedString1("aa", "aba")) // false
    fmt.Println(isFlipedString1("bluefrog", "leetcode")) // false
}