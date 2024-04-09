package main

// 1137. N-th Tribonacci Number
// The Tribonacci sequence Tn is defined as follows: 
//     T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

// Given n, return the value of Tn.

// Example 1:
// Input: n = 4
// Output: 4
// Explanation:
// T_3 = 0 + 1 + 1 = 2
// T_4 = 1 + 1 + 2 = 4

// Example 2:
// Input: n = 25
// Output: 1389537
 
// Constraints:
//     0 <= n <= 37
//     The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.

import "fmt"

func tribonacci(n int) int {
    if 0 == n {
        return 0
    }
    if 1 == n || 2 == n{
        return 1
    }
    if 3 == n {
        return 2
    }
    dp := make([]int, n + 1)
    dp[1], dp[2], dp[3] = 1, 1, 2
    for i := 4; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3]
    }
    return dp[n]
}

func main() {
    fmt.Println(tribonacci(1)) // 1
    fmt.Println(tribonacci(2)) // 1
    fmt.Println(tribonacci(3)) // 2
    fmt.Println(tribonacci(4)) // 4
    fmt.Println(tribonacci(5)) // 7
    fmt.Println(tribonacci(6)) // 13
    fmt.Println(tribonacci(7)) // 24
    fmt.Println(tribonacci(8)) // 44
    fmt.Println(tribonacci(25)) // 1389537
}