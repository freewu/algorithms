package main

// 3359. Find Sorted Submatrices With Maximum Element at Most K
// You are given a 2D matrix grid of size m x n. You are also given a non-negative integer k.

// Return the number of submatrices of grid that satisfy the following conditions:
//     The maximum element in the submatrix less than or equal to k.
//     Each row in the submatrix is sorted in non-increasing order.

// A submatrix (x1, y1, x2, y2) is a matrix that forms by choosing all cells grid[x][y] where x1 <= x <= x2 and y1 <= y <= y2.

// Example 1:
// Input: grid = [[4,3,2,1],[8,7,6,1]], k = 3
// Output: 8
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/01/mine.png" />
// The 8 submatrices are:
// [[1]]
// [[1]]
// [[2,1]]
// [[3,2,1]]
// [[1],[1]]
// [[2]]
// [[3]]
// [[3,2]]

// Example 2:
// Input: grid = [[1,1,1],[1,1,1],[1,1,1]], k = 1
// Output: 36
// Explanation:
// There are 36 submatrices of grid. All submatrices have their maximum element equal to 1.

// Example 3:
// Input: grid = [[1]], k = 1
// Output: 1

// Constraints:
//     1 <= m == grid.length <= 10^3
//     1 <= n == grid[i].length <= 10^3
//     1 <= grid[i][j] <= 10^9
//     1 <= k <= 10^9

import "fmt"

func countSubmatrices(grid [][]int, k int) int64 {
    res, m, n := 0, len(grid), len(grid[0])
    widths := make([]int, m)
    for j := 0; j < n; j++ {
        curr := 0
        stack := make([][2]int, 0) // Deque
        for i := 0; i < m; i++ {
            if grid[i][j] > k {
                widths[i] = 0
            } else if j > 0 && grid[i][j] > grid[i][j - 1] {
                widths[i] = 1
            } else {
                widths[i]++
            }
            width, height := widths[i],  1
            for len(stack) != 0 && stack[len(stack) - 1][0] >= width {
                prev := stack[len(stack)-1] 
                stack = stack[:len(stack)-1] // pop
                curr -= (prev[0] - width) * prev[1]
                height += prev[1]
            }
            curr += width
            res += curr
            stack = append(stack, [2]int{width, height})
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: grid = [[4,3,2,1],[8,7,6,1]], k = 3
    // Output: 8
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/01/mine.png" />
    // The 8 submatrices are:
    // [[1]]
    // [[1]]
    // [[2,1]]
    // [[3,2,1]]
    // [[1],[1]]
    // [[2]]
    // [[3]]
    // [[3,2]]
    fmt.Println(countSubmatrices([][]int{{4,3,2,1},{8,7,6,1}}, 3)) // 8
    // Example 2:
    // Input: grid = [[1,1,1],[1,1,1],[1,1,1]], k = 1
    // Output: 36
    // Explanation:
    // There are 36 submatrices of grid. All submatrices have their maximum element equal to 1.
    fmt.Println(countSubmatrices([][]int{{1,1,1},{1,1,1},{1,1,1}}, 1)) // 36
    // Example 3:
    // Input: grid = [[1]], k = 1
    // Output: 1
    fmt.Println(countSubmatrices([][]int{{1}}, 1)) // 1
}