package main

// 529. Minesweeper
// Let's play the minesweeper game (Wikipedia, online game)!
// You are given an m x n char matrix board representing the game board where:
//     'M' represents an unrevealed mine,
//     'E' represents an unrevealed empty square,
//     'B' represents a revealed blank square that has no adjacent mines (i.e., above, below, left, right, and all 4 diagonals),
//     digit ('1' to '8') represents how many mines are adjacent to this revealed square, and
//     'X' represents a revealed mine.

// You are also given an integer array click where click = [clickr, clickc] represents the next click position among all the unrevealed squares ('M' or 'E').
// Return the board after revealing this position according to the following rules:
//     If a mine 'M' is revealed, then the game is over. You should change it to 'X'.
//     If an empty square 'E' with no adjacent mines is revealed, then change it to a revealed blank 'B' and all of its adjacent unrevealed squares should be revealed recursively.
//     If an empty square 'E' with at least one adjacent mine is revealed, then change it to a digit ('1' to '8') representing the number of adjacent mines.

// Return the board when no more squares will be revealed.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/09/untitled.jpeg" />
// Input: board = [["E","E","E","E","E"],["E","E","M","E","E"],["E","E","E","E","E"],["E","E","E","E","E"]], click = [3,0]
// Output: [["B","1","E","1","B"],["B","1","M","1","B"],["B","1","1","1","B"],["B","B","B","B","B"]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/09/untitled-2.jpeg" />
// Input: board = [["B","1","E","1","B"],["B","1","M","1","B"],["B","1","1","1","B"],["B","B","B","B","B"]], click = [1,2]
// Output: [["B","1","E","1","B"],["B","1","X","1","B"],["B","1","1","1","B"],["B","B","B","B","B"]]
 
// Constraints:
//     m == board.length
//     n == board[i].length
//     1 <= m, n <= 50
//     board[i][j] is either 'M', 'E', 'B', or a digit from '1' to '8'.
//     click.length == 2
//     0 <= clickr < m
//     0 <= clickc < n
//     board[clickr][clickc] is either 'M' or 'E'.

import "fmt"

// bfs
func updateBoard(board [][]byte, click []int) [][]byte {
    row, col := len(board), len(board[0])
    ci, cj := click[0], click[1]
    cur := board[ci][cj]
    if cur == 'M' { // 直接中雷
        board[ci][cj] = 'X'
        return board
    }
    if (cur >= '1' && cur <= '8' || cur == 'B') && cur != 'E' {
        return board
    }
    q := [][2]int{ {ci, cj}}
    board[ci][cj] = 'B'
    var dirs = [][]int{ {-1, 0}, {1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1} } // 8个方向
    for len(q) > 0 {
        tmp := q
        q = nil
        for _, top := range tmp {
            hasMine := 0
            for _, dir := range dirs {
                x := top[0] + dir[0]
                y := top[1] + dir[1]
                if x >= 0 && x < row && y >= 0 && y < col {
                    if board[x][y] == 'M' {
                        hasMine += 1
                    }
                }
            }
            // for 8 dirs end check 
            if hasMine > 0 {
                board[top[0]][top[1]] = '0' + byte(hasMine)
            } else {
                for _, dir := range dirs {
                    x := top[0] + dir[0]
                    y := top[1] + dir[1]
                    if x >= 0 && x < row && y >= 0 && y < col && board[x][y] == 'E' {
                        board[x][y] = 'B'
                        q = append(q, [2]int{x, y})    
                    }
                }
            }
        }
    }
    return board
}

// dfs
const (
    unMine      = 'M'
    emptySquare = 'E'
    blank       = 'B'
    reMine      = 'X'
)
func updateBoard1(board [][]byte, click []int) [][]byte {
    x, y := click[0], click[1]
    if board[x][y] == unMine {
        board[x][y] = reMine
        return board
    }
    m, n := len(board), len(board[0])
    var getAdjMineCount func(i, j int) int
    getAdjMineCount = func(x, y int) int {
        var mineCount int
        i0, j0 := x-1, y-1
        for i := i0; i < i0+3; i++ {
            if i < 0 || i >= m {
                continue
            }
            for j := j0; j < j0+3; j++ {
                if j < 0 || j >= n {
                    continue
                }
                if board[i][j] == unMine {
                    mineCount++
                }
            }
        }
        return mineCount
    }
    var dfs func(x, y int)
    dfs = func(x, y int) {
        mineCount := getAdjMineCount(x, y)
        if mineCount > 0 {
            board[x][y] = byte('0' + mineCount)
            return
        }
        board[x][y] = blank
        i0, j0 := x-1, y-1
        for i := i0; i < i0+3; i++ {
            if i < 0 || i >= m {
                continue
            }
            for j := j0; j < j0+3; j++ {
                if j < 0 || j >= n {
                    continue
                }
                if board[i][j] == emptySquare {
                    dfs(i, j)
                }
            }
        }
    }
    dfs(x, y)
    return board
}


func main() {
    board := [][]byte{
        {'E','E','E','E','E'},
        {'E','E','M','E','E'},
        {'E','E','E','E','E'},
        {'E','E','E','E','E'},
    }
    fmt.Println(updateBoard(board,[]int{3,0}))
    fmt.Println(updateBoard(board,[]int{1,2}))

    fmt.Println(updateBoard1(board,[]int{3,0}))
    fmt.Println(updateBoard1(board,[]int{1,2}))
}