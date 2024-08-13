package main

// 1504. Count Submatrices With All Ones
// Given an m x n binary matrix mat, return the number of submatrices that have all ones.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/27/ones1-grid.jpg" />
// Input: mat = [[1,0,1],[1,1,0],[1,1,0]]
// Output: 13
// Explanation: 
// There are 6 rectangles of side 1x1.
// There are 2 rectangles of side 1x2.
// There are 3 rectangles of side 2x1.
// There is 1 rectangle of side 2x2. 
// There is 1 rectangle of side 3x1.
// Total number of rectangles = 6 + 2 + 3 + 1 + 1 = 13.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/27/ones2-grid.jpg" />
// Input: mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]
// Output: 24
// Explanation: 
// There are 8 rectangles of side 1x1.
// There are 5 rectangles of side 1x2.
// There are 2 rectangles of side 1x3. 
// There are 4 rectangles of side 2x1.
// There are 2 rectangles of side 2x2. 
// There are 2 rectangles of side 3x1. 
// There is 1 rectangle of side 3x2. 
// Total number of rectangles = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24.

// Constraints:
//     1 <= m, n <= 150
//     mat[i][j] is either 0 or 1.

import "fmt"

func numSubmat(mat [][]int) int {
    rows, cols := len(mat), len(mat[0])
    res, dp, inf := 0, make([][]int, rows), 1 << 31
    for i := 0; i < rows; i++ {
        dp[i] = make([]int, cols)
        count := 0
        for j := cols - 1; j >= 0; j-- {
            if mat[i][j] == 1 {
                count++
            } else {
                count = 0
            }
            dp[i][j] = count
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            mn := inf
            for k := i; k < rows; k++ {
                // get the min of consecutive 1s
                mn = min(mn, dp[k][j])
                res += mn
            }
        }
    }
    return res
}

// stack
func numSubmat1(mat [][]int) int {
    res, n := 0, len(mat[0])
    s, sum := make([]int, n+1), make([]int, n+1)
    for _, row := range mat {
        stack := []int{ 0 }
        for j, c := range row {
            sum[j+1] = sum[j+1] * c + c
            for sum[j+1] < sum[stack[len(stack)-1]] {
                stack = stack[:len(stack) - 1] // pop
            }
            k := stack[len(stack)-1] // peak
            s[j+1] = (j+1-k) * sum[j+1] + s[k]
            res += s[j+1]
            stack = append(stack, j+1) // push
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/27/ones1-grid.jpg" />
    // Input: mat = [[1,0,1],[1,1,0],[1,1,0]]
    // Output: 13
    // Explanation: 
    // There are 6 rectangles of side 1x1.
    // There are 2 rectangles of side 1x2.
    // There are 3 rectangles of side 2x1.
    // There is 1 rectangle of side 2x2. 
    // There is 1 rectangle of side 3x1.
    // Total number of rectangles = 6 + 2 + 3 + 1 + 1 = 13.
    mat1 := [][]int{
        {1,0,1},
        {1,1,0},
        {1,1,0},
    }
    fmt.Println(numSubmat(mat1)) // 13
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/27/ones2-grid.jpg" />
    // Input: mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]
    // Output: 24
    // Explanation: 
    // There are 8 rectangles of side 1x1.
    // There are 5 rectangles of side 1x2.
    // There are 2 rectangles of side 1x3. 
    // There are 4 rectangles of side 2x1.
    // There are 2 rectangles of side 2x2. 
    // There are 2 rectangles of side 3x1. 
    // There is 1 rectangle of side 3x2. 
    // Total number of rectangles = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24.
    mat2 := [][]int{
        {0,1,1,0},
        {0,1,1,1},
        {1,1,1,0},
    }
    fmt.Println(numSubmat(mat2)) // 24

    fmt.Println(numSubmat1(mat1)) // 13
    fmt.Println(numSubmat1(mat2)) // 24
}