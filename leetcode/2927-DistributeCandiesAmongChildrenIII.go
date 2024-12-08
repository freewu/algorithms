package main

// 2927. Distribute Candies Among Children III
// You are given two positive integers n and limit.

// Return the total number of ways to distribute n candies among 3 children such that no child gets more than limit candies.

// Example 1:
// Input: n = 5, limit = 2
// Output: 3
// Explanation: There are 3 ways to distribute 5 candies such that no child gets more than 2 candies: (1, 2, 2), (2, 1, 2) and (2, 2, 1).

// Example 2:
// Input: n = 3, limit = 3
// Output: 10
// Explanation: There are 10 ways to distribute 3 candies such that no child gets more than 3 candies: (0, 0, 3), (0, 1, 2), (0, 2, 1), (0, 3, 0), (1, 0, 2), (1, 1, 1), (1, 2, 0), (2, 0, 1), (2, 1, 0) and (3, 0, 0).

// Constraints:
//     1 <= n <= 10^8
//     1 <= limit <= 10^8

import "fmt"

func distributeCandies(n int, limit int) int64 {
    comb2 := func(n int) int { return n * (n - 1) / 2 }
    if n > 3 * limit { return 0 }
    res := comb2(n + 2)
    if n > limit {
        res -= 3 * comb2(n - limit + 1)
    }
    if n - 2 >= 2 * limit {
        res += 3 * comb2(n - 2 * limit)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 5, limit = 2
    // Output: 3
    // Explanation: There are 3 ways to distribute 5 candies such that no child gets more than 2 candies: (1, 2, 2), (2, 1, 2) and (2, 2, 1).
    fmt.Println(distributeCandies(5, 2)) // 3
    // Example 2:
    // Input: n = 3, limit = 3
    // Output: 10
    // Explanation: There are 10 ways to distribute 3 candies such that no child gets more than 3 candies: (0, 0, 3), (0, 1, 2), (0, 2, 1), (0, 3, 0), (1, 0, 2), (1, 1, 1), (1, 2, 0), (2, 0, 1), (2, 1, 0) and (3, 0, 0).
    fmt.Println(distributeCandies(3, 3)) // 10

    fmt.Println(distributeCandies(1, 1)) // 3
    fmt.Println(distributeCandies(1, 100_000_000)) // 3
    fmt.Println(distributeCandies(100_000_000, 1)) // 0
    fmt.Println(distributeCandies(100_000_000, 100_000_000)) // 5000000150000001
}