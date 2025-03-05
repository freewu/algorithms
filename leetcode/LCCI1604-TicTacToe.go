package main

// 面试题 16.04. Tic-Tac-Toe LCCI
// Design an algorithm to figure out if someone has won a game of tic-tac-toe. 
// Input is a string array of size N x N, including characters " ", "X" and "O", where " " represents a empty grid.

// The rules of tic-tac-toe are as follows:
//     1. Players place characters into an empty grid(" ") in turn.
//     2. The first player always place character "O", and the second one place "X".
//     3. Players are only allowed to place characters in empty grid. 
//        Replacing a character is not allowed.
//     4. If there is any row, column or diagonal filled with N same characters, the game ends. 
//        The player who place the last charater wins.
//     5. When there is no empty grid, the game ends.
//     6. If the game ends, players cannot place any character further.

// If there is any winner, return the character that the winner used. 
// If there's a draw, return "Draw". If the game doesn't end and there is no winner, return "Pending".

// Example 1:
// Input:  board = ["O X"," XO","X O"]
// Output:  "X"

// Example 2:
// Input:  board = ["OOX","XXO","OXO"]
// Output:  "Draw"
// Explanation:  no player wins and no empty grid left

// Example 3:
// Input:  board = ["OOX","XXO","OX "]
// Output:  "Pending"
// Explanation:  no player wins but there is still a empty grid

// Note:
//     1 <= board.length == board[i].length <= 100
//     Input follows the rules.

import "fmt"

func tictactoe(board []string) string {
    hbyte, zbyte := []byte{}, []byte{}
    x, o := "", "" // 根据棋盘大小设定赢棋长度
    for i := 0; i < len(board); i++ {
        x += "X"
        o += "O"
    }
    l1, l2, empty := []byte{}, []byte{},  false
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board); j++ {
            hbyte, zbyte = append(hbyte, board[i][j]), append(zbyte, board[j][i])
            if i == j {
                l1 = append(l1, board[i][i])
            }
            if i + j == len(board) -1 {
                l2 = append(l2, board[i][j])
            }
            if board[i][j] == ' ' { // 空格数
                empty = true
            } 
        }
        if x == string(hbyte) || x == string(zbyte) {
            return "X"
        }
        if o == string(hbyte) || o == string(zbyte) {
            return "O"
        }
        hbyte, zbyte = []byte{}, []byte{}
    }
    if x == string(l1) || x == string(l2) {
        return "X"
    }
    if o == string(l1) || o == string(l2) {
        return "O"
    }
    if empty { // 还有空棋盘
        return "Pending"
    }
    return "Draw"
}

func main() {
    // Example 1:
    // Input:  board = ["O X"," XO","X O"]
    // Output:  "X"
    fmt.Println(tictactoe([]string{"O X"," XO","X O"})) // "X"
    // Example 2:
    // Input:  board = ["OOX","XXO","OXO"]
    // Output:  "Draw"
    // Explanation:  no player wins and no empty grid left
    fmt.Println(tictactoe([]string{"OOX","XXO","OXO"})) // "Draw"
    // Example 3:
    // Input:  board = ["OOX","XXO","OX "]
    // Output:  "Pending"
    // Explanation:  no player wins but there is still a empty grid
    fmt.Println(tictactoe([]string{"OOX","XXO","OX "})) // "Pending"
}