package main

// 1915. Number of Wonderful Substrings
// A wonderful string is a string where at most one letter appears an odd number of times.
//     For example, "ccjjc" and "abab" are wonderful, but "ab" is not.

// Given a string word that consists of the first ten lowercase English letters ('a' through 'j'), 
// return the number of wonderful non-empty substrings in word. 
// If the same substring appears multiple times in word, then count each occurrence separately.

// A substring is a contiguous sequence of characters in a string.

// Example 1:
// Input: word = "aba"
// Output: 4
// Explanation: The four wonderful substrings are underlined below:
// - "(a)ba" -> "a"
// - "a(b)a" -> "b"
// - "ab(a)" -> "a"
// - "(aba)" -> "aba"

// Example 2:
// Input: word = "aabb"
// Output: 9
// Explanation: The nine wonderful substrings are underlined below:
// - "(a)abb" -> "a"
// - "(aa)bb" -> "aa"
// - "(aab)b" -> "aab"
// - "(aabb)" -> "aabb"
// - "a(a)bb" -> "a"
// - "a(abb)" -> "abb"
// - "aa(b)b" -> "b"
// - "aa(bb)" -> "bb"
// - "aab(b)" -> "b"

// Example 3:
// Input: word = "he"
// Output: 2
// Explanation: The two wonderful substrings are underlined below:
// - "(h)e" -> "h"
// - "h(e)" -> "e"
 
// Constraints:
//     1 <= word.length <= 10^5
//     word consists of lowercase English letters from 'a' to 'j'.

import "fmt"

func wonderfulSubstrings(word string) int64 {
    res, count := int64(0), make([]int64, 1024)  // 计数数组，用于记录每个前缀的奇数字符出现次数的状态
    count[0] = 1  // 空字符串的状态为0
    mask := 0  // 当前前缀的状态
    for _, char := range word {
        bit := int(char - 'a')  // 计算当前字符对应的二进制位
        mask ^= 1 << bit  // 更新当前前缀的状态
        res += count[mask]  // 计算与当前前缀状态相差一个字符的状态
        for i := 0; i < 10; i++ {
            res += count[mask ^ (1 << i)]  // 计算与当前前缀状态相差一个字符的状态（每个字符出现偶数次）
        }
        count[mask]++  // 更新计数数组
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "aba"
    // Output: 4
    // Explanation: The four wonderful substrings are underlined below:
    // - "(a)ba" -> "a"
    // - "a(b)a" -> "b"
    // - "ab(a)" -> "a"
    // - "(aba)" -> "aba"
    fmt.Println(wonderfulSubstrings("aba")) // 4
    // Example 2:
    // Input: word = "aabb"
    // Output: 9
    // Explanation: The nine wonderful substrings are underlined below:
    // - "(a)abb" -> "a"
    // - "(aa)bb" -> "aa"
    // - "(aab)b" -> "aab"
    // - "(aabb)" -> "aabb"
    // - "a(a)bb" -> "a"
    // - "a(abb)" -> "abb"
    // - "aa(b)b" -> "b"
    // - "aa(bb)" -> "bb"
    // - "aab(b)" -> "b"
    fmt.Println(wonderfulSubstrings("aabb")) // 9
    // Example 3:
    // Input: word = "he"
    // Output: 2
    // Explanation: The two wonderful substrings are underlined below:
    // - "(h)e" -> "h"
    // - "h(e)" -> "e"
    fmt.Println(wonderfulSubstrings("he")) // 2
}