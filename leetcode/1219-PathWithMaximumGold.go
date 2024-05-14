package main

// 1219. Path with Maximum Gold
// In a gold mine grid of size m x n, 
// each cell in this mine has an integer representing the amount of gold in that cell, 
// 0 if it is empty.
// Return the maximum amount of gold you can collect under the conditions:
//     Every time you are located in a cell you will collect all the gold in that cell.
//     From your position, you can walk one step to the left, right, up, or down.
//     You can't visit the same cell more than once.
//     Never visit a cell with 0 gold.
//     You can start and stop collecting gold from any position in the grid that has some gold.
    
// Example 1:
// Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
// Output: 24
// Explanation:
// [[0,6,0],
//  [5,8,7],
//  [0,9,0]]
// Path to get the maximum gold, 9 -> 8 -> 7.

// Example 2:
// Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
// Output: 28
// Explanation:
// [[1,0,7],
//  [2,0,6],
//  [3,4,5],
//  [0,3,0],
//  [9,0,20]]
// Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 15
//     0 <= grid[i][j] <= 100
//     There are at most 25 cells containing gold.

import "fmt"

func getMaximumGold(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(row int, col int) int
    dfs = func(row int, col int) int { // 计算出向4个方向走得到黄金最大的
        gold := grid[row][col]
        grid[row][col] = 0 // marked as visited
        currentMax := 0
        if row > 0 && grid[row-1][col] != 0 { // move up
            currentMax = max(currentMax, dfs(row-1, col))
        }
        if row < rows-1 && grid[row+1][col] != 0 { // move down
            currentMax = max(currentMax, dfs(row+1, col))
        }
        if col > 0 && grid[row][col-1] != 0 { // move left
            currentMax = max(currentMax, dfs(row, col-1))
        }
        if col < cols-1 && grid[row][col+1] != 0 { // move right
            currentMax = max(currentMax, dfs(row, col+1))
        }
        result := gold + currentMax
        grid[row][col] = gold
        return result
    }
    res := 0
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if grid[row][col] != 0 {
                res = max(res, dfs(row, col))
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
    // Output: 24
    // Explanation:
    // [[0,6,0],
    //  [5,8,7],
    //  [0,9,0]]
    // Path to get the maximum gold, 9 -> 8 -> 7.
    fmt.Println(getMaximumGold([][]int{{0,6,0},{5,8,7},{0,9,0}})) // 24
    // Example 2:
    // Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
    // Output: 28
    // Explanation:
    // [[1,0,7],
    //  [2,0,6],
    //  [3,4,5],
    //  [0,3,0],
    //  [9,0,20]]
    // Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
    fmt.Println(getMaximumGold([][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}})) // 28
}