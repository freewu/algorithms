package main

// 660. Remove 9
// Start from integer 1, remove any integer that contains 9 such as 9, 19, 29...
// Now, you will have a new integer sequence [1, 2, 3, 4, 5, 6, 7, 8, 10, 11, ...].
// Given an integer n, return the nth (1-indexed) integer in the new sequence.

// Example 1:
// Input: n = 9
// Output: 10

// Example 2:
// Input: n = 10
// Output: 11

// Constraints:
//     1 <= n <= 8 * 10^8

import "fmt"

// 九进制
func newInteger(n int) int {
    if n == 0 {
        return 0
    }
    return newInteger(n / 9) *10 + n % 9
}

func main() {
    // Example 1:
    // Input: n = 9
    // Output: 10
    fmt.Println(newInteger(9)) // 10
    // Example 2:
    // Input: n = 10
    // Output: 11
    fmt.Println(newInteger(10)) // 11
    fmt.Println(newInteger(1024)) // 1357
}