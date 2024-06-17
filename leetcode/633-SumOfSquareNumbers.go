package main

// 633. Sum of Square Numbers
// Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

// Example 1:
// Input: c = 5
// Output: true
// Explanation: 1 * 1 + 2 * 2 = 5

// Example 2:
// Input: c = 3
// Output: false
 
// Constraints:
//     0 <= c <= 2^31 - 1

import "fmt"
import "math"

// 双指针
func judgeSquareSum(c int) bool {
    low, high := 0, int(math.Sqrt(float64(c)))
    for t := high * high + low * low; low <= high && t != c; t = high * high + low * low {
        if t < c {
            low++
        } else {
            high--
        }
    }
    return c == high * high + low * low
}

func main() {
    // Example 1:
    // Input: c = 5
    // Output: true
    // Explanation: 1 * 1 + 2 * 2 = 5
    fmt.Println(judgeSquareSum(5)) // true
    // Example 2:
    // Input: c = 3
    // Output: false
    fmt.Println(judgeSquareSum(3)) // false
}