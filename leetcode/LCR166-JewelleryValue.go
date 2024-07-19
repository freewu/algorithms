package main

// LCR 166. 珠宝的最高价值
// 现有一个记作二维矩阵 frame 的珠宝架，其中 frame[i][j] 为该位置珠宝的价值。拿取珠宝的规则为：
//     只能从架子的左上角开始拿珠宝
//     每次可以移动到右侧或下侧的相邻位置
//     到达珠宝架子的右下角时，停止拿取

// 注意：珠宝的价值都是大于 0 的。除非这个架子上没有任何珠宝，比如 frame = [[0]]。

// 示例 1:
// 输入: frame = [[1,3,1],[1,5,1],[4,2,1]]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最高价值的珠宝

// 提示：
//     0 < frame.length <= 200
//     0 < frame[0].length <= 200

import "fmt"

func jewelleryValue(frame [][]int) int {
    m, n := len(frame), len(frame[0])
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                continue
            }
            if i == 0 {
                frame[i][j] += frame[i][j-1]
            } else if j == 0 {
                frame[i][j] += frame[i-1][j]
            } else {
                frame[i][j] += max(frame[i][j-1], frame[i-1][j])
            }
        }
    }
    return frame[m-1][n-1]
}

func jewelleryValue1(frame [][]int) int {
    m, n := len(frame), len(frame[0])
    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i, row := range frame {
        for j, x := range row {
            dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1]) + x
        }
    }
    return dp[m][n]
}

func main() {
    // 示例 1:
    // 输入: frame = [[1,3,1],[1,5,1],[4,2,1]]
    // 输出: 12
    // 解释: 路径 1→3→5→2→1 可以拿到最高价值的珠宝
    fmt.Println(jewelleryValue([][]int{{1,3,1},{1,5,1},{4,2,1}}))

    fmt.Println(jewelleryValue1([][]int{{1,3,1},{1,5,1},{4,2,1}}))
}