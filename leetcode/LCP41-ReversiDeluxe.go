package main

// LCP 41. 黑白翻转棋
// 在 n*m 大小的棋盘中，有黑白两种棋子，黑棋记作字母 "X", 白棋记作字母 "O"，空余位置记作 "."。
// 当落下的棋子与其他相同颜色的棋子在行、列或对角线完全包围（中间不存在空白位置）另一种颜色的棋子，则可以翻转这些棋子的颜色。

// <img src="https://pic.leetcode-cn.com/1630396029-eTgzpN-6da662e67368466a96d203f67bb6e793.gif" />
// <img src="https://pic.leetcode-cn.com/1630396240-nMvdcc-8e4261afe9f60e05a4f740694b439b6b.gif" />
// <img src="https://pic.leetcode-cn.com/1630396291-kEtzLL-6fcb682daeecb5c3f56eb88b23c81d33.gif" />

// 「力扣挑战赛」黑白翻转棋项目中，将提供给选手一个未形成可翻转棋子的棋盘残局，其状态记作 chessboard。
// 若下一步可放置一枚黑棋，请问选手最多能翻转多少枚白棋。

// 注意：
// 若翻转白棋成黑棋后，棋盘上仍存在可以翻转的白棋，将可以 继续 翻转白棋
// 输入数据保证初始棋盘状态无可以翻转的棋子且存在空余位置

// 示例 1：
// 输入：chessboard = ["....X.","....X.","XOOO..","......","......"]
// 输出：3
// 解释： 可以选择下在 [2,4] 处，能够翻转白方三枚棋子。

// 示例 2：
// 输入：chessboard = [".X.",".O.","XO."]
// 输出：2
// 解释： 可以选择下在 [2,2] 处，能够翻转白方两枚棋子。
// <img src="https://pic.leetcode-cn.com/1626683255-OBtBud-2126c1d21b1b9a9924c639d449cc6e65.gif" />

// 示例 3：
// 输入：chessboard = [".......",".......",".......","X......",".O.....","..O....","....OOX"]
// 输出：4
// 解释： 可以选择下在 [6,3] 处，能够翻转白方四枚棋子。
// <img src="https://pic.leetcode-cn.com/1630393770-Puyked-803f2f04098b6174397d6c696f54d709.gif" />

// 提示：
//     1 <= chessboard.length, chessboard[i].length <= 8
//     chessboard[i] 仅包含 "."、"O" 和 "X"

import "fmt"

func flipChess(chessboard []string) int {
    res, m, n := 0, len(chessboard), len(chessboard[0])
    max := func (x, y int) int { if x > y { return x; }; return y; }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    bfs := func(i, j int) int {
        queue := [][2]int{{i, j}}
        g := make([][]byte, m)
        for i := range g {
            g[i] = make([]byte, n)
            copy(g[i], []byte(chessboard[i]))
        }
        g[i][j] = 'X'
        count := 0
        for len(queue) > 0 {
            p := queue[0]
            queue = queue[1:]
            i, j = p[0], p[1]
            for a := -1; a <= 1; a++ {
                for b := -1; b <= 1; b++ {
                    if a == 0 && b == 0 { continue }
                    x, y := i + a, j + b
                    for x >= 0 && x < m && y >= 0 && y < n && g[x][y] == 'O' {
                        x, y = x + a, y + b
                    }
                    if x >= 0 && x < m && y >= 0 && y < n && g[x][y] == 'X' {
                        x -= a
                        y -= b
                        count += max(abs(x - i), abs(y - j))
                        for x != i || y != j {
                            g[x][y] = 'X'
                            queue = append(queue, [2]int{x, y})
                            x -= a
                            y -= b
                        }
                    }
                }
            }
        }
        return count
    }
    for i, row := range chessboard {
        for j, c := range row {
            if c == '.' {
                res = max(res, bfs(i, j))
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：chessboard = ["....X.","....X.","XOOO..","......","......"]
    // 输出：3
    // 解释： 可以选择下在 [2,4] 处，能够翻转白方三枚棋子。
    fmt.Println(flipChess([]string{"....X.","....X.","XOOO..","......","......"})) // 3
    // 示例 2：
    // 输入：chessboard = [".X.",".O.","XO."]
    // 输出：2
    // 解释： 可以选择下在 [2,2] 处，能够翻转白方两枚棋子。
    // <img src="https://pic.leetcode-cn.com/1626683255-OBtBud-2126c1d21b1b9a9924c639d449cc6e65.gif" />
    fmt.Println(flipChess([]string{".X.",".O.","XO."})) // 2
    // 示例 3：
    // 输入：chessboard = [".......",".......",".......","X......",".O.....","..O....","....OOX"]
    // 输出：4
    // 解释： 可以选择下在 [6,3] 处，能够翻转白方四枚棋子。
    // <img src="https://pic.leetcode-cn.com/1630393770-Puyked-803f2f04098b6174397d6c696f54d709.gif" />
    fmt.Println(flipChess([]string{".......",".......",".......","X......",".O.....","..O....","....OOX"})) // 4
}