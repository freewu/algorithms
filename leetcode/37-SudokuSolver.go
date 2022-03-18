package main

import "fmt"

/**
37. Sudoku Solver

Write a program to solve a Sudoku puzzle by filling the empty cells.
A sudoku solution must satisfy all of the following rules:

	Each of the digits 1-9 must occur exactly once in each row.
	Each of the digits 1-9 must occur exactly once in each column.
	Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
	Empty cells are indicated by the character '.'.


解题思路
	DFS 暴力回溯枚举
	数独要求每横行，每竖行，每九宫格内，1-9 的数字不能重复，每次放下一个数字的时候，在这 3 个地方都需要判断一次。
*/

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) {
	pos, find := []position{}, false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{x: i, y: j})
			}
		}
	}
	putSudoku(&board, pos, 0, &find)
}

func putSudoku(board *[][]byte, pos []position, index int, succ *bool) {
	if *succ == true {
		return
	}
	if index == len(pos) {
		*succ = true
		return
	}
	for i := 1; i < 10; i++ {
		if checkSudoku(board, pos[index], i) && !*succ {
			(*board)[pos[index].x][pos[index].y] = byte(i) + '0'
			putSudoku(board, pos, index+1, succ)
			if *succ == true {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'
		}
	}
}

func checkSudoku(board *[][]byte, pos position, val int) bool {
	// 判断横行是否有重复数字
	for i := 0; i < len((*board)[0]); i++ {
		if (*board)[pos.x][i] != '.' && int((*board)[pos.x][i]-'0') == val {
			return false
		}
	}
	// 判断竖行是否有重复数字
	for i := 0; i < len((*board)); i++ {
		if (*board)[i][pos.y] != '.' && int((*board)[i][pos.y]-'0') == val {
			return false
		}
	}
	// 判断九宫格是否有重复数字
	posx, posy := pos.x-pos.x%3, pos.y-pos.y%3
	for i := posx; i < posx+3; i++ {
		for j := posy; j < posy+3; j++ {
			if (*board)[i][j] != '.' && int((*board)[i][j]-'0') == val {
				return false
			}
		}
	}
	return true
}

// best solution
func solveSudokuBest(board [][]byte) {
	var rows, cols, boxes [9][10]bool
	bI := func(row, col int) int {
		return (row/3)*3 + (col / 3)
	}
	solved := false

	var placeNumber func(d, row, col int)
	placeNumber = func(d, row, col int) {
		rows[row][d] = true
		cols[col][d] = true
		boxes[bI(row, col)][d] = true
		board[row][col] = Itoa(d)
	}

	canPlace := func(d, row, col int) bool {
		return !(rows[row][d] || cols[col][d] || boxes[bI(row, col)][d])
	}

	removeNumber := func(d, row, col int) {
		rows[row][d] = false
		cols[col][d] = false
		boxes[bI(row, col)][d] = false
		board[row][col] = '.'
	}

	var backtrack func(row, col int)
	placeNextNumbers := func(row, col int) {
		if row == 8 && col == 8 {
			solved = true
		} else {
			if col == 8 {
				backtrack(row+1, 0)
			} else {
				backtrack(row, col+1)
			}
		}
	}

	backtrack = func(row, col int) {
		if board[row][col] == '.' {
			for d := 1; d <= 9; d++ {
				if canPlace(d, row, col) {
					placeNumber(d, row, col)
					placeNextNumbers(row, col)
					if !solved {
						removeNumber(d, row, col)
					}
				}
			}
		} else {
			placeNextNumbers(row, col)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				placeNumber(Atoi(board[i][j]), i, j)
			}
		}
	}
	backtrack(0, 0)
}

func Atoi(b byte) int {
	return int(b - '0')
}

func Itoa(i int) byte {
	return '0' + byte(i)
}

func main() {
	bytes := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println("sudoku problem:")
	for i := range bytes {
		fmt.Printf("%c\n",bytes[i])
	}
	fmt.Println("solution:")
	solveSudokuBest(bytes)
	//solveSudoku(bytes)
	for i := range bytes {
		fmt.Printf("%c\n",bytes[i])
	}
}