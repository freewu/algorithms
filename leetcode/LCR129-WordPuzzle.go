package main

// LCR 129. 字母迷宫
// 字母迷宫游戏初始界面记作 m x n 二维字符串数组 grid，请判断玩家是否能在 grid 中找到目标单词 target。
// 注意：寻找单词时 必须 按照字母顺序，通过水平或垂直方向相邻的单元格内的字母构成，同时，同一个单元格内的字母 不允许被重复使用 。
// <img src="https://assets.leetcode.com/uploads/2020/11/04/word2.jpg" />

// 示例 1：
// 输入：grid = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], target = "ABCCED"
// 输出：true

// 示例 2：
// 输入：grid = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], target = "SEE"
// 输出：true

// 示例 3：
// 输入：grid = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], target = "ABCB"
// 输出：false

// 提示：
//     m == grid.length
//     n = grid[i].length
//     1 <= m, n <= 6
//     1 <= target.length <= 15
//     grid 和 target 仅由大小写英文字母组成

// 解题思路:
//     在地图上的任意一个起点开始，
//     向 4 个方向分别 DFS 搜索，
//     直到所有的单词字母都找到了就输出 true，否则输出 false

import "fmt"

func wordPuzzle(board [][]byte, word string) bool {
    dir := [][]int{ {-1, 0}, {0, 1}, {1, 0}, {0, -1} }
    visited := make([][]bool, len(board))
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, len(board[0]))
    }
    isInBoard := func (board [][]byte, x, y int) bool { 
        return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
    }
    var dfs func(index, x, y int) bool
    dfs = func(index, x, y int) bool {
        if index == len(word)-1 {
            return board[x][y] == word[index]
        }
        if board[x][y] == word[index] {
            visited[x][y] = true
            for i := 0; i < 4; i++ {
                nx := x + dir[i][0]
                ny := y + dir[i][1]
                if isInBoard(board, nx, ny) && !visited[nx][ny] && dfs(index + 1, nx, ny) {
                    return true
                }
            }
            visited[x][y] = false
        }
        return false
    }
    for i, v := range board {
        for j := range v {
            if dfs(0, i, j) {
                return true
            }
        }
    }
    return false
}

// best solution dfs
func wordPuzzle1(board [][]byte, word string) bool {
    type Cell struct {
        Row, Col int
    }
    direction := []Cell{{0,1},{0,-1},{1,0},{-1,0}}
    n, m := len(board), len(board[0])
    if len(word) > m*n {
        return false
    }
    counter := map[byte]int{} // counter for pruning
    for i := range word { // count char freq in words
        counter[word[i]]++
    }
    explored := make([][]bool, n)
    for i := range explored {
        explored[i] = make([]bool, m)
        for j := range explored[i] {
            _, ok := counter[board[i][j]];if ok{
                counter[board[i][j]]--
            }
        }
    }
    for _, count := range counter {  // if counter has a negative number
        if count > 0 {              // it means there not enough chars in board
            return false            // to search the word
        }
    }
    var dfs func(int, Cell) bool
    dfs = func(i int, v Cell) bool {
        if i == len(word)-1 {
            return true
        }
        explored[v.Row][v.Col] = true
        for _, dir := range direction {
            u := Cell{v.Row + dir.Row, v.Col + dir.Col}
            if u.Row >= 0 && u.Row < n && u.Col >= 0 && u.Col < m {
                if !explored[u.Row][u.Col] && board[u.Row][u.Col] == word[i+1] {
                    if dfs(i+1, u) {
                        return true
                    } else {
                        explored[u.Row][u.Col] = false // backtrack
                    }
                }
            }
        }
        return false
    }
    for i := range board {
        for j := range board[i] {
            if board[i][j] == word[0] {
                if dfs(0, Cell{i, j}) {
                    return true
                } else {
                    explored[i][j] = false // backtrack
                }
            }
        }
    }
    return false
}

// bfs
func wordPuzzle2(board [][]byte, word string) bool {
    var backtrack func (i, j int, word string, board [][]byte) bool
    backtrack = func (i, j int, word string, board [][]byte) bool {
        if len(word) == 0 {
            return true
        }
        // boundary check
        if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
            return false
        }
        if board[i][j] == word[0] {
            temp := board[i][j]
            board[i][j] = '0'
            if backtrack(i, j+1, word[1:], board) || backtrack(i, j-1, word[1:], board) || backtrack(i+1, j, word[1:], board) || backtrack(i-1, j, word[1:], board) {
                return true
            }
            board[i][j] = temp
        }
        return false
    }
    // bfs approach?
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] {
                temp := board[i][j]
                board[i][j] = '0'
                if backtrack(i, j+1, word[1:], board) || backtrack(i, j-1, word[1:], board) || backtrack(i+1, j, word[1:], board) || backtrack(i-1, j, word[1:], board) {
                    return true
                }
                board[i][j] = temp
            }
        }
    }
    return false
}

func main() {
    bytes := [][]byte{
        {'A', 'B', 'C', 'E'},
        {'S', 'F', 'C', 'S'}, 
        {'A', 'D', 'E', 'E'},
    }
    fmt.Printf("wordPuzzle(bytes,\"ABCCED\") = %v\n",wordPuzzle(bytes,"ABCCED")) // true
    fmt.Printf("wordPuzzle(bytes,\"SEE\") = %v\n",wordPuzzle(bytes,"SEE")) // true
    fmt.Printf("wordPuzzle(bytes,\"ABCB\") = %v\n",wordPuzzle(bytes,"ABCB")) // false

    fmt.Printf("wordPuzzle1(bytes,\"ABCCED\") = %v\n",wordPuzzle1(bytes,"ABCCED")) // true
    fmt.Printf("wordPuzzle1(bytes,\"SEE\") = %v\n",wordPuzzle1(bytes,"SEE")) // true
    fmt.Printf("wordPuzzle1(bytes,\"ABCB\") = %v\n",wordPuzzle1(bytes,"ABCB")) // false

    fmt.Printf("wordPuzzle2(bytes,\"ABCCED\") = %v\n",wordPuzzle2(bytes,"ABCCED")) // true
    fmt.Printf("wordPuzzle2(bytes,\"SEE\") = %v\n",wordPuzzle2(bytes,"SEE")) // true
    fmt.Printf("wordPuzzle2(bytes,\"ABCB\") = %v\n",wordPuzzle2(bytes,"ABCB")) // false
}
