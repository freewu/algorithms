package main

import "fmt"

/**
200. Number of Islands

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:

	Input: grid = [
	  ["1","1","1","1","0"],
	  ["1","1","0","1","0"],
	  ["1","1","0","0","0"],
	  ["0","0","0","0","0"]
	]
	Output: 1

Example 2:

	Input: grid = [
	  ["1","1","0","0","0"],
	  ["1","1","0","0","0"],
	  ["0","0","1","0","0"],
	  ["0","0","0","1","1"]
	]
	Output: 3


Constraints:

	m == grid.length
	n == grid[i].length
	1 <= m, n <= 300
	grid[i][j] is '0' or '1'.
 */

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res, visited := 0, make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				searchIslands(grid, &visited, i, j)
				res++
			}
		}
	}
	return res
}

func searchIslands(grid [][]byte, visited *[][]bool, x, y int) {
	(*visited)[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(grid, nx, ny) && !(*visited)[nx][ny] && grid[nx][ny] == '1' {
			searchIslands(grid, visited, nx, ny)
		}
	}
}

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
func numIslandsBest(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i:=0; i<m; i++ {
		visited[i] = make([]bool, n)
	}

	result := 0

	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				searchIslandsBest(grid, visited, i, j)
				result++
			}
		}
	}
	return result
}

func searchIslandsBest(grid [][]byte, visited [][]bool, i, j int) {
	visited[i][j] = true
	for m:=0; m<4; m++ {
		ni := i + dir[m][0]
		nj := j + dir[m][1]
		if isInBoard(grid, ni, nj) && visited[ni][nj] == false && grid[ni][nj] == '1' {
			searchIslandsBest(grid, visited, ni, nj)
		}

	}
}

func main() {
	map1 := [][]byte{ { '1','1','1','1','0'} , {'1','1','0','1','0'}, {'1','1','0','0','0'},{'0','0','0','0','0'} }
	map2 := [][]byte{ { '1','1','0','0','0'} , {'1','1','0','0','0'}, {'0','0','1','0','0'},{'0','0','0','1','1'} }
	fmt.Printf("map1 = %v\n",map1)
	fmt.Printf("map2 = %v\n",map2)

	fmt.Printf("numIslands(map1) = %v\n",numIslands(map1)) // 1
	fmt.Printf("numIslands(map2) = %v\n",numIslands(map2)) // 3

	fmt.Printf("numIslandsBest(map1) = %v\n",numIslandsBest(map1)) // 1
	fmt.Printf("numIslandsBest(map2) = %v\n",numIslandsBest(map2)) // 3
}
