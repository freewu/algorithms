package main

// 1034. Coloring A Border
// You are given an m x n integer matrix grid, and three integers row, col, and color. 
// Each value in the grid represents the color of the grid square at that location.

// Two squares are called adjacent if they are next to each other in any of the 4 directions.

// Two squares belong to the same connected component if they have the same color and they are adjacent.

// The border of a connected component is all the squares in the connected component 
// that are either adjacent to (at least) a square not in the component, or on the boundary of the grid (the first or last row or column).

// You should color the border of the connected component that contains the square grid[row][col] with color.

// Return the final grid.

// Example 1:
// Input: grid = [[1,1],[1,2]], row = 0, col = 0, color = 3
// Output: [[3,3],[3,2]]

// Example 2:
// Input: grid = [[1,2,2],[2,3,2]], row = 0, col = 1, color = 3
// Output: [[1,3,3],[2,3,3]]

// Example 3:
// Input: grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2
// Output: [[2,2,2],[2,1,2],[2,2,2]]

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     1 <= grid[i][j], color <= 1000
//     0 <= row < m
//     0 <= col < n

import "fmt"

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
    n, m, before := len(grid), len(grid[0]), grid[row][col]
    var flood func(grid [][]int, r int, c int, before int)
    flood = func(grid [][]int, r int, c int, before int) {
        n, m := len(grid), len(grid[0])
        if r >= 0 && r < n && c >= 0 && c < m && grid[r][c] == before {
            grid[r][c] = 0
            flood(grid, r - 1, c, before)
            flood(grid, r + 1, c, before)
            flood(grid, r, c - 1, before)
            flood(grid, r, c + 1, before)
        }
    }
    flood(grid, row, col, before)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 0 {
                if i == 0 || j == 0 || i == n - 1 || j == m - 1 {
                    grid[i][j] = -1
                } else if grid[i - 1][j] > 0 || grid[i + 1][j] > 0 || grid[i][j - 1] > 0 || grid[i][j + 1] > 0 {
                    grid[i][j] = -1
                }
            }
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 0 {
                grid[i][j] = before
            } else if grid[i][j] == -1 {
                grid[i][j] = color
            }
        }
    }
    return grid
}

func colorBorder1(grid [][]int, row int, col int, color int) [][]int {
    m, n, cur := len(grid), len(grid[0]), grid[row][col]
    arr, visit := []int{}, make([][]bool, m)
    for i := 0; i < m; i++ {
        visit[i] = make([]bool, n)
    }
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != cur || visit[i][j] { // 超出边界或已访问
            return
        }
        visit[i][j] = true
        if i == 0 || j == 0 || i == m-1 || j == n-1 {
            arr = append(arr, i, j)
        } else {
            if cur != grid[i+1][j] ||
                cur != grid[i-1][j] ||
                cur != grid[i][j+1] ||
                cur != grid[i][j-1] {
                arr = append(arr, i, j)
            }
        }
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }
    dfs(row, col)
    for i := 0; i < len(arr); i += 2 {
        grid[arr[i]][arr[i+1]] = color
    }
    return grid
}

func main() {
    // Example 1:
    // Input: grid = [[1,1],[1,2]], row = 0, col = 0, color = 3
    // Output: [[3,3],[3,2]]
    fmt.Println(colorBorder([][]int{{1,1},{1,2}}, 0, 0, 3)) // [[3,3],[3,2]]
    // Example 2:
    // Input: grid = [[1,2,2],[2,3,2]], row = 0, col = 1, color = 3
    // Output: [[1,3,3],[2,3,3]]
    fmt.Println(colorBorder([][]int{{1,2,2},{2,3,2}}, 0, 1, 3)) // [[1,3,3],[2,3,3]]
    // Example 3:
    // Input: grid = [[1,1,1],[1,1,1],[1,1,1]], row = 1, col = 1, color = 2
    // Output: [[2,2,2],[2,1,2],[2,2,2]]
    fmt.Println(colorBorder([][]int{{1,1,1},{1,1,1},{1,1,1}}, 1, 1, 2)) // [[2,2,2],[2,1,2],[2,2,2]]

    fmt.Println(colorBorder1([][]int{{1,1},{1,2}}, 0, 0, 3)) // [[3,3],[3,2]]
    fmt.Println(colorBorder1([][]int{{1,2,2},{2,3,2}}, 0, 1, 3)) // [[1,3,3],[2,3,3]]
    fmt.Println(colorBorder1([][]int{{1,1,1},{1,1,1},{1,1,1}}, 1, 1, 2)) // [[2,2,2],[2,1,2],[2,2,2]]
}