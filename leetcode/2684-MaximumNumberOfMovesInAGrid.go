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

// dp
func maxMoves(grid [][]int) int {
    res, move := 0, make([][]int, len(grid))
    for i := range move {
        move[i] = make([]int, len(grid[0]))
    }
    max := func (x,y int) int { if x > y { return x; }; return y; }
    for c := len(grid[0])-1; c >= 0; c-- {
        for r := len(grid)-1; r >= 0; r-- {
            if c + 1 >= len(grid[r]) { continue }
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
    for _, v := range move {
        if v[0] > res {
            res = v[0]
        }
    }
    return res
}

// bfs
func maxMoves1(grid [][]int) int {
    directions := [][]int{{-1, 1}, {0, 1}, {1, 1}}
    res, m, n := 0, len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    type Step struct {
        x, y, c int
    }
    max := func (x,y int) int { if x > y { return x; }; return y; }
    bfs := func(step Step) {
        queue := []Step{ step }
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:] // pop
            origin := grid[cur.x][cur.y]
            res = max(res, cur.c)
            for _, d := range directions {
                x, y := cur.x+d[0], cur.y+d[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] <= origin || visited[x][y] { continue }
                visited[x][y] = true
                queue = append(queue, Step{x, y, cur.c + 1})
            }
        }
    }
    for i := 0; i < m; i++ {
        if !visited[i][0] {
            bfs(Step{i, 0, 0})
        }
    }
    return res
}

func maxMoves2(grid [][]int) int {
    type Point struct { row, col int }
    res, m, n := 0, len(grid), len(grid[0])
    queue, visited := make([]Point, 0), make([][]bool, m) // 创建一个 visited 数组
    for i := range visited {
        visited[i] = make([]bool, n) // 初始化每一行
    }
    // 从第一列的每个单元格开始
    for i := 0; i < m; i++ {
        queue = append(queue, Point{i, 0})
        visited[i][0] = true // 将起始点标记为已访问
    }
    max := func (x,y int) int { if x > y { return x; }; return y; }
    for len(queue) != 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            top := queue[0]
            queue = queue[1:] // pop
            row, col := top.row, top.col
            // 检查可以移动的三个方向
            directions := []Point{{row - 1, col + 1}, {row, col + 1}, {row + 1, col + 1}}
            for _, d := range directions {
                newRow, newCol := d.row, d.col
                // 确保不超出边界并且值严格大于当前值
                if newRow >= 0 && newRow < m && newCol < n && !visited[newRow][newCol] && grid[newRow][newCol] > grid[row][col] {
                    queue = append(queue, Point{ newRow, newCol })
                    visited[newRow][newCol] = true // 标记为已访问
                    res = max(res, newCol) // 更新最大步数
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/04/11/yetgriddrawio-10.png"/>
    // Input: grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
    // Output: 3
    // Explanation: We can start at the cell (0, 0) and make the following moves:
    // - (0, 0) -> (0, 1).
    // - (0, 1) -> (1, 2).
    // - (1, 2) -> (2, 3).
    // It can be shown that it is the maximum number of moves that can be made.
    grid1 := [][]int {
        []int{2,4,3,5},
        []int{5,4,9,3},
        []int{3,4,2,11},
        []int{10,9,13,15},
    }
    fmt.Println(maxMoves(grid1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/04/12/yetgrid4drawio.png"/>
    // Input: grid = [[3,2,4],[2,1,9],[1,1,7]]
    // Output: 0
    // Explanation: Starting from any cell in the first column we cannot perform any moves.
    grid2 := [][]int {
        []int{3,2,4},
        []int{2,1,9},
        []int{1,1,7},
    }
    fmt.Println(maxMoves(grid2)) // 0

    fmt.Println(maxMoves1(grid1)) // 3
    fmt.Println(maxMoves1(grid2)) // 0

    fmt.Println(maxMoves2(grid1)) // 3
    fmt.Println(maxMoves2(grid2)) // 0
}