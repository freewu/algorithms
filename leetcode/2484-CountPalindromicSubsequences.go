package main

// 2484. Count Palindromic Subsequences
// Given a string of digits s, return the number of palindromic subsequences of s having length 5. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Note:
//     1. A string is palindromic if it reads the same forward and backward.
//     2. A subsequence is a string that can be derived from another string by deleting some 
//        or no characters without changing the order of the remaining characters.

// Example 1:
// Input: s = "103301"
// Output: 2
// Explanation: 
// There are 6 possible subsequences of length 5: "10330","10331","10301","10301","13301","03301". 
// Two of them (both equal to "10301") are palindromic.

// Example 2:
// Input: s = "0000000"
// Output: 21
// Explanation: All 21 subsequences are "00000", which is palindromic.

// Example 3:
// Input: s = "9999900000"
// Output: 2
// Explanation: The only two palindromic subsequences are "99999" and "00000".

// Constraints:
//     1 <= s.length <= 10^4
//     s consists of digits.

import "fmt"

func countPalindromes(s string) int {
    res, n, mod := 0, len(s), 1_000_000_007
    prefix, suffix, count, t := [10010][10][10]int{}, [10010][10][10]int{}, make([]int, 10), make([]int, n)
    for i, c := range s {
        t[i] = int(c - '0')
    }
    for i := 1; i <= n; i++ {
        v := t[i-1]
        for j := 0; j < 10; j++ {
            for k := 0; k < 10; k++ {
                prefix[i][j][k] = prefix[i-1][j][k]
            }
        }
        for j := 0; j < 10; j++ {
            prefix[i][j][v] += count[j]
        }
        count[v]++
    }
    count = make([]int, 10)
    for i := n; i > 0; i-- {
        v := t[i-1]
        for j := 0; j < 10; j++ {
            for k := 0; k < 10; k++ {
                suffix[i][j][k] = suffix[i+1][j][k]
            }
        }
        for j := 0; j < 10; j++ {
            suffix[i][j][v] += count[j]
        }
        count[v]++
    }
    for i := 1; i <= n; i++ {
        for j := 0; j < 10; j++ {
            for k := 0; k < 10; k++ {
                res += prefix[i-1][j][k] * suffix[i+1][j][k]
                res %= mod
            }
        }
    }
    return res
}

func countPalindromes1(s string) int {
    res, n := 0, len(s)
    // 题目：要求可以重复，只要序号不重复皆可
    // 那么可以考虑中间字符c，只要在两边找到相同字符组合即可
    // 那么我们使用 suffix[a][b], prefix[a][b]表示后缀 / 前缀为ab的组合的个数
    // 则以当前字符为中心的长度为5的回文序列为 prefix[a][b] * suffix[a][b]
    suffix1, suffix2 := [10]int{}, [10][10]int{}
    for i := n - 1; i >= 0; i-- {
        v := s[i] - '0'
        for j := 0; j < 10; j++ {
            suffix2[v][j] += suffix1[j]
        }
        suffix1[v]++
    }
    // 在遍历的时候一边初始化 prefix1 和 prefix2
    prefix1, prefix2 := [10]int{}, [10][10]int{}
    for _, d := range s[:n - 1] {
        v := d - '0'
        // 先撤销suff1和suff2在统计中心字符的子序列个数
        suffix1[v]--
        for i, c := range suffix1 {
            suffix2[v][i] -= c
        }
        // 遍历中心字符,计算中心序列总和总和
        for i := 0; i < 10; i++ {
            for j := 0; j < 10; j++ {
                res += (prefix2[i][j] * suffix2[j][i]) // 这里需要注意是 prefix2[i][j] 和 suffix2[j][i]
                res %= 1_000_000_007
                // 因为我们定义了a到b之间的有序性，因此也需要相反才能得出正确答案
            }
        }
        // 统计 prefix1, prefix2
        for i := 0; i < 10; i++ {
            prefix2[i][v] += prefix1[i]
        }
        prefix1[v]++
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "103301"
    // Output: 2
    // Explanation: 
    // There are 6 possible subsequences of length 5: "10330","10331","10301","10301","13301","03301". 
    // Two of them (both equal to "10301") are palindromic.
    fmt.Println(countPalindromes("103301")) // 2
    // Example 2:
    // Input: s = "0000000"
    // Output: 21
    // Explanation: All 21 subsequences are "00000", which is palindromic.
    fmt.Println(countPalindromes("0000000")) // 21
    // Example 3:
    // Input: s = "9999900000"
    // Output: 2
    // Explanation: The only two palindromic subsequences are "99999" and "00000".
    fmt.Println(countPalindromes("9999900000")) // 2
    fmt.Println(countPalindromes("112233445566778899")) // 0

    fmt.Println(countPalindromes1("103301")) // 2
    fmt.Println(countPalindromes1("0000000")) // 21
    fmt.Println(countPalindromes1("9999900000")) // 2
    fmt.Println(countPalindromes1("112233445566778899")) // 0
}