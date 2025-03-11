package main

// 2174. Remove All Ones With Row and Column Flips II
// You are given a 0-indexed m x n binary matrix grid.
// In one operation, you can choose any i and j that meet the following conditions:
//     0 <= i < m
//     0 <= j < n
//     grid[i][j] == 1

// and change the values of all cells in row i and column j to zero.
// Return the minimum number of operations needed to remove all 1's from grid.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/13/image-20220213162716-1.png" />
// Input: grid = [[1,1,1],[1,1,1],[0,1,0]]
// Output: 2
// Explanation:
// In the first operation, change all cell values of row 1 and column 1 to zero.
// In the second operation, change all cell values of row 0 and column 0 to zero.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/02/13/image-20220213162737-2.png" />
// Input: grid = [[0,1,0],[1,0,1],[0,1,0]]
// Output: 2
// Explanation:
// In the first operation, change all cell values of row 1 and column 0 to zero.
// In the second operation, change all cell values of row 2 and column 1 to zero.
// Note that we cannot perform an operation using row 1 and column 1 because grid[1][1] != 1.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/02/13/image-20220213162752-3.png" />
// Input: grid = [[0,0],[0,0]]
// Output: 0
// Explanation:
// There are no 1's to remove so return 0.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 15
//     1 <= m * n <= 15
//     grid[i][j] is either 0 or 1.

import "fmt"

// 状态压缩 + 位运算
func removeOnes(grid [][]int) int {
    points := [][]int{}
    m, n, mask := len(grid), len(grid[0]), 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                points = append(points, []int{i, j})
                mask += 1 << (i*n + j)
            }
        }
    }
    bitLen := func (i int)int{
        res := 0
        for i > 0 {
            res++
            i &= i - 1
        }
        return res
    }
    res, total, all := len(points), 1 << len(points), (1 << (m*n))-1
    for i := 1; i < total; i++ {
        temp, count := mask, bitLen(i)
        if count > res { continue }
        for j := 0; j < len(points); j++ {
            if (1<<j)&i > 0 {
                for k := 0; k < m; k++ {
                    temp &= all - (1 << (k*n + points[j][1]))
                }
                for k := 0; k < n; k++ {
                    temp &= all - (1 << (points[j][0]*n + k))
                }
            }
        }
        if temp == 0 {
            res = count
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/13/image-20220213162716-1.png" />
    // Input: grid = [[1,1,1],[1,1,1],[0,1,0]]
    // Output: 2
    // Explanation:
    // In the first operation, change all cell values of row 1 and column 1 to zero.
    // In the second operation, change all cell values of row 0 and column 0 to zero.
    grid1 := [][]int{
        {1,1,1},
        {1,1,1},
        {0,1,0},
    }
    fmt.Println(removeOnes(grid1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/13/image-20220213162737-2.png" />
    // Input: grid = [[0,1,0],[1,0,1],[0,1,0]]
    // Output: 2
    // Explanation:
    // In the first operation, change all cell values of row 1 and column 0 to zero.
    // In the second operation, change all cell values of row 2 and column 1 to zero.
    // Note that we cannot perform an operation using row 1 and column 1 because grid[1][1] != 1.
    grid2 := [][]int{
        {0,1,0},
        {1,0,1},
        {0,1,0},
    }
    fmt.Println(removeOnes(grid2)) // 2
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/02/13/image-20220213162752-3.png" />
    // Input: grid = [[0,0],[0,0]]
    // Output: 0
    // Explanation:
    // There are no 1's to remove so return 0.
    grid3 := [][]int{
        {0,0},
        {0,0},
    }
    fmt.Println(removeOnes(grid3)) // 0
}