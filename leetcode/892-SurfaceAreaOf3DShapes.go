package main

// 892. Surface Area of 3D Shapes
// You are given an n x n grid where you have placed some 1 x 1 x 1 cubes. 
// Each value v = grid[i][j] represents a tower of v cubes placed on top of cell (i, j).

// After placing these cubes, you have decided to glue any directly adjacent cubes to each other, forming several irregular 3D shapes.

// Return the total surface area of the resulting shapes.
// Note: The bottom face of each shape counts toward its surface area.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/08/tmp-grid2.jpg" />
// Input: grid = [[1,2],[3,4]]
// Output: 34

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/08/tmp-grid4.jpg" />
// Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
// Output: 32

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/01/08/tmp-grid5.jpg" />
// Input: grid = [[2,2,2],[2,1,2],[2,2,2]]
// Output: 46

// Constraints:
//     n == grid.length == grid[i].length
//     1 <= n <= 50
//     0 <= grid[i][j] <= 50

import "fmt"

func surfaceArea(grid [][]int) int {
    res, n := 0, len(grid)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 {
                res += grid[i][j] * 4 + 2
            }
            if i > 0 {
                res -= (min(grid[i][j], grid[i - 1][j]) * 2)
            }
            if j > 0 {
                res -= (min(grid[i][j], grid[i][j - 1]) * 2)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/01/08/tmp-grid2.jpg" />
    // Input: grid = [[1,2],[3,4]]
    // Output: 34
    fmt.Println(surfaceArea([][]int{{1,2},{3,4}})) // 34
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/01/08/tmp-grid4.jpg" />
    // Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
    // Output: 32
    fmt.Println(surfaceArea([][]int{{1,1,1},{1,0,1},{1,1,1}})) // 32
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/01/08/tmp-grid5.jpg" />
    // Input: grid = [[2,2,2],[2,1,2],[2,2,2]]
    // Output: 46
    fmt.Println(surfaceArea([][]int{{2,2,2},{2,1,2},{2,2,2}})) // 46
}