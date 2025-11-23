package main

// 3750. Minimum Number of Flips to Reverse Binary String
// You are given a positive integer n.

// Let s be the binary representation of n without leading zeros.

// The reverse of a binary string s is obtained by writing the characters of s in the opposite order.

// You may flip any bit in s (change 0 → 1 or 1 → 0). Each flip affects exactly one bit.

// Return the minimum number of flips required to make s equal to the reverse of its original form.

// Example 1:
// Input: n = 7
// Output: 0
// Explanation:
// The binary representation of 7 is "111".
// Its reverse is also "111", which is the same. 
// Hence, no flips are needed.

// Example 2:
// Input: n = 10
// Output: 4
// Explanation:
// The binary representation of 10 is "1010". Its reverse is "0101". 
// All four bits must be flipped to make them equal. 
// Thus, the minimum number of flips required is 4.

// Constraints:
//     1 <= n <= 10^9

import "fmt"

func minimumFlips(n int) int {
    res, r, m := 0, 0, n
    for m > 0 { // Reverse bits of n into r
        r <<= 1
        r |= (m & 1)
        m >>= 1
    }
    for n > 0 { // Count flips needed to match reversed bits
        if (n & 1) != (r & 1) {
            res++
        }
        n >>= 1
        r >>= 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 7
    // Output: 0
    // Explanation:
    // The binary representation of 7 is "111".
    // Its reverse is also "111", which is the same. 
    // Hence, no flips are needed.
    fmt.Println(minimumFlips(7)) // 0
    // Example 2:
    // Input: n = 10
    // Output: 4
    // Explanation:
    // The binary representation of 10 is "1010". Its reverse is "0101". 
    // All four bits must be flipped to make them equal. 
    // Thus, the minimum number of flips required is 4.
    fmt.Println(minimumFlips(10)) // 4

    fmt.Println(minimumFlips(1)) // 0
    fmt.Println(minimumFlips(8)) // 2
    fmt.Println(minimumFlips(64)) // 2
    fmt.Println(minimumFlips(999)) // 4
    fmt.Println(minimumFlips(1000)) // 8
    fmt.Println(minimumFlips(1024)) // 2
    fmt.Println(minimumFlips(999_999_999)) // 14
    fmt.Println(minimumFlips(1_000_000_000)) // 18
}