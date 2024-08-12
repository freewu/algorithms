package main

// LCR 130. 衣橱整理
// 家居整理师将待整理衣橱划分为 m x n 的二维矩阵 grid，其中 grid[i][j] 代表一个需要整理的格子。
// 整理师自 grid[0][0] 开始 逐行逐列 地整理每个格子。

// 整理规则为：在整理过程中，可以选择 向右移动一格 或 向下移动一格，但不能移动到衣柜之外。
// 同时，不需要整理 digit(i) + digit(j) > cnt 的格子，其中 digit(x) 表示数字 x 的各数位之和。

// 请返回整理师 总共需要整理多少个格子。

// 示例 1：
// 输入：m = 4, n = 7, cnt = 5
// 输出：18

// 提示：
//     1 <= n, m <= 100
//     0 <= cnt <= 20

import "fmt"

func wardrobeFinishing(m int, n int, cnt int) int {
    visited := make([][]bool, m)
    for i, _ := range visited { 
        visited[i] = make([]bool, n) 
    }
    sum := func (num int) int {
        res := 0
        for num != 0 {
            res += num % 10
            num /= 10
        }
        return res
    }
    var dfs func (i, j int) int
    dfs = func (i, j int) int {
        if ( i >= m || j >= n || visited[i][j] || cnt < sum(i) + sum(j)) { 
            return 0 
        }
        visited[i][j] = true
        return 1 + dfs(i+1, j) + dfs(i, j+1)
    }
    return dfs(0, 0)
}

func main() {
    // 示例 1：
    // 输入：m = 4, n = 7, cnt = 5
    // 输出：18
    fmt.Println(wardrobeFinishing(4,7,5)) // 18
}