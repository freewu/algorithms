package main

// 723. Candy Crush
// This question is about implementing a basic elimination algorithm for Candy Crush.

// Given an m x n integer array board representing the grid of candy where board[i][j] represents the type of candy.
// A value of board[i][j] == 0 represents that the cell is empty.

// The given board represents the state of the game following the player's move. 
// Now, you need to restore the board to a stable state by crushing candies according to the following rules:
//     1. If three or more candies of the same type are adjacent vertically or horizontally, 
//     crush them all at the same time - these positions become empty.
//     2. After crushing all candies simultaneously, if an empty space on the board has candies on top of itself, 
//     then these candies will drop until they hit a candy or bottom at the same time. 
//     No new candies will drop outside the top boundary.
//     3. After the above steps, there may exist more candies that can be crushed. If so, you need to repeat the above steps.
//     4. If there does not exist more candies that can be crushed (i.e., the board is stable), then return the current board.

// You need to perform the above rules until the board becomes stable, then return the stable board.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/candy_crush_example_2.png" />
// Input: board = [[110,5,112,113,114],[210,211,5,213,214],[310,311,3,313,314],[410,411,412,5,414],[5,1,512,3,3],[610,4,1,613,614],[710,1,2,713,714],[810,1,2,1,1],[1,1,2,2,2],[4,1,4,4,1014]]
// Output: [[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[110,0,0,0,114],[210,0,0,0,214],[310,0,0,113,314],[410,0,0,213,414],[610,211,112,313,614],[710,311,412,613,714],[810,411,512,713,1014]]

// Example 2:
// Input: board = [[1,3,5,5,2],[3,4,3,3,1],[3,2,4,5,2],[2,4,4,5,5],[1,4,4,1,1]]
// Output: [[1,3,0,0,0],[3,4,0,5,2],[3,2,0,3,1],[2,4,0,5,2],[1,4,3,1,1]]

// Constraints:
//     m == board.length
//     n == board[i].length
//     3 <= m, n <= 50
//     1 <= board[i][j] <= 2000

import "fmt"

// 模拟
func candyCrush(board [][]int) [][]int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    m, n, erase := len(board), len(board[0]), true
    for erase {
        erase = false
        for _, line := range board { // 横向消除
            for j, num := range line {
                num = abs(num)
                if num > 0 && j > 1 {
                    if abs(line[j]) == abs(line[j-1]) && abs(line[j-1]) == abs(line[j-2]) {
                        line[j], line[j-1], line[j-2] = -num, -num, -num
                        erase = true
                    }
                }
            }
        }
        for j := 0; j < n; j++ { // 纵向消除
            for i := 0; i < m; i++ {
                num := abs(board[i][j])
                if num > 0 && i > 1 {
                    if abs(board[i][j]) == abs(board[i-1][j]) && abs(board[i-1][j]) == abs(board[i-2][j]) {
                        board[i][j], board[i-1][j], board[i-2][j] = -num, -num, -num
                        erase = true
                    }
                }
            }
        }
        for r := 0; r < n; r++ { // 下落
            up, dp := m-1, m-1
            for ; up >= 0; up-- {
                board[dp][r], board[up][r] = board[up][r], board[dp][r]
                if board[dp][r] > 0 {
                    dp--
                }
            }
            for ; dp >= 0; dp-- {
                board[dp][r] = 0
            }
        }
    }
    return board
}

