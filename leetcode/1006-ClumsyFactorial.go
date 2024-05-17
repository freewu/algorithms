package main

// 1006. Clumsy Factorial
// The factorial of a positive integer n is the product of all positive integers less than or equal to n.
//     For example, factorial(10) = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1.

// We make a clumsy factorial using the integers in decreasing order by swapping out the multiply operations 
// for a fixed rotation of operations with multiply '*', divide '/', add '+', and subtract '-' in this order.
//     For example, clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1.

// However, these operations are still applied using the usual order of operations of arithmetic. 
// We do all multiplication and division steps before any addition or subtraction steps, 
// and multiplication and division steps are processed left to right.

// Additionally, the division that we use is floor division such that 10 * 9 / 8 = 90 / 8 = 11.
// Given an integer n, return the clumsy factorial of n.

// Example 1:
// Input: n = 4
// Output: 7
// Explanation: 7 = 4 * 3 / 2 + 1

// Example 2:
// Input: n = 10
// Output: 12
// Explanation: 12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func clumsy(n int) int {
    var f func (n int) int
    f = func (n int) int {
        switch n {
        case 1:
            return 1
        case 2:
            return 2 - 1
        case 3:
            return 3 - 2 * 1
        case 4:
            return 4 - 3 * 2 / 1
        default:
            return n - (n - 1) * (n - 2) / (n - 3) + f(n - 4)
        }
    }
    switch n {
        case 1:
            return 1
        case 2:
            return 2 * 1
        case 3:
            return 3 * 2 / 1
        default:
            return n * (n - 1) / (n - 2) + f(n - 3)
    }
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: 7
    // Explanation: 7 = 4 * 3 / 2 + 1
    fmt.Println(clumsy(4)) // 7
    // Example 2:
    // Input: n = 10
    // Output: 12
    // Explanation: 12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1
    fmt.Println(clumsy(10)) // 12
}