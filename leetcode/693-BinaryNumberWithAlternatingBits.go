package main

// 693. Binary Number with Alternating Bits
// Given a positive integer, check whether it has alternating bits: namely, 
// if two adjacent bits will always have different values.

// Example 1:
// Input: n = 5
// Output: true
// Explanation: The binary representation of 5 is: 101

// Example 2:
// Input: n = 7
// Output: false
// Explanation: The binary representation of 7 is: 111.

// Example 3:
// Input: n = 11
// Output: false
// Explanation: The binary representation of 11 is: 1011.

// Constraints:
//     1 <= n <= 2^31 - 1

import "fmt"

func hasAlternatingBits(n int) bool {
    i, swap := 1, false
    for i < n {
        if swap {
            i = 2 * i + 1
        } else {
            i *= 2
        }
        swap = !swap // 交替
    }
    return i == n
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: true
    // Explanation: The binary representation of 5 is: 101
    fmt.Println(hasAlternatingBits(5)) // true  101
    // Example 2:
    // Input: n = 7
    // Output: false
    // Explanation: The binary representation of 7 is: 111.
    fmt.Println(hasAlternatingBits(7)) // false 111
    // Example 3:
    // Input: n = 11
    // Output: false
    // Explanation: The binary representation of 11 is: 1011.
    fmt.Println(hasAlternatingBits(11)) // false  1011

    fmt.Println(hasAlternatingBits(1)) //
    fmt.Println(hasAlternatingBits(999)) // false 111111111
    fmt.Println(hasAlternatingBits(1024)) // false 
    fmt.Println(hasAlternatingBits(1_000_000_007)) // false 
    fmt.Println(hasAlternatingBits(2 << 31 - 1)) // false 
}