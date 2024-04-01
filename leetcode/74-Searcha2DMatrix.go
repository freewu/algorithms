package main

// 74. Search a 2D Matrix
// You are given an m x n integer matrix matrix with the following two properties:
//     Each row is sorted in non-decreasing order.
//     The first integer of each row is greater than the last integer of the previous row.

// Given an integer target, return true if target is in matrix or false otherwise.
// You must write a solution in O(log(m * n)) time complexity.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/05/mat.jpg" />
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg" />
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false
 
// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m, n <= 100
//     -10^4 <= matrix[i][j], target <= 10^4

import "fmt"

// 二分
// 给出一个二维矩阵，矩阵的特点是随着矩阵的下标增大而增大。要求设计一个算法能在这个矩阵中高效的找到一个数，如果找到就输出 true，找不到就输出 false。
// 虽然是一个二维矩阵，但是由于它特殊的有序性，所以完全可以按照下标把它看成一个一维矩阵，只不过需要行列坐标转换。最后利用二分搜索直接搜索即可
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    m, low, high := len(matrix[0]), 0, len(matrix[0]) * len(matrix) - 1
    for low <= high {
        mid := low + (high-low) >> 1
        if matrix[mid / m][mid % m] == target {
            return true
        } else if matrix[mid / m][mid % m] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return false
}

// 遍历大法 O(n*m)
func searchMatrix1(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    colum := len(matrix[0]) - 1 // 列数
    for row := 0; row < len(matrix); row++ {
        // 找到大于目标的那一行开始查找
        if matrix[row][colum] >= target {
            // 循环每列
            for i := 0; i <= colum; i++ {
                if matrix[row][i] == target {
                    return true
                }
            }
            return false
        }
    }
    return false
}

func main() {
    matrix := [][]int{[]int{1, 3, 5, 7},[]int{10, 11, 16, 20},[]int{23, 30, 34, 50}}
    fmt.Printf("matrix = %v\n",matrix)
    fmt.Printf("searchMatrix(matrix,3) = %v\n",searchMatrix(matrix,3))  // true
    fmt.Printf("searchMatrix(matrix,2) = %v\n",searchMatrix(matrix,13)) // false

    fmt.Printf("searchMatrix1(matrix,3) = %v\n",searchMatrix1(matrix,3))  // true
    fmt.Printf("searchMatrix1(matrix,2) = %v\n",searchMatrix1(matrix,13)) // false
}
