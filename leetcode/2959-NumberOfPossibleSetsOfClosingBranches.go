package main

// 2959. Number of Possible Sets of Closing Branches
// There is a company with n branches across the country, some of which are connected by roads. 
// Initially, all branches are reachable from each other by traveling some roads.

// The company has realized that they are spending an excessive amount of time traveling between their branches. 
// As a result, they have decided to close down some of these branches (possibly none). 
// However, they want to ensure that the remaining branches have a distance of at most maxDistance from each other.

// The distance between two branches is the minimum total traveled length needed to reach one branch from another.

// You are given integers n, maxDistance, and a 0-indexed 2D array roads, 
// where roads[i] = [ui, vi, wi] represents the undirected road between branches ui and vi with length wi.

// Return the number of possible sets of closing branches, so that any branch has a distance of at most maxDistance from any other.
// Note that, after closing a branch, the company will no longer have access to any roads connected to it.
// Note that, multiple roads are allowed.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/11/08/example11.png" />
// Input: n = 3, maxDistance = 5, roads = [[0,1,2],[1,2,10],[0,2,10]]
// Output: 5
// Explanation: The possible sets of closing branches are:
// - The set [2], after closing, active branches are [0,1] and they are reachable to each other within distance 2.
// - The set [0,1], after closing, the active branch is [2].
// - The set [1,2], after closing, the active branch is [0].
// - The set [0,2], after closing, the active branch is [1].
// - The set [0,1,2], after closing, there are no active branches.
// It can be proven, that there are only 5 possible sets of closing branches.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/11/08/example22.png" />
// Input: n = 3, maxDistance = 5, roads = [[0,1,20],[0,1,10],[1,2,2],[0,2,2]]
// Output: 7
// Explanation: The possible sets of closing branches are:
// - The set [], after closing, active branches are [0,1,2] and they are reachable to each other within distance 4.
// - The set [0], after closing, active branches are [1,2] and they are reachable to each other within distance 2.
// - The set [1], after closing, active branches are [0,2] and they are reachable to each other within distance 2.
// - The set [0,1], after closing, the active branch is [2].
// - The set [1,2], after closing, the active branch is [0].
// - The set [0,2], after closing, the active branch is [1].
// - The set [0,1,2], after closing, there are no active branches.
// It can be proven, that there are only 7 possible sets of closing branches.

// Example 3:
// Input: n = 1, maxDistance = 10, roads = []
// Output: 2
// Explanation: The possible sets of closing branches are:
// - The set [], after closing, the active branch is [0].
// - The set [0], after closing, there are no active branches.
// It can be proven, that there are only 2 possible sets of closing branches.

// Constraints:
//     1 <= n <= 10
//     1 <= maxDistance <= 10^5
//     0 <= roads.length <= 1000
//     roads[i].length == 3
//     0 <= ui, vi <= n - 1
//     ui != vi
//     1 <= wi <= 1000
//     All branches are reachable from each other by traveling some roads.

import "fmt"

// Floyd最短路径算法
func numberOfSets(n int, maxDistance int, roads [][]int) int {
    res, opened, dp := 0, make([]int, n), make([][]int, n)
    for mask := 0; mask < (1 << n); mask++ {
        for i := 0; i < n; i++ {
            opened[i] = mask & (1 << i)
        }
        for i := range dp {
            dp[i] = make([]int, n)
            for j := range dp[i] {
                dp[i][j] = 1000000
            }
        }
        for _, road := range roads {
            i, j, r := road[0], road[1], road[2]
            if opened[i] > 0 && opened[j] > 0 {
                if r < dp[i][j] {
                    dp[i][j], dp[j][i] = r, r
                }
            }
        }
        // Floyd-Warshall algorithm
        for k := 0; k < n; k++ {
            if opened[k] > 0 {
                for i := 0; i < n; i++ {
                    if opened[i] > 0 {
                        for j := i + 1; j < n; j++ {
                            if opened[j] > 0 {
                                if dp[i][k] + dp[k][j] < dp[i][j] {
                                    dp[i][j] = dp[i][k] + dp[k][j]
                                    dp[j][i] = dp[i][j] // Ensure symmetry
                                }
                            }
                        }
                    }
                }
            }
        }
        // Validate
        good := 1
        for i := 0; i < n; i++ {
            if opened[i] > 0 {
                for j := i + 1; j < n; j++ {
                    if opened[j] > 0 && dp[i][j] > maxDistance {
                        good = 0
                        break
                    }
                }
                if good == 0 {
                    break
                }
            }
        }
        res += good
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/11/08/example11.png" />
    // Input: n = 3, maxDistance = 5, roads = [[0,1,2],[1,2,10],[0,2,10]]
    // Output: 5
    // Explanation: The possible sets of closing branches are:
    // - The set [2], after closing, active branches are [0,1] and they are reachable to each other within distance 2.
    // - The set [0,1], after closing, the active branch is [2].
    // - The set [1,2], after closing, the active branch is [0].
    // - The set [0,2], after closing, the active branch is [1].
    // - The set [0,1,2], after closing, there are no active branches.
    // It can be proven, that there are only 5 possible sets of closing branches.
    fmt.Println(numberOfSets(3,5,[][]int{{0,1,2},{1,2,10},{0,2,10}})) // 5
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/11/08/example22.png" />
    // Input: n = 3, maxDistance = 5, roads = [[0,1,20],[0,1,10],[1,2,2],[0,2,2]]
    // Output: 7
    // Explanation: The possible sets of closing branches are:
    // - The set [], after closing, active branches are [0,1,2] and they are reachable to each other within distance 4.
    // - The set [0], after closing, active branches are [1,2] and they are reachable to each other within distance 2.
    // - The set [1], after closing, active branches are [0,2] and they are reachable to each other within distance 2.
    // - The set [0,1], after closing, the active branch is [2].
    // - The set [1,2], after closing, the active branch is [0].
    // - The set [0,2], after closing, the active branch is [1].
    // - The set [0,1,2], after closing, there are no active branches.
    // It can be proven, that there are only 7 possible sets of closing branches.
    fmt.Println(numberOfSets(3,5,[][]int{{0,1,20},{0,1,10},{1,2,2},{0,2,2}})) // 7
    // Example 3:
    // Input: n = 1, maxDistance = 10, roads = []
    // Output: 2
    // Explanation: The possible sets of closing branches are:
    // - The set [], after closing, the active branch is [0].
    // - The set [0], after closing, there are no active branches.
    // It can be proven, that there are only 2 possible sets of closing branches.
    fmt.Println(numberOfSets(1,10,[][]int{})) // 2
}