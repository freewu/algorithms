package main

// 829. Consecutive Numbers Sum
// Given an integer n, return the number of ways you can write n as the sum of consecutive positive integers.

// Example 1:
// Input: n = 5
// Output: 2
// Explanation: 5 = 2 + 3

// Example 2:
// Input: n = 9
// Output: 3
// Explanation: 9 = 4 + 5 = 2 + 3 + 4

// Example 3:
// Input: n = 15
// Output: 4
// Explanation: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func consecutiveNumbersSum(n int) int {
    res := 0
    for i := 1; n > 0; i++ {
        if n % i == 0 {
            res++
        }
        n -= i
    }
    return res
}

func consecutiveNumbersSum1(n int) int {
    res := 0
    for i :=1; i*i < 2*n; i++ {
        if(2 * n % i == 0 && (2 * n / i-i+1) % 2 == 0) {// i 是 2n 的约数 且 2a 是偶数
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: 2
    // Explanation: 5 = 2 + 3
    fmt.Println(consecutiveNumbersSum(5)) // 2 (5 = 2 + 3)
    // Example 2:
    // Input: n = 9
    // Output: 3
    // Explanation: 9 = 4 + 5 = 2 + 3 + 4
    fmt.Println(consecutiveNumbersSum(9)) // 3 (9 = 4 + 5 = 2 + 3 + 4)
    // Example 3:
    // Input: n = 15
    // Output: 4
    // Explanation: 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
    fmt.Println(consecutiveNumbersSum(15)) // 3 (15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5)

    fmt.Println(consecutiveNumbersSum1(5)) // 2 (5 = 2 + 3)
    fmt.Println(consecutiveNumbersSum1(9)) // 3 (9 = 4 + 5 = 2 + 3 + 4)
    fmt.Println(consecutiveNumbersSum1(15)) // 3 (15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5)
}