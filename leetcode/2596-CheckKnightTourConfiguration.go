package main

// 2596. Check Knight Tour Configuration
// There is a knight on an n x n chessboard. 
// In a valid configuration, the knight starts at the top-left cell of the board and visits every cell on the board exactly once.

// You are given an n x n integer matrix grid consisting of distinct integers from the range [0, n * n - 1] where grid[row][col] indicates 
// that the cell (row, col) is the grid[row][col]th cell that the knight visited. The moves are 0-indexed.

// Return true if grid represents a valid configuration of the knight's movements or false otherwise.

// Note that a valid knight move consists of moving two squares vertically and one square horizontally, or two squares horizontally and one square vertically. 
// The figure below illustrates all the possible eight moves of a knight from some cell.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/knight.png" /> 
// Input: grid = [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
// Output: true
// Explanation: The above diagram represents the grid. It can be shown that it is a valid configuration.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/28/yetgriddrawio-6.png" /> 
// Input: grid = [[0,3,6],[5,8,1],[2,7,4]]
// Output: false
// Explanation: The above diagram represents the grid. The 8th move of the knight is not valid considering its position after the 7th move.

// Constraints:
//     n == grid.length == grid[i].length
//     3 <= n <= 7
//     0 <= grid[row][col] < n * n
//     All integers in grid are unique.

import "fmt"

func checkValidGrid(grid [][]int) bool {
    if grid[0][0] != 0 {
        return false
    }
    dirs := [][]int{{-1,-2},{-1,2},{1,-2},{1,2},{-2,-1},{-2,1},{2,-1},{2,1}}
    row, col, jump, n := 0, 0, 0, len(grid)
    for jump < n*n - 1 {
        found := false
        for _, v := range dirs {
            r, c := row + v[0], col + v [1]
            if 0 <= r && r < n && 0 <= c && c < n && grid[r][c] == jump+1 { // 边界检测
                row, col = r, c
                jump++
                found = true
                break
            }
        }
        if !found {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/knight.png" /> 
    // Input: grid = [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13,2,7,22]]
    // Output: true
    // Explanation: The above diagram represents the grid. It can be shown that it is a valid configuration.
    grid1 := [][]int{
        {0,11,16,5,20},
        {17,4,19,10,15},
        {12,1,8,21,6},
        {3,18,23,14,9},
        {24,13,2,7,22},
    }
    fmt.Println(checkValidGrid(grid1)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/28/yetgriddrawio-6.png" /> 
    // Input: grid = [[0,3,6],[5,8,1],[2,7,4]]
    // Output: false
    // Explanation: The above diagram represents the grid. The 8th move of the knight is not valid considering its position after the 7th move.
    grid2 := [][]int{
        {0,3,6},
        {5,8,1},
        {2,7,4},
    }
    fmt.Println(checkValidGrid(grid2)) // false

    grid3 := [][]int{
        {8,3,6},
        {5,0,1},
        {2,7,4},
    }
    fmt.Println(checkValidGrid(grid3)) // false
}