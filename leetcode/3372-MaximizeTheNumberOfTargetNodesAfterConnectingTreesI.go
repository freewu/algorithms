package main

// 3372. Maximize the Number of Target Nodes After Connecting Trees I
// There exist two undirected trees with n and m nodes, 
// with distinct labels in ranges [0, n - 1] and [0, m - 1], respectively.

// You are given two 2D integer arrays edges1 and edges2 of lengths n - 1 and m - 1, 
// respectively, where edges1[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the first tree and edges2[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the second tree. 
// You are also given an integer k.

// Node u is target to node v if the number of edges on the path from u to v is less than or equal to k. 
// Note that a node is always target to itself.

// Return an array of n integers answer, 
// where answer[i] is the maximum possible number of nodes target to node i of the first tree if you have to connect one node from the first tree to another node in the second tree.

// Note that queries are independent from each other. 
// That is, for every query you will remove the added edge before proceeding to the next query.

// Example 1:
// Input: edges1 = [[0,1],[0,2],[2,3],[2,4]], edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]], k = 2
// Output: [9,7,9,8,8]
// Explanation:
// For i = 0, connect node 0 from the first tree to node 0 from the second tree.
// For i = 1, connect node 1 from the first tree to node 0 from the second tree.
// For i = 2, connect node 2 from the first tree to node 4 from the second tree.
// For i = 3, connect node 3 from the first tree to node 4 from the second tree.
// For i = 4, connect node 4 from the first tree to node 4 from the second tree.

// Example 2:
// Input: edges1 = [[0,1],[0,2],[0,3],[0,4]], edges2 = [[0,1],[1,2],[2,3]], k = 1
// Output: [6,3,3,3,3]
// Explanation:
// For every i, connect node i of the first tree with any node of the second tree.

// Constraints:
//     2 <= n, m <= 1000
//     edges1.length == n - 1
//     edges2.length == m - 1
//     edges1[i].length == edges2[i].length == 2
//     edges1[i] = [ai, bi]
//     0 <= ai, bi < n
//     edges2[i] = [ui, vi]
//     0 <= ui, vi < m
//     The input is generated such that edges1 and edges2 represent valid trees.
//     0 <= k <= 1000

import "fmt"

// // 解答错误 372 / 817
// func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
//     n, m := len(edges1) + 1, len(edges2) + 1
//     tree1, tree2 := make([][]int, n), make([][]int, m)
//     for i := range tree1 { tree1[i] = make([]int, n) }
//     for i := range tree2 { tree2[i] = make([]int, m) }
//     for _, v := range edges1 {
//         tree1[v[0]], tree1[v[1]] = append(tree1[v[0]], v[1]), append(tree1[v[1]], v[0])
//     }
//     for _, v := range edges2 {
//         tree2[v[0]], tree2[v[1]] = append(tree2[v[0]], v[1]), append(tree2[v[1]], v[0])
//     }
//     var dfs func(graph [][]int, src int, k int, visited []bool) int
//     dfs = func(graph [][]int, src int, k int, visited []bool) int {
//         if k < 0 { return 0 }
//         res := 1
//         for _, nbr := range graph[src] {
//             if !visited[nbr] {
//                 visited[nbr] = true
//                 res += dfs(graph, nbr, k - 1, visited)
//             }
//         }
//         return res
//     }
//     res, temp := make([]int, n), make([]int, n)
//     for i := 0; i < n; i++ {
//         visited := make([]bool, n)
//         visited[i] = true
//         temp[i] = dfs(tree1, i, k, visited)
//     }
//     mx := 0
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for i := 0; i < m; i++  {
//         visited := make([]bool, m)
//         visited[i] = true
//         mx = max(mx, dfs(tree2, i, k - 1, visited))
//     }
//     for i := 0; i < n; i++  {
//         res[i] = temp[i] + mx
//     }
//     return res
// }

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
    buildGraph := func(edges [][]int) [][]int {
        graph := make([][]int, len(edges) + 1)
        for _, v := range edges {
            graph[v[0]], graph[v[1]] = append(graph[v[0]], v[1]), append(graph[v[1]], v[0])
        }
        return graph
    }
    var dfs func(graph [][]int, root int, par int, k int, count int) int 
    dfs = func(graph [][]int, root int, par int, k int, count int) int {
        if k < 0 { return 0 }
        for _, node := range graph[root] {
            if node != par {
                count += dfs(graph, node, root, k - 1, 1)
            }
        }
        return count
    }
    graph1, graph2 := buildGraph(edges1), buildGraph(edges2)
    count, n, m := 0,  len(edges1) + 1, len(edges2) + 1
    res := []int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < m; i++ {
        count = max(count, dfs(graph2, i, -1, k - 1, 1))
    } 
    for i := 0; i < n; i++ {
        res = append(res, count + dfs(graph1, i, -1, k, 1))
    }
    return res
}

