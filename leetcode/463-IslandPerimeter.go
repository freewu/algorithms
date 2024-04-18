package main

// 463. Island Perimeter
// You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.
// Grid cells are connected horizontally/vertically (not diagonally). 
// The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

// The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. 
// One cell is a square with side length 1.
// The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/island.png" />
// Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
// Output: 16
// Explanation: The perimeter is the 16 yellow stripes in the image above.

// Example 2:
// <img src="" />
// Input: grid = [[1]]
// Output: 4

// Example 3:
// <img src="" />
// Input: grid = [[1,0]]
// Output: 4
 
// Constraints:
//     row == grid.length
//     col == grid[i].length
//     1 <= row, col <= 100
//     grid[i][j] is 0 or 1.
//     There is exactly one island in grid.

import "fmt"

func islandPerimeter(grid [][]int) int {
    res := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            // 判断四周边界的情况依次加一
            if grid[i][j] == 1 {
                if i-1 < 0 || grid[i-1][j] == 0 {
                    res++
                }
                if i + 1 >= len(grid) || grid[i+1][j] == 0 {
                    res++
                }
                if j - 1 < 0 || grid[i][j-1] == 0 {
                    res++
                }
                if j + 1 >= len(grid[0]) || grid[i][j+1] == 0 {
                    res++
                }
            }
        }
    }
    return res
}

// dfs
func islandPerimeter1(grid [][]int) int {
    var dfs func(x, y int, grid [][]int) int 
    dfs = func(x, y int, grid [][]int) int {
        if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[0])-1 {
            return 1
        }
        if grid[x][y] == 0 {
            return 1
        }
        if grid[x][y] == 2 {
            return 0
        }
        grid[x][y] = 2
        return dfs(x+1, y, grid) + dfs(x-1, y, grid) + dfs(x, y+1, grid) + dfs(x, y-1, grid)
    }
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                return dfs(i, j, grid)
            }
        }
    }
    return 0
}

func islandPerimeter2(grid [][]int) int {
    dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
    rows, cols, res := len(grid), len(grid[0]), 0
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if grid[row][col] == 1 {
                for _, dir := range dirs {
                    r, c := row + dir[0], col + dir[1]
                    // Check to see if the neighbour is an edge tile or water.
                    if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
                        res++
                    }
                }
            }
        }
    }
    return res
}

func main() {
    fmt.Println(islandPerimeter([][]int {{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}})) // 16
    fmt.Println(islandPerimeter([][]int{{1}})) // 4
    fmt.Println(islandPerimeter([][]int{{1,0}})) // 4

    fmt.Println(islandPerimeter1([][]int {{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}})) // 16
    fmt.Println(islandPerimeter1([][]int{{1}})) // 4
    fmt.Println(islandPerimeter1([][]int{{1,0}})) // 4

    fmt.Println(islandPerimeter2([][]int {{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}})) // 16
    fmt.Println(islandPerimeter2([][]int{{1}})) // 4
    fmt.Println(islandPerimeter2([][]int{{1,0}})) // 4
}