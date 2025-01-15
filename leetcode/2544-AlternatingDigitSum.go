package main

// 2544. Alternating Digit Sum
// You are given a positive integer n. 
// Each digit of n has a sign according to the following rules:
//     1. The most significant digit is assigned a positive sign.
//     2. Each other digit has an opposite sign to its adjacent digits.

// Return the sum of all digits with their corresponding sign.

// Example 1:
// Input: n = 521
// Output: 4
// Explanation: (+5) + (-2) + (+1) = 4.

// Example 2:
// Input: n = 111
// Output: 1
// Explanation: (+1) + (-1) + (+1) = 1.

// Example 3:
// Input: n = 886996
// Output: 0
// Explanation: (+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0.

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func alternateDigitSum(n int) int {
    res := 0
    for ; n != 0; n /= 10 {
        res = (n % 10 - res)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 521
    // Output: 4
    // Explanation: (+5) + (-2) + (+1) = 4.
    fmt.Println(alternateDigitSum(521)) // 4
    // Example 2:
    // Input: n = 111
    // Output: 1
    // Explanation: (+1) + (-1) + (+1) = 1.
    fmt.Println(alternateDigitSum(111)) // 1
    // Example 3:
    // Input: n = 886996
    // Output: 0
    // Explanation: (+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0.
    fmt.Println(alternateDigitSum(886996)) // 0

    fmt.Println(alternateDigitSum(0)) // 0
    fmt.Println(alternateDigitSum(1024)) // -1
    fmt.Println(alternateDigitSum(1_000_000_000)) // 1
    fmt.Println(alternateDigitSum(999_999_999)) // 9
}