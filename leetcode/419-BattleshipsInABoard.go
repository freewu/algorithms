package main

// 419. Battleships in a Board
// Given an m x n matrix board where each cell is a battleship 'X' or empty '.', 
// return the number of the battleships on board.

// Battleships can only be placed horizontally or vertically on board. 
// In other words, they can only be made of the shape 1 x k (1 row, k columns) or k x 1 (k rows, 1 column), where k can be of any size. 
// At least one horizontal or vertical cell separates between two battleships (i.e., there are no adjacent battleships).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/10/battelship-grid.jpg" />
// Input: board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
// Output: 2

// Example 2:
// Input: board = [["."]]
// Output: 0

// Constraints:
//     m == board.length
//     n == board[i].length
//     1 <= m, n <= 200
//     board[i][j] is either '.' or 'X'.
 
// Follow up: Could you do it in one-pass, using only O(1) extra memory and without modifying the values board?

import "fmt"

func countBattleships(board [][]byte) int {
    ship := 0
    for i, array := range board {
        for j, v := range array {
            if v == 'X' {
                switch {
                case i == 0 && j == 0: // edge, ship++
                    ship++
                case i != 0 && j == 0: // ↑ is not ship, ship++
                    if board[i-1][j] != 'X' {
                        ship++
                    }
                case i == 0 && j != 0:
                    if board[i][j-1] != 'X' { // ← is not ship, ship++
                        ship++
                    }
                case i != 0 && j != 0:
                    if board[i][j-1] != 'X' && board[i-1][j] != 'X' { // ↑ and ← both not ship, ship++
                        ship++
                    }
                }
            }
        }
    }
    return ship
}

func countBattleships1(board [][]byte) int {
    if len(board) == 0 {
        return 0
    }
    count := 0
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == 'X' {
                count++
                board[i][j] = 'x'
                if j+1 < len(board[0]) {
                    index := j + 1
                    for index < len(board[0]) {
                        if board[i][index] == 'X' {
                            board[i][index] = 'x'
                            index++
                        } else {
                            break
                        }
                    }
                }
                if i+1 < len(board) {
                    index := i + 1
                    for index < len(board) {
                        if board[index][j] == 'X' {
                            board[index][j] = 'x'
                            index++
                        } else {
                            break
                        }
                    }
                }
            }
        }
    }
    return count
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/10/battelship-grid.jpg" />
    // Input: board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
    // Output: 2
    board1 := [][]byte{
        {'X','.','.','X'},
        {'.','.','.','X'},
        {'.','.','.','X'},
    }
    fmt.Println(countBattleships(board1)) // 2
    // Example 2:
    // Input: board = [["."]]
    // Output: 0
    board2 := [][]byte{{'.'}}
    fmt.Println(countBattleships(board2)) // 0

    fmt.Println(countBattleships1(board1)) // 2
    fmt.Println(countBattleships1(board2)) // 0
}