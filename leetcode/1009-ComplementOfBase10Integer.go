package main

// 1009. Complement of Base 10 Integer
// The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
// For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
// Given an integer n, return its complement.

// Example 1:
// Input: n = 5
// Output: 2
// Explanation: 5 is "101" in binary, with complement "010" in binary, which is 2 in base-10.

// Example 2:
// Input: n = 7
// Output: 0
// Explanation: 7 is "111" in binary, with complement "000" in binary, which is 0 in base-10.

// Example 3:
// Input: n = 10
// Output: 5
// Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.
 
// Constraints:
//     0 <= n < 10^9

import "fmt"
import "math/bits"

func bitwiseComplement1(n int) int {
    return (1 << (bits.Len(uint(n))) - 1) ^ n
}

func bitwiseComplement(n int) int {
    if n == 0 {
        return 1
    }
    res, fac := 0, 1
    for n > 0 {
        res = res + (1- n % 2) * fac
        fac = fac * 2
        n = n / 2
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: 2
    // Explanation: 5 is "101" in binary, with complement "010" in binary, which is 2 in base-10.
    fmt.Println(bitwiseComplement(5)) // 2
    // Example 2:
    // Input: n = 7
    // Output: 0
    // Explanation: 7 is "111" in binary, with complement "000" in binary, which is 0 in base-10.
    fmt.Println(bitwiseComplement(7)) // 0
    // Example 3:
    // Input: n = 10
    // Output: 5
    // Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.
    fmt.Println(bitwiseComplement(10)) // 5
    fmt.Println(bitwiseComplement(1023)) // 0
    fmt.Println(bitwiseComplement(1024)) // 1023

    fmt.Println(bitwiseComplement1(5)) // 2
    fmt.Println(bitwiseComplement1(7)) // 0
    fmt.Println(bitwiseComplement1(10)) // 5
    fmt.Println(bitwiseComplement1(1023)) // 0
    fmt.Println(bitwiseComplement1(1024)) // 1023
}