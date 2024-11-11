package main

// 3260. Find the Largest Palindrome Divisible by K
// You are given two positive integers n and k.

// An integer x is called k-palindromic if:
//     x is a palindrome.
//     x is divisible by k.

// Return the largest integer having n digits (as a string) that is k-palindromic.

// Note that the integer must not have leading zeros.

// Example 1:
// Input: n = 3, k = 5
// Output: "595"
// Explanation:
// 595 is the largest k-palindromic integer with 3 digits.

// Example 2:
// Input: n = 1, k = 4
// Output: "8"
// Explanation:
// 4 and 8 are the only k-palindromic integers with 1 digit.

// Example 3:
// Input: n = 5, k = 6
// Output: "89898"

// Constraints:
//     1 <= n <= 10^5
//     1 <= k <= 9

import "fmt"
import "bytes"
import "math/bits"

func largestPalindrome(n int, k int) string {
    five, sixes, eights := byte('5'), "66", "888"
    nine := []byte{'9'}
    pattern6 := []string{"77", "8"}
    pattern7 := []string{"7", "77", "5", "77", "7", "99", "4", "44", "6", "44", "4", "99"}
    res := bytes.Repeat(nine, n) // will be used as is for k = 1, 3, 9
    switch k {
    case 2, 4, 8:
        wrap := min(bits.Len(uint(k)) - 1, n)
        copy(res[:wrap], eights)
        copy(res[n - wrap:], eights)
    case 5:
        res[0], res[n - 1] = five, five
    case 6:
        if n < 3 { return sixes[:n] }
        res[0], res[n - 1] = eights[0], eights[0]
        copy(res[(n - 1) / 2:], pattern6[n % 2])
    case 7:
        copy(res[(n - 1) / 2:], pattern7[(n - 1) % 12])
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: n = 3, k = 5
    // Output: "595"
    // Explanation:
    // 595 is the largest k-palindromic integer with 3 digits.
    fmt.Println(largestPalindrome(3, 5)) // "595"
    // Example 2:
    // Input: n = 1, k = 4
    // Output: "8"
    // Explanation:
    // 4 and 8 are the only k-palindromic integers with 1 digit.
    fmt.Println(largestPalindrome(1, 4)) // "8"
    // Example 3:
    // Input: n = 5, k = 6
    // Output: "89898"
    fmt.Println(largestPalindrome(5, 6)) // "89898"

    fmt.Println(largestPalindrome(1, 1)) // "9"
    //fmt.Println(largestPalindrome(10000, 9)) // "999...9"
    fmt.Println(largestPalindrome(1, 9)) // "9"
    //fmt.Println(largestPalindrome(10000, 1)) // "999...9"
}