package main

// 741. Cherry Pickup
// You are given an n x n grid representing a field of cherries, each cell is one of three possible integers.
//     0 means the cell is empty, so you can pass through,
//     1 means the cell contains a cherry that you can pick up and pass through, or
//     -1 means the cell contains a thorn that blocks your way.

// Return the maximum number of cherries you can collect by following the rules below:
//     Starting at the position (0, 0) and reaching (n - 1, n - 1) by moving right or down through valid path cells (cells with value 0 or 1).
//     After reaching (n - 1, n - 1), returning to (0, 0) by moving left or up through valid path cells.
//     When passing through a path cell containing a cherry, you pick it up, and the cell becomes an empty cell 0.
//     If there is no valid path between (0, 0) and (n - 1, n - 1), then no cherries can be collected.
    
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/grid.jpg" />
// Input: grid = [[0,1,-1],[1,0,-1],[1,1,1]]
// Output: 5
// Explanation: The player started at (0, 0) and went down, down, right right to reach (2, 2).
// 4 cherries were picked up during this single trip, and the matrix becomes [[0,1,-1],[0,0,-1],[0,0,0]].
// Then, the player went left, up, up, left to return home, picking up one more cherry.
// The total number of cherries picked up is 5, and this is the maximum possible.

// Example 2:
// Input: grid = [[1,1,-1],[1,-1,1],[-1,1,1]]
// Output: 0
 
// Constraints:
//     n == grid.length
//     n == grid[i].length
//     1 <= n <= 50
//     grid[i][j] is -1, 0, or 1.
//     grid[0][0] != -1
//     grid[n - 1][n - 1] != -1

import "fmt"

func cherryPickup(grid [][]int) int {
    n, dp := len(grid), make([][][]int, len(grid))
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var solve func(int, int, int) int
    solve = func(r1, c1, c2 int) int {
        r2 := r1 + c1 - c2
        // Base cases
        if r1 >= n || c1 >= n || r2 >= n || c2 >= n || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
            return -1e8
        }
        if r1 == n-1 && c1 == n-1 {
            return grid[r1][c1]
        }
        if dp[r1][c1][c2] != -1 {
            return dp[r1][c1][c2]
        }
        cnt := grid[r1][c1]
        if r1 != r2 {
            cnt += grid[r2][c2]
        }
        a := solve(r1+1, c1, c2)
        b := solve(r1+1, c1, c2+1)
        c := solve(r1, c1+1, c2)
        d := solve(r1, c1+1, c2+1)
        dp[r1][c1][c2] = cnt + max(max(a, b), max(c, d))
        return dp[r1][c1][c2]
    }
    return max(0, solve(0, 0, 0))
}

func cherryPickup1(grid [][]int) int {
    n, dp, inf := len(grid), make([][]int, len(grid)), -1 << 31 -1
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    dp[0][0] = grid[0][0]
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for k := 1; k < n*2-1; k++ {
        for x1 := min(k, n-1); x1 >= max(k-n+1, 0); x1-- {
            for x2 := min(k, n-1); x2 >= x1; x2-- {
                y1, y2 := k-x1, k-x2
                if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
                    dp[x1][x2] = inf
                    continue
                }
                res := dp[x1][x2] // 都往右
                if x1 > 0 {
                    res = max(res, dp[x1-1][x2]) // 往下，往右
                }
                if x2 > 0 {
                    res = max(res, dp[x1][x2-1]) // 往右，往下
                }
                if x1 > 0 && x2 > 0 {
                    res = max(res, dp[x1-1][x2-1]) // 都往下
                }
                res += grid[x1][y1]
                if x2 != x1 { // 避免重复摘同一个樱桃
                    res += grid[x2][y2]
                }
                dp[x1][x2] = res
            }
        }
    }
    return max(dp[n-1][n-1], 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/14/grid.jpg" />
    // Input: grid = [[0,1,-1],[1,0,-1],[1,1,1]]
    // Output: 5
    // Explanation: The player started at (0, 0) and went down, down, right right to reach (2, 2).
    // 4 cherries were picked up during this single trip, and the matrix becomes [[0,1,-1],[0,0,-1],[0,0,0]].
    // Then, the player went left, up, up, left to return home, picking up one more cherry.
    // The total number of cherries picked up is 5, and this is the maximum possible.
    fmt.Println(cherryPickup([][]int{{0,1,-1},{1,0,-1},{1,1,1}})) // 5
    // Example 2:
    // Input: grid = [[1,1,-1],[1,-1,1],[-1,1,1]]
    // Output: 0
    fmt.Println(cherryPickup([][]int{{1,1,-1},{1,-1,1},{-1,1,1}})) // 0

    fmt.Println(cherryPickup1([][]int{{0,1,-1},{1,0,-1},{1,1,1}})) // 5
    fmt.Println(cherryPickup1([][]int{{1,1,-1},{1,-1,1},{-1,1,1}})) // 0
}