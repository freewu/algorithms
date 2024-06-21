package main

// 2482. Difference Between Ones and Zeros in Row and Column
// You are given a 0-indexed m x n binary matrix grid.
// A 0-indexed m x n difference matrix diff is created with the following procedure:
//     Let the number of ones in the ith row be onesRowi.
//     Let the number of ones in the jth column be onesColj.
//     Let the number of zeros in the ith row be zerosRowi.
//     Let the number of zeros in the jth column be zerosColj.
//     diff[i][j] = onesRowi + onesColj - zerosRowi - zerosColj

// Return the difference matrix diff.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/06/image-20221106171729-5.png" / >
// Input: grid = [[0,1,1],[1,0,1],[0,0,1]]
// Output: [[0,0,4],[0,0,4],[-2,-2,2]]
// Explanation:
// - diff[0][0] = onesRow0 + onesCol0 - zerosRow0 - zerosCol0 = 2 + 1 - 1 - 2 = 0 
// - diff[0][1] = onesRow0 + onesCol1 - zerosRow0 - zerosCol1 = 2 + 1 - 1 - 2 = 0 
// - diff[0][2] = onesRow0 + onesCol2 - zerosRow0 - zerosCol2 = 2 + 3 - 1 - 0 = 4 
// - diff[1][0] = onesRow1 + onesCol0 - zerosRow1 - zerosCol0 = 2 + 1 - 1 - 2 = 0 
// - diff[1][1] = onesRow1 + onesCol1 - zerosRow1 - zerosCol1 = 2 + 1 - 1 - 2 = 0 
// - diff[1][2] = onesRow1 + onesCol2 - zerosRow1 - zerosCol2 = 2 + 3 - 1 - 0 = 4 
// - diff[2][0] = onesRow2 + onesCol0 - zerosRow2 - zerosCol0 = 1 + 1 - 2 - 2 = -2
// - diff[2][1] = onesRow2 + onesCol1 - zerosRow2 - zerosCol1 = 1 + 1 - 2 - 2 = -2
// - diff[2][2] = onesRow2 + onesCol2 - zerosRow2 - zerosCol2 = 1 + 3 - 2 - 0 = 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/06/image-20221106171747-6.png" / >
// Input: grid = [[1,1,1],[1,1,1]]
// Output: [[5,5,5],[5,5,5]]
// Explanation:
// - diff[0][0] = onesRow0 + onesCol0 - zerosRow0 - zerosCol0 = 3 + 2 - 0 - 0 = 5
// - diff[0][1] = onesRow0 + onesCol1 - zerosRow0 - zerosCol1 = 3 + 2 - 0 - 0 = 5
// - diff[0][2] = onesRow0 + onesCol2 - zerosRow0 - zerosCol2 = 3 + 2 - 0 - 0 = 5
// - diff[1][0] = onesRow1 + onesCol0 - zerosRow1 - zerosCol0 = 3 + 2 - 0 - 0 = 5
// - diff[1][1] = onesRow1 + onesCol1 - zerosRow1 - zerosCol1 = 3 + 2 - 0 - 0 = 5
// - diff[1][2] = onesRow1 + onesCol2 - zerosRow1 - zerosCol2 = 3 + 2 - 0 - 0 = 5

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 10^5
//     grid[i][j] is either 0 or 1.

import "fmt"

func onesMinusZeros(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    row, col := make([][]int, m), make([][]int, m)
    for i := range row { row[i] = make([]int, n) }
    for i := range col { col[i] = make([]int, n) }
    for i := 0; i < m; i++ { // 只需要统计 1 的数量就行, 余下就是0
        ones := 0
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 { ones++ }
        }
        for j := 0; j < n; j++ {
            row[i][j] = ones
        }
    }
    for j := 0; j < n; j++ {
        ones := 0
        for i := 0; i < m; i++ {
            if grid[i][j] == 1 { ones++ }
        }
        for i := 0; i < m; i++ {
            col[i][j] = ones
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            grid[i][j] = row[i][j] + col[i][j] - (n - row[i][j]) - (m - col[i][j])
        }
    }
    return grid
}

