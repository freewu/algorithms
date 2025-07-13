package main

// 3613. Minimize Maximum Component Cost
// You are given an undirected connected graph with n nodes labeled from 0 to n - 1 and a 2D integer array edges where edges[i] = [ui, vi, wi] denotes an undirected edge between node ui and node vi with weight wi, and an integer k.

// You are allowed to remove any number of edges from the graph such that the resulting graph has at most k connected components.

// The cost of a component is defined as the maximum edge weight in that component. 
// If a component has no edges, its cost is 0.

// Return the minimum possible value of the maximum cost among all components after such removals.

// Example 1:
// Input: n = 5, edges = [[0,1,4],[1,2,3],[1,3,2],[3,4,6]], k = 2
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/19/minimizemaximumm.jpg" />
// Remove the edge between nodes 3 and 4 (weight 6).
// The resulting components have costs of 0 and 4, so the overall maximum cost is 4.

// Example 2:
// Input: n = 4, edges = [[0,1,5],[1,2,5],[2,3,5]], k = 1
// Output: 5
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/19/minmax2.jpg" />
// No edge can be removed, since allowing only one component (k = 1) requires the graph to stay fully connected.
// That single component’s cost equals its largest edge weight, which is 5.

// Constraints:
//     1 <= n <= 5 * 10^4
//     0 <= edges.length <= 10^5
//     edges[i].length == 3
//     0 <= ui, vi < n
//     1 <= wi <= 10^6
//     1 <= k <= n
//     The input graph is connected.

import "fmt"
import "sort"

type UnionFind struct {
    parent []int
    rank []int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    return &UnionFind{parent, make([]int, n)}
}

func (u *UnionFind) find(x int) int {
    if u.parent[x] != x {
        u.parent[x] = u.find(u.parent[x])
    }
    return u.parent[x]
}

func (u *UnionFind) union(x, y int) bool {
    rootX := u.find(x)
    rootY := u.find(y)
    if rootX == rootY { return false }
    if u.rank[rootX] < u.rank[rootY] {
        u.parent[rootX] = rootY
    } else if u.rank[rootX] > u.rank[rootY] {
        u.parent[rootY] = rootX
    } else {
        u.parent[rootY] = rootX
        u.rank[rootX] += 1
    }
    return true
}

func minCost(n int, edges [][]int, k int) int {
    if len(edges) == 0 { return 0 }
    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })
    uf := NewUnionFind(n)
    mst := []int{}
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        if uf.union(u, v) {
           mst = append(mst, w) 
        }
    }
    sort.Sort(sort.Reverse(sort.IntSlice(mst)))
    if len(mst) < k-1 {
        return 0
    }
    if k == 1 {
        return mst[0]
    }
    mst = mst[k-1:]
    if len(mst) == 0 {
        return 0
    }
    return mst[0]
}


type UnionFind1 struct {
    fa   []int
    size int
}

func newUnionFind1(n int) UnionFind1 {
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }
    return UnionFind1{fa, n}
}

func (uf UnionFind1) find(x int) int {
    if uf.fa[x] != x {
        uf.fa[x] = uf.find(uf.fa[x])
    }
    return uf.fa[x]
}

func (uf *UnionFind1) union(x, y int) {
    rootX, rootY := uf.find(x), uf.find(y)
    if rootX == rootY { return }
    uf.fa[rootX] = rootY
    uf.size--
}

func minCost1(n int, edges [][]int, k int) int {
    if n == k { return 0 }
    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })
    uf := newUnionFind1(n)
    for _, e := range edges {
        uf.union(e[0], e[1])
        if uf.size == k {
            return e[2]
        }
    }
    return 0
}

func minCost2(n int, edges [][]int, k int) int {
    if k == n { return 0 }
    count := n
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    var find func(x int) int
    find = func(x int) int {
        if parent[x] == x { return x }
        parent[x] = find(parent[x])
        return parent[x]
    }
    union := func(x, y int) {
        pX, pY := find(x), find(y)
        if pX != pY {
            count--
            parent[pX] = pY
        }
    }
    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })
    for _, e := range edges {
        union(e[0], e[1])
        if count <= k {
            return e[2]
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 5, edges = [[0,1,4],[1,2,3],[1,3,2],[3,4,6]], k = 2
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/19/minimizemaximumm.jpg" />
    // Remove the edge between nodes 3 and 4 (weight 6).
    // The resulting components have costs of 0 and 4, so the overall maximum cost is 4.
    fmt.Println(minCost(5,[][]int{{0,1,4},{1,2,3},{1,3,2},{3,4,6}}, 2)) // 4
    // Example 2:
    // Input: n = 4, edges = [[0,1,5],[1,2,5],[2,3,5]], k = 1
    // Output: 5
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/19/minmax2.jpg" />
    // No edge can be removed, since allowing only one component (k = 1) requires the graph to stay fully connected.
    // That single component’s cost equals its largest edge weight, which is 5.
    fmt.Println(minCost(4,[][]int{{0,1,5},{1,2,5},{2,3,5}}, 1)) // 5

    fmt.Println(minCost1(5,[][]int{{0,1,4},{1,2,3},{1,3,2},{3,4,6}}, 2)) // 4
    fmt.Println(minCost1(4,[][]int{{0,1,5},{1,2,5},{2,3,5}}, 1)) // 5

    fmt.Println(minCost2(5,[][]int{{0,1,4},{1,2,3},{1,3,2},{3,4,6}}, 2)) // 4
    fmt.Println(minCost2(4,[][]int{{0,1,5},{1,2,5},{2,3,5}}, 1)) // 5
}