package main

// 3146. Permutation Difference between Two Strings
// You are given two strings s and t such that every character occurs at most once in s and t is a permutation of s.

// The permutation difference between s and t is defined as the sum of the absolute difference 
// between the index of the occurrence of each character in s and the index of the occurrence of the same character in t.

// Return the permutation difference between s and t.

// Example 1:
// Input: s = "abc", t = "bac"
// Output: 2
// Explanation:
// For s = "abc" and t = "bac", the permutation difference of s and t is equal to the sum of:
//     The absolute difference between the index of the occurrence of "a" in s and the index of the occurrence of "a" in t.
//     The absolute difference between the index of the occurrence of "b" in s and the index of the occurrence of "b" in t.
//     The absolute difference between the index of the occurrence of "c" in s and the index of the occurrence of "c" in t.
// That is, the permutation difference between s and t is equal to |0 - 1| + |2 - 2| + |1 - 0| = 2.

// Example 2:
// Input: s = "abcde", t = "edbac"
// Output: 12
// Explanation: The permutation difference between s and t is equal to |0 - 3| + |1 - 2| + |2 - 4| + |3 - 1| + |4 - 0| = 12.

// Constraints:
//     1 <= s.length <= 26
//     Each character occurs at most once in s.
//     t is a permutation of s.
//     s consists only of lowercase English letters.

import "fmt"

func findPermutationDifference(s string, t string) int {
    res, mp := 0, make([]int, 26) // 记录出现位置
    for i, c := range s {
        mp[c - 'a'] = i
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for i, c := range t {
        res += abs(i - mp[c - 'a'])
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abc", t = "bac"
    // Output: 2
    // Explanation:
    // For s = "abc" and t = "bac", the permutation difference of s and t is equal to the sum of:
    //     The absolute difference between the index of the occurrence of "a" in s and the index of the occurrence of "a" in t.
    //     The absolute difference between the index of the occurrence of "b" in s and the index of the occurrence of "b" in t.
    //     The absolute difference between the index of the occurrence of "c" in s and the index of the occurrence of "c" in t.
    // That is, the permutation difference between s and t is equal to |0 - 1| + |2 - 2| + |1 - 0| = 2.
    fmt.Println(findPermutationDifference("abc", "bac")) // 2
    // Example 2:
    // Input: s = "abcde", t = "edbac"
    // Output: 12
    // Explanation: The permutation difference between s and t is equal to |0 - 3| + |1 - 2| + |2 - 4| + |3 - 1| + |4 - 0| = 12.
    fmt.Println(findPermutationDifference("abcde", "edbac")) // 12
}