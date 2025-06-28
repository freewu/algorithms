package main

// 3596. Minimum Cost Path with Alternating Directions I
// You are given two integers m and n representing the number of rows and columns of a grid, respectively.

// The cost to enter cell (i, j) is defined as (i + 1) * (j + 1).

// You start at cell (0, 0) on move 1.

// At each step, you move to an adjacent cell, following an alternating pattern:
//     1. On odd-numbered moves, you must move either right or down.
//     2. On even-numbered moves, you must move either left or up.

// Return the minimum total cost required to reach (m - 1, n - 1). If it is impossible, return -1.

// Example 1:
// Input: m = 1, n = 1
// Output: 1
// Explanation:
// You start at cell (0, 0).
// The cost to enter (0, 0) is (0 + 1) * (0 + 1) = 1.
// Since you're at the destination, the total cost is 1.

// Example 2:
// Input: m = 2, n = 1
// Output: 3
// Explanation:
// You start at cell (0, 0) with cost (0 + 1) * (0 + 1) = 1.
// Move 1 (odd): You can move down to (1, 0) with cost (1 + 1) * (0 + 1) = 2.
// Thus, the total cost is 1 + 2 = 3.

// Constraints:
//     1 <= m, n <= 10^6

import "fmt"

// func minCost(m int, n int) int {
//     if m == 1 && n == 1  { return 1  } // Immediate Arrival Check
//     if (m + n) % 2 == 0  { return -1 } // Reachability Check
//     return (m * (m + 1) / 2) + (n * (n + 1) / 2) - 1
// }

// 打表法
// def minCost(self, m: int, n: int) -> int:
//         return {(1,1):1,(1,2):3,(2,1):3}.get((m,n),-1)

func minCost(m int, n int) int {
    mp := map[[2]int]int{ 
        [2]int{1,1} : 1,
        [2]int{1,2} : 3,
        [2]int{2,1} : 3,
    }
    res, ok := mp[[2]int{ m, n}] 
    if !ok { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: m = 1, n = 1
    // Output: 1
    // Explanation:
    // You start at cell (0, 0).
    // The cost to enter (0, 0) is (0 + 1) * (0 + 1) = 1.
    // Since you're at the destination, the total cost is 1.
    fmt.Println(minCost(1,1)) // 1
    // Example 2:
    // Input: m = 2, n = 1
    // Output: 3
    // Explanation:
    // You start at cell (0, 0) with cost (0 + 1) * (0 + 1) = 1.
    // Move 1 (odd): You can move down to (1, 0) with cost (1 + 1) * (0 + 1) = 2.
    // Thus, the total cost is 1 + 2 = 3.
    fmt.Println(minCost(2,1)) // 3

    fmt.Println(minCost(1, 4)) // -1

    // fmt.Println(minCost(1,1024)) // 524800
    // fmt.Println(minCost(1024,1024)) // -1
    // fmt.Println(minCost(1024,1)) // 524800
    // fmt.Println(minCost(1,1_000_000)) // 500000500000
    // fmt.Println(minCost(1_000_000,1_000_000)) // -1
    // fmt.Println(minCost(1_000_000,1)) // 500000500000
}
