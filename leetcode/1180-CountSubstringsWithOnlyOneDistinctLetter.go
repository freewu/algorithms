package main

// 1180. Count Substrings with Only One Distinct Letter
// Given a string s, return the number of substrings that have only one distinct letter.

// Example 1:
// Input: s = "aaaba"
// Output: 8
// Explanation: The substrings with one distinct letter are "aaa", "aa", "a", "b".
// "aaa" occurs 1 time.
// "aa" occurs 2 times.
// "a" occurs 4 times.
// "b" occurs 1 time.
// So the answer is 1 + 2 + 4 + 1 = 8.

// Example 2:
// Input: s = "aaaaaaaaaa"
// Output: 55
 
// Constraints:
//     1 <= s.length <= 1000
//     s[i] consists of only lowercase English letters.

import "fmt"

func countLetters(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        for k := i; k < len(s); k++ {
            if s[i] == s[k] {
                res = res + 1
            } else {
                break
            }
        }
    }
    return res
}

func countLetters1(s string) int {
    res := 0
    for i, size := 0, len(s); i < size; i++ {
        l, r:= i, i
        for r < size {
            if s[l] == s[r] {
                res++
            } else {
                break
            }
            r++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aaaba"
    // Output: 8
    // Explanation: The substrings with one distinct letter are "aaa", "aa", "a", "b".
    // "aaa" occurs 1 time.
    // "aa" occurs 2 times.
    // "a" occurs 4 times.
    // "b" occurs 1 time.
    // So the answer is 1 + 2 + 4 + 1 = 8.
    fmt.Println(countLetters("aaaba")) // 8
    // Example 2:
    // Input: s = "aaaaaaaaaa"
    // Output: 55
    fmt.Println(countLetters("aaaaaaaaaa")) // 55

    fmt.Println(countLetters1("aaaba")) // 8
    fmt.Println(countLetters1("aaaaaaaaaa")) // 55
}