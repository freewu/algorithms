package main

import "fmt"

/**
51. N-Queens
he n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle.
Each solution contains a distinct board configuration of the n-queens’ placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example 1:

	Input: n = 4
	Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
	Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

Example 2:

	Input: n = 1
	Output: [["Q"]]
 */

// 解法一 DFS
func solveNQueens(n int) [][]string {
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}
	putQueen(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

// 尝试在一个n皇后问题中, 摆放第index行的皇后位置
func putQueen(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string) {
	if index == n {
		*res = append(*res, generateBoard(n, row))
		return
	}
	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen(n, index+1, col, dia1, dia2, row, res)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}

// 解法二 二进制操作法
func solveNQueens2(n int) [][]string {
	res := [][]string{}
	placements := make([]string, n)
	for i := range placements {
		buf := make([]byte, n)
		for j := range placements {
			if i == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		placements[i] = string(buf)
	}
	var construct func(prev []int)
	construct = func(prev []int) {
		if len(prev) == n {
			plan := make([]string, n)
			for i := 0; i < n; i++ {
				plan[i] = placements[prev[i]]
			}
			res = append(res, plan)
			return
		}
		occupied := 0
		for i := range prev {
			dist := len(prev) - i
			bit := 1 << prev[i]
			occupied |= bit | bit << dist | bit >> dist
		}
		prev = append(prev, -1)
		for i := 0; i < n; i++ {
			if (occupied >> i ) &1 != 0 {
				continue
			}
			prev[len(prev)-1] = i
			construct(prev)
		}
	}
	construct(make([]int, 0, n))
	return res
}

// best solution
func solveNQueensBest(n int) [][]string {
	var res [][]string
	board := makeBoardBest(n)
	btBest(board,0,&res)
	return res
}

func btBest(board [][]string, row int, res *[][]string){
	if row==len(board) {
		temp := make([]string,len(board))

		for row:=0;row<len(board);row++{
			rowStr := ""
			for col:=0;col<len(board[0]);col++{
				rowStr += board[row][col]
			}
			temp[row] = rowStr
		}
		*res = append(*res,temp)
	}
	for col:=0;col<len(board);col++{
		if isValidBest(row,col,board) {
			board[row][col] = "Q"
			btBest(board,row+1,res)
			board[row][col] = "."
		}
	}
}

func isValidBest(row,col int,board [][]string) bool{
	for i:=0;i<row;i++{
		if board[i][col] == "Q"{
			return false
		}
	}
	//左上角45度
	for i,j:=row-1,col-1;i>=0&&j>=0;i,j=i-1,j-1{
		if board[i][j] == "Q"{
			return false
		}
	}
	//右上角45度
	for i,j:=row-1,col+1;i>=0&&j<len(board[0]);i,j=i-1,j+1{
		if board[i][j] == "Q"{
			return false
		}
	}
	return true
}

func makeBoardBest(n int)[][]string{
	board := make([][]string,n)
	for i:=0; i<len(board);i++{
		row:= make([]string,n)
		for j:=0;j<len(row);j++{
			row[j] = "."
		}
		board[i] = row
	}
	return board
}

func main() {
	fmt.Printf("solveNQueens(1) = %v\n",solveNQueens(1))
	fmt.Printf("solveNQueens(2) = %v\n",solveNQueens(2))
	fmt.Printf("solveNQueens(3) = %v\n",solveNQueens(3))
	fmt.Printf("solveNQueens(4) = %v\n",solveNQueens(4))
	fmt.Printf("solveNQueens(5) = %v\n",solveNQueens(5))
	fmt.Printf("solveNQueens(6) = %v\n",solveNQueens(6))
	fmt.Printf("solveNQueens(7) = %v\n",solveNQueens(7))
	fmt.Printf("solveNQueens(8) = %v\n",solveNQueens(8))

	fmt.Printf("solveNQueens2(1) = %v\n",solveNQueens2(1))
	fmt.Printf("solveNQueens2(2) = %v\n",solveNQueens2(2))
	fmt.Printf("solveNQueens2(3) = %v\n",solveNQueens2(3))
	fmt.Printf("solveNQueens2(4) = %v\n",solveNQueens2(4))
	fmt.Printf("solveNQueens2(5) = %v\n",solveNQueens2(5))
	fmt.Printf("solveNQueens2(6) = %v\n",solveNQueens2(6))
	fmt.Printf("solveNQueens2(7) = %v\n",solveNQueens2(7))
	fmt.Printf("solveNQueens2(8) = %v\n",solveNQueens2(8))

	fmt.Printf("solveNQueensBest(1) = %v\n",solveNQueensBest(1))
	fmt.Printf("solveNQueensBest(2) = %v\n",solveNQueensBest(2))
	fmt.Printf("solveNQueensBest(3) = %v\n",solveNQueensBest(3))
	fmt.Printf("solveNQueensBest(4) = %v\n",solveNQueensBest(4))
	fmt.Printf("solveNQueensBest(5) = %v\n",solveNQueensBest(5))
	fmt.Printf("solveNQueensBest(6) = %v\n",solveNQueensBest(6))
	fmt.Printf("solveNQueensBest(7) = %v\n",solveNQueensBest(7))
	fmt.Printf("solveNQueensBest(8) = %v\n",solveNQueensBest(8))
}