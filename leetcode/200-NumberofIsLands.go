package main

// 200. Number of Islands
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.

// Example 1:
// Input: grid = [
//     ["1","1","1","1","0"],
//     ["1","1","0","1","0"],
//     ["1","1","0","0","0"],
//     ["0","0","0","0","0"]
// ]
// Output: 1

// Example 2:
// Input: grid = [
//     ["1","1","0","0","0"],
//     ["1","1","0","0","0"],
//     ["0","0","1","0","0"],
//     ["0","0","0","1","1"]
// ]
// Output: 3

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 300
//     grid[i][j] is '0' or '1'.

import "fmt"

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

var dir = [][]int{ []int{-1, 0}, []int{0, 1}, []int{1, 0},  []int{0, -1} }

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

func numIslands1(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := 0; i < m; i++ {
        visited[i] = make([]bool, n)
    }

    res := 0
    var searchIslands func (grid [][]byte, visited [][]bool, i, j int)
    searchIslands = func (grid [][]byte, visited [][]bool, i, j int) {
        visited[i][j] = true
        isInBoard := func (board [][]byte, x, y int) bool { return x >= 0 && x < len(board) && y >= 0 && y < len(board[0]); }
        for m:=0; m<4; m++ {
            ni := i + dir[m][0]
            nj := j + dir[m][1]
            if isInBoard(grid, ni, nj) && visited[ni][nj] == false && grid[ni][nj] == '1' {
                searchIslands(grid, visited, ni, nj)
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' && visited[i][j] == false {
                searchIslands(grid, visited, i, j)
                res++
            }
        }
    }
    return res
}

func numIslands2(grid [][]byte) int {
    row, col, count := len(grid), len(grid[0]), 0
    var dfs func(int, int)
    dfs = func(i, j int) {
        if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
            return
        }
        if grid[i][j] == '1' {
            grid[i][j] = '.' // 修改访问过的标记
        } else {
            return
        }
        // 以下的判断可以放在下一层的return判断中
        dfs(i + 1, j)
        dfs(i, j + 1)
        dfs(i - 1, j)
        dfs(i, j - 1)
    }
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if grid[i][j] == '1' {
                dfs(i, j)
                count += 1
            }
        }
    }
    return count
}

func main() {
    map1 := [][]byte{ { '1','1','1','1','0'} , {'1','1','0','1','0'}, {'1','1','0','0','0'},{'0','0','0','0','0'} }
    map2 := [][]byte{ { '1','1','0','0','0'} , {'1','1','0','0','0'}, {'0','0','1','0','0'},{'0','0','0','1','1'} }
    fmt.Printf("map1 = %v\n",map1)
    fmt.Printf("map2 = %v\n",map2)

    fmt.Printf("numIslands(map1) = %v\n",numIslands(map1)) // 1
    fmt.Printf("numIslands(map2) = %v\n",numIslands(map2)) // 3

    fmt.Printf("numIslands1(map1) = %v\n",numIslands1(map1)) // 1
    fmt.Printf("numIslands1(map2) = %v\n",numIslands1(map2)) // 3

    fmt.Printf("numIslands2(map1) = %v\n",numIslands2(map1)) // 1
    fmt.Printf("numIslands2(map2) = %v\n",numIslands2(map2)) // 3
}
