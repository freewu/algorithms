package main

// 3650. Minimum Cost Path with Edge Reversals
// You are given a directed, weighted graph with n nodes labeled from 0 to n - 1, 
// and an array edges where edges[i] = [ui, vi, wi] represents a directed edge from node ui to node vi with cost wi.

// Each node ui has a switch that can be used at most once: 
//     when you arrive at ui and have not yet used its switch, 
//     you may activate it on one of its incoming edges vi → ui reverse that edge to ui → vi and immediately traverse it.

// The reversal is only valid for that single move, and using a reversed edge costs 2 * wi.

// Return the minimum total cost to travel from node 0 to node n - 1. If it is not possible, return -1.

// Example 1:
// Input: n = 4, edges = [[0,1,3],[3,1,1],[2,3,4],[0,2,2]]
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/07/e1drawio.png" />
// Use the path 0 → 1 (cost 3).
// At node 1 reverse the original edge 3 → 1 into 1 → 3 and traverse it at cost 2 * 1 = 2.
// Total cost is 3 + 2 = 5.

// Example 2:
// Input: n = 4, edges = [[0,2,1],[2,1,1],[1,3,1],[2,3,3]]
// Output: 3
// Explanation:
// No reversal is needed. Take the path 0 → 2 (cost 1), then 2 → 1 (cost 1), then 1 → 3 (cost 1).
// Total cost is 1 + 1 + 1 = 3.

// Constraints:
//     2 <= n <= 5 * 10^4
//     1 <= edges.length <= 10^5
//     edges[i] = [ui, vi, wi]
//     0 <= ui, vi <= n - 1
//     1 <= wi <= 1000

import "fmt"
import "container/heap"

type Pair struct{ dis, x int }
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minCost(n int, edges [][]int) int {
    type Edge struct{ to, weight int }
    g := make([][]Edge, n) // 邻接表
    for _, e := range edges {
        x, y, weight := e[0], e[1], e[2]
        g[x] = append(g[x], Edge{y, weight})
        g[y] = append(g[y], Edge{x, weight * 2}) // 反转边
    }
    dis := make([]int, n)
    for i := range dis {
        dis[i] = 1 << 31
    }
    dis[0] = 0 // 起点到自己的距离是 0
    // 堆中保存 (起点到节点 x 的最短路长度，节点 x)
    hp := &MinHeap{{}}
    for hp.Len() > 0 {
        p := heap.Pop(hp).(Pair)
        disX, x := p.dis, p.x
        if disX > dis[x] { continue } // x 之前出堆过
        if x == n - 1 { return disX } // 到达终点
        for _, e := range g[x] {
            y := e.to
            newDisY := disX + e.weight
            if newDisY < dis[y] {
                dis[y] = newDisY // 更新 x 的邻居的最短路
                // 懒更新堆：只插入数据，不更新堆中数据
                // 相同节点可能有多个不同的 newDisY，除了最小的 newDisY，其余值都会触发上面的 continue
                heap.Push(hp, Pair{newDisY, y})
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 4, edges = [[0,1,3],[3,1,1],[2,3,4],[0,2,2]]
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/07/e1drawio.png" />
    // Use the path 0 → 1 (cost 3).
    // At node 1 reverse the original edge 3 → 1 into 1 → 3 and traverse it at cost 2 * 1 = 2.
    // Total cost is 3 + 2 = 5.
    fmt.Println(minCost(4,[][]int{{0,1,3},{3,1,1},{2,3,4},{0,2,2}})) // 5
    // Example 2:
    // Input: n = 4, edges = [[0,2,1],[2,1,1],[1,3,1],[2,3,3]]
    // Output: 3
    // Explanation:
    // No reversal is needed. Take the path 0 → 2 (cost 1), then 2 → 1 (cost 1), then 1 → 3 (cost 1).
    // Total cost is 1 + 1 + 1 = 3.
    fmt.Println(minCost(4,[][]int{{0,2,1},{2,1,1},{1,3,1},{2,3,3}})) // 3
}