func onesMinusZeros1(grid [][]int) [][]int {
    m, n, r, c := len(grid), len(grid[0]), []int{}, []int{}
    for i := range grid{
        x := 0
        for j := range grid[i] {
            x += grid[i][j]
            if grid[i][j] == 0 {
                x -= 1
            }
        }
        r = append(r, x)
    }
    for i := 0; i < n; i++ {
        x := 0
        for j := 0; j <m; j++ {
            x += grid[j][i]
            if grid[j][i] == 0 {
                x -= 1
            }
        }
        c = append(c,x)
    }
    for i := range grid {
        for j := range grid[i] {
            grid[i][j] = r[i] + c[j]
        }
    }
    return grid
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/06/image-20221106171729-5.png" / >
    // Input: grid = [[0,1,1],[1,0,1],[0,0,1]]
    // Output: [[0,0,4],[0,0,4],[-2,-2,2]]
    // Explanation:
    // - diff[0][0] = onesRow0 + onesCol0 - zerosRow0 - zerosCol0 = 2 + 1 - 1 - 2 = 0 
    // - diff[0][1] = onesRow0 + onesCol1 - zerosRow0 - zerosCol1 = 2 + 1 - 1 - 2 = 0 
    // - diff[0][2] = onesRow0 + onesCol2 - zerosRow0 - zerosCol2 = 2 + 3 - 1 - 0 = 4 
    // - diff[1][0] = onesRow1 + onesCol0 - zerosRow1 - zerosCol0 = 2 + 1 - 1 - 2 = 0 
    // - diff[1][1] = onesRow1 + onesCol1 - zerosRow1 - zerosCol1 = 2 + 1 - 1 - 2 = 0 
    // - diff[1][2] = onesRow1 + onesCol2 - zerosRow1 - zerosCol2 = 2 + 3 - 1 - 0 = 4 
    // - diff[2][0] = onesRow2 + onesCol0 - zerosRow2 - zerosCol0 = 1 + 1 - 2 - 2 = -2
    // - diff[2][1] = onesRow2 + onesCol1 - zerosRow2 - zerosCol1 = 1 + 1 - 2 - 2 = -2
    // - diff[2][2] = onesRow2 + onesCol2 - zerosRow2 - zerosCol2 = 1 + 3 - 2 - 0 = 2
    grid1 := [][]int{
        {0,1,1},
        {1,0,1},
        {0,0,1},
    }
    fmt.Println(onesMinusZeros(grid1)) // [[0,0,4],[0,0,4],[-2,-2,2]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/06/image-20221106171747-6.png" / >
    // Input: grid = [[1,1,1],[1,1,1]]
    // Output: [[5,5,5],[5,5,5]]
    // Explanation:
    // - diff[0][0] = onesRow0 + onesCol0 - zerosRow0 - zerosCol0 = 3 + 2 - 0 - 0 = 5
    // - diff[0][1] = onesRow0 + onesCol1 - zerosRow0 - zerosCol1 = 3 + 2 - 0 - 0 = 5
    // - diff[0][2] = onesRow0 + onesCol2 - zerosRow0 - zerosCol2 = 3 + 2 - 0 - 0 = 5
    // - diff[1][0] = onesRow1 + onesCol0 - zerosRow1 - zerosCol0 = 3 + 2 - 0 - 0 = 5
    // - diff[1][1] = onesRow1 + onesCol1 - zerosRow1 - zerosCol1 = 3 + 2 - 0 - 0 = 5
    // - diff[1][2] = onesRow1 + onesCol2 - zerosRow1 - zerosCol2 = 3 + 2 - 0 - 0 = 5
    grid2 := [][]int{
        {1,1,1},
        {1,1,1},
    }
    fmt.Println(onesMinusZeros(grid2)) // [[5,5,5],[5,5,5]]

    grid11 := [][]int{
        {0,1,1},
        {1,0,1},
        {0,0,1},
    }
    fmt.Println(onesMinusZeros1(grid11)) // [[0,0,4],[0,0,4],[-2,-2,2]]
    grid12 := [][]int{
        {1,1,1},
        {1,1,1},
    }
    fmt.Println(onesMinusZeros1(grid12)) // [[5,5,5],[5,5,5]]
}