package main 

// 317. Shortest Distance from All Buildings
// You are given an m x n grid grid of values 0, 1, or 2, where:
//     each 0 marks an empty land that you can pass by freely,
//     each 1 marks a building that you cannot pass through, and
//     each 2 marks an obstacle that you cannot pass through.

// You want to build a house on an empty land that reaches all buildings in the shortest total travel distance. 
// You can only move up, down, left, and right.

// Return the shortest travel distance for such a house. 
// If it is not possible to build such a house according to the above rules, return -1.

// The total travel distance is the sum of the distances between the houses of the friends and the meeting point.
// The distance is calculated using Manhattan Distance, where distance(p1, p2) = |p2.x - p1.x| + |p2.y - p1.y|.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/buildings-grid.jpg" />
// Input: grid = [[1,0,2,0,1],[0,0,0,0,0],[0,0,1,0,0]]
// Output: 7
// Explanation: Given three buildings at (0,0), (0,4), (2,2), and an obstacle at (0,2).
// The point (1,2) is an ideal empty land to build a house, as the total travel distance of 3+3+1=7 is minimal.
// So return 7.

// Example 2:
// Input: grid = [[1,0]]
// Output: 1

// Example 3:
// Input: grid = [[1]]
// Output: -1
 
// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 50
//     grid[i][j] is either 0, 1, or 2.
//     There will be at least one building in the grid.

import "fmt"

func shortestDistance(grid [][]int) int {
    res, m, n := 1 << 32 - 1, len(grid), len(grid[0])
    if m == 0 { return 0 }
    sumDist := make([][]int,m) // 记录当前空地 0 的距离和（每段距离：到某一个building 1)
    sumCnt := make([][]int,m) // 记录当前空地 0 中被1(building)扫过的次数
    for i := range sumDist {
        sumDist[i] = make([]int,n)
        sumCnt[i] = make([]int,n)
    }
    cnt1 :=0 // number of the building (1)
    q := [][]int{} // {r, c, dist}
    d := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 四个方向
    inArea := func(r, c int) bool { // 边界检测
        return r >= 0 && r < m && c >= 0 && c < n
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < m; i++ {
        for j:=0; j < n; j++ {
            if grid[i][j] == 1 {
                cnt1++
                visited := make([][]bool, m)
                for i := range visited {
                    visited[i] = make([]bool, n)
                }
                q = append(q,[]int{i,j,0})
                for len(q)  >0 {
                    size := len(q)
                    for k := 0; k < size; k++ {
                    item := q[0]
                    q = q[1:]
                    for _, v := range d {
                        nr, nc := v[0] + item[0], v[1] + item[1]
                        if inArea(nr,nc) && grid[nr][nc] == 0 && !visited[nr][nc] {
                            sumCnt[nr][nc]++
                            visited[nr][nc] = true
                            sumDist[nr][nc]+=item[2]+1
                            q = append(q,[]int{nr,nc,item[2]+1})
                        }
                    }
                    }
                }
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 && sumCnt[i][j] == cnt1 {
                res = min(res, sumDist[i][j])
            }
        }
    }
    if res == 1 << 32 - 1 {
        return -1
    }
    return res
}

func shortestDistance1(grid [][]int) int {
    if 0 == len(grid) || 0 == len(grid[0]) {
        return -1
    }
    copyGrid := func (grid [][]int) [][]int {
        m, n := len(grid), len(grid[0])
        res := make([][]int, m)
        for i := 0; i < m; i++ {
            res[i] = make([]int, n)
            copy(res[i], grid[i])
        }
        return res
    }
    res, v, inf := 0, 0, 1 << 32 - 1
    m, n := len(grid), len(grid[0])
    sum := copyGrid(grid)
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 1 { // 不能能通过
                continue
            }
            res = inf
            dist := copyGrid(grid)
            qi, qj := []int{i}, []int{j}
            for 0 != len(qi) {
                ti, tj := qi[0], qj[0]
                qi, qj = qi[1:], qj[1:]
                for _, dir := range dirs {
                    y, x := ti+dir[0], tj+dir[1]
                    if y >= 0 && y < m && x >= 0 && x < n && grid[y][x] == v {
                        grid[y][x]--
                        dist[y][x] = dist[ti][tj] + 1
                        sum[y][x] += dist[y][x] - 1
                        qi, qj = append(qi, y), append(qj, x)
                        if sum[y][x] < res {
                            res = sum[y][x]
                        }
                    }
                }
            }
            v--
        }
    }
    if res == inf {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/buildings-grid.jpg" />
    // Input: grid = [[1,0,2,0,1],[0,0,0,0,0],[0,0,1,0,0]]
    // Output: 7
    // Explanation: Given three buildings at (0,0), (0,4), (2,2), and an obstacle at (0,2).
    // The point (1,2) is an ideal empty land to build a house, as the total travel distance of 3+3+1=7 is minimal.
    // So return 7.
    fmt.Println(shortestDistance([][]int{{1,0,2,0,1},{0,0,0,0,0},{0,0,1,0,0}})) // 7
    // Example 2:
    // Input: grid = [[1,0]]
    // Output: 1
    fmt.Println(shortestDistance([][]int{{1,0}})) // 1
    // Example 3:
    // Input: grid = [[1]]
    // Output: -1
    fmt.Println(shortestDistance([][]int{{1}})) // -1

    fmt.Println(shortestDistance1([][]int{{1,0,2,0,1},{0,0,0,0,0},{0,0,1,0,0}})) // 7
    fmt.Println(shortestDistance1([][]int{{1,0}})) // 1
    fmt.Println(shortestDistance1([][]int{{1}})) // -1
}
