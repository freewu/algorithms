package main

// 1351. Count Negative Numbers in a Sorted Matrix
// Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, 
// return the number of negative numbers in grid.

// Example 1:
// Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
// Output: 8
// Explanation: There are 8 negatives number in the matrix.

// Example 2:
// Input: grid = [[3,2],[1,0]]
// Output: 0

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 100
//     -100 <= grid[i][j] <= 100

import "fmt"

// O(n * n)
func countNegatives(grid [][]int) int {
    res := 0
    for i:= 0; i < len(grid); i++ {
        for j:= 0; j < len(grid[i]); j++ {
            if grid[i][j] < 0 {
                res++
            }
        }
    }
    return res
}

// 二分
func countNegatives1(grid [][]int) int {
    n,m := len(grid), len(grid[0])
    l, r := 0 , m - 1
    for l <= r {
        mid := l + ((r - l) >> 1)
        if grid[0][mid] >= 0 {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    res := m - l
    for i := 1; i < n; i++ {
        for r > 0 && grid[i][r] < 0 {
            r--
        }
        if grid[i][r] >= 0 {
            res += m - r - 1
        } else {
            res += m - r
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
    // Output: 8
    // Explanation: There are 8 negatives number in the matrix.
    fmt.Println(countNegatives([][]int{ {4,3,2,-1}, {3,2,1,-1}, {1,1,-1,-2}, {-1,-1,-2,-3}  })) // 8
    // Example 2:
    // Input: grid = [[3,2],[1,0]]
    // Output: 0
    fmt.Println(countNegatives([][]int{ {3,2}, {1,0}})) // 0

    fmt.Println(countNegatives1([][]int{ {4,3,2,-1}, {3,2,1,-1}, {1,1,-1,-2}, {-1,-1,-2,-3}  })) // 8
    fmt.Println(countNegatives1([][]int{ {3,2}, {1,0}})) // 0
}