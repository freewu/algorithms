package main

// 3821. Find Nth Smallest Integer With K One Bits
// You are given two positive integers n and k.

// Return an integer denoting the nth smallest positive integer that has exactly k ones in its binary representation. 
// It is guaranteed that the answer is strictly less than 2^50.

// Example 1:
// Input: n = 4, k = 2
// Output: 9
// Explanation:
// The 4 smallest positive integers that have exactly k = 2 ones in their binary representations are:
// 3 = 112
// 5 = 1012
// 6 = 1102
// 9 = 10012

// Example 2:
// Input: n = 3, k = 1
// Output: 4
// Explanation:
// The 3 smallest positive integers that have exactly k = 1 one in their binary representations are:
// 1 = 12
// 2 = 102
// 4 = 1002

// Constraints:
//     1 <= n <= 2^50
//     1 <= k <= 50
//     The answer is strictly less than 2^50.

import "fmt"

const MX = 50
var comb [MX][MX + 1]int64

func init() {
    for i := range comb { // 预处理组合数
        comb[i][0] = 1
        for j := 1; j <= i; j++ {
            comb[i][j] = comb[i-1][j-1] + comb[i-1][j]
        }
    }
}

func nthSmallest(n int64, k int) int64 {
    res := int64(0)
    for i := MX - 1; k > 0; i-- {
        c := comb[i][k] // 第 i 位填 0 的方案数
        if n > c { // n 比较大，第 i 位必须填 1
            n -= c
            res |= 1 << i
            k-- // 维护剩余的 1 的个数
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, k = 2
    // Output: 9
    // Explanation:
    // The 4 smallest positive integers that have exactly k = 2 ones in their binary representations are:
    // 3 = 112
    // 5 = 1012
    // 6 = 1102
    // 9 = 10012
    fmt.Println(nthSmallest(4, 2)) // 9
    // Example 2:
    // Input: n = 3, k = 1
    // Output: 4
    // Explanation:
    // The 3 smallest positive integers that have exactly k = 1 one in their binary representations are:
    // 1 = 12
    // 2 = 102
    // 4 = 1002
    fmt.Println(nthSmallest(3, 1)) // 4

    fmt.Println(nthSmallest(1, 1)) // 1
    fmt.Println(nthSmallest(1, 50)) // 1125899906842623
    fmt.Println(nthSmallest(1024, 1)) // 562949953421312
    fmt.Println(nthSmallest(1024, 50)) // 1125899906842623
    fmt.Println(nthSmallest(2 << 50, 1)) // 562949953421312
    fmt.Println(nthSmallest(2 << 50, 50)) // 1125899906842623
}