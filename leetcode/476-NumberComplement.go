package main

// 476. Number Complement
// The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
// For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
// Given an integer num, return its complement.

// Example 1:
// Input: num = 5
// Output: 2
// Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.

// Example 2:
// Input: num = 1
// Output: 0
// Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
 
// Constraints:
//     1 <= num < 2^31

import "fmt"
import "math/bits"

func findComplement(num int) int {
    res, fac := 0, 1
    for num > 0 {
        res = res + (1-num % 2) * fac
        fac = fac * 2
        num = num / 2
    }
    return res
}

func findComplement1(num int) int {
    return (1 << (bits.Len(uint(num))) - 1) ^ num
}

func main() {
    // Example 1:
    // Input: num = 5
    // Output: 2
    // Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
    fmt.Println(findComplement(5)) // 2
    // Example 2:
    // Input: num = 1
    // Output: 0
    // Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.
    fmt.Println(findComplement(1)) // 0

    fmt.Println(findComplement1(5)) // 2
    fmt.Println(findComplement1(1)) // 0
}