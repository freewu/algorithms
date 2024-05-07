package main

// 1463. Cherry Pickup II
// You are given a rows x cols matrix grid representing a field of cherries where grid[i][j] represents the number of cherries that you can collect from the (i, j) cell.
// You have two robots that can collect cherries for you:
//     Robot #1 is located at the top-left corner (0, 0), and
//     Robot #2 is located at the top-right corner (0, cols - 1).

// Return the maximum number of cherries collection using both robots by following the rules below:
//     From a cell (i, j), robots can move to cell (i + 1, j - 1), (i + 1, j), or (i + 1, j + 1).
//     When any robot passes through a cell, It picks up all cherries, and the cell becomes an empty cell.
//     When both robots stay in the same cell, only one takes the cherries.
//     Both robots cannot move outside of the grid at any moment.
//     Both robots should reach the bottom row in grid.
 

// Example 1:
// <img src = "https://assets.leetcode.com/uploads/2020/04/23/sample_1_1802.png" />
// Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]
// Output: 24
// Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
// Cherries taken by Robot #1, (3 + 2 + 5 + 2) = 12.
// Cherries taken by Robot #2, (1 + 5 + 5 + 1) = 12.
// Total of cherries: 12 + 12 = 24.

// Example 2:
// <img src = "https://assets.leetcode.com/uploads/2020/04/23/sample_2_1802.png" />
// Input: grid = [[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]
// Output: 28
// Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
// Cherries taken by Robot #1, (1 + 9 + 5 + 2) = 17.
// Cherries taken by Robot #2, (1 + 3 + 4 + 3) = 11.
// Total of cherries: 17 + 11 = 28.

// Example 3:
// Input: grid = [[1,0,0,3],[0,0,0,3],[0,0,3,3],[9,0,3,3]]
// Output: 22

// Example 4:
// Input: grid = [[1,1],[1,1]]
// Output: 4

// Constraints:
//     rows == grid.length
//     cols == grid[i].length
//     2 <= rows, cols <= 70
//     0 <= grid[i][j] <= 100

import "fmt"

func cherryPickup(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    // 定义 dp[i][j][k] 代表第一个机器人从 (0,0) 走到 (i,k) 坐标，第二个机器人从 (0,n-1) 走到 (i,k) 坐标，两者最多能收集樱桃的数目
    dp := make([][][]int, rows)
    for i := 0; i < rows; i++ {
        dp[i] = make([][]int, cols)
        for j := 0; j < cols; j++ {
            dp[i][j] = make([]int, cols)
        }
    }
    isInBoard := func (dp [][][]int, i, j, k int) int {
        if i < 0 || j < 0 || j >= len(dp[0]) || k < 0 || k >= len(dp[0]) {
            return 0
        }
        return dp[i][j][k]
    }
    for i := 0; i < rows; i++ {
        for j := 0; j <= i && j < cols; j++ {
            for k := cols - 1; k >= cols-1-i && k >= 0; k-- {
                max := 0
                for a := j - 1; a <= j+1; a++ {
                    for b := k - 1; b <= k+1; b++ {
                        sum := isInBoard(dp, i-1, a, b)
                        if a == b && i > 0 && a >= 0 && a < cols {
                            sum -= grid[i-1][a]
                        }
                        if sum > max {
                            max = sum
                        }
                    }
                }
                if j == k {
                    max += grid[i][j]
                } else {
                    max += grid[i][j] + grid[i][k]
                }
                dp[i][j][k] = max
            }
        }
    }
    // 边界条件 dp[i][0][n-1] = grid[0][0] + grid[0][n-1]，最终答案存储在 dp[m-1] 行中，循环找出 dp[m-1][j1][j2] 中的最大值
    count := 0
    for j := 0; j < cols && j < rows; j++ {
        for k := cols - 1; k >= 0 && k >= cols-rows; k-- {
            if dp[rows-1][j][k] > count {
                count = dp[rows-1][j][k]
            }
        }
    }
    return count
}

func cherryPickup1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][][]int, m+1)
    for i := range dp {
        dp[i] = make([][]int, n+2)
        for j := range dp[i] {
            dp[i][j] = make([]int, n+2)
        }
    }
    max := func(arr ...int) int {
        mx := 0
        for i,v := range arr {
            if i == 0 {
                mx = v 
            } else if (v > mx) {
                mx = v
            }
        }
        return mx
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := m-1; i >= 0; i-- {
        for j := 0; j < min(n, i+1); j++ {
            for k := j+1; k < n; k++ {
                dp[i][j+1][k+1] = max(dp[i+1][j][k], dp[i+1][j][k+1], dp[i+1][j][k+2], dp[i+1][j+1][k], dp[i+1][j+1][k+1], dp[i+1][j+1][k+2],dp[i+1][j+2][k],dp[i+1][j+2][k+1], dp[i+1][j+2][k+2]) + grid[i][j] + grid[i][k]
            }
        }
    }
    return dp[0][1][n]
}

func main() {
    fmt.Println(cherryPickup([][]int{[]int{3,1,1},[]int{2,5,1},[]int{1,5,5},[]int{2,1,1}})) // 24
    fmt.Println(cherryPickup([][]int{[]int{1,0,0,0,0,0,1},[]int{2,0,0,0,0,3,0},[]int{2,0,9,0,0,0,0},[]int{0,3,0,5,4,0,0},[]int{1,0,2,3,0,0,6}})) // 28
    fmt.Println(cherryPickup([][]int{[]int{1,0,0,3},[]int{0,0,0,3},[]int{0,0,3,3},[]int{9,0,3,3}})) // 22
    fmt.Println(cherryPickup([][]int{[]int{1,1},[]int{1,1}})) // 4

    fmt.Println(cherryPickup1([][]int{[]int{3,1,1},[]int{2,5,1},[]int{1,5,5},[]int{2,1,1}})) // 24
    fmt.Println(cherryPickup1([][]int{[]int{1,0,0,0,0,0,1},[]int{2,0,0,0,0,3,0},[]int{2,0,9,0,0,0,0},[]int{0,3,0,5,4,0,0},[]int{1,0,2,3,0,0,6}})) // 28
    fmt.Println(cherryPickup1([][]int{[]int{1,0,0,3},[]int{0,0,0,3},[]int{0,0,3,3},[]int{9,0,3,3}})) // 22
    fmt.Println(cherryPickup1([][]int{[]int{1,1},[]int{1,1}})) // 4
}