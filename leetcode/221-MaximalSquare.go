package main

// 221. Maximal Square
// Given an m x n binary matrix filled with 0's and 1's, 
// find the largest square containing only 1's and return its area.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg" />
// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 4

// Example 2:
// Input: matrix = [["0","1"],["1","0"]]
// Output: 1

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/11/26/max2grid.jpg" />
// Input: matrix = [["0"]]
// Output: 0
 
// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 300
//     matrix[i][j] is '0' or '1'.

import "fmt"

// func maximalSquare(matrix [][]byte) int {
//     res, heights := 0, make([]int, len(matrix[0]) + 1)
//     heights[len(heights)-1] = -1
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for _, row := range matrix {
//         for i := range row {
//             if row[i] == '1' {
//                 heights[i]++
//             } else {
//                 heights[i] = 0
//             }
//         }
//         stack := []int{}
//         for i, currentHeight := range heights {
//             for len(stack) > 0 && heights[stack[len(stack)-1]] > currentHeight {
//                 prev := heights[stack[len(stack)-1]]
//                 stack = stack[:len(stack)-1]
//                 width := i
//                 if len(stack) > 0 {
//                     width = i - stack[len(stack)-1] - 1
//                 }
//                 //fmt.Println(prev)
//                 if prev == width {
//                     res = max(res, prev * width)
//                 }
//             }
//             stack = append(stack, i)
//         } 
//     }
//     return res
// }

// dp
func maximalSquare(matrix [][]byte) int {
    res, m, n := 0, len(matrix), len(matrix[0])
    dp := make([][]int,m)
    for i := 0; i < m; i++ {
        dp[i]=make([]int,n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j]=='1'{
                res = 1
                dp[i][j]=1
             } 
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if dp[i][j] == 1 {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
                if res < dp[i][j] {
                    res = dp[i][j]
                }
            }
        }
    }
    return res * res 
}

func maximalSquare1(matrix [][]byte) int {
    res, rows, cols := 0, len(matrix), len(matrix[0])
    dp := make([][]int, rows + 1)
    for i, _ := range dp {
        dp[i] = make([]int, cols + 1)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == '1' {
                t := dp[i][j]
                if dp[i+1][j] < t {
                    t = dp[i+1][j]
                }
                if dp[i][j+1] < t {
                    t = dp[i][j+1]
                }
                dp[i+1][j+1] = t + 1
                if t + 1 > res {
                    res = t + 1
                }
            }
        }
    }
    return res * res
}

func main() {
    fmt.Println(maximalSquare([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}})) // 4
    fmt.Println(maximalSquare([][]byte{{'1','0'},{'0','1'}})) // 1
    fmt.Println(maximalSquare([][]byte{{'0'}})) // 0

    fmt.Println(maximalSquare1([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}})) // 4
    fmt.Println(maximalSquare1([][]byte{{'1','0'},{'0','1'}})) // 1
    fmt.Println(maximalSquare1([][]byte{{'0'}})) // 0
}