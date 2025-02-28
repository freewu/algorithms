package main

// 面试题 08.01. Three Steps Problem LCCI
// A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3 steps at a time. 
// Implement a method to count how many possible ways the child can run up the stairs. 
// The result may be large, so return it modulo 1000000007.

// Example1:
// Input: n = 3 
// Output: 4

// Example2:
// Input: n = 5
// Output: 13

// Note:
//     1 <= n <= 1000000

import "fmt"

// 动态规划
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func waysToStep(n int) int {
    if n == 0 { return 1 }
    dp := make([]int, n + 1)
    dp[0] = 1
    if n >= 1 {
        dp[1] = 1
    }
    if n >= 2 {
        dp[2] = 2
    }
    for i := 3; i <= n; i++ {
        dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1_000_000_007
    }
    return dp[n]
}

// 动态规划，使用滚动数组优化空间
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func waysToStep1(n int) int {
    if n == 0 { return 1 }
    dp0, dp1, dp2 := 1, 1, 2
    if n == 1 { return dp1 }
    if n == 2 { return dp2 }
    for i := 3; i <= n; i++ {
        next := (dp0 + dp1 + dp2) % 1_000_000_007
        dp0, dp1, dp2 = dp1, dp2, next
    }
    return dp2
}

func main() {
    // Example1:
    // Input: n = 3 
    // Output: 4
    fmt.Println(waysToStep(3)) // 4
    // Example2:
    // Input: n = 5
    // Output: 13
    fmt.Println(waysToStep(5)) // 13

    fmt.Println(waysToStep(1)) // 1
    fmt.Println(waysToStep(2)) // 2
    fmt.Println(waysToStep(999)) // 887456284
    fmt.Println(waysToStep(1024)) // 54616240
    fmt.Println(waysToStep(999999)) // 71313044
    fmt.Println(waysToStep(1000000)) // 746580045

    fmt.Println(waysToStep1(3)) // 4
    fmt.Println(waysToStep1(5)) // 13
    fmt.Println(waysToStep1(1)) // 1
    fmt.Println(waysToStep1(2)) // 2
    fmt.Println(waysToStep1(999)) // 887456284
    fmt.Println(waysToStep1(1024)) // 54616240
    fmt.Println(waysToStep1(999999)) // 71313044
    fmt.Println(waysToStep1(1000000)) // 746580045
}