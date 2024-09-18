package main

// 2132. Stamping the Grid
// You are given an m x n binary matrix grid where each cell is either 0 (empty) or 1 (occupied).

// You are then given stamps of size stampHeight x stampWidth. 
// We want to fit the stamps such that they follow the given restrictions and requirements:
//     1. Cover all the empty cells.
//     2. Do not cover any of the occupied cells.
//     3. We can put as many stamps as we want.
//     4. Stamps can overlap with each other.
//     5. Stamps are not allowed to be rotated.
//     6. Stamps must stay completely inside the grid.

// Return true if it is possible to fit the stamps while following the given restrictions and requirements. 
// Otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/03/ex1.png" />
// Input: grid = [[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]], stampHeight = 4, stampWidth = 3
// Output: true
// Explanation: We have two overlapping stamps (labeled 1 and 2 in the image) that are able to cover all the empty cells.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/03/ex2.png" />
// Input: grid = [[1,0,0,0],[0,1,0,0],[0,0,1,0],[0,0,0,1]], stampHeight = 2, stampWidth = 2 
// Output: false 
// Explanation: There is no way to fit the stamps onto all the empty cells without the stamps going outside the grid.

// Constraints:
//     m == grid.length
//     n == grid[r].length
//     1 <= m, n <= 10^5
//     1 <= m * n <= 2 * 10^5
//     grid[r][c] is either 0 or 1.
//     1 <= stampHeight, stampWidth <= 10^5

import "fmt"

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
    height, width, FILLED := len(grid), len(grid[0]), 1
    pre := make([]int, width)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := range pre { pre[i] = 1 } // fill 1
    for r := 0; r < height; r++ {
        for c := 0; c < width; c++ {
            if grid[r][c] != FILLED {
                if pre[c] == 1 {
                    grid[r][c] = -1
                } else {
                    grid[r][c] = max(pre[c] - 1, -stampHeight)
                }
            }
        }
        pre = grid[r]
    }
    fill := func(r, c, height, width int) {
        for i := c - width + 1; i <= c; i++ {
            j := r
            for j >= r - height + 1 && grid[j][i] != FILLED {
                grid[j][i] = FILLED
                j--
            }
        }
    }
    // scan each height, and once we find consecutive `grid` of `-stampHeight` of length `stampWidth` or longer,
    // replace the `grid` value with `FILLED`
    for r := stampHeight - 1; r < height; r++ {
        count := 0
        for c := 0; c < width; c++ {
            if grid[r][c] == -stampHeight { 
                count++ 
            } else {
                count = 0
            }
            if count == stampWidth { fill(r, c, stampHeight, stampWidth) }
            if count > stampWidth { fill(r, c, stampHeight, 1) }
            // width set to 1 for performance; the previous width 0 - cnt would have been already filled
        }
    }
    // check if there is remaining cell of value that is not `FILLED`
    for _, ints := range grid {
        for _, v := range ints {
            if v != FILLED { return false }
        }
    }
    return true
}

func possibleToStamp1(grid [][]int, stampHeight, stampWidth int) bool {
    m, n := len(grid), len(grid[0])
    // 1. 计算 grid 的二维前缀和
    s := make([][]int, m+1)
    s[0] = make([]int, n+1)
    for i, row := range grid {
        s[i+1] = make([]int, n+1)
        for j, v := range row {
            s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + v
        }
    }
    // 2. 计算二维差分
    // 为方便第 3 步的计算，在 d 数组的最上面和最左边各加了一行（列），所以下标要 +1
    d := make([][]int, m+2)
    for i := range d {
        d[i] = make([]int, n+2)
    }
    for i2 := stampHeight; i2 <= m; i2++ {
        for j2 := stampWidth; j2 <= n; j2++ {
            i1 := i2 - stampHeight + 1
            j1 := j2 - stampWidth + 1
            if s[i2][j2]-s[i2][j1-1]-s[i1-1][j2]+s[i1-1][j1-1] == 0 {
                d[i1][j1]++
                d[i1][j2+1]--
                d[i2+1][j1]--
                d[i2+1][j2+1]++
            }
        }
    }
    // 3. 还原二维差分矩阵对应的计数矩阵（原地计算）
    for i, row := range grid {
        for j, v := range row {
            d[i+1][j+1] += d[i+1][j] + d[i][j+1] - d[i][j]
            if v == 0 && d[i+1][j+1] == 0 {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/03/ex1.png" />
    // Input: grid = [[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]], stampHeight = 4, stampWidth = 3
    // Output: true
    // Explanation: We have two overlapping stamps (labeled 1 and 2 in the image) that are able to cover all the empty cells.
    grid1 := [][]int{
        {1,0,0,0},
        {1,0,0,0},
        {1,0,0,0},
        {1,0,0,0},
        {1,0,0,0},
    }
    fmt.Println(possibleToStamp(grid1, 4, 3)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/03/ex2.png" />
    // Input: grid = [[1,0,0,0],[0,1,0,0],[0,0,1,0],[0,0,0,1]], stampHeight = 2, stampWidth = 2 
    // Output: false 
    // Explanation: There is no way to fit the stamps onto all the empty cells without the stamps going outside the grid.
    grid2 := [][]int{
        {1,0,0,0},
        {0,1,0,0},
        {0,0,1,0},
        {0,0,0,1},
    }
    fmt.Println(possibleToStamp(grid2, 2, 2)) // false

    grid11 := [][]int{
        {1,0,0,0},
        {1,0,0,0},
        {1,0,0,0},
        {1,0,0,0},
        {1,0,0,0},
    }
    fmt.Println(possibleToStamp1(grid11, 4, 3)) // true
    grid12 := [][]int{
        {1,0,0,0},
        {0,1,0,0},
        {0,0,1,0},
        {0,0,0,1},
    }
    fmt.Println(possibleToStamp1(grid12, 2, 2)) // false
}