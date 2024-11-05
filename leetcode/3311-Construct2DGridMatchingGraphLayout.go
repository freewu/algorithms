package main

// 3311. Construct 2D Grid Matching Graph Layout
// You are given a 2D integer array edges representing an undirected graph having n nodes, 
// where edges[i] = [ui, vi] denotes an edge between nodes ui and vi.

// Construct a 2D grid that satisfies these conditions:
//     1. The grid contains all nodes from 0 to n - 1 in its cells, with each node appearing exactly once.
//     2. Two nodes should be in adjacent grid cells (horizontally or vertically) 
//        if and only if there is an edge between them in edges.

// It is guaranteed that edges can form a 2D grid that satisfies the conditions.

// Return a 2D integer array satisfying the conditions above. 
// If there are multiple solutions, return any of them.

// Example 1:
// Input: n = 4, edges = [[0,1],[0,2],[1,3],[2,3]]
// Output: [[3,1],[2,0]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/11/screenshot-from-2024-08-11-14-07-59.png" />

// Example 2:
// Input: n = 5, edges = [[0,1],[1,3],[2,3],[2,4]]
// Output: [[4,2,3,1,0]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/11/screenshot-from-2024-08-11-14-06-02.png" />

// Example 3:
// Input: n = 9, edges = [[0,1],[0,4],[0,5],[1,7],[2,3],[2,4],[2,5],[3,6],[4,6],[4,7],[6,8],[7,8]]
// Output: [[8,6,3],[7,4,2],[1,0,5]]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/11/screenshot-from-2024-08-11-14-06-38.png" />

// Constraints:
//     2 <= n <= 5 * 10^4
//     1 <= edges.length <= 10^5
//     edges[i] = [ui, vi]
//     0 <= ui < vi < n
//     All the edges are distinct.
//     The input is generated such that edges can form a 2D grid that satisfies the conditions.

import "fmt"

func constructGridLayout(n int, edges [][]int) [][]int {
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    start := 0 // 找到度数最小的点
    for i, v := range graph {
        if len(v) < len(graph[start]) {
            start = i
        }
    }
    row, visited, degreeStart := []int{start}, make([]bool, n), len(graph[start]) // 起点的度数
    visited[start] = true
    for { // 注意题目保证 n >= 2，可以至少循环一次
        next := -1
        for _, y := range graph[start] {
            if !visited[y] && (next < 0 || len(graph[y]) < len(graph[next])) {
                next = y
            }
        }
        start = next
        row = append(row, start)
        visited[start] = true
        if len(graph[start]) == degreeStart { break }
    }
    k := len(row)
    res := make([][]int, n/k)
    res[0] = row
    for i := 1; i < len(res); i++ {
        res[i] = make([]int, k)
        for j, x := range res[i-1] {
            for _, y := range graph[x] {
                // 上左右的邻居都访问过了，没访问过的邻居只会在 x 下面
                if !visited[y] {
                    visited[y] = true
                    res[i][j] = y
                    break
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, edges = [[0,1],[0,2],[1,3],[2,3]]
    // Output: [[3,1],[2,0]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/11/screenshot-from-2024-08-11-14-07-59.png" />
    fmt.Println(constructGridLayout(4, [][]int{{0,1},{0,2},{1,3},{2,3}})) // [[3,1],[2,0]]
    // Example 2:
    // Input: n = 5, edges = [[0,1],[1,3],[2,3],[2,4]]
    // Output: [[4,2,3,1,0]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/11/screenshot-from-2024-08-11-14-06-02.png" />
    fmt.Println(constructGridLayout(5, [][]int{{0,1},{1,3},{2,3},{2,4}})) // [[4,2,3,1,0]]
    // Example 3:
    // Input: n = 9, edges = [[0,1],[0,4],[0,5],[1,7],[2,3],[2,4],[2,5],[3,6],[4,6],[4,7],[6,8],[7,8]]
    // Output: [[8,6,3],[7,4,2],[1,0,5]]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/11/screenshot-from-2024-08-11-14-06-38.png" />
    fmt.Println(constructGridLayout(9, [][]int{{0,1},{0,4},{0,5},{1,7},{2,3},{2,4},{2,5},{3,6},{4,6},{4,7},{6,8},{7,8}})) // [[8,6,3],[7,4,2],[1,0,5]]
}