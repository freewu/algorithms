package main

// 3559. Number of Ways to Assign Edge Weights II
// There is an undirected tree with n nodes labeled from 1 to n, rooted at node 1. 
// The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi.

// Initially, all edges have a weight of 0. You must assign each edge a weight of either 1 or 2.

// The cost of a path between any two nodes u and v is the total weight of all edges in the path connecting them.

// You are given a 2D integer array queries. 
// For each queries[i] = [ui, vi], determine the number of ways to assign weights to edges in the path such that the cost of the path between ui and vi is odd.

// Return an array answer, where answer[i] is the number of valid assignments for queries[i].

// Since the answer may be large, apply modulo 10^9 + 7 to each answer[i].

// Note: For each query, disregard all edges not in the path between node ui and vi.

// Example 1:
// <img src="https://pic.leetcode.cn/1748074049-lsGWuV-screenshot-2025-03-24-at-060006.png" />
// Input: edges = [[1,2]], queries = [[1,1],[1,2]]
// Output: [0,1]
// Explanation:
// Query [1,1]: The path from Node 1 to itself consists of no edges, so the cost is 0. Thus, the number of valid assignments is 0.
// Query [1,2]: The path from Node 1 to Node 2 consists of one edge (1 → 2). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.

// Example 2:
// <img src="https://pic.leetcode.cn/1748074095-sRyffx-screenshot-2025-03-24-at-055820.png" />
// Input: edges = [[1,2],[1,3],[3,4],[3,5]], queries = [[1,4],[3,4],[2,5]]
// Output: [2,1,4]
// Explanation:
// Query [1,4]: The path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4). Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.
// Query [3,4]: The path from Node 3 to Node 4 consists of one edge (3 → 4). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
// Query [2,5]: The path from Node 2 to Node 5 consists of three edges (2 → 1, 1 → 3, and 3 → 5). Assigning (1,2,2), (2,1,2), (2,2,1), or (1,1,1) makes the cost odd. Thus, the number of valid assignments is 4.

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] == [ui, vi]
//     1 <= queries.length <= 10^5
//     queries[i] == [ui, vi]
//     1 <= ui, vi <= n
//     edges represents a valid tree.

import "fmt"
import "math/bits"

const mod = 1_000_000_007
const mx = 17

var pow2 = [1e5]int{1}

func init() {
    for i := 1; i < len(pow2); i++ {
        pow2[i] = pow2[i - 1] * 2 % mod // 预处理 2 的幂
    }
}

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
    n := len(edges) + 1
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0]-1, e[1]-1
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    pa := make([][mx]int, n)
    dep := make([]int, n)
    var dfs func(int, int)
    dfs = func(x, p int) {
        pa[x][0] = p
        for _, y := range g[x] {
            if y != p {
                dep[y] = dep[x] + 1
                dfs(y, x)
            }
        }
    }
    dfs(0, -1)
    for i := 0; i < mx - 1; i++ {
        for x := range pa {
            if p := pa[x][i]; p != -1 {
                pa[x][i+1] = pa[p][i]
            } else {
                pa[x][i+1] = -1
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
    getDis := func(x, y int) int { return dep[x] + dep[y] - dep[getLCA(x, y)] * 2 }
    res := make([]int, len(queries))
    for i, q := range queries {
        if q[0] != q[1] {
            res[i] = pow2[getDis(q[0]-1, q[1]-1)-1]
        }
    }
    return res
}

// 超出时间限制 585 / 589
func assignEdgeWeights1(edges [][]int, queries [][]int) []int {
    n, mod:= len(edges) + 1, 1_000_000_007
    // 邻接表存储树结构
    adj := make([][]int, n+1)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    // 记录每个节点的父节点
    parent, depth := make([]int, n + 1), make([]int, n + 1)
    // 预处理每个节点的深度和父节点
    var dfs func(u, p, dep int) 
    dfs = func(u, p, dep int) {
        parent[u], depth[u] = p, dep
        for _, v := range adj[u] {
            if v != p {
                dfs(v, u, dep + 1)
            }
        }
    }
    dfs(1, 0, 0)
    pathLen := func(u, v int) int { // 计算两个节点之间路径上的边数
        if depth[u] < depth[v] {
            u, v = v, u
        }
        diff := depth[u] - depth[v]
        for i := 0; i < diff; i++ {
            u = parent[u]
        }
        if u == v {
            return diff
        }
        for u != v {
            u = parent[u]
            v = parent[v]
            diff += 2
        }
        return diff
    }
    pow := func(a, b int) int {
        res := 1
        for b > 0 {
            if b & 1 == 1 {
                res = res * a % mod
            }
            a = a * a % mod
            b >>= 1
        }
        return res
    }
    res := make([]int, len(queries))
    for i, query := range queries {
        u, v := query[0], query[1]
        if u == v {
            res[i] = 0
            continue
        }
        len := pathLen(u, v)
        if len % 2 == 0 {
            res[i] = pow(2, len-1) % mod
        } else {
            res[i] = pow(2, len-1) % mod
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://pic.leetcode.cn/1748074049-lsGWuV-screenshot-2025-03-24-at-060006.png" />
    // Input: edges = [[1,2]], queries = [[1,1],[1,2]]
    // Output: [0,1]
    // Explanation:
    // Query [1,1]: The path from Node 1 to itself consists of no edges, so the cost is 0. Thus, the number of valid assignments is 0.
    // Query [1,2]: The path from Node 1 to Node 2 consists of one edge (1 → 2). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
    fmt.Println(assignEdgeWeights([][]int{{1,2}}, [][]int{{1,1},{1,2}})) // [0,1]
    // Example 2:
    // <img src="https://pic.leetcode.cn/1748074095-sRyffx-screenshot-2025-03-24-at-055820.png" />
    // Input: edges = [[1,2],[1,3],[3,4],[3,5]], queries = [[1,4],[3,4],[2,5]]
    // Output: [2,1,4]
    // Explanation:
    // Query [1,4]: The path from Node 1 to Node 4 consists of two edges (1 → 3 and 3 → 4). Assigning weights (1,2) or (2,1) results in an odd cost. Thus, the number of valid assignments is 2.
    // Query [3,4]: The path from Node 3 to Node 4 consists of one edge (3 → 4). Assigning weight 1 makes the cost odd, while 2 makes it even. Thus, the number of valid assignments is 1.
    // Query [2,5]: The path from Node 2 to Node 5 consists of three edges (2 → 1, 1 → 3, and 3 → 5). Assigning (1,2,2), (2,1,2), (2,2,1), or (1,1,1) makes the cost odd. Thus, the number of valid assignments is 4.
    fmt.Println(assignEdgeWeights([][]int{{1,2},{1,3},{3,4},{3,5}}, [][]int{{1,4},{3,4},{2,5}})) // [2,1,4]

    fmt.Println(assignEdgeWeights1([][]int{{1,2}}, [][]int{{1,1},{1,2}})) // [0,1]
    fmt.Println(assignEdgeWeights1([][]int{{1,2},{1,3},{3,4},{3,5}}, [][]int{{1,4},{3,4},{2,5}})) // [2,1,4]
}