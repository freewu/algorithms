package main

// 289. Game of Life 
// According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

// The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). 
// Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):
//     Any live cell with fewer than two live neighbors dies as if caused by under-population.
//     Any live cell with two or three live neighbors lives on to the next generation.
//     Any live cell with more than three live neighbors dies, as if by over-population.
//     Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

// The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. 
// Given the current state of the m x n grid board, return the next state.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/26/grid1.jpg" />
// Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/26/grid2.jpg" />
// Input: board = [[1,1],[1,0]]
// Output: [[1,1],[1,1]]
 
// Constraints:
//     m == board.length
//     n == board[i].length
//     1 <= m, n <= 25
//     board[i][j] is 0 or 1.

// Follow up:
//     Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
//     In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?

import "fmt"

func gameOfLife(board [][]int)  {
    rows := len(board)
    cols := len(board[0])
    for r := range board {
        for c := range board[r] {
            var neighbors int
            for nr := max(0, r-1); nr < min(rows, r+2); nr++ {
                for nc := max(0, c-1); nc < min(cols, c+2); nc++ {
                    if nr == r && nc == c {
                        continue
                    }
                    // the cell may already hold the second bit,
                    // so `if board[nr][nc] == 1` doesn't work
                    if board[nr][nc] & 1 == 1 {
                        neighbors++
                    }
                }
            }
            if neighbors == 3 || board[r][c] & 1 != 0 && neighbors ==2 {
                board[r][c] |= 1<<1
            }
        }
    }
    for r := range board {
        for c := range board[r] {
            board[r][c] = board[r][c]>>1
        }
    }
}

func gameOfLife1(board [][]int)  {
    rows :=  len(board)
    cols := len(board[0])
    round := []int{-1,0,1}
    for row:=0;row<rows;row++{
        for col := 0; col < cols; col++ {
            lives := 0
            for i:=0;i<3;i++{
                for j:=0 ;j<3;j++{
                    if !(round[i] == 0&& round[j]==0){
                        r := row + round[i]
                        c := col + round[j]
                        if (r<rows && r>=0)&&(c < cols && c >= 0) && (board[r][c] == 1||board[r][c]==-1){
                            lives += 1
                        }
                    }
                }
            }
            if board[row][col] == 1 && (lives<2 ||lives>3){
                board[row][col] = -1
            }
            if (board[row][col] == 0 && lives == 3) {
                    // 2 代表这个细胞过去是死的现在活了
                    board[row][col] = 2;
            }

        }
    }
    for row := 0;row<rows;row++{
        for  col := 0;col<cols;col++{
            if (board[row][col] > 0) {
                board[row][col] = 1;
            } else {
                board[row][col] = 0;
            }
        }
    }
}

func main() {
    matrix1 := [][]int{
        []int{0,1,0},
        []int{0,0,1},
        []int{1,1,1},
        []int{0,0,0},
    }
    fmt.Println(matrix1) // [[0 1 0] [0 0 1] [1 1 1] [0 0 0]]
    gameOfLife(matrix1)
    fmt.Println(matrix1) // [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

    matrix2 := [][]int{
        []int{1,1},
        []int{1,0},
    }
    fmt.Println(matrix2) // [[1 1] [1 0]]
    gameOfLife(matrix2)
    fmt.Println(matrix2) // [[1,1],[1,1]]

    matrix11 := [][]int{
        []int{0,1,0},
        []int{0,0,1},
        []int{1,1,1},
        []int{0,0,0},
    }
    fmt.Println(matrix11) // [[0 1 0] [0 0 1] [1 1 1] [0 0 0]]
    gameOfLife1(matrix11)
    fmt.Println(matrix11) // [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

    matrix21 := [][]int{
        []int{1,1},
        []int{1,0},
    }
    fmt.Println(matrix21) // [[1 1] [1 0]]
    gameOfLife1(matrix21)
    fmt.Println(matrix21) // [[1,1],[1,1]]
}