package main

// 4043. Count Rotations With Exactly K Equal Adjacent Pairs
// You are given a string s of length n and an integer k.

// A cyclic rotation of s is obtained by choosing a prefix of s whose length is between 0 and n - 1 (inclusive), 
// and moving it to the end of the string while preserving the order of all characters.

// For every cyclic rotation of s, let its score be the number of indices i such that 0 <= i < n - 1 and the characters at positions i and i + 1 are equal.

// Return the number of cyclic rotations of s whose score equals k.

// Example 1:
// Input: s = "aab", k = 1
// Output: 2
// Explanation:
// The cyclic rotations of s are:
// "aab": The characters at positions 0 and 1 are equal, so score = 1.
// "aba": No two adjacent characters are equal, so score = 0.
// "baa": The characters at positions 1 and 2 are equal, so score = 1.
// Since score equals k for 2 cyclic rotations of s, the answer is 2.

// Example 2:
// Input: s = "abca", k = 0
// Output: 1
// Explanation:
// The cyclic rotations of s are:
// "abca": No two adjacent characters are equal, so score = 0.
// "bcaa": The characters at positions 2 and 3 are equal, so score = 1.
// "caab": The characters at positions 1 and 2 are equal, so score = 1.
// "aabc": The characters at positions 0 and 1 are equal, so score = 1.
// Since score equals k for only 1 cyclic rotation of s, the answer is 1.

// Constraints:
//     2 <= n == s.length <= 100
//     s only consists of lowercase English letters.
//     0 <= k <= n - 1

import "fmt"

func countRotations(s string, k int) int {
    res := 0
    for i := 1; i < len(s); i++ {
        if s[i] ==  s[i-1] {
            res++
        }
    }
    if s[0] == s[len(s)-1] {
        res++
    }
    if k == res {
        return len(s) - res
    }
    if k == res-1 {
        return res
    }
    return 0
}

func main() {
    // Example 1:
    // Input: s = "aab", k = 1
    // Output: 2
    // Explanation:
    // The cyclic rotations of s are:
    // "aab": The characters at positions 0 and 1 are equal, so score = 1.
    // "aba": No two adjacent characters are equal, so score = 0.
    // "baa": The characters at positions 1 and 2 are equal, so score = 1.
    // Since score equals k for 2 cyclic rotations of s, the answer is 2.
    fmt.Println(countRotations("aab", 1)) // 2
    // Example 2:
    // Input: s = "abca", k = 0
    // Output: 1
    // Explanation:
    // The cyclic rotations of s are:
    // "abca": No two adjacent characters are equal, so score = 0.
    // "bcaa": The characters at positions 2 and 3 are equal, so score = 1.
    // "caab": The characters at positions 1 and 2 are equal, so score = 1.
    // "aabc": The characters at positions 0 and 1 are equal, so score = 1.
    // Since score equals k for only 1 cyclic rotation of s, the answer is 1.
    fmt.Println(countRotations("abca", 0)) // 1

    fmt.Println(countRotations("bluefrog", 1)) // 0
    fmt.Println(countRotations("leetcode", 1)) // 7
    fmt.Println(countRotations("freewu", 1)) // 5
}