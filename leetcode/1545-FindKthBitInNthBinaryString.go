package main

// 1545. Find Kth Bit in Nth Binary String
// Given two positive integers n and k, the binary string Sn is formed as follows:
//     S1 = "0"
//     Si = Si - 1 + "1" + reverse(invert(Si - 1)) for i > 1

// Where + denotes the concatenation operation, reverse(x) returns the reversed string x, 
// and invert(x) inverts all the bits in x (0 changes to 1 and 1 changes to 0).

// For example, the first four strings in the above sequence are:
//     S1 = "0"
//     S2 = "011"
//     S3 = "0111001"
//     S4 = "011100110110001"

// Return the kth bit in Sn. It is guaranteed that k is valid for the given n.

// Example 1:
// Input: n = 3, k = 1
// Output: "0"
// Explanation: S3 is "0111001".
// The 1st bit is "0".

// Example 2:
// Input: n = 4, k = 11
// Output: "1"
// Explanation: S4 is "011100110110001".
// The 11th bit is "1".

// Constraints:
//     1 <= n <= 20
//     1 <= k <= 2^n - 1

import "fmt"

func findKthBit(n int, k int) byte {
    if n == 1 || k == 1 { return '0' }
    l := 1 << n - 1 // 长度为 2 的 n 次方 - 1
    mid := l >> 1 + 1 // 15-8,7-4,3-2，中间位置为l/2+1
    if k == mid { // 最中间就是1
        return '1'
    } else if k < mid {
        return findKthBit(n-1, k)
    } else {
        invert := func(ch byte) byte { if ch == '0' { return '1'; } ; return '0'; }
        return invert(findKthBit(n-1, l-k+1)) // 15-1,7-1,3-1，reverse 之前的位置为 l - k + 1
    }
}

func findKthBit1(n, k int) byte {
    var solve func(n, k int) byte
    solve = func(n, k int) byte {
        if n == 1 || k == 1 { return '0' }
        mid := 1 << (n - 1)
        if k == mid { return '1' }
        if k < mid  { return solve(n-1, k) }
        k = mid * 2 - k
        if solve(n-1, k) == '1' {
            return '0'
        }
        return '1'
    }
    return solve(n, k)
}

func main() {
    // Example 1:
    // Input: n = 3, k = 1
    // Output: "0"
    // Explanation: S3 is "0111001".
    // The 1st bit is "0".
    fmt.Printf("%c\n", findKthBit(3,1)) // '0'
    // Example 2:
    // Input: n = 4, k = 11
    // Output: "1"
    // Explanation: S4 is "011100110110001".
    // The 11th bit is "1".
    fmt.Printf("%c\n", findKthBit(4,11)) // '1'

    fmt.Printf("%c\n", findKthBit1(3,1)) // '0'
    fmt.Printf("%c\n", findKthBit1(4,11)) // '1'
} 