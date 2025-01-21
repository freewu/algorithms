package main

// 2858. Minimum Edge Reversals So Every Node Is Reachable
// There is a simple directed graph with n nodes labeled from 0 to n - 1. 
// The graph would form a tree if its edges were bi-directional.

// You are given an integer n and a 2D integer array edges, 
// where edges[i] = [ui, vi] represents a directed edge going from node ui to node vi.

// An edge reversal changes the direction of an edge, 
// i.e., a directed edge going from node ui to node vi becomes a directed edge going from node vi to node ui.

// For every node i in the range [0, n - 1], 
// your task is to independently calculate the minimum number of edge reversals required so it is possible to reach any other node starting from node i through a sequence of directed edges.

// Return an integer array answer, 
// where answer[i] is the minimum number of edge reversals required so it is possible to reach any other node starting from node i through a sequence of directed edges.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/26/image-20230826221104-3.png" />
// Input: n = 4, edges = [[2,0],[2,1],[1,3]]
// Output: [1,1,0,2]
// Explanation: The image above shows the graph formed by the edges.
// For node 0: after reversing the edge [2,0], it is possible to reach any other node starting from node 0.
// So, answer[0] = 1.
// For node 1: after reversing the edge [2,1], it is possible to reach any other node starting from node 1.
// So, answer[1] = 1.
// For node 2: it is already possible to reach any other node starting from node 2.
// So, answer[2] = 0.
// For node 3: after reversing the edges [1,3] and [2,1], it is possible to reach any other node starting from node 3.
// So, answer[3] = 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/26/image-20230826225541-2.png" />
// Input: n = 3, edges = [[1,2],[2,0]]
// Output: [2,0,1]
// Explanation: The image above shows the graph formed by the edges.
// For node 0: after reversing the edges [2,0] and [1,2], it is possible to reach any other node starting from node 0.
// So, answer[0] = 2.
// For node 1: it is already possible to reach any other node starting from node 1.
// So, answer[1] = 0.
// For node 2: after reversing the edge [1, 2], it is possible to reach any other node starting from node 2.
// So, answer[2] = 1.

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ui == edges[i][0] < n
//     0 <= vi == edges[i][1] < n
//     ui != vi
//     The input is generated such that if the edges were bi-directional, the graph would be a tree.

import "fmt"

func minEdgeReversals(n int, edges [][]int) []int {
    graph := make(map[int]map[int]int, n)
    for i := 0; i < n; i++ {
        graph[i] = make(map[int]int)
    }
    for _, v := range edges {
        graph[v[0]][v[1]], graph[v[1]][v[0]] = 0, 1
    }
    dp := make(map[int]map[int]int)
    for i := 0; i < n; i++ {
        dp[i] = make(map[int]int)
    }
    var dfs func(i, parent int) int
    dfs = func(i, parent int) int {
        sum := 0
        if v, ok := dp[i][parent]; ok {  return v }
        for k, v := range graph[i] {
            if k == parent { continue }
            sum += dfs(k, i) + v
        }
        dp[i][parent] = sum
        return sum
    }
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = dfs(i, -1)
    }
    return res
}

func minEdgeReversals1(n int, edges [][]int) []int {
    type Pair struct{ to, dir int }
    graph := make([][]Pair, n)
    for _, v := range edges {
        x, y := v[0], v[1]
        graph[x] = append(graph[x], Pair{ y, 1 })
        graph[y] = append(graph[y], Pair{ x, -1 }) // 从 y 到 x 需要反向
    }
    res := make([]int, n)
    var dfs func(int, parent int)
    dfs = func(i, parent int) {
        for _, p := range graph[i] {
            if p.to != parent {
                if p.dir < 0 {
                    res[0]++
                }
                dfs(p.to, i)
            }
        }
    }
    dfs(0, -1)
    var reroot func(x, parent int)
    reroot = func(x, parent int) {
        for _, p := range graph[x] {
            if p.to != parent {
                res[p.to] = res[x] + p.dir
                reroot(p.to, x)
            }
        }
    }
    reroot(0,-1)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/08/26/image-20230826221104-3.png" />
    // Input: n = 4, edges = [[2,0],[2,1],[1,3]]
    // Output: [1,1,0,2]
    // Explanation: The image above shows the graph formed by the edges.
    // For node 0: after reversing the edge [2,0], it is possible to reach any other node starting from node 0.
    // So, answer[0] = 1.
    // For node 1: after reversing the edge [2,1], it is possible to reach any other node starting from node 1.
    // So, answer[1] = 1.
    // For node 2: it is already possible to reach any other node starting from node 2.
    // So, answer[2] = 0.
    // For node 3: after reversing the edges [1,3] and [2,1], it is possible to reach any other node starting from node 3.
    // So, answer[3] = 2.
    fmt.Println(minEdgeReversals(4, [][]int{{2,0},{2,1},{1,3}})) // [1,1,0,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/08/26/image-20230826225541-2.png" />
    // Input: n = 3, edges = [[1,2],[2,0]]
    // Output: [2,0,1]
    // Explanation: The image above shows the graph formed by the edges.
    // For node 0: after reversing the edges [2,0] and [1,2], it is possible to reach any other node starting from node 0.
    // So, answer[0] = 2.
    // For node 1: it is already possible to reach any other node starting from node 1.
    // So, answer[1] = 0.
    // For node 2: after reversing the edge [1, 2], it is possible to reach any other node starting from node 2.
    // So, answer[2] = 1.
    fmt.Println(minEdgeReversals(3, [][]int{{1,2},{2,0}})) // [2,0,1]

    fmt.Println(minEdgeReversals1(4, [][]int{{2,0},{2,1},{1,3}})) // [1,1,0,2]
    fmt.Println(minEdgeReversals1(3, [][]int{{1,2},{2,0}})) // [2,0,1]
}