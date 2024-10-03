package main

// 2946. Matrix Similarity After Cyclic Shifts
// You are given an m x n integer matrix mat and an integer k. 
// The matrix rows are 0-indexed.

// The following proccess happens k times:
//     Even-indexed rows (0, 2, 4, ...) are cyclically shifted to the left.
//     <img src="https://assets.leetcode.com/uploads/2024/05/19/lshift.jpg" />

//     Odd-indexed rows (1, 3, 5, ...) are cyclically shifted to the right.
//     <img src="https://assets.leetcode.com/uploads/2024/05/19/rshift-stlone.jpg" />

// Return true if the final modified matrix after k steps is identical to the original matrix, and false otherwise.

// Example 1:
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 4
// Output: false
// Explanation:
// In each step left shift is applied to rows 0 and 2 (even indices), and right shift to row 1 (odd index).
// <img src="https://assets.leetcode.com/uploads/2024/05/19/t1-2.jpg" />

// Example 2:
// Input: mat = [[1,2,1,2],[5,5,5,5],[6,3,6,3]], k = 2
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/19/t1-3.jpg" />

// Example 3:
// Input: mat = [[2,2],[2,2]], k = 3
// Output: true
// Explanation:
// As all the values are equal in the matrix, even after performing cyclic shifts the matrix will remain the same.

// Constraints:
//     1 <= mat.length <= 25
//     1 <= mat[i].length <= 25
//     1 <= mat[i][j] <= 25
//     1 <= k <= 50

import "fmt"
import "slices"

func areSimilar(mat [][]int, k int) bool {
    m, n := len(mat), len(mat[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] != mat[i][ ( j + k) % n] {
                return false
            }
        }
    }    
    return true
}

func areSimilar1(mat [][]int, k int) bool {
    n := len(mat[0])
    k %= n
    if k == 0 {
        return true
    }
    for _, r := range mat {
        if !slices.Equal(r, append(r[k:], r[:k]...)){
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 4
    // Output: false
    // Explanation:
    // In each step left shift is applied to rows 0 and 2 (even indices), and right shift to row 1 (odd index).
    // <img src="https://assets.leetcode.com/uploads/2024/05/19/t1-2.jpg" />
    mat1 := [][]int{
        {1,2,3},
        {4,5,6},
        {7,8,9},
    }
    fmt.Println(areSimilar(mat1, 4)) // false
    // Example 2:
    // Input: mat = [[1,2,1,2],[5,5,5,5],[6,3,6,3]], k = 2
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/19/t1-3.jpg" />
    mat2 := [][]int{
        {1,2,1,2},
        {5,5,5,5},
        {6,3,6,3},
    }
    fmt.Println(areSimilar(mat2, 2)) // true
    // Example 3:
    // Input: mat = [[2,2],[2,2]], k = 3
    // Output: true
    // Explanation:
    // As all the values are equal in the matrix, even after performing cyclic shifts the matrix will remain the same.
    mat3 := [][]int{
        {2,2},
        {2,2},
    }
    fmt.Println(areSimilar(mat3, 3)) // true

    fmt.Println(areSimilar1(mat1, 4)) // false
    fmt.Println(areSimilar1(mat2, 2)) // true
    fmt.Println(areSimilar1(mat3, 3)) // true
}