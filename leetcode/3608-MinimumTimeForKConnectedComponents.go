package main

// 3608. Minimum Time for K Connected Components
// You are given an integer n and an undirected graph with n nodes labeled from 0 to n - 1. 
// This is represented by a 2D array edges, where edges[i] = [ui, vi, timei] indicates an undirected edge between nodes ui and vi that can be removed at timei.

// You are also given an integer k.

// Initially, the graph may be connected or disconnected. 
// Your task is to find the minimum time t such that after removing all edges with time <= t, the graph contains at least k connected components.

// Return the minimum time t.

// A connected component is a subgraph of a graph in which there exists a path between any two vertices, and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

// Example 1:
// Input: n = 2, edges = [[0,1,3]], k = 2
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/31/screenshot-2025-06-01-at-022724.png" />
// Initially, there is one connected component {0, 1}.
// At time = 1 or 2, the graph remains unchanged.
// At time = 3, edge [0, 1] is removed, resulting in k = 2 connected components {0}, {1}. Thus, the answer is 3.

// Example 2:
// Input: n = 3, edges = [[0,1,2],[1,2,4]], k = 3
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/31/screenshot-2025-06-01-at-022812.png" />
// Initially, there is one connected component {0, 1, 2}.
// At time = 2, edge [0, 1] is removed, resulting in two connected components {0}, {1, 2}.
// At time = 4, edge [1, 2] is removed, resulting in k = 3 connected components {0}, {1}, {2}. Thus, the answer is 4.

// Example 3:
// Input: n = 3, edges = [[0,2,5]], k = 2
// Output: 0
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/05/31/screenshot-2025-06-01-at-022930.png" />
// Since there are already k = 2 disconnected components {1}, {0, 2}, no edge removal is needed. Thus, the answer is 0.
 
// Constraints:
//     1 <= n <= 10^5
//     0 <= edges.length <= 10^5
//     edges[i] = [ui, vi, timei]
//     0 <= ui, vi < n
//     ui != vi
//     1 <= timei <= 10^9
//     1 <= k <= n
//     There are no duplicate edges.

import "fmt"
import "slices"

type UnionFind struct {
    fa []int // 代表元
    cc int   // 连通块个数
}

func newUnionFind(n int) UnionFind {
    fa := make([]int, n)
    // 一开始有 n 个集合 {0}, {1}, ..., {n-1}
    // 集合 i 的代表元是自己
    for i := range fa {
        fa[i] = i
    }
    return UnionFind{fa, n}
}

// 返回 x 所在集合的代表元
// 同时做路径压缩，也就是把 x 所在集合中的所有元素的 fa 都改成代表元
func (u UnionFind) find(x int) int {
    // 如果 fa[x] == x，则表示 x 是代表元
    if u.fa[x] != x {
        u.fa[x] = u.find(u.fa[x]) // fa 改成代表元
    }
    return u.fa[x]
}

// 把 from 所在集合合并到 to 所在集合中
func (u *UnionFind) merge(from, to int) {
    x, y := u.find(from), u.find(to)
    if x == y { return } // from 和 to 在同一个集合，不做合并
    u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
    u.cc--      // 成功合并，连通块个数减一
}

func minTime(n int, edges [][]int, k int) int {
    slices.SortFunc(edges, func(a, b []int) int { return b[2] - a[2] })
    u := newUnionFind(n)
    for _, e := range edges {
        u.merge(e[0], e[1])
        if u.cc < k { // 这条边不能留，即移除所有 time <= e[2] 的边
            return e[2]
        }
    }
    return 0 // 无需移除任何边
}

func main() {
    // Example 1:
    // Input: n = 2, edges = [[0,1,3]], k = 2
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/31/screenshot-2025-06-01-at-022724.png" />
    // Initially, there is one connected component {0, 1}.
    // At time = 1 or 2, the graph remains unchanged.
    // At time = 3, edge [0, 1] is removed, resulting in k = 2 connected components {0}, {1}. Thus, the answer is 3.
    fmt.Println(minTime(2, [][]int{{0,1,3}}, 2)) // 3
    // Example 2:
    // Input: n = 3, edges = [[0,1,2],[1,2,4]], k = 3
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/31/screenshot-2025-06-01-at-022812.png" />
    // Initially, there is one connected component {0, 1, 2}.
    // At time = 2, edge [0, 1] is removed, resulting in two connected components {0}, {1, 2}.
    // At time = 4, edge [1, 2] is removed, resulting in k = 3 connected components {0}, {1}, {2}. Thus, the answer is 4.
    fmt.Println(minTime(3, [][]int{{0,1,2},{1,2,4}}, 3)) // 4
    // Example 3:
    // Input: n = 3, edges = [[0,2,5]], k = 2
    // Output: 0
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/05/31/screenshot-2025-06-01-at-022930.png" />
    // Since there are already k = 2 disconnected components {1}, {0, 2}, no edge removal is needed. Thus, the answer is 0.
    fmt.Println(minTime(3, [][]int{{0,2,5}}, 2)) // 0
}