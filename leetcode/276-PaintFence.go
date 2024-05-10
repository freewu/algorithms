package main

// 276. Paint Fence
// You are painting a fence of n posts with k different colors. You must paint the posts following these rules:
//     Every post must be painted exactly one color.
//     There cannot be three or more consecutive posts with the same color.

// Given the two integers n and k, return the number of ways you can paint the fence.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/28/paintfenceex1.png" />
// Input: n = 3, k = 2
// Output: 6
// Explanation: All the possibilities are shown.
// Note that painting all the posts red or all the posts green is invalid because there cannot be three posts in a row with the same color.

// Example 2:
// Input: n = 1, k = 1
// Output: 1

// Example 3:
// Input: n = 7, k = 2
// Output: 42
 
// Constraints:
//     1 <= n <= 50
//     1 <= k <= 10^5
//     The testcases are generated such that the answer is in the range [0, 2^31 - 1] for the given n and k.

import "fmt"

// 假设f[i]表示前i个栅栏的方案数
// 最多连续两个栅栏同色，那么为每一个栅栏上色的时候，有两个选择：1与前一个同色，表示为f[i][0]；2与前一个异色，表示为f[i][1]；
// 如果选择1，那么方案数不变，且前一个栅栏一定是选择2：f[i][0] = f[i-1][1]
// 如果选择2，那么当前的栅栏可以选择k-1种颜色，且前一个栅栏可以选择1和2：f[i][1] = (k-1)*(f[i-1][0]+f[i-1][1])
// 最终的结果是第n个栅栏的两种选择的方案数之合
// 可知dp状态的转移只与前一个状态有关，所以可以空间压缩为两个变量，分别表示两个选择的方案数
func numWays(n int, k int) int {
    f0, f1 := 0, k
    for i := 1; i < n; i++ {
        f0, f1 = f1, (k-1) * (f0 + f1)
    }
    return f0 + f1
}

func numWays1(n int, k int) int {
    if n <= 0 || k <= 0 {
        return 0
    }
    if n == 1 {
        return k
    }
    if n == 2 {
        return k * k
    }

    dp := make([]int, n)
    dp[0] = k
    dp[1] = k * k
    for i := 2; i < n; i++ {
        // 与前面不同的颜色, 有 k-1 涂法: dp[i] = dp[i-2]*(k-1)
        // 与前面相同的颜色, 有 k-1 涂法: dp[i] = dp[i-1]*(k-1)
        dp[i] = dp[i-1]*(k-1) + dp[i-2]*(k-1)
    }
    return dp[n-1]
}

func main() {
    // Example 1:
    // Input: n = 3, k = 2
    // Output: 6
    // Explanation: All the possibilities are shown.
    // Note that painting all the posts red or all the posts green is invalid because there cannot be three posts in a row with the same color.
    fmt.Println(numWays(3, 2)) // 6
    // Example 2:
    // Input: n = 1, k = 1
    // Output: 1
    fmt.Println(numWays(1, 1)) // 1
    // Example 3:
    // Input: n = 7, k = 2
    // Output: 42
    fmt.Println(numWays(7, 2)) // 42

    fmt.Println(numWays1(3, 2)) // 6
    fmt.Println(numWays1(1, 1)) // 1
    fmt.Println(numWays1(7, 2)) // 42
}