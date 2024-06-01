package main

// 2928. Distribute Candies Among Children I
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
//     1 <= n <= 50
//     1 <= limit <= 50

import "fmt"

// 暴力
func distributeCandies(n, limit int) int {
    res := 0
    for i := 0; i <= limit; i++ {
        for j := 0; j <= limit; j++ {
            for k := 0; k <= limit; k++ {
                if i + j + k == n {
                    res++
                }
            }
        }
    }
    return res
}

func distributeCandies1(n int, limit int) int {
    res := 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i <= min(limit, n); i++ {
        if n - i > 2 * limit {
            continue
        }
        res += min(n - i, limit) - max(0, n - i - limit) + 1
    }
    return res
}

func distributeCandies2(n int, limit int) int {
    res := 0
    for i := 0; i <= limit; i++ {
        for j := 0; j <= limit; j++ {
            if i + j > n {
                break
            }
            if n - i - j <= limit {
                res++
            }
        }
    }
    return res
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

    fmt.Println(distributeCandies1(5, 2)) // 3
    fmt.Println(distributeCandies1(3, 3)) // 10

    fmt.Println(distributeCandies2(5, 2)) // 3
    fmt.Println(distributeCandies2(3, 3)) // 10
}