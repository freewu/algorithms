package main

// 3239. Minimum Number of Flips to Make Binary Grid Palindromic I
// You are given an m x n binary matrix grid.

// A row or column is considered palindromic if its values read the same forward and backward.

// You can flip any number of cells in grid from 0 to 1, or from 1 to 0.

// Return the minimum number of cells that need to be flipped to make either all rows palindromic or all columns palindromic.

// Example 1:
// Input: grid = [[1,0,0],[0,0,0],[0,0,1]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/07/screenshot-from-2024-07-08-00-20-10.png" />
// Flipping the highlighted cells makes all the rows palindromic.

// Example 2:
// Input: grid = [[0,1],[0,1],[0,0]]
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/07/screenshot-from-2024-07-08-00-31-23.png" />
// Flipping the highlighted cell makes all the columns palindromic.

// Example 3:
// Input: grid = [[1],[0]]
// Output: 0
// Explanation:
// All rows are already palindromic.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m * n <= 2 * 10^5
//     0 <= grid[i][j] <= 1

import "fmt"

func minFlips(grid [][]int) int {
    m, l, r, c := 0, 0, len(grid), len(grid[0])
    for i := 0; i < r; i++ {
        for j := 0; j < c / 2; j++ {
            if grid[i][j] != grid[i][c-j-1] {
                m++
            }
        }
    }
    for i := 0; i < c; i++ {
        for j := 0; j < r / 2; j++ {
            if grid[j][i] != grid[r-1-j][i] {
                l++
            }
        }
    }
    if m < l { 
        return m 
    }
    return l
}

func main() {
    // Example 1:
    // Input: grid = [[1,0,0],[0,0,0],[0,0,1]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/07/screenshot-from-2024-07-08-00-20-10.png" />
    // Flipping the highlighted cells makes all the rows palindromic.
    fmt.Println(minFlips([][]int{{1,0,0},{0,0,0},{0,0,1}})) // 2
    // Example 2:
    // Input: grid = [[0,1],[0,1],[0,0]]
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/07/screenshot-from-2024-07-08-00-31-23.png" />
    // Flipping the highlighted cell makes all the columns palindromic.
    fmt.Println(minFlips([][]int{{0,1},{0,1},{0,0}})) // 1
    // Example 3:
    // Input: grid = [[1],[0]]
    // Output: 0
    // Explanation:
    // All rows are already palindromic.
    fmt.Println(minFlips([][]int{{1},{0}})) // 0
}