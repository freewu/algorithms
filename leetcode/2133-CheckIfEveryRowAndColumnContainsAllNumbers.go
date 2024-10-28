package main

// 2133. Check if Every Row and Column Contains All Numbers
// An n x n matrix is valid if every row and every column contains all the integers from 1 to n (inclusive).

// Given an n x n integer matrix matrix, return true if the matrix is valid. Otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/21/example1drawio.png" />
// Input: matrix = [[1,2,3],[3,1,2],[2,3,1]]
// Output: true
// Explanation: In this case, n = 3, and every row and column contains the numbers 1, 2, and 3.
// Hence, we return true.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/21/example2drawio.png" />
// Input: matrix = [[1,1,1],[1,2,3],[1,2,3]]
// Output: false
// Explanation: In this case, n = 3, but the first row and the first column do not contain the numbers 2 or 3.
// Hence, we return false.
 
// Constraints:
//     n == matrix.length == matrix[i].length
//     1 <= n <= 100
//     1 <= matrix[i][j] <= n

import "fmt"

func checkValid(matrix [][]int) bool {
    n := len(matrix)
    for i := 0; i < n; i++ {
        arr1, arr2 := make([]int, n), make([]int, n)
        for j := 0; j < n;  j++ {
            arr1[matrix[i][j]-1]++
            arr2[matrix[j][i]-1]++
            if arr1[matrix[i][j]-1] > 1 || arr2[matrix[j][i]-1] > 1 { return false }
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/21/example1drawio.png" />
    // Input: matrix = [[1,2,3],[3,1,2],[2,3,1]]
    // Output: true
    // Explanation: In this case, n = 3, and every row and column contains the numbers 1, 2, and 3.
    // Hence, we return true.
    fmt.Println(checkValid([][]int{{1,2,3},{3,1,2},{2,3,1}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/21/example2drawio.png" />
    // Input: matrix = [[1,1,1],[1,2,3],[1,2,3]]
    // Output: false
    // Explanation: In this case, n = 3, but the first row and the first column do not contain the numbers 2 or 3.
    // Hence, we return false.
    fmt.Println(checkValid([][]int{{1,1,1},{1,2,3},{1,2,3}})) // false
}