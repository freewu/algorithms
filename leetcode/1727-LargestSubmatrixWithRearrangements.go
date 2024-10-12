package main

// 1727. Largest Submatrix With Rearrangements
// You are given a binary matrix matrix of size m x n, 
// and you are allowed to rearrange the columns of the matrix in any order.

// Return the area of the largest submatrix within matrix 
// where every element of the submatrix is 1 after reordering the columns optimally.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/29/screenshot-2020-12-30-at-40536-pm.png" />
// Input: matrix = [[0,0,1],[1,1,1],[1,0,1]]
// Output: 4
// Explanation: You can rearrange the columns as shown above.
// The largest submatrix of 1s, in bold, has an area of 4.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/29/screenshot-2020-12-30-at-40852-pm.png" />
// Input: matrix = [[1,0,1,0,1]]
// Output: 3
// Explanation: You can rearrange the columns as shown above.
// The largest submatrix of 1s, in bold, has an area of 3.

// Example 3:
// Input: matrix = [[1,1,0],[1,0,1]]
// Output: 2
// Explanation: Notice that you must rearrange entire columns, and there is no way to make a submatrix of 1s larger than an area of 2.

// Constraints:
//     m == matrix.length
//     n == matrix[i].length
//     1 <= m * n <= 10^5
//     matrix[i][j] is either 0 or 1.

import "fmt"
import "sort"

func largestSubmatrix(matrix [][]int) int {
    res, m, n := 0, len(matrix), len(matrix[0])
    height := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                height[j] = 0
            } else {
                height[j]++
            }
        }
        heightCopy := make([]int, n)
        copy(heightCopy, height)
        sort.Ints(heightCopy)
        for i := 0; i < n; i++ {
            res = max(res, heightCopy[i]*(n-i))
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/29/screenshot-2020-12-30-at-40536-pm.png" />
    // Input: matrix = [[0,0,1],[1,1,1],[1,0,1]]
    // Output: 4
    // Explanation: You can rearrange the columns as shown above.
    // The largest submatrix of 1s, in bold, has an area of 4.
    fmt.Println(largestSubmatrix([][]int{{0,0,1},{1,1,1},{1,0,1}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/29/screenshot-2020-12-30-at-40852-pm.png" />
    // Input: matrix = [[1,0,1,0,1]]
    // Output: 3
    // Explanation: You can rearrange the columns as shown above.
    // The largest submatrix of 1s, in bold, has an area of 3.
    fmt.Println(largestSubmatrix([][]int{{1,0,1,0,1}})) // 3
    // Example 3:
    // Input: matrix = [[1,1,0],[1,0,1]]
    // Output: 2
    // Explanation: Notice that you must rearrange entire columns, and there is no way to make a submatrix of 1s larger than an area of 2.
    fmt.Println(largestSubmatrix([][]int{{1,1,0},{1,0,1}})) // 2
}