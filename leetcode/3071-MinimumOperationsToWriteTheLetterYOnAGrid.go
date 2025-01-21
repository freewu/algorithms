package main

// 3071. Minimum Operations to Write the Letter Y on a Grid
// You are given a 0-indexed n x n grid where n is odd, and grid[r][c] is 0, 1, or 2.

// We say that a cell belongs to the Letter Y if it belongs to one of the following:
//     1. The diagonal starting at the top-left cell and ending at the center cell of the grid.
//     2. The diagonal starting at the top-right cell and ending at the center cell of the grid.
//     3. The vertical line starting at the center cell and ending at the bottom border of the grid.

// The Letter Y is written on the grid if and only if:
//     1. All values at cells belonging to the Y are equal.
//     2. All values at cells not belonging to the Y are equal.
//     3. The values at cells belonging to the Y are different from the values at cells not belonging to the Y.

// Return the minimum number of operations needed to write the letter Y on the grid given 
// that in one operation you can change the value at any cell to 0, 1, or 2.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/01/22/y2.png" />
// Input: grid = [[1,2,2],[1,1,0],[0,1,0]]
// Output: 3
// Explanation: We can write Y on the grid by applying the changes highlighted in blue in the image above. After the operations, all cells that belong to Y, denoted in bold, have the same value of 1 while those that do not belong to Y are equal to 0.
// It can be shown that 3 is the minimum number of operations needed to write Y on the grid.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/01/22/y3.png" />
// Input: grid = [[0,1,0,1,0],[2,1,0,1,2],[2,2,2,0,1],[2,2,2,2,2],[2,1,2,2,2]]
// Output: 12
// Explanation: We can write Y on the grid by applying the changes highlighted in blue in the image above. After the operations, all cells that belong to Y, denoted in bold, have the same value of 0 while those that do not belong to Y are equal to 2. 
// It can be shown that 12 is the minimum number of operations needed to write Y on the grid.

// Constraints:
//     3 <= n <= 49 
//     n == grid.length == grid[i].length
//     0 <= grid[i][j] <= 2
//     n is odd.

import "fmt"

func minimumOperationsToWriteY(grid [][]int) int {
    n := len(grid)
    count1, count2 := [3]int{}, [3]int{}
    for i, row := range grid {
        for j, v := range row {
            if (i == j && i <= n / 2) || (i + j == n - 1 && i <= n / 2) || (j == n / 2 && i >= n / 2) {
                count1[v]++
            } else {
                count2[v]++
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := n * n
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i != j {
                res = min(res, n * n - count1[i] - count2[j])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/01/22/y2.png" />
    // Input: grid = [[1,2,2],[1,1,0],[0,1,0]]
    // Output: 3
    // Explanation: We can write Y on the grid by applying the changes highlighted in blue in the image above. After the operations, all cells that belong to Y, denoted in bold, have the same value of 1 while those that do not belong to Y are equal to 0.
    // It can be shown that 3 is the minimum number of operations needed to write Y on the grid.
    fmt.Println(minimumOperationsToWriteY([][]int{{1,2,2},{1,1,0},{0,1,0}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/01/22/y3.png" />
    // Input: grid = [[0,1,0,1,0],[2,1,0,1,2],[2,2,2,0,1],[2,2,2,2,2],[2,1,2,2,2]]
    // Output: 12
    // Explanation: We can write Y on the grid by applying the changes highlighted in blue in the image above. After the operations, all cells that belong to Y, denoted in bold, have the same value of 0 while those that do not belong to Y are equal to 2. 
    // It can be shown that 12 is the minimum number of operations needed to write Y on the grid.
    fmt.Println(minimumOperationsToWriteY([][]int{{0,1,0,1,0},{2,1,0,1,2},{2,2,2,0,1},{2,2,2,2,2},{2,1,2,2,2}})) // 12
}