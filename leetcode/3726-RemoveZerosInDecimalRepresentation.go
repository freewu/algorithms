package main

// 3726. Remove Zeros in Decimal Representation
// You are given a positive integer n.

// Return the integer obtained by removing all zeros from the decimal representation of n.

// Example 1:
// Input: n = 1020030
// Output: 123
// Explanation:
// After removing all zeros from 1020030, we get 123.

// Example 2:
// Input: n = 1
// Output: 1
// Explanation:
// 1 has no zero in its decimal representation. Therefore, the answer is 1.

// Constraints:
//     1 <= n <= 10^15

import "fmt"

func removeZeros(n int64) int64 {
    res, mult := int64(0), int64(1)
    for n > 0 {
        rem := n % 10
        if rem != 0 {
            res += rem * mult
            mult *= 10
        }
        n /= 10
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 1020030
    // Output: 123
    // Explanation:
    // After removing all zeros from 1020030, we get 123.
    fmt.Println(removeZeros(1020030)) // 123
    // Example 2:
    // Input: n = 1
    // Output: 1
    // Explanation:
    // 1 has no zero in its decimal representation. Therefore, the answer is 1.
    fmt.Println(removeZeros(1)) // 1

    fmt.Println(removeZeros(8)) // 8
    fmt.Println(removeZeros(64)) // 64
    fmt.Println(removeZeros(99)) // 99
    fmt.Println(removeZeros(100)) // 1
    fmt.Println(removeZeros(1024)) // 124
    fmt.Println(removeZeros(1_000_000_007)) // 17
    fmt.Println(removeZeros(1_000_000_000_000_000)) // 1
}