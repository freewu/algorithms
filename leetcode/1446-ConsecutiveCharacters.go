package main

// 1446. Consecutive Characters
// The power of the string is the maximum length of a non-empty substring that contains only one unique character.
// Given a string s, return the power of s.

// Example 1:
// Input: s = "leetcode"
// Output: 2
// Explanation: The substring "ee" is of length 2 with the character 'e' only.

// Example 2:
// Input: s = "abbcccddddeeeeedcba"
// Output: 5
// Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
 
// Constraints:
//     1 <= s.length <= 500
//     s consists of only lowercase English letters.

import "fmt"

func maxPower(s string) int {
    res, count := 0, 1
    for i := 1; i < len(s); i++ { 
        if s[i - 1] == s[i] { // 判断是否与上一个字符是否相同
            count++
        } else {
            if count > res {
                res = count
            }
            count = 1
        }
    }
    if count > res {
        res = count
    }
    return res
}

func maxPower1(s string) int {
    res, count := 1, 1
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            count++
            if count > res {
                res = count
            }
        } else {
            count = 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "leetcode"
    // Output: 2
    // Explanation: The substring "ee" is of length 2 with the character 'e' only.
    fmt.Println(maxPower("leetcode")) // 2
    // Example 2:
    // Input: s = "abbcccddddeeeeedcba"
    // Output: 5
    // Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
    fmt.Println(maxPower("abbcccddddeeeeedcba")) // 5

    fmt.Println(maxPower("cc")) // 2

    fmt.Println(maxPower1("leetcode")) // 2
    fmt.Println(maxPower1("abbcccddddeeeeedcba")) // 5
    fmt.Println(maxPower1("cc")) // 2
}