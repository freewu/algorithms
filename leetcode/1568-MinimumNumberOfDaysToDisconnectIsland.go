package main

// 1568. Minimum Number of Days to Disconnect Island
// You are given an m x n binary grid grid where 1 represents land and 0 represents water. 
// An island is a maximal 4-directionally (horizontal or vertical) connected group of 1's.

// The grid is said to be connected if we have exactly one island, otherwise is said disconnected.
// In one day, we are allowed to change any single land cell (1) into a water cell (0).
// Return the minimum number of days to disconnect the grid.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/24/land1.jpg" />
// Input: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]
// Output: 2
// Explanation: We need at least 2 days to get a disconnected grid.
// Change land grid[1][1] and grid[0][2] to water and get 2 disconnected island.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/12/24/land2.jpg" />
// Input: grid = [[1,1]]
// Output: 2
// Explanation: Grid of full water is also disconnected ([[1,1]] -> [[0,0]]), 0 islands.

// Constraints:
//     m == grid.length
//     n == grid[i].length
//     1 <= m, n <= 30
//     grid[i][j] is either 0 or 1.

import "fmt"

func minDays(grid [][]int) int {
    dirs := [][]int{[]int{0,1}, []int{0,-1}, []int{-1,0}, []int{1, 0}}
    graph := map[[2]int][][2]int{}
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1{
                if graph[[2]int{i,j}] == nil {
                    graph[[2]int{i,j}] = [][2]int{}
                }
                for _, d := range dirs {
                    x, y := i + d[0], j + d[1]
                    if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] == 1 {
                        graph[[2]int{i,j}] = append(graph[[2]int{i,j}], [2]int{x,y})
                    }
                }
            }
        }
    }
    // 特殊情况处理
    if len(graph) == 1 {
        return 1
    }
    if len(graph) == 0 {
        return 0
    }
    dfn, low := map[[2]int]int{}, map[[2]int]int{}
    for u := range graph {
        dfn[u], low[u] = 0, 0
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    ts := 1
    cuts := map[[2]int]bool{} // 求割点
    var trajan func(u, parent [2]int)
    trajan = func(u, parent [2]int) {
        dfn[u], low[u] = ts, ts
        ts++
        children := 0
        for _, v := range graph[u] {
            if v == parent{
                continue
            }
            if dfn[v] != 0 { // 处理回边
                low[u] = min(low[u], dfn[v])
            } else {
                children++
                trajan(v, u)
                low[u] = min(low[u], low[v])
                // 判断 割点
                if parent == [2]int{-1,-1} && children >= 2 {
                    cuts[u] = true
                } else if parent != [2]int{-1,-1} && low[v] >= dfn[u] {
                    cuts[u] = true
                }
            }
        }
    }
    // trajan 算法
    cnt := 0
    for u := range graph {
        if dfn[u] != 0 {
            continue
        }
        cnt++
        if cnt > 1 { // 如果 连通分量是 大于1， 则 直接返回 0 不需要操作
            return 0
        }
        trajan(u, [2]int{-1, -1})
    }
    if len(cuts) > 0 { 
        return 1
    }
    return 2
}

func minDays1(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    cntOne := 0
    dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    g := make(map[int][]int, m * n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                cntOne++
                for _, d := range dirs {
                    nx, ny := i+d[0], j+d[1]
                    if nx < 0 || nx >= m || ny < 0 || ny >= n {
                        continue
                    }
                    if grid[nx][ny] == 1 {
                        g[i*n+j] = append(g[i*n+j], nx*n+ny)
                    }
                }
            }
        }
    }
    if cntOne == 1 {
        return 1
    }
    if len(g) == 0 {
        return 0
    }
    isCut := make([]bool, m*n)
    dfn := make([]int, m*n) // DFS 到结点 v 的时间（从 1 开始）
    low := make([]int, m*n) //v节点最小时间 的时间（从 1 开始）
    // low[v] 定义为以下两种情况的最小值
    // 1. dfn[v]
    // 2. subtree(v) 的返祖边所指向的节点的 dfn，也就是经过恰好一条不在 DFS 树上的边，能够到达 subtree(v) 的节点的 dfn
    dfsClock, cnt := 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var tarjan func(int, int) int
    tarjan = func(v, fa int) int { // 无需考虑重边
        dfsClock++
        dfn[v] = dfsClock
        low[v] = dfsClock
        childCnt := 0
        for _, w := range g[v] {
            if dfn[w] == 0 {
                childCnt++
                lowW := tarjan(w, v)
                low[v] = min(low[v], lowW)
                if lowW >= dfn[v] { // 以 w 为根的子树中没有反向边能连回 v 的祖先（可以连到 v 上，这也算割点）
                    isCut[v] = true
                    cnt++
                }
            } else if w!=fa { // （w!=fa 可以省略，但为了保证某些题目没有重复统计所以保留）   找到 v 的反向边 v-w，用 dfn[w] 来更新 lowV
                low[v] = min(low[v], low[w])
            }
        }
        if fa == -1 && childCnt == 1 { // 特判：在 DFS 树上只有一个儿子的树根，删除后并没有增加连通分量的个数，这种情况下不是割点
            if isCut[v] {
                cnt--
            }
            isCut[v] = false
        }
        return low[v]
    }
    cntR := 0
    for i := range grid {
        for j, v := range grid[i] {
            if v == 1 && dfn[i*n+j] == 0 {
                if cntR > 0 {
                    return 0
                }
                cntR++
                tarjan(i*n+j, -1)
            }
        }
    }
    if cnt > 0 {
        return 1
    }
    return 2
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/24/land1.jpg" />
    // Input: grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]
    // Output: 2
    // Explanation: We need at least 2 days to get a disconnected grid.
    // Change land grid[1][1] and grid[0][2] to water and get 2 disconnected island.
    fmt.Println(minDays([][]int{{0,1,1,0},{0,1,1,0},{0,0,0,0}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/12/24/land2.jpg" />
    // Input: grid = [[1,1]]
    // Output: 2
    // Explanation: Grid of full water is also disconnected ([[1,1]] -> [[0,0]]), 0 islands.
    fmt.Println(minDays([][]int{{1,1}})) // 2

    fmt.Println(minDays1([][]int{{0,1,1,0},{0,1,1,0},{0,0,0,0}})) // 2
    fmt.Println(minDays1([][]int{{1,1}})) // 2
}