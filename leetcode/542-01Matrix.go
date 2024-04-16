package main

// 542. 01 Matrix
// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/01-1-grid.jpg" />
// Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
// Output: [[0,0,0],[0,1,0],[0,0,0]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/24/01-2-grid.jpg" />
// Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
// Output: [[0,0,0],[0,1,0],[1,2,1]]
 
// Constraints:
//     m == mat.length
//     n == mat[i].length
//     1 <= m, n <= 10^4
//     1 <= m * n <= 10^4
//     mat[i][j] is either 0 or 1.
//     There is at least one 0 in mat.

import "fmt"

// bfs
func updateMatrix(mat [][]int) [][]int {
    if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
        return [][]int{}
    }
    m, n := len(mat), len(mat[0])
    queue := make([][]int, 0)
    MAX_VALUE := m * n
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                queue = append(queue, []int{i, j})
            } else {
                mat[i][j] = MAX_VALUE
            }
        }
    }
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    for len(queue) > 0 {
        cell := queue[0]
        queue = queue[1:]
        for _, dir := range directions {
            r, c := cell[0]+dir[0], cell[1]+dir[1]
            if r >= 0 && r < m && c >= 0 && c < n && mat[r][c] > mat[cell[0]][cell[1]]+1 {
                queue = append(queue, []int{r, c})
                mat[r][c] = mat[cell[0]][cell[1]] + 1
            }
        }
    }
    return mat
}

// dp
func updateMatrix1(mat [][]int) [][]int {
    m, n := len(mat), len(mat[0])
    dp, mx := make([][]int, m), m * n
    dicts := [][]int{ {1, 0}, {-1, 0}, {0, 1}, {0, -1}, }
    for i := range dp {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = mx 
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                dp[i][j] = 0
                continue
            }
            if (i-1 >= 0 && dp[i-1][j] == 0) || (i+1 < m && dp[i+1][j] == 0) || 
                (j-1 >= 0 && dp[i][j-1] == 0) || (j+1 < n && dp[i][j+1] == 0) {
                dp[i][j] = 1
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dp[i][j] > 1 {
                for _, dict := range dicts {
                    x, y := i+dict[0], j+dict[1]
                    if x < 0 || x >= m || y < 0 || y >= n {
                        continue
                    }
                    dp[i][j] = min(dp[i][j], dp[x][y]+1)
                }
            }
        }
    }
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if dp[i][j] > 1 {
                for _, dict := range dicts {
                    x, y := i+dict[0], j+dict[1]
                    if x < 0 || x >= m || y < 0 || y >= n {
                        continue
                    }
                    dp[i][j] = min(dp[i][j], dp[x][y]+1)
                }
            }
        }
    }
    return dp
}

func main() {
    fmt.Println(updateMatrix([][]int{{0,0,0},{0,1,0},{0,0,0}})) // [[0,0,0],[0,1,0],[0,0,0]]
    fmt.Println(updateMatrix([][]int{{0,0,0},{0,1,0},{1,1,1}})) // [[0,0,0],[0,1,0],[1,2,1]]

    fmt.Println(updateMatrix1([][]int{{0,0,0},{0,1,0},{0,0,0}})) // [[0,0,0],[0,1,0],[0,0,0]]
    fmt.Println(updateMatrix1([][]int{{0,0,0},{0,1,0},{1,1,1}})) // [[0,0,0],[0,1,0],[1,2,1]]
}