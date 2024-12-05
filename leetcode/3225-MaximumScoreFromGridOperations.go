package main

// 3225. Maximum Score From Grid Operations
// You are given a 2D matrix grid of size n x n. 
// Initially, all cells of the grid are colored white. 
// In one operation, you can select any cell of indices (i, j), 
// and color black all the cells of the jth column starting from the top row down to the ith row.

// The grid score is the sum of all grid[i][j] such that cell (i, j) is white and it has a horizontally adjacent black cell.

// Return the maximum score that can be achieved after some number of operations.

// Example 1:
// Input: grid = [[0,0,0,0,0],[0,0,3,0,0],[0,1,0,0,0],[5,0,0,3,0],[0,0,0,0,2]]
// Output: 11
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/11/one.png" />
// In the first operation, we color all cells in column 1 down to row 3, and in the second operation, we color all cells in column 4 down to the last row. 
// The score of the resulting grid is grid[3][0] + grid[1][2] + grid[3][3] which is equal to 11.

// Example 2:
// Input: grid = [[10,9,0,0,15],[7,1,0,8,0],[5,20,0,11,0],[0,0,0,1,2],[8,12,1,10,3]]
// Output: 94
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/11/two-1.png" />
// We perform operations on 1, 2, and 3 down to rows 1, 4, and 0, respectively. 
// The score of the resulting grid is grid[0][0] + grid[1][0] + grid[2][1] + grid[4][1] + grid[1][3] + grid[2][3] + grid[3][3] + grid[4][3] + grid[0][4] which is equal to 94.

// Constraints:
//     1 <= n == grid.length <= 100
//     n == grid[i].length
//     0 <= grid[i][j] <= 10^9

import "fmt"

func maximumScore(grid [][]int) int64 {
    n :=  len(grid)
    dp := make([][][2]int, n) // i, lastHeight, inclColScore
    for i := range dp {
        dp[i] = make([][2]int, n + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n - 1; i++ { // i is the column number
        for lastHeight := range dp[0] { // height of last column (i), height is the amount of colored cells
            lastColScore, nextColScore := 0, 0 // score from column i added by coloring column i+1 & score from column i+1 added by coloring column i
            for row := 0; row < lastHeight; row++ { // score from col i+1 calculated according to height of i
                nextColScore += grid[row][i+1]
            }
            for height := range dp[0] { // height of next column (i + 1)
                if height > 0 && height <= lastHeight { // next column doesn't contibute score from the colored cells
                    nextColScore -= grid[height-1][i+1]
                }
                if height > lastHeight { // add score from the neighbour of the newly colored cell
                    lastColScore += grid[height-1][i]
                }
                inclColScore, exclColScore := 1, 0 // indecies of the 2-cell array
                dp[i+1][height][exclColScore] = max(dp[i+1][height][exclColScore], max(dp[i][lastHeight][exclColScore] + lastColScore, dp[i][lastHeight][inclColScore]))
                dp[i+1][height][inclColScore] = max(dp[i+1][height][inclColScore], max(dp[i][lastHeight][inclColScore] + nextColScore, dp[i][lastHeight][exclColScore] + nextColScore + lastColScore))
            }
        }
    }
    res, i := 0, len(dp) - 1
    for lastHeight := range dp[i] { // find max score
        res = max(res, max(dp[i][lastHeight][0], dp[i][lastHeight][1]))
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: grid = [[0,0,0,0,0],[0,0,3,0,0],[0,1,0,0,0],[5,0,0,3,0],[0,0,0,0,2]]
    // Output: 11
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/11/one.png" />
    // In the first operation, we color all cells in column 1 down to row 3, and in the second operation, we color all cells in column 4 down to the last row. 
    // The score of the resulting grid is grid[3][0] + grid[1][2] + grid[3][3] which is equal to 11.
    fmt.Println(maximumScore([][]int{{0,0,0,0,0},{0,0,3,0,0},{0,1,0,0,0},{5,0,0,3,0},{0,0,0,0,2}})) // 11
    // Example 2:
    // Input: grid = [[10,9,0,0,15],[7,1,0,8,0],[5,20,0,11,0],[0,0,0,1,2],[8,12,1,10,3]]
    // Output: 94
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/11/two-1.png" />
    // We perform operations on 1, 2, and 3 down to rows 1, 4, and 0, respectively. 
    // The score of the resulting grid is grid[0][0] + grid[1][0] + grid[2][1] + grid[4][1] + grid[1][3] + grid[2][3] + grid[3][3] + grid[4][3] + grid[0][4] which is equal to 94.
    fmt.Println(maximumScore([][]int{{10,9,0,0,15},{7,1,0,8,0},{5,20,0,11,0},{0,0,0,1,2},{8,12,1,10,3}})) // 94
}