package main

// 3789. Minimum Cost to Acquire Required Items
// You are given five integers cost1, cost2, costBoth, need1, and need2.

// There are three types of items available:
//     1. An item of type 1 costs cost1 and contributes 1 unit to the type 1 requirement only.
//     2. An item of type 2 costs cost2 and contributes 1 unit to the type 2 requirement only.
//     3. An item of type 3 costs costBoth and contributes 1 unit to both type 1 and type 2 requirements.

// You must collect enough items so that the total contribution toward type 1 is at least need1 and the total contribution toward type 2 is at least need2.

// Return an integer representing the minimum possible total cost to achieve these requirements.

// Example 1:
// Input: cost1 = 3, cost2 = 2, costBoth = 1, need1 = 3, need2 = 2
// Output: 3
// Explanation:
// After buying three type 3 items, which cost 3 * 1 = 3, the total contribution to type 1 is 3 (>= need1 = 3) and to type 2 is 3 (>= need2 = 2).
// Any other valid combination would cost more, so the minimum total cost is 3.

// Example 2:
// Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 2, need2 = 3
// Output: 22
// Explanation:
// We buy need1 = 2 items of type 1 and need2 = 3 items of type 2: 2 * 5 + 3 * 4 = 10 + 12 = 22.
// Any other valid combination would cost more, so the minimum total cost is 22.

// Example 3:
// Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 0, need2 = 0
// Output: 0
// Explanation:
// Since no items are required (need1 = need2 = 0), we buy nothing and pay 0.

// Constraints:
//     1 <= cost1, cost2, costBoth <= 10^6
//     0 <= need1, need2 <= 10^9

import "fmt"

func minimumCost(cost1, cost2, costBoth, need1, need2 int) int64 {
    res1 := cost1 * need1 + cost2 * need2 // 各买各的
    if need1 > need2 {
        need1, need2 = need2, need1
        cost2 = cost1
    }
    res2 := costBoth * need2 // 我包了
    res3 := costBoth * need1 + cost2 * (need2 - need1) // 混合策略
    return int64(min(res1, res2, res3))
}

func main() {
    // Example 1:
    // Input: cost1 = 3, cost2 = 2, costBoth = 1, need1 = 3, need2 = 2
    // Output: 3
    // Explanation:
    // After buying three type 3 items, which cost 3 * 1 = 3, the total contribution to type 1 is 3 (>= need1 = 3) and to type 2 is 3 (>= need2 = 2).
    // Any other valid combination would cost more, so the minimum total cost is 3.
    fmt.Println(minimumCost(3, 2, 1, 3, 2)) // 3 
    // Example 2:
    // Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 2, need2 = 3
    // Output: 22
    // Explanation:
    // We buy need1 = 2 items of type 1 and need2 = 3 items of type 2: 2 * 5 + 3 * 4 = 10 + 12 = 22.
    // Any other valid combination would cost more, so the minimum total cost is 22.
    fmt.Println(minimumCost(5, 4, 15, 2, 3)) // 22
    // Example 3:
    // Input: cost1 = 5, cost2 = 4, costBoth = 15, need1 = 0, need2 = 0
    // Output: 0
    // Explanation:
    // Since no items are required (need1 = need2 = 0), we buy nothing and pay 0. 
    fmt.Println(minimumCost(5, 4, 15, 0, 0)) // 0

    fmt.Println(minimumCost(1, 1, 1, 0, 0)) // 0
    fmt.Println(minimumCost(1_000_000, 1_000_000, 1_000_000, 0, 1_000_000_000)) // 1000000000000000
    fmt.Println(minimumCost(1_000_000, 1_000_000, 1_000_000, 1_000_000_000, 0)) // 1000000000000000
    fmt.Println(minimumCost(1, 1_000_000, 1_000_000, 1_000_000_000, 1_000_000_000)) // 1000000000000000
    fmt.Println(minimumCost(1_000_000, 1, 1_000_000, 1_000_000_000, 1_000_000_000)) // 1000000000000000
    fmt.Println(minimumCost(1_000_000, 1_000_000, 1, 1_000_000_000, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1, 1, 1_000_000, 1_000_000_000, 1_000_000_000)) // 2000000000
    fmt.Println(minimumCost(1_000_000, 1, 1, 1_000_000_000, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1, 1_000_000, 1, 1_000_000_000, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1_000_000, 1, 0, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1_000_000, 1_000_000, 0, 0)) // 0
    fmt.Println(minimumCost(1, 1, 1, 1_000_000_000, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1, 1, 0, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1, 1, 1_000_000, 0, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1, 1, 0, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1, 1_000_000, 1, 0, 1_000_000_000)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1, 1, 1_000_000_000, 0)) // 1000000000
    fmt.Println(minimumCost(1, 1, 1_000_000, 1_000_000_000, 0)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1, 1, 1_000_000_000, 0)) // 1000000000
    fmt.Println(minimumCost(1, 1_000_000, 1, 1_000_000_000, 0)) // 1000000000
    fmt.Println(minimumCost(1_000_000, 1_000_000, 1_000_000, 1_000_000_000, 1_000_000_000)) // 1000000000000000
}