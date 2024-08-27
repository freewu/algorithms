package main

// 2556. Disconnect Path in a Binary Matrix by at Most One Flip
// You are given a 0-indexed m x n binary matrix grid. 
// You can move from a cell (row, col) to any of the cells (row + 1, col) or (row, col + 1) that has the value 1. 
// The matrix is disconnected if there is no path from (0, 0) to (m - 1, n - 1).

// You can flip the value of at most one (possibly none) cell. 
// You cannot flip the cells (0, 0) and (m - 1, n - 1).

// Return true if it is possible to make the matrix disconnect or false otherwise.
// Note that flipping a cell changes its value from 0 to 1 or from 1 to 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/07/yetgrid2drawio.png" />
// Input: grid = [[1,1,1],[1,0,0],[1,1,1]]
// Output: true
// Explanation: We can change the cell shown in the diagram above. There is no path from (0, 0) to (2, 2) in the resulting grid.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/07/yetgrid3drawio.png" />
// Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
// Output: false
// Explanation: It is not possible to change at most one cell such that there is not path from (0, 0) to (2, 2).

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     grid[i][j] is either 0 or 1.
//     grid[0][0] == grid[m - 1][n - 1] == 1

import "fmt"

func isPossibleToCutPath(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    cache := make(map[int]bool)  // key: 10000*i+j

    var dfs func(m, n, i, j int) bool
    dfs = func (m, n, i, j int) bool {
        if i == m - 1 && j == n - 1 { // arrrive the destination
            return true 
        }
        if _, ok := cache[10000 * i + j]; ok { 
            return false 
        }
        directions := [][]int{ []int{ i+1, j}, []int{i, j+1} }
        for _, d := range(directions) {
            nexti, nextj := d[0], d[1]
            if nexti < m && nextj < n && grid[nexti][nextj] == 1 {
                if dfs(m, n, nexti, nextj) { 
                    grid[i][j] = 0
                    return true
                }
            }
        }
        cache[10000 * i + j] = false // positon (i, j) can't reach the destination
        return false
    }
    if !dfs(m, n, 0, 0) { 
        return true 
    }
    cache[10000] = true
    if !dfs(m, n, 0, 0) { 
        return true 
    }
    return false
}

func isPossibleToCutPath1(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    var dfs func(r,c int)bool
    dfs = func(r,c int)bool{
        if r == m-1 && c == n-1 { // 到达
            return true
        }
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == 0 { // 边界检测
            return false
        }
        grid[r][c] = 0
        return dfs(r+1,c) || dfs(r,c+1)
    }
    grid[0][0] = 1
    res := !dfs(0,0)
    grid[0][0] = 1
    res = res || !dfs(0,0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/07/yetgrid2drawio.png" />
    // Input: grid = [[1,1,1],[1,0,0],[1,1,1]]
    // Output: true
    // Explanation: We can change the cell shown in the diagram above. There is no path from (0, 0) to (2, 2) in the resulting grid.
    fmt.Println(isPossibleToCutPath([][]int{{1,1,1},{1,0,0},{1,1,1}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/07/yetgrid3drawio.png" />
    // Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
    // Output: false
    // Explanation: It is not possible to change at most one cell such that there is not path from (0, 0) to (2, 2).
    fmt.Println(isPossibleToCutPath([][]int{{1,1,1},{1,0,1},{1,1,1}})) // false

    fmt.Println(isPossibleToCutPath1([][]int{{1,1,1},{1,0,0},{1,1,1}})) // true
    fmt.Println(isPossibleToCutPath1([][]int{{1,1,1},{1,0,1},{1,1,1}})) // false
}