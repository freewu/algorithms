package main

// 1254. Number of Closed Islands
// Given a 2D grid consists of 0s (land) and 1s (water).  
// An island is a maximal 4-directionally connected group of 0s and a closed island is an island totally (all left, top, right, bottom) surrounded by 1s.
// Return the number of closed islands.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/10/31/sample_3_1610.png" />
// Input: grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
// Output: 2
// Explanation: 
// Islands in gray are closed because they are completely surrounded by water (group of 1s).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/10/31/sample_4_1610.png" />
// Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
// Output: 1

// Example 3:
// Input: grid = [[1,1,1,1,1,1,1],
//                [1,0,0,0,0,0,1],
//                [1,0,1,1,1,0,1],
//                [1,0,1,0,1,0,1],
//                [1,0,1,1,1,0,1],
//                [1,0,0,0,0,0,1],
//                [1,1,1,1,1,1,1]]
// Output: 2
 
// Constraints:
//     1 <= grid.length, grid[0].length <= 100
//     0 <= grid[i][j] <=1

import "fmt"

func closedIsland(grid [][]int) int {
    if len(grid) <= 2 || len(grid[0]) <= 2 {
        return 0
    }
    var dfs func (grid [][]int, i, j int, footprint [][]bool, change bool)
    dfs = func (grid [][]int, i, j int, footprint [][]bool, change bool) {
        if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
            return
        }
        if footprint[i][j] == true {
            return
        }
        if grid[i][j] == 1 {
            return
        }
        footprint[i][j] = true
        if change == true {
            grid[i][j] = 1
        }
        dfs(grid, i-1, j, footprint, change)
        dfs(grid, i+1, j, footprint, change)
        dfs(grid, i, j-1, footprint, change)
        dfs(grid, i, j+1, footprint, change)
    }
    footprint := make([][]bool, len(grid))
    for i := range footprint {
        footprint[i] = make([]bool, len(grid[0]))
    }
    for _, i := range []int{0, len(grid)-1} {
        for j, v := range grid[i] {
            if v == 0 {
                dfs(grid, i, j, footprint, true)
            }
        }
    }
    for _, j := range []int{0, len(grid[0])-1} {
        for i := range grid {
            if grid[i][j] == 0 {
                dfs(grid, i, j, footprint, true)
            }
        }
    }
    closed := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 && footprint[i][j] == false {
                dfs(grid, i, j, footprint, false)
                closed++
            }
        }
    }
    return closed
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/10/31/sample_3_1610.png" />
    // Input: grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
    // Output: 2
    // Explanation: 
    // Islands in gray are closed because they are completely surrounded by water (group of 1s).
    grid1 := [][]int{
        {1,1,1,1,1,1,1,0},
        {1,0,0,0,0,1,1,0},
        {1,0,1,0,1,1,1,0},
        {1,0,0,0,0,1,0,1},
        {1,1,1,1,1,1,1,0},
    }
    fmt.Println(closedIsland(grid1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/10/31/sample_4_1610.png" />
    // Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
    // Output: 1
    grid2 := [][]int{
        {0,0,1,0,0},
        {0,1,0,1,0},
        {0,1,1,1,0},
    }
    fmt.Println(closedIsland(grid2)) // 1
    // Example 3:
    // Input: grid = [[1,1,1,1,1,1,1],
    //                [1,0,0,0,0,0,1],
    //                [1,0,1,1,1,0,1],
    //                [1,0,1,0,1,0,1],
    //                [1,0,1,1,1,0,1],
    //                [1,0,0,0,0,0,1],
    //                [1,1,1,1,1,1,1]]
    // Output: 2
    grid3 := [][]int{
        {1,1,1,1,1,1,1},
        {1,0,0,0,0,0,1},
        {1,0,1,1,1,0,1},
        {1,0,1,0,1,0,1},
        {1,0,1,1,1,0,1},
        {1,0,0,0,0,0,1},
        {1,1,1,1,1,1,1},
    }
    fmt.Println(closedIsland(grid3)) // 2
}