package main

// 694. Number of Distinct Islands
// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) 
// You may assume all four edges of the grid are surrounded by water.
// An island is considered to be the same as another if and only if one island can be translated (and not rotated or reflected) to equal the other.
// Return the number of distinct islands.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/01/distinctisland1-1-grid.jpg" />
// Input: grid = [[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]
// Output: 1

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/01/distinctisland1-2-grid.jpg" />
// Input: grid = [[1,1,0,1,1],[1,0,0,0,0],[0,0,0,0,1],[1,1,0,1,1]]
// Output: 3
 
// Constraints:
//         m == grid.length
//         n == grid[i].length
//         1 <= m, n <= 50
//         grid[i][j] is either 0 or 1.

import "fmt"
import "sort"

type Point struct {
	x, y int
}

func equalIslands(island1, island2 []Point) bool {
	if len(island1) != len(island2) {
		return false
	}
	for i := 0; i < len(island1); i++ {
		if island1[i].x != island2[i].x || island1[i].y != island2[i].y {
			return false
		}
	}
	return true
}

// 暴力破解
func numDistinctIslands(grid [][]int) int {
	seen := make(map[Point]bool)
	var uniqueIslands [][]Point
	var currentIsland []Point

	currentIslandUnique := func() bool {
		for _, otherIsland := range uniqueIslands {
			if equalIslands(otherIsland, currentIsland) {
				return false
			}
		}
		return true
	}

    // 执行 DFS 以查找当前岛中的所有单元。
    var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
			return
		}
		if seen[Point{row, col}] || grid[row][col] == 0 {
			return
		}
		seen[Point{row, col}] = true
		currentIsland = append(currentIsland, Point{row, col})
		dfs(row + 1, col)
		dfs(row - 1, col)
		dfs(row, col + 1)
		dfs(row, col - 1)
	}

    // 只要还有岛屿，就重复启动 DFS。
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			currentIsland = nil
			dfs(row, col)
			if len(currentIsland) == 0 {
				continue
			}
            // 把我们刚发现的岛平移到左上角。
			minCol := len(grid[0]) - 1
			for i := 0; i < len(currentIsland); i++ {
				minCol = min(minCol, currentIsland[i].y)
			}
			for i := range currentIsland {
				currentIsland[i].x -= row
				currentIsland[i].y -= minCol
			}
            // 如果这个岛是不同的，就把它添加到列表中。
			if currentIslandUnique() {
				uniqueIslands = append(uniqueIslands, append([]Point(nil), currentIsland...))
			}
		}
	}
	return len(uniqueIslands)
}

// 根据本地坐标哈希
func numDistinctIslands1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	var currRowOrigin, currColOrigin int
	var currentIsland []int
	seen := make(map[int]bool)

	// 执行 DFS 以查找当前岛中的所有单元。
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= m || col >= n || seen[row * n + col] || grid[row][col] == 0 {
			return
		}
		seen[row * n + col] = true
		currentIsland = append(currentIsland, (row-currRowOrigin) * m * n + col - currColOrigin)
		dfs(row + 1, col)
		dfs(row - 1, col)
		dfs(row, col + 1)
		dfs(row, col - 1)
	}

	var uniqueIslands [][]int
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			currentIsland = []int{}
			currRowOrigin, currColOrigin = row, col
			dfs(row, col)
			if len(currentIsland) > 0 {
				sort.Ints(currentIsland)
				uniqueIslands = append(uniqueIslands, currentIsland)
			}
		}
	}

	// 使用排序后的岛屿进行去重
	sort.Slice(uniqueIslands, func(i, j int) bool {
		return fmt.Sprint(uniqueIslands[i]) < fmt.Sprint(uniqueIslands[j])
	})
	uniqueCount := 0
	for i := 0; i < len(uniqueIslands); i++ {
		if i == 0 || fmt.Sprint(uniqueIslands[i]) != fmt.Sprint(uniqueIslands[i-1]) {
			uniqueCount++
		}
	}

	return uniqueCount
}

// 根据路径签名进行哈希
func numDistinctIslands2(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var currentIsland string

	var dfs func(int, int, byte)
	dfs = func(row, col int, dir byte) {
		if row < 0 || col < 0 || row >= m || col >= n || visited[row][col] || grid[row][col] == 0 {
			return
		}

		visited[row][col] = true
		currentIsland += string(dir)
		dfs(row + 1, col, 'D')
		dfs(row - 1, col, 'U')
		dfs(row, col + 1, 'R')
		dfs(row, col - 1, 'L')
        // 回溯 会通过在字符串后面添加 区分需要 回溯的
		currentIsland += "0"
	}

	islands := make(map[string]bool)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			currentIsland = ""
			dfs(row, col, '0')
			if len(currentIsland) == 0 {
				continue
			}
			islands[currentIsland] = true;
		}
	}

	return len(islands)
}

// best solution
func numDistinctIslands3(grid [][]int) int {
	ROW := len(grid)
	COL := len(grid[0])
	hashMap := make(map[string]bool)
	dirs4 := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var path []byte
	var dfs func(cr, cc int)
	dfs = func(cr, cc int) {
		for p, d := range dirs4 {
			nr := cr + d[0]
			nc := cc + d[1]
			if nr >= 0 && nc >= 0 && nr < ROW && nc < COL && grid[nr][nc] == 1 {
				grid[nr][nc] = 0
				path = append(path, byte(p)+'0')
				dfs(nr, nc)
				path = append(path, byte(p^1)+'0')
			}
		}
	}
	for r := 0; r < ROW; r++ {
		for c := 0; c < COL; c++ {
			if grid[r][c] == 1 {
				path = []byte{}
				dfs(r, c)
				hashMap[string(path)] = true
			}
		}
	}
	return len(hashMap)
}

func main() {
    fmt.Println(numDistinctIslands(
        [][]int{
            []int{1,1,0,0,0},
            []int{1,1,0,0,0},
            []int{0,0,0,1,1},
            []int{0,0,0,1,1},
        },
    )) // 1
    fmt.Println(numDistinctIslands(
        [][]int{
            []int{1,1,0,1,1},
            []int{1,0,0,0,0},
            []int{0,0,0,0,1},
            []int{1,1,0,1,1},
        },
    )) // 3

    fmt.Println(numDistinctIslands1(
        [][]int{
            []int{1,1,0,0,0},
            []int{1,1,0,0,0},
            []int{0,0,0,1,1},
            []int{0,0,0,1,1},
        },
    )) // 1
    fmt.Println(numDistinctIslands1(
        [][]int{
            []int{1,1,0,1,1},
            []int{1,0,0,0,0},
            []int{0,0,0,0,1},
            []int{1,1,0,1,1},
        },
    )) // 3

    fmt.Println(numDistinctIslands2(
        [][]int{
            []int{1,1,0,0,0},
            []int{1,1,0,0,0},
            []int{0,0,0,1,1},
            []int{0,0,0,1,1},
        },
    )) // 1
    fmt.Println(numDistinctIslands2(
        [][]int{
            []int{1,1,0,1,1},
            []int{1,0,0,0,0},
            []int{0,0,0,0,1},
            []int{1,1,0,1,1},
        },
    )) // 3

    fmt.Println(numDistinctIslands3(
        [][]int{
            []int{1,1,0,0,0},
            []int{1,1,0,0,0},
            []int{0,0,0,1,1},
            []int{0,0,0,1,1},
        },
    )) // 1
    fmt.Println(numDistinctIslands3(
        [][]int{
            []int{1,1,0,1,1},
            []int{1,0,0,0,0},
            []int{0,0,0,0,1},
            []int{1,1,0,1,1},
        },
    )) // 3
}