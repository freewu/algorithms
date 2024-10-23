package main

// 3240. Minimum Number of Flips to Make Binary Grid Palindromic II
// You are given an m x n binary matrix grid.

// A row or column is considered palindromic if its values read the same forward and backward.

// You can flip any number of cells in grid from 0 to 1, or from 1 to 0.

// Return the minimum number of cells that need to be flipped to make all rows and columns palindromic, 
// and the total number of 1's in grid divisible by 4.

// Example 1:
// Input: grid = [[1,0,0],[0,1,0],[0,0,1]]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/01/image.png" />

// Example 2:
// Input: grid = [[0,1],[0,1],[0,0]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/08/screenshot-from-2024-07-09-01-37-48.png" />

// Example 3:
// Input: grid = [[1],[1]]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/01/screenshot-from-2024-08-01-23-05-26.png" />

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m * n <= 2 * 10^5
//     0 <= grid[i][j] <= 1

import "fmt"

func minFlips(grid [][]int) int {
    res, m, n := 0, len(grid), len(grid[0])
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // Process each 2x2 block of the grid
    for i := 0; i < m/2; i++ {
        for j := 0; j < n/2; j++ {
            // Count the number of 1s in the current 2x2 block
            count := grid[i][j]
            count += grid[m-1-i][j]
            count += grid[i][n-1-j]
            count += grid[m-1-i][n-1-j]
            // Calculate flips needed to make all cells the same
            res += min(count, 4 - count)
        }
    }
    // Handle center row if the number of rows is odd
    diff, p0, p1 := 0, 0, 0
    if m % 2 ==1 {
        for j := 0; j < n/2; j++ {
            if grid[m/2][j] != grid[m/2][n-j-1] {
                diff++
            } else {
                if grid[m/2][j] == 0 {
                    p0++
                } else {
                    p1++
                }
            }
        }
    }
    if n % 2 == 1 {
        for i := 0; i < m / 2; i++ {
            if grid[i][n/2] != grid[m-i-1][n/2] {
                diff++
            } else {
                if(grid[i][n/2] == 0) {
                    p0++
                } else {
                    p1++
                }
            }
        }
    }
    // Adjust for the center cell if both dimensions are odd
    if n % 2 == 1 && m % 2 == 1 {
        if grid[m/2][n/2] == 1 { res++ }
    }
    // Calculate the minimum number of additional flips needed for center row and column
    res1 := 1 << 31
    if diff % 2 == p1 % 2 {
        res1 = diff
    } else {
        if diff % 2 == 0 {
            if diff == 0 {
                res1 = 2
            } else {
                res1 = diff
            }
        } else {
            res1 = diff
        }
    }
    return res + res1
}

func minFlips1(grid [][]int) (ans int) {
    res, m, n := 0, len(grid), len(grid[0])
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, row := range grid[:m/2] {
        row2 := grid[m-1-i]
        for j, x := range row[:n/2] {
            cnt1 := x + row[n-1-j] + row2[j] + row2[n-1-j]
            res += min(cnt1, 4-cnt1) // 全为 1 或全为 0
        }
    }
    if m % 2 > 0 && n % 2 > 0 { // 正中间的数必须是 0
        res += grid[m/2][n/2]
    }
    diff, cnt1 := 0, 0
    if m % 2 > 0 {
        // 统计正中间这一排
        row := grid[m/2]
        for j, x := range row[:n/2] {
            if x != row[n-1-j] {
                diff++
            } else {
                cnt1 += x * 2
            }
        }
    }
    if n % 2 > 0 {
        // 统计正中间这一列
        for i, row := range grid[:m/2] {
            if row[n/2] != grid[m-1-i][n/2] {
                diff++
            } else {
                cnt1 += row[n/2] * 2
            }
        }
    }
    if diff > 0 {
        res += diff
    } else {
        res += cnt1 % 4
    }
    return res
}

func main() {
    // Example 1:
    // Input: grid = [[1,0,0],[0,1,0],[0,0,1]]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/01/image.png" />
    fmt.Println(minFlips([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3
    // Example 2:
    // Input: grid = [[0,1],[0,1],[0,0]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/08/screenshot-from-2024-07-09-01-37-48.png" />
    fmt.Println(minFlips([][]int{{0,1},{0,1},{0,0}})) // 2
    // Example 3:
    // Input: grid = [[1],[1]]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/01/screenshot-from-2024-08-01-23-05-26.png" />
    fmt.Println(minFlips([][]int{{1},{1}})) // 2

    fmt.Println(minFlips1([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3
    fmt.Println(minFlips1([][]int{{0,1},{0,1},{0,0}})) // 2
    fmt.Println(minFlips1([][]int{{1},{1}})) // 2
}