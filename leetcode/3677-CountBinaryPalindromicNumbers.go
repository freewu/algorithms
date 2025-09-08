package main

// 3677. Count Binary Palindromic Numbers
// You are given a non-negative integer n.

// A non-negative integer is called binary-palindromic if its binary representation (written without leading zeros) reads the same forward and backward.

// Return the number of integers k such that 0 <= k <= n and the binary representation of k is a palindrome.

// Note: The number 0 is considered binary-palindromic, and its representation is "0".

// Example 1:
// Input: n = 9
// Output: 6
// Explanation:
// The integers k in the range [0, 9] whose binary representations are palindromes are:
// 0 → "0"
// 1 → "1"
// 3 → "11"
// 5 → "101"
// 7 → "111"
// 9 → "1001"
// All other values in [0, 9] have non-palindromic binary forms. Therefore, the count is 6.

// Example 2:
// Input: n = 0
// Output: 1
// Explanation:
// Since "0" is a palindrome, the count is 1.

// Constraints:
//     0 <= n <= 10^15

import "fmt"
import "math/bits"

func countBinaryPalindromes(n int64) int {
    if n == 0 { return 1 }
    m := bits.Len(uint(n))
    k := (m - 1) / 2
    // 二进制长度小于 m
    res := 2 << k - 1 
    if m % 2 == 0 {
        res += 1 << k
    }
    // 二进制长度等于 m，且回文数的左半小于 n 的左半
    left := n >> (m / 2)
    res += int(left) - 1<<k
    // 二进制长度等于 m，且回文数的左半等于 n 的左半
    right := bits.Reverse32(uint32(left >> (m % 2))) >> (32 - m / 2)
    if left << (m / 2) | int64(right) <= n {
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 9
    // Output: 6
    // Explanation:
    // The integers k in the range [0, 9] whose binary representations are palindromes are:
    // 0 → "0"
    // 1 → "1"
    // 3 → "11"
    // 5 → "101"
    // 7 → "111"
    // 9 → "1001"
    // All other values in [0, 9] have non-palindromic binary forms. Therefore, the count is 6.
    fmt.Println(countBinaryPalindromes(9)) // 6
    // Example 2:
    // Input: n = 0
    // Output: 1
    // Explanation:
    // Since "0" is a palindrome, the count is 1.
    fmt.Println(countBinaryPalindromes(0)) // 1

    fmt.Println(countBinaryPalindromes(1)) // 2
    fmt.Println(countBinaryPalindromes(1024)) // 63
    fmt.Println(countBinaryPalindromes(10 << 15)) // 1151
}