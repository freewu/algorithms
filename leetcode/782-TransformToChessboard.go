package main

// 782. Transform to Chessboard
// You are given an n x n binary grid board. 
// In each move, you can swap any two rows with each other, or any two columns with each other.

// Return the minimum number of moves to transform the board into a chessboard board. 
// If the task is impossible, return -1.

// A chessboard board is a board where no 0's and no 1's are 4-directionally adjacent.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/chessboard1-grid.jpg" />
// Input: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
// Output: 2
// Explanation: One potential sequence of moves is shown.
// The first move swaps the first and second column.
// The second move swaps the second and third row.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/chessboard2-grid.jpg" />
// Input: board = [[0,1],[1,0]]
// Output: 0
// Explanation: Also note that the board with 0 in the top left corner, is also a valid chessboard.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/chessboard3-grid.jpg" />
// Input: board = [[1,0],[1,0]]
// Output: -1
// Explanation: No matter what sequence of moves you make, you cannot end with a valid chessboard.

// Constraints:
//     n == board.length
//     n == board[i].length
//     2 <= n <= 30
//     board[i][j] is either 0 or 1.

import "fmt"

func movesToChessboard(board [][]int) int {
    n, rowSum, colSum, rowSwap, colSwap := len(board), 0, 0, 0, 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if board[0][0] ^ board[i][0] ^ board[0][j] ^ board[i][j] == 1 {
                return -1
            }
        }
    }
    for i := 0; i < n; i++ {
        rowSum += board[0][i]
        colSum += board[i][0]
        if board[i][0] == i % 2 {
            rowSwap++
        }
        if board[0][i] == i % 2 {
            colSwap++
        }
    }
    if rowSum != n/2 && rowSum != (n+1)/2 { return -1 }
    if colSum != n/2 && colSum != (n+1)/2 { return -1 }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    if n%2 == 1 {
        if colSwap%2 == 1 {
            colSwap = n - colSwap
        }
        if rowSwap%2 == 1 {
            rowSwap = n - rowSwap
        }
    } else {
        colSwap = min(n-colSwap, colSwap)
        rowSwap = min(n-rowSwap, rowSwap)
    }
    return (colSwap + rowSwap) / 2
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/chessboard1-grid.jpg" />
    // Input: board = [[0,1,1,0],[0,1,1,0],[1,0,0,1],[1,0,0,1]]
    // Output: 2
    // Explanation: One potential sequence of moves is shown.
    // The first move swaps the first and second column.
    // The second move swaps the second and third row.
    fmt.Println(movesToChessboard([][]int{{0,1,1,0},{0,1,1,0},{1,0,0,1},{1,0,0,1}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/chessboard2-grid.jpg" />
    // Input: board = [[0,1],[1,0]]
    // Output: 0
    // Explanation: Also note that the board with 0 in the top left corner, is also a valid chessboard.
    fmt.Println(movesToChessboard([][]int{{0,1},{1,0}})) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/chessboard3-grid.jpg" />
    // Input: board = [[1,0],[1,0]]
    // Output: -1
    // Explanation: No matter what sequence of moves you make, you cannot end with a valid chessboard.
    fmt.Println(movesToChessboard([][]int{{1,0},{1,0}})) // -1
}