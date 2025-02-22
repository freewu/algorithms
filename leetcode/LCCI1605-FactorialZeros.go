package main

// 面试题 16.05. Factorial Zeros LCCI
// Write an algorithm which computes the number of trailing zeros in n factorial.

// Example 1:
// Input: 3
// Output: 0
// Explanation: 3! = 6, no trailing zero.

// Example 2:
// Input: 5
// Output: 1
// Explanation: 5! = 120, one trailing zero.
// Note: Your solution should be in logarithmic time complexity.

import "fmt"

func trailingZeroes(n int) int {
    res := 0
    for n > 0 {
        n /= 5
        res += n
    }
    return res
}

func main() {
    // Example 1:
    // Input: 3
    // Output: 0
    // Explanation: 3! = 6, no trailing zero.
    fmt.Println(trailingZeroes(3)) // 0
    // Example 2:
    // Input: 5
    // Output: 1
    // Explanation: 5! = 120, one trailing zero.
    // Note: Your solution should be in logarithmic time complexity.
    fmt.Println(trailingZeroes(5)) // 1

    fmt.Println(trailingZeroes(999)) // 246
    fmt.Println(trailingZeroes(1024)) // 253
}