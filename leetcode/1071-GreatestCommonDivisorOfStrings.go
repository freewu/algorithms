package main

// 1071. Greatest Common Divisor of Strings
// For two strings s and t, we say "t divides s" if 
// and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

// Example 1:
// Input: str1 = "ABCABC", str2 = "ABC"
// Output: "ABC"

// Example 2:
// Input: str1 = "ABABAB", str2 = "ABAB"
// Output: "AB"

// Example 3:
// Input: str1 = "LEET", str2 = "CODE"
// Output: ""

// Constraints:
//     1 <= str1.length, str2.length <= 1000
//     str1 and str2 consist of English uppercase letters.

import "fmt"

// 递归
func gcdOfStrings(str1 string, str2 string) string {
    if str1 == str2 {
        return str1
    }
    if len(str2) > len(str1) {
        str1, str2 = str2, str1
    }
    if str1[:len(str2)] != str2 {
        return ""
    }
    return gcdOfStrings(str1[len(str2):], str2)
}

func main() {
    // Example 1:
    // Input: str1 = "ABCABC", str2 = "ABC"
    // Output: "ABC"
    fmt.Println(gcdOfStrings("ABCABC","ABC")) // "ABC"
    // Example 2:
    // Input: str1 = "ABABAB", str2 = "ABAB"
    // Output: "AB"
    fmt.Println(gcdOfStrings("ABABAB","ABAB")) // "AB"
    // Example 3:
    // Input: str1 = "LEET", str2 = "CODE"
    // Output: ""
    fmt.Println(gcdOfStrings("LEET","CODE")) // ""
}