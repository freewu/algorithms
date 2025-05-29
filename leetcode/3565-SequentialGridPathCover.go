package main

// 3565. Sequential Grid Path Cover
// You are given a 2D array grid of size m x n, and an integer k. 
// There are k cells in grid containing the values from 1 to k exactly once, and the rest of the cells have a value 0.

// You can start at any cell, and move from a cell to its neighbors (up, down, left, or right). 
// You must find a path in grid which:
//     1. Visits each cell in grid exactly once.
//     2. Visits the cells with values from 1 to k in order.

// Return a 2D array result of size (m * n) x 2, where result[i] = [xi, yi] represents the ith cell visited in the path. 
// If there are multiple such paths, you may return any one.

// If no such path exists, return an empty array.

// Example 1:
// Input: grid = [[0,0,0],[0,1,2]], k = 2
// Output: [[0,0],[1,0],[1,1],[1,2],[0,2],[0,1]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/16/ezgifcom-animated-gif-maker1.gif" />

// Example 2:
// Input: grid = [[1,0,4],[3,0,2]], k = 4
// Output: []
// Explanation:
// There is no possible path that satisfies the conditions.

// Constraints:
//     1 <= m == grid.length <= 6
//     1 <= n == grid[i].length <= 6
//     1 <= k <= m * n
//     0 <= grid[i][j] <= k
//     grid contains all integers between 1 and k exactly once.

import "fmt"

func findPath(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    res, visited := [][]int{}, make(map[int]bool)
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i, x, y, pre int) bool
    dfs = func(i, x, y, pre int) bool {
        if grid[x][y] == 0 || pre < grid[x][y] {
            res = append(res, []int{x, y}) // push
            visited[x * n + y] = true
            if i == m * n - 1 {  return true }
            for _, dir := range dirs {
                nx, ny := x + dir[0], y + dir[1]
                if nx >= 0 && nx < m && ny >= 0 && ny < n {
                    if !visited[nx * n + ny] {
                        if dfs(i + 1, nx, ny, max(grid[x][y], pre)) {
                            return true
                        }
                    }
                }
            }
            res = res[:len(res) - 1] // pop
            delete(visited, x * n + y)
        }
        return false
    }
    for i := 0 ; i < m; i++ {
        for j := 0 ; j < n; j++ {
            if dfs(0, i, j, 0) {
                return res
            }
        }
    }
    return [][]int{}
}

func main() {
    // Example 1:
    // Input: grid = [[0,0,0],[0,1,2]], k = 2
    // Output: [[0,0],[1,0],[1,1],[1,2],[0,2],[0,1]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/16/ezgifcom-animated-gif-maker1.gif" />
    fmt.Println(findPath([][]int{{0,0,0},{0,1,2}}, 2)) // [[0,0],[1,0],[1,1],[1,2],[0,2],[0,1]]
    // Example 2:
    // Input: grid = [[1,0,4],[3,0,2]], k = 4
    // Output: []
    // Explanation:
    // There is no possible path that satisfies the conditions.
    fmt.Println(findPath([][]int{{1,0,4},{3,0,2}}, 4)) // []
}
