package main

// 3183. The Number of Ways to Make the Sum
// You have an infinite number of coins with values 1, 2, and 6, and only 2 coins with value 4.

// Given an integer n, return the number of ways to make the sum of n with the coins you have.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Note that the order of the coins doesn't matter and [2, 2, 3] is the same as [2, 3, 2].

// Example 1:
// Input: n = 4
// Output: 4
// Explanation:
// Here are the four combinations: [1, 1, 1, 1], [1, 1, 2], [2, 2], [4].

// Example 2:
// Input: n = 12
// Output: 22
// Explanation:
// Note that [4, 4, 4] is not a valid combination since we cannot use 4 three times.

// Example 3:
// Input: n = 5
// Output: 4
// Explanation:
// Here are the four combinations: [1, 1, 1, 1, 1], [1, 1, 1, 2], [1, 2, 2], [1, 4].

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func numberOfWays(n int) int {
    mod := 1_000_000_007
    coins := []int{1, 2, 6}
    facts := make([]int, n + 1)
    facts[0] = 1
    for _, v := range coins {
        for i := v; i <= n; i++ {
            facts[i] = (facts[i] + facts[i - v]) % mod
        }
    }
    res := facts[n]
    if n >= 4 { res = (res + facts[n - 4]) % mod }
    if n >= 8 { res = (res + facts[n - 8]) % mod }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4
    // Output: 4
    // Explanation:
    // Here are the four combinations: [1, 1, 1, 1], [1, 1, 2], [2, 2], [4].
    fmt.Println(numberOfWays(4)) // 4
    // Example 2:
    // Input: n = 12
    // Output: 22
    // Explanation:
    // Note that [4, 4, 4] is not a valid combination since we cannot use 4 three times.
    fmt.Println(numberOfWays(12)) // 22
    // Example 3:
    // Input: n = 5
    // Output: 4
    // Explanation:
    // Here are the four combinations: [1, 1, 1, 1, 1], [1, 1, 1, 2], [1, 2, 2], [1, 4].
    fmt.Println(numberOfWays(5)) // 4

    fmt.Println(numberOfWays(1)) // 1
    fmt.Println(numberOfWays(2)) // 2
    fmt.Println(numberOfWays(64)) // 529
    fmt.Println(numberOfWays(1024)) // 131329
    fmt.Println(numberOfWays(99_999)) // 249974994
    fmt.Println(numberOfWays(100_000)) // 250024994
}