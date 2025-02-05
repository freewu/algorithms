package main

// 2929. Distribute Candies Among Children II
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
//     1 <= n <= 10^6
//     1 <= limit <= 10^6

import "fmt"

func distributeCandies(n int, limit int) int64 {
    res := int64(0)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i <= min(limit, n); i++ {
        if float64(n - i) / 2. > float64(limit) { continue }
        e := min(limit, n - i)
        s := n - i - e
        res += int64(e - s + 1)
    } 
    return res
}

func distributeCandies1(n int, limit int) int64 {
    // this is equivalent to solving the problem of
    // x1 + x2 + x3 = 5, for 0 <= x1, x2, x3 <= limit
    // the constant time solution is with the math formula...
    // first we solve the problem regardless of the limit
    // that is solve x1 + x2 + x3 = 5 for x1,x2,x3 >= 0
    // then we solve x1 + x2 + x3 = 5, for above AND x1 <= limit (WLOG so times 3 for other combinations)
    // then we solve x1 + x2 + x3 = 5, for above AND x2 <= limit (WLOG so times 3 for other combinations)
    // then, the actual problem is just part 1 minus the sum of parts 2 and 3
    calc := func(n int) int64 {
        if n < 2 { return 0 }
        return int64(n) * int64(n-1) / 2
    }
    return calc(n + 2) - 3 * calc(n - limit + 1) + 3 * calc(n - 2 * limit) - calc(n - 3 * limit - 1)
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
    fmt.Println(distributeCandies(1000000, 1000000)) // 500001500001
    fmt.Println(distributeCandies(1, 1000000)) // 3
    fmt.Println(distributeCandies(1000000, 1)) // 0

    fmt.Println(distributeCandies1(5, 2)) // 3
    fmt.Println(distributeCandies1(3, 3)) // 10
    fmt.Println(distributeCandies1(1, 1)) // 3
    fmt.Println(distributeCandies1(1000000, 1000000)) // 500001500001
    fmt.Println(distributeCandies1(1, 1000000)) // 3
    fmt.Println(distributeCandies1(1000000, 1)) // 0
}