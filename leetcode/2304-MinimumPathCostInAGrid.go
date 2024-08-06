package main

// 2304. Minimum Path Cost in a Grid
// You are given a 0-indexed m x n integer matrix grid consisting of distinct integers from 0 to m * n - 1. 
// You can move in this matrix from a cell to any other cell in the next row. 
// That is, if you are in cell (x, y) such that x < m - 1, you can move to any of the cells (x + 1, 0), (x + 1, 1), ..., (x + 1, n - 1). Note that it is not possible to move from cells in the last row.

// Each possible move has a cost given by a 0-indexed 2D array moveCost of size (m * n) x n, 
// where moveCost[i][j] is the cost of moving from a cell with value i to a cell in column j of the next row. 
// The cost of moving from cells in the last row of grid can be ignored.

// The cost of a path in grid is the sum of all values of cells visited plus the sum of costs of all the moves made. 
// Return the minimum cost of a path that starts from any cell in the first row and ends at any cell in the last row.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/04/28/griddrawio-2.png" />
// Input: grid = [[5,3],[4,0],[2,1]], moveCost = [[9,8],[1,5],[10,12],[18,6],[2,4],[14,3]]
// Output: 17
// Explanation: The path with the minimum possible cost is the path 5 -> 0 -> 1.
// - The sum of the values of cells visited is 5 + 0 + 1 = 6.
// - The cost of moving from 5 to 0 is 3.
// - The cost of moving from 0 to 1 is 8.
// So the total cost of the path is 6 + 3 + 8 = 17.

// Example 2:
// Input: grid = [[5,1,2],[4,0,3]], moveCost = [[12,10,15],[20,23,8],[21,7,1],[8,1,13],[9,10,25],[5,3,2]]
// Output: 6
// Explanation: The path with the minimum possible cost is the path 2 -> 3.
// - The sum of the values of cells visited is 2 + 3 = 5.
// - The cost of moving from 2 to 3 is 1.
// So the total cost of this path is 5 + 1 = 6.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 50
//     grid consists of distinct integers from 0 to m * n - 1.
//     moveCost.length == m * n
//     moveCost[i].length == n
//     1 <= moveCost[i][j] <= 100

import "fmt"

// Top-down with memoization:
func minPathCost(grid [][]int, moveCost [][]int) int {
    n, inf := len(grid[0]), 1 << 31
    // Initialize dp matrix
    minCost := make([][]int, 0)
    for range grid {
        minCost = append(minCost, make([]int, n))
    }
    var dp func(i, j int) int
    dp = func(i, j int) int {
        if i == 0 {
            return grid[0][j]
        }
        if minCost[i][j] != 0 {
            return minCost[i][j]
        }
        localMin := inf
        for k, prev := range grid[i-1] {
            temp := dp(i-1, k) + grid[i][j] + moveCost[prev][j]
            if temp < localMin {
                localMin = temp
            }
        }
        minCost[i][j] = localMin
        return localMin
    }
    globalMin := inf
    for col, _ := range grid[len(grid) - 1] {
        currCost := dp(len(grid) - 1, col)
        if currCost < globalMin {
            globalMin = currCost
        }
    }
    return globalMin
}

// "Bellman-Ford" bottom-up:
func minPathCost1(grid [][]int, moveCost [][]int) int {
    // Initialize matrix (first row same as grid)
    minCost, inf := make([][]int, 0), 1 << 31
    minCost = append(minCost, grid[0])
    for range grid[1:] {
        minCost = append(minCost, make([]int, len(grid[0])))
    }
    for i:=1; i<len(grid); i++ {
        for j:=0; j<len(grid[0]); j++ {
            localMin := inf
            for prevIndex, prev := range grid[i-1] {
                temp := grid[i][j] + minCost[i-1][prevIndex] + moveCost[prev][j]
                if temp < localMin {
                    localMin = temp
                }
            }
            minCost[i][j] = localMin
        }
    }
    globalMin := inf
    for _, val := range minCost[len(grid)-1] {
        if val < globalMin {
            globalMin = val
        }
    }
    return globalMin
}

func minPathCost2(grid [][]int, moveCost [][]int) int {
    pre, dp, m, n, inf := make([]int, len(grid[0])), make([]int, len(grid[0])), len(grid), len(grid[0]), 1 << 31
    for i := range dp {
        dp[i] = grid[m-1][i]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := m - 2; i >=0; i-- {
        copy(pre, dp)
        for j := 0; j < n; j++ {
            res := inf
            for k := 0; k < n; k++ {
                res = min(inf, grid[i][j] + moveCost[grid[i][j]][k] + pre[k])
            } 
            dp[j] = res
        }
    }
    res := dp[0]
    for i := range dp {
        res = min(res, dp[i])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/04/28/griddrawio-2.png" />
    // Input: grid = [[5,3],[4,0],[2,1]], moveCost = [[9,8],[1,5],[10,12],[18,6],[2,4],[14,3]]
    // Output: 17
    // Explanation: The path with the minimum possible cost is the path 5 -> 0 -> 1.
    // - The sum of the values of cells visited is 5 + 0 + 1 = 6.
    // - The cost of moving from 5 to 0 is 3.
    // - The cost of moving from 0 to 1 is 8.
    // So the total cost of the path is 6 + 3 + 8 = 17.
    fmt.Println(minPathCost([][]int{{5,3},{4,0},{2,1}},[][]int{{9,8},{1,5},{10,12},{18,6},{2,4},{14,3}})) // 17
    // Example 2:
    // Input: grid = [[5,1,2],[4,0,3]], moveCost = [[12,10,15],[20,23,8],[21,7,1],[8,1,13],[9,10,25],[5,3,2]]
    // Output: 6
    // Explanation: The path with the minimum possible cost is the path 2 -> 3.
    // - The sum of the values of cells visited is 2 + 3 = 5.
    // - The cost of moving from 2 to 3 is 1.
    // So the total cost of this path is 5 + 1 = 6.
    fmt.Println(minPathCost([][]int{{5,1,2},{4,0,3}},[][]int{{12,10,15},{20,23,8},{21,7,1},{8,1,13},{9,10,25},{5,3,2}})) // 6

    fmt.Println(minPathCost1([][]int{{5,3},{4,0},{2,1}},[][]int{{9,8},{1,5},{10,12},{18,6},{2,4},{14,3}})) // 17
    fmt.Println(minPathCost1([][]int{{5,1,2},{4,0,3}},[][]int{{12,10,15},{20,23,8},{21,7,1},{8,1,13},{9,10,25},{5,3,2}})) // 6

    fmt.Println(minPathCost2([][]int{{5,3},{4,0},{2,1}},[][]int{{9,8},{1,5},{10,12},{18,6},{2,4},{14,3}})) // 17
    fmt.Println(minPathCost2([][]int{{5,1,2},{4,0,3}},[][]int{{12,10,15},{20,23,8},{21,7,1},{8,1,13},{9,10,25},{5,3,2}})) // 6
}