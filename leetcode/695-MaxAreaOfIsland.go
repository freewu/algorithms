package main

// 695. Max Area of Island
// You are given an m x n binary matrix grid. 
// An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) 
// You may assume all four edges of the grid are surrounded by water.

// The area of an island is the number of cells with a value 1 in the island.
// Return the maximum area of an island in grid. If there is no island, return 0.

// Example 1:
// <img scr="" />
// Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
// Output: 6
// Explanation: The answer is not 11, because the island must be connected 4-directionally.

// Example 2:
// Input: grid = [[0,0,0,0,0,0,0,0]]
// Output: 0
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     grid[i][j] is either 0 or 1.

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
    dir := [][]int{ {-1, 0}, {0, 1}, {1, 0}, {0, -1}, }
    isInGrid := func (grid [][]int, x, y int) bool {
        return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
    }
    var areaOfIsland func(grid [][]int, x, y int) int
    areaOfIsland = func(grid [][]int, x, y int) int {
        if !isInGrid(grid, x, y) || grid[x][y] == 0 { // 靠边缘的岛屿不能计算在内
            return 0
        }
        grid[x][y] = 0
        total := 1
        for i := 0; i < 4; i++ { // 向上下左右四个方向
            nx := x + dir[i][0]
            ny := y + dir[i][1]
            total += areaOfIsland(grid, nx, ny)
        }
        return total
    }
    res := 0
    for i, row := range grid {
        for j, col := range row {
            if col == 0 {
                continue
            }
            area := areaOfIsland(grid, i, j)
            if area > res { // 动态维护岛屿的最大面积
                res = area
            }
        }
    }
    return res
}

func maxAreaOfIsland1(grid [][]int) int {
    dir := [][]int{ {-1, 0}, {1, 0}, {0, -1}, {0, 1} }
    res, count, visited := 0, 0, make([][]bool, len(grid))
    for i := range visited {
        visited[i] = make([]bool, len(grid[i]))
    }
    var dfs func(grid [][]int, i, j int, visited [][]bool)
    dfs = func(grid [][]int, i, j int, visited [][]bool) {
        for _, d := range dir {
            x, y := i+d[0], j+d[1]
            if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) { // 靠边缘的岛屿不能计算在内
                continue
            }
            if grid[x][y] == 1 && !visited[x][y] {
                count++
                visited[x][y] = true
                dfs(grid, x, y, visited)
            }
        }
    }
    for i := range grid {
        for j := range grid[i] {
            if !visited[i][j] && grid[i][j] == 1 {
                count = 1
                visited[i][j] = true
                dfs(grid, i, j, visited)
                if count > res {
                    res = count
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img scr="" />
    // Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
    // Output: 6
    // Explanation: The answer is not 11, because the island must be connected 4-directionally.
    island1 := [][]int{
        {0,0,1,0,0,0,0,1,0,0,0,0,0},
        {0,0,0,0,0,0,0,1,1,1,0,0,0},
        {0,1,1,0,1,0,0,0,0,0,0,0,0},
        {0,1,0,0,1,1,0,0,1,0,1,0,0},
        {0,1,0,0,1,1,0,0,1,1,1,0,0},
        {0,0,0,0,0,0,0,0,0,0,1,0,0},
        {0,0,0,0,0,0,0,1,1,1,0,0,0},
        {0,0,0,0,0,0,0,1,1,0,0,0,0},
    }
    fmt.Println(maxAreaOfIsland(island1)) // 6
    // Example 2:
    // Input: grid = [[0,0,0,0,0,0,0,0]]
    // Output: 0
    island2 := [][]int{
        {0,0,0,0,0,0,0,0},
    }
    fmt.Println(maxAreaOfIsland(island2)) // 0

    fmt.Println(maxAreaOfIsland1(island1)) // 6
    fmt.Println(maxAreaOfIsland1(island2)) // 0
}