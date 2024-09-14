package main

// 2658. Maximum Number of Fish in a Grid
// You are given a 0-indexed 2D matrix grid of size m x n, where (r, c) represents:
//     A land cell if grid[r][c] = 0, or
//     A water cell containing grid[r][c] fish, if grid[r][c] > 0.

// A fisher can start at any water cell (r, c) and can do the following operations any number of times:
//     Catch all the fish at cell (r, c), or
//     Move to any adjacent water cell.

// Return the maximum number of fish the fisher can catch if he chooses his starting cell optimally, or 0 if no water cell exists.

// An adjacent cell of the cell (r, c), is one of the cells (r, c + 1), (r, c - 1), (r + 1, c) or (r - 1, c) if it exists.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/03/29/example.png" />
// Input: grid = [[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]
// Output: 7
// Explanation: The fisher can start at cell (1,3) and collect 3 fish, then move to cell (2,3) and collect 4 fish.

// Example 2:
// <img src="" />
// Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]
// Output: 1
// Explanation: The fisher can start at cells (0,0) or (3,3) and collect a single fish. 

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10
//     0 <= grid[i][j] <= 10

import "fmt"

func findMaxFish(grid [][]int) int {
    moves := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    res, n, m := 0, len(grid), len(grid[0])
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(row,col int, grid[][]int) int
    dfs = func(row, col int, grid[][]int) int {
        if row < 0 || row == n || col < 0 || col == m { return 0 } // 边界检测
        if grid[row][col] == 0 { return 0 } // 陆地
        res := grid[row][col]
        grid[row][col] = 0
        for _, move := range moves {
            res += dfs(row + move[0] , col + move[1] ,grid)
        }
        return res
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            res = max(res, dfs(i,j,grid))
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/03/29/example.png" />
    // Input: grid = [[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]
    // Output: 7
    // Explanation: The fisher can start at cell (1,3) and collect 3 fish, then move to cell (2,3) and collect 4 fish.
    grid1 := [][]int{
        {0,2,1,0},
        {4,0,0,3},
        {1,0,0,4},
        {0,3,2,0},
    }
    fmt.Println(findMaxFish(grid1)) // 7
    // Example 2:
    // <img src="" />
    // Input: grid = [[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]
    // Output: 1
    // Explanation: The fisher can start at cells (0,0) or (3,3) and collect a single fish. 
    grid2 := [][]int{
        {1,0,0,0},
        {0,0,0,0},
        {0,0,0,0},
        {0,0,0,1},
    }
    fmt.Println(findMaxFish(grid2)) // 1
}