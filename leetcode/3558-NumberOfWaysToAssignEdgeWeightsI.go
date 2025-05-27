package main

// 3558. Number of Ways to Assign Edge Weights I
// There is an undirected tree with n nodes labeled from 1 to n, rooted at node 1. 
// The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi.

// Initially, all edges have a weight of 0. You must assign each edge a weight of either 1 or 2.

// The cost of a path between any two nodes u and v is the total weight of all edges in the path connecting them.

// Select any one node x at the maximum depth. Return the number of ways to assign edge weights in the path from node 1 to x such that its total cost is odd.

// Since the answer may be large, return it modulo 10^9 + 7.

// Note: Ignore all edges not in the path from node 1 to x.

// Example 1:
// <img src="https://pic.leetcode.cn/1748074049-lsGWuV-screenshot-2025-03-24-at-060006.png" />
// Input: edges = [[1,2]]
// Output: 1
// Explanation:
// The path from Node 1 to Node 2 consists of one edge (1 → 2).
// Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.

// Example 2:
// <img src="https://pic.leetcode.cn/1748074095-sRyffx-screenshot-2025-03-24-at-055820.png" />
// Input: edges = [[1,2],[1,3],[3,4],[3,5]]
// Output: 2
// Explanation:
// The maximum depth is 2, with nodes 4 and 5 at the same depth. Either node can be selected for processing.
// For example, the path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4).
// Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] == [ui, vi]
//     1 <= ui, vi <= n
//     edges represents a valid tree.

import "fmt"
import "slices"

func assignEdgeWeights(edges [][]int) int {
    n, mod := len(edges) + 1, 1_000_000_007
    g := make([][]int, n + 1)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x, fa int) int
    dfs = func(x, fa int) (d int) {
        for _, y := range g[x] {
            if y != fa { // 不递归到父节点
                d = max(d, dfs(y, x) + 1)
            }
        }
        return
    }
    pow := func(x, n int) int {
        res := 1
        for ; n > 0; n /= 2 {
            if n % 2 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    k := dfs(1, 0)
    return pow(2, k - 1)
}

func assignEdgeWeights1(edges [][]int) int {
    n := len(edges) + 1
    if n == 1 {
        return 0
    }
    g := NewGraph(n, 2 * n)
    for _, edge := range edges {
        u, v := edge[0]-1, edge[1]-1
        g.AddEdge(u, v)
        g.AddEdge(v, u)
    }
    dep := make([]int, n)
    var dfs func(p int, u int)
    dfs = func(p int, u int) {
        for i := g.nodes[u]; i > 0; i = g.next[i] {
            v := g.to[i]
            if p != v {
                dep[v] = dep[u] + 1
                dfs(u, v)
            }
        }
    }
    dfs(0, 0)
    x := slices.Max(dep)
    res := 1
    for i := 0; i < x - 1; i++ {
        res = res * 2 % 1_000_000_007
    }
    return res
}

type Graph struct {
    nodes []int
    next  []int
    to    []int
    cur   int
}

func NewGraph(n int, e int) *Graph {
    nodes := make([]int, n)
    next := make([]int, e + 3)
    to := make([]int, e + 3)
    return &Graph{ nodes, next, to, 0 }
}

func (g *Graph) AddEdge(u, v int) {
    g.cur++
    g.next[g.cur] = g.nodes[u]
    g.nodes[u] = g.cur
    g.to[g.cur] = v
}

func main() {
    // Example 1:
    // <img src="https://pic.leetcode.cn/1748074049-lsGWuV-screenshot-2025-03-24-at-060006.png" />
    // Input: edges = [[1,2]]
    // Output: 1
    // Explanation:
    // The path from Node 1 to Node 2 consists of one edge (1 → 2).
    // Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
    fmt.Println(assignEdgeWeights([][]int{{1,2}})) // 1
    // Example 2:
    // <img src="https://pic.leetcode.cn/1748074095-sRyffx-screenshot-2025-03-24-at-055820.png" />
    // Input: edges = [[1,2],[1,3],[3,4],[3,5]]
    // Output: 2
    // Explanation:
    // The maximum depth is 2, with nodes 4 and 5 at the same depth. Either node can be selected for processing.
    // For example, the path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4).
    // Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.
    fmt.Println(assignEdgeWeights([][]int{{1,2},{1,3},{3,4},{3,5}})) // 2

    fmt.Println(assignEdgeWeights1([][]int{{1,2}})) // 1
    fmt.Println(assignEdgeWeights1([][]int{{1,2},{1,3},{3,4},{3,5}})) // 2
}