package main

// 3276. Select Cells in Grid With Maximum Score
// You are given a 2D matrix grid consisting of positive integers.

// You have to select one or more cells from the matrix such that the following conditions are satisfied:
//     1. No two selected cells are in the same row of the matrix.
//     2. The values in the set of selected cells are unique.

// Your score will be the sum of the values of the selected cells.

// Return the maximum score you can achieve.

// Example 1:
// Input: grid = [[1,2,3],[4,3,2],[1,1,1]]
// Output: 8
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/29/grid1drawio.png" />
// We can select the cells with values 1, 3, and 4 that are colored above.

// Example 2:
// Input: grid = [[8,7,6],[8,3,2]]
// Output: 15
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/29/grid8_8drawio.png" />
// We can select the cells with values 7 and 8 that are colored above.

// Constraints:
//     1 <= grid.length, grid[i].length <= 10
//     1 <= grid[i][j] <= 100

import "fmt"

func maxScore(grid [][]int) int {
    n, mx := len(grid), 0
    g := [101][11]bool{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, row := range grid {
        for _, v := range row {
            g[v][i] = true
            mx = max(mx, v)
        }
    }
    f := make([][]int, mx + 1)
    for i := range f {
        f[i] = make([]int, 1 << n)
    }
    for i := 1; i <= mx; i++ {
        for j := 0; j < 1 << n; j++ {
            f[i][j] = f[i-1][j]
            for k := 0; k < n; k++ {
                if g[i][k] && (j >> k & 1) == 1 {
                    f[i][j] = max(f[i][j], f[i-1][j ^ 1 << k] + i)
                }
            }
        }
    }
    return f[mx][1 << n - 1]
}

func maxScore1(grid [][]int) int {
    pos := make(map[int]int)
    for i, row := range grid {
        for _, x := range row {
            pos[x] |= 1 << i // 保存所有包含 x 的行号（压缩成二进制数）
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    f := make([]int, 1 << len(grid))
    for x, mask := range pos {
        for j := range f {
            for t, lb := mask, 0; t > 0; t ^= lb {
                lb = t & -t    // lb = 1 << k，其中 k 是行号
                if j & lb == 0 { // 没选过第 k 行的数
                    f[j] = max(f[j], f[j|lb] + x)
                }
            }
        }
    }
    return f[0]
}

func main() {
    // Example 1:
    // Input: grid = [[1,2,3],[4,3,2],[1,1,1]]
    // Output: 8
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/29/grid1drawio.png" />
    // We can select the cells with values 1, 3, and 4 that are colored above.
    fmt.Println(maxScore([][]int{{1,2,3}, {4,3,2}, {1,1,1}})) // 8
    // Example 2:
    // Input: grid = [[8,7,6],[8,3,2]]
    // Output: 15
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/29/grid8_8drawio.png" />
    // We can select the cells with values 7 and 8 that are colored above.
    fmt.Println(maxScore([][]int{{8,7,6}, {8,3,2}})) // 15

    fmt.Println(maxScore1([][]int{{1,2,3}, {4,3,2}, {1,1,1}})) // 8
    fmt.Println(maxScore1([][]int{{8,7,6}, {8,3,2}})) // 15
}