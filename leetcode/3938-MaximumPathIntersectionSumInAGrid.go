package main

// 3938. Maximum Path Intersection Sum in a Grid
// You are given an m x n integer matrix grid.

// Two players move across the grid:
//     1. Player 1 starts at the top-left cell (0, 0) and can move only right or down. Their destination is the bottom-right cell (m - 1, n - 1).
//     2. Player 2 starts at the bottom-left cell (m - 1, 0) and can move only right or up. Their destination is the top-right cell (0, n - 1).

// Each player must choose a valid path from their respective starting cell to their destination.

// A cell is called shared if it belongs to both chosen paths.

// Return an integer denoting the maximum possible sum of values of all shared cells.

// Example 1:
// ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/19/image.png" />
// Input: grid = [[1,2,0,-3],[1,-2,1,0],[-4,2,-1,3],[3,-3,3,-2],[-1,-5,0,1]]
// Output: 4
// Explanation:
// The diagram shows one optimal choice of paths.
// Player 1 follows the red/purple path from the top-left cell to the bottom-right cell:
// (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3) → (3, 3) → (4, 3)
// Player 2 follows the blue/purple path from the bottom-left cell to the top-right cell:
// (4, 0) → (4, 1) → (3, 1) → (2, 1) → (2, 2) → (2, 3) → (1, 3) → (0, 3)
// The shared cells are (2, 1), (2, 2), and (2, 3).
// The sum is 2 + (-1) + 3 = 4, which is the maximum possible sum.

// Example 2:
// ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/19/chatgpt-image-may-19-2026-01_39_39-pm.png" />
// Input: grid = [[4,-2,-3],[-1,-3,-1],[-4,2,-1]]
// Output: 3
// Explanation:
// One optimal pair of paths is shown in the diagram.
// Player 1 follows the red/purple path:
// (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
// Player 2 follows the blue/purple path:
// (2, 0) → (1, 0) → (0, 0) → (0, 1) → (0, 2)
// The shared cells are (0, 0) and (1, 0).
// The sum is 4 + (-1) = 3, which is the maximum possible.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 1000
//     4 <= m * n <= 5 * 10^5
//     -100 <= grid[i][j] <= 100

import "fmt"
import "slices"

func maxScore(grid [][]int) int {
    res, m, n := -1 << 61,len(grid), len(grid[0])
    // 单独计算子数组长为 1 的情况，此时子数组不能在 grid 的边界上
    if m > 2 && n > 2 {
        for _, row := range grid[1 : m-1] {
            res = max(res, slices.Max(row[1:n-1]))
        }
    }
    maxSubArray := func(nums []int) int { // 最大子数组和（子数组长度 >= 2）
        res, f := -1 << 61,nums[0]
        for _, x := range nums[1:] {
            res = max(res, f + x) // f+x 保证子数组至少有两个数
            f = max(f, 0) + x
        }
        return res
    }
    // 每行的最大子数组和（子数组长度 >= 2）
    for _, row := range grid {
        res = max(res, maxSubArray(row))
    }
    // 每列的最大子数组和（子数组长度 >= 2）
    col := make([]int, m)
    for j := range n {
        for i, row := range grid {
            col[i] = row[j]
        }
        res = max(res, maxSubArray(col))
    }
    return res
}

func maxScore1(grid [][]int) int {
    res, n, m := -1 << 61, len(grid), len(grid[0])
    if n == 1 || m == 1 {
        sum := 0
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                sum += grid[i][j]
            }
        }
        return sum
    }
    for i := 0; i < n; i++ {
        curr := grid[i][0]
        for j := 1; j < m; j++ {
            val := grid[i][j]
            if curr + val > res {
                res = curr + val
            }
            curr = max(val, curr + val)
        }    
    }
    for j := 0; j < m; j++ {
        curr := grid[0][j]
        for i := 1; i < n; i++ {
            val := grid[i][j]
            if curr + val > res {
                res = curr + val
            }

            curr = max(val, curr + val)
        }    
    }
    for i := 1; i < n-1; i++ {
        for j := 1; j < m-1; j++ {
            if grid[i][j] > res {
                res = grid[i][j]
            }
        }
    }
    return res
}

func maxScore2(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    if n == 1 {
        res := 0
        for _, x := range grid[0] {
            res += x
        }
        return res
    }
    if m == 1 {
        res := 0
        for i := 0; i < n; i++ {
            res += grid[i][0]
        }
        return res
    }
    res := -1 << 61
    for i := range n {
        curr := grid[i][0]
        for j := 1; j < m; j++ {
            res = max(res, curr + grid[i][j])
            curr = max(0, curr) + grid[i][j]
        }
    }
    for j := range m {
        curr := grid[0][j]
        for i := 1; i < n; i++ {
            res = max(res, curr + grid[i][j])
            curr = max(0, curr) + grid[i][j]
        }
    }
    for i := 1; i < n-1; i++ {
        for j := 1; j < m-1; j++ {
            res = max(res, grid[i][j])
        }
    }
    return res
}

func main() {
    // Example 1:
    // ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/19/image.png" />
    // Input: grid = [[1,2,0,-3],[1,-2,1,0],[-4,2,-1,3],[3,-3,3,-2],[-1,-5,0,1]]
    // Output: 4
    // Explanation:
    // The diagram shows one optimal choice of paths.
    // Player 1 follows the red/purple path from the top-left cell to the bottom-right cell:
    // (0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3) → (3, 3) → (4, 3)
    // Player 2 follows the blue/purple path from the bottom-left cell to the top-right cell:
    // (4, 0) → (4, 1) → (3, 1) → (2, 1) → (2, 2) → (2, 3) → (1, 3) → (0, 3)
    // The shared cells are (2, 1), (2, 2), and (2, 3).
    // The sum is 2 + (-1) + 3 = 4, which is the maximum possible sum.
    fmt.Println(maxScore([][]int{{1,2,0,-3},{1,-2,1,0},{-4,2,-1,3},{3,-3,3,-2},{-1,-5,0,1}})) // 4
    // Example 2:
    // ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/19/chatgpt-image-may-19-2026-01_39_39-pm.png" />
    // Input: grid = [[4,-2,-3],[-1,-3,-1],[-4,2,-1]]
    // Output: 3
    // Explanation:
    // One optimal pair of paths is shown in the diagram.
    // Player 1 follows the red/purple path:
    // (0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)
    // Player 2 follows the blue/purple path:
    // (2, 0) → (1, 0) → (0, 0) → (0, 1) → (0, 2)
    // The shared cells are (0, 0) and (1, 0).
    // The sum is 4 + (-1) = 3, which is the maximum possible.
    fmt.Println(maxScore([][]int{{4,-2,-3},{-1,-3,-1},{-4,2,-1}})) // 3

    fmt.Println(maxScore1([][]int{{1,2,0,-3},{1,-2,1,0},{-4,2,-1,3},{3,-3,3,-2},{-1,-5,0,1}})) // 4
    fmt.Println(maxScore1([][]int{{4,-2,-3},{-1,-3,-1},{-4,2,-1}})) // 3

    fmt.Println(maxScore2([][]int{{1,2,0,-3},{1,-2,1,0},{-4,2,-1,3},{3,-3,3,-2},{-1,-5,0,1}})) // 4
    fmt.Println(maxScore2([][]int{{4,-2,-3},{-1,-3,-1},{-4,2,-1}})) // 3
}