package main

// 2290. Minimum Obstacle Removal to Reach Corner
// You are given a 0-indexed 2D integer array grid of size m x n. 
// Each cell has one of two values:
//     0 represents an empty cell,
//     1 represents an obstacle that may be removed.

// You can move up, down, left, or right from and to an empty cell.

// Return the minimum number of obstacles to remove so you can move from the upper left corner (0, 0) to the lower right corner (m - 1, n - 1).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/04/06/example1drawio-1.png" />
// Input: grid = [[0,1,1],[1,1,0],[1,1,0]]
// Output: 2
// Explanation: We can remove the obstacles at (0, 1) and (0, 2) to create a path from (0, 0) to (2, 2).
// It can be shown that we need to remove at least 2 obstacles, so we return 2.
// Note that there may be other ways to remove 2 obstacles to create a path.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/04/06/example1drawio.png" />
// Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]]
// Output: 0
// Explanation: We can move from (0, 0) to (2, 4) without removing any obstacles, so we return 0.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10^5
//     2 <= m * n <= 10^5
//     grid[i][j] is either 0 or 1.
//     grid[0][0] == grid[m - 1][n - 1] == 0

import "fmt"

func minimumObstacles(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    visited[0][0] = true
    curr, next := [][2]int{{0, 0}}, [][2]int{}
    dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    check := func(i, j int) bool { return i >= 0 && j >= 0 && i < m && j < n }
    for steps := 0; len(curr) > 0; steps++ {
        next = next[:0]
        // Note! A range loop here would not work!
        for k := 0; k < len(curr); k++ {
            i, j := curr[k][0], curr[k][1]
            for _, d := range dirs {
                ii, jj := i+d[0], j+d[1]
                if !check(ii, jj) || visited[ii][jj] { continue }
                if ii == m - 1 && jj == n - 1 { return steps }
                visited[ii][jj] = true
                if grid[ii][jj] == 1 {
                    next = append(next, [2]int{ii, jj})
                } else {
                    curr = append(curr, [2]int{ii, jj})
                }
            }
        }
        curr, next = next, curr
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/04/06/example1drawio-1.png" />
    // Input: grid = [[0,1,1],[1,1,0],[1,1,0]]
    // Output: 2
    // Explanation: We can remove the obstacles at (0, 1) and (0, 2) to create a path from (0, 0) to (2, 2).
    // It can be shown that we need to remove at least 2 obstacles, so we return 2.
    // Note that there may be other ways to remove 2 obstacles to create a path.
    fmt.Println(minimumObstacles([][]int{{0,1,1},{1,1,0},{1,1,0}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/04/06/example1drawio.png" />
    // Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]]
    // Output: 0
    // Explanation: We can move from (0, 0) to (2, 4) without removing any obstacles, so we return 0.
    fmt.Println(minimumObstacles([][]int{{0,1,0,0,0},{0,1,0,1,0},{0,0,0,1,0}})) // 0
}