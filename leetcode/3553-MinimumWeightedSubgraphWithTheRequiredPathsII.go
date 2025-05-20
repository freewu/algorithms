package main

// 3553. Minimum Weighted Subgraph With the Required Paths II
// You are given an undirected weighted tree with n nodes, numbered from 0 to n - 1. 
// It is represented by a 2D integer array edges of length n - 1, 
// where edges[i] = [ui, vi, wi] indicates that there is an edge between nodes ui and vi with weight wi.​

// Additionally, you are given a 2D integer array queries, where queries[j] = [src1j, src2j, destj].

// Return an array answer of length equal to queries.length, 
// where answer[j] is the minimum total weight of a subtree such that it is possible to reach destj from both src1j and src2j using edges in this subtree.

// A subtree here is any connected subset of nodes and edges of the original tree forming a valid tree.

// Example 1:
// Input: edges = [[0,1,2],[1,2,3],[1,3,5],[1,4,4],[2,5,6]], queries = [[2,3,4],[0,2,5]]
// Output: [12,11]
// Explanation:
// The blue edges represent one of the subtrees that yield the optimal answer.
// <img src="https://assets.leetcode.com/uploads/2025/04/02/tree1-4.jpg" />
// answer[0]: The total weight of the selected subtree that ensures a path from src1 = 2 and src2 = 3 to dest = 4 is 3 + 5 + 4 = 12.
// answer[1]: The total weight of the selected subtree that ensures a path from src1 = 0 and src2 = 2 to dest = 5 is 2 + 3 + 6 = 11.

// Example 2:
// Input: edges = [[1,0,8],[0,2,7]], queries = [[0,1,2]]
// Output: [15]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/02/tree1-5.jpg" />
// answer[0]: The total weight of the selected subtree that ensures a path from src1 = 0 and src2 = 1 to dest = 2 is 8 + 7 = 15.
 
// Constraints:
//     3 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 3
//     0 <= ui, vi < n
//     1 <= wi <= 10^4
//     1 <= queries.length <= 10^5
//     queries[j].length == 3
//     0 <= src1j, src2j, destj < n
//     src1j, src2j, and destj are pairwise distinct.
//     The input is generated such that edges represents a valid tree.

import "fmt"
import "math/bits"

func minimumWeight(edges [][]int, queries [][]int) []int {
    n := len(edges) + 1
    type Edge struct{ to, weight int }
    g := make([][]Edge, n)
    for _, e := range edges {
        x, y, w := e[0], e[1], e[2]
        g[x] = append(g[x], Edge{y, w})
        g[y] = append(g[y], Edge{x, w})
    }
    const mx = 17
    pa, dep, dis  := make([][mx]int, n), make([]int, n), make([]int, n)
    var dfs func(x, p int) 
    dfs = func(x, p int) {
        pa[x][0] = p
        for _, e := range g[x] {
            y := e.to
            if y == p { continue }
            dep[y] = dep[x] + 1
            dis[y] = dis[x] + e.weight
            dfs(y, x)
        }
    }
    dfs(0, -1)
    for i := 0; i < mx - 1; i++ {
        for x := range pa {
            p := pa[x][i]
            if p != -1 {
                pa[x][i + 1] = pa[p][i]
            } else {
                pa[x][i + 1] = -1
            }
        }
    }
    uptoDep := func(x, d int) int {
        for k := uint(dep[x] - d); k > 0; k &= k - 1 {
            x = pa[x][bits.TrailingZeros(k)]
        }
        return x
    }
    getLCA := func(x, y int) int {
        if dep[x] > dep[y] {
            x, y = y, x
        }
        y = uptoDep(y, dep[x])
        if y == x {
            return x
        }
        for i := mx - 1; i >= 0; i-- {
            if pv, pw := pa[x][i], pa[y][i]; pv != pw {
                x, y = pv, pw
            }
        }
        return pa[x][0]
    }
    getDis := func(x, y int) int { return dis[x] + dis[y] - dis[getLCA(x, y)]*2 }
    res := make([]int, len(queries))
    for i, q := range queries {
        a, b, c := q[0], q[1], q[2]
        res[i] = (getDis(a, b) + getDis(b, c) + getDis(a, c)) / 2
    }
    return res
}

func main() {
    // Example 1:
    // Input: edges = [[0,1,2],[1,2,3],[1,3,5],[1,4,4],[2,5,6]], queries = [[2,3,4],[0,2,5]]
    // Output: [12,11]
    // Explanation:
    // The blue edges represent one of the subtrees that yield the optimal answer.
    // <img src="https://assets.leetcode.com/uploads/2025/04/02/tree1-4.jpg" />
    // answer[0]: The total weight of the selected subtree that ensures a path from src1 = 2 and src2 = 3 to dest = 4 is 3 + 5 + 4 = 12.
    // answer[1]: The total weight of the selected subtree that ensures a path from src1 = 0 and src2 = 2 to dest = 5 is 2 + 3 + 6 = 11.
    fmt.Println(minimumWeight([][]int{{0,1,2},{1,2,3},{1,3,5},{1,4,4},{2,5,6}},[][]int{{2,3,4},{0,2,5}})) // [12,11]
    // Example 2:
    // Input: edges = [[1,0,8],[0,2,7]], queries = [[0,1,2]]
    // Output: [15]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/02/tree1-5.jpg" />
    // answer[0]: The total weight of the selected subtree that ensures a path from src1 = 0 and src2 = 1 to dest = 2 is 8 + 7 = 15.
    fmt.Println(minimumWeight([][]int{{1,0,8},{0,2,7}}, [][]int{{0,1,2}})) // [15]
}