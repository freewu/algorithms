package main

// 1970. Last Day Where You Can Still Cross
// There is a 1-based binary matrix where 0 represents land and 1 represents water. 
// You are given integers row and col representing the number of rows and columns in the matrix, respectively.

// Initially on day 0, the entire matrix is land. However, each day a new cell becomes flooded with water. 
// You are given a 1-based 2D array cells, where cells[i] = [ri, ci] represents that on the ith day, 
// the cell on the rith row and cith column (1-based coordinates) will be covered with water (i.e., changed to 1).

// You want to find the last day that it is possible to walk from the top to the bottom by only walking on land cells. 
// You can start from any cell in the top row and end at any cell in the bottom row. 
// You can only travel in the four cardinal directions (left, right, up, and down).

// Return the last day where it is possible to walk from the top to the bottom by only walking on land cells.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/27/1.png" />
// Input: row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]
// Output: 2
// Explanation: The above image depicts how the matrix changes each day starting from day 0.
// The last day where it is possible to cross from top to bottom is on day 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/07/27/2.png" />
// Input: row = 2, col = 2, cells = [[1,1],[1,2],[2,1],[2,2]]
// Output: 1
// Explanation: The above image depicts how the matrix changes each day starting from day 0.
// The last day where it is possible to cross from top to bottom is on day 1.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/07/27/3.png" />
// Input: row = 3, col = 3, cells = [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]
// Output: 3
// Explanation: The above image depicts how the matrix changes each day starting from day 0.
// The last day where it is possible to cross from top to bottom is on day 3.

// Constraints:
//     2 <= row, col <= 2 * 10^4
//     4 <= row * col <= 2 * 10^4
//     cells.length == row * col
//     1 <= ri <= row
//     1 <= ci <= col
//     All the values of cells are unique.

import "fmt"

// bfs
func latestDayToCross(row int, col int, cells [][]int) int {
    res, left, right := 0, 0, row * col
    isPossible := func(m int, n int, t int, cells [][]int) bool {
        grid := make([][]int, m+1)
        for i := range grid {
            grid[i] = make([]int, n+1)
        }
        directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
        for i := 0; i < t; i++ {
            grid[cells[i][0]][cells[i][1]] = 1
        }
        queue := [][]int{}
        for i := 1; i <= n; i++ {
            if grid[1][i] == 0 {
                queue = append(queue, []int{1, i})
                grid[1][i] = 1
            }
        }
        for len(queue) > 0 {
            cell := queue[0]
            queue = queue[1:]
            r, c := cell[0], cell[1]
            for _, dir := range directions {
                nr, nc := r+dir[0], c+dir[1]
                if nr > 0 && nc > 0 && nr <= m && nc <= n && grid[nr][nc] == 0 {
                    grid[nr][nc] = 1
                    if nr == m {
                        return true
                    }
                    queue = append(queue, []int{nr, nc})
                }
            }
        }
        return false
    }
    for left < right-1 {
        mid := left + (right-left)/2
        if isPossible(row, col, mid, cells) {
            left = mid
            res = mid
        } else {
            right = mid
        }
    }
    return res
}


