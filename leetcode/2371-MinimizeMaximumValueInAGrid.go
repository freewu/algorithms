package main

// 2371. Minimize Maximum Value in a Grid
// You are given an m x n integer matrix grid containing distinct positive integers.

// You have to replace each integer in the matrix with a positive integer satisfying the following conditions:
//     1. The relative order of every two elements 
//        that are in the same row or column should stay the same after the replacements.
//     2. The maximum number in the matrix after the replacements should be as small as possible.

// The relative order stays the same if for all pairs of elements in the original matrix 
// such that grid[r1][c1] > grid[r2][c2] where either r1 == r2 or c1 == c2, 
// then it must be true that grid[r1][c1] > grid[r2][c2] after the replacements.

// For example, if grid = [[2, 4, 5], [7, 3, 9]] then a good replacement could be either 
// grid = [[1, 2, 3], [2, 1, 4]] or grid = [[1, 2, 3], [3, 1, 4]].

// Return the resulting matrix. 
// If there are multiple answers, return any of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/08/09/grid2drawio.png" />
// Input: grid = [[3,1],[2,5]]
// Output: [[2,1],[1,2]]
// Explanation: The above diagram shows a valid replacement.
// The maximum number in the matrix is 2. It can be shown that no smaller value can be obtained.

// Example 2:
// Input: grid = [[10]]
// Output: [[1]]
// Explanation: We replace the only number in the matrix with 1.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     1 <= grid[i][j] <= 10^9
//     grid consists of distinct integers.

import "fmt"
import "sort"

func minScore(grid [][]int) [][]int {
    a, m, n := [][]int{}, len(grid), len(grid[0])
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            a = append(a, []int{i, j})
        }
    }
    sort.Slice(a, func(i, j int) bool {
        return grid[a[i][0]][a[i][1]] < grid[a[j][0]][a[j][1]]
    })
    maxCol, maxRow := make([]int, n), make([]int, m)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range a {
        i, j := v[0], v[1]
        grid[i][j] = max(maxCol[j], maxRow[i]) + 1
        maxRow[i], maxCol[j] = grid[i][j], grid[i][j]
    }
    return grid
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/08/09/grid2drawio.png" />
    // Input: grid = [[3,1],[2,5]]
    // Output: [[2,1],[1,2]]
    // Explanation: The above diagram shows a valid replacement.
    // The maximum number in the matrix is 2. It can be shown that no smaller value can be obtained.
    fmt.Println(minScore([][]int{{3,1},{2,5}})) // [[2,1],[1,2]]
    // Example 2:
    // Input: grid = [[10]]
    // Output: [[1]]
    // Explanation: We replace the only number in the matrix with 1.
    fmt.Println(minScore([][]int{{10}})) // [[1]]
}