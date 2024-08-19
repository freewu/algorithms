package main

// 999. Available Captures for Rook
// You are given an 8 x 8 matrix representing a chessboard. 
// There is exactly one white rook represented by 'R', some number of white bishops 'B', and some number of black pawns 'p'. 
// Empty squares are represented by '.'.

// A rook can move any number of squares horizontally or vertically (up, down, left, right) 
// until it reaches another piece or the edge of the board. 
// A rook is attacking a pawn if it can move to the pawn's square in one move.

// Note: A rook cannot move through other pieces, such as bishops or pawns. 
// This means a rook cannot attack a pawn if there is another piece blocking the path.

// Return the number of pawns the white rook is attacking.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/02/20/1253_example_1_improved.PNG" />
// Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
// Output: 3
// Explanation:
// In this example, the rook is attacking all the pawns.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/02/19/1253_example_2_improved.PNG" />
// Input: board = [[".",".",".",".",".",".","."],[".","p","p","p","p","p",".","."],[".","p","p","B","p","p",".","."],[".","p","B","R","B","p",".","."],[".","p","p","B","p","p",".","."],[".","p","p","p","p","p",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
// Output: 0
// Explanation:
// The bishops are blocking the rook from attacking any of the pawns.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/02/20/1253_example_3_improved.PNG" />
// Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","p",".",".",".","."],["p","p",".","R",".","p","B","."],[".",".",".",".",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."]]
// Output: 3
// Explanation:
// The rook is attacking the pawns at positions b5, d6, and f5.

// Constraints:
//     board.length == 8
//     board[i].length == 8
//     board[i][j] is either 'R', '.', 'B', or 'p'
//     There is exactly one cell with board[i][j] == 'R'

import "fmt"

func numRookCaptures(board [][]byte) int {
    res, y, x := 0, -1, -1
    for i := 0; i < 8; i++ { // 找到车的位置
        for j := 0; j < 8; j++ {
            if board[i][j] == 'R' {
                y = i
                x = j
                break
            }
        }
        if y != -1 { break }
    }
    i, j := y + 1, x // 向下 
    for i < 8 {
        if board[i][j] == 'B' { break } // 遇到象停止
        if board[i][j] == 'p' {
            res++
            break
        }
        i++ 
    }
    i, j = y, x - 1 // 向左
    for j >= 0 {
        if board[i][j] == 'B' { break } // 遇到象停止
        if board[i][j] == 'p' {
            res++
            break
        }
        j--
    }
    i, j = y - 1, x // 向上
    for i >= 0 {
        if board[i][j] == 'B' { break } // 遇到象停止
        if board[i][j] == 'p' {
            res++
            break
        }
        i--
    }
    i, j = y, x + 1 // 向右
    for j < 8 {
        if board[i][j] == 'B' { break } // 遇到象停止
        if board[i][j] == 'p' {
            res++
            break
        }
        j++
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/02/20/1253_example_1_improved.PNG" />
    // Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
    // Output: 3
    // Explanation:
    // In this example, the rook is attacking all the pawns.
    board1 := [][]byte{
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','p','.','.','.','.'},
        {'.','.','.','R','.','.','.','p'},
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','p','.','.','.','.'},
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','.','.','.','.','.'},
    }
    fmt.Println(numRookCaptures(board1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/02/19/1253_example_2_improved.PNG" />
    // Input: board = [[".",".",".",".",".",".","."],[".","p","p","p","p","p",".","."],[".","p","p","B","p","p",".","."],[".","p","B","R","B","p",".","."],[".","p","p","B","p","p",".","."],[".","p","p","p","p","p",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
    // Output: 0
    // Explanation:
    // The bishops are blocking the rook from attacking any of the pawns.
    board2 := [][]byte{
        {'.','.','.','.','.','.','.','.'},
        {'.','p','p','p','p','p','.','.'},
        {'.','p','p','B','p','p','.','.'},
        {'.','p','B','R','B','p','.','.'},
        {'.','p','p','B','p','p','.','.'},
        {'.','p','p','p','p','p','.','.'},
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','.','.','.','.','.'},
    }
    fmt.Println(numRookCaptures(board2)) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/02/20/1253_example_3_improved.PNG" />
    // Input: board = [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","p",".",".",".","."],["p","p",".","R",".","p","B","."],[".",".",".",".",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."]]
    // Output: 3
    // Explanation:
    // The rook is attacking the pawns at positions b5, d6, and f5.
    board3 := [][]byte{
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','p','.','.','.','.'},
        {'.','.','.','p','.','.','.','.'},
        {'p','p','.','R','.','p','B','.'},
        {'.','.','.','.','.','.','.','.'},
        {'.','.','.','B','.','.','.','.'},
        {'.','.','.','p','.','.','.','.'},
        {'.','.','.','.','.','.','.','.'},
    }
    fmt.Println(numRookCaptures(board3)) // 3
}