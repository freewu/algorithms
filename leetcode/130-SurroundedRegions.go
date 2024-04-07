package main

// 130. Surrounded Regions
// Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.

// Example 1:
//     X X X X
//     X O O X
//     X X O X
//     X O X X
// After running your function, the board should be:
//     X X X X
//     X X X X
//     X X X X
//     X O X X
// <img src="https://assets.leetcode.com/uploads/2021/02/19/xogrid.jpg"/>
// Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// Explanation: Notice that an 'O' should not be flipped if:
// - It is on the border, or
// - It is adjacent to an 'O' that should not be flipped.
// The bottom 'O' is on the border, so it is not flipped.
// The other three 'O' form a surrounded region, so they are flipped.

// Example 2:
// Input: board = [["X"]]
// Output: [["X"]]
 
// Constraints:
//     m == board.length
//     n == board[i].length
//     1 <= m, n <= 200
//     board[i][j] is 'X' or 'O'.

// 解题思路
//     给出一张二维地图，要求把地图上非边缘上的 ‘O’ 都用 ‘X’ 替换掉

import "fmt"

// dfs
func solve(board [][]byte) {
    dir := [][]int{{-1, 0},  {0, 1},  {1, 0},  {0, -1}}
    var dfs func(i, j int, board [][]byte) 
    dfs = func(i, j int, board [][]byte) {
        if i < 0 || i > len(board) - 1 || j < 0 || j > len(board[i])-1 {
            return
        }
        if board[i][j] == 'O' {
            board[i][j] = '*'
            for k := 0; k < 4; k++ { //  处理四周
                dfs(i+ dir[k][0], j + dir[k][1], board)
            }
        }
    }
    for i := range board {
        //fmt.Printf("before: %v\n",board[i])
        for j := range board[i] {
            if i == 0 || i == len(board) - 1 || j == 0 || j == len(board[i]) - 1 { // 处理外围一圈 如果是 O 先设置为 *
                // fmt.Printf("board[%v][%v] = %v\n",i,j,board[i][j])
                if board[i][j] == 'O' {
                    dfs(i, j, board)
                }
            }
        }
        //fmt.Printf("after: %v\n",board[i])
    }
    // 循环二维数组  * => O O => X
    for i := range board {
        for j := range board[i] {
            if board[i][j] == '*' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}

func solve1(board [][]byte)  {
    n, m := len(board), len(board[0])
    vis := [][]bool{}
    for i := 0; i < n; i ++ {
        vis = append(vis, make([]bool, m))
    }
    dir := [4][2]int{{1,0},{0,1},{-1,0},{0,-1}}
    var dfs func (x, y, n, m int, board [][]byte) 
    dfs = func (x, y, n, m int, board [][]byte) {
        if vis[x][y] {
            return
        }
        vis[x][y] = true
        for i := 0; i < 4; i ++ {
            a, b := dir[i][0] + x, dir[i][1] + y
            if a >= 0 && a < n && b >= 0 && b < m && board[a][b] == 'O' {
                dfs(a, b, n, m, board)
            }
        }
    }
    for i := 0; i < n; i ++ {
        for j := 0; j < m; j ++ {
            if (i == 0 || j == 0 || i == n - 1 || j == m - 1) && board[i][j] == 'O' {
                dfs(i, j, n, m, board)
            }
        }
    }
    for i := 0; i < n; i ++ {
        for j := 0; j < m; j ++ {
            if !vis[i][j] {
                board[i][j] = 'X'
            }
        }
    }
}

func main() {
    board1 := [][]byte{{'X'}}
    fmt.Printf("board before solve = %v\n",board1) // [[88]]
    solve(board1)
    fmt.Printf("board after solve = %v\n",board1) // [[88]]

    board2 := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
    fmt.Printf("board before solve = %v\n",board2) // [[88 88 88 88] [88 79 79 88] [88 88 79 88] [88 79 88 88]]
    solve(board2)
    fmt.Printf("board after solve = %v\n",board2) // [[88 88 88 88] [88 88 88 88] [88 88 88 88] [88 79 88 88]]
    
    board11 := [][]byte{{'X'}}
    fmt.Printf("board before solve = %v\n",board11) // [[88]]
    solve(board11)
    fmt.Printf("board after solve = %v\n",board11) // [[88]]

    board12 := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
    fmt.Printf("board before solve = %v\n",board12) // [[88 88 88 88] [88 79 79 88] [88 88 79 88] [88 79 88 88]]
    solve(board12)
    fmt.Printf("board after solve = %v\n",board12) // [[88 88 88 88] [88 88 88 88] [88 88 88 88] [88 79 88 88]]
}