package main

// 3112. Minimum Time to Visit Disappearing Nodes
// There is an undirected graph of n nodes. 
// You are given a 2D array edges, where edges[i] = [ui, vi, lengthi] describes an edge between node ui and node vi with a traversal time of lengthi units.

// Additionally, you are given an array disappear, where disappear[i] denotes the time when the node i disappears from the graph and you won't be able to visit it.
// Notice that the graph might be disconnected and might contain multiple edges.

// Return the array answer, with answer[i] denoting the minimum 
// units of time required to reach node i from node 0. If node i is unreachable from node 0 then answer[i] is -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/03/09/example1.png" />
// Input:  n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,1,5]
// Output:  [0,-1,4]
// Explanation:
// We are starting our journey from node 0, and our goal is to find the minimum time required to reach each node before it disappears.
// For node 0, we don't need any time as it is our starting point.
// For node 1, we need at least 2 units of time to traverse edges[0]. Unfortunately, it disappears at that moment, so we won't be able to visit it.
// For node 2, we need at least 4 units of time to traverse edges[2].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/03/09/example2.png" />
// Input:  n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,3,5]
// Output:  [0,2,3]
// Explanation:
// We are starting our journey from node 0, and our goal is to find the minimum time required to reach each node before it disappears.
// For node 0, we don't need any time as it is the starting point.
// For node 1, we need at least 2 units of time to traverse edges[0].
// For node 2, we need at least 3 units of time to traverse edges[0] and edges[1].

// Example 3:
// Input: n = 2, edges = [[0,1,1]], disappear = [1,1]
// Output: [0,-1]
// Explanation:
// Exactly when we reach node 1, it disappears.

// Constraints:
//     1 <= n <= 5 * 10^4
//     0 <= edges.length <= 10^5
//     edges[i] == [ui, vi, lengthi]
//     0 <= ui, vi <= n - 1
//     1 <= lengthi <= 10^5
//     disappear.length == n
//     1 <= disappear[i] <= 10^5

import "fmt"
import "container/heap"

// dijkstra
func minimumTime(n int, edges [][]int, disappear []int) []int {
    adj := make([][]struct{ v, length int }, n)
    for _, edge := range edges {
        u, v, length := edge[0], edge[1], edge[2]
        adj[u] = append(adj[u], struct{ v, length int }{v, length})
        adj[v] = append(adj[v], struct{ v, length int }{u, length})
    }
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, Item{0, 0})
    res := make([]int, n)
    for i := range res {
        res[i] = -1
    }
    res[0] = 0
    for pq.Len() > 0 {
        item := heap.Pop(pq).(Item)
        t, u := item.priority, item.value
        if t != res[u] {
            continue
        }
        for _, edge := range adj[u] {
            v, length := edge.v, edge.length
            if t + length < disappear[v] && (res[v] == -1 || t + length < res[v]) {
                heap.Push(pq, Item{t + length, v})
                res[v] = t + length
            }
        }
    }
    return res
}

type Item struct {
    priority, value int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func minimumTime1(n int, edges [][]int, disappear []int) []int {
    g := make([][]pair, n)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w})
    }
    dist := make([]int, n)
    for i := range dist {
        dist[i] = 1 << 30
    }
    dist[0] = 0
    pq := hp{{0, 0}}
    for len(pq) > 0 {
        du, u := pq[0].dis, pq[0].u
        heap.Pop(&pq)
        if du > dist[u] { continue }
        for _, nxt := range g[u] {
            v, w := nxt.dis, nxt.u
            if dist[v] > dist[u]+w && dist[u]+w < disappear[v] {
                dist[v] = dist[u] + w
                heap.Push(&pq, pair{dist[v], v})
            }
        }
    }
    res := make([]int, n)
    for i := 0; i < n; i++ {
        if dist[i] < disappear[i] {
            res[i] = dist[i]
        } else {
            res[i] = -1
        }
    }
    return res
}

type pair struct{ dis, u int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/03/09/example1.png" />
    // Input:  n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,1,5]
    // Output:  [0,-1,4]
    // Explanation:
    // We are starting our journey from node 0, and our goal is to find the minimum time required to reach each node before it disappears.
    // For node 0, we don't need any time as it is our starting point.
    // For node 1, we need at least 2 units of time to traverse edges[0]. Unfortunately, it disappears at that moment, so we won't be able to visit it.
    // For node 2, we need at least 4 units of time to traverse edges[2].
    fmt.Println(minimumTime(3,[][]int{{0,1,2},{1,2,1},{0,2,4}}, []int{1,1,5})) // [0,-1,4]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/03/09/example2.png" />
    // Input:  n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,3,5]
    // Output:  [0,2,3]
    // Explanation:
    // We are starting our journey from node 0, and our goal is to find the minimum time required to reach each node before it disappears.
    // For node 0, we don't need any time as it is the starting point.
    // For node 1, we need at least 2 units of time to traverse edges[0].
    // For node 2, we need at least 3 units of time to traverse edges[0] and edges[1].
    fmt.Println(minimumTime(3,[][]int{{0,1,2},{1,2,1},{0,2,4}}, []int{1,3,5})) // [0,2,3]
    // Example 3:
    // Input: n = 2, edges = [[0,1,1]], disappear = [1,1]
    // Output: [0,-1]
    // Explanation:
    // Exactly when we reach node 1, it disappears.
    fmt.Println(minimumTime(2,[][]int{{0,1,1}}, []int{1,1})) // [0,-1]

    fmt.Println(minimumTime1(3,[][]int{{0,1,2},{1,2,1},{0,2,4}}, []int{1,1,5})) // [0,-1,4]
    fmt.Println(minimumTime1(3,[][]int{{0,1,2},{1,2,1},{0,2,4}}, []int{1,3,5})) // [0,2,3]
    fmt.Println(minimumTime1(2,[][]int{{0,1,1}}, []int{1,1})) // [0,-1]
}