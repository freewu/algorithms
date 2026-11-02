package main

// 4096. Simulate Bomb Explosions and Falling Boxes
// You are given a 2D character array grid of size m x n, where each cell contains one of the following symbols:
//     1. '#': a box.
//     2. '-': an empty cell.
//     3. '*': a bomb.

// Each second, the following events occur in order:
//     1. Explosions: A bomb explodes if it is orthogonally adjacent to at least one box (shares a side with it).
//        An exploding bomb destroys all boxes in its 8 neighboring cells, including diagonally adjacent cells, turning them into empty cells.
//        All explosions happen simultaneously, based on the grid at the beginning of that second.
//        Empty cells and bombs remain unchanged.
//     2. Gravity: After the explosions, all non-static boxes simultaneously fall down by exactly one row.
//        A box is static if it is in the bottommost row, immediately above a bomb, or immediately above another static box.

// Bombs never move or disappear, and the same bomb may explode multiple times.
// Boxes above non-static boxes can fall together with them.

// Return the final grid after 10^{100} seconds.

// Example 1:
// Input: grid = [["#","#","#"],["-","-","*"],["-","-","-"]]
// Output: [["-","-","-"],["-","-","*"],["#","-","-"]]
// Explanation:
// Let (r, c) represent the cell in the r^{th} row from the top and c^{th} column from the left.
// For example, (0, 0) denotes the top-left most cell.
// The 1st second:
// - Since the bomb at (1, 2) is orthogonally adjacent to the box at (0, 2), the boxes at (0, 1) and (0, 2) become empty cells.
// - The box at (0, 0) falls to (1, 0).
// The 2nd second:
// - The bomb at (1, 2) is not orthogonally adjacent to any box. Therefore, nothing happens.
// - The box at (1, 0) falls to (2, 0).
// Nothing happens afterwards.

// Example 2:
// Input: grid = [["-","*"],["#","-"]]
// Output: [["-","*"],["#","-"]]
// Explanation:
//  Nothing happens because the bomb is not orthogonally adjacent to any box.

// Example 3:
// Input: grid = [["-","*","#"],["#","-","-"]]
// Output: [["-","*","-"],["-","-","-"]]
// Explanation:
// The 1st second:
// - The bomb at (0, 1) is orthogonally adjacent to the box at (0, 2).
//   Therefore, the boxes at (0, 2) and (1, 0) become empty cells.
// Nothing happens afterwards.

// Example 4:
// Input: grid = [["#","*","-","#"],["-","-","*","-"]]
// Output: [["-","*","-","-"],["-","-","*","-"]]
// Explanation:
// The 1st second:
// - Since the bomb at (0, 1) is orthogonally adjacent to the box at (0, 0), the box at (0, 0) becomes an empty cell.
//   The other bomb at (1, 2) is not orthogonally adjacent to any box.
// - The box at (0, 3) falls to (1, 3).
// The 2nd second:
// - Since the bomb at (1, 2) is orthogonally adjacent to the box at (1, 3), the box at (1, 3) becomes an empty cell.
//   The other bomb at (0, 1) is not orthogonally adjacent to any box.

// Nothing happens afterwards.
//     - 1 <= m == grid.length <= 100
//     - 1 <= n == grid[i].length <= 100
//     - grid[i][j] is '#', '-', or '*'.

import "fmt"

func finalGrid(grid [][]byte) [][]byte {
	m := len(grid)
	if m == 0 {
		return grid
	}
	n := len(grid[0])
	canBomb := func(x, y int) bool {
		if grid[x][y] != '*' {
			return false
		}
		dirs := [][]int{
			{x - 1, y},
			{x, y - 1},
			{x, y + 1},
			{x + 1, y},
		}
		for _, d := range dirs {
			i, j := d[0], d[1]
			if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '#' {
				return true
			}
		}
		return false
	}
	bomb := func(x, y int) bool {
		if grid[x][y] != '!' {
			return false
		}
		startI, endI, startJ, endJ := max(0, x-1), min(m, x+2), max(0, y-1), min(n, y+2)
		for i := startI; i < endI; i++ {
			for j := startJ; j < endJ; j++ {
				if grid[i][j] == '#' {
					grid[i][j] = '-'
				}
			}
		}
		grid[x][y] = '*'
		return true
	}
	explosion := func() bool {
		// 先标记所有需要炸的位置为 !，不能边遍历边改
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if canBomb(i, j) {
					grid[i][j] = '!'
				}
			}
		}
		changed := false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if bomb(i, j) {
					changed = true
				}
			}
		}
		return changed
	}
	gravity := func() bool {
		changed := false
		// 从倒数第二行往上遍历
		for i := m - 2; i >= 0; i-- {
			for j := 0; j < n; j++ {
				if grid[i][j] == '#' && grid[i+1][j] == '-' {
					grid[i][j], grid[i+1][j] = grid[i+1][j], grid[i][j]
					changed = true
				}
			}
		}
		return changed
	}
	for explosion() || gravity() {
	}
	return grid
}

