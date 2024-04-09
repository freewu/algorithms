package main

// 509. Fibonacci Number
// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, 
// such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
//     F(0) = 0, F(1) = 1
//     F(n) = F(n - 1) + F(n - 2), for n > 1.

// Given n, calculate F(n).

// Example 1:
// Input: n = 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

// Example 2:
// Input: n = 3
// Output: 2
// Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

// Example 3:
// Input: n = 4
// Output: 3
// Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
 
// Constraints:
//     0 <= n <= 30

import "fmt"

// 递归
func fib(n int) int {
    if 0 == n {
        return 0
    }
    if 1 == n || 2 == n{
        return 1
    }
    return fib(n - 1) + fib(n - 2)
}

// 缓存迭代
func fib1(n int) int {
    if 0 == n {
        return 0
    }
    if 1 == n || 2 == n{
        return 1
    }
    dp := make([]int, n + 1)
    dp[1], dp[2] = 1, 1
    for i := 3; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[n]
}

func main() {
    fmt.Println(fib(1)) // 1
    fmt.Println(fib(2)) // 1
    fmt.Println(fib(3)) // 2
    fmt.Println(fib(4)) // 3
    fmt.Println(fib(5)) // 5
    fmt.Println(fib(6)) // 8
    fmt.Println(fib(7)) // 13

    fmt.Println(fib1(1)) // 1
    fmt.Println(fib1(2)) // 1
    fmt.Println(fib1(3)) // 2
    fmt.Println(fib1(4)) // 3
    fmt.Println(fib1(5)) // 5
    fmt.Println(fib1(6)) // 8
    fmt.Println(fib1(7)) // 13
}