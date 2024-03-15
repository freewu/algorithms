package main

// 2684. Maximum Number of Moves in a Grid
// You are given a 0-indexed m x n matrix grid consisting of positive integers.
// You can start at any cell in the first column of the matrix, and traverse the grid in the following way:
//     From a cell (row, col), you can move to any of the cells: 
//         (row - 1, col + 1), (row, col + 1) and (row + 1, col + 1) such that the value of the cell you move to, 
//         should be strictly bigger than the value of the current cell.

// Return the maximum number of moves that you can perform.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/04/11/yetgriddrawio-10.png"/>
// Input: grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
// Output: 3
// Explanation: We can start at the cell (0, 0) and make the following moves:
// - (0, 0) -> (0, 1).
// - (0, 1) -> (1, 2).
// - (1, 2) -> (2, 3).
// It can be shown that it is the maximum number of moves that can be made.
// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/04/12/yetgrid4drawio.png"/>
// Input: grid = [[3,2,4],[2,1,9],[1,1,7]]
// Output: 0
// Explanation: Starting from any cell in the first column we cannot perform any moves.
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 1000
//     4 <= m * n <= 10^5
//     1 <= grid[i][j] <= 10^6

import "fmt"

// db
func maxMoves(grid [][]int) int {
    max := func (a,b int) int { if a > b { return a; }; return b; }
	move := make([][]int, len(grid))
	for i := range move {
		move[i] = make([]int, len(grid[0]))
	}
	for c := len(grid[0])-1; c >= 0; c-- {
		for r := len(grid)-1; r >= 0; r-- {
			if c + 1 >= len(grid[r]) {
				continue
			}
			if grid[r][c] < grid[r][c+1] {
				move[r][c] = max(move[r][c], move[r][c+1]+1)
			}
			if r + 1 < len(grid) {
				if grid[r][c] < grid[r+1][c+1] {
					move[r][c] = max(move[r][c], move[r+1][c+1]+1)
				}
			}
			if r - 1 >= 0 {
				if grid[r][c] < grid[r-1][c+1] {
					move[r][c] = max(move[r][c], move[r-1][c+1]+1)
				}
			}
		}
	}
	res := 0
	for _, m := range move {
		if m[0] > res {
			res = m[0]
		}
	}
	return res
}

type step struct {
	x, y, c int
}

var DIR = [][]int{{-1, 1}, {0, 1}, {1, 1}}

// bfs
func maxMoves1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
    //fmt.Println(m, n)
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	res := 0
	bfs := func(st step) {
		q := []step{st}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			origin := grid[cur.x][cur.y]
			res = max(res, cur.c)
			for _, d := range DIR {
				x, y := cur.x+d[0], cur.y+d[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] <= origin || vis[x][y]{
					continue
				}
                vis[x][y] = true
				q = append(q, step{x, y, cur.c + 1})
			}
		}
	}
	for i := 0; i < m; i++ {
		if !vis[i][0] {
			bfs(step{i, 0, 0})
		}
	}
	return res
}

func main() {
    // Explanation: We can start at the cell (0, 0) and make the following moves:
    // - (0, 0) -> (0, 1).
    // - (0, 1) -> (1, 2).
    // - (1, 2) -> (2, 3).
    fmt.Println(
        maxMoves(
            [][]int {
                []int{2,4,3,5},
                []int{5,4,9,3},
                []int{3,4,2,11},
                []int{10,9,13,15},
            },
        ),
    ) // 3
    fmt.Println(
        maxMoves(
            [][]int {
                []int{3,2,4},
                []int{2,1,9},
                []int{1,1,7},
            },
        ),
    ) // 0

    fmt.Println(
        maxMoves1(
            [][]int {
                []int{2,4,3,5},
                []int{5,4,9,3},
                []int{3,4,2,11},
                []int{10,9,13,15},
            },
        ),
    ) // 3
    fmt.Println(
        maxMoves1(
            [][]int {
                []int{3,2,4},
                []int{2,1,9},
                []int{1,1,7},
            },
        ),
    ) // 0
}