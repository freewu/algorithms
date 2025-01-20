package main

// 2595. Number of Even and Odd Bits
// You are given a positive integer n.

// Let even denote the number of even indices in the binary representation of n with value 1.

// Let odd denote the number of odd indices in the binary representation of n with value 1.

// Note that bits are indexed from right to left in the binary representation of a number.

// Return the array [even, odd].

// Example 1:
// Input: n = 50
// Output: [1,2]
// Explanation:
// The binary representation of 50 is 110010.
// It contains 1 on indices 1, 4, and 5.

// Example 2:
// Input: n = 2
// Output: [0,1]
// Explanation:
// The binary representation of 2 is 10.
// It contains 1 only on index 1.

// Constraints:
//     1 <= n <= 1000

import "fmt"

func evenOddBit(n int) []int {
    even, odd, index := 0, 0, 0
    for n > 0 {
        if index % 2 == 0 && n & 1 == 1 {
            even++
        } else if n & 1 == 1 {
            odd++
        }
        index++
        n >>= 1
    }
    return []int{ even, odd }
}

func evenOddBit1(n int) []int {
    even, odd := 0, 0
    for i := 0; n != 0; i++ {
        if n & 1 == 1 {
            if i & 1 == 0 {
                even++
            } else {
                odd++
            }
        }
        n >>= 1
    }
    return []int{ even, odd }
}

func main() {
    // Example 1:
    // Input: n = 50
    // Output: [1,2]
    // Explanation:
    // The binary representation of 50 is 110010.
    // It contains 1 on indices 1, 4, and 5.
    fmt.Println(evenOddBit(50)) // [1,2]
    // Example 2:
    // Input: n = 2
    // Output: [0,1]
    // Explanation:
    // The binary representation of 2 is 10.
    // It contains 1 only on index 1.
    fmt.Println(evenOddBit(2)) // [0,1]

    fmt.Println(evenOddBit(1)) // [1 0]
    fmt.Println(evenOddBit(8)) // [0,1]
    fmt.Println(evenOddBit(64)) // [1 0]
    fmt.Println(evenOddBit(999)) // [4 4]
    fmt.Println(evenOddBit(1000)) // [2 4]

    fmt.Println(evenOddBit1(50)) // [1,2]
    fmt.Println(evenOddBit1(2)) // [0,1]
    fmt.Println(evenOddBit1(1)) // [1 0]
    fmt.Println(evenOddBit1(8)) // [0,1]
    fmt.Println(evenOddBit1(64)) // [1 0]
    fmt.Println(evenOddBit1(999)) // [4 4]
    fmt.Println(evenOddBit1(1000)) // [2 4]
}