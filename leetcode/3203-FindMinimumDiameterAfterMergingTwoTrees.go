package main

// 3203. Find Minimum Diameter After Merging Two Trees
// There exist two undirected trees with n and m nodes, numbered from 0 to n - 1 and from 0 to m - 1, respectively. 
// You are given two 2D integer arrays edges1 and edges2 of lengths n - 1 and m - 1, 
// respectively, where edges1[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the first tree and edges2[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the second tree.

// You must connect one node from the first tree with another node from the second tree with an edge.

// Return the minimum possible diameter of the resulting tree.

// The diameter of a tree is the length of the longest path between any two nodes in the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/04/22/example11-transformed.png" />
// Input: edges1 = [[0,1],[0,2],[0,3]], edges2 = [[0,1]]
// Output: 3
// Explanation:
// We can obtain a tree of diameter 3 by connecting node 0 from the first tree with any node from the second tree.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2024/04/22/example211.png" />
// Input: edges1 = [[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]], edges2 = [[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]]
// Output: 5
// Explanation:
// We can obtain a tree of diameter 5 by connecting node 0 from the first tree with node 0 from the second tree.

// Constraints:
//     1 <= n, m <= 10^5
//     edges1.length == n - 1
//     edges2.length == m - 1
//     edges1[i].length == edges2[i].length == 2
//     edges1[i] = [ai, bi]
//     0 <= ai, bi < n
//     edges2[i] = [ui, vi]
//     0 <= ui, vi < m
//     The input is generated such that edges1 and edges2 represent valid trees.

import "fmt"

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
    d1, d2 := treeDiameter(edges1), treeDiameter(edges2)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(max(d1, d2), (d1 + 1) / 2 + (d2 + 1)/ 2 + 1)
}

func treeDiameter(edges [][]int) int {
    res, n := 0, len(edges) + 1
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x, parent int) int
    dfs = func(x, parent int) int {
        mx := 0
        for _, y := range graph[x] {
            if y != parent {
                next := dfs(y, x) + 1
                res = max(res, mx + next)
                mx = max(mx, next)
            }
        }
        return mx
    }
    dfs(0, -1)
    return res
    // res, n, a := 0, len(edges) + 1, 0
    // graph := make([][]int, n)
    // for _, v := range edges {
    //     graph[v[0]] = append(graph[v[0]], v[1])
    //     graph[v[1]] = append(graph[v[1]], v[0])
    // }
    // var dfs func(i, parent, t int)
    // dfs = func(i, parent, t int) {
    //     for _, j := range graph[i] {
    //         if j != parent { 
    //             dfs(j, i, t + 1) 
    //         }
    //     }
    //     if res < t {
    //         res, a = t, i
    //     }
    // }
    // dfs(0, -1, 0)
    // dfs(a, -1, 0)
    // return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2024/04/22/example11-transformed.png" />
    // Input: edges1 = [[0,1],[0,2],[0,3]], edges2 = [[0,1]]
    // Output: 3
    // Explanation:
    // We can obtain a tree of diameter 3 by connecting node 0 from the first tree with any node from the second tree.
    fmt.Println(minimumDiameterAfterMerge([][]int{{0,1},{0,2},{0,3}}, [][]int{{0,1}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2024/04/22/example211.png" />
    // Input: edges1 = [[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]], edges2 = [[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]]
    // Output: 5
    // Explanation:
    // We can obtain a tree of diameter 5 by connecting node 0 from the first tree with node 0 from the second tree.
    fmt.Println(minimumDiameterAfterMerge([][]int{{0,1},{0,2},{0,3},{2,4},{2,5},{3,6},{2,7}}, [][]int{{0,1},{0,2},{0,3},{2,4},{2,5},{3,6},{2,7}})) // 5
}