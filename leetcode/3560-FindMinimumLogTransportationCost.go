package main

// 3560. Find Minimum Log Transportation Cost
// You are given integers n, m, and k.

// There are two logs of lengths n and m units, which need to be transported in three trucks where each truck can carry one log with length at most k units.

// You may cut the logs into smaller pieces, where the cost of cutting a log of length x into logs of length len1 and len2 is cost = len1 * len2 such that len1 + len2 = x.

// Return the minimum total cost to distribute the logs onto the trucks. 
// If the logs don't need to be cut, the total cost is 0.

// Example 1:
// Input: n = 6, m = 5, k = 5
// Output: 5
// Explanation:
// Cut the log with length 6 into logs with length 1 and 5, at a cost equal to 1 * 5 == 5. 
// Now the three logs of length 1, 5, and 5 can fit in one truck each.

// Example 2:
// Input: n = 4, m = 4, k = 6
// Output: 0
// Explanation:
// The two logs can fit in the trucks already, hence we don't need to cut the logs.

// Constraints:
//     2 <= k <= 10^5
//     1 <= n, m <= 2 * k
//     The input is generated such that it is always possible to transport the logs.

import "fmt"

func minCuttingCost(n int, m int, k int) int64 {
    if n <= k && m <= k { return 0 }   
    res := 0 
    if n > k {
        res += k * (n  -k)
    }
    if m > k {
        res += k * (m - k)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 6, m = 5, k = 5
    // Output: 5
    // Explanation:
    // Cut the log with length 6 into logs with length 1 and 5, at a cost equal to 1 * 5 == 5. 
    // Now the three logs of length 1, 5, and 5 can fit in one truck each.
    fmt.Println(minCuttingCost(6, 5, 5)) // 5
    // Example 2:
    // Input: n = 4, m = 4, k = 6
    // Output: 0
    // Explanation:
    // The two logs can fit in the trucks already, hence we don't need to cut the logs.
    fmt.Println(minCuttingCost(4, 4, 6)) // 0

    fmt.Println(minCuttingCost(1, 1, 2)) // 0
    fmt.Println(minCuttingCost(1, 1, 100_000)) // 0
    fmt.Println(minCuttingCost(200_000, 200_000, 100_000)) // 20000000000
    fmt.Println(minCuttingCost(200_000, 1, 100_000)) // 10000000000
    fmt.Println(minCuttingCost(1, 200_000, 100_000)) // 10000000000
    fmt.Println(minCuttingCost(2, 2, 1)) // 2
}