package main

// 64. Minimum Path Sum
// Given a m x n grid filled with non-negative numbers, 
// find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

// Note: You can only move either down or right at any point in time.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/minpath.jpg" />
// Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
// Output: 7
// Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

// Example 2:
// Input: grid = [[1,2,3],[4,5,6]]
// Output: 12
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 200
//     0 <= grid[i][j] <= 200

import "fmt"

// 原地 DP，无辅助空间
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    for i := 1; i < m; i++ {
        grid[i][0] += grid[i-1][0]
    }
    for j := 1; j < n; j++ {
        grid[0][j] += grid[0][j-1]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }
    return grid[m-1][n-1]
}

// 最原始的方法，辅助空间 O(n^2)
func minPathSum1(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    if m == 0 || n == 0 {
        return 0
    }
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    // initFirstCol
    for i := 0; i < len(dp); i++ {
        if i == 0 {
            dp[i][0] = grid[i][0]
        } else {
            dp[i][0] = grid[i][0] + dp[i-1][0]
        }
    }
    // initFirstRow
    for i := 0; i < len(dp[0]); i++ {
        if i == 0 {
            dp[0][i] = grid[0][i]
        } else {
            dp[0][i] = grid[0][i] + dp[0][i-1]
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }
    return dp[m-1][n-1]
}

// best solution
func minPathSum2(grid [][]int) int {
    // parallel grid which finds the minimum per square
    dp := make([][]int, len(grid))
    for i,_ := range grid {
        dp[i] = make([]int, len(grid[i]))
    }
    for i := 0; i < len(grid); i += 1{
        for j := 0 ; j < len(grid[i]); j += 1 {
            if i == 0 && j == 0 {
                dp[i][j] = grid[i][j]
            } else if i == 0 {
                dp[i][j] = dp[i][j-1] + grid[i][j]
            } else if j == 0 {
                dp[i][j] = dp[i-1][j] + grid[i][j]
            } else { // i > 0 && j > 0
                m := dp[i-1][j]
                if dp[i][j-1] < m {
                    m = dp[i][j-1]
                }
                dp[i][j] = m + grid[i][j]
            }
        }
    }
    return dp[len(dp) - 1][len(dp[0]) - 1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/05/minpath.jpg" />
    // Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
    // Output: 7
    // Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
    fmt.Printf("minPathSum([][]int{ []int{1,3,1},[]int{1,5,1},[]int{4,2,1}}) = %v\n",minPathSum([][]int{ []int{1,3,1},[]int{1,5,1},[]int{4,2,1}}))
    // Example 2:
    // Input: grid = [[1,2,3],[4,5,6]]
    // Output: 12
    fmt.Printf("minPathSum([][]int{ []int{1,2,3},[]int{4,5,6} }) = %v\n",minPathSum([][]int{  []int{1,2,3},[]int{4,5,6} }))

    fmt.Printf("minPathSum1([][]int{ []int{1,3,1},[]int{1,5,1},[]int{4,2,1}}) = %v\n",minPathSum1([][]int{ []int{1,3,1},[]int{1,5,1},[]int{4,2,1}}))
    fmt.Printf("minPathSum1([][]int{ []int{1,2,3},[]int{4,5,6} }) = %v\n",minPathSum1([][]int{  []int{1,2,3},[]int{4,5,6} }))

    fmt.Printf("minPathSum2([][]int{ []int{1,3,1},[]int{1,5,1},[]int{4,2,1}}) = %v\n",minPathSum2([][]int{ []int{1,3,1},[]int{1,5,1},[]int{4,2,1}}))
    fmt.Printf("minPathSum2([][]int{ []int{1,2,3},[]int{4,5,6} }) = %v\n",minPathSum2([][]int{  []int{1,2,3},[]int{4,5,6} }))
}
