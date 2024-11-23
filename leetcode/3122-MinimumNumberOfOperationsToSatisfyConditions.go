package main

// 3122. Minimum Number of Operations to Satisfy Conditions
// You are given a 2D matrix grid of size m x n. 
// In one operation, you can change the value of any cell to any non-negative number. 
// You need to perform some operations such that each cell grid[i][j] is:
//     Equal to the cell below it, i.e. grid[i][j] == grid[i + 1][j] (if it exists).
//     Different from the cell to its right, i.e. grid[i][j] != grid[i][j + 1] (if it exists).

// Return the minimum number of operations needed.

// Example 1:
// Input: grid = [[1,0,2],[1,0,2]]
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/04/15/examplechanged.png" />
// All the cells in the matrix already satisfy the properties.

// Example 2:
// Input: grid = [[1,1,1],[0,0,0]]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/03/27/example21.png" />
// The matrix becomes [[1,0,1],[1,0,1]] which satisfies the properties, by doing these 3 operations:
//     Change grid[1][0] to 1.
//     Change grid[0][1] to 0.
//     Change grid[1][2] to 1.

// Example 3:
// Input: grid = [[1],[2],[3]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/03/31/changed.png" />
// There is a single column. We can change the value to 1 in each cell using 2 operations.

// Constraints:
//     1 <= n, m <= 1000
//     0 <= grid[i][j] <= 9

import "fmt"

func minimumOperations(grid [][]int) int {
    dp := make([][]int, 1005)
    for i := range dp {
        dp[i] = make([]int, 15)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    calc := func(num, col int, grid [][]int) int {
        res := 0
        for i := 0; i < len(grid); i++ {
            if grid[i][col] != num {
                res++
            }
        }
        return res
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i, prev, n int) int 
    dfs = func(i, prev, n int) int {
        if i >= n { return 0 }
        if dp[i][prev + 1] != -1 { return dp[i][prev + 1] }
        res := 1 << 31
        for j := 0; j < 10; j++ {
            if j != prev {
                res = min(res, calc(j, i, grid) + dfs(i + 1, j, n))
            }
        }
        dp[i][prev + 1] = res
        return res
    }
    return dfs(0, -1, len(grid[0]))
}

func minimumOperations1(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    count := make([][10]int, n)
    for _, row := range grid {
        for j, v := range row {
            count[j][v]++
        }
    }
    memo := make([][11]int, n)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i < 0 { return 0 }
        if memo[i][j] != -1 { return memo[i][j] }
        res := 0
        for k, c := range count[i] {
            if k != j {
                res = max(res, dfs(i - 1, k) + c)
            }
        }
        memo[i][j] = res
        return res
    }
    return m * n - dfs(n - 1, 10)
}

func main() {
    // Example 1:
    // Input: grid = [[1,0,2],[1,0,2]]
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/04/15/examplechanged.png" />
    // All the cells in the matrix already satisfy the properties.
    fmt.Println(minimumOperations([][]int{{1,0,2},{1,0,2}})) // 0
    // Example 2:
    // Input: grid = [[1,1,1],[0,0,0]]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/03/27/example21.png" />
    // The matrix becomes [[1,0,1],[1,0,1]] which satisfies the properties, by doing these 3 operations:
    //     Change grid[1][0] to 1.
    //     Change grid[0][1] to 0.
    //     Change grid[1][2] to 1.
    fmt.Println(minimumOperations([][]int{{1,1,1},{0,0,0}})) // 3
    // Example 3:
    // Input: grid = [[1],[2],[3]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/03/31/changed.png" />
    // There is a single column. We can change the value to 1 in each cell using 2 operations.
    fmt.Println(minimumOperations([][]int{{1},{2},{3}})) // 2

    fmt.Println(minimumOperations1([][]int{{1,0,2},{1,0,2}})) // 0
    fmt.Println(minimumOperations1([][]int{{1,1,1},{0,0,0}})) // 3
    fmt.Println(minimumOperations1([][]int{{1},{2},{3}})) // 2
}