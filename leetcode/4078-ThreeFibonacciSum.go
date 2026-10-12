package main

// 4078. Three Fibonacci Sum
// You are given an integer n.

// The Fibonacci sequence starts with 0 and 1, and each subsequent number is the sum of the previous two numbers. 
// Its first few terms are 0, 1, 1, 2, 3, 5, 8, 13, ....

// Return true if n can be expressed as the sum of three consecutive terms in the Fibonacci sequence, and false otherwise.

// Example 1:
// Input: n = 16
// Output: true
// Explanation:
// The three consecutive terms 3, 5, and 8 have a sum of 3 + 5 + 8 = 16.

// Example 2:
// Input: n = 8
// Output: false
// Explanation:
// Although 1 + 2 + 5 = 8, these terms are not consecutive in the Fibonacci sequence. 
// No three consecutive terms have a sum of 8.

// Example 3:
// Input: n = 2
// Output: true
// Explanation:
// The first three terms have a sum of 0 + 1 + 1 = 2.

// Constraints:
//     - 1 <= n <= 10^{9}

import "fmt"

func threeFibonacciSum(n int) bool {
    for a, b := 0, 1 ; 2 * (a + b) <= n; a, b = b, a + b {
        if 2 * (a + b) == n {
            return true
        }
    }
    return false
}

func threeFibonacciSum1(n int) bool {
    a, b, c := 0, 1, 1
    for a + b + c < n {
        a, b, c = b, c, a + b
    }
    return a + b + c == n
}

func main() {
    // Example 1:
    // Input: n = 16
    // Output: true
    // Explanation:
    // The three consecutive terms 3, 5, and 8 have a sum of 3 + 5 + 8 = 16.
    fmt.Println(threeFibonacciSum(16)) // true
    // Example 2:
    // Input: n = 8
    // Output: false
    // Explanation:
    // Although 1 + 2 + 5 = 8, these terms are not consecutive in the Fibonacci sequence. 
    // No three consecutive terms have a sum of 8.
    fmt.Println(threeFibonacciSum(8)) // false
    // Example 3:
    // Input: n = 2
    // Output: true
    // Explanation:
    // The first three terms have a sum of 0 + 1 + 1 = 2.
    fmt.Println(threeFibonacciSum(2)) // true

    fmt.Println(threeFibonacciSum(1)) // false
    fmt.Println(threeFibonacciSum(3)) // false
    fmt.Println(threeFibonacciSum(5)) // false
    fmt.Println(threeFibonacciSum(64)) // false
    fmt.Println(threeFibonacciSum(99)) // false
    fmt.Println(threeFibonacciSum(100)) // false
    fmt.Println(threeFibonacciSum(101)) // false
    fmt.Println(threeFibonacciSum(1024)) // false
    fmt.Println(threeFibonacciSum(999_999)) // false
    fmt.Println(threeFibonacciSum(1_000_000)) // false

    fmt.Println(threeFibonacciSum1(16)) // true
    fmt.Println(threeFibonacciSum1(8)) // false
    fmt.Println(threeFibonacciSum1(2)) // true
    fmt.Println(threeFibonacciSum1(1)) // false
    fmt.Println(threeFibonacciSum1(3)) // false
    fmt.Println(threeFibonacciSum1(5)) // false
    fmt.Println(threeFibonacciSum1(64)) // false
    fmt.Println(threeFibonacciSum1(99)) // false
    fmt.Println(threeFibonacciSum1(100)) // false
    fmt.Println(threeFibonacciSum1(101)) // false
    fmt.Println(threeFibonacciSum1(1024)) // false
    fmt.Println(threeFibonacciSum1(999_999)) // false
    fmt.Println(threeFibonacciSum1(1_000_000)) // false
}