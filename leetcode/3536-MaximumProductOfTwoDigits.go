package main

// 3536. Maximum Product of Two Digits
// You are given a positive integer n.
// Return the maximum product of any two digits in n.
// Note: You may use the same digit twice if it appears more than once in n.

// Example 1:
// Input: n = 31
// Output: 3
// Explanation:
// The digits of n are [3, 1].
// The possible products of any two digits are: 3 * 1 = 3.
// The maximum product is 3.

// Example 2:
// Input: n = 22
// Output: 4
// Explanation:
// The digits of n are [2, 2].
// The possible products of any two digits are: 2 * 2 = 4.
// The maximum product is 4.

// Example 3:
// Input: n = 124
// Output: 8
// Explanation:
// The digits of n are [1, 2, 4].
// The possible products of any two digits are: 1 * 2 = 2, 1 * 4 = 4, 2 * 4 = 8.
// The maximum product is 8.

// Constraints:
//     10 <= n <= 10^9

import "fmt"
import "sort"

func maxProduct(n int) int {
    arr := []int{}
    for n > 0 {
        arr = append(arr, n % 10)
        n = n / 10
    }
    sort.Ints(arr)
    return arr[len(arr) - 2] * arr[len(arr) - 1]
}

func main() {
    // Example 1:
    // Input: n = 31
    // Output: 3
    // Explanation:
    // The digits of n are [3, 1].
    // The possible products of any two digits are: 3 * 1 = 3.
    // The maximum product is 3.
    fmt.Println(maxProduct(31)) // 3
    // Example 2:
    // Input: n = 22
    // Output: 4
    // Explanation:
    // The digits of n are [2, 2].
    // The possible products of any two digits are: 2 * 2 = 4.
    // The maximum product is 4.
    fmt.Println(maxProduct(22)) // 4
    // Example 3:
    // Input: n = 124
    // Output: 8
    // Explanation:
    // The digits of n are [1, 2, 4].
    // The possible products of any two digits are: 1 * 2 = 2, 1 * 4 = 4, 2 * 4 = 8.
    // The maximum product is 8.
    fmt.Println(maxProduct(124)) // 8

    fmt.Println(maxProduct(10)) // 0
    fmt.Println(maxProduct(1024)) // 8
    fmt.Println(maxProduct(999_999_999)) // 91
    fmt.Println(maxProduct(1_000_000_000)) // 0
}