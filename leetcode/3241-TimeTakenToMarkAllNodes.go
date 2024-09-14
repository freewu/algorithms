package main

// 3241. Time Taken to Mark All Nodes
// There exists an undirected tree with n nodes numbered 0 to n - 1. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the tree.

// Initially, all nodes are unmarked. For each node i:
//     If i is odd, the node will get marked at time x if there is at least one node adjacent to it which was marked at time x - 1.
//     If i is even, the node will get marked at time x if there is at least one node adjacent to it which was marked at time x - 2.

// Return an array times where times[i] is the time when all nodes get marked in the tree, if you mark node i at time t = 0.

// Note that the answer for each times[i] is independent, i.e. when you mark node i all other nodes are unmarked.

// Example 1:
// Input: edges = [[0,1],[0,2]]
// Output: [2,4,3]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" />
// For i = 0:
// Node 1 is marked at t = 1, and Node 2 at t = 2.
// For i = 1:
// Node 0 is marked at t = 2, and Node 2 at t = 4.
// For i = 2:
// Node 0 is marked at t = 2, and Node 1 at t = 3.

// Example 2:
// Input: edges = [[0,1]]
// Output: [1,2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png" />
// For i = 0:
// Node 1 is marked at t = 1.
// For i = 1:
// Node 0 is marked at t = 2.

// Example 3:
// Input: edges = [[2,4],[0,1],[2,3],[0,2]]
// Output: [4,6,3,5,5]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" />
 
// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= edges[i][0], edges[i][1] <= n - 1
//     The input is generated such that edges represents a valid tree.

import "fmt"

func timeTaken(edges [][]int) []int {
    res, adj := make([]int, len(edges)+1), make([][]int, len(edges)+1)
    for _, e := range edges { // 接邻表
        u, v := e[0], e[1]
        adj[u], adj[v] = append(adj[u], v), append(adj[v], u)
    }
    max := func(x, y int) int { if x > y { return x; }; return y; }
    var dfs func(res []int, adj [][]int, r, par int) int 
    dfs = func(res []int, adj [][]int, r, par int) int {
        for _, v := range adj[r] {
            if v == par { continue }
            res[r] = max(res[r], dfs(res, adj, v, r))
        }
        return res[r] + 2 - ( r % 2)
    }
    var rdfs func(res []int, adj [][]int, r, par, prv int) 
    rdfs = func(res []int, adj [][]int, r, par, prv int) {
        res[r] = max(res[r], prv)
        h1, h2 := prv, 0
        for _, v := range adj[r] {
            if v == par {
                continue
            }
            switch nt := res[v]+2-(v%2); {
            case nt > h1:
                h1, h2 = nt, h1
            case nt > h2:
                h2 = nt
            }
        }
        for _, v := range adj[r] {
            if v == par { continue }
            if c := res[v] + 2 - (v % 2); c == h1 {
                rdfs(res, adj, v, r, h2+2-(r%2))
            } else {
                rdfs(res, adj, v, r, h1+2-(r%2))
            }
        }
    }
    dfs(res, adj, 0, 0)
    rdfs(res, adj, 0, 0, 0)
    return res
}

func timeTaken1(edges [][]int) []int {
    g := make([][]int, len(edges)+1)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    res, nodes := make([]int, len(g)), make([]struct{ fi, se, w int }, len(g))
    var dfs func(x, fa int) int
    dfs = func(x, fa int) int {
        p := &nodes[x]
        for _, y := range g[x] {
            if y == fa { continue }
            fi := dfs(y, x) + 2 - y%2
            if fi > p.fi {
                p.se = p.fi
                p.fi = fi
                p.w = y
            } else if fi > p.se {
                p.se = fi
            }
        }
        return p.fi
    }
    dfs(0, -1)
    max := func(x, y int) int { if x > y { return x; }; return y; }
    var reroot func(x, fa, fromUp int)
    reroot = func(x, fa, fromUp int) {
        p := nodes[x]
        res[x] = max(p.fi, fromUp)
        for _, y := range g[x] {
            if y == fa { continue }
            w := 2 - x%2
            if y == p.w {
                reroot(y, x, max(p.se, fromUp) + w)
            } else {
                reroot(y, x, max(p.fi, fromUp) + w)
            }
        }
    }
    reroot(0, -1, 0)
    return res
}

func main() {
    // Example 1:
    // Input: edges = [[0,1],[0,2]]
    // Output: [2,4,3]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" />
    // For i = 0:
    // Node 1 is marked at t = 1, and Node 2 at t = 2.
    // For i = 1:
    // Node 0 is marked at t = 2, and Node 2 at t = 4.
    // For i = 2:
    // Node 0 is marked at t = 2, and Node 1 at t = 3.
    fmt.Println(timeTaken([][]int{{2,4},{0,1},{2,3},{0,2}})) // [2,4,3]
    // Example 2:
    // Input: edges = [[0,1]]
    // Output: [1,2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png" />
    // For i = 0:
    // Node 1 is marked at t = 1.
    // For i = 1:
    // Node 0 is marked at t = 2.
    fmt.Println(timeTaken([][]int{{0, 1}})) // [1,2]
    // Example 3:
    // Input: edges = [[2,4],[0,1],[2,3],[0,2]]
    // Output: [4,6,3,5,5]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" />
    fmt.Println(timeTaken([][]int{{0,1},{0,2}})) // [4,6,3,5,5]

    fmt.Println(timeTaken1([][]int{{2,4},{0,1},{2,3},{0,2}})) // [2,4,3]
    fmt.Println(timeTaken1([][]int{{0, 1}})) // [1,2]
    fmt.Println(timeTaken1([][]int{{0,1},{0,2}})) // [4,6,3,5,5]
}