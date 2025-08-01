package main

// 3545. Minimum Deletions for At Most K Distinct Characters
// You are given a string s consisting of lowercase English letters, and an integer k.

// Your task is to delete some (possibly none) of the characters in the string so that the number of distinct characters in the resulting string is at most k.

// Return the minimum number of deletions required to achieve this.

// Example 1:
// Input: s = "abc", k = 2
// Output: 1
// Explanation:
// s has three distinct characters: 'a', 'b' and 'c', each with a frequency of 1.
// Since we can have at most k = 2 distinct characters, remove all occurrences of any one character from the string.
// For example, removing all occurrences of 'c' results in at most k distinct characters. Thus, the answer is 1.

// Example 2:
// Input: s = "aabb", k = 2
// Output: 0
// Explanation:
// s has two distinct characters ('a' and 'b') with frequencies of 2 and 2, respectively.
// Since we can have at most k = 2 distinct characters, no deletions are required. Thus, the answer is 0.

// Example 3:
// Input: s = "yyyzz", k = 1
// Output: 2
// Explanation:
// s has two distinct characters ('y' and 'z') with frequencies of 3 and 2, respectively.
// Since we can have at most k = 1 distinct character, remove all occurrences of any one character from the string.
// Removing all 'z' results in at most k distinct characters. Thus, the answer is 2.

// Constraints:
//     1 <= s.length <= 16
//     1 <= k <= 16
//     s consists only of lowercase English letters.

import "fmt"
import "sort"

func minDeletion(s string, k int) int {
    res, count := 0, make([]int, 26)
    for _, v := range s {
        count[v - 'a']++
    }
    sort.Ints(count)
    for i := 0; i < 26 - k;i++ {
        res += count[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abc", k = 2
    // Output: 1
    // Explanation:
    // s has three distinct characters: 'a', 'b' and 'c', each with a frequency of 1.
    // Since we can have at most k = 2 distinct characters, remove all occurrences of any one character from the string.
    // For example, removing all occurrences of 'c' results in at most k distinct characters. Thus, the answer is 1.
    fmt.Println(minDeletion("abc", 2)) // 1
    // Example 2:
    // Input: s = "aabb", k = 2
    // Output: 0
    // Explanation:
    // s has two distinct characters ('a' and 'b') with frequencies of 2 and 2, respectively.
    // Since we can have at most k = 2 distinct characters, no deletions are required. Thus, the answer is 0.
    fmt.Println(minDeletion("aabb", 2)) // 0
    // Example 3:
    // Input: s = "yyyzz", k = 1
    // Output: 2
    // Explanation:
    // s has two distinct characters ('y' and 'z') with frequencies of 3 and 2, respectively.
    // Since we can have at most k = 1 distinct character, remove all occurrences of any one character from the string.
    // Removing all 'z' results in at most k distinct characters. Thus, the answer is 2.
    fmt.Println(minDeletion("yyyzz", 1)) // 2

    fmt.Println(minDeletion("bluefrog", 2)) // 6
    fmt.Println(minDeletion("leetcode", 2)) // 4
}