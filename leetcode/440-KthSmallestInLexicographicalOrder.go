package main

// 440. K-th Smallest in Lexicographical Order
// Given two integers n and k, return the kth lexicographically smallest integer in the range [1, n].

// Example 1:
// Input: n = 13, k = 2
// Output: 10
// Explanation: The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.

// Example 2:
// Input: n = 1, k = 1
// Output: 1

// Constraints:
//     1 <= k <= n <= 10^9

import "fmt"

func findKthNumber(n int, k int) int {
    res := 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    getChild := func (v int,n int) int {
        res, base, levelMin := 1, 10, v * 10
        for levelMin <= n {
            res += min(base, n - levelMin + 1)
            base *= 10
            levelMin *= 10
        }
        return res 
    }
    for k > 0 {
        nums := getChild(res, n)
        if nums < k {
            k -= nums
            res++
            continue
        }
        if k== 1 {
            return res
        }
        k--
        res *= 10
    }
    return 0
}

func main() {
    // Example 1:
    // Input: n = 13, k = 2
    // Output: 10
    // Explanation: The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.
    fmt.Println(findKthNumber(13, 2)) // 10
    // Example 2:
    // Input: n = 1, k = 1
    // Output: 1
    fmt.Println(findKthNumber(1, 1)) // 1
}