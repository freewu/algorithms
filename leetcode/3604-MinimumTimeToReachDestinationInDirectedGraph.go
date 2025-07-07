package main

// 3604. Minimum Time to Reach Destination in Directed Graph
// You are given an integer n and a directed graph with n nodes labeled from 0 to n - 1. 
// This is represented by a 2D array edges, where edges[i] = [ui, vi, starti, endi] indicates an edge from node ui to vi that can only be used at any integer time t such that starti <= t <= endi.

// You start at node 0 at time 0.

// In one unit of time, you can either:
//     1. Wait at your current node without moving, or
//     2. Travel along an outgoing edge from your current node if the current time t satisfies starti <= t <= endi.

// Return the minimum time required to reach node n - 1. 
// If it is impossible, return -1.

// Example 1:
// Input: n = 3, edges = [[0,1,0,1],[1,2,2,5]]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004535.png" />
// The optimal path is:
// At time t = 0, take the edge (0 → 1) which is available from 0 to 1. You arrive at node 1 at time t = 1, then wait until t = 2.
// At time t = 2, take the edge (1 → 2) which is available from 2 to 5. You arrive at node 2 at time 3.
// Hence, the minimum time to reach node 2 is 3.

// Example 2:
// Input: n = 4, edges = [[0,1,0,3],[1,3,7,8],[0,2,1,5],[2,3,4,7]]
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004757.png" />
// The optimal path is:
// Wait at node 0 until time t = 1, then take the edge (0 → 2) which is available from 1 to 5. You arrive at node 2 at t = 2.
// Wait at node 2 until time t = 4, then take the edge (2 → 3) which is available from 4 to 7. You arrive at node 3 at t = 5.
// Hence, the minimum time to reach node 3 is 5.

// Example 3:
// Input: n = 3, edges = [[1,0,1,3],[1,2,3,5]]
// Output: -1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004914.png" />
// Since there is no outgoing edge from node 0, it is impossible to reach node 2. Hence, the output is -1.

// Constraints:
//     1 <= n <= 10^5
//     0 <= edges.length <= 10^5
//     edges[i] == [ui, vi, starti, endi]
//     0 <= ui, vi <= n - 1
//     ui != vi
//     0 <= starti <= endi <= 10^9

import "fmt"
import "container/heap"

type Item struct {
    node int
    time int
}

type MinHeap []Item

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}

func minTime(n int, edges [][]int) int {
    type Edge struct { to, start, end int }
    graph := make([][]Edge, n)
    for _, e := range edges {
        u, v, s, t := e[0], e[1], e[2], e[3]
        graph[u] = append(graph[u], Edge{v, s, t})
    }
    dist := make([]int, n)
    for i := 0; i < n; i++ {
        dist[i] = 1 << 61
    }
    dist[0] = 0
    h := &MinHeap{}
    heap.Push(h, Item{0, 0})
    for h.Len() > 0 {
        cur := heap.Pop(h).(Item)
        node, t := cur.node, cur.time
        if node == n - 1  { return t }
        if t > dist[node] { continue }
        for _,  e := range graph[node] {
            nextTime := max(t, e.start) + 1
            if nextTime - 1 > e.end { continue }
            if nextTime < dist[e.to] {
                dist[e.to] = nextTime
                heap.Push(h, Item{e.to, nextTime})
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1,0,1],[1,2,2,5]]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004535.png" />
    // The optimal path is:
    // At time t = 0, take the edge (0 → 1) which is available from 0 to 1. You arrive at node 1 at time t = 1, then wait until t = 2.
    // At time t = 2, take the edge (1 → 2) which is available from 2 to 5. You arrive at node 2 at time 3.
    // Hence, the minimum time to reach node 2 is 3.
    fmt.Println(minTime(3, [][]int{{0,1,0,1},{1,2,2,5}})) // 3
    // Example 2:
    // Input: n = 4, edges = [[0,1,0,3],[1,3,7,8],[0,2,1,5],[2,3,4,7]]
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004757.png" />
    // The optimal path is:
    // Wait at node 0 until time t = 1, then take the edge (0 → 2) which is available from 1 to 5. You arrive at node 2 at t = 2.
    // Wait at node 2 until time t = 4, then take the edge (2 → 3) which is available from 4 to 7. You arrive at node 3 at t = 5.
    // Hence, the minimum time to reach node 3 is 5.
    fmt.Println(minTime(4, [][]int{{0,1,0,3},{1,3,7,8},{0,2,1,5},{2,3,4,7}})) // 5
    // Example 3:
    // Input: n = 3, edges = [[1,0,1,3],[1,2,3,5]]
    // Output: -1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/06/05/screenshot-2025-06-06-at-004914.png" />
    // Since there is no outgoing edge from node 0, it is impossible to reach node 2. Hence, the output is -1.
    fmt.Println(minTime(3, [][]int{{1,0,1,3},{1,2,3,5}})) // -1
}