func maxTargetNodes1(edges1 [][]int, edges2 [][]int, k int) []int {
    buildTree := func(edges [][]int, k int) func(int, int, int) int {
        graph := make([][]int, len(edges) + 1)
        for _, v := range edges {
            graph[v[0]], graph[v[1]] = append(graph[v[0]], v[1]), append(graph[v[1]], v[0])
        }
        var dfs func(x, fa, d int) int
        dfs = func(x, fa, d int) int {
            if d > k { return 0 }
            count := 1
            for _, y := range graph[x] {
                if y != fa { count += dfs(y, x, d + 1) }
            }
            return count
        }
        return dfs
    }
    mx := 0
    if k > 0 {
        dfs := buildTree(edges2, k-1) // 注意这里传的是 k-1
        for i := 0; i < len(edges2) + 1; i++ {
            mx = max(mx, dfs(i, -1, 0))
        }
    }
    dfs := buildTree(edges1, k)
    res := make([]int, len(edges1) + 1)
    for i := range res {
        res[i] = dfs(i, -1, 0) + mx
    }
    return res
}

func main() {
    // Example 1:
    // Input: edges1 = [[0,1],[0,2],[2,3],[2,4]], edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]], k = 2
    // Output: [9,7,9,8,8]
    // Explanation:
    // For i = 0, connect node 0 from the first tree to node 0 from the second tree.
    // For i = 1, connect node 1 from the first tree to node 0 from the second tree.
    // For i = 2, connect node 2 from the first tree to node 4 from the second tree.
    // For i = 3, connect node 3 from the first tree to node 4 from the second tree.
    // For i = 4, connect node 4 from the first tree to node 4 from the second tree.
    fmt.Println(maxTargetNodes([][]int{{0,1},{0,2},{2,3},{2,4}}, [][]int{{0,1},{0,2},{0,3},{2,7},{1,4},{4,5},{4,6}}, 2)) // [9,7,9,8,8]
    // Example 2:
    // Input: edges1 = [[0,1],[0,2],[0,3],[0,4]], edges2 = [[0,1],[1,2],[2,3]], k = 1
    // Output: [6,3,3,3,3]
    // Explanation:
    // For every i, connect node i of the first tree with any node of the second tree.
    fmt.Println(maxTargetNodes([][]int{{0,1},{0,2},{0,3},{0,4}}, [][]int{{0,1},{1,2},{2,3}}, 1)) // [6,3,3,3,3]

    fmt.Println(maxTargetNodes([][]int{{2,0},{3,1},{3,2},{3,4}}, [][]int{{0,3},{0,4},{2,5},{0,2},{7,0},{1,6},{1,7}}, 1)) // [3,3,4,5,3]

    fmt.Println(maxTargetNodes1([][]int{{0,1},{0,2},{2,3},{2,4}}, [][]int{{0,1},{0,2},{0,3},{2,7},{1,4},{4,5},{4,6}}, 2)) // [9,7,9,8,8]
    fmt.Println(maxTargetNodes1([][]int{{0,1},{0,2},{0,3},{0,4}}, [][]int{{0,1},{1,2},{2,3}}, 1)) // [6,3,3,3,3]
    fmt.Println(maxTargetNodes1([][]int{{2,0},{3,1},{3,2},{3,4}}, [][]int{{0,3},{0,4},{2,5},{0,2},{7,0},{1,6},{1,7}}, 1)) // [3,3,4,5,3]
}