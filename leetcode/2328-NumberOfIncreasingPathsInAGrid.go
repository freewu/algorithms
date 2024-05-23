package main

// 2328. Number of Increasing Paths in a Grid
// You are given an m x n integer matrix grid, where you can move from a cell to any adjacent cell in all 4 directions.
// Return the number of strictly increasing paths in the grid such that you can start from any cell and end at any cell. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Two paths are considered different if they do not have exactly the same sequence of visited cells.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/05/10/griddrawio-4.png" />
// Input: grid = [[1,1],[3,4]]
// Output: 8
// Explanation: The strictly increasing paths are:
// - Paths with length 1: [1], [1], [3], [4].
// - Paths with length 2: [1 -> 3], [1 -> 4], [3 -> 4].
// - Paths with length 3: [1 -> 3 -> 4].
// The total number of paths is 4 + 3 + 1 = 8.

// Example 2:
// Input: grid = [[1],[2]]
// Output: 3
// Explanation: The strictly increasing paths are:
// - Paths with length 1: [1], [2].
// - Paths with length 2: [1 -> 2].
// The total number of paths is 2 + 1 = 3.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     1 <= grid[i][j] <= 10^5

import "fmt"

func countPaths(matrix [][]int) int {
    n, m, mod := len(matrix), len(matrix[0]), int(10e8) + 7
    res, memo := 0, make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, m)
    }
    var dfs func(r, c int, lastInt int) int
    dfs = func(r, c, lastInt int) int {
        if r < 0 || r >= len(matrix) || c < 0 || c >= m || matrix[r][c] <= lastInt {
            return 0
        }
        if memo[r][c] != 0 {
            return memo[r][c]
        }
        prev := matrix[r][c]
        memo[r][c] = (dfs(r-1, c, prev) + dfs(r+1, c, prev) + dfs(r, c+1, prev) + dfs(r, c-1, prev) + 1) % mod
        return memo[r][c]
    }
    for r := range matrix {
        for c := range matrix[r] {
            res = (res + dfs(r, c, -1)) % mod
        }
    }
    return res
}

func countPaths1(matrix [][]int) int {
    res, inf, mod, m, n := 0, 1 << 32 - 1, int(1e9) + 7,len(matrix), len(matrix[0])
    init := func (m, n, defaultVal int) [][]int {
        arr := make([][]int, m)
        for i := range arr {
            tmp := make([]int, n)
            for j := range tmp {
                tmp[j] = defaultVal
            }
            arr[i] = tmp
        }
        return arr
    }
    cache := init(m, n, -1)
    var dfs func(i, j, pre int) int
    dfs = func(i, j, pre int) int {
        if i < 0 || i >= m || j < 0 || j >= n || matrix[i][j] >= pre {
            return 0
        }
        if cache[i][j] != -1 {
            return cache[i][j]
        }
        v := matrix[i][j]
        res := dfs(i+1, j, v) + dfs(i-1, j, v) + dfs(i, j+1, v) + dfs(i, j-1, v)+1
        res %= mod
        cache[i][j] = res
        return res
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res = (res + dfs(i, j, inf)) % mod
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/05/10/griddrawio-4.png" />
    // Input: grid = [[1,1],[3,4]]
    // Output: 8
    // Explanation: The strictly increasing paths are:
    // - Paths with length 1: [1], [1], [3], [4].
    // - Paths with length 2: [1 -> 3], [1 -> 4], [3 -> 4].
    // - Paths with length 3: [1 -> 3 -> 4].
    // The total number of paths is 4 + 3 + 1 = 8.
    fmt.Println(countPaths([][]int{{1,1},{3,4}})) // 8
    // Example 2:
    // Input: grid = [[1],[2]]
    // Output: 3
    // Explanation: The strictly increasing paths are:
    // - Paths with length 1: [1], [2].
    // - Paths with length 2: [1 -> 2].
    // The total number of paths is 2 + 1 = 3.
    fmt.Println(countPaths([][]int{{1},{2}})) // 3

    fmt.Println(countPaths1([][]int{{1,1},{3,4}})) // 8
    fmt.Println(countPaths1([][]int{{1},{2}})) // 3
}