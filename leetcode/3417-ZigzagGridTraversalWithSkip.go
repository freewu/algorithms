package main

// 3417. Zigzag Grid Traversal With Skip
// You are given an m x n 2D array grid of positive integers.

// Your task is to traverse grid in a zigzag pattern while skipping every alternate cell.

// Zigzag pattern traversal is defined as following the below actions:
//     1. Start at the top-left cell (0, 0).
//     2. Move right within a row until the end of the row is reached.
//     3. Drop down to the next row, then traverse left until the beginning of the row is reached.
//     4. Continue alternating between right and left traversal until every row has been traversed.

// Note that you must skip every alternate cell during the traversal.

// Return an array of integers result containing, 
// in order, the value of the cells visited during the zigzag traversal with skips.

// Example 1:
// Input: grid = [[1,2],[3,4]]
// Output: [1,4]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example0.png" />

// Example 2:
// Input: grid = [[2,1],[2,1],[2,1]]
// Output: [2,1,2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example1.png" />

// Example 3:
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,3,5,7,9]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example2.png" />

// Constraints:
//     2 <= n == grid.length <= 50
//     2 <= m == grid[i].length <= 50
//     1 <= grid[i][j] <= 2500

import "fmt"

func zigzagTraversal(grid [][]int) []int {
    res := []int{}
    for i := 0; i < len(grid); i++ {
        if i % 2 == 1 {
            for j := len(grid[0]) - 1; j >= 0; j-- {
                if j % 2 == 1 {
                    res = append(res, grid[i][j])
                }
            }
        } else {
            for j := 0; j < len(grid[0]); j++ {
                if j % 2 == 0 {
                    res = append(res, grid[i][j])
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,2],[3,4]]
    // Output: [1,4]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example0.png" />
    fmt.Println(zigzagTraversal([][]int{{1,2},{3,4}})) // [1,4]
    // Example 2:
    // Input: grid = [[2,1],[2,1],[2,1]]
    // Output: [2,1,2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example1.png" />
    fmt.Println(zigzagTraversal([][]int{{2,1},{2,1},{2,1}})) // [2,1,2]
    // Example 3:
    // Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
    // Output: [1,3,5,7,9]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/23/4012_example2.png" />
    fmt.Println(zigzagTraversal([][]int{{1,2,3},{4,5,6},{7,8,9}})) // [1,3,5,7,9]
}