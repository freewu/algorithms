package main

// 51. N-Queens
// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle. 
// You may return the answer in any order.

// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/13/queens.jpg" />
// Input: n = 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

// Example 2:
// Input: n = 1
// Output: [["Q"]]
 
// Constraints:
//     1 <= n <= 9
 
import "fmt"

// dfs
func solveNQueens(n int) [][]string {
    col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, [][]string{}
    // 生成棋盘
    generateBoard := func (n int, row *[]int) []string {
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
    // 尝试在一个n皇后问题中, 摆放第index行的皇后位置
    var putQueen func(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string)
    putQueen = func(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *[][]string) {
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
    putQueen(n, 0, &col, &dia1, &dia2, &row, &res)
    return res
}

// 二进制
func solveNQueens1(n int) [][]string {
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
func solveNQueens2(n int) [][]string {
	res := [][]string{}
    // 生成棋盘
    makeBoard := func (n int) [][]string {
        board := make([][]string,n)
        for i:=0; i < len(board);i++{
            row:= make([]string,n)
            for j:=0; j < len(row); j++{
                row[j] = "."
            }
            board[i] = row
        }
        return board
    }
    // 判断是否符合规则
    isValid := func (row,col int,board [][]string) bool{
        for i:= 0; i < row; i++{
            if board[i][col] == "Q" {
                return false
            }
        }
        // 左上角45度 /
        for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i - 1,j - 1 {
            if board[i][j] == "Q" {
                return false
            }
        }
        // 右上角45度 \
        for i, j := row-1, col + 1; i >= 0 && j < len(board[0]); i, j = i - 1, j + 1 {
            if board[i][j] == "Q" {
                return false
            }
        }
        return true
    }
    var dfs func(board [][]string, row int, res *[][]string)
    dfs = func(board [][]string, row int, res *[][]string) {
        if row == len(board) {
            temp := make([]string,len(board))
            for row := 0; row < len(board); row++ {
                rowStr := ""
                for col := 0; col < len(board[0]); col++ {
                    rowStr += board[row][col]
                }
                temp[row] = rowStr
            }
            *res = append(*res, temp)
        }
        for col := 0; col < len(board); col++ {
            if isValid(row,col,board) {
                board[row][col] = "Q"
                dfs(board,row+1,res)
                board[row][col] = "."
            }
        }
    }
    board := makeBoard(n)
    dfs(board,0, &res)
    return res
}

// 判断当前位置是否可以放置皇后
func isValid(board [][]string, row, col int) bool {
    // 检查列是否有皇后互相冲突
    for i := 0; i < row; i++ {
        if board[i][col] == "Q" {
            return false
        }
    }
    // 检查左上方是否有皇后互相冲突
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    // 检查右上方是否有皇后互相冲突
    for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    return true
}

// 回溯函数
func backtrackNhuanghou(board [][]string, row int, res *[][]string) {
    // 触发结束条件
    if row == len(board) {
        temp := make([]string, len(board))
        for i := 0; i < len(board); i++ {
            temp[i] = strings.Join(board[i], "")
        }
        *res = append(*res, temp)
        return
    }
    // 穷举每一列
    for col := 0; col < len(board[row]); col++ {
        // 排除不合法选择
        if !isValid(board, row, col) {
            continue
        }
        // 做选择
        board[row][col] = "Q"
        // 进入下一行决策
        backtrackNhuanghou(board, row+1, res)
        // 撤销选择
        board[row][col] = "."
    }
}

// 主函数，用于解决N皇后问题
func solveNQueens3(n int) [][]string {
    // '.' 表示空，'Q' 表示皇后，初始化空棋盘。
    board := make([][]string, n)
    for i := 0; i < n; i++ {
        board[i] = make([]string, n)
        for j := 0; j < n; j++ {
            board[i][j] = "."
        }
    }
    var res [][]string
    backtrackNhuanghou(board, 0, &res)
    return res
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

    fmt.Printf("solveNQueens1(1) = %v\n",solveNQueens1(1))
    fmt.Printf("solveNQueens1(2) = %v\n",solveNQueens1(2))
    fmt.Printf("solveNQueens1(3) = %v\n",solveNQueens1(3))
    fmt.Printf("solveNQueens1(4) = %v\n",solveNQueens1(4))
    fmt.Printf("solveNQueens1(5) = %v\n",solveNQueens1(5))
    fmt.Printf("solveNQueens1(6) = %v\n",solveNQueens1(6))
    fmt.Printf("solveNQueens1(7) = %v\n",solveNQueens1(7))
    fmt.Printf("solveNQueens1(8) = %v\n",solveNQueens1(8))

    fmt.Printf("solveNQueens3(1) = %v\n",solveNQueens3(1))
    fmt.Printf("solveNQueens3(2) = %v\n",solveNQueens3(2))
    fmt.Printf("solveNQueens3(3) = %v\n",solveNQueens3(3))
    fmt.Printf("solveNQueens3(4) = %v\n",solveNQueens3(4))
    fmt.Printf("solveNQueens3(5) = %v\n",solveNQueens3(5))
    fmt.Printf("solveNQueens3(6) = %v\n",solveNQueens3(6))
    fmt.Printf("solveNQueens3(7) = %v\n",solveNQueens3(7))
    fmt.Printf("solveNQueens3(8) = %v\n",solveNQueens3(8))
}