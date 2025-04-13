package main

// 3518. Smallest Palindromic Rearrangement II
// You are given a palindromic string s and an integer k.

// Return the k-th lexicographically smallest palindromic permutation of s. 
// If there are fewer than k distinct palindromic permutations, return an empty string.

// Note: Different rearrangements that yield the same palindromic string are considered identical and are counted once.

// Example 1:
// Input: s = "abba", k = 2
// Output: "baab"
// Explanation:
// The two distinct palindromic rearrangements of "abba" are "abba" and "baab".
// Lexicographically, "abba" comes before "baab". Since k = 2, the output is "baab".

// Example 2:
// Input: s = "aa", k = 2
// Output: ""
// Explanation:
// There is only one palindromic rearrangement: "aa".
// The output is an empty string since k = 2 exceeds the number of possible rearrangements.

// Example 3:
// Input: s = "bacab", k = 1
// Output: "abcba"
// Explanation:
// The two distinct palindromic rearrangements of "bacab" are "abcba" and "bacab".
// Lexicographically, "abcba" comes before "bacab". Since k = 1, the output is "abcba".

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of lowercase English letters.
//     s is guaranteed to be palindromic.
//     1 <= k <= 10^6

import "fmt"
import "slices"

func smallestPalindrome(s string, k int) string {
    n := len(s)
    m := n / 2
    freq := make([]int, 26)
    for _, v := range s[:m] {
        freq[v - 'a']++
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    comb := func(n, m int) int {
        m = min(m, n-m)
        res := 1
        for i := 1; i <= m; i++ {
            res = res * (n + 1 - i) / i
            if res >= k { // 太大了
                return k
            }
        }
        return res
    }
    // 计算长为 n 的字符串的排列个数
    perm := func(n int) int {
        res := 1
        for _, c := range freq {
            if c == 0 { continue }
            res *= comb(n, c) // 先从 n 个里面选 c 个位置填当前字母
            if res >= k { return k } // 太大了
            n -= c // 从剩余位置中选位置填下一个字母
        }
        return res
    }
    if perm(m) < k { return "" } // k 太大
    // 构造回文串的左半部分
    left := make([]byte, m)
    for i := range left {
        for j := range freq {
            if freq[j] == 0 { continue }
            freq[j]-- // 假设填字母 j，看是否有足够的排列
            p := perm(m - i - 1) // 剩余位置的排列个数
            if p >= k { // 有足够的排列
                left[i] = 'a' + byte(j)
                break
            }
            k -= p // k 太大，要填更大的字母
            freq[j]++
        }
    }
    res := string(left)
    if n % 2 > 0 { // 奇数个
        res += string(s[n/2])
    }
    slices.Reverse(left)
    return res + string(left)
}

func main() {
    // Example 1:
    // Input: s = "abba", k = 2
    // Output: "baab"
    // Explanation:
    // The two distinct palindromic rearrangements of "abba" are "abba" and "baab".
    // Lexicographically, "abba" comes before "baab". Since k = 2, the output is "baab".
    fmt.Println(smallestPalindrome("abba", 2)) // "baab"
    // Example 2:
    // Input: s = "aa", k = 2
    // Output: ""
    // Explanation:
    // There is only one palindromic rearrangement: "aa".
    // The output is an empty string since k = 2 exceeds the number of possible rearrangements.
    fmt.Println(smallestPalindrome("aa", 2)) // ""
    // Example 3:
    // Input: s = "bacab", k = 1
    // Output: "abcba"
    // Explanation:
    // The two distinct palindromic rearrangements of "bacab" are "abcba" and "bacab".
    // Lexicographically, "abcba" comes before "bacab". Since k = 1, the output is "abcba".
    fmt.Println(smallestPalindrome("bacab", 1)) // "abcba"

    fmt.Println(smallestPalindrome("bluefrogbluefrog", 2)) // "abcba"
    fmt.Println(smallestPalindrome("leetcodeleetcode", 2)) // "abcba"
}