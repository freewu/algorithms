package main

// 面试题 10.09. Sorted Matrix Search LCCI
// Given an M x N matrix in which each row and each column is sorted in ascending order, write a method to find an element.

// Example:
// Given matrix:
// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
// Given target = 5, return true.
// Given target = 20, return false.

import "fmt"
import "sort"

// brute force
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {  return false }
    for _, row := range matrix {
        for _, v := range row {
            if v == target {
                return true
            }
        }
    }
    return false
}

// 二分
func searchMatrix1(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    for _, row := range matrix {
        i := sort.SearchInts(row, target)
        if i < len(row) && row[i] == target {
            return true
        }
    }
    return false
}

// Z 字形查找
func searchMatrix2(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    if m == 0 || n == 0 { return false }
    x, y := 0, n - 1
    for x < m && y >= 0 {
        if matrix[x][y] == target { return true }
        if matrix[x][y] > target {
            y--
        } else {
            x++
        }
    }
    return false
}

func main() {
    // Example:
    // Given matrix:
    // [
    //   [1,   4,  7, 11, 15],
    //   [2,   5,  8, 12, 19],
    //   [3,   6,  9, 16, 22],
    //   [10, 13, 14, 17, 24],
    //   [18, 21, 23, 26, 30]
    // ]
    // Given target = 5, return true.
    // Given target = 20, return false.
    grid1 := [][]int{
        {1,   4,  7, 11, 15},
        {2,   5,  8, 12, 19},
        {3,   6,  9, 16, 22},
        {10, 13, 14, 17, 24},
        {18, 21, 23, 26, 30},
    }
    fmt.Println(searchMatrix(grid1, 5)) // true
    fmt.Println(searchMatrix(grid1, 20)) // false

    fmt.Println(searchMatrix1(grid1, 5)) // true
    fmt.Println(searchMatrix1(grid1, 20)) // false

    fmt.Println(searchMatrix2(grid1, 5)) // true
    fmt.Println(searchMatrix2(grid1, 20)) // false
}