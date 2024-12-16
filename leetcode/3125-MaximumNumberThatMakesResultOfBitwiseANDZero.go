package main

// 3125. Maximum Number That Makes Result of Bitwise AND Zero
// Given an integer n, return the maximum integer x such that x <= n, 
// and the bitwise AND of all the numbers in the range [x, n] is 0.
// Example 1:

// Input: n = 7
// Output: 3
// Explanation:
// The bitwise AND of [6, 7] is 6.
// The bitwise AND of [5, 6, 7] is 4.
// The bitwise AND of [4, 5, 6, 7] is 4.
// The bitwise AND of [3, 4, 5, 6, 7] is 0.

// Example 2:
// Input: n = 9
// Output: 7
// Explanation:
// The bitwise AND of [7, 8, 9] is 0.

// Example 3:
// Input: n = 17
// Output: 15
// Explanation:
// The bitwise AND of [15, 16, 17] is 0.

// Constraints:
//     1 <= n <= 10^15

import "fmt"
import "math/bits"

func maxNumber(n int64) int64 {
    return int64(1 << (bits.Len64(uint64(n)) - 1)) - 1
}

func maxNumber1(n int64) int64 {
    getBitsOfN := func(n int64) int64 {
        bits := int64(0)
        for n > 0 {
            n >>= 1
            bits++
        }
        return bits
    }
    return 1<<(getBitsOfN(n) - 1) - 1
}

func main() {
    // Input: n = 7
    // Output: 3
    // Explanation:
    // The bitwise AND of [6, 7] is 6.
    // The bitwise AND of [5, 6, 7] is 4.
    // The bitwise AND of [4, 5, 6, 7] is 4.
    // The bitwise AND of [3, 4, 5, 6, 7] is 0.
    fmt.Println(maxNumber(7)) // 3
    // Example 2:
    // Input: n = 9
    // Output: 7
    // Explanation:
    // The bitwise AND of [7, 8, 9] is 0.
    fmt.Println(maxNumber(9)) // 7
    // Example 3:
    // Input: n = 17
    // Output: 15
    // Explanation:
    // The bitwise AND of [15, 16, 17] is 0.
    fmt.Println(maxNumber(17)) // 15

    fmt.Println(maxNumber(1)) // 0
    fmt.Println(maxNumber(2)) // 1
    fmt.Println(maxNumber(64)) // 63
    fmt.Println(maxNumber(1024)) // 1023
    fmt.Println(maxNumber(99999)) // 65535
    fmt.Println(maxNumber(10e16)) // 72057594037927935

    fmt.Println(maxNumber1(7)) // 3
    fmt.Println(maxNumber1(9)) // 7
    fmt.Println(maxNumber1(17)) // 15
    fmt.Println(maxNumber1(1)) // 0
    fmt.Println(maxNumber1(2)) // 1
    fmt.Println(maxNumber1(64)) // 63
    fmt.Println(maxNumber1(1024)) // 1023
    fmt.Println(maxNumber1(99999)) // 65535
    fmt.Println(maxNumber1(10e16)) // 72057594037927935
}