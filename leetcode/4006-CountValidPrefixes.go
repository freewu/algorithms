package main

// 4006. Count Valid Prefixes
// You are given a binary string s.

// A prefix of s is considered valid if its characters can be rearranged to form an alternating string.

// Return the number of valid prefixes of s.

// A binary string is a string consisting only of '0' and '1'.

// A prefix of a string is a substring that starts from the beginning of the string and extends to any point within it.

// A substring is a contiguous non-empty sequence of characters within a string.

// A string is considered alternating if no two adjacent characters are equal.

// Example 1:
// Input: s = "00101"
// Output: 3
// Explanation:
// The valid prefixes are:
// "0": It is already an alternating string.
// "001": It can be rearranged into "010", which is an alternating string.
// "00101": It can be rearranged into "01010", which is an alternating string.
// Thus, the answer is 3.

// Example 2:
// Input: s = "101"
// Output: 3
// Explanation:
// All prefixes of s = "101" are already alternating strings. Thus, the answer is 3.

// Constraints:
//     1 <= s.length <= 100
//     s consists only of '0' and '1'.

import "fmt"

func countValidPrefixes(s string) int {
    res, balance := 0, 0
    for _, ch := range s {
        if ch == '1' {
            balance++
        } else {
            balance--
        }
        if balance >= -1 && balance <= 1 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "00101"
    // Output: 3
    // Explanation:
    // The valid prefixes are:
    // "0": It is already an alternating string.
    // "001": It can be rearranged into "010", which is an alternating string.
    // "00101": It can be rearranged into "01010", which is an alternating string.
    // Thus, the answer is 3.
    fmt.Println(countValidPrefixes("00101")) // 3 
    // Example 2:
    // Input: s = "101"
    // Output: 3
    // Explanation:
    // All prefixes of s = "101" are already alternating strings. Thus, the answer is 3.
    fmt.Println(countValidPrefixes("101")) // 3 

    fmt.Println(countValidPrefixes("0000000000")) // 1
    fmt.Println(countValidPrefixes("1111111111")) // 1
    fmt.Println(countValidPrefixes("0000011111")) // 3
    fmt.Println(countValidPrefixes("1111100000")) // 3 
    fmt.Println(countValidPrefixes("0101010101")) // 10
    fmt.Println(countValidPrefixes("1010101010")) // 10
}