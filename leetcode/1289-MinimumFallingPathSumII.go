package main

// 1289. Minimum Falling Path Sum II
// Given an n x n integer matrix grid, return the minimum sum of a falling path with non-zero shifts.
// A falling path with non-zero shifts is a choice of exactly one element from each row of grid such that no two elements chosen in adjacent rows are in the same column.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/10/falling-grid.jpg" />
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
// Output: 13
// Explanation: 
// The possible falling paths are:
// [1,5,9], [1,5,7], [1,6,7], [1,6,8],
// [2,4,8], [2,4,9], [2,6,7], [2,6,8],
// [3,4,8], [3,4,9], [3,5,7], [3,5,9]
// The falling path with the smallest sum is [1,5,7], so the answer is 13.

// Example 2:
// Input: grid = [[7]]
// Output: 7
 
// Constraints:
//     n == grid.length == grid[i].length
//     1 <= n <= 200
//     -99 <= grid[i][j] <= 99

import "fmt"

func minFallingPathSum(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    dp := make(map[int]map[int]int, n) // 定义一个map，存储从当前元素开始 非零偏移下降路径 数字和的最小值
    for i := 0; i < n; i++ {
        dp[i] = make(map[int]int, m - 1)
    }
    var dfs func(int, int) (bool, int)
    dfs = func(i int, j int) (bool, int) {
        // 边界情况判断
        if i < 0 || j < 0 || i > n - 1 || j > m - 1 {
            return false, 0
        }
        num := grid[i][j]
        if i == n - 1 { // 如果为最后一行，直接返回
            return true, num
        }
        if mn, ok := dp[i][j]; ok { // 判断当前元素是否已经存在最小值的结果，存在就直接返回
            return true, mn
        }
        mn, isFirst := 0, true
        for k := 0; k < m; k++ { // 遍历下一行
            if k == j { // 排除不符合条件的情况
                continue
            }
            b, res := dfs(i + 1, k)
            if b && isFirst {
                isFirst = false
                mn = res
                continue
            }
            if b && res < mn { // 找到当前元素开始的 非零偏移下降路径 数字和的最小值
                mn = res
            }
        }
        // 写入map记录
        mn += num
        dp[i][j] = mn
        return true, mn
    }
    res := 0
    for j := 0; j < m; j++ { // 遍历首行，找到最小值
        _, v := dfs(0, j)
        if j == 0 {
            res = v
            continue
        }
        if v < res {
            res = v
        }
    }
    return res
}

func minFallingPathSum1(grid [][]int) (ans int) {
    n, inf := len(grid), 1 << 32 -1
    prevFirstMin, prevSecondMin, prevFirstMinIdx := 0, 0, -1
    for i := 0; i < n; i++ {
        curFirstMin, curSecondMin, curFirstMinIdx := inf, inf, -1
        for j := 0; j < n; j++ {
            curSum := grid[i][j]
            if j == prevFirstMinIdx {
                curSum += prevSecondMin
            } else {
                curSum += prevFirstMin
            }
            if curSum < curFirstMin {
                curFirstMin, curSecondMin = curSum, curFirstMin
                curFirstMinIdx = j
            } else if curSum < curSecondMin {
                curSecondMin = curSum
            }
        }
        prevFirstMin, prevSecondMin, prevFirstMinIdx = curFirstMin, curSecondMin, curFirstMinIdx
    }
    return prevFirstMin
}

func main() {
    // Example 1:
    // Input: grid = [[1,2,3],[4,5,6],[7,8,9]]
    // Output: 13
    // Explanation: 
    // The possible falling paths are:
    // [1,5,9], [1,5,7], [1,6,7], [1,6,8],
    // [2,4,8], [2,4,9], [2,6,7], [2,6,8],
    // [3,4,8], [3,4,9], [3,5,7], [3,5,9]
    // The falling path with the smallest sum is [1,5,7], so the answer is 13.
    fmt.Println(minFallingPathSum([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 13
    // Example 2:
    // Input: grid = [[7]]
    // Output: 7
    fmt.Println(minFallingPathSum([][]int{{7}})) // 7

    fmt.Println(minFallingPathSum1([][]int{{1,2,3},{4,5,6},{7,8,9}})) // 13
    fmt.Println(minFallingPathSum1([][]int{{7}})) // 7
}