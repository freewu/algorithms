package main

// 3933. Largest Local Values in a Matrix II
// You are given an n x m integer matrix matrix containing non-negative integers.

// A non-zero cell (row, col) checks the cells near it as follows:
//     1. Let x = matrix[row][col].
//     2. Consider every cell within x rows and x columns of (row, col).
//     3. Ignore cells that are outside the matrix.
//     4. Ignore the cells where both the row distance and column distance are exactly x.

// The cell (row, col) is a local maximum if it is non-zero and no considered cell has a value greater than x.

// Return an integer denoting the number of local maximums in matrix.

// ​​​​​​​Example 1:
// Input: matrix = [[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,2,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0]]
// Output: 1
// ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/13/chatgpt-image-may-14-2026-01_53_19-am.png" />
// Explanation:
// For the non-zero cell (3, 3), x = matrix[3][3] = 2.
// The highlighted cells are the considered cells within x rows and x columns of (3, 3).
// The four cells with both row and column distances equal to x = 2 are ignored.
// No considered cell has a value greater than 2, so (3, 3) is a local maximum.
// There are no other non-zero cells, so the answer is 1.

// Example 2:
// Input: matrix = [[1,2],[3,4]]
// Output: 1
// Explanation:
// Only the cell with value 4 is a local maximum. Every other non-zero cell considers a cell with a greater value.

// Example 3:
// Input: matrix = [[1,0,1],[0,1,0],[1,0,1]]
// Output: 5
// Explanation:
// For a cell with value 1, the considered cells are the cell itself and its 4-directionally adjacent cells that are inside the matrix.
// Each of the five cells with value 1 only considers cells with values 0 or 1, so all five of them are local maximums.

// Example 4:
// Input: matrix = [[1,1],[1,1]]
// Output: 4
// Explanation:
// All cells have the same value. Therefore, no cell considers another cell with a greater value, so all 4 cells are local maximums.

// Constraints:
//     1 <= n == matrix.length <= 200
//     1 <= m == matrix[i].length <= 200
//     0 <= matrix[i][j] <= 200

import "fmt"
import "math/bits"

func countLocalMaximums(matrix [][]int) int {
    res, n, m := 0, len(matrix), len(matrix[0])
    wn, wm := bits.Len(uint(n)), bits.Len(uint(m))
    // dp[i][j][k1][k2] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
    dp := make([][][8][8]int, n)
    for i := range dp {
        dp[i] = make([][8][8]int, m)
    }
    for i, row := range matrix {
        for j, x := range row {
            dp[i][j][0][0] = x // 初始值
        }
    }
    // 单独计算 k1 = 0
    for k2 := 1; k2 < wm; k2++ {
        for i := range n {
            for j := range m - 1<<k2 + 1 {
                dp[i][j][0][k2] = max(dp[i][j][0][k2-1], dp[i][j+1<<(k2-1)][0][k2-1])
            }
        }
    }
    for k1 := 1; k1 < wn; k1++ {
        for k2 := range wm {
            for i := range n - 1 << k1 + 1 {
                for j := range m - 1 << k2 + 1 {
                    dp[i][j][k1][k2] = max(dp[i][j][k1-1][k2], dp[i+1<<(k1-1)][j][k1-1][k2])
                }
            }
        }
    }
    // 返回子矩阵最大值
    // 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
    query := func(r1, c1, r2, c2 int) int {
        r1, c1, r2, c2 = max(r1, 0), max(c1, 0), min(r2, n), min(c2, m)
        k1 := bits.Len8(uint8(r2-r1)) - 1
        k2 := bits.Len8(uint8(c2-c1)) - 1
        // 视作四个子矩阵的并集
        return max(dp[r1][c1][k1][k2], dp[r2-1<<k1][c1][k1][k2], dp[r1][c2-1<<k2][k1][k2], dp[r2-1<<k1][c2-1<<k2][k1][k2])
    }
    for i, row := range matrix {
        for j, x := range row {
            if x > 0 && max(query(i-x, j-x+1, i+x+1, j+x), query(i-x+1, j-x, i+x, j+x+1)) <= x {
                res++
            }
        }
    }
    return res
}

func countLocalMaximums1(matrix [][]int) int {
    type Pair struct {
        r, c int
    }
    res, row, col := 0, len(matrix), len(matrix[0])
    pos := make([][]Pair, 201)
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            val := matrix[i][j]
            pos[val] = append(pos[val], Pair{i, j}) 
        }
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    check := func(r, c, val int) bool {
        for next := val + 1; next <= 200; next++ {
            for _, p := range pos[next] {
                rd, cd := abs(p.r - r), abs(p.c - c)
                if rd <= val && cd <= val {
                    if rd == val && cd == val {
                        continue
                    }
                    return false
                }
            }
        }
        return true
    }
    for val := 1; val <= 200; val++ {
        for _, v := range pos[val] {
            if check(v.r, v.c, val) {
                res++
            }
        }
    }
    return res
}

func main() {
    // ​​​​​​​Example 1:
    // Input: matrix = [[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,2,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0]]
    // Output: 1
    // ​​​​​​​​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/13/chatgpt-image-may-14-2026-01_53_19-am.png" />
    // Explanation:
    // For the non-zero cell (3, 3), x = matrix[3][3] = 2.
    // The highlighted cells are the considered cells within x rows and x columns of (3, 3).
    // The four cells with both row and column distances equal to x = 2 are ignored.
    // No considered cell has a value greater than 2, so (3, 3) is a local maximum.
    // There are no other non-zero cells, so the answer is 1.
    fmt.Println(countLocalMaximums([][]int{
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,2,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
    })) // 1
    // Example 2:
    // Input: matrix = [[1,2],[3,4]]
    // Output: 1
    // Explanation:
    // Only the cell with value 4 is a local maximum. Every other non-zero cell considers a cell with a greater value.
    fmt.Println(countLocalMaximums([][]int{ {1,2},{3,4} })) // 1
    // Example 3:
    // Input: matrix = [[1,0,1],[0,1,0],[1,0,1]]
    // Output: 5
    // Explanation:
    // For a cell with value 1, the considered cells are the cell itself and its 4-directionally adjacent cells that are inside the matrix.
    // Each of the five cells with value 1 only considers cells with values 0 or 1, so all five of them are local maximums.
    fmt.Println(countLocalMaximums([][]int{ {1,0,1},{0,1,0},{1,0,1} })) // 5
    // Example 4:
    // Input: matrix = [[1,1],[1,1]]
    // Output: 4
    // Explanation:
    // All cells have the same value. Therefore, no cell considers another cell with a greater value, so all 4 cells are local maximums.
    fmt.Println(countLocalMaximums([][]int{ {1,1},{1,1} })) // 4

    fmt.Println(countLocalMaximums1([][]int{
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,2,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0},
    })) // 1
    fmt.Println(countLocalMaximums1([][]int{ {1,2},{3,4} })) // 1
    fmt.Println(countLocalMaximums1([][]int{ {1,0,1},{0,1,0},{1,0,1} })) // 5
    fmt.Println(countLocalMaximums1([][]int{ {1,1},{1,1} })) // 4
}
