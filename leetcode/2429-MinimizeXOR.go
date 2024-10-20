package main

// 2429. Minimize XOR
// Given two positive integers num1 and num2, find the positive integer x such that:
//     x has the same number of set bits as num2, and
//     The value x XOR num1 is minimal.

// Note that XOR is the bitwise XOR operation.

// Return the integer x. The test cases are generated such that x is uniquely determined.

// The number of set bits of an integer is the number of 1's in its binary representation.

// Example 1:
// Input: num1 = 3, num2 = 5
// Output: 3
// Explanation:
// The binary representations of num1 and num2 are 0011 and 0101, respectively.
// The integer 3 has the same number of set bits as num2, and the value 3 XOR 3 = 0 is minimal.

// Example 2:
// Input: num1 = 1, num2 = 12
// Output: 3
// Explanation:
// The binary representations of num1 and num2 are 0001 and 1100, respectively.
// The integer 3 has the same number of set bits as num2, and the value 3 XOR 1 = 2 is minimal.

// Constraints:
//     1 <= num1, num2 <= 10^9

import "fmt"
import "math/bits"

func minimizeXor(num1 int, num2 int) int {
    bitCount := func(x int) int { // return bit 1 count
        res := 0
        for x > 0 {
            x = x & (x - 1)
            res++
        }
        return res
    }
    n1bc, n2bc := bitCount(num1), bitCount(num2)
    for i := 0; i < n1bc - n2bc; i++ { // change 1 to 0 from low to high
        num1 &= (num1 - 1)
    }
    // change 0 to 1 form low to high
    for i := 0; i < n2bc - n1bc; i++ {
        num1 = num1 | (num1 + 1)
    }
    return num1
}

func minimizeXor1(num1 int, num2 int) int {
    n1bc, n2bc := bits.OnesCount(uint(num1)), bits.OnesCount(uint(num2))
    for i := 0; i < n1bc - n2bc; i++ { // change 1 to 0 from low to high
        num1 &= (num1 - 1)
    }
    // change 0 to 1 form low to high
    for i := 0; i < n2bc - n1bc; i++ {
        num1 = num1 | (num1 + 1)
    }
    return num1
}

func main() {
    // Example 1:
    // Input: num1 = 3, num2 = 5
    // Output: 3
    // Explanation:
    // The binary representations of num1 and num2 are 0011 and 0101, respectively.
    // The integer 3 has the same number of set bits as num2, and the value 3 XOR 3 = 0 is minimal.
    fmt.Println(minimizeXor(3,5)) // 3
    // Example 2:
    // Input: num1 = 1, num2 = 12
    // Output: 3
    // Explanation:
    // The binary representations of num1 and num2 are 0001 and 1100, respectively.
    // The integer 3 has the same number of set bits as num2, and the value 3 XOR 1 = 2 is minimal.
    fmt.Println(minimizeXor(1,12)) // 3

    fmt.Println(minimizeXor(1,1)) // 1
    fmt.Println(minimizeXor(1,1_000_000_000)) // 8191
    fmt.Println(minimizeXor(1_000_000_000,1_000_000_000)) // 1000000000
    fmt.Println(minimizeXor(1_000_000_000,1)) // 536870912

    fmt.Println(minimizeXor1(3,5)) // 3
    fmt.Println(minimizeXor1(1,12)) // 3
    fmt.Println(minimizeXor1(1,1)) // 1
    fmt.Println(minimizeXor1(1,1_000_000_000)) // 8191
    fmt.Println(minimizeXor1(1_000_000_000,1_000_000_000)) // 1000000000
    fmt.Println(minimizeXor1(1_000_000_000,1)) // 536870912
}