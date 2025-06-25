package main

// 2311. Longest Binary Subsequence Less Than or Equal to K
// You are given a binary string s and a positive integer k.

// Return the length of the longest subsequence of s that makes up a binary number less than or equal to k.

// Note:
//     1. The subsequence can contain leading zeroes.
//     2. The empty string is considered to be equal to 0.
//     3. A subsequence is a string that can be derived from another string by deleting some 
//        or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "1001010", k = 5
// Output: 5
// Explanation: The longest subsequence of s that makes up a binary number less than or equal to 5 is "00010", as this number is equal to 2 in decimal.
// Note that "00100" and "00101" are also possible, which are equal to 4 and 5 in decimal, respectively.
// The length of this subsequence is 5, so 5 is returned.

// Example 2:
// Input: s = "00101001", k = 1
// Output: 6
// Explanation: "000001" is the longest subsequence of s that makes up a binary number less than or equal to 1, as this number is equal to 1 in decimal.
// The length of this subsequence is 6, so 6 is returned.

// Constraints:
//     1 <= s.length <= 1000
//     s[i] is either '0' or '1'.
//     1 <= k <= 10^9

import "fmt"
import "math"
import "strings"
import "math/bits"
import "strconv"

// 解答错误 151 / 153 
// func longestSubsequence(s string, k int) int {
//     res, sum, flag := 0, 0, false
//     pow := func (b int, p int) int { // 求幂
//         res := 1
//         for i := 0; i < p; i++ {
//             res *= b
//         }
//         return res
//     }
//     for i := len(s) - 1; i >= 0; i--{
//         if rune(s[i]) == '0' {
//             res++
//         } else if !flag {
//             sum += pow(2, res)
//             if sum > k {
//                 flag = true
//             } else {
//                 res++
//             }
//         }
//     }
//     return res
// }

func longestSubsequence(s string, k int) int {
    res, sum, flag := 0, float64(0), false
    for i := len(s) - 1; i >= 0; i--{
        if s[i] == '0'{
            res++
        } else if !flag {
            sum += math.Pow(float64(2), float64(res))
            if sum > float64(k) {
                flag = true
            } else {
                res++
            }
        }
    }
    return res
}

func longestSubsequence1(s string, k int) int {
    n, m := len(s), bits.Len(uint(k))
    if n < m { return n }
    res := m
    if v, _ := strconv.ParseInt(s[n-m:], 2, 0); int(v) > k {
        res--
    }
    return res + strings.Count(s[:n-m], "0")
}

func main() {
    // Example 1:
    // Input: s = "1001010", k = 5
    // Output: 5
    // Explanation: The longest subsequence of s that makes up a binary number less than or equal to 5 is "00010", as this number is equal to 2 in decimal.
    // Note that "00100" and "00101" are also possible, which are equal to 4 and 5 in decimal, respectively.
    // The length of this subsequence is 5, so 5 is returned.
    fmt.Println(longestSubsequence("1001010", 5)) // 5
    // Example 2:
    // Input: s = "00101001", k = 1
    // Output: 6
    // Explanation: "000001" is the longest subsequence of s that makes up a binary number less than or equal to 1, as this number is equal to 1 in decimal.
    // The length of this subsequence is 6, so 6 is returned.
    fmt.Println(longestSubsequence("00101001", 1)) // 6

    fmt.Println(longestSubsequence("1111111111", 1)) // 1
    fmt.Println(longestSubsequence("1111100000", 1)) // 5
    fmt.Println(longestSubsequence("0000011111", 1)) // 6
    fmt.Println(longestSubsequence("0000000000", 1)) // 10
    fmt.Println(longestSubsequence("0101010101", 1)) // 6
    fmt.Println(longestSubsequence("1010101010", 1)) // 5

    fmt.Println(longestSubsequence1("1001010", 5)) // 5
    fmt.Println(longestSubsequence1("00101001", 1)) // 6
    fmt.Println(longestSubsequence1("1111111111", 1)) // 1
    fmt.Println(longestSubsequence1("1111100000", 1)) // 5
    fmt.Println(longestSubsequence1("0000011111", 1)) // 6
    fmt.Println(longestSubsequence1("0000000000", 1)) // 10
    fmt.Println(longestSubsequence1("0101010101", 1)) // 6
    fmt.Println(longestSubsequence1("1010101010", 1)) // 5
}