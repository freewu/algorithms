package main

// 52. N-Queens II
// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/13/queens.jpg" />
// Input: n = 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as shown.

// Example 2:
// Input: n = 1
// Output: 1
 
// Constraints:
//     1 <= n <= 9

import "fmt"

func totalNQueens(n int) int {
    var backtrack func(n, row, columns, diagonals, antiDiagonals int) int 
    backtrack = func(n, row, columns, diagonals, antiDiagonals int) int {
        if n == row {
            return 1
        }
        res := 0
        for c := 0; c < n; c++ {
            column, diagonal, antiDiagonal := 1 << c, 1 << (row - c + n), 1 << (row + c)
            if columns & column > 0 ||
                diagonals & diagonal > 0 ||
                antiDiagonals & antiDiagonal > 0 {
                continue
            }
            columns ^= column
            diagonals ^= diagonal
            antiDiagonals ^= antiDiagonal
            res += backtrack(n, row + 1, columns, diagonals, antiDiagonals)
            columns ^= column
            diagonals ^= diagonal
            antiDiagonals ^= antiDiagonal
        }
        return res
    }
    return backtrack(n, 0, 0, 0, 0)
}

// dfs
func totalNQueens1(n int) int {
    res, col, used := 0, make([]int, n), make([]bool, n)
    diag1, diag2 := make([]bool, n * 2 - 1), make([]bool, n * 2 - 1)
    var dfs func(int)
    dfs = func(r int) {
        if r == n {
            res++
            return
        }
        for c, on := range used {
            rc := r - c + n - 1
            if !on && !diag1[r + c] && !diag2[rc] {
                col[r] = c
                used[c], diag1[r + c], diag2[rc] = true, true, true
                dfs(r + 1)
                used[c], diag1[r + c], diag2[rc] = false, false, false
            }
        }
    }
    dfs(0)
    return res
}

func main() {
    fmt.Println(totalNQueens(1)) // 1
    fmt.Println(totalNQueens(2)) // 0
    fmt.Println(totalNQueens(3)) // 0
    fmt.Println(totalNQueens(4)) // 2
    fmt.Println(totalNQueens(5)) // 10
    fmt.Println(totalNQueens(6)) // 4
    fmt.Println(totalNQueens(7)) // 40
    fmt.Println(totalNQueens(8)) // 92
    fmt.Println(totalNQueens(9)) // 352

    fmt.Println(totalNQueens1(1)) // 1
    fmt.Println(totalNQueens1(2)) // 0
    fmt.Println(totalNQueens1(3)) // 0
    fmt.Println(totalNQueens1(4)) // 2
    fmt.Println(totalNQueens1(5)) // 10
    fmt.Println(totalNQueens1(6)) // 4
    fmt.Println(totalNQueens1(7)) // 40
    fmt.Println(totalNQueens1(8)) // 92
    fmt.Println(totalNQueens1(9)) // 352
}
