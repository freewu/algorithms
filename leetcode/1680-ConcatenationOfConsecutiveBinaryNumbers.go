package main

// 1680. Concatenation of Consecutive Binary Numbers
// Given an integer n, return the decimal value of the binary string formed 
// by concatenating the binary representations of 1 to n in order, modulo 10^9 + 7.

// Example 1:
// Input: n = 1
// Output: 1
// Explanation: "1" in binary corresponds to the decimal value 1.

// Example 2:
// Input: n = 3
// Output: 27
// Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
// After concatenating them, we have "11011", which corresponds to the decimal value 27.

// Example 3:
// Input: n = 12
// Output: 505379714
// Explanation: The concatenation results in "1101110010111011110001001101010111100".
// The decimal value of that is 118505380540.
// After modulo 10^9 + 7, the result is 505379714.

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func concatenatedBinary(n int) int {
    res, curr, count := 1, 2, 1
    for i := 2; i <= n; i++ {
        if i == curr {
            curr <<= 1
            count++
        } 
        res = ((res << count) + i) % 1_000_000_007
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 1
    // Output: 1
    // Explanation: "1" in binary corresponds to the decimal value 1.
    fmt.Println(concatenatedBinary(1)) // 1
    // Example 2:
    // Input: n = 3
    // Output: 27
    // Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
    // After concatenating them, we have "11011", which corresponds to the decimal value 27.
    fmt.Println(concatenatedBinary(3)) // 27
    // Example 3:
    // Input: n = 12
    // Output: 505379714
    // Explanation: The concatenation results in "1101110010111011110001001101010111100".
    // The decimal value of that is 118505380540.
    // After modulo 10^9 + 7, the result is 505379714.
    fmt.Println(concatenatedBinary(12)) // 505379714

    fmt.Println(concatenatedBinary(1024)) // 41183187
    fmt.Println(concatenatedBinary(10000)) // 356435599
    fmt.Println(concatenatedBinary(9999)) // 987753695
}