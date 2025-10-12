package main

// 3707. Equal Score Substrings
// You are given a string s consisting of lowercase English letters.

// The score of a string is the sum of the positions of its characters in the alphabet, where 'a' = 1, 'b' = 2, ..., 'z' = 26.

// Determine whether there exists an index i such that the string can be split into two non-empty substrings s[0..i] and s[(i + 1)..(n - 1)] that have equal scores.

// Return true if such a split exists, otherwise return false.

// Example 1:
// Input: s = "adcb"
// Output: true
// Explanation:
// Split at index i = 1:
// Left substring = s[0..1] = "ad" with score = 1 + 4 = 5
// Right substring = s[2..3] = "cb" with score = 3 + 2 = 5
// Both substrings have equal scores, so the output is true.

// Example 2:
// Input: s = "bace"
// Output: false
// Explanation:​​​​​​
// ​​​​​​​No split produces equal scores, so the output is false.

// Constraints:
//     2 <= s.length <= 100
//     s consists of lowercase English letters.

import "fmt"

func scoreBalance(s string) bool {
    sum, left := 0, 0
    for _, v := range s { // 计算字符串的总分数
        sum += int(v & 31)
    }
    for _, v := range s { // 字母位置是正数，可以遍历到 s 末尾（末尾一定不满足要求）
        left += int(v & 31)
        if left * 2 == sum {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: s = "adcb"
    // Output: true
    // Explanation:
    // Split at index i = 1:
    // Left substring = s[0..1] = "ad" with score = 1 + 4 = 5
    // Right substring = s[2..3] = "cb" with score = 3 + 2 = 5
    // Both substrings have equal scores, so the output is true.
    fmt.Println(scoreBalance("adcb")) // true
    // Example 2:
    // Input: s = "bace"
    // Output: false
    // Explanation:​​​​​​
    // ​​​​​​​No split produces equal scores, so the output is false.
    fmt.Println(scoreBalance("bace")) // false

    fmt.Println(scoreBalance("bluefrog")) // false
    fmt.Println(scoreBalance("leetcode")) // false
    fmt.Println(scoreBalance("abcdefghijklmnopqrstuvwxyz")) // false
    fmt.Println(scoreBalance("zyxwvutsrqponmlkjihgfedcba")) // false
}