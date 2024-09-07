package main

// 3123. Find Edges in Shortest Paths
// You are given an undirected weighted graph of n nodes numbered from 0 to n - 1. 
// The graph consists of m edges represented by a 2D array edges, where edges[i] = [ai, bi, wi] indicates that there is an edge between nodes ai and bi with weight wi.

// Consider all the shortest paths from node 0 to node n - 1 in the graph. 
// You need to find a boolean array answer where answer[i] is true if the edge edges[i] is part of at least one shortest path. 
// Otherwise, answer[i] is false.

// Return the array answer.

// Note that the graph may not be connected.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/03/05/graph35drawio-1.png" />
// Input: n = 6, edges = [[0,1,4],[0,2,1],[1,3,2],[1,4,3],[1,5,1],[2,3,1],[3,5,3],[4,5,2]]
// Output: [true,true,true,false,true,true,true,false]
// Explanation:
// The following are all the shortest paths between nodes 0 and 5:
// The path 0 -> 1 -> 5: The sum of weights is 4 + 1 = 5.
// The path 0 -> 2 -> 3 -> 5: The sum of weights is 1 + 1 + 3 = 5.
// The path 0 -> 2 -> 3 -> 1 -> 5: The sum of weights is 1 + 1 + 2 + 1 = 5.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/03/05/graphhhh.png" />
// Input: n = 4, edges = [[2,0,1],[0,1,1],[0,3,4],[3,2,2]]
// Output: [true,false,false,true]
// Explanation:
// There is one shortest path between nodes 0 and 3, which is the path 0 -> 2 -> 3 with the sum of weights 1 + 2 = 3.

// Constraints:
//     2 <= n <= 5 * 10^4
//     m == edges.length
//     1 <= m <= min(5 * 10^4, n * (n - 1) / 2)
//     0 <= ai, bi < n
//     ai != bi
//     1 <= wi <= 10^5
//     There are no repeated edges.

import "fmt"
import "container/heap"

func findAnswer(n int, edges [][]int) []bool {
    type Edge struct { to, w, i int }
    m, inf := len(edges), 1 << 31
    graph, lowDist, res := make([][]Edge, n) ,make([]int, n) , make([]bool, m) // 图, 最短路径
    for i, e := range edges { //初始化
        x, y, w := e[0], e[1], e[2]
        graph[x] = append(graph[x], Edge{y, w, i})
        graph[y] = append(graph[y], Edge{x, w, i})
    }
    for i := 1; i < n; i++ {
        lowDist[i] = inf
    }
    // dijkstra
    h := &MinHeap{{}}
    for len(*h) > 0 {
        p := heap.Pop(h).(Pair)
        distX, x := p.dis, p.x
        if distX > lowDist[x] {
            continue
        }
        for _, e := range graph[x] {
            y, w := e.to, e.w
            newDist := distX + w
            if newDist < lowDist[y] {
                lowDist[y] = newDist
                heap.Push(h, Pair{newDist, y})
            }
        }
    }
    // 特判
    if lowDist[n-1] == inf {
        return res
    }
    // 从终点反向dfs
    visited := make([]bool, n) // 是否已访问
    var dfs func(y int)       // 反向dfs找最短路路径: lowDist[x] + w = lowDist[y]
    dfs = func(y int) {
        if visited[y] { return } // 已访问
        visited[y] = true
        for _, e := range graph[y] {
            x := e.to
            if lowDist[x] + e.w != lowDist[y] {
                continue
            }
            res[e.i] = true
            dfs(x)
        }
    }
    dfs(n - 1)
    return res
}

// 最小堆
type Pair struct { dis, x int }
type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any) { *h = append(*h, v.(Pair))}
func (h *MinHeap) Pop() (v any) {
    a := *h
    *h, v = a[:len(a)-1], a[len(a)-1]
    return
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/03/05/graph35drawio-1.png" />
    // Input: n = 6, edges = [[0,1,4],[0,2,1],[1,3,2],[1,4,3],[1,5,1],[2,3,1],[3,5,3],[4,5,2]]
    // Output: [true,true,true,false,true,true,true,false]
    // Explanation:
    // The following are all the shortest paths between nodes 0 and 5:
    // The path 0 -> 1 -> 5: The sum of weights is 4 + 1 = 5.
    // The path 0 -> 2 -> 3 -> 5: The sum of weights is 1 + 1 + 3 = 5.
    // The path 0 -> 2 -> 3 -> 1 -> 5: The sum of weights is 1 + 1 + 2 + 1 = 5.
    fmt.Println(findAnswer(6,[][]int{{0,1,4},{0,2,1},{1,3,2},{1,4,3},{1,5,1},{2,3,1},{3,5,3},{4,5,2}})) // [true,true,true,false,true,true,true,false]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/03/05/graphhhh.png" />
    // Input: n = 4, edges = [[2,0,1],[0,1,1],[0,3,4],[3,2,2]]
    // Output: [true,false,false,true]
    // Explanation:
    // There is one shortest path between nodes 0 and 3, which is the path 0 -> 2 -> 3 with the sum of weights 1 + 2 = 3.
    fmt.Println(findAnswer(4,[][]int{{2,0,1},{0,1,1},{0,3,4},{3,2,2}})) // [true,false,false,true]
}