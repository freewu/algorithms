package main

// 3249. Count the Number of Good Nodes
// There is an undirected tree with n nodes labeled from 0 to n - 1, and rooted at node 0. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// A node is good if all the subtrees rooted at its children have the same size.

// Return the number of good nodes in the given tree.

// A subtree of treeName is a tree consisting of a node in treeName and all of its descendants.

// Example 1:
// Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
// Output: 7
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/26/tree1.png" />
// All of the nodes of the given tree are good.

// Example 2:
// Input: edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]
// Output: 6
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-193552.png" />
// There are 6 good nodes in the given tree. They are colored in the image above.

// Example 3:
// Input: edges = [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]
// Output: 12
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/08/rob.jpg" />
// All nodes except node 9 are good.

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     The input is generated such that edges represents a valid tree.

import "fmt"

func countGoodNodes(edges [][]int) int {
    res, adj := 0, make(map [int][]int)
    for _, v := range edges { // 接邻表
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    var dfs func(n, from int) int
    dfs = func(n, from int) int {
        sum, subTreeSize, allSameSize := 1, -1, true
        for _, m := range adj[n] {
            if m == from { continue }
            v := dfs(m, n)
            sum += v
            if subTreeSize == -1 {
                subTreeSize = v
            } else if v != subTreeSize {
                allSameSize = false
            }
        }
        if allSameSize {
            res++
        }
        return sum
    }
    dfs(0, -1)
    return res
}

func countGoodNodes1(edges [][]int) int {
    res, n := 0, len(edges) + 1
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    var dfs func(int, int) int
    dfs = func(x, fa int) int {
        size, sz0, ok := 1, 0, true
        for _, y := range graph[x] {
            if y == fa { continue }
            sz := dfs(y, x)
            if sz0 == 0 {
                sz0 = sz // 记录第一个儿子子树的大小
            } else if sz != sz0 { // 存在大小不一样的儿子子树
                ok = false // 注意不能 break，其他子树 y 仍然要递归
            }
            size += sz
        }
        if ok {
            res++
        }
        return size
    }
    dfs(0, -1)
    return res
}

func main() {
    // Example 1:
    // Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
    // Output: 7
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/26/tree1.png" />
    // All of the nodes of the given tree are good.
    fmt.Println(countGoodNodes([][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}})) // 7
    // Example 2:
    // Input: edges = [[0,1],[1,2],[2,3],[3,4],[0,5],[1,6],[2,7],[3,8]]
    // Output: 6
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-193552.png" />
    // There are 6 good nodes in the given tree. They are colored in the image above.
    fmt.Println(countGoodNodes([][]int{{0,1},{1,2},{2,3},{3,4},{0,5},{1,6},{2,7},{3,8}})) // 6
    // Example 3:
    // Input: edges = [[0,1],[1,2],[1,3],[1,4],[0,5],[5,6],[6,7],[7,8],[0,9],[9,10],[9,12],[10,11]]
    // Output: 12
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/08/rob.jpg" />
    // All nodes except node 9 are good.
    fmt.Println(countGoodNodes([][]int{{0,1},{1,2},{1,3},{1,4},{0,5},{5,6},{6,7},{7,8},{0,9},{9,10},{9,12},{10,11}})) // 12

    fmt.Println(countGoodNodes1([][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}})) // 7
    fmt.Println(countGoodNodes1([][]int{{0,1},{1,2},{2,3},{3,4},{0,5},{1,6},{2,7},{3,8}})) // 6
    fmt.Println(countGoodNodes1([][]int{{0,1},{1,2},{1,3},{1,4},{0,5},{5,6},{6,7},{7,8},{0,9},{9,10},{9,12},{10,11}})) // 12
}