package main

// 3345. Smallest Divisible Digit Product I
// You are given two integers n and t. 
// Return the smallest number greater than or equal to n such that the product of its digits is divisible by t.

// Example 1:
// Input: n = 10, t = 2
// Output: 10
// Explanation:
// The digit product of 10 is 0, which is divisible by 2, making it the smallest number greater than or equal to 10 that satisfies the condition.

// Example 2:
// Input: n = 15, t = 3
// Output: 16
// Explanation:
// The digit product of 16 is 6, which is divisible by 3, making it the smallest number greater than or equal to 15 that satisfies the condition.

// Constraints:
//     1 <= n <= 100
//     1 <= t <= 10

import "fmt"

func smallestNumber(n int, t int) int {
    productOfDigits := func(n int) int {
        product := 1
        for n > 0 {
            product *= (n % 10)
            n = n / 10
        }
        return product
    }
    for i := n; ; i++ {
        if productOfDigits(i) % t == 0 {
            return i
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 10, t = 2
    // Output: 10
    // Explanation:
    // The digit product of 10 is 0, which is divisible by 2, making it the smallest number greater than or equal to 10 that satisfies the condition.
    fmt.Println(smallestNumber(10, 2)) // 10
    // Example 2:
    // Input: n = 15, t = 3
    // Output: 16
    // Explanation:
    // The digit product of 16 is 6, which is divisible by 3, making it the smallest number greater than or equal to 15 that satisfies the condition.
    fmt.Println(smallestNumber(15, 3)) // 16

    fmt.Println(smallestNumber(100, 10)) // 100
    fmt.Println(smallestNumber(1, 1)) // 1
    fmt.Println(smallestNumber(1, 10)) // 10
    fmt.Println(smallestNumber(100, 1)) // 100
}