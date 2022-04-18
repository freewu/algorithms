package main

import "fmt"

/**
130. Surrounded Regions
Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.
A region is captured by flipping all 'O's into 'X's in that surrounded region.

Constraints:

	m == board.length
	n == board[i].length
	1 <= m, n <= 200
	board[i][j] is 'X' or 'O'.

Example 1:

	X X X X
	X O O X
	X X O X
	X O X X

	After running your function, the board should be:

	X X X X
	X X X X
	X X X X
	X O X X

	Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
	Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
	Explanation: Surrounded regions should not be on the border,
	which means that any 'O' on the border of the board are not flipped to 'X'.
	Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
	Two cells are connected if they are adjacent cells connected horizontally or vertically.

Example 2:

	Input: board = [["X"]]
	Output: [["X"]]

解题思路
	给出一张二维地图，要求把地图上非边缘上的 ‘O’ 都用 ‘X’ 替换掉

 */

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// DFS
func solve(board [][]byte) {
	for i := range board {
		//fmt.Printf("before: %v\n",board[i])
		for j := range board[i] {
			if i == 0 || i == len(board) - 1 || j == 0 || j == len(board[i]) - 1 { // 处理外围一圈 如果是 O 先设置为 *
				 // fmt.Printf("board[%v][%v] = %v\n",i,j,board[i][j])
				if board[i][j] == 'O' {
					dfs130(i, j, board)
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

func dfs130(i, j int, board [][]byte) {
	if i < 0 || i > len(board) - 1 || j < 0 || j > len(board[i])-1 {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '*'
		// 处理四周
		for k := 0; k < 4; k++ {
			dfs130(i+ dir[k][0], j + dir[k][1], board)
		}
	}
}

func main() {
	board := [][]byte{{'X'}}
	fmt.Printf("board before solve = %v\n",board) // [[88]]
	solve(board)
	fmt.Printf("board after solve = %v\n",board) // [[88]]

	board1 := [][]byte{{'X','X','X','X'},{'X','O','O','X'},{'X','X','O','X'},{'X','O','X','X'}}
	fmt.Printf("board before solve = %v\n",board1) // [[88 88 88 88] [88 79 79 88] [88 88 79 88] [88 79 88 88]]
	solve(board1)
	fmt.Printf("board after solve = %v\n",board1) // [[88 88 88 88] [88 88 88 88] [88 88 88 88] [88 79 88 88]]
}