package main

// 3754. Concatenate Non-Zero Digits and Multiply by Sum I
// You are given an integer n.

// Form a new integer x by concatenating all the non-zero digits of n in their original order. If there are no non-zero digits, x = 0.

// Let sum be the sum of digits in x.

// Return an integer representing the value of x * sum.

// Example 1:
// Input: n = 10203004
// Output: 12340
// Explanation:
// The non-zero digits are 1, 2, 3, and 4. Thus, x = 1234.
// The sum of digits is sum = 1 + 2 + 3 + 4 = 10.
// Therefore, the answer is x * sum = 1234 * 10 = 12340.

// Example 2:
// Input: n = 1000
// Output: 1
// Explanation:
// The non-zero digit is 1, so x = 1 and sum = 1.
// Therefore, the answer is x * sum = 1 * 1 = 1.

// Constraints:
//     0 <= n <= 10^9

import "fmt"

func sumAndMultiply(n int) int64 {
    s, x, k := 0, 0, 1
    for n > 0 {
        if n % 10 > 0 { 
            x += k * (n % 10)
            k *= 10 
        }
        s += n % 10
        n = n / 10
    }
    return int64(s) * int64(x)
}

func main() {
    // Example 1:
    // Input: n = 10203004
    // Output: 12340
    // Explanation:
    // The non-zero digits are 1, 2, 3, and 4. Thus, x = 1234.
    // The sum of digits is sum = 1 + 2 + 3 + 4 = 10.
    // Therefore, the answer is x * sum = 1234 * 10 = 12340.
    fmt.Println(sumAndMultiply(10203004)) // 12340
    // Example 2:
    // Input: n = 1000
    // Output: 1
    // Explanation:
    // The non-zero digit is 1, so x = 1 and sum = 1.
    // Therefore, the answer is x * sum = 1 * 1 = 1.
    fmt.Println(sumAndMultiply(1000)) // 1

    fmt.Println(sumAndMultiply(0)) // 0
    fmt.Println(sumAndMultiply(1)) // 1
    fmt.Println(sumAndMultiply(8)) // 64
    fmt.Println(sumAndMultiply(64)) // 640
    fmt.Println(sumAndMultiply(99)) // 1782
    fmt.Println(sumAndMultiply(100)) // 1
    fmt.Println(sumAndMultiply(1024)) // 868
    fmt.Println(sumAndMultiply(999_999_999)) // 80999999919
    fmt.Println(sumAndMultiply(1_000_000_000)) // 1
}