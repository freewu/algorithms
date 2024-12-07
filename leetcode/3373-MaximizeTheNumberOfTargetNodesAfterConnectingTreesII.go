package main

// 3373. Maximize the Number of Target Nodes After Connecting Trees II
// There exist two undirected trees with n and m nodes, labeled from [0, n - 1] and [0, m - 1], respectively.

// You are given two 2D integer arrays edges1 and edges2 of lengths n - 1 and m - 1, 
// respectively, where edges1[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the first tree and edges2[i] = [ui, vi] indicates 
// that there is an edge between nodes ui and vi in the second tree.

// Node u is target to node v if the number of edges on the path from u to v is even. 
// Note that a node is always target to itself.

// Return an array of n integers answer, 
// where answer[i] is the maximum possible number of nodes that are target to node i of the first tree if you had to connect one node from the first tree to another node in the second tree.

// Note that queries are independent from each other. 
// That is, for every query you will remove the added edge before proceeding to the next query.

// Example 1:
// Input: edges1 = [[0,1],[0,2],[2,3],[2,4]], edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]]
// Output: [8,7,7,8,8]
// Explanation:
// For i = 0, connect node 0 from the first tree to node 0 from the second tree.
// For i = 1, connect node 1 from the first tree to node 4 from the second tree.
// For i = 2, connect node 2 from the first tree to node 7 from the second tree.
// For i = 3, connect node 3 from the first tree to node 0 from the second tree.
// For i = 4, connect node 4 from the first tree to node 4 from the second tree.

// Example 2:
// Input: edges1 = [[0,1],[0,2],[0,3],[0,4]], edges2 = [[0,1],[1,2],[2,3]]
// Output: [3,6,6,6,6]
// Explanation:
// For every i, connect node i of the first tree with any node of the second tree.

// Constraints:
//     2 <= n, m <= 10^5
//     edges1.length == n - 1
//     edges2.length == m - 1
//     edges1[i].length == edges2[i].length == 2
//     edges1[i] = [ai, bi]
//     0 <= ai, bi < n
//     edges2[i] = [ui, vi]
//     0 <= ui, vi < m
//     The input is generated such that edges1 and edges2 represent valid trees.

import "fmt"

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
    buildGraph := func(edges [][]int) (int, [][]int) {
        n := len(edges) + 1
        graph := make([][]int, n)
        for _, v := range edges {
            graph[v[0]], graph[v[1]] = append(graph[v[0]], v[1]), append(graph[v[1]], v[0])
        }
        return n, graph
    }
    n, graph1 := buildGraph(edges1)
    m, graph2 := buildGraph(edges2)
    helper := func(graph [][]int) ([]int, int) {
        n := len(graph)
        distance := make([]int, n)
        var dfs func(u, fa, i int)
        dfs = func(u, fa, i int) {
            distance[u] = i
            for _, v := range graph[u] {
                if v != fa {
                    dfs(v, u, i^1)
                }
            }
        }
        dfs(0, -1, 0)
        count := 0
        for _, v := range distance {
            count += v
        }
        return distance, count
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, count1 := helper(graph1)
        _,  count2 := helper(graph2)
    count2 = max(count2, m - count2)
    for i, v := range res {
        if v == 0 {
            res[i] = n - count1 + count2
        } else {
            res[i] = count1 + count2
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: edges1 = [[0,1],[0,2],[2,3],[2,4]], edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]]
    // Output: [8,7,7,8,8]
    // Explanation:
    // For i = 0, connect node 0 from the first tree to node 0 from the second tree.
    // For i = 1, connect node 1 from the first tree to node 4 from the second tree.
    // For i = 2, connect node 2 from the first tree to node 7 from the second tree.
    // For i = 3, connect node 3 from the first tree to node 0 from the second tree.
    // For i = 4, connect node 4 from the first tree to node 4 from the second tree.
    fmt.Println(maxTargetNodes([][]int{{0,1},{0,2},{2,3},{2,4}}, [][]int{{0,1},{0,2},{0,3},{2,7},{1,4},{4,5},{4,6}})) // [8,7,7,8,8]
    // Example 2:
    // Input: edges1 = [[0,1],[0,2],[0,3],[0,4]], edges2 = [[0,1],[1,2],[2,3]]
    // Output: [3,6,6,6,6]
    // Explanation:
    // For every i, connect node i of the first tree with any node of the second tree.
    fmt.Println(maxTargetNodes([][]int{{0,1},{0,2},{0,3},{0,4}}, [][]int{{0,1},{1,2},{2,3}})) // [3,6,6,6,6]
}