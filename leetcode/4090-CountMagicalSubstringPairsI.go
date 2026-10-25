package main

// 4090. Count Magical Substring Pairs I
// You are given a string s of length n consisting of lowercase English letters.

// For every triplet of indices (i, j, k) satisfying 0 <= i <= j < k < n, define:
//     1. s1 = s[i..j]
//     2. s2 = s[j + 1..k]

// A pair of substrings (s1, s2) is called magical if the number of distinct characters in s1 is equal to the number of distinct characters in s2.

// Return the number of triplets (i, j, k) for which (s1, s2) is a magical pair.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "aba"
// Output: 2
// Explanation:
// There are 2 valid triplets:
//  No. (i, j, k) s1 s2 distinct(s1) distinct(s2)
//  1 (0, 0, 1) "a" "b" 1 1
//  2 (1, 1, 2) "b" "a" 1 1

// Example 2:
// Input: s = "kkk"
// Output: 4
// Explanation:
// There are 4 valid triplets:
//  No. (i, j, k) s1 s2 distinct(s1) distinct(s2)
//  1 (0, 0, 1) "k" "k" 1 1
//  2 (0, 0, 2) "k" "kk" 1 1
//  3 (0, 1, 2) "kk" "k" 1 1
//  4 (1, 1, 2) "k" "k" 1 1

// Example 3:
// Input: s = "abcac"
// Output: 7
// Explanation:
// There are 7 valid triplets:
//  No. (i, j, k) s1 s2 distinct(s1) distinct(s2)
//  1 (0, 0, 1) "a" "b" 1 1
//  2 (1, 1, 2) "b" "c" 1 1
//  3 (2, 2, 3) "c" "a" 1 1
//  4 (3, 3, 4) "a" "c" 1 1
//  5 (0, 1, 3) "ab" "ca" 2 2
//  6 (0, 1, 4) "ab" "cac" 2 2
//  7 (1, 2, 4) "bc" "ac" 2 2

// Constraints:
//     - 2 <= s.length <= 2000
//     - s consists of lowercase English letters.

import "fmt"
import "slices"

func countMagicalPairs(s string) int {
    logTrick := func(s []byte) [][26]int {
        // 子串包含的字母集合，最小左端点
        type Pair struct{ mask, left int }
        data := []Pair{}
        res := make([][26]int, len(s))
        for i, v := range s {
            bit := 1 << (v - 'a')
            // 计算右端点为 i 的子串的 mask
            for j := range data {
                data[j].mask |= bit
            }
            // v 单独作为子串
            data = append(data, Pair{bit, i})
            // 原地去重（相同 mask 只保留最左边的）
            index := 1
            for j := 1; j < len(data); j++ {
                if data[j].mask != data[j-1].mask {
                    data[index] = data[j]
                    index++
                }
            }
            data = data[:index]
            for k, p := range data {
                right := i
                if k < len(data)-1 {
                    right = data[k+1].left - 1
                }
                // 对于左端点在 [p.left, right]，右端点为 i 的子串，mask 都相同
                res[i][len(data)-1-k] = right - p.left + 1
            }
        }
        return res
    }
    res, arr := 0, []byte(s)
    pre := logTrick(arr)
    slices.Reverse(arr)
    suf := logTrick(arr) // 相当于倒着遍历 s
    slices.Reverse(suf)
    for i := range len(pre) - 1 {
        for j := 0; j < 26 && pre[i][j] > 0 && suf[i+1][j] > 0; j++ {
            res += pre[i][j] * suf[i+1][j]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aba"
    // Output: 2
    // Explanation:
    // There are 2 valid triplets:
    //  No. (i, j, k) s1 s2 distinct(s1) distinct(s2)
    //  1 (0, 0, 1) "a" "b" 1 1
    //  2 (1, 1, 2) "b" "a" 1 1
    fmt.Println(countMagicalPairs("aba")) // 2
    // Example 2:
    // Input: s = "kkk"
    // Output: 4
    // Explanation:
    // There are 4 valid triplets:
    //  No. (i, j, k) s1 s2 distinct(s1) distinct(s2)
    //  1 (0, 0, 1) "k" "k" 1 1
    //  2 (0, 0, 2) "k" "kk" 1 1
    //  3 (0, 1, 2) "kk" "k" 1 1
    //  4 (1, 1, 2) "k" "k" 1 1
    fmt.Println(countMagicalPairs("kkk")) // 4
    // Example 3:
    // Input: s = "abcac"
    // Output: 7
    // Explanation:
    // There are 7 valid triplets:
    //  No. (i, j, k) s1 s2 distinct(s1) distinct(s2)
    //  1 (0, 0, 1) "a" "b" 1 1
    //  2 (1, 1, 2) "b" "c" 1 1
    //  3 (2, 2, 3) "c" "a" 1 1
    //  4 (3, 3, 4) "a" "c" 1 1
    //  5 (0, 1, 3) "ab" "ca" 2 2
    //  6 (0, 1, 4) "ab" "cac" 2 2
    //  7 (1, 2, 4) "bc" "ac" 2 2
    fmt.Println(countMagicalPairs("abcac")) // 7

    fmt.Println(countMagicalPairs("leetcode")) // 18
    fmt.Println(countMagicalPairs("bluefrog")) // 16
    fmt.Println(countMagicalPairs("freewu")) // 11
}