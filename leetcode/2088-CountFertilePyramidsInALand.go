package main

// 2088. Count Fertile Pyramids in a Land
// A farmer has a rectangular grid of land with m rows and n columns that can be divided into unit cells. 
// Each cell is either fertile (represented by a 1) or barren (represented by a 0). 
// All cells outside the grid are considered barren.

// A pyramidal plot of land can be defined as a set of cells with the following criteria:
//     1. The number of cells in the set has to be greater than 1 and all cells must be fertile.
//     2. The apex of a pyramid is the topmost cell of the pyramid. 
//        The height of a pyramid is the number of rows it covers. Let (r, c) be the apex of the pyramid, and its height be h. 
//        Then, the plot comprises of cells (i, j) where r <= i <= r + h - 1 and c - (i - r) <= j <= c + (i - r).

// An inverse pyramidal plot of land can be defined as a set of cells with similar criteria:
//     1. The number of cells in the set has to be greater than 1 and all cells must be fertile.
//     2. The apex of an inverse pyramid is the bottommost cell of the inverse pyramid. 
//        The height of an inverse pyramid is the number of rows it covers. Let (r, c) be the apex of the pyramid, and its height be h. 
//        Then, the plot comprises of cells (i, j) where r - h + 1 <= i <= r and c - (r - i) <= j <= c + (r - i).

// Some examples of valid and invalid pyramidal (and inverse pyramidal) plots are shown below. 
// Black cells indicate fertile cells.
// <img src="https://assets.leetcode.com/uploads/2021/11/08/image.png" />

// Given a 0-indexed m x n binary matrix grid representing the farmland, 
// return the total number of pyramidal and inverse pyramidal plots that can be found in grid.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/22/1.JPG" />
// Input: grid = [[0,1,1,0],[1,1,1,1]]
// Output: 2
// Explanation: The 2 possible pyramidal plots are shown in blue and red respectively.
// There are no inverse pyramidal plots in this grid. 
// Hence total number of pyramidal and inverse pyramidal plots is 2 + 0 = 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/22/2.JPG" />
// Input: grid = [[1,1,1],[1,1,1]]
// Output: 2
// Explanation: The pyramidal plot is shown in blue, and the inverse pyramidal plot is shown in red. 
// Hence the total number of plots is 1 + 1 = 2.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/22/3.JPG" />
// Input: grid = [[1,1,1,1,0],[1,1,1,1,1],[1,1,1,1,1],[0,1,0,0,1]]
// Output: 13
// Explanation: There are 7 pyramidal plots, 3 of which are shown in the 2nd and 3rd figures.
// There are 6 inverse pyramidal plots, 2 of which are shown in the last figure.
// The total number of plots is 7 + 6 = 13.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     grid[i][j] is either 0 or 1.

import "fmt"

func countPyramids(grid [][]int) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    count := func(grid [][]int) int{
        n, m, res := len(grid), len(grid[0]), 0
        for i := 1; i < n; i++ {
            for j := 1; j < m - 1; j++ {
                if grid[i][j] > 0 && grid[i-1][j] > 0 {
                    grid[i][j] = min(grid[i-1][j-1], grid[i-1][j+1]) + 1
                    res += grid[i][j]-1
                }
            }
        }
        for i := 1;i < n; i++ {
            for j := 1; j < m-1; j++ {
                if grid[i][j] > 0 {
                    grid[i][j]=1
                }
            }
        }
        return res
    }
    res := count(grid)
    for i,j := 0, len(grid) - 1; i < j; i,j = i+1, j-1 {
        grid[i],grid[j]=grid[j],grid[i]
    }
    res += count(grid)
    return res
}

func countPyramids1(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    count := func() {
        dp[m-1] = grid[m-1]
        for i := m - 2; i >= 0; i-- {
            dp[i][0] = grid[i][0]
            dp[i][n-1] = grid[i][n-1]
            for j := 1; j < n-1; j++ {
                if grid[i][j] == 0 {
                    dp[i][j] = 0
                } else {
                    dp[i][j] = min(min(dp[i+1][j-1], dp[i+1][j+1]), dp[i+1][j]) + 1
                    res += dp[i][j] - 1
                }
            }
        }
    }
    count() // 求正金字塔个数
    for i := 0; i < m/2; i++ { // 上下颠倒
        grid[i], grid[m-1-i] = grid[m-1-i], grid[i]
    }
    count() // 再求一遍正金字塔个数
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/22/1.JPG" />
    // Input: grid = [[0,1,1,0],[1,1,1,1]]
    // Output: 2
    // Explanation: The 2 possible pyramidal plots are shown in blue and red respectively.
    // There are no inverse pyramidal plots in this grid. 
    // Hence total number of pyramidal and inverse pyramidal plots is 2 + 0 = 2.
    gird1 := [][]int{
        {0,1,1,0},
        {1,1,1,1},
    }
    fmt.Println(countPyramids(gird1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/22/2.JPG" />
    // Input: grid = [[1,1,1],[1,1,1]]
    // Output: 2
    // Explanation: The pyramidal plot is shown in blue, and the inverse pyramidal plot is shown in red. 
    // Hence the total number of plots is 1 + 1 = 2.
    gird2 := [][]int{
        {1,1,1},
        {1,1,1},
    }
    fmt.Println(countPyramids(gird2)) // 2
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/22/3.JPG" />
    // Input: grid = [[1,1,1,1,0],[1,1,1,1,1],[1,1,1,1,1],[0,1,0,0,1]]
    // Output: 13
    // Explanation: There are 7 pyramidal plots, 3 of which are shown in the 2nd and 3rd figures.
    // There are 6 inverse pyramidal plots, 2 of which are shown in the last figure.
    // The total number of plots is 7 + 6 = 13.
    gird3 := [][]int{
        {1,1,1,1,0},
        {1,1,1,1,1},
        {1,1,1,1,1},
        {0,1,0,0,1},
    }
    fmt.Println(countPyramids(gird3)) // 13

    gird11 := [][]int{
        {0,1,1,0},
        {1,1,1,1},
    }
    fmt.Println(countPyramids1(gird11)) // 2
    gird12 := [][]int{
        {1,1,1},
        {1,1,1},
    }
    fmt.Println(countPyramids1(gird12)) // 2
    gird13 := [][]int{
        {1,1,1,1,0},
        {1,1,1,1,1},
        {1,1,1,1,1},
        {0,1,0,0,1},
    }
    fmt.Println(countPyramids1(gird13)) // 13
}