package main

// 764. Largest Plus Sign
// You are given an integer n. 
// You have an n x n binary grid grid with all values initially 1's except for some indices given in the array mines. 
// The ith element of the array mines is defined as mines[i] = [xi, yi] where grid[xi][yi] == 0.

// Return the order of the largest axis-aligned plus sign of 1's contained in grid. If there is none, return 0.

// An axis-aligned plus sign of 1's of order k has some center grid[r][c] == 1 along with four arms of length k - 1 going up, down, left, and right, and made of 1's. 
// Note that there could be 0's or 1's beyond the arms of the plus sign, only the relevant area of the plus sign is checked for 1's.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/plus1-grid.jpg" />
// Input: n = 5, mines = [[4,2]]
// Output: 2
// Explanation: In the above grid, the largest plus sign can only be of order 2. One of them is shown.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/plus2-grid.jpg" />
// Input: n = 1, mines = [[0,0]]
// Output: 0
// Explanation: There is no plus sign, so return 0.

// Constraints:
//     1 <= n <= 500
//     1 <= mines.length <= 5000
//     0 <= xi, yi < n
//     All the pairs (xi, yi) are unique.

import "fmt"

// 模拟
func orderOfLargestPlusSign(n int, mines [][]int) int {
    res, grid := 0, make([][]int, n)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            grid[i][j] = 1
        }
    }
    for _, v := range mines {
        grid[v[0]][v[1]] = 0
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                mx, dir := 1, 1
                for j - dir >= 0 && j + dir < n && i - dir >= 0 && i + dir < n && // 边界检测
                    grid[i-dir][j] == 1 && grid[i+dir][j] == 1 &&
                    grid[i][j-dir] == 1 && grid[i][j+dir] == 1 {
                    dir++
                    mx++
                }
                if mx > res { 
                    res = mx 
                }
            }
        }
    }
    return res
}

func orderOfLargestPlusSign1(n int, mines [][]int) int {
    grid := make([][]int, n)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            grid[i][j] = n
        }
    }
    for _, v := range mines {
        grid[v[0]][v[1]] = 0
    }
    for i := 0; i < n; i++ {
        left, right, up, down := 0, 0, 0, 0
        for j, k := 0, n - 1; j < n; j, k = j + 1, k - 1 {
            if grid[i][j] == 0 {
                left = 0
            } else {
                left += 1
            }
            if left < grid[i][j] { grid[i][j] = left }
            if grid[i][k] == 0 {
                right = 0
            } else {
                right += 1
            }
            if right < grid[i][k] { grid[i][k] = right }
            if grid[j][i] == 0 {
                up = 0
            } else {
                up += 1
            }
            if up < grid[j][i] { grid[j][i] = up }
            if grid[k][i] == 0 {
                down = 0
            } else {
                down += 1
            }
            if down < grid[k][i] { grid[k][i] = down }
        }
    }
    res := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] > res { 
                res = grid[i][j] 
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/plus1-grid.jpg" />
    // Input: n = 5, mines = [[4,2]]
    // Output: 2
    // Explanation: In the above grid, the largest plus sign can only be of order 2. One of them is shown.
    fmt.Println(orderOfLargestPlusSign(5,[][]int{{4,2}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/plus2-grid.jpg" />
    // Input: n = 1, mines = [[0,0]]
    // Output: 0
    // Explanation: There is no plus sign, so return 0.
    fmt.Println(orderOfLargestPlusSign(1,[][]int{{0,0}})) // 0

    fmt.Println(orderOfLargestPlusSign(5,[][]int{{0,0}})) // 3

    fmt.Println(orderOfLargestPlusSign1(5,[][]int{{4,2}})) // 2
    fmt.Println(orderOfLargestPlusSign1(1,[][]int{{0,0}})) // 0
    fmt.Println(orderOfLargestPlusSign1(5,[][]int{{0,0}})) // 3
}