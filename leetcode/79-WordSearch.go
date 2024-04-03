package main

// 79. Word Search
// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells, 
// where adjacent cells are horizontally or vertically neighboring. 
// The same letter cell may not be used more than once.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/04/word2.jpg" />
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg" />
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// Output: true

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/15/word3.jpg" />
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// Output: false

// Constraints:
//     m == board.length
//     n = board[i].length
//     1 <= m, n <= 6
//     1 <= word.length <= 15
//     board and word consists of only lowercase and uppercase English letters.
    
// Follow up: Could you use search pruning to make your solution faster with a larger board?

// 解题思路:
//     在地图上的任意一个起点开始，
//     向 4 个方向分别 DFS 搜索，
//     直到所有的单词字母都找到了就输出 true，否则输出 false


import "fmt"

func exist(board [][]byte, word string) bool {
    dir := [][]int{ {-1, 0}, {0, 1}, {1, 0}, {0, -1} }
    visited := make([][]bool, len(board))
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, len(board[0]))
    }
    var dfs func(board [][]byte, visited [][]bool, word string, index, x, y int) bool
    dfs = func(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
        if index == len(word)-1 {
            return board[x][y] == word[index]
        }
        isInBoard := func (board [][]byte, x, y int) bool { return x >= 0 && x < len(board) && y >= 0 && y < len(board[0]); }
        if board[x][y] == word[index] {
            visited[x][y] = true
            for i := 0; i < 4; i++ {
                nx := x + dir[i][0]
                ny := y + dir[i][1]
                if isInBoard(board, nx, ny) && !visited[nx][ny] && dfs(board, visited, word, index+1, nx, ny) {
                    return true
                }
            }
            visited[x][y] = false
        }
        return false
    }
    for i, v := range board {
        for j := range v {
            if dfs(board, visited, word, 0, i, j) {
                return true
            }
        }
    }
    return false
}

// best solution dfs
type Cell struct {
    Row, Col int
}
var direction = []Cell{{0,1},{0,-1},{1,0},{-1,0}}
func exist1(board [][]byte, word string) bool {
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
func exist2(board [][]byte, word string) bool {
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
    fmt.Printf("exist(bytes,\"ABCCED\") = %v\n",exist(bytes,"ABCCED")) // true
    fmt.Printf("exist(bytes,\"SEE\") = %v\n",exist(bytes,"SEE")) // true
    fmt.Printf("exist(bytes,\"ABCB\") = %v\n",exist(bytes,"ABCB")) // false

    fmt.Printf("exist1(bytes,\"ABCCED\") = %v\n",exist1(bytes,"ABCCED")) // true
    fmt.Printf("exist1(bytes,\"SEE\") = %v\n",exist1(bytes,"SEE")) // true
    fmt.Printf("exist1(bytes,\"ABCB\") = %v\n",exist1(bytes,"ABCB")) // false

    fmt.Printf("exist2(bytes,\"ABCCED\") = %v\n",exist2(bytes,"ABCCED")) // true
    fmt.Printf("exist2(bytes,\"SEE\") = %v\n",exist2(bytes,"SEE")) // true
    fmt.Printf("exist2(bytes,\"ABCB\") = %v\n",exist2(bytes,"ABCB")) // false
}
