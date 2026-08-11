package main

// 4018. Total Sum of Interaction Cost in Tree Groups II
// You are given an integer n and an undirected tree rooted at node 0 with n nodes numbered from 0 to n - 1. 
// The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected edge between nodes ui and vi.

// You are also given an integer array group of length n, where group[i] denotes the group label assigned to node i.
//     1. Two nodes u and v belong to the same group if and only if group[u] == group[v].
//     2. The interaction cost between two nodes is the shortest distance between them in the tree.

// Return the sum of interaction costs over all pairs of node indices (u, v) such that 0 <= u < v < n and group[u] == group[v].

// The shortest distance between two nodes is the number of edges on the unique path connecting them in the tree.

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], group = [1,1,1]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40329am.png" />
// All nodes belong to group 1. The interaction costs between the pairs of nodes are:
// Nodes [0, 1]: 1
// Nodes [1, 2]: 1
// Nodes [0, 2]: 2
// Thus, the total interaction cost is 1 + 1 + 2 = 4.

// Example 2:
// Input: n = 3, edges = [[0,1],[1,2]], group = [3,2,3]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40416am.png" />
// Nodes 0 and 2 belong to group 3. The interaction cost between this pair is 2.
// Node 1 belongs to a different group and forms no valid pair. Therefore, the total interaction cost is 2.

// Example 3:
// Input: n = 4, edges = [[0,1],[0,2],[0,3]], group = [1,1,4,4]
// Output: 3
// Explanation:
// ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40819am.png" />
// Nodes belonging to the same groups and their interaction costs are:
// Group 1: Nodes [0, 1]: 1
// Group 4: Nodes [2, 3]: 2
// Thus, the total interaction cost is 1 + 2 = 3.

// Example 4:
// Input: n = 2, edges = [[0,1]], group = [1,2]
// Output: 0
// Explanation:
// All nodes belong to different groups and there are no valid pairs. Therefore, the total interaction cost is 0.

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi <= n - 1
//     group.length == n
//     1 <= group[i] <= n
//     The input is generated such that edges represents a valid tree.

import "fmt"

func interactionCosts(n int, edges [][]int, group []int) int64 {
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    depth, subSize,cntLca  := make([]int, n), make([]int, n),  make([]int64, n)
    // first dfs: compute size & depth
    var dfsSize func(u, p int)
    dfsSize = func(u, p int) {
        subSize[u] = 1
        for _, v := range adj[u] {
            if v == p {
                continue
            }
            depth[v] = depth[u] + 1
            dfsSize(v, u)
            subSize[u] += subSize[v]
        }
    }
    dfsSize(0, -1)
    // DSU on tree small‑to‑large
    var dfs func(u, p int, keep bool) map[int]int64
    dfs = func(u, p int, keep bool) map[int]int64 {
        // find heavy child
        heavy, mx := -1, 0
        for _, v := range adj[u] {
            if v != p && subSize[v] > mx {
                mx, heavy = subSize[v], v
            }
        }
        res := make(map[int]int64)
        if heavy != -1 {
            res = dfs(heavy, u, true)
        } else {
            res = make(map[int]int64)
        }
        // iterate light children
        for _, v := range adj[u] {
            if v == p || v == heavy {
                continue
            }
            childMap := dfs(v, u, false)
            // count cross pairs between res and childMap: these pairs' lca is u
            for g, c := range childMap {
                cntLca[u] += res[g] * c
            }
            // merge childMap into res
            for g, c := range childMap {
                res[g] += c
            }
        }
        // add current node u itself
        gU := group[u]
        cntLca[u] += res[gU] // pairs between u and existing nodes in subtree, lca=u
        res[gU]++
        if !keep {
            // discard map
            return nil
        }
        return res
    }
    dfs(0, -1, true)
    // compute first term sum_{u<v} (depth[u]+depth[v])
    groups := make(map[int][]int)
    for i := 0; i < n; i++ {
        g := group[i]
        groups[g] = append(groups[g], i)
    }
    sum := int64(0)
    for _, nodes := range groups {  
        m := len(nodes)
        if m < 2 {
            continue
        }
        s := int64(0)
        for _, nd := range nodes {
            s += int64(depth[nd])
        }
        sum += s * int64(m-1)
    }
    sumLca := int64(0)
    for x := 0; x < n; x++ {
        sumLca += cntLca[x] * int64(depth[x])
    }
    return sum - 2 * sumLca
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], group = [1,1,1]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40329am.png" />
    // All nodes belong to group 1. The interaction costs between the pairs of nodes are:
    // Nodes [0, 1]: 1
    // Nodes [1, 2]: 1
    // Nodes [0, 2]: 2
    // Thus, the total interaction cost is 1 + 1 + 2 = 4.
    fmt.Println(interactionCosts(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 4
    // Example 2:
    // Input: n = 3, edges = [[0,1],[1,2]], group = [3,2,3]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40416am.png" />
    // Nodes 0 and 2 belong to group 3. The interaction cost between this pair is 2.
    // Node 1 belongs to a different group and forms no valid pair. Therefore, the total interaction cost is 2.
    fmt.Println(interactionCosts(3, [][]int{{0,1},{1,2}}, []int{3,2,3})) // 2
    // Example 3:
    // Input: n = 4, edges = [[0,1],[0,2],[0,3]], group = [1,1,4,4]
    // Output: 3
    // Explanation:
    // ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40819am.png" />
    // Nodes belonging to the same groups and their interaction costs are:
    // Group 1: Nodes [0, 1]: 1
    // Group 4: Nodes [2, 3]: 2
    // Thus, the total interaction cost is 1 + 2 = 3.
    fmt.Println(interactionCosts(4, [][]int{{0,1},{0,2},{0,3}}, []int{1,1,4,4})) // 3
    // Example 4:
    // Input: n = 2, edges = [[0,1]], group = [1,2]
    // Output: 0
    // Explanation:
    // All nodes belong to different groups and there are no valid pairs. Therefore, the total interaction cost is 0.
    fmt.Println(interactionCosts(2, [][]int{{0,1}}, []int{1,2})) // 0
}