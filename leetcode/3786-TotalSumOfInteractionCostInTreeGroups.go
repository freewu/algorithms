package main

// 3786. Total Sum of Interaction Cost in Tree Groups
// You are given an integer n and an undirected tree with n nodes numbered from 0 to n - 1. 
// This is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected edge between nodes ui and vi.

// You are also given an integer array group of length n, where group[i] denotes the group label assigned to node i.
//     1. Two nodes u and v are considered part of the same group if group[u] == group[v].
//     2. The interaction cost between u and v is defined as the number of edges on the unique path connecting them in the tree.

// Return an integer denoting the sum of interaction costs over all unordered pairs (u, v) with u != v such that group[u] == group[v].

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], group = [1,1,1]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/09/24/screenshot-2025-09-24-at-50538-pm.png" />
// All nodes belong to group 1. The interaction costs between the pairs of nodes are:
// Nodes (0, 1): 1
// Nodes (1, 2): 1
// Nodes (0, 2): 2
// Thus, the total interaction cost is 1 + 1 + 2 = 4.

// Example 2:
// Input: n = 3, edges = [[0,1],[1,2]], group = [3,2,3]
// Output: 2
// Explanation:
// Nodes 0 and 2 belong to group 3. The interaction cost between this pair is 2.
// Node 1 belongs to a different group and forms no valid pair. Therefore, the total interaction cost is 2.

// Example 3:
// Input: n = 4, edges = [[0,1],[0,2],[0,3]], group = [1,1,4,4]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/09/24/screenshot-2025-09-24-at-51312-pm.png" />
// Nodes belonging to the same groups and their interaction costs are:
// Group 1: Nodes (0, 1): 1
// Group 4: Nodes (2, 3): 2
// Thus, the total interaction cost is 1 + 2 = 3.

// Example 4:
// Input: n = 2, edges = [[0,1]], group = [9,8]
// Output: 0
// Explanation:
// All nodes belong to different groups and there are no valid pairs. Therefore, the total interaction cost is 0.

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi <= n - 1
//     group.length == n
//     1 <= group[i] <= 20
//     The input is generated such that edges represents a valid tree.

import "fmt"

func interactionCosts(n int, edges [][]int, group []int) int64 {
    type Node struct {
        count int
        dist int64
    }
    adj := make([][]int, n)
    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], edge[1])
        adj[edge[1]] = append(adj[edge[1]], edge[0])
    }
    nodes := make([][20]Node, n)
    res := int64(0)
    var dfs func(int, int)
    dfs = func(node, parent int) {
        for _, child := range adj[node] {
            if child == parent { continue }
            dfs(child, node)
            for i := range nodes[node] {
                nodes[node][i].count += nodes[child][i].count
                nodes[node][i].dist += nodes[child][i].dist
            }
        }
        nodes[node][group[node]-1].count++
        for _, child := range adj[node] {
            if child == parent { continue }
            for g := range nodes[node] {
                res += int64(nodes[node][g].count - nodes[child][g].count)*nodes[child][g].dist
            }
        }
        for g := range nodes[node] {
            nodes[node][g].dist += int64(nodes[node][g].count)
        }
    }
    dfs(0, -1)
    return res
}

func interactionCosts1(n int, edges [][]int, group []int) int64 {
    type stackItem struct {
        node    int
        parent  int
        visited bool
    }
    // Build adjacency list
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    // Compute total counts for each group
    var total [21]int
    for _, g := range group {
        total[g]++
    }
    subtreeCounts := make([][21]int, n)
    res := int64(0)
    // Iterative post-order DFS traversal
    stack := []stackItem{{node: 0, parent: -1, visited: false}}
    for len(stack) > 0 {
        item := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        u, parent, visited := item.node, item.parent, item.visited
        if !visited {
            // Push back as visited for later processing
            stack = append(stack, stackItem{u, parent, true})
            // Push children (neighbors excluding parent)
            for _, v := range adj[u] {
                if v != parent {
                    stack = append(stack, stackItem{v, u, false})
                }
            }
        } else {
            // Calculate subtree counts for current node
            var cnt [21]int
            cnt[group[u]] = 1
            for _, v := range adj[u] {
                if v != parent {
                    childCnt := subtreeCounts[v]
                    for g := 1; g <= 20; g++ {
                        cnt[g] += childCnt[g]
                    }
                }
            }
            subtreeCounts[u] = cnt
            // Calculate contribution of the edge to parent if not root
            if parent != -1 {
                var contrib int64 = 0
                for g := 1; g <= 20; g++ {
                    c := cnt[g]
                    total := total[g]
                    contrib += int64(c) * int64(total - c)
                }
                res += contrib
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], group = [1,1,1]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/09/24/screenshot-2025-09-24-at-50538-pm.png" />
    // All nodes belong to group 1. The interaction costs between the pairs of nodes are:
    // Nodes (0, 1): 1
    // Nodes (1, 2): 1
    // Nodes (0, 2): 2
    // Thus, the total interaction cost is 1 + 1 + 2 = 4.
    fmt.Println(interactionCosts(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 4
    // Example 2:
    // Input: n = 3, edges = [[0,1],[1,2]], group = [3,2,3]
    // Output: 2
    // Explanation:
    // Nodes 0 and 2 belong to group 3. The interaction cost between this pair is 2.
    // Node 1 belongs to a different group and forms no valid pair. Therefore, the total interaction cost is 2.
    fmt.Println(interactionCosts(3, [][]int{{0,1},{1,2}}, []int{3,2,3})) // 2
    // Example 3:
    // Input: n = 4, edges = [[0,1],[0,2],[0,3]], group = [1,1,4,4]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/09/24/screenshot-2025-09-24-at-51312-pm.png" />
    // Nodes belonging to the same groups and their interaction costs are:
    // Group 1: Nodes (0, 1): 1
    // Group 4: Nodes (2, 3): 2
    // Thus, the total interaction cost is 1 + 2 = 3.
    fmt.Println(interactionCosts(4, [][]int{{0,1},{0,2},{0,3}}, []int{1,1,4,4})) // 3
    // Example 4:
    // Input: n = 2, edges = [[0,1]], group = [9,8]
    // Output: 0
    // Explanation:
    // All nodes belong to different groups and there are no valid pairs. Therefore, the total interaction cost is 0.
    fmt.Println(interactionCosts(2, [][]int{{0,1}}, []int{9,8})) // 0

    fmt.Println(interactionCosts1(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 4
    fmt.Println(interactionCosts1(3, [][]int{{0,1},{1,2}}, []int{3,2,3})) // 2
    fmt.Println(interactionCosts1(4, [][]int{{0,1},{0,2},{0,3}}, []int{1,1,4,4})) // 3
    fmt.Println(interactionCosts1(2, [][]int{{0,1}}, []int{9,8})) // 0
}