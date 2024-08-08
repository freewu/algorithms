package main

// 794. Valid Tic-Tac-Toe State
// Given a Tic-Tac-Toe board as a string array board, return true if and only if it is possible to reach this board position during the course of a valid tic-tac-toe game.
// The board is a 3 x 3 array that consists of characters ' ', 'X', and 'O'. 
// The ' ' character represents an empty square.

// Here are the rules of Tic-Tac-Toe:
//     Players take turns placing characters into empty squares ' '.
//     The first player always places 'X' characters, while the second player always places 'O' characters.
//     'X' and 'O' characters are always placed into empty squares, never filled ones.
//     The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
//     The game also ends if all squares are non-empty.
//     No more moves can be played if the game is over.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/15/tictactoe1-grid.jpg" />
// Input: board = ["O  ","   ","   "]
// Output: false
// Explanation: The first player always plays "X".

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/15/tictactoe2-grid.jpg" />
// Input: board = ["XOX"," X ","   "]
// Output: false
// Explanation: Players take turns making moves.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/05/15/tictactoe4-grid.jpg" />
// Input: board = ["XOX","O O","XOX"]
// Output: true

// Constraints:
//     board.length == 3
//     board[i].length == 3
//     board[i][j] is either 'X', 'O', or ' '.

import "fmt"

func validTicTacToe(board []string) bool {
    win := func(board []string, char string) bool {
        check := char + char + char
        for i := 0; i < 3; i++ {
            if board[i] == check { return true }
            col := string([]byte{board[0][i], board[1][i], board[2][i]})
            if col == check { return true }
        }
        diag1 := string([]byte{board[0][0], board[1][1], board[2][2]})
        diag2 := string([]byte{board[0][2], board[1][1], board[2][0]})
        if diag1 == check || diag2 == check { return true }
        return false
    }
    x, o, xWon, oWon := 0, 0, win(board, "X"), win(board, "O")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == 'X' {
                x++
            } else if board[i][j] == 'O' {
                o++
            }
        }
    }
    if o > x || x > o + 1 { return false }
    if xWon && (oWon || x == o) { return false }
    if oWon && x > o { return false }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/15/tictactoe1-grid.jpg" />
    // Input: board = ["O  ","   ","   "]
    // Output: false
    // Explanation: The first player always plays "X".
    fmt.Println(validTicTacToe([]string{"O  ","   ","   "})) // false
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/15/tictactoe2-grid.jpg" />
    // Input: board = ["XOX"," X ","   "]
    // Output: false
    // Explanation: Players take turns making moves.
    fmt.Println(validTicTacToe([]string{"XOX"," X ","   "})) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/05/15/tictactoe4-grid.jpg" />
    // Input: board = ["XOX","O O","XOX"]
    // Output: true
    fmt.Println(validTicTacToe([]string{"XOX","O O","XOX"})) // true
}