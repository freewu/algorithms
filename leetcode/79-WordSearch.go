package main

import (
	"fmt"
)

/**
79. Word Search
Given an m x n grid of characters board and a string word, return true if word exists in the grid.
The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
The same letter cell may not be used more than once.

Constraints:

	m == board.length
	n = board[i].length
	1 <= m, n <= 6
	1 <= word.length <= 15
	board and word consists of only lowercase and uppercase English letters.

Example 1:

	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
	Output: true

Example 2:

	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
	Output: true

Example 3:

	Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
	Output: false

Follow up: Could you use search pruning to make your solution faster with a larger board?

解题思路:
	在地图上的任意一个起点开始，
	向 4 个方向分别 DFS 搜索，
	直到所有的单词字母都找到了就输出 true，否则输出 false
 */

var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}

// best solution
type Cell struct {
	Row, Col int
}

var direction = []Cell{{0,1},{0,-1},{1,0},{-1,0}}

func existBest(board [][]byte, word string) bool {
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

func main() {
	bytes := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
	fmt.Printf("exist(bytes,\"ABCCED\") = %v\n",exist(bytes,"ABCCED")) // true
	fmt.Printf("exist(bytes,\"SEE\") = %v\n",exist(bytes,"SEE")) // true
	fmt.Printf("exist(bytes,\"ABCB\") = %v\n",exist(bytes,"ABCB")) // false

	fmt.Printf("existBest(bytes,\"ABCCED\") = %v\n",existBest(bytes,"ABCCED")) // true
	fmt.Printf("existBest(bytes,\"SEE\") = %v\n",existBest(bytes,"SEE")) // true
	fmt.Printf("existBest(bytes,\"ABCB\") = %v\n",existBest(bytes,"ABCB")) // false
}
