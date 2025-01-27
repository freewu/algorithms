package main

// 面试题 01.08. Zero Matrix LCCI
// Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.

// Example 1:
// Input: 
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// Output: 
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]

// Example 2:
// Input: 
// [
//   [0,1,2,0],
//   [3,4,5,2],
//   [1,3,1,5]
// ]
// Output: 
// [
//   [0,0,0,0],
//   [0,4,5,0],
//   [0,3,1,0]
// ]

import "fmt"

func setZeroes(matrix [][]int)  {
    row, col := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
    for i, r := range matrix {
        for j, v := range r {
            if v == 0 { // 找到 0 的坐标
                row[i], col[j] = true, true
            }
        }
    }
    for i, r := range matrix {
        for j := range r {
            if row[i] || col[j] {
                r[j] = 0
            }
        }
    }
}

func main() {
    // Example 1:
    // Input: 
    // [
    //   [1,1,1],
    //   [1,0,1],
    //   [1,1,1]
    // ]
    // Output: 
    // [
    //   [1,0,1],
    //   [0,0,0],
    //   [1,0,1]
    // ]
    matrix1 := [][]int{
        []int{1,1,1},
        []int{1,0,1},
        []int{1,1,1},
    }
    fmt.Println("matrix1 before: ", matrix1) // [[1 1 1] [1 0 1] [1 1 1]]
    setZeroes(matrix1)
    fmt.Println("matrix1 after: ",  matrix1) // [[1 0 1] [0 0 0] [1 0 1]]
    // Example 2: 
    // Input: 
    // [
    //   [0,1,2,0],
    //   [3,4,5,2],
    //   [1,3,1,5]
    // ]
    // Output: 
    // [
    //   [0,0,0,0],
    //   [0,4,5,0],
    //   [1,3,1,5]
    // ]
    matrix2 := [][]int{
        []int{0,1,2,0},
        []int{3,4,5,2},
        []int{1,3,1,5},
    }
    fmt.Println("matrix2 before: ", matrix2) // [[0 1 2 0] [3 4 5 2] [1 3 1 5]]
    setZeroes(matrix2)
    fmt.Println("matrix2 after: ",  matrix2) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}