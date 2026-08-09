package main

// 4016. Maximum Area of Two Non-Overlapping Square Submatrices
// You are given a 2D integer matrix mat of size m × n, where:
//     1. mat[r][c] == 1 means the cell at row r and column c is usable.
//     2. mat[r][c] == 0 means it is not usable.

// Your task is to find two submatrices that satisfy the following conditions:
//     1. Both submatrices must be squares of the same side length k.
//     2. The two submatrices must not share any cell.
//     3. Each submatrix can only cover cells where mat[r][c] == 1.
    
// Return the maximum possible area of each of the two squares. 
// If it is not possible to choose two such squares, return 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2026/06/13/image.png" />
// Input: mat = [[1,1,1,0],[1,1,1,1],[0,0,1,1]]
// Output: 4
// Explanation:
// The largest equal non-overlapping squares have side length k = 2 with area 4.
// First square starts at top-left (0, 0) and covers cells (0, 0), (0, 1), (1, 0), and (1, 1).
// Second square starts at top-left (1, 2) and covers cells (1, 2), (1, 3), (2, 2), and (2, 3).
// Thus, the answer is 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2026/06/13/screenshot-2026-06-13-at-83728pm.png" />
// Input: mat = [[0,1],[1,0]]
// Output: 1
// Explanation:
// The largest equal non-overlapping squares have side length k = 1 with area 1.
// First square starts at top-left (0, 1) and covers cell (0, 1).
// Second square starts at top-left (1, 0) and covers cell (1, 0).
// Thus, the answer is 1.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2026/06/13/screenshot-2026-06-13-at-83751pm.png" />
// Input: mat = [[0,0],[0,1]]
// Output: 0
// Explanation:
// There is only one usable cell, so it is impossible to choose two non-overlapping squares. Thus, the answer is 0.

// Constraints:
//     mat.length == m
//     mat[i].length == n
//     1 <= m, n <= 500
//     mat[i][j] is either 0 or 1.

import "fmt"

func maxArea(mat [][]int) int {
    calc := func(mat [][]int) int {
        m, n := len(mat), len(mat[0])
        // 221. 最大正方形（空间优化写法）计算 mat 下半部分的最大正方形的边长
        sufMax := make([]int, m)
        f := make([]int, n+1)
        mx := 0
        for i := m - 1; i > 0; i-- {
            last := 0
            for j, x := range mat[i] {
                if x == 1 {
                    tmp := f[j+1]
                    f[j+1] = min(last, f[j+1], f[j]) + 1
                    last = tmp
                    mx = max(mx, f[j+1])
                } else {
                    f[j+1] = 0
                    last = 0
                }
            }
            sufMax[i] = mx
        }
        res := 0
        // 计算 mat 上半部分的最大正方形的边长
        preMax := 0
        clear(f)
        for i, row := range mat[:m-1] {
            last := 0
            for j, x := range row {
                if x == 1 {
                    tmp := f[j+1]
                    f[j+1] = min(last, f[j+1], f[j]) + 1
                    last = tmp
                    preMax = max(preMax, f[j+1])
                } else {
                    f[j+1] = 0
                    last = 0
                }
            }
            if sufMax[i+1] <= res {
                break // 最优性优化：继续循环不会让 res 变大
            }
            res = max(res, min(preMax, sufMax[i+1])) // 题目要求两个正方形的边长相等
        }
        return res * res
    }
    transpose := func(mat [][]int) [][]int { // 转置矩阵 mat
        m, n := len(mat), len(mat[0])
        arr := make([][]int, n)
        for i := range arr {
            arr[i] = make([]int, m)
            for j, row := range mat {
                arr[i][j] = row[i]
            }
        }
        return arr
    }
    return max(calc(mat), calc(transpose(mat)))
}

func maxArea1(mat [][]int) int {
    get := func(i, j int) int {
        if i < 0 || j < 0 {
            return 0
        }
        return mat[i][j]
    }
    for i, row := range mat {
        for j := range row {
            if mat[i][j] == 1 {
                mat[i][j] = 1 + min(get(i, j-1), min(get(i-1, j-1), get(i-1, j)))
            }
        }
    }
    feasible := func(k int) bool {
        if k == 0 {
            return false
        }
        minR, minC, found := 1 << 61, 1 << 61, false
        for i, row := range mat {
            for j := range row {
                if mat[i][j] >= k {
                    found = true
                    if r := i - k + 1; r < minR {
                        minR = r
                    }
                    if c := j - k + 1; c < minC {
                        minC = c
                    }
                }
            }
        }
        if !found {
            return false
        }
        for i, row := range mat {
            for j := range row {
                if mat[i][j] >= k {
                    r, c := i - k + 1, j-k + 1
                    if r >= minR + k || c >= minC + k {
                        return true
                    }
                }
            }
        }
        return false
    }
    high := min(len(mat), len(mat[0]))
    l, h, res := 1, high, 0
    for l <= h {
        mid := (l + h) / 2
        if feasible(mid) {
            res = mid
            l = mid + 1
        } else {
            h = mid - 1
        }
    }
    return res * res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2026/06/13/image.png" />
    // Input: mat = [[1,1,1,0],[1,1,1,1],[0,0,1,1]]
    // Output: 4
    // Explanation:
    // The largest equal non-overlapping squares have side length k = 2 with area 4.
    // First square starts at top-left (0, 0) and covers cells (0, 0), (0, 1), (1, 0), and (1, 1).
    // Second square starts at top-left (1, 2) and covers cells (1, 2), (1, 3), (2, 2), and (2, 3).
    // Thus, the answer is 4.
    fmt.Println(maxArea([][]int{{1,1,1,0},{1,1,1,1},{0,0,1,1}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2026/06/13/screenshot-2026-06-13-at-83728pm.png" />
    // Input: mat = [[0,1],[1,0]]
    // Output: 1
    // Explanation:
    // The largest equal non-overlapping squares have side length k = 1 with area 1.
    // First square starts at top-left (0, 1) and covers cell (0, 1).
    // Second square starts at top-left (1, 0) and covers cell (1, 0).
    // Thus, the answer is 1.
    fmt.Println(maxArea([][]int{{0,1},{1,0}})) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2026/06/13/screenshot-2026-06-13-at-83751pm.png" />
    // Input: mat = [[0,0],[0,1]]
    // Output: 0
    // Explanation:
    // There is only one usable cell, so it is impossible to choose two non-overlapping squares. Thus, the answer is 0.
    fmt.Println(maxArea([][]int{{0,0},{0,1}})) // 0

    fmt.Println(maxArea1([][]int{{1,1,1,0},{1,1,1,1},{0,0,1,1}})) // 4
    fmt.Println(maxArea1([][]int{{0,1},{1,0}})) // 1
    fmt.Println(maxArea1([][]int{{0,0},{0,1}})) // 0
}