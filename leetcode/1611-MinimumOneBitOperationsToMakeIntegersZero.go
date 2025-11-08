package main

// 1611. Minimum One Bit Operations to Make Integers Zero
// Given an integer n, you must transform it into 0 using the following operations any number of times:
//     1. Change the rightmost (0th) bit in the binary representation of n.
//     2. Change the ith bit in the binary representation of n if the (i-1)th bit is set to 1 and the (i-2)th through 0th bits are set to 0.

// Return the minimum number of operations to transform n into 0.

// Example 1:
// Input: n = 3
// Output: 2
// Explanation: The binary representation of 3 is "11".
// "11" -> "01" with the 2nd operation since the 0th bit is 1.
// "01" -> "00" with the 1st operation.

// Example 2:
// Input: n = 6
// Output: 4
// Explanation: The binary representation of 6 is "110".
// "110" -> "010" with the 2nd operation since the 1st bit is 1 and 0th through 0th bits are 0.
// "010" -> "011" with the 1st operation.
// "011" -> "001" with the 2nd operation since the 0th bit is 1.
// "001" -> "000" with the 1st operation.

// Constraints:
//     0 <= n <= 10^9

import "fmt"
import "math/bits"

func minimumOneBitOperations1(n int) int {
    res, curr := 0, 1
    for n != 0 {
        if n % 2 == 1 {
            res = curr - res
        }
        curr = 2 * curr + 1
        n /= 2
    }
    return res
}

func minimumOneBitOperations(n int) int {
    if n <= 1 { return n  }
    m := bits.Len(uint(n))
    return 1 << m - 1- minimumOneBitOperations(1 << (m - 1) ^ n)
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 2
    // Explanation: The binary representation of 3 is "11".
    // "11" -> "01" with the 2nd operation since the 0th bit is 1.
    // "01" -> "00" with the 1st operation.
    fmt.Println(minimumOneBitOperations(3)) // 2
    // Example 2:
    // Input: n = 6
    // Output: 4
    // Explanation: The binary representation of 6 is "110".
    // "110" -> "010" with the 2nd operation since the 1st bit is 1 and 0th through 0th bits are 0.
    // "010" -> "011" with the 1st operation.
    // "011" -> "001" with the 2nd operation since the 0th bit is 1.
    // "001" -> "000" with the 1st operation.
    fmt.Println(minimumOneBitOperations(6)) // 4

    fmt.Println(minimumOneBitOperations(0)) // 0
    fmt.Println(minimumOneBitOperations(1)) // 1
    fmt.Println(minimumOneBitOperations(8)) // 15
    fmt.Println(minimumOneBitOperations(64)) // 127
    fmt.Println(minimumOneBitOperations(99)) // 66
    fmt.Println(minimumOneBitOperations(100)) // 71
    fmt.Println(minimumOneBitOperations(128)) // 255
    fmt.Println(minimumOneBitOperations(1024)) // 2047
    fmt.Println(minimumOneBitOperations(99999999)) // 111591253
    fmt.Println(minimumOneBitOperations(100000000)) // 111590912

    fmt.Println(minimumOneBitOperations1(3)) // 2
    fmt.Println(minimumOneBitOperations1(6)) // 4
    fmt.Println(minimumOneBitOperations1(0)) // 0
    fmt.Println(minimumOneBitOperations1(1)) // 1
    fmt.Println(minimumOneBitOperations1(8)) // 15
    fmt.Println(minimumOneBitOperations1(64)) // 127
    fmt.Println(minimumOneBitOperations1(99)) // 66
    fmt.Println(minimumOneBitOperations1(100)) // 71
    fmt.Println(minimumOneBitOperations1(128)) // 255
    fmt.Println(minimumOneBitOperations1(1024)) // 2047
    fmt.Println(minimumOneBitOperations1(99999999)) // 111591253
    fmt.Println(minimumOneBitOperations1(100000000)) // 111590912
}