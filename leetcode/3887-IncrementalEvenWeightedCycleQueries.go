package main

// 3887. Incremental Even-Weighted Cycle Queries
// You are given a positive integer n.

// There is an undirected graph with n nodes labeled from 0 to n - 1. 
// Initially, the graph has no edges.

// You are also given a 2D integer array edges, where edges[i] = [ui, vi, wi] represents an edge between nodes ui and vi with weight wi. 
// The weight wi is either 0 or 1.

// Process the edges in edges in the given order. 
// For each edge, add it to the graph only if, after adding it, the sum of the weights of the edges in every cycle in the resulting graph is even.

// Return an integer denoting the number of edges that are successfully added to the graph.

// Example 1:
// Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,1]]
// Output: 2
// Explanation:
// <img alt src="https://assets.leetcode.com/uploads/2026/03/21/hmadizgovu.png" />
// [0, 1, 1]: We add the edge between vertex 0 and vertex 1 with weight 1.
// [1, 2, 1]: We add the edge between vertex 1 and vertex 2 with weight 1.
// [0, 2, 1]: The edge between vertex 0 and vertex 2 (the dashed edge in the diagram) is not added because the cycle 0 - 1 - 2 - 0 has total edge weight 1 + 1 + 1 = 3, which is an odd number.

// Example 2:
// Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,0]]
// Output: 3
// Explanation:
// <img alt src="https://assets.leetcode.com/uploads/2026/03/21/rbdgrefwok.png" />
// [0, 1, 1]: We add the edge between vertex 0 and vertex 1 with weight 1.
// [1, 2, 1]: We add the edge between vertex 1 and vertex 2 with weight 1.
// [0, 2, 0]: We add the edge between vertex 0 and vertex 2 with weight 0.
// Note that the cycle 0 - 1 - 2 - 0 has total edge weight 1 + 1 + 0 = 2, which is an even number.
 
// Constraints:
//     3 <= n <= 5 * 10^4
//     1 <= edges.length <= 5 * 10^4
//     edges[i] = [ui, vi, wi]
//     0 <= ui < vi < n
//     All edges are distinct.
//     wi = 0 or wi = 1

import "fmt"

type UnionFind struct {
    fa  []int // fa[x] 是 x 的代表元
    dis []int // dis[x] = 从 x 到 fa[x] 的路径异或和
}

func newUnionFind(n int) UnionFind {
    fa := make([]int, n)
    dis := make([]int, n)
    for i := range fa {
        fa[i] = i
    }
    return UnionFind{fa, dis}
}

func (u UnionFind) find(x int) int {
    if u.fa[x] != x {
        root := u.find(u.fa[x])
        u.dis[x] ^= u.dis[u.fa[x]]
        u.fa[x] = root
    }
    return u.fa[x]
}

func (u UnionFind) merge(from, to, value int) bool {
    x, y := u.find(from), u.find(to)
    if x == y {
        return u.dis[from]^u.dis[to] == value
    }
    u.dis[x] = value ^ u.dis[to] ^ u.dis[from]
    u.fa[x] = y
    return true
}

func numberOfEdgesAdded(n int, edges [][]int) int {
    res := 0
    uf := newUnionFind(n)
    for _, e := range edges {
        if uf.merge(e[0], e[1], e[2]) {
            res++
        }
    }
    return res
}

type DSU struct {
    p []int // parent
    s []int // xor weight to parent
}

func (d *DSU) find(i int) (int, int) {
    if d.p[i] == i { return i, 0 }
    root, weight := d.find(d.p[i])
    d.p[i] = root
    d.s[i] = d.s[i] ^ weight
    return d.p[i], d.s[i]
}

func numberOfEdgesAdded1(n int, edges [][]int) int {
    res, p, s := 0, make([]int, n), make([]int, n)
    for i := range p {
        p[i] = i
    }
    d := &DSU{p: p, s: s}
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        rootU, weightU := d.find(u)
        rootV, weightV := d.find(v)
        if rootU != rootV {
            // Union components
            d.p[rootU] = rootV
            d.s[rootU] = weightU ^ weightV ^ w
            res++
        } else {
            // Check if existing path parity matches new edge weight
            if (weightU ^ weightV) == w {
                res++
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,1]]
    // Output: 2
    // Explanation:
    // <img alt src="https://assets.leetcode.com/uploads/2026/03/21/hmadizgovu.png" />
    // [0, 1, 1]: We add the edge between vertex 0 and vertex 1 with weight 1.
    // [1, 2, 1]: We add the edge between vertex 1 and vertex 2 with weight 1.
    // [0, 2, 1]: The edge between vertex 0 and vertex 2 (the dashed edge in the diagram) is not added because the cycle 0 - 1 - 2 - 0 has total edge weight 1 + 1 + 1 = 3, which is an odd number.
    fmt.Println(numberOfEdgesAdded(3, [][]int{{0,1,1},{1,2,1},{0,2,1}})) // 2
    // Example 2:
    // Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,0]]
    // Output: 3
    // Explanation:
    // <img alt src="https://assets.leetcode.com/uploads/2026/03/21/rbdgrefwok.png" />
    // [0, 1, 1]: We add the edge between vertex 0 and vertex 1 with weight 1.
    // [1, 2, 1]: We add the edge between vertex 1 and vertex 2 with weight 1.
    // [0, 2, 0]: We add the edge between vertex 0 and vertex 2 with weight 0.
    // Note that the cycle 0 - 1 - 2 - 0 has total edge weight 1 + 1 + 0 = 2, which is an even number.
    fmt.Println(numberOfEdgesAdded(3, [][]int{{0,1,1},{1,2,1},{0,2,0}})) // 3

    fmt.Println(numberOfEdgesAdded1(3, [][]int{{0,1,1},{1,2,1},{0,2,1}})) // 2
    fmt.Println(numberOfEdgesAdded1(3, [][]int{{0,1,1},{1,2,1},{0,2,0}})) // 3
}