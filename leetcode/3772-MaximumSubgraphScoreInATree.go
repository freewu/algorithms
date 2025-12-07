package main

// 3772. Maximum Subgraph Score in a Tree
// You are given an undirected tree with n nodes, numbered from 0 to n - 1. 
// It is represented by a 2D integer array edges​​​​​​​ of length n - 1, where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// You are also given an integer array good of length n, where good[i] is 1 if the ith node is good, and 0 if it is bad.

// Define the score of a subgraph as the number of good nodes minus the number of bad nodes in that subgraph.

// For each node i, find the maximum possible score among all connected subgraphs that contain node i.

// Return an array of n integers where the ith element is the maximum score for node i.

// A subgraph is a graph whose vertices and edges are subsets of the original graph.

// A connected subgraph is a subgraph in which every pair of its vertices is reachable from one another using only its edges.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2025/11/17/tree1fixed.png" />
// Input: n = 3, edges = [[0,1],[1,2]], good = [1,0,1]
// Output: [1,1,1]
// Explanation:
// Green nodes are good and red nodes are bad.
// For each node, the best connected subgraph containing it is the whole tree, which has 2 good nodes and 1 bad node, resulting in a score of 1.
// Other connected subgraphs containing a node may have the same score.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2025/11/17/tree2.png" />
// Input: n = 5, edges = [[1,0],[1,2],[1,3],[3,4]], good = [0,1,0,1,1]
// Output: [2,3,2,3,3]
// Explanation:
// Node 0: The best connected subgraph consists of nodes 0, 1, 3, 4, which has 3 good nodes and 1 bad node, resulting in a score of 3 - 1 = 2.
// Nodes 1, 3, and 4: The best connected subgraph consists of nodes 1, 3, 4, which has 3 good nodes, resulting in a score of 3.
// Node 2: The best connected subgraph consists of nodes 1, 2, 3, 4, which has 3 good nodes and 1 bad node, resulting in a score of 3 - 1 = 2.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2025/11/17/tree3.png" />
// Input: n = 2, edges = [[0,1]], good = [0,0]
// Output: [-1,-1]
// Explanation:
// For each node, including the other node only adds another bad node, so the best score for both nodes is -1.

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ai, bi]
//     0 <= ai, bi < n
//     good.length == n
//     0 <= good[i] <= 1
//     The input is generated such that edges represents a valid tree.

import "fmt"

func maxSubgraphScore(n int, edges [][]int, good []int) []int {
    graph := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
    }
    // subScore[x] 表示（以 0 为根时）子树 x 的最大得分（一定包含节点 x）
    subScore := make([]int, n)
    // 计算并返回 subScore[x]
    var dfs func(int, int) int
    dfs = func(x, fa int) int {
        for _, y := range graph[x] {
            if y != fa {
                // 如果子树 y 的得分是负数，不选子树 y，否则选子树 y
                subScore[x] += max(dfs(y, x), 0)
            }
        }
        subScore[x] += good[x] * 2 - 1 // subScore[x] 一定包含 x
        return subScore[x]
    }
    dfs(0, -1)
    res := make([]int, n)
    // 计算子图 x 的最大得分 scoreX，其中 faScore 表示来自父节点 fa 的最大得分（一定包含节点 fa）
    var reroot func(int, int, int)
    reroot = func(x, fa, faScore int) {
        scoreX := subScore[x] + max(faScore, 0)
        res[x] = scoreX
        for _, y := range graph[x] {
            if y != fa {
                // scoreX-max(subScore[y],0) 是不含子树 y 的最大得分
                reroot(y, x, scoreX - max(subScore[y], 0))
            }
        }
    }
    reroot(0, -1, 0)
    return res
}

func maxSubgraphScore1(n int, edges [][]int, good []int) []int {
    if n == 0 { return []int{} }
    values := make([]int, n)
    for i := 0; i < n; i++ { // 转换值
        if good[i] == 0 {
            values[i] = -1
        } else {
            values[i] = 1
        }
    }
    graph := make([][]int, n)
    for _, edge := range edges { // 构建树
        u, v := edge[0], edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    res, dp := make([]int, n), make([]int, n) // // dp[u] 表示以u为根的子树的最大得分
    var dfs1 func(u, parent int) 
    dfs1 = func(u, parent int) {  // 第一次dfs：计算以0为根的dp值
        dp[u] = values[u]
        for _, v := range graph[u] {
            if v != parent {
                dfs1(v, u)
                if dp[v] > 0 {
                    dp[u] += dp[v]
                }
            }
        }
    }
    dfs1(0, -1)
    var dfs2 func(u, parent, parentScore int)
    dfs2 = func(u, parent, parentScore int) { // 第二次dfs：换根计算每个节点作为根的答案
        // 计算以u为根的总得分
        total := values[u]
        if parentScore > 0 {
            total += parentScore
        }
        for _, v := range graph[u] {
            if v != parent && dp[v] > 0 {
                total += dp[v]
            }
        }
        res[u] = total
        // 为每个子节点计算新的parentScore
        for _, v := range graph[u] {
            if v != parent {
                newParentScore := total
                if dp[v] > 0 {
                    newParentScore -= dp[v]
                }
                // 如果去掉v后得分变负，则不从父节点传递任何正得分
                if newParentScore < 0 {
                    newParentScore = 0
                }
                dfs2(v, u, newParentScore)
            }
        }
    }
    dfs2(0, -1, 0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2025/11/17/tree1fixed.png" />
    // Input: n = 3, edges = [[0,1],[1,2]], good = [1,0,1]
    // Output: [1,1,1]
    // Explanation:
    // Green nodes are good and red nodes are bad.
    // For each node, the best connected subgraph containing it is the whole tree, which has 2 good nodes and 1 bad node, resulting in a score of 1.
    // Other connected subgraphs containing a node may have the same score.
    fmt.Println(maxSubgraphScore(3, [][]int{{0,1},{1,2}}, []int{1,0,1})) // [1,1,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2025/11/17/tree2.png" />
    // Input: n = 5, edges = [[1,0],[1,2],[1,3],[3,4]], good = [0,1,0,1,1]
    // Output: [2,3,2,3,3]
    // Explanation:
    // Node 0: The best connected subgraph consists of nodes 0, 1, 3, 4, which has 3 good nodes and 1 bad node, resulting in a score of 3 - 1 = 2.
    // Nodes 1, 3, and 4: The best connected subgraph consists of nodes 1, 3, 4, which has 3 good nodes, resulting in a score of 3.
    // Node 2: The best connected subgraph consists of nodes 1, 2, 3, 4, which has 3 good nodes and 1 bad node, resulting in a score of 3 - 1 = 2.
    fmt.Println(maxSubgraphScore(5, [][]int{{1,0},{1,2},{1,3},{3,4}}, []int{0,1,0,1,1})) // [2,3,2,3,3]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2025/11/17/tree3.png" />
    // Input: n = 2, edges = [[0,1]], good = [0,0]
    // Output: [-1,-1]
    // Explanation:
    // For each node, including the other node only adds another bad node, so the best score for both nodes is -1.
    fmt.Println(maxSubgraphScore(2, [][]int{{0,1}}, []int{0,0})) // [-1,-1]

    fmt.Println(maxSubgraphScore1(3, [][]int{{0,1},{1,2}}, []int{1,0,1})) // [1,1,1]
    fmt.Println(maxSubgraphScore1(5, [][]int{{1,0},{1,2},{1,3},{3,4}}, []int{0,1,0,1,1})) // [2,3,2,3,3]
    fmt.Println(maxSubgraphScore1(2, [][]int{{0,1}}, []int{0,0})) // [-1,-1]
}