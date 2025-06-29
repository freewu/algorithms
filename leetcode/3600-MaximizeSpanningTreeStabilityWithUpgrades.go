package main

// 3600. Maximize Spanning Tree Stability with Upgrades
// You are given an integer n, representing n nodes numbered from 0 to n - 1 and a list of edges, 
// where edges[i] = [ui, vi, si, musti]:

//     1. ui and vi indicates an undirected edge between nodes ui and vi.
//     2. si is the strength of the edge.
//     3. musti is an integer (0 or 1). If musti == 1, the edge must be included in the spanning tree. 
//        These edges cannot be upgraded.

// You are also given an integer k, the maximum number of upgrades you can perform. 
// Each upgrade doubles the strength of an edge, and each eligible edge (with musti == 0) can be upgraded at most once.

// The stability of a spanning tree is defined as the minimum strength score among all edges included in it.

// Return the maximum possible stability of any valid spanning tree. 
// If it is impossible to connect all nodes, return -1.

// Note: A spanning tree of a graph with n nodes is a subset of the edges that connects all nodes together (i.e. the graph is connected) without forming any cycles, and uses exactly n - 1 edges.

// Example 1:
// Input: n = 3, edges = [[0,1,2,1],[1,2,3,0]], k = 1
// Output: 2
// Explanation:
// Edge [0,1] with strength = 2 must be included in the spanning tree.
// Edge [1,2] is optional and can be upgraded from 3 to 6 using one upgrade.
// The resulting spanning tree includes these two edges with strengths 2 and 6.
// The minimum strength in the spanning tree is 2, which is the maximum possible stability.

// Example 2:
// Input: n = 3, edges = [[0,1,4,0],[1,2,3,0],[0,2,1,0]], k = 2
// Output: 6
// Explanation:
// Since all edges are optional and up to k = 2 upgrades are allowed.
// Upgrade edges [0,1] from 4 to 8 and [1,2] from 3 to 6.
// The resulting spanning tree includes these two edges with strengths 8 and 6.
// The minimum strength in the tree is 6, which is the maximum possible stability.

// Example 3:
// Input: n = 3, edges = [[0,1,1,1],[1,2,1,1],[2,0,1,1]], k = 0
// Output: -1
// Explanation:
// All edges are mandatory and form a cycle, which violates the spanning tree property of acyclicity. 
// Thus, the answer is -1.

// Constraints:
//     2 <= n <= 10^5
//     1 <= edges.length <= 10^5
//     edges[i] = [ui, vi, si, musti]
//     0 <= ui, vi < n
//     ui != vi
//     1 <= si <= 10^5
//     musti is either 0 or 1.
//     0 <= k <= n
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
// 返回是否合并成功
func (u *UnionFind) merge(from, to int) bool {
    x, y := u.find(from), u.find(to)
    if x == y { // from 和 to 在同一个集合，不做合并
        return false
    }
    u.fa[x] = y // 合并集合。修改后就可以认为 from 和 to 在同一个集合了
    u.cc--      // 成功合并，连通块个数减一
    return true
}

// 根据 Kruskal 算法求最大生成树，把剩余的边按照边权（先不乘 2）从大到小合并
func maxStability(n int, edges [][]int, k int) int {
    uf,all := newUnionFind(n), newUnionFind(n)
    mn := 1 << 31
    for _, e := range edges {
        x, y, s, must := e[0], e[1], e[2], e[3]
        if must > 0 {
            if !uf.merge(x, y) { return -1 } // 必选边成环
            mn = min(mn, s)
        }
        all.merge(x, y)
    }
    if all.cc > 1 { return -1 } // 图不连通
    if uf.cc == 1 { return mn } // 只需选必选边
    // Kruskal 算法求最大生成树
    slices.SortFunc(edges, func(a, b []int) int { return b[2] - a[2] })
    arr := []int{}
    for _, e := range edges {
        x, y, s, must := e[0], e[1], e[2], e[3]
        if must == 0 && uf.merge(x, y) {
            arr = append(arr, s)
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    // 答案为如下三者的最小值：
    // 1. must = 1 中的最小边权
    // 2. arr 中最小边权 * 2
    // 3. arr 中第 k+1 小边权
    m := len(arr)
    res := min(mn, arr[m-1] * 2)
    if k < m {
        res = min(res, arr[m-1-k])
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1,2,1],[1,2,3,0]], k = 1
    // Output: 2
    // Explanation:
    // Edge [0,1] with strength = 2 must be included in the spanning tree.
    // Edge [1,2] is optional and can be upgraded from 3 to 6 using one upgrade.
    // The resulting spanning tree includes these two edges with strengths 2 and 6.
    // The minimum strength in the spanning tree is 2, which is the maximum possible stability.
    fmt.Println(maxStability(3, [][]int{{0,1,2,1},{1,2,3,0}}, 1)) // 2
    // Example 2:
    // Input: n = 3, edges = [[0,1,4,0],[1,2,3,0],[0,2,1,0]], k = 2
    // Output: 6
    // Explanation:
    // Since all edges are optional and up to k = 2 upgrades are allowed.
    // Upgrade edges [0,1] from 4 to 8 and [1,2] from 3 to 6.
    // The resulting spanning tree includes these two edges with strengths 8 and 6.
    // The minimum strength in the tree is 6, which is the maximum possible stability.
    fmt.Println(maxStability(3, [][]int{{0,1,4,0},{1,2,3,0},{0,2,1,0}}, 2)) // 6
    // Example 3:
    // Input: n = 3, edges = [[0,1,1,1],[1,2,1,1],[2,0,1,1]], k = 0
    // Output: -1
    // Explanation:
    // All edges are mandatory and form a cycle, which violates the spanning tree property of acyclicity. 
    // Thus, the answer is -1.
    fmt.Println(maxStability(3, [][]int{{0,1,1,1},{1,2,1,1},{2,0,1,1}}, 0)) // -1
}