package main

// 867. Transpose Matrix
// Given a 2D integer array matrix, return the transpose of matrix.
// The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
// <img src="https://assets.leetcode.com/uploads/2021/02/10/hint_transpose.png" / >

// [1,2,3]      [1,4,7]
// [4,5,6]  =>  [2,5,8]
// [7,8,9]      [3,6,9]

// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[1,4,7],[2,5,8],[3,6,9]]

// Example 2:
// Input: matrix = [[1,2,3],[4,5,6]]
// Output: [[1,4],[2,5],[3,6]]
 
// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 1000
//     1 <= m * n <= 10^5
//     -10^9 <= matrix[i][j] <= 10^9

import "fmt"

func transpose(matrix [][]int) [][]int {
    // 造一个全0的结果
    res := make([][]int, len(matrix[0]))
    for i := 0; i < len(matrix[0]); i++ {
        t := make([]int,len(matrix))
        res[i] = t 
    }
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            res[j][i] = matrix[i][j]
        }
    }
    return res
}

func transpose1(matrix [][]int) [][]int {
    n,m := len(matrix),len(matrix[0])
    t := make([][]int,m)
    for i := range t{
        t[i] = make([]int,n)
        for j := range t[i] {
            t[i][j] = -1
        }
    }
    for i,row := range matrix{
        for j,v := range row{
            t[j][i] =v
        }
    }
    return t
}

func main() {
    fmt.Println(
        transpose(
            [][]int {
                []int{ 1,2,3 },
                []int{ 4,5,6 },
                []int{ 7,8,9 },
            },
        ),
    ) // [[1,4,7],[2,5,8],[3,6,9]]

    fmt.Println(
        transpose(
            [][]int {
                []int{ 1,2,3 },
                []int{ 4,5,6 },
            },
        ),
    ) // [[1,4],[2,5],[3,6]]

    fmt.Println(
        transpose1(
            [][]int {
                []int{ 1,2,3 },
                []int{ 4,5,6 },
                []int{ 7,8,9 },
            },
        ),
    ) // [[1,4,7],[2,5,8],[3,6,9]]

    fmt.Println(
        transpose1(
            [][]int {
                []int{ 1,2,3 },
                []int{ 4,5,6 },
            },
        ),
    ) // [[1,4],[2,5],[3,6]]
}