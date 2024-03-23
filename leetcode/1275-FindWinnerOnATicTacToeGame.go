package main

// 1275. Find Winner on a Tic Tac Toe Game
// Tic-tac-toe is played by two players A and B on a 3 x 3 grid. The rules of Tic-Tac-Toe are:
//     Players take turns placing characters into empty squares ' '.
//     The first player A always places 'X' characters, while the second player B always places 'O' characters.
//     'X' and 'O' characters are always placed into empty squares, never on filled ones.
//     The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
//     The game also ends if all squares are non-empty.
//     No more moves can be played if the game is over.

// Given a 2D integer array moves where moves[i] = [rowi, coli] indicates that the ith move will be played on grid[rowi][coli]. 
// return the winner of the game if it exists (A or B). In case the game ends in a draw return "Draw". 
// If there are still movements to play return "Pending".
// You can assume that moves is valid (i.e., it follows the rules of Tic-Tac-Toe), the grid is initially empty, and A will play first.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/09/22/xo1-grid.jpg" />
// Input: moves = [[0,0],[2,0],[1,1],[2,1],[2,2]]
// Output: "A"
// Explanation: A wins, they always play first.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/09/22/xo2-grid.jpg" />
// Input: moves = [[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]]
// Output: "B"
// Explanation: B wins.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/09/22/xo3-grid.jpg" />
// Input: moves = [[0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2]]
// Output: "Draw"
// Explanation: The game ends in a draw since there are no moves to make.
 
// Constraints:
//     1 <= moves.length <= 9
//     moves[i].length == 2
//     0 <= rowi, coli <= 2
//     There are no repeated elements on moves.
//     moves follow the rules of tic tac toe.

import "fmt"

func tictactoe(moves [][]int) string {
    grid := [3][3]int{} // 初始化一个 3 * 3 的二维组
    // 将步骤在二维组中重放
    for i, move := range moves {
        grid[move[0]][move[1]] = i & 1 + 1 // odd turn will be 1, even turn will be 2
    }
    // 检查 row col 是否值一样
    for i := 0; i < 3; i++ {
        if player := grid[i][i]; player != 0 {
            if (grid[i][0] == grid[i][1] && grid[i][0] == grid[i][2]) || (grid[0][i] == grid[1][i] && grid[0][i] == grid[2][i]) {
                return string(byte('@' + player))
            }
        }
    }
    // 检查交叉 是否有值一样
    if player := grid[1][1]; player != 0 {
        if (grid[0][0] == grid[2][2] && grid[0][0] == player) || (grid[2][0] == grid[0][2] && grid[2][0] == player) {
            return string(byte('@' + player))
        }
    }
    // 不够 9 步说明还未下完
    if len(moves) != 9 {
        return "Pending"
    }

    return "Draw"
}

func tictactoe1(moves [][]int) string {
    // 初始化棋盘、玩家标识、棋盘坐标
    q, p, x, y := [][]int{{-4, -4, -4}, {-4, -4, -4}, {-4, -4, -4}}, 0, 0, 0
    // 开始下棋,"X"=0,"O"=1
    for _, v := range moves {
        x, y = v[0], v[1]
        // 落子
        q[x][y] = p % 2
        // 检验胜负
        // 行列
        if q[0][y]+q[1][y]+q[2][y] == p%2*3 || q[x][0]+q[x][1]+q[x][2] == p%2*3 {
            if p%2 == 0 {
                return "A"
            } else {
                return "B"
            }
        }
        // 对角线
        if (x == y || y+x == 2) && (q[0][0]+q[1][1]+q[2][2] == p%2*3 || q[2][0]+q[1][1]+q[0][2] == p%2*3) {
            if p%2 == 0 {
                return "A"
            } else {
                return "B"
            }
        }
        p++
    }
    // 游戏未结束
    if p < 9 {
        return "Pending"
    }
    // 满棋子
    return "Draw"
}

func main() {
    fmt.Println(tictactoe(
        [][]int {
            []int{0,0},             
            []int{2,0},
            []int{1,1},
            []int{2,1},             
            []int{2,2},
        },
    )) // A
    fmt.Println(tictactoe(
        [][]int {
            []int{0,0},             
            []int{1,1},
            []int{0,1},             
            []int{0,2},
            []int{1,0},
            []int{2,0},
        },
    )) // B
    fmt.Println(tictactoe(
        [][]int {
            []int{0,0},             
            []int{1,1},
            []int{2,0},             
            []int{1,0},
            []int{1,2},
            []int{2,1},
            []int{0,1},
            []int{0,2},
            []int{2,2},
        },
    )) // Draw

    fmt.Println(tictactoe1(
        [][]int {
            []int{0,0},             
            []int{2,0},
            []int{1,1},
            []int{2,1},             
            []int{2,2},
        },
    )) // A
    fmt.Println(tictactoe1(
        [][]int {
            []int{0,0},             
            []int{1,1},
            []int{0,1},             
            []int{0,2},
            []int{1,0},
            []int{2,0},
        },
    )) // B
    fmt.Println(tictactoe1(
        [][]int {
            []int{0,0},             
            []int{1,1},
            []int{2,0},             
            []int{1,0},
            []int{1,2},
            []int{2,1},
            []int{0,1},
            []int{0,2},
            []int{2,2},
        },
    )) // Draw
}