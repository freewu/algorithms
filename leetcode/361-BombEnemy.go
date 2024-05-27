package main

// 361. Bomb Enemy
// Given an m x n matrix grid where each cell is either a wall 'W', an enemy 'E' or empty '0', 
// return the maximum enemies you can kill using one bomb. 
// You can only place the bomb in an empty cell.

// The bomb kills all the enemies in the same row and column from the planted point until it hits the wall since it is too strong to be destroyed.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/27/bomb1-grid.jpg" />
// Input: grid = [["0","E","0","0"],["E","0","W","E"],["0","E","0","0"]]
// Output: 3

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/27/bomb2-grid.jpg" />
// Input: grid = [["W","W","W"],["0","0","0"],["E","E","E"]]
// Output: 1
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 500
//     grid[i][j] is either 'W', 'E', or '0'.

import "fmt"

// 模拟 暴力解法
func maxKilledEnemies(grid [][]byte) int {
    res, m, n := 0, len(grid), len(grid[0])
    if m == 0 {
        return 0
    }
    inArea := func (row, col int) bool { // 边界检测
        if row < 0 || col < 0 || row >= m || col >= n { return false }
        return true
    }
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 四个方向
    for row, v1 := range grid {
        for col, v2 := range v1 {
            if v2 == 'E' || v2 == 'W' { // 非空位置无法放置
                continue
            }
            count := 0
            for _, dir := range dirs { // 遍历所有方向
                for newRow, newCol := row + dir[0], col+dir[1];
                    inArea(newRow, newCol) && grid[newRow][newCol] != 'W';
                    newRow, newCol = newRow+dir[0], newCol + dir[1] { // 一直往该方向延伸直到撞墙
                    if grid[newRow][newCol] == 'E' {
                        count++
                    }
                }
            }
            if count > res {
                res = count
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/27/bomb1-grid.jpg" />
    // Input: grid = [["0","E","0","0"],["E","0","W","E"],["0","E","0","0"]]
    // Output: 3
    fmt.Println(maxKilledEnemies([][]byte{{'0','E','0','0'},{'E','0','W','E'},{'0','E','0','0'}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/27/bomb2-grid.jpg" />
    // Input: grid = [["W","W","W"],["0","0","0"],["E","E","E"]]
    // Output: 1
    fmt.Println(maxKilledEnemies([][]byte{{'W','W','W'},{'0','0','0'},{'E','E','E'}})) // 1
}