func finalGrid1(grid [][]byte) [][]byte {
	n, m := len(grid), len(grid[0])
	dir4 := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	dir8 := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	for {
		type Pair struct{ x, y int }
		todo := []Pair{}
		for i, r := range grid {
			for j, v := range r {
				if v != '*' {
					continue
				}
				for _, dir := range dir4 {
					x, y := i+dir.x, j+dir.y
					if 0 <= x && x < n && 0 <= y && y < m {
						if grid[x][y] == '#' {
							todo = append(todo, Pair{i, j})
							break
						}
					}
				}
			}
		}
		upd := len(todo) > 0
		for _, p := range todo {
			for _, dir := range dir8 {
				x, y := p.x+dir.x, p.y+dir.y
				if 0 <= x && x < n && 0 <= y && y < m {
					if grid[x][y] == '#' {
						grid[x][y] = '-'
						upd = true
					}
				}
			}
		}
		for i := n - 2; i >= 0; i-- {
			for j, v := range grid[i] {
				if v == '#' && grid[i+1][j] == '-' {
					grid[i+1][j] = '#'
					grid[i][j] = '-'
					upd = true
				}
			}
		}
		if !upd {
			return grid
		}
	}
}

func main() {
	// Example 1:
	// Input: grid = [["#","#","#"],["-","-","*"],["-","-","-"]]
	// Output: [["-","-","-"],["-","-","*"],["#","-","-"]]
	// Explanation:
	// Let (r, c) represent the cell in the r^{th} row from the top and c^{th} column from the left.
	// For example, (0, 0) denotes the top-left most cell.
	// The 1st second:
	// - Since the bomb at (1, 2) is orthogonally adjacent to the box at (0, 2), the boxes at (0, 1) and (0, 2) become empty cells.
	// - The box at (0, 0) falls to (1, 0).
	// The 2nd second:
	// - The bomb at (1, 2) is not orthogonally adjacent to any box. Therefore, nothing happens.
	// - The box at (1, 0) falls to (2, 0).
	// Nothing happens afterwards.
	fmt.Println(finalGrid([][]byte{{byte('#'), byte('#'), byte('#')}, {byte('-'), byte('-'), byte('*')}, {byte('-'), byte('-'), byte('-')}})) // [["-","-","-"],["-","-","*"],["#","-","-"]]
	// Example 2:
	// Input: grid = [["-","*"],["#","-"]]
	// Output: [["-","*"],["#","-"]]
	// Explanation:
	//  Nothing happens because the bomb is not orthogonally adjacent to any box.
	fmt.Println(finalGrid([][]byte{{byte('-'), byte('*')}, {byte('#'), byte('-')}})) // [["-","*","#","-"]]
	// Example 3:
	// Input: grid = [["-","*","#"],["#","-","-"]]
	// Output: [["-","*","-"],["-","-","-"]]
	// Explanation:
	// The 1st second:
	// - The bomb at (0, 1) is orthogonally adjacent to the box at (0, 2).
	//   Therefore, the boxes at (0, 2) and (1, 0) become empty cells.
	// Nothing happens afterwards.
	fmt.Println(finalGrid([][]byte{{byte('-'), byte('*')}, {byte('#'), byte('#')}})) // [["-","*","-","-"],["-","-","*"],["#","-","-"]]
	// Example 4
	// Input: grid = [["#","*","-","#"],["-","-","*","-"]]
	// Output: [["-","*","-","-"],["-","-","*","-"]]
	// Explanation:
	// The 1st second:
	// - Since the bomb at (0, 1) is orthogonally adjacent to the box at (0, 0), the box at (0, 0) becomes an empty cell.
	//   The other bomb at (1, 2) is not orthogonally adjacent to any box.
	// - The box at (0, 3) falls to (1, 3).
	// The 2nd second:
	// - Since the bomb at (1, 2) is orthogonally adjacent to the box at (1, 3), the box at (1, 3) becomes an empty cell.
	//   The other bomb at (0, 1) is not orthogonally adjacent to any box.
	fmt.Println(finalGrid([][]byte{{byte('#'), byte('#'), byte('#')}, {byte('-'), byte('-'), byte('*')}, {byte('-'), byte('-'), byte('-')}})) // [["-","-","-"],["-","-","*"],["#","-","-"]]

	fmt.Println(finalGrid1([][]byte{{byte('-'), byte('*')}, {byte('#'), byte('-')}}))                                                          // [["-","*","#","-"]]
	fmt.Println(finalGrid1([][]byte{{byte('-'), byte('*')}, {byte('#'), byte('#')}}))                                                          // [["-","*","-","-"],["-","-","*"],["#","-","-"]]
	fmt.Println(finalGrid1([][]byte{{byte('#'), byte('#'), byte('#')}, {byte('-'), byte('-'), byte('*')}, {byte('-'), byte('-'), byte('-')}})) // [["-","-","-"],["-","-","*"],["#","-","-"]]
}
