package main

// 3882. Minimum XOR Path in a Grid
// You are given a 2D integer array grid of size m * n.

// You start at the top-left cell (0, 0) and want to reach the bottom-right cell (m - 1, n - 1).

// At each step, you may move either right or down.

// The cost of a path is defined as the bitwise XOR of all the values in the cells along that path, including the start and end cells.

// Return the minimum possible XOR value among all valid paths from (0, 0) to (m - 1, n - 1).

// Example 1:
// Input: grid = [[1,2],[3,4]]
// Output: 6
// Explanation:
// There are two valid paths:
// (0, 0) → (0, 1) → (1, 1) with XOR: 1 XOR 2 XOR 4 = 7
// (0, 0) → (1, 0) → (1, 1) with XOR: 1 XOR 3 XOR 4 = 6
// The minimum XOR value among all valid paths is 6.

// Example 2:
// Input: grid = [[6,7],[5,8]]
// Output: 9
// Explanation:
// There are two valid paths:
// (0, 0) → (0, 1) → (1, 1) with XOR: 6 XOR 7 XOR 8 = 9
// (0, 0) → (1, 0) → (1, 1) with XOR: 6 XOR 5 XOR 8 = 11
// The minimum XOR value among all valid paths is 9.

// Example 3:
// Input: grid = [[2,7,5]]
// Output: 0
// Explanation:
// There is only one valid path:
// (0, 0) → (0, 1) → (0, 2) with XOR: 2 XOR 7 XOR 5 = 0
// The XOR value of this path is 0, which is the minimum possible.

// Constraints:
//     1 <= m == grid.length <= 1000
//     1 <= n == grid[i].length <= 1000
//     m * n <= 1000
//     0 <= grid[i][j] <= 1023​

import "fmt"

func minCost(grid [][]int) int {
    m, n, mx := len(grid), len(grid[0]), 1024
    dp := make([][][]bool, m)
    for i := 0; i < m; i++ {
        dp[i] = make([][]bool, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]bool, mx)
        }
    }
    dp[0][0][grid[0][0]] = true
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 { continue }
            v := grid[i][j]
            if i > 0 {
                for x := 0; x < mx; x++ {
                    if dp[i-1][j][x] {
                        dp[i][j][x^v] = true
                    }
                }
            }
            if j > 0 {
                for x := 0; x < mx; x++ {
                    if dp[i][j-1][x] {
                        dp[i][j][x^v] = true
                    }
                }
            }
        }
    }
    for i := 0; i < mx; i++ {
        if dp[m-1][n-1][i] {
            return i
        }
    }
    return -1
}

func minCost1(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    dp := make([][][1024]bool,m)
    for i := 0; i < m; i++ {
        dp[i] = make([][1024]bool,n)
    }
    dp[0][0][grid[0][0]] = true
    for i, x := 1, grid[0][0]; i < m; i++ {
        x ^= grid[i][0]
        dp[i][0][x] = true
    }
    for i, x := 1, grid[0][0]; i < n; i++ {
        x ^= grid[0][i]
        dp[0][i][x] = true
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            for k := 0; k < 1024; k++ {
                dp[i][j][k^grid[i][j]] = dp[i-1][j][k] || dp[i][j-1][k]
            }
        }
    }
    for res<1024 && !dp[m-1][n-1][res] {
        res++
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,2],[3,4]]
    // Output: 6
    // Explanation:
    // There are two valid paths:
    // (0, 0) → (0, 1) → (1, 1) with XOR: 1 XOR 2 XOR 4 = 7
    // (0, 0) → (1, 0) → (1, 1) with XOR: 1 XOR 3 XOR 4 = 6
    // The minimum XOR value among all valid paths is 6.
    fmt.Println(minCost([][]int{{1,2},{3,4}})) // 6
    // Example 2:
    // Input: grid = [[6,7],[5,8]]
    // Output: 9
    // Explanation:
    // There are two valid paths:
    // (0, 0) → (0, 1) → (1, 1) with XOR: 6 XOR 7 XOR 8 = 9
    // (0, 0) → (1, 0) → (1, 1) with XOR: 6 XOR 5 XOR 8 = 11
    // The minimum XOR value among all valid paths is 9.
    fmt.Println(minCost([][]int{{6,7},{5,8}})) // 9
    // Example 3:
    // Input: grid = [[2,7,5]]
    // Output: 0
    // Explanation:
    // There is only one valid path:
    // (0, 0) → (0, 1) → (0, 2) with XOR: 2 XOR 7 XOR 5 = 0
    // The XOR value of this path is 0, which is the minimum possible.
    fmt.Println(minCost([][]int{{2,7,5}})) // 0

    fmt.Println(minCost([][]int{{1,2,3,4,5,6,7,8,9},{1,2,3,4,5,6,7,8,9}})) // 0
    fmt.Println(minCost([][]int{{1,2,3,4,5,6,7,8,9},{9,8,7,6,5,4,3,2,1}})) // 0
    fmt.Println(minCost([][]int{{9,8,7,6,5,4,3,2,1},{1,2,3,4,5,6,7,8,9}})) // 0
    fmt.Println(minCost([][]int{{9,8,7,6,5,4,3,2,1},{9,8,7,6,5,4,3,2,1}})) // 0

    fmt.Println(minCost1([][]int{{1,2},{3,4}})) // 6
    fmt.Println(minCost1([][]int{{6,7},{5,8}})) // 9
    fmt.Println(minCost1([][]int{{2,7,5}})) // 0
    fmt.Println(minCost1([][]int{{1,2,3,4,5,6,7,8,9},{1,2,3,4,5,6,7,8,9}})) // 0
    fmt.Println(minCost1([][]int{{1,2,3,4,5,6,7,8,9},{9,8,7,6,5,4,3,2,1}})) // 0
    fmt.Println(minCost1([][]int{{9,8,7,6,5,4,3,2,1},{1,2,3,4,5,6,7,8,9}})) // 0
    fmt.Println(minCost1([][]int{{9,8,7,6,5,4,3,2,1},{9,8,7,6,5,4,3,2,1}})) // 0
}