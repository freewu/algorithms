package main

// 2577. Minimum Time to Visit a Cell In a Grid
// You are given a m x n matrix grid consisting of non-negative integers 
// where grid[row][col] represents the minimum time required to be able to visit the cell (row, col), 
// which means you can visit the cell (row, col) only when the time you visit it is greater than or equal to grid[row][col].

// You are standing in the top-left cell of the matrix in the 0th second, 
// and you must move to any adjacent cell in the four directions: 
//     up, down, left, and right. Each move you make takes 1 second.

// Return the minimum time required in which you can visit the bottom-right cell of the matrix. 
// If you cannot visit the bottom-right cell, then return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/02/14/yetgriddrawio-8.png" />
// Input: grid = [[0,1,3,2],[5,1,2,5],[4,3,8,6]]
// Output: 7
// Explanation: One of the paths that we can take is the following:
// - at t = 0, we are on the cell (0,0).
// - at t = 1, we move to the cell (0,1). It is possible because grid[0][1] <= 1.
// - at t = 2, we move to the cell (1,1). It is possible because grid[1][1] <= 2.
// - at t = 3, we move to the cell (1,2). It is possible because grid[1][2] <= 3.
// - at t = 4, we move to the cell (1,1). It is possible because grid[1][1] <= 4.
// - at t = 5, we move to the cell (1,2). It is possible because grid[1][2] <= 5.
// - at t = 6, we move to the cell (1,3). It is possible because grid[1][3] <= 6.
// - at t = 7, we move to the cell (2,3). It is possible because grid[2][3] <= 7.
// The final time is 7. It can be shown that it is the minimum time possible.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/02/14/yetgriddrawio-9.png" />
// Input: grid = [[0,2,4],[3,2,1],[1,0,4]]
// Output: -1
// Explanation: There is no path from the top left to the bottom-right cell.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     2 <= m, n <= 1000
//     4 <= m * n <= 10^5
//     0 <= grid[i][j] <= 10^5
//     grid[0][0] == 0

import "fmt"
import "container/heap"
import "sort"

// Dijkstra 
func minimumTime(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    if grid[0][1] > 1 && grid[1][0] > 1 { // 无法「等待」
        return -1
    }
    dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = 1 << 31
        }
    }
    dis[0][0] = 0
    h := &hp{{}}
    for { // 可以等待，就一定可以到达终点
        p := heap.Pop(h).(Tuple)
        d, i, j := p.d, p.i, p.j
        if d > dis[i][j] {
            continue
        }
        if i == m-1 && j == n-1 { // 找到终点，此时 d 一定是最短路
            return d
        }
        for _, q := range dirs { // 枚举周围四个格子
            x, y := i+q.x, j+q.y
            if 0 <= x && x < m && 0 <= y && y < n {
                nd := max(d+1, grid[x][y])
                nd += (nd - x - y) % 2 // nd 必须和 x+y 同奇偶
                if nd < dis[x][y] {
                    dis[x][y] = nd // 更新最短路
                    heap.Push(h, Tuple{nd, x, y})
                }
            }
        }
    }
    return -1
}

type Tuple struct{ d, i, j int }
type hp []Tuple
func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(Tuple)) }
func (h *hp) Pop() interface{}     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }


// 二分法
func minimumTime1(grid [][]int) int {
    type Pair struct{ x, y int }
    dirs := []Pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    m, n := len(grid), len(grid[0])
    if grid[0][1] > 1 && grid[1][0] > 1 { // 无法「等待」
        return -1
    }
    visited := make([][]int, m)
    for i := range visited {
        visited[i] = make([]int, n)
    }
    endTime := sort.Search(1e5+m+n, func(endTime int) bool {
        if endTime < grid[m-1][n-1] || endTime < m+n-2 {
            return false
        }
        visited[m-1][n-1] = endTime
        q := []Pair{{m - 1, n - 1}}
        for t := endTime - 1; len(q) > 0; t-- {
            tmp := q
            q = nil
            for _, p := range tmp {
                for _, d := range dirs { // 枚举周围四个格子
                    x, y := p.x+d.x, p.y+d.y
                    if 0 <= x && x < m && 0 <= y && y < n && visited[x][y] != endTime && grid[x][y] <= t {
                        if x == 0 && y == 0 {
                            return true
                        }
                        visited[x][y] = endTime // 用二分的值来标记，避免重复创建 visited 数组
                        q = append(q, Pair{x, y})
                    }
                }
            }
        }
        return false
    })
    return endTime + (endTime+m+n) % 2
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/02/14/yetgriddrawio-8.png" />
    // Input: grid = [[0,1,3,2],[5,1,2,5],[4,3,8,6]]
    // Output: 7
    // Explanation: One of the paths that we can take is the following:
    // - at t = 0, we are on the cell (0,0).
    // - at t = 1, we move to the cell (0,1). It is possible because grid[0][1] <= 1.
    // - at t = 2, we move to the cell (1,1). It is possible because grid[1][1] <= 2.
    // - at t = 3, we move to the cell (1,2). It is possible because grid[1][2] <= 3.
    // - at t = 4, we move to the cell (1,1). It is possible because grid[1][1] <= 4.
    // - at t = 5, we move to the cell (1,2). It is possible because grid[1][2] <= 5.
    // - at t = 6, we move to the cell (1,3). It is possible because grid[1][3] <= 6.
    // - at t = 7, we move to the cell (2,3). It is possible because grid[2][3] <= 7.
    // The final time is 7. It can be shown that it is the minimum time possible.
    grid1 := [][]int{
        {0,1,3,2},
        {5,1,2,5},
        {4,3,8,6},
    }
    fmt.Println(minimumTime(grid1)) // 7
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/02/14/yetgriddrawio-9.png" />
    // Input: grid = [[0,2,4],[3,2,1],[1,0,4]]
    // Output: -1
    // Explanation: There is no path from the top left to the bottom-right cell.
    grid2 := [][]int{
        {0,2,4},
        {3,2,1},
        {1,0,4},
    }
    fmt.Println(minimumTime(grid2)) // -1

    fmt.Println(minimumTime1(grid1)) // 7
    fmt.Println(minimumTime1(grid2)) // -1
}