package main

// 2128. Remove All Ones With Row and Column Flips
// You are given an m x n binary matrix grid.
// In one operation, you can choose any row or column and flip each value in that row or column (i.e., changing all 0's to 1's, and all 1's to 0's).
// Return true if it is possible to remove all 1's from grid using any number of operations or false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/01/03/image-20220103191300-1.png" />
// Input: grid = [[0,1,0],[1,0,1],[0,1,0]]
// Output: true
// Explanation: One possible way to remove all 1's from grid is to:
// - Flip the middle row
// - Flip the middle column

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/01/03/image-20220103181204-7.png" />
// Input: grid = [[1,1,0],[0,0,0],[0,0,0]]
// Output: false
// Explanation: It is impossible to remove all 1's from grid.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/01/03/image-20220103181224-8.png" />
// Input: grid = [[0]]
// Output: true
// Explanation: There are no 1's in grid.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 300
//     grid[i][j] is either 0 or 1.

import "fmt"

func removeOnes(grid [][]int) bool {
    n, m := len(grid), len(grid[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[0][j] != grid[i][j] ^ grid[i][0] ^ grid[0][0] {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/01/03/image-20220103191300-1.png" />
    // Input: grid = [[0,1,0],[1,0,1],[0,1,0]]
    // Output: true
    // Explanation: One possible way to remove all 1's from grid is to:
    // - Flip the middle row
    // - Flip the middle column
    fmt.Println(removeOnes([][]int{{0,1,0},{1,0,1},{0,1,0}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/01/03/image-20220103181204-7.png" />
    // Input: grid = [[1,1,0],[0,0,0],[0,0,0]]
    // Output: false
    // Explanation: It is impossible to remove all 1's from grid.
    fmt.Println(removeOnes([][]int{{1,1,0},{0,0,0},{0,0,0}})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/01/03/image-20220103181224-8.png" />
    // Input: grid = [[0]]
    // Output: true
    // Explanation: There are no 1's in grid.
    fmt.Println(removeOnes([][]int{{0}})) // true
}