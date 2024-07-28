package main

// 3195. Find the Minimum Area to Cover All Ones I
// You are given a 2D binary array grid. 
// Find a rectangle with horizontal and vertical sides with the smallest area, such that all the 1's in grid lie inside this rectangle.

// Return the minimum possible area of the rectangle.

// Example 1:
// Input: grid = [[0,1,0],[1,0,1]]
// Output: 6
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/08/examplerect0.png" />
// The smallest rectangle has a height of 2 and a width of 3, so it has an area of 2 * 3 = 6.

// Example 2:
// Input: grid = [[1,0],[0,0]]
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/08/examplerect1.png" />
// The smallest rectangle has both height and width 1, so its area is 1 * 1 = 1.

// Constraints:
//     1 <= grid.length, grid[i].length <= 1000
//     grid[i][j] is either 0 or 1.
//     The input is generated such that there is at least one 1 in grid.

import "fmt"

func minimumArea(grid [][]int) int {
    r1, r2, c1, c2 := -1, -1, -1, -1
    for row := range grid {
        for col, v := range grid[row] {
            if v == 1 {
                if r1 == -1 { r1 = row }
                r2 = row
                if c1 == -1 { 
                    c1, c2 = col, col
                } else {
                    if col < c1 { c1 = col }
                    if col > c2 { c2 = col }
                }
            }
        }
    }
    return ((r2 - r1) + 1) * ((c2 - c1) + 1)
}

func main() {
    // Example 1:
    // Input: grid = [[0,1,0],[1,0,1]]
    // Output: 6
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/08/examplerect0.png" />
    // The smallest rectangle has a height of 2 and a width of 3, so it has an area of 2 * 3 = 6.
    fmt.Println(minimumArea([][]int{{0,1,0},{1,0,1}})) // 6
    // Example 2:
    // Input: grid = [[1,0],[0,0]]
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/08/examplerect1.png" />
    // The smallest rectangle has both height and width 1, so its area is 1 * 1 = 1.
    fmt.Println(minimumArea([][]int{{1,0},{0,0}})) // 1
}