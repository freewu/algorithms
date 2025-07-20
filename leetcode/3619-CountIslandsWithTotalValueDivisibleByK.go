package main

// 3619. Count Islands With Total Value Divisible by K
// You are given an m x n matrix grid and a positive integer k. 
// An island is a group of positive integers (representing land) that are 4-directionally connected (horizontally or vertically).

// The total value of an island is the sum of the values of all cells in the island.

// Return the number of islands with a total value divisible by k.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2025/03/06/example1griddrawio-1.png" />
// Input: grid = [[0,2,1,0,0],[0,5,0,0,5],[0,0,1,0,0],[0,1,4,7,0],[0,2,0,0,8]], k = 5
// Output: 2
// Explanation:
// The grid contains four islands. The islands highlighted in blue have a total value that is divisible by 5, while the islands highlighted in red do not.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2025/03/06/example2griddrawio.png" />
// Input: grid = [[3,0,3,0], [0,3,0,3], [3,0,3,0]], k = 3
// Output: 6
// Explanation:
// The grid contains six islands, each with a total value that is divisible by 3.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     0 <= grid[i][j] <= 10^6
//     1 <= k <= 10^6

import "fmt"

func countIslands(grid [][]int, k int) int {
    res, m, n := 0, len(grid), len(grid[0])
    var dfs func(i, j int) int
    dfs = func(i,j int) int {
        paths := [][]int{
            {i + 1, j},
            {i - 1, j},
            {i, j + 1},
            {i, j - 1},
        }
        val := grid[i][j]
        grid[i][j] = 0
        for _, path := range paths {
            r, c := path[0], path[1]
            if -1 < r && r < m && -1 < c && c < n && grid[r][c] != 0 {
                val += dfs(r, c)
            }
        }
        return val
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                total := dfs(i,j)
                if total % k == 0 {
                    res++
                }
            }
        }
    }
    return res
}

func countIslands1(grid [][]int, k int) int {
    res, m, n := 0, len(grid), len(grid[0])
    var dfs func( i, j int, count *int)
    dfs = func( i, j int, count *int) {
        if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] <= 0 { return }
        *count = *count + grid[i][j]
        grid[i][j]=-1
        dfs(i-1,j,count)
        dfs(i,j+1,count)
        dfs(i+1,j,count)
        dfs(i,j-1,count)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            count := 0
            if grid[i][j] > 0 {
                dfs(i, j, &count)
                if count % k == 0 {
                    res++
                }
            }
        }
    }
    return res
}

func main() { 
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2025/03/06/example1griddrawio-1.png" />
    // Input: grid = [[0,2,1,0,0],[0,5,0,0,5],[0,0,1,0,0],[0,1,4,7,0],[0,2,0,0,8]], k = 5
    // Output: 2
    // Explanation:
    // The grid contains four islands. The islands highlighted in blue have a total value that is divisible by 5, while the islands highlighted in red do not.
    fmt.Println(countIslands([][]int{{0,2,1,0,0},{0,5,0,0,5},{0,0,1,0,0},{0,1,4,7,0},{0,2,0,0,8}}, 5)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2025/03/06/example2griddrawio.png" />
    // Input: grid = [[3,0,3,0], [0,3,0,3], [3,0,3,0]], k = 3
    // Output: 6
    // Explanation:
    // The grid contains six islands, each with a total value that is divisible by 3.
    fmt.Println(countIslands([][]int{{3,0,3,0}, {0,3,0,3}, {3,0,3,0}}, 3)) // 6

    fmt.Println(countIslands1([][]int{{0,2,1,0,0},{0,5,0,0,5},{0,0,1,0,0},{0,1,4,7,0},{0,2,0,0,8}}, 5)) // 2
    fmt.Println(countIslands1([][]int{{3,0,3,0}, {0,3,0,3}, {3,0,3,0}}, 3)) // 6
}