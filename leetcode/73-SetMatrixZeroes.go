package main

// 73. Set Matrix Zeroes
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
// You must do it in place.

// Constraints:
//     m == matrix.length
//     n == matrix[0].length
//     1 <= m, n <= 200
//     -2^31 <= matrix[i][j] <= 2^31 - 1

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg" />
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg" />
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Follow up:
//     A straight forward solution using O(mn) space is probably a bad idea.
//     A simple improvement uses O(m + n) space, but still not the best solution.

// 解题思路:
//     给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0


import "fmt"

func setZeroes(matrix [][]int) {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    isFirstRowExistZero, isFirstColExistZero := false, false
    // 判断第1列中是否有 0 存在
    for i := 0; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            isFirstColExistZero = true
            break
        }
    }
    // 判断第1行中是否有 0 存在
    for j := 0; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            isFirstRowExistZero = true
            break
        }
    }
    // 从第2行第2列开始 循环判断在中间的数是否存在 0，有则把 第1列 & 第1行 设置为 0
    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    // 处理[1:]行全部置 0
    for i := 1; i < len(matrix); i++ {
        if matrix[i][0] == 0 {
            for j := 1; j < len(matrix[0]); j++ {
                matrix[i][j] = 0
            }
        }
    }
    // 处理[1:]列全部置 0
    for j := 1; j < len(matrix[0]); j++ {
        if matrix[0][j] == 0 {
            for i := 1; i < len(matrix); i++ {
                matrix[i][j] = 0
            }
        }
    }
    if isFirstRowExistZero {
        for j := 0; j < len(matrix[0]); j++ {
            matrix[0][j] = 0
        }
    }
    if isFirstColExistZero {
        for i := 0; i < len(matrix); i++ {
            matrix[i][0] = 0
        }
    }
}

// best solution
func setZeroes1(matrix [][]int)  {
    m,n := len(matrix), len(matrix[0])
    rows := make([]int, m)
    columns := make([]int, n)
    for i:=0; i < m; i++ {
        for j:= 0; j < n; j++ {
            if matrix[i][j] == 0 {
                rows[i] = 1
                columns[j] = 1
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j<n; j++ {
            if rows[i] == 1 {
                matrix[i][j] = 0
            }
        }
    }
    for j:=0; j < n; j++ {
        for i := 0; i < m; i++ {
            if columns[j] == 1 {
                matrix[i][j] = 0
            }
        }
    }
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg" />
    // Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
    // Output: [[1,0,1],[0,0,0],[1,0,1]]
    matrix1 := [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}
    fmt.Printf("before matrix1 = %v\n",matrix1)
    for _,m := range matrix1 {
        fmt.Printf("%v\n",m)
    }
    setZeroes(matrix1)
    fmt.Printf("after matrix1 = %v\n",matrix1)
    for _,m := range matrix1 {
        fmt.Printf("%v\n",m)
    }
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg" />
    // Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
    // Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
    matrix2 := [][]int{[]int{0,1,2,0}, []int{3,4,5,2}, []int{1,3,1,5}}
    fmt.Printf("before matrix2 = %v\n",matrix2)
    for _,m := range matrix2 {
        fmt.Printf("%v\n",m)
    }
    setZeroes(matrix2)
    fmt.Printf("after matrix2 = %v\n",matrix2)
    for _,m := range matrix2 {
        fmt.Printf("%v\n",m)
    }

    matrix11 := [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}
    fmt.Printf("before matrix11 = %v\n",matrix1)
    for _,m := range matrix11 {
        fmt.Printf("%v\n",m)
    }
    setZeroes1(matrix11)
    fmt.Printf("after matrix11 = %v\n",matrix1)
    for _,m := range matrix11 {
        fmt.Printf("%v\n",m)
    }

    matrix12 := [][]int{[]int{0,1,2,0}, []int{3,4,5,2}, []int{1,3,1,5}}
    fmt.Printf("before matrix12 = %v\n",matrix2)
    for _,m := range matrix12 {
        fmt.Printf("%v\n",m)
    }
    setZeroes1(matrix12)
    fmt.Printf("after matrix12 = %v\n",matrix2)
    for _,m := range matrix12 {
        fmt.Printf("%v\n",m)
    }
}
