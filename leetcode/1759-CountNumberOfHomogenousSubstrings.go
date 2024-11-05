package main

// 1759. Count Number of Homogenous Substrings
// Given a string s, return the number of homogenous substrings of s. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// A string is homogenous if all the characters of the string are the same.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "abbcccaa"
// Output: 13
// Explanation: The homogenous substrings are listed as below:
// "a"   appears 3 times.
// "aa"  appears 1 time.
// "b"   appears 2 times.
// "bb"  appears 1 time.
// "c"   appears 3 times.
// "cc"  appears 2 times.
// "ccc" appears 1 time.
// 3 + 1 + 2 + 1 + 3 + 2 + 1 = 13.

// Example 2:
// Input: s = "xy"
// Output: 2
// Explanation: The homogenous substrings are "x" and "y".

// Example 3:
// Input: s = "zzzzz"
// Output: 15

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase letters.

import "fmt"

func countHomogenous(s string) int {
    res, count := 0, 0
    for i := 0; i < len(s); i++ {
        if i > 0 && s[i] == s[i-1] {
            count++
        } else {
            count = 1
        }
        res = (res + count) % 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abbcccaa"
    // Output: 13
    // Explanation: The homogenous substrings are listed as below:
    // "a"   appears 3 times.
    // "aa"  appears 1 time.
    // "b"   appears 2 times.
    // "bb"  appears 1 time.
    // "c"   appears 3 times.
    // "cc"  appears 2 times.
    // "ccc" appears 1 time.
    // 3 + 1 + 2 + 1 + 3 + 2 + 1 = 13.
    fmt.Println(countHomogenous("abbcccaa")) // 13
    // Example 2:
    // Input: s = "xy"
    // Output: 2
    // Explanation: The homogenous substrings are "x" and "y".
    fmt.Println(countHomogenous("xy")) // 2
    // Input: s = "zzzzz"
    // Output: 15
    fmt.Println(countHomogenous("zzzzz")) // 15
}