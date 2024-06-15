package main

// 521. Longest Uncommon Subsequence I
// Given two strings a and b, return the length of the longest uncommon subsequence between a and b. 
// If no such uncommon subsequence exists, return -1.
// An uncommon subsequence between two strings is a string that is a subsequence of exactly one of them.

// Example 1:
// Input: a = "aba", b = "cdc"
// Output: 3
// Explanation: One longest uncommon subsequence is "aba" because "aba" is a subsequence of "aba" but not "cdc".
// Note that "cdc" is also a longest uncommon subsequence.

// Example 2:
// Input: a = "aaa", b = "bbb"
// Output: 3
// Explanation: The longest uncommon subsequences are "aaa" and "bbb".

// Example 3:
// Input: a = "aaa", b = "aaa"
// Output: -1
// Explanation: Every subsequence of string a is also a subsequence of string b. Similarly, every subsequence of string b is also a subsequence of string a. So the answer would be -1.
 
// Constraints:
//     1 <= a.length, b.length <= 100
//     a and b consist of lower-case English letters.

import "fmt"

func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }
    if len(a) > len(b) {
        return len(a)
    }
    return len(b)
}

func findLUSlength1(a string, b string) int {
    if a == b {
        return -1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(len(a), len(b))
}

func main() {
    // Example 1:
    // Input: a = "aba", b = "cdc"
    // Output: 3
    // Explanation: One longest uncommon subsequence is "aba" because "aba" is a subsequence of "aba" but not "cdc".
    // Note that "cdc" is also a longest uncommon subsequence.
    fmt.Println(findLUSlength("aba", "cdc")) // 3
    // Example 2:
    // Input: a = "aaa", b = "bbb"
    // Output: 3
    // Explanation: The longest uncommon subsequences are "aaa" and "bbb".
    fmt.Println(findLUSlength("aaa", "cdc")) // 3
    // Example 3:
    // Input: a = "aaa", b = "aaa"
    // Output: -1
    // Explanation: Every subsequence of string a is also a subsequence of string b. Similarly, every subsequence of string b is also a subsequence of string a. So the answer would be -1.
    fmt.Println(findLUSlength("aaa", "aaa")) // -1

    fmt.Println(findLUSlength1("aba", "cdc")) // 3
    fmt.Println(findLUSlength1("aaa", "cdc")) // 3
    fmt.Println(findLUSlength1("aaa", "aaa")) // -1
}