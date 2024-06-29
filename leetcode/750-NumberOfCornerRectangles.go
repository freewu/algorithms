package main

// 750. Number Of Corner Rectangles
// Given an m x n integer matrix grid where each entry is only 0 or 1, return the number of corner rectangles.

// A corner rectangle is four distinct 1's on the grid that forms an axis-aligned rectangle. 
// Note that only the corners need to have the value 1. Also, all four 1's used must be distinct.

// Example 1:
// <img src="" />
// Input: grid = [[1,0,0,1,0],[0,0,1,0,1],[0,0,0,1,0],[1,0,1,0,1]]
// Output: 1
// Explanation: There is only one corner rectangle, with corners grid[1][2], grid[1][4], grid[3][2], grid[3][4].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/12/cornerrec2-grid.jpg" />
// Input: grid = [[1,1,1],[1,1,1],[1,1,1]]
// Output: 9
// Explanation: There are four 2x2 rectangles, four 2x3 and 3x2 rectangles, and one 3x3 rectangle.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/12/cornerrec3-grid.jpg" />
// Input: grid = [[1,1,1,1]]
// Output: 0
// Explanation: Rectangles must have four distinct corners.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 200
//     grid[i][j] is either 0 or 1.
//     The number of 1's in the grid is in the range [1, 6000].

import "fmt"

// 我们用 count[i, j] 来记录 row[i] = row[j] = 1 的次数。
// 当我们处理新的一行时，对于每一对 new_row[i] = new_row[j] = 1，我们添加 count[i, j] 到答案中，然后 count[i, j]++
func countCornerRectangles(grid [][]int) int {
    res, n, m := 0, len(grid), len(grid[0])
    if n < 1 || m < 1 {
        return 0
    }
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            count := 0
            for k := 0; k < m; k++ {
                if grid[i][k] == 1 && grid[j][k] == 1 { // count[i, j] 来记录 row[i] = row[j] = 1 的次数
                    count++
                    res += count - 1
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: grid = [[1,0,0,1,0],[0,0,1,0,1],[0,0,0,1,0],[1,0,1,0,1]]
    // Output: 1
    // Explanation: There is only one corner rectangle, with corners grid[1][2], grid[1][4], grid[3][2], grid[3][4].
    grid1 := [][]int{
        {1,0,0,1,0},
        {0,0,1,0,1},
        {0,0,0,1,0},
        {1,0,1,0,1},
    }
    fmt.Println(countCornerRectangles(grid1)) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/12/cornerrec2-grid.jpg" />
    // Input: grid = [[1,1,1],[1,1,1],[1,1,1]]
    // Output: 9
    // Explanation: There are four 2x2 rectangles, four 2x3 and 3x2 rectangles, and one 3x3 rectangle.
    grid2 := [][]int{
        {1,1,1},
        {1,1,1},
        {1,1,1},
    }
    fmt.Println(countCornerRectangles(grid2)) // 9
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/12/cornerrec3-grid.jpg" />
    // Input: grid = [[1,1,1,1]]
    // Output: 0
    // Explanation: Rectangles must have four distinct corners.
    grid3 := [][]int{
        {1,1,1,1},
    }
    fmt.Println(countCornerRectangles(grid3)) // 0
}