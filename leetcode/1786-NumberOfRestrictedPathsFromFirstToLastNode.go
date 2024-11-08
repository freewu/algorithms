package main

// 1786. Number of Restricted Paths From First to Last Node
// There is an undirected weighted connected graph. 
// You are given a positive integer n which denotes that the graph has n nodes labeled from 1 to n, 
// and an array edges where each edges[i] = [ui, vi, weighti] denotes that there is an edge between nodes ui 
// and vi with weight equal to weighti.

// A path from node start to node end is a sequence of nodes [z0, z1, z2, ..., zk] 
// such that z0 = start and zk = end and there is an edge between zi and zi+1 where 0 <= i <= k-1.

// The distance of a path is the sum of the weights on the edges of the path. 
// Let distanceToLastNode(x) denote the shortest distance of a path between node n and node x. 
// A restricted path is a path that also satisfies that distanceToLastNode(zi) > distanceToLastNode(zi+1) where 0 <= i <= k-1.

// Return the number of restricted paths from node 1 to node n. 
// Since that number may be too large, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/17/restricted_paths_ex1.png" />
// Input: n = 5, edges = [[1,2,3],[1,3,3],[2,3,1],[1,4,2],[5,2,2],[3,5,1],[5,4,10]]
// Output: 3
// Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The three restricted paths are:
// 1) 1 --> 2 --> 5
// 2) 1 --> 2 --> 3 --> 5
// 3) 1 --> 3 --> 5

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/17/restricted_paths_ex22.png" />
// Input: n = 7, edges = [[1,3,1],[4,1,2],[7,3,4],[2,5,3],[5,6,1],[6,7,2],[7,5,3],[2,6,4]]
// Output: 1
// Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The only restricted path is 1 --> 3 --> 7.

// Constraints:
//     1 <= n <= 2 * 10^4
//     n - 1 <= edges.length <= 4 * 10^4
//     edges[i].length == 3
//     1 <= ui, vi <= n
//     ui != vi
//     1 <= weighti <= 10^5
//     There is at most one edge between any two nodes.
//     There is at least one path between any two nodes.

import "fmt"
import "container/heap"
import "sort"

type Edge struct {
    v, w int
}

type DistHeap struct {
    h    []int // nodes 
    pos  []int // position in heap
    dist []int // distance
}
   
func (d *DistHeap) Len() int { return len(d.h) }
func (d *DistHeap) Less(i, j int) bool {  return d.dist[d.h[i]] < d.dist[d.h[j]] }
func (d *DistHeap) Swap(i, j int) { 
    u, v := d.h[i], d.h[j]
    d.h[i], d.h[j] = v, u
    d.pos[u], d.pos[v] = j, i
}
func (d *DistHeap) Push(x interface{}) {
    // no push, only pop
}
func (d *DistHeap) Pop() interface{} {
    n := len(d.h)
    x := d.h[n-1]
    d.h = d.h[:n-1] 
    return x
}

func countRestrictedPaths(n int, edges [][]int) int {
    adj, inf, mod := make([][]Edge, n), 1 << 31, 1_000_000_007
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        u-- // re-label
        v--
        adj[u] = append(adj[u], Edge{v: v, w: w})
        adj[v] = append(adj[v], Edge{v: u, w: w})
    }
    dijkstra := func(adj [][]Edge, n int) int {
        dh := &DistHeap{ h: make([]int, n), pos: make([]int, n), dist: make([]int, n), }
        for i := range dh.dist {
            dh.dist[i] = inf
            dh.h[i] = i
            dh.pos[i] = i
        }
        dh.Swap(0, n-1)
        dh.dist[n-1] = 0
        // now start from n-1
        for dh.Len() > 0 {
            u := heap.Pop(dh).(int)
            for _, e := range adj[u] {
                if dh.dist[u] + e.w < dh.dist[e.v] {
                    dh.dist[e.v] = dh.dist[u] + e.w
                    heap.Fix(dh, dh.pos[e.v])
                }
            }
        }
        dist := dh.dist
        // reuse
        total := dh.pos
        for i := range total {
            total[i] = 0 // reset
        }
        total[n-1] = 1
        // reuse
        ind := dh.h[:n]
        for i := range ind {
            ind[i] = i
        }
        // sort by last distance
        sort.Slice(ind, func(i, j int) bool {
            return dist[ind[i]] < dist[ind[j]]
        }) 
        for _, u := range ind {
            if u == n - 1 { continue }
            if dist[u] == inf { break }
            du := dist[u]
            for _, e := range adj[u] {
                if dist[e.v] < du {
                    total[u] = ( total[u] + total[e.v]) % mod
                }
            }
        }
        return total[0]
    }
    return dijkstra(adj, n)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/02/17/restricted_paths_ex1.png" />
    // Input: n = 5, edges = [[1,2,3],[1,3,3],[2,3,1],[1,4,2],[5,2,2],[3,5,1],[5,4,10]]
    // Output: 3
    // Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The three restricted paths are:
    // 1) 1 --> 2 --> 5
    // 2) 1 --> 2 --> 3 --> 5
    // 3) 1 --> 3 --> 5
    fmt.Println(countRestrictedPaths(5, [][]int{{1,2,3},{1,3,3},{2,3,1},{1,4,2},{5,2,2},{3,5,1},{5,4,10}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/02/17/restricted_paths_ex22.png" />
    // Input: n = 7, edges = [[1,3,1],[4,1,2],[7,3,4],[2,5,3],[5,6,1],[6,7,2],[7,5,3],[2,6,4]]
    // Output: 1
    // Explanation: Each circle contains the node number in black and its distanceToLastNode value in blue. The only restricted path is 1 --> 3 --> 7.
    fmt.Println(countRestrictedPaths(7, [][]int{{1,3,1},{4,1,2},{7,3,4},{2,5,3},{5,6,1},{6,7,2},{7,5,3},{2,6,4}})) // 1
}