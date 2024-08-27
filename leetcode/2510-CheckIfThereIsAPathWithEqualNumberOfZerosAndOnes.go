package main

// 2510. Check if There is a Path With Equal Number of 0's And 1's
// You are given a 0-indexed m x n binary matrix grid. 
// You can move from a cell (row, col) to any of the cells (row + 1, col) or (row, col + 1).

// Return true if there is a path from (0, 0) to (m - 1, n - 1) that visits an equal number of 0's and 1's. 
// Otherwise return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/20/yetgriddrawio-4.png" />
// Input: grid = [[0,1,0,0],[0,1,0,0],[1,0,1,0]]
// Output: true
// Explanation: 
// The path colored in blue in the above diagram is a valid path because we have 3 cells with a value of 1 and 3 with a value of 0. 
// Since there is a valid path, we return true.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/20/yetgrid2drawio-1.png" />
// Input: grid = [[1,1,0],[0,0,1],[1,0,0]]
// Output: false
// Explanation: There is no path in this grid with an equal number of 0's and 1's.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 100
//     grid[i][j] is either 0 or 1.

import "fmt"

func isThereAPath(grid [][]int) bool {
    dirs := [][]int{{1,0}, {0,1}}
    m, n, res := len(grid), len(grid[0]), false
    visited := make([][][]bool, m)
    for i := range grid {
        visited[i] = make([][]bool, n)
        for j := range grid[0] {
            visited[i][j] = make([]bool, 500)
        }
    }
    if (m + n) % 2 == 0 {
        return false
    }
    var dfs func(x, y, num0, num1 int)
    dfs = func(x, y, num0, num1 int) {
        if x >= m || y >= n {
            return
        }
        if grid[x][y] == 0 {
            num0++
        } else {
            num1++
        }
        if x == m - 1 && y == n - 1 {
            if num0 == num1 {
                res = true
            }
            return
        }
        if visited[x][y][250 + num0 - num1] {
            return
        }
        visited[x][y][250 + num0 - num1] = true
        for _, v := range dirs {
            dfs(x + v[0], y + v[1], num0, num1)
        }
    }
    dfs(0,0,0,0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/20/yetgriddrawio-4.png" />
    // Input: grid = [[0,1,0,0],[0,1,0,0],[1,0,1,0]]
    // Output: true
    // Explanation: The path colored in blue in the above diagram is a valid path because we have 3 cells with a value of 1 and 3 with a value of 0. Since there is a valid path, we return true.
    grid1 := [][]int{
        {0,1,0,0},
        {0,1,0,0},
        {1,0,1,0},
    }
    fmt.Println(isThereAPath(grid1)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/20/yetgrid2drawio-1.png" />
    // Input: grid = [[1,1,0],[0,0,1],[1,0,0]]
    // Output: false
    // Explanation: There is no path in this grid with an equal number of 0's and 1's.
    grid2 := [][]int{
        {1,1,0},
        {0,0,1},
        {1,0,0},
    }
    fmt.Println(isThereAPath(grid2)) // false
}