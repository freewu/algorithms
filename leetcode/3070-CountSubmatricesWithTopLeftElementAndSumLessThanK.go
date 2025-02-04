package main

// 3070. Count Submatrices with Top-Left Element and Sum Less Than k
// You are given a 0-indexed integer matrix grid and an integer k.

// Return the number of submatrices that contain the top-left element of the grid, and have a sum less than or equal to k.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/01/01/example1.png" />
// Input: grid = [[7,6,3],[6,6,1]], k = 18
// Output: 4
// Explanation: There are only 4 submatrices, shown in the image above, that contain the top-left element of grid, and have a sum less than or equal to 18.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/01/01/example21.png" />
// Input: grid = [[7,2,9],[1,5,0],[2,6,6]], k = 20
// Output: 6
// Explanation: There are only 6 submatrices, shown in the image above, that contain the top-left element of grid, and have a sum less than or equal to 20.

// Constraints:
//     m == grid.length 
//     n == grid[i].length
//     1 <= n, m <= 1000 
//     0 <= grid[i][j] <= 1000
//     1 <= k <= 10^9

import "fmt"

func countSubmatrices(grid [][]int, k int) int {
    res, n, m := 0, len(grid), len(grid[0])
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if i > 0 {
                grid[i][j] += grid[i - 1][j]
            }
            if j > 0 {
                grid[i][j] += grid[i][j - 1]
            }
            if i > 0 && j > 0 {
                grid[i][j] -= grid[i - 1][j - 1]
            }
            if grid[i][j] <= k {
                res++
            }
        }
    }
    return res
}

func countSubmatrices1(grid [][]int, k int) int {
    res, n := 0, len(grid[0])
    sum := make([]int,n + 1)
    for i := range grid {
        t := 0
        for j := range grid[i] {
            t += grid[i][j]
            sum[j + 1] += t
            if sum[j + 1] > k {
                continue
            } else if sum[j + 1] <= k {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/01/01/example1.png" />
    // Input: grid = [[7,6,3],[6,6,1]], k = 18
    // Output: 4
    // Explanation: There are only 4 submatrices, shown in the image above, that contain the top-left element of grid, and have a sum less than or equal to 18.
    fmt.Println(countSubmatrices([][]int{{7,6,3},{6,6,1}}, 18)) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/01/01/example21.png" />
    // Input: grid = [[7,2,9],[1,5,0],[2,6,6]], k = 20
    // Output: 6
    // Explanation: There are only 6 submatrices, shown in the image above, that contain the top-left element of grid, and have a sum less than or equal to 20.
    fmt.Println(countSubmatrices([][]int{{7,2,9},{1,5,0},{2,6,6}}, 20)) // 6

    fmt.Println(countSubmatrices1([][]int{{7,6,3},{6,6,1}}, 18)) // 4
    fmt.Println(countSubmatrices1([][]int{{7,2,9},{1,5,0},{2,6,6}}, 20)) // 6
}