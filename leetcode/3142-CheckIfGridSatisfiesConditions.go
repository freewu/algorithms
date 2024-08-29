package main

// 3142. Check if Grid Satisfies Conditions
// You are given a 2D matrix grid of size m x n. 
// You need to check if each cell grid[i][j] is:
//     Equal to the cell below it, i.e. grid[i][j] == grid[i + 1][j] (if it exists).
//     Different from the cell to its right, i.e. grid[i][j] != grid[i][j + 1] (if it exists).

// Return true if all the cells satisfy these conditions, otherwise, return false.

// Example 1:
// Input: grid = [[1,0,2],[1,0,2]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/15/examplechanged.png" />
// All the cells in the grid satisfy the conditions.

// Example 2:
// Input: grid = [[1,1,1],[0,0,0]]
// Output: false
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/03/27/example21.png" />
// All cells in the first row are equal.

// Example 3:
// Input: grid = [[1],[2],[3]]
// Output: false
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/03/31/changed.png" />
// Cells in the first column have different values.

// Constraints:
//     1 <= n, m <= 10
//     0 <= grid[i][j] <= 9

import "fmt"

func satisfiesConditions(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i < m - 1 && grid[i][j] != grid[i + 1][j] { // 如果它下面的格子存在，那么它需要等于它下面的格子，也就是 grid[i][j] == grid[i + 1][j]
                return false
            }
            if j < n - 1 && grid[i][j] == grid[i][j + 1] { // 如果它右边的格子存在，那么它需要不等于它右边的格子，也就是 grid[i][j] != grid[i][j + 1]
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: grid = [[1,0,2],[1,0,2]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/15/examplechanged.png" />
    // All the cells in the grid satisfy the conditions.
    fmt.Println(satisfiesConditions([][]int{{1,0,2},{1,0,2}})) // true
    // Example 2:
    // Input: grid = [[1,1,1],[0,0,0]]
    // Output: false
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/03/27/example21.png" />
    // All cells in the first row are equal.
    fmt.Println(satisfiesConditions([][]int{{1,1,1},{0,0,0}})) // false
    // Example 3:
    // Input: grid = [[1],[2],[3]]
    // Output: false
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/03/31/changed.png" />
    // Cells in the first column have different values.
    fmt.Println(satisfiesConditions([][]int{{1},{2},{3}})) // false

    fmt.Println(satisfiesConditions([][]int{{1,2,3}})) // true
}