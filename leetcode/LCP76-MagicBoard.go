package main

// LCP 76. 魔法棋盘
// 在大小为 n * m 的棋盘中，有两种不同的棋子：黑色，红色。
// 当两颗颜色不同的棋子同时满足以下两种情况时，将会产生魔法共鸣：
//     1. 两颗异色棋子在同一行或者同一列
//     2. 两颗异色棋子之间恰好只有一颗棋子

// 注：异色棋子之间可以有空位

// 由于棋盘上被施加了魔法禁制，棋盘上的部分格子变成问号。
// chessboard[i][j] 表示棋盘第 i 行 j 列的状态：
//     1. 若为 . ，表示当前格子确定为空
//     2. 若为 B ，表示当前格子确定为 黑棋
//     3. 若为 R ，表示当前格子确定为 红棋
//     4. 若为 ? ，表示当前格子待定

// 现在，探险家小扣的任务是确定所有问号位置的状态（留空/放黑棋/放红棋），使最终的棋盘上，任意两颗棋子间都 无法 产生共鸣。
// 请返回可以满足上述条件的放置方案数量。

// 示例1：
// 输入：n = 3, m = 3, chessboard = ["..R","..B","?R?"]
// 输出：5
// 解释：给定的棋盘如图：image.png所有符合题意的最终局面如图：image.png

// 示例2：
// 输入：n = 3, m = 3, chessboard = ["?R?","B?B","?R?"]
// 输出：105

// 提示：
//     n == chessboard.length
//     m == chessboard[i].length
//     1 <= n*m <= 30
//     chessboard 中仅包含 "."、"B"、"R"、"?"

import "fmt"

func getSchemeCount(n int, m int, chessboard []string) int64 {
    // 每一行的含义：{当前序列末尾添加 B 后的状态，当前序列末尾添加 R 后的状态}
    trans := [7][2]int{
        {1, 2},  // 空
        {3, 6},  // 只有一个 B
        {5, 4},  // 只有一个 R
        {3, -1}, // 连续多个 B
        {-1, 4}, // 连续多个 R
        {-1, 6}, // BR 交替，且以 B 结尾
        {5, -1}, // BR 交替，且以 R 结尾
    }
    arr := make([][]byte, n)
    for i, row := range chessboard {
        arr[i] = []byte(row)
    }
    rotate := func(arr [][]byte) [][]byte {
        n, m := len(arr), len(arr[0])
        res := make([][]byte, m)
        for i := range res {
            res[i] = make([]byte, n)
        }
        for i, r := range arr {
            for j, v := range r {
                res[j][n - i - 1] = v
            }
        }
        return res
    }
    if n < m {
        arr = rotate(arr) // 保证 n >= m
        n, m = m, n
    }
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, 1 << (m * 3))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs1 func(r, mask int) int
    dfs1 = func(r, mask int) int {
        if r == n { return 1 } // 找到 1 个合法方案
        if memo[r][mask] != -1 { return memo[r][mask] }
        // 写一个爆搜，生成所有的合法状态
        var dfs2 func(c, rowMask, colMask int) int
        dfs2 = func(c, rowMask, colMask int) (res int) {
            if c == m { // 方案合法
                return dfs1(r + 1, colMask) // 枚举下一行
            }
            next := func(color int) int {
                rm := trans[rowMask][color] // 新的 rowMask
                if rm < 0 { return 0 }// 非法
                c3 := c * 3
                cm := trans[colMask >> c3 & 7][color] // 新的 colMask 的第 c 列
                if cm < 0 { return 0 } // 非法
                return dfs2(c + 1, rm, colMask&^(7<<c3)|cm<<c3) // 修改 colMask 的第 c 列
            }
            switch arr[r][c] {
            case 'B': // 填 B
                return next(0)
            case 'R': // 填 R
                return next(1)
            case '?': // 留空 / 填 B / 填 R
                return dfs2(c+1, rowMask, colMask) + next(0) + next(1)
            default: // 留空
                return dfs2(c+1, rowMask, colMask)
            }
        }
        memo[r][mask] = dfs2(0, 0, mask)
        return memo[r][mask]
    }
    return int64(dfs1(0, 0))
}

func main() {
    // 示例1：
    // 输入：n = 3, m = 3, chessboard = ["..R","..B","?R?"]
    // 输出：5
    // 解释：给定的棋盘如图：image.png所有符合题意的最终局面如图：image.png
    fmt.Println(getSchemeCount(3, 3, []string{"..R","..B","?R?"})) // 5
    // 示例2：
    // 输入：n = 3, m = 3, chessboard = ["?R?","B?B","?R?"]
    // 输出：105
    fmt.Println(getSchemeCount(3, 3, []string{"?R?","B?B","?R?"})) // 105
}