func candyCrush1(board [][]int) [][]int {
    rl, cl, crushLimit := len(board), len(board[0]), 3
    vw := make([]int, 0, crushLimit)
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    tag := func(r, c int) { if board[r][c] < 0 { return; }; board[r][c] *= -1; }
    check := func(r, c int, isH bool) bool {
        needCrush, needCheck := false, true
        if len(vw) < crushLimit {
            vw = append(vw, board[r][c])
            if len(vw) != crushLimit {
                needCheck = false
            }
        } else {
            vw[0], vw[1], vw[2] = vw[1], vw[2], board[r][c]
        }
        if needCheck && vw[0] != 0 && abs(vw[0]) == abs(vw[1]) && abs(vw[0]) == abs(vw[2]) {
            needCrush = true
            if isH {
                tag(r, c)
                tag(r, c-1)
                tag(r, c-2)
            } else {
                tag(r, c)
                tag(r+1, c)
                tag(r+2, c)
            }
        }
        return needCrush
    }
    checkCrush := func() bool {
        needCrush := false
        // 检测时, 若遇到相邻超过3个相同值, 则取反
        for r := 0; r < rl; r++ { // 先横向检测
            vw = []int{}
            for c := 0; c < cl; c++ {
                if check(r, c, true) {
                    needCrush = true
                }
            }
        }
        for c := 0; c < cl; c++ { // 再纵向检测
            vw = []int{}
            for r := rl - 1; r >= 0; r-- {
                if board[r][c] == 0 {
                    break
                }
                if check(r, c, false) {
                    needCrush = true
                }
            }
        }
        return needCrush
    }
    reLayout := func() {
        for c := 0; c < cl; c++ { // 纵向遍历, 消除负值.
            r, write := rl - 1, rl - 1
            for ; r >= 0; r-- {
                if board[r][c] == 0 {
                    break
                }
                if board[r][c] > 0 {
                    board[write][c] = board[r][c]
                    write--
                }
            }
            // [write, r)
            for write > r {
                board[write][c] = 0
                write--
            }
        }
    }
    for {
        if !checkCrush() {
            break
        }
        reLayout()
    }
    return board
}

func main() {
    // Example 1:
    // Input: board = [[110,5,112,113,114],[210,211,5,213,214],[310,311,3,313,314],[410,411,412,5,414],[5,1,512,3,3],[610,4,1,613,614],[710,1,2,713,714],[810,1,2,1,1],[1,1,2,2,2],[4,1,4,4,1014]]
    // Output: [[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[110,0,0,0,114],[210,0,0,0,214],[310,0,0,113,314],[410,0,0,213,414],[610,211,112,313,614],[710,311,412,613,714],[810,411,512,713,1014]]
    board1 := [][]int{
        {110,5,112,113,114},
        {210,211,5,213,214},
        {310,311,3,313,314},
        {410,411,412,5,414},
        {5,1,512,3,3},
        {610,4,1,613,614},
        {710,1,2,713,714},
        {810,1,2,1,1},
        {1,1,2,2,2},
        {4,1,4,4,1014},
    }
    fmt.Println(candyCrush(board1)) // [[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[110,0,0,0,114],[210,0,0,0,214],[310,0,0,113,314],[410,0,0,213,414],[610,211,112,313,614],[710,311,412,613,714],[810,411,512,713,1014]]
    // Example 2:
    // Input: board = [[1,3,5,5,2],[3,4,3,3,1],[3,2,4,5,2],[2,4,4,5,5],[1,4,4,1,1]]
    // Output: [[1,3,0,0,0],[3,4,0,5,2],[3,2,0,3,1],[2,4,0,5,2],[1,4,3,1,1]]
    board2 := [][]int{
        {1,3,5,5,2},
        {3,4,3,3,1},
        {3,2,4,5,2},
        {2,4,4,5,5},
        {1,4,4,1,1},
    }
    fmt.Println(candyCrush(board2)) // [[1,3,0,0,0],[3,4,0,5,2],[3,2,0,3,1],[2,4,0,5,2],[1,4,3,1,1]]

    board11 := [][]int{
        {110,5,112,113,114},
        {210,211,5,213,214},
        {310,311,3,313,314},
        {410,411,412,5,414},
        {5,1,512,3,3},
        {610,4,1,613,614},
        {710,1,2,713,714},
        {810,1,2,1,1},
        {1,1,2,2,2},
        {4,1,4,4,1014},
    }
    fmt.Println(candyCrush1(board11)) // [[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[110,0,0,0,114],[210,0,0,0,214],[310,0,0,113,314],[410,0,0,213,414],[610,211,112,313,614],[710,311,412,613,714],[810,411,512,713,1014]]
    board12 := [][]int{
        {1,3,5,5,2},
        {3,4,3,3,1},
        {3,2,4,5,2},
        {2,4,4,5,5},
        {1,4,4,1,1},
    }
    fmt.Println(candyCrush1(board12)) // [[1,3,0,0,0],[3,4,0,5,2],[3,2,0,3,1],[2,4,0,5,2],[1,4,3,1,1]]
}