// dfs
func latestDayToCross1(row int, col int, cells [][]int) int {
    l, r := 0, len(cells)-1
    n, m := row, col
    directions := []int{0, 1, 0, -1, 0}
    helper := func(vis [][]int, v [][]int, i int, j int) bool {
        vis[i][j] = 1
        stack := [][]int{{i, j}}
        res := false
        for len(stack) > 0 {
            cell := stack[len(stack)-1] // pop
            stack = stack[:len(stack)-1]
            r, c := cell[0], cell[1]
            if r == n - 1 {
                res = true
                break
            }
            for k := 0; k < 4; k++ {
                nr, nc := r + directions[k], c + directions[k+1]
                if nr >= 0 && nc >= 0 && nr < n && nc < m && vis[nr][nc] == 0 && v[nr][nc] == 0 {
                    vis[nr][nc] = 1
                    stack = append(stack, []int{nr, nc})
                }
            }
        }
        return res
    }
    for l <= r {
        mid := (l + r) / 2
        v := make([][]int, row)
        for i := range v {
            v[i] = make([]int, col)
        }
        for i := 0; i <= mid; i++ {
            v[cells[i][0]-1][cells[i][1]-1] = 1
        }
        vis := make([][]int, row)
        for i := range vis {
            vis[i] = make([]int, col)
        }
        flag := false
        for i := 0; i < col; i++ {
            if vis[0][i] == 0 && v[0][i] == 0 {
                flag = flag || helper(vis, v, 0, i)
            }
        }
        if flag {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return l
}

// Union Find
func latestDayToCross2(row, col int, cells [][]int) int {
    grid := make([][]int, row)
    for i := range grid {
        grid[i] = make([]int, col)
        for j := range grid[i] {
            grid[i][j] = 1
        }
    }
    p := make([]int, row*col+3)
    for i := range p {
        p[i] = -1
    }
    directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
    var find func (x int) int
    find = func (x int) int {
        if p[x] < 0 {
            return x
        }
        p[x] = find(p[x])
        return p[x]
    }
    union := func (x, y int) {
        x = find(x)
        y = find(y)
        if x == y {
            return
        }
        p[y] = x
    }
    startG, endG := row*col + 1, row*col + 2
    for j := 0; j < col; j++ {
        union(j, startG)
        union((row-1)*col+j, endG)
    }
    for i := row*col - 1; i >= 0; i-- {
        r := cells[i][0] - 1
        c := cells[i][1] - 1
        grid[r][c] = 0
        key := r*col + c
        for k := 0; k < 4; k++ {
            nr := r + directions[k][0]
            nc := c + directions[k][1]
            if nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc] == 0 {
                nKey := nr*col + nc
                union(key, nKey)
            }
            if find(startG) == find(endG) {
                return i
            }
        }
    }
    return row * col - 1
}

func latestDayToCross3(m, n int, cells [][]int) int {
    var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // 左右上下
    // 0：水
    // 1：陆地（待访问）
    // 2：陆地（已访问）
    state := make([][]int8, m)
    for i := range state {
        state[i] = make([]int8, n)
    }
    // 能否从第一行到达 (r, c)
    canReachFromTop := func(r, c int) bool {
        if r == 0 { // 已经是第一行
            return true
        }
        for _, d := range dirs {
            x, y := r+d.x, c+d.y
            if 0 <= x && x < m && 0 <= y && y < n && state[x][y] == 2 {
                return true
            }
        }
        return false
    }
    // 从 (r, c) 出发，能否到达最后一行
    var dfs func(int, int) bool
    dfs = func(r, c int) bool {
        if r == m-1 {
            return true
        }
        state[r][c] = 2 // 已访问的陆地
        for _, d := range dirs {
            x, y := r+d.x, c+d.y
            if 0 <= x && x < m && 0 <= y && y < n && state[x][y] == 1 && dfs(x, y) {
                return true
            }
        }
        return false
    }
    for day := len(cells) - 1; ; day-- {
        cell := cells[day]
        r, c := cell[0]-1, cell[1]-1 // 改成从 0 开始的下标
        state[r][c] = 1 // 待访问的陆地
        if canReachFromTop(r, c) && dfs(r, c) {
            return day
        }
    }
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/27/1.png" />
    // Input: row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]
    // Output: 2
    // Explanation: The above image depicts how the matrix changes each day starting from day 0.
    // The last day where it is possible to cross from top to bottom is on day 2.
    fmt.Println(latestDayToCross(2,2,[][]int{{1,1},{2,1},{1,2},{2,2}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/07/27/2.png" />
    // Input: row = 2, col = 2, cells = [[1,1],[1,2],[2,1],[2,2]]
    // Output: 1
    // Explanation: The above image depicts how the matrix changes each day starting from day 0.
    // The last day where it is possible to cross from top to bottom is on day 1.
    fmt.Println(latestDayToCross(2,2,[][]int{{1,1},{1,2},{2,1},{2,2}})) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/07/27/3.png" />
    // Input: row = 3, col = 3, cells = [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]
    // Output: 3
    // Explanation: The above image depicts how the matrix changes each day starting from day 0.
    // The last day where it is possible to cross from top to bottom is on day 3.
    fmt.Println(latestDayToCross(3,3,[][]int{{1,2},{2,1},{3,3},{2,2},{1,1},{1,3},{2,3},{3,2},{3,1}})) // 3

    fmt.Println(latestDayToCross1(2,2,[][]int{{1,1},{2,1},{1,2},{2,2}})) // 2
    fmt.Println(latestDayToCross1(2,2,[][]int{{1,1},{1,2},{2,1},{2,2}})) // 1
    fmt.Println(latestDayToCross1(3,3,[][]int{{1,2},{2,1},{3,3},{2,2},{1,1},{1,3},{2,3},{3,2},{3,1}})) // 3

    fmt.Println(latestDayToCross2(2,2,[][]int{{1,1},{2,1},{1,2},{2,2}})) // 2
    fmt.Println(latestDayToCross2(2,2,[][]int{{1,1},{1,2},{2,1},{2,2}})) // 1
    fmt.Println(latestDayToCross2(3,3,[][]int{{1,2},{2,1},{3,3},{2,2},{1,1},{1,3},{2,3},{3,2},{3,1}})) // 3

    fmt.Println(latestDayToCross3(2,2,[][]int{{1,1},{2,1},{1,2},{2,2}})) // 2
    fmt.Println(latestDayToCross3(2,2,[][]int{{1,1},{1,2},{2,1},{2,2}})) // 1
    fmt.Println(latestDayToCross3(3,3,[][]int{{1,2},{2,1},{3,3},{2,2},{1,1},{1,3},{2,3},{3,2},{3,1}})) // 3
}