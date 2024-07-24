package main

// 343. Integer Break
// Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.
// Return the maximum product you can get.

// Example 1:
// Input: n = 2
// Output: 1
// Explanation: 2 = 1 + 1, 1 × 1 = 1.

// Example 2:
// Input: n = 10
// Output: 36
// Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.

// Constraints:
//     2 <= n <= 58

import "fmt"

// dp
func integerBreak(n int) int {
    dp := make([]int, n + 1)
    dp[0], dp[1] = 1, 1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 1; j < i; j++ {
            dp[i] = max(dp[i], j* max(dp[i-j], i-j))
        }
    }
    return dp[n]
}

func integerBreak1(n int) int {
    if n <= 3 {
        return n - 1
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := 1
    for n > 0 {
        if n > 4 {
            res *= min(n, 3)
            n -= 3
        } else {
            res *= n
            n = 0
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 1
    // Explanation: 2 = 1 + 1, 1 × 1 = 1.
    fmt.Println(integerBreak(2)) // 1
    // Example 2:
    // Input: n = 10
    // Output: 36
    // Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
    fmt.Println(integerBreak(10)) // 36
    fmt.Println(integerBreak(12)) // 81

    fmt.Println(integerBreak1(2)) // 1
    fmt.Println(integerBreak1(10)) // 36
    fmt.Println(integerBreak1(12)) // 81
}