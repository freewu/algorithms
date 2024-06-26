package main

// 191. Number of 1 Bits
// Write a function that takes the binary representation of a positive integer and returns the number of set bits it has (also known as the Hamming weight).

// Example 1:
// Input: n = 11
// Output: 3
// Explanation:
// The input binary string 1011 has a total of three set bits.

// Example 2:
// Input: n = 128
// Output: 1
// Explanation:
// The input binary string 10000000 has a total of one set bit.

// Example 3:
// Input: n = 2147483645
// Output: 30
// Explanation:
// The input binary string 1111111111111111111111111111101 has a total of thirty set bits.

// Constraints:
//     1 <= n <= 2^31 - 1

// Follow up: If this function is called many times, how would you optimize it?

import "fmt"
import "math/bits"

func hammingWeight(n int) int {
    res := 0
    for i := 0; i < 32; i++ {
        // fmt.Println((n & (1 << i)) >> i )
        if (n & (1 << i)) >> i == 1 { // 取第 i 位判断中否为 1
            res++
        }
    }
    return res
}

// lib
func hammingWeight1(n int) int {
    return bits.OnesCount32(uint32(n))
}

func hammingWeight2(n int) int {
    res := 0
    for n != 0 {
        if n & 1 == 1 {
            res++
        }
        n >>= 1
    }
    return res
}

func hammingWeight3(n int) int {
    res := 0
    for n != 0 {
        res += n % 2
        n = n >> 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 11
    // Output: 3
    // Explanation:
    // The input binary string 1011 has a total of three set bits.
    fmt.Println(hammingWeight(11)) // 3 (1011)
    // Example 2:
    // Input: n = 128
    // Output: 1
    // Explanation:
    // The input binary string 10000000 has a total of one set bit.
    fmt.Println(hammingWeight(128)) // 1 (10000000)
    // Example 3:
    // Input: n = 2147483645
    // Output: 30
    // Explanation:
    // The input binary string 1111111111111111111111111111101 has a total of thirty set bits.
    fmt.Println(hammingWeight(2147483645)) // 30 (1111111111111111111111111111101)

    fmt.Println(hammingWeight1(11)) // 3 (1011)
    fmt.Println(hammingWeight1(128)) // 1 (10000000)
    fmt.Println(hammingWeight1(2147483645)) // 30 (1111111111111111111111111111101)

    fmt.Println(hammingWeight2(11)) // 3 (1011)
    fmt.Println(hammingWeight2(128)) // 1 (10000000)
    fmt.Println(hammingWeight2(2147483645)) // 30 (1111111111111111111111111111101)

    fmt.Println(hammingWeight3(11)) // 3 (1011)
    fmt.Println(hammingWeight3(128)) // 1 (10000000)
    fmt.Println(hammingWeight3(2147483645)) // 30 (1111111111111111111111111111101)
}