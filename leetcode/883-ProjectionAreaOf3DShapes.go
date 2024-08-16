package main

// 883. Projection Area of 3D Shapes
// You are given an n x n grid where we place some 1 x 1 x 1 cubes that are axis-aligned with the x, y, and z axes.
// Each value v = grid[i][j] represents a tower of v cubes placed on top of the cell (i, j).
// We view the projection of these cubes onto the xy, yz, and zx planes.

// A projection is like a shadow, that maps our 3-dimensional figure to a 2-dimensional plane. 
// We are viewing the "shadow" when looking at the cubes from the top, the front, and the side.

// Return the total area of all three projections.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/02/shadow.png" />
// Input: grid = [[1,2],[3,4]]
// Output: 17
// Explanation: Here are the three projections ("shadows") of the shape made with each axis-aligned plane.

// Example 2:
// Input: grid = [[2]]
// Output: 5

// Example 3:
// Input: grid = [[1,0],[0,2]]
// Output: 8

// Constraints:
//     n == grid.length == grid[i].length
//     1 <= n <= 50
//     0 <= grid[i][j] <= 50

import "fmt"

func projectionArea(grid [][]int) int {
    res, n := 0, len(grid)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        x, y := 0, 0
        for j := 0; j < n; j++ {
            x = max(x, grid[i][j])
            y = max(y, grid[j][i])
            if grid[i][j] > 0 {
                res++
            }
        }
        res += (x + y)
    }
    return res
}

func projectionArea1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    res, xarr, yarr := 0, make([]int, m), make([]int, n)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] > 0 {
                res += 1
            }
            if grid[i][j] > xarr[i] { xarr[i] = grid[i][j] }
            if grid[i][j] > yarr[j] { yarr[j] = grid[i][j] }
        }
    }
    for _, v := range xarr { res += v }
    for _, v := range yarr { res += v }
    return res
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/02/shadow.png" />
    // Input: grid = [[1,2],[3,4]]
    // Output: 17
    // Explanation: Here are the three projections ("shadows") of the shape made with each axis-aligned plane.
    fmt.Println(projectionArea([][]int{{1, 2},{3, 4}})) // 17
    // Example 2:
    // Input: grid = [[2]]
    // Output: 5
    fmt.Println(projectionArea([][]int{{2}})) // 5
    // Example 3:
    // Input: grid = [[1,0],[0,2]]
    // Output: 8
    fmt.Println(projectionArea([][]int{{1, 0},{0, 2}})) // 8

    fmt.Println(projectionArea1([][]int{{1, 2},{3, 4}})) // 17
    fmt.Println(projectionArea1([][]int{{2}})) // 5
    fmt.Println(projectionArea1([][]int{{1, 0},{0, 2}})) // 8
}