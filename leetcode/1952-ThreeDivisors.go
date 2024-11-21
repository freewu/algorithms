package main

// 1952. Three Divisors
// Given an integer n, return true if n has exactly three positive divisors. 
// Otherwise, return false.

// An integer m is a divisor of n if there exists an integer k such that n = k * m.

// Example 1:
// Input: n = 2
// Output: false
// Explantion: 2 has only two divisors: 1 and 2.

// Example 2:
// Input: n = 4
// Output: true
// Explantion: 4 has three divisors: 1, 2, and 4.

// Constraints:
//     1 <= n <= 10^4

import "fmt"

func isThree(n int) bool {
    count := 0
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            count++
        } else if 3 < count {
            break
        }
    }
    return count == 3
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: false
    // Explantion: 2 has only two divisors: 1 and 2.
    fmt.Println(isThree(2)) // false
    // Example 2:
    // Input: n = 4
    // Output: true
    // Explantion: 4 has three divisors: 1, 2, and 4.
    fmt.Println(isThree(4)) // true

    fmt.Println(isThree(1)) // false
    fmt.Println(isThree(1024)) // false
    fmt.Println(isThree(17)) // false
    fmt.Println(isThree(4000)) // false
}