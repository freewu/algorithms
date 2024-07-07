package main

// 1020. Number of Enclaves
// You are given an m x n binary matrix grid, where 0 represents a sea cell and 1 represents a land cell.
// A move consists of walking from one land cell to another adjacent (4-directionally) land cell or walking off the boundary of the grid.
// Return the number of land cells in grid for which we cannot walk off the boundary of the grid in any number of moves. 

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/18/enclaves1.jpg" />
// Input: grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
// Output: 3
// Explanation: There are three 1s that are enclosed by 0s, and one 1 that is not enclosed because its on the boundary.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/18/enclaves2.jpg" />
// Input: grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
// Output: 0
// Explanation: All 1s are either on the boundary or can reach the boundary.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 500
//     grid[i][j] is either 0 or 1.

import "fmt"

func numEnclaves(grid [][]int) int {
    var dfs func(grid [][]int, row,col int)
    dfs = func(grid [][]int, row,col int) {
        if row < 0 || row >= len(grid) || col <0 || col >= len(grid[row]) { // 边界检测
            return
        }
        if grid[row][col] == 0 { // 水
            return
        }
        grid[row][col] = 0
        dfs(grid,row + 1,col) // 向右
        dfs(grid,row - 1,col) // 向左
        dfs(grid,row,col + 1) // 向下
        dfs(grid,row,col - 1) // 向上
    }
    for i := 0 ; i < len(grid); i++ {
        dfs(grid,i,0)
        dfs(grid,i,len(grid[0]) - 1)
    }
    for i := 0; i < len(grid[0]); i++{
        dfs(grid,0,i)
        dfs(grid,len(grid)-1,i)
    }
    res := 0
    for i := 0 ; i <len(grid); i++{
        for j := 0 ; j < len(grid[i]) ; j++{
            if grid[i][j] == 1 { // 陆地
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/02/18/enclaves1.jpg" />
    // Input: grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
    // Output: 3
    // Explanation: There are three 1s that are enclosed by 0s, and one 1 that is not enclosed because its on the boundary.
    grid1 := [][]int{
        {0,0,0,0},
        {1,0,1,0},
        {0,1,1,0},
        {0,0,0,0},
    }
    fmt.Println(numEnclaves(grid1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/02/18/enclaves2.jpg" />
    // Input: grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
    // Output: 0
    // Explanation: All 1s are either on the boundary or can reach the boundary.
    grid2 := [][]int{
        {0,1,1,0},
        {0,0,1,0},
        {0,0,1,0},
        {0,0,0,0},
    }
    fmt.Println(numEnclaves(grid2)) // 3
}