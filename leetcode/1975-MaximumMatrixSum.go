package main

// 1975. Maximum Matrix Sum
// You are given an n x n integer matrix. 
// You can do the following operation any number of times:
//     Choose any two adjacent elements of matrix and multiply each of them by -1.

// Two elements are considered adjacent if and only if they share a border.

// Your goal is to maximize the summation of the matrix's elements. 
// Return the maximum sum of the matrix's elements using the operation mentioned above.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/16/pc79-q2ex1.png" />
// Input: matrix = [[1,-1],[-1,1]]
// Output: 4
// Explanation: We can follow the following steps to reach sum equals 4:
// - Multiply the 2 elements in the first row by -1.
// - Multiply the 2 elements in the first column by -1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/16/pc79-q2ex2.png" />
// Input: matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
// Output: 16
// Explanation: We can follow the following step to reach sum equals 16:
// - Multiply the 2 last elements in the second row by -1.

// Constraints:
//     n == matrix.length == matrix[i].length
//     2 <= n <= 250
//     -10^5 <= matrix[i][j] <= 10^5

import "fmt"

func maxMatrixSum(matrix [][]int) int64 {
    sum, flag, mn := 0, false, 100001
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for _, row := range matrix {
        for _, cell := range row {
            if cell < 0 { flag = !flag }
            cell := abs(cell)
            sum += cell
            if mn > cell {
                mn = cell
            }
        }
    }
    if flag { sum -= mn * 2 }
    return int64(sum)
}

func maxMatrixSum1(matrix [][]int) int64 {
    sum, count, mn  := 0, 0, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, row := range matrix {
        for _, col := range row {
            if col < 0 {
                count++
                col = -col
            }
            sum += col
            mn = min(mn, col)
        }
    }
    if count & 1 == 0 {
        return int64(sum)
    }
    return int64(sum - mn << 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/16/pc79-q2ex1.png" />
    // Input: matrix = [[1,-1],[-1,1]]
    // Output: 4
    // Explanation: We can follow the following steps to reach sum equals 4:
    // - Multiply the 2 elements in the first row by -1.
    // - Multiply the 2 elements in the first column by -1.
    fmt.Println(maxMatrixSum([][]int{{1,-1},{-1,1}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/07/16/pc79-q2ex2.png" />
    // Input: matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
    // Output: 16
    // Explanation: We can follow the following step to reach sum equals 16:
    // - Multiply the 2 last elements in the second row by -1.
    fmt.Println(maxMatrixSum([][]int{{1,2,3},{-1,-2,-3},{1,2,3}})) // 16

    fmt.Println(maxMatrixSum1([][]int{{1,-1},{-1,1}})) // 4
    fmt.Println(maxMatrixSum1([][]int{{1,2,3},{-1,-2,-3},{1,2,3}})) // 16
}