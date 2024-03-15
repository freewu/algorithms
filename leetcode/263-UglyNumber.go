package main

// 263. Ugly Number
// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
// Given an integer n, return true if n is an ugly number.

// Example 1:
// Input: n = 6
// Output: true
// Explanation: 6 = 2 × 3

// Example 2:
// Input: n = 1
// Output: true
// Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.

// Example 3:
// Input: n = 14
// Output: false
// Explanation: 14 is not ugly since it includes the prime factor 7.
 
// Constraints:
//     -2^31 <= n <= 2^31 - 1

import "fmt"

// 递归
func isUgly(n int) bool {
    if n == 0 {
        return false
    }
    if n == 1 {
        return true
    }
    if n % 2 == 0 { // 能被 2 整除
        return isUgly(n / 2)
    }
    if n % 3 == 0 { // 能被 3 整除
        return isUgly(n / 3)
    }
    if n % 5 == 0 { // 能被 5 整除
        return isUgly(n / 5)
    }
    return false
}

// 迭代
func isUgly1(n int) bool {
    if n <= 0 {
        return false
    }
    var factors = []int{2, 3, 5}
    for _, f := range factors {
        for n % f == 0 {
            n /= f
        }
    }
    return n == 1
}

func main() {
    fmt.Println(isUgly(1)) // true
    fmt.Println(isUgly(6)) // true
    fmt.Println(isUgly(14)) // false

    fmt.Println(isUgly1(1)) // true
    fmt.Println(isUgly1(6)) // true
    fmt.Println(isUgly1(14)) // false
}