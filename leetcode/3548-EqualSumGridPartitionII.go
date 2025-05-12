package main

// 3548. Equal Sum Grid Partition II
// You are given an m x n matrix grid of positive integers. 
// Your task is to determine if it is possible to make either one horizontal or one vertical cut on the grid such that:
//     1. Each of the two resulting sections formed by the cut is non-empty.
//     2. The sum of elements in both sections is equal, or can be made equal by discounting at most one single cell in total (from either section).
//     3. If a cell is discounted, the rest of the section must remain connected.

// Return true if such a partition exists; otherwise, return false.

// Note: A section is connected if every cell in it can be reached from any other cell by moving up, down, left, or right through other cells in the section.

// Example 1:
// Input: grid = [[1,4],[2,3]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/03/30/lc.jpeg" />
// A horizontal cut after the first row gives sums 1 + 4 = 5 and 2 + 3 = 5, which are equal. Thus, the answer is true.

// Example 2:
// Input: grid = [[1,2],[3,4]]
// Output: true
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/01/chatgpt-image-apr-1-2025-at-05_28_12-pm.png" />
// A vertical cut after the first column gives sums 1 + 3 = 4 and 2 + 4 = 6.
// By discounting 2 from the right section (6 - 2 = 4), both sections have equal sums and remain connected. Thus, the answer is true.

// Example 3:
// Input: grid = [[1,2,4],[2,3,5]]
// Output: false
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/01/chatgpt-image-apr-2-2025-at-02_50_29-am.png" />
// A horizontal cut after the first row gives 1 + 2 + 4 = 7 and 2 + 3 + 5 = 10.
// By discounting 3 from the bottom section (10 - 3 = 7), both sections have equal sums, but they do not remain connected as it splits the bottom section into two parts ([2] and [5]). Thus, the answer is false.

// Example 4:
// Input: grid = [[4,1,8],[3,2,6]]
// Output: false
// Explanation:
// No valid cut exists, so the answer is false.

// Constraints:
//     1 <= m == grid.length <= 10^5
//     1 <= n == grid[i].length <= 10^5
//     2 <= m * n <= 10^5
//     1 <= grid[i][j] <= 10^5

import "fmt"
import "slices"

func canPartitionGrid(grid [][]int) bool {
    total := 0
    for _, row := range grid {
        for _, v := range row {
            total += v
        }
    }
    // 能否水平分割
    check := func(grid [][]int) bool {
        m, n := len(grid), len(grid[0])
        f := func() bool {
            mp := map[int]bool{0: true} // 0 对应不删除数字
            sum := 0
            for i, row := range grid[:m-1] {
                for j, v := range row {
                    sum += v
                    if i > 0 || j == 0 || j == n - 1 { // 第一行，不能删除中间元素
                        mp[v] = true
                    }
                }
                if n == 1 { // 特殊处理只有一列的情况，此时只能删除第一个数或者分割线上那个数
                    if sum * 2 == total || sum * 2 - total == grid[0][0] || sum * 2 - total == row[0] {
                        return true
                    }
                    continue
                }
                if mp[sum * 2 - total] {
                    return true
                }
                if i == 0 { // 如果分割到更下面，那么可以删第一行的元素
                    for _, v := range row {
                        mp[v] = true
                    }
                }
            }
            return false
        }
        // 删除上半部分中的一个数
        if f() { return true }
        // 删除下半部分中的一个数
        slices.Reverse(grid)
        return f()
    }
    // 顺时针旋转矩阵 90°
    rotate := func(grid [][]int) [][]int {
        m, n := len(grid), len(grid[0])
        res := make([][]int, n)
        for i := range res {
            res[i] = make([]int, m)
        }
        for i, row := range grid {
            for j, v := range row {
                res[j][m - 1 - i] = v
            }
        }
        return res
    }
    // 水平分割 or 垂直分割
    return check(grid) || check(rotate(grid))
}

func main() {
    // Example 1:
    // Input: grid = [[1,4],[2,3]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/03/30/lc.jpeg" />
    // A horizontal cut after the first row gives sums 1 + 4 = 5 and 2 + 3 = 5, which are equal. Thus, the answer is true.
    fmt.Println(canPartitionGrid([][]int{{1,4},{2,3}})) // true
    // Example 2:
    // Input: grid = [[1,2],[3,4]]
    // Output: true
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/01/chatgpt-image-apr-1-2025-at-05_28_12-pm.png" />
    // A vertical cut after the first column gives sums 1 + 3 = 4 and 2 + 4 = 6.
    // By discounting 2 from the right section (6 - 2 = 4), both sections have equal sums and remain connected. Thus, the answer is true.
    fmt.Println(canPartitionGrid([][]int{{1,2},{3,4}})) // true
    // Example 3:
    // Input: grid = [[1,2,4],[2,3,5]]
    // Output: false
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/01/chatgpt-image-apr-2-2025-at-02_50_29-am.png" />
    // A horizontal cut after the first row gives 1 + 2 + 4 = 7 and 2 + 3 + 5 = 10.
    // By discounting 3 from the bottom section (10 - 3 = 7), both sections have equal sums, but they do not remain connected as it splits the bottom section into two parts ([2] and [5]). Thus, the answer is false.
    fmt.Println(canPartitionGrid([][]int{{1,2,4},{2,3,5}})) // false
    // Example 4:
    // Input: grid = [[4,1,8],[3,2,6]]
    // Output: false
    // Explanation:
    // No valid cut exists, so the answer is false.
    fmt.Println(canPartitionGrid([][]int{{4,1,8},{3,2,6}})) // false
}