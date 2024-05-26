package main

// 311. Sparse Matrix Multiplication
// Given two sparse matrices mat1 of size m x k and mat2 of size k x n, return the result of mat1 x mat2. 
// You may assume that multiplication is always possible.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/12/mult-grid.jpg" />
// Input: mat1 = [[1,0,0],[-1,0,3]], mat2 = [[7,0,0],[0,0,0],[0,0,1]]
// Output: [[7,0,0],[-7,0,3]]

// Example 2:
// Input: mat1 = [[0]], mat2 = [[0]]
// Output: [[0]]
 
// Constraints:
//     m == mat1.length
//     k == mat1[i].length == mat2.length
//     n == mat2[i].length
//     1 <= m, n, k <= 100
//     -100 <= mat1[i][j], mat2[i][j] <= 100

import "fmt"

// 暴力
func multiply(mat1 [][]int, mat2 [][]int) [][]int {
    m,n,l := len(mat1),len(mat2[0]), len(mat1[0])
    res := make([][]int, m)
    for i:= 0; i < m; i++ {
        res[i] = make([]int, n)
    }
    for i := 0; i < m; i++ { // 三层循环进行计算
        for j := 0; j < n; j++ {
            for k := 0; k < l; k++ {
                res[i][j] += mat1[i][k] * mat2[k][j]
            }
        }
    }
    return res  
}

// 运用乘法左边矩阵的稀疏性
func multiply1(mat1 [][]int, mat2 [][]int) [][]int {
    m, n, l := len(mat1), len(mat2[0]), len(mat1[0])
    res := make([][]int, m)
    for i:= 0; i < m; i++ {
        res[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for k := 0; k < l; k++ {
            // mat1[i][k] 若为 0， mat1[i][k] * mat2[k][j]也为 0，故可以跳过，减少循环次数
            if mat1[i][k] == 0 { continue  }
            for j := 0; j < n; j++ {
                res[i][j] += mat1[i][k] * mat2[k][j]
            }
        } 
    }
    return res
}

// 先对两个矩阵做一个处理（时间复杂度O(mn)，只保留非0数据），之后根据矩阵乘法的性质进行计算
func multiply2(mat1 [][]int, mat2 [][]int) [][]int {
    m,n:= len(mat1),len(mat2[0])
    res := make([][]int,m)
    for i:= 0; i < m; i++ {
        res[i] = make([]int, n)
    }
    getNoneZeroMat := func (matrix [][]int) [][]int{
        m, n := len(matrix),len(matrix[0])
        res := [][]int{}
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                if matrix[i][j] != 0 {
                    res = append(res, []int{i,j,matrix[i][j]})
                }
            }
        }
        return res
    }
    noneZeroA := getNoneZeroMat(mat1)
    noneZeroB := getNoneZeroMat(mat2)
    for _, m1 := range noneZeroA{  
        for _, m2 := range noneZeroB{ // 每个 m1 包含的数据：i,j,matrix[i][j]
            if m1[1] == m2[0] { // 这里这么判断的原因：mat1和mat2相乘，只有mat1的数据的列数和mat2的行数相等才会进行计算
                res[m1[0]][m2[1]] += m1[2]*m2[2]
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/mult-grid.jpg" />
    // Input: mat1 = [[1,0,0],[-1,0,3]], mat2 = [[7,0,0],[0,0,0],[0,0,1]]
    // Output: [[7,0,0],[-7,0,3]]
    fmt.Println(multiply([][]int{{1,0,0},{-1,0,3}},[][]int{{7,0,0},{0,0,0},{0,0,1}})) // [[7,0,0],[-7,0,3]]
    // Example 2:
    // Input: mat1 = [[0]], mat2 = [[0]]
    // Output: [[0]]
    fmt.Println(multiply([][]int{{0}},[][]int{{0}})) // [[0]]

    fmt.Println(multiply1([][]int{{1,0,0},{-1,0,3}},[][]int{{7,0,0},{0,0,0},{0,0,1}})) // [[7,0,0],[-7,0,3]]
    fmt.Println(multiply1([][]int{{0}},[][]int{{0}})) // [[0]]

    fmt.Println(multiply2([][]int{{1,0,0},{-1,0,3}},[][]int{{7,0,0},{0,0,0},{0,0,1}})) // [[7,0,0],[-7,0,3]]
    fmt.Println(multiply2([][]int{{0}},[][]int{{0}})) // [[0]]
}