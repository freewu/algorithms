package main

// 37. Sudoku Solver

// Write a program to solve a Sudoku puzzle by filling the empty cells.

// A sudoku solution must satisfy all of the following rules:
//     1. Each of the digits 1-9 must occur exactly once in each row.
//     2. Each of the digits 1-9 must occur exactly once in each column.
//     3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

// The '.' character indicates empty cells.

// Example 1:
// <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png" />
// Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
// Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
// Explanation: The input board is shown above and the only valid solution is shown below:
// <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png">

// Constraints:
//     board.length == 9
//     board[i].length == 9
//     board[i][j] is a digit or '.'.
//     It is guaranteed that the input board has only one solution.

// 解题思路
//     DFS 暴力回溯枚举
//     数独要求每横行，每竖行，每九宫格内，1-9 的数字不能重复，每次放下一个数字的时候，在这 3 个地方都需要判断一次。

import "fmt"

type Position struct {
    x int
    y int
}

func solveSudoku(board [][]byte) {
    pos, find := []Position{}, false
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == '.' {
                pos = append(pos, Position{x: i, y: j})
            }
        }
    }
    putSudoku(&board, pos, 0, &find)
}

func putSudoku(board *[][]byte, pos []Position, index int, succ *bool) {
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

func checkSudoku(board *[][]byte, pos Position, val int) bool {
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
func solveSudoku1(board [][]byte) {
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

func solveSudoku2(board [][]byte) {
    // Bitmasks for used digits in rows, columns, and 3x3 boxes.
    // Bit d (1<<d) indicates digit d is present. We use bits 1..9.
    var row, col, box [9]int
    // Collect empty cells as indices r*9 + c
    empties := make([]int, 0, 81)
    // Initialize masks and empties
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            ch := board[r][c]
            if ch == '.' {
                empties = append(empties, r*9+c)
            } else {
                d := int(ch - '0')
                b := (r/3)*3 + (c / 3)
                bit := 1 << d
                row[r] |= bit
                col[c] |= bit
                box[b] |= bit
            }
        }
    }
    // Precompute bit -> digit (1<<d) -> d
    var bitToDigit [1 << 10]int
    for d := 1; d <= 9; d++ {
        bitToDigit[1<<d] = d
    }
    // Popcount using Kernighan’s trick
    popcount := func(x int) int {
        c := 0
        for x != 0 {
            x &= x - 1
            c++
        }
        return c
    }
    // Get candidate bitmask for cell (r,c)
    // Valid bits are from 1..9 => mask 0b1111111110 == 0x3FE
    candidates := func(r, c int) int {
        b := (r/3)*3 + (c / 3)
        used := row[r] | col[c] | box[b]
        return (^used) & 0x3FE
    }
    // Choose the empty cell with the fewest candidates (MRV).
    // Returns index into empties slice, or -1 if none (solved).
    selectMRV := func() int {
        bestIdx := -1
        bestCount := 10 // larger than any possible count
        for i := 0; i < len(empties); i++ {
            pos := empties[i]
            r, c := pos/9, pos%9
            cand := candidates(r, c)
            cnt := popcount(cand)
            if cnt < bestCount {
                bestCount = cnt
                bestIdx = i
                if cnt == 1 {
                    break // optimal choice found
                }
            }
        }
        return bestIdx
    }
    // Depth-first search with backtracking.
    var dfs func() bool
    dfs = func() bool {
        if len(empties) == 0 {
            return true // solved
        }
        // Pick most constrained empty cell
        m := selectMRV()
        if m == -1 {
            return true // no empties => solved
        }
        // Work on this cell; remove from empties by swapping with last
        pos := empties[m]
        r, c := pos/9, pos%9
        // Move chosen position to the end for efficient removal/restore
        empties[m], empties[len(empties)-1] = empties[len(empties)-1], empties[m]
        empties = empties[:len(empties)-1]
        b := (r/3)*3 + (c / 3)
        cand := candidates(r, c)
        // Try each candidate digit (iterate set bits)
        for x := cand; x != 0; x &= x - 1 {
            bit := x & -x
            d := bitToDigit[bit]
            // Place digit
            board[r][c] = byte('0' + d)
            row[r] |= bit
            col[c] |= bit
            box[b] |= bit
            if dfs() {
                return true
            }
            // Undo
            board[r][c] = '.'
            row[r] &^= bit
            col[c] &^= bit
            box[b] &^= bit
        }
        // Restore the empty slot
        empties = append(empties, pos)
        // Put it back where it was to keep MRV scanning predictable (optional)
        empties[len(empties)-1], empties[m] = empties[m], empties[len(empties)-1]
        return false
    }
    dfs()
}

func main() {
    // Example 1:
    // <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png" />
    // Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
    // Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
    // Explanation: The input board is shown above and the only valid solution is shown below:
    // <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png">
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
    solveSudoku(bytes)
    //solveSudoku(bytes)
    for i := range bytes {
        fmt.Printf("%c\n",bytes[i])
    }

    bytes11 := [][]byte{
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
    fmt.Println("sudoku problem1:")
    for i := range bytes11 {
        fmt.Printf("%c\n",bytes11[i])
    }
    fmt.Println("solution1:")
    solveSudoku1(bytes11)
    //solveSudoku(bytes)
    for i := range bytes11 {
        fmt.Printf("%c\n",bytes11[i])
    }

    bytes21 := [][]byte{
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
    fmt.Println("sudoku problem2:")
    for i := range bytes21 {
        fmt.Printf("%c\n",bytes21[i])
    }
    fmt.Println("solution2:")
    solveSudoku2(bytes21)
    //solveSudoku(bytes)
    for i := range bytes21 {
        fmt.Printf("%c\n",bytes21[i])
    }
}