package main

// 507. Perfect Number
// A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself. 
// A divisor of an integer x is an integer that can divide x evenly.

// Given an integer n, return true if n is a perfect number, otherwise return false.

// Example 1:
// Input: num = 28
// Output: true
// Explanation: 28 = 1 + 2 + 4 + 7 + 14
// 1, 2, 4, 7, and 14 are all divisors of 28.

// Example 2:
// Input: num = 7
// Output: false

// Constraints:
//     1 <= num <= 10^8

import "fmt"
import "math"

// 对于一个 正整数，如果它和除了它自身以外的所有正因子之和相等，我们称它为“完美数”
func checkPerfectNumber(num int) bool {
    if num <= 1 {
        return false
    }
    sum, bound := 1, int(math.Sqrt(float64(num))) + 1
    for i := 2; i < bound; i++ {
        if num % i != 0 {
            continue
        }
        corrDiv := num / i
        sum += corrDiv + i
    }
    return sum == num
}

// 打表
func checkPerfectNumber1(num int) bool {
    return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}

func main() {
    // Example 1:
    // Input: num = 28
    // Output: true
    // Explanation: 28 = 1 + 2 + 4 + 7 + 14
    // 1, 2, 4, 7, and 14 are all divisors of 28.
    fmt.Println(checkPerfectNumber(28)) // true
    // Example 2:
    // Input: num = 7
    // Output: false
    fmt.Println(checkPerfectNumber(7)) // false

    fmt.Println(checkPerfectNumber1(28)) // true
    fmt.Println(checkPerfectNumber1(7)) // false
}