package main

// 738. Monotone Increasing Digits
// An integer has monotone increasing digits if and only if each pair of adjacent digits x and y satisfy x <= y.
// Given an integer n, return the largest number that is less than or equal to n with monotone increasing digits.

// Example 1:
// Input: n = 10
// Output: 9

// Example 2:
// Input: n = 1234
// Output: 1234

// Example 3:
// Input: n = 332
// Output: 299
 
// Constraints:
//     0 <= n <= 10^9

import "fmt"

func monotoneIncreasingDigits(n int) int {
    res, i, digits := 0, 0, []int{}
    for n > 0 {
        digits = append(digits, n % 10)
        n = n / 10
    }
    for i < len(digits) - 1 {
        if digits[i] < digits[i+1] {
            for j := 0; j <= i; j++ {
                digits[j] = 9
            }
            digits[i+1]--
        }
        i++
    }
    for i := len(digits) - 1; i >= 0; i-- {
        res = res * 10 + digits[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 10
    // Output: 9
    fmt.Println(monotoneIncreasingDigits(10)) // 9
    // Example 2:
    // Input: n = 1234
    // Output: 1234
    fmt.Println(monotoneIncreasingDigits(1234)) // 1234
    // Example 3:
    // Input: n = 332
    // Output: 299
    fmt.Println(monotoneIncreasingDigits(332)) // 299
}