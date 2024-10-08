package main

// 329. Longest Increasing Path in a Matrix
// Given an m x n integers matrix, return the length of the longest increasing path in matrix.
// From each cell, you can either move in four directions: 
//     left, right, up, or down. 
    
// You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg" />
// Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
// Output: 4
// Explanation: The longest increasing path is [1, 2, 6, 9].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg" />
// Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
// Output: 4
// Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.

// Example 3:
// Input: matrix = [[1]]
// Output: 1
 
// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 200
//     0 <= matrix[i][j] <= 2^31 - 1

import "fmt"

// dp
func longestIncreasingPath(matrix [][]int) int {
    dp, res := make([][]int, len(matrix)), 0
    var dir = [][]int{ {-1, 0}, {0, 1}, {1, 0}, {0, -1} }
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(matrix[0]))
    }
    max := func(x, y int) int { if x > y { return x; }; return y; }
    // 判断是否在边界里
    isInBoard := func(board [][]int, x, y int) bool { return x >= 0 && x < len(board) && y >= 0 && y < len(board[0]); }
    var searchPath func(board, cache [][]int, lastNum, x, y int) int
    searchPath = func(board, cache [][]int, lastNum, x, y int) int {
        if board[x][y] <= lastNum {
            return 0
        }
        if cache[x][y] > 0 {
            return cache[x][y]
        }
        count := 1
        for i := 0; i < 4; i++ {
            nx := x + dir[i][0]
            ny := y + dir[i][1]
            if isInBoard(board, nx, ny) {
                count = max(count, searchPath(board, cache, board[x][y], nx, ny) + 1)
            }
        }
        cache[x][y] = count
        return count
    }
    for i, v := range matrix {
        for j := range v {
            searchPath(matrix, dp, -1 << 31, i, j)
            res = max(res, dp[i][j])
        }
    }
    return res
}

func longestIncreasingPath1(matrix [][]int) int {
    m, n, res := len(matrix), len(matrix[0]), 0 
    cache := make([][]int, m)
    for i,_ := range cache {
        cache[i] = make([]int, n)
    }
    max := func(x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int, int) int
    dfs = func(i, j, v0 int) int {
        if i < 0 || j < 0 || i == m || j == n {
            return 0
        }
        v := matrix[i][j]
        if v <= v0 {
            return 0
        }
        if cache[i][j] > 0 {
            return cache[i][j]
        }
        l := 1 + max(max(max(dfs(i-1, j, v), dfs(i+1, j, v)), dfs(i, j-1, v)), dfs(i, j+1, v))
        cache[i][j] = l
        return l
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = max(res, dfs(i, j, -1))
        }
    }
    return res
}

func main() {
    // Explanation: The longest increasing path is [1, 2, 6, 9].
    fmt.Println(longestIncreasingPath([][]int{{9,9,4},{6,6,8},{2,1,1}})) // 4
    // Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
    fmt.Println(longestIncreasingPath([][]int{{3,4,5},{3,2,6},{2,2,1}})) // 4
    fmt.Println(longestIncreasingPath([][]int{{1}})) // 1

    fmt.Println(longestIncreasingPath1([][]int{{9,9,4},{6,6,8},{2,1,1}})) // 4
    fmt.Println(longestIncreasingPath1([][]int{{3,4,5},{3,2,6},{2,2,1}})) // 4
    fmt.Println(longestIncreasingPath1([][]int{{1}})) // 1
}