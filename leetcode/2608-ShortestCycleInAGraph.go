package main

// 2608. Shortest Cycle in a Graph
// There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1. 
// The edges in the graph are represented by a given 2D integer array edges, where edges[i] = [ui, vi] denotes an edge between vertex ui and vertex vi. 
// Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.

// Return the length of the shortest cycle in the graph. 
// If no cycle exists, return -1.

// A cycle is a path that starts and ends at the same node, and each edge in the path is used only once.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/01/04/cropped.png" />
// Input: n = 7, edges = [[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]]
// Output: 3
// Explanation: The cycle with the smallest length is : 0 -> 1 -> 2 -> 0 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/01/04/croppedagin.png" />
// Input: n = 4, edges = [[0,1],[0,2]]
// Output: -1
// Explanation: There are no cycles in this graph.

// Constraints:
//     2 <= n <= 1000
//     1 <= edges.length <= 1000
//     edges[i].length == 2
//     0 <= ui, vi < n
//     ui != vi
//     There are no repeated edges.

import "fmt"

// // bfs
// func findShortestCycle(n int, edges [][]int) int {
//     next := make(map[int][]int)
//     for _, edge := range edges {
//         s, e := edge[0], edge[1]
//         if _, ok := next[s]; !ok {
//             next[s] = []int{}
//         }
//         next[s] = append(next[s], e)
//         if _, ok := next[e]; !ok {
//             next[e] = []int{}
//         }
//         next[e] = append(next[e], s)
//     }
//     res := 1 << 31
//     for i := 0; i < n; i++ {
//         curr := make(map[int]int)
//         queue := [][]int{ []int{i, -1, 0} }  // current position, previous position, counter
//         for len(queue) > 0 {
//             p, prev, counter := queue[0][0], queue[0][1], queue[0][2]
//             queue = queue[1:] // pop
//             if v, ok := curr[p]; ok {
//                 // start at position i, find a min cycle
//                 if counter + v < res  { 
//                     res = counter + v 
//                 }
//                 break
//             }
//             curr[p] = counter
//             for _, nxt := range next[p] {
//                 if nxt == prev { continue }
//                 queue = append(queue, []int{ nxt, p, counter + 1}) // push
//             }
//         }
//     }
//     if res == 1 << 31 { return -1 }
//     return res
// }

// Time Limit Exceeded  87 / 88 
// func findShortestCycle(n int, edges [][]int) int {
//     res := 1 << 31
//     graph := make(map[int][]int)
//     for _,edge := range(edges){
//         a, b := edge[0], edge[1]
//         graph[a] = append(graph[a],b)
//         graph[b] = append(graph[b],a)
//     }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     for i := 0; i < n; i++ {
//         dist := map[int]int{i:0}
//         par := make(map[int]int)
//         q := []int{ i }
//         for len(q) > 0 {
//             node := q[0]
//             q = q[1:] // pop
//             for _, nei := range(graph[node]) {
//                 if _,ok := dist[nei]; !ok {
//                     dist[nei] = 1 + dist[node]
//                     par[nei] = node
//                     q = append(q,nei)
//                 } else if  par[node] != nei && par[nei] != node {
//                     res = min(res, dist[node] + dist[nei] + 1)
//                 }
//             }
//         }
//     }
//     if res == 1 << 31{ return -1 }
//     return res
// }

func findShortestCycle(n int, edges [][]int) int {
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x) // 建图
    }
    type Pair struct{ x, fa int }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    res := 1 << 31
    dis := make([]int, n) // dis[i] 表示从 start 到 i 的最短路长度
    for start := 0; start < n; start++ { // 枚举每个起点跑 BFS
        for j := range dis {
            dis[j] = -1
        }
        dis[start] = 0
        q := []Pair{{start, -1}}
        for len(q) > 0 {
            p := q[0]
            q = q[1:]
            x, fa := p.x, p.fa
            for _, y := range g[x] {
                if dis[y] < 0 { // 第一次遇到
                    dis[y] = dis[x] + 1
                    q = append(q, Pair{y, x})
                } else if y != fa { // 第二次遇到
                    res = min(res, dis[x] + dis[y] + 1)
                }
            }
        }
    }
    if res == 1 << 31 { // 无环图
        return -1
    }
    return res
}


func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/01/04/cropped.png" />
    // Input: n = 7, edges = [[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]]
    // Output: 3
    // Explanation: The cycle with the smallest length is : 0 -> 1 -> 2 -> 0 
    fmt.Println(findShortestCycle(7,[][]int{{0,1},{1,2},{2,0},{3,4},{4,5},{5,6},{6,3}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/01/04/croppedagin.png" />
    // Input: n = 4, edges = [[0,1],[0,2]]
    // Output: -1
    // Explanation: There are no cycles in this graph.
    fmt.Println(findShortestCycle(7,[][]int{{0,1},{0,2}})) // -1

    fmt.Println(findShortestCycle(6,[][]int{{4,2},{5,1},{5,0},{0,3},{5,2},{1,4},{1,3},{3,4}})) // 3
    fmt.Println(findShortestCycle(12,[][]int{{0,3},{0,5},{3,4},{4,5},{1,9},{1,11},{9,10},{11,10},{2,6},{2,8},{6,7},{8,7},{0,1},{0,2},{1,2}})) // 3
    
}