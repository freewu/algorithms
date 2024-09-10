package main

// 2973. Find Number of Coins to Place in Tree Nodes
// You are given an undirected tree with n nodes labeled from 0 to n - 1, and rooted at node 0. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// You are also given a 0-indexed integer array cost of length n, 
// where cost[i] is the cost assigned to the ith node.

// You need to place some coins on every node of the tree. T
// he number of coins to be placed at node i can be calculated as:
//     1. If size of the subtree of node i is less than 3, place 1 coin.
//     2. Otherwise, place an amount of coins equal to the maximum product of cost values assigned to 3 distinct nodes in the subtree of node i. If this product is negative, place 0 coins.

// Return an array coin of size n such that coin[i] is the number of coins placed at node i.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012641.png" />
// Input: edges = [[0,1],[0,2],[0,3],[0,4],[0,5]], cost = [1,2,3,4,5,6]
// Output: [120,1,1,1,1,1]
// Explanation: For node 0 place 6 * 5 * 4 = 120 coins. All other nodes are leaves with subtree of size 1, place 1 coin on each of them.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012614.png" />
// Input: edges = [[0,1],[0,2],[1,3],[1,4],[1,5],[2,6],[2,7],[2,8]], cost = [1,4,2,3,5,7,8,-4,2]
// Output: [280,140,32,1,1,1,1,1,1]
// Explanation: The coins placed on each node are:
// - Place 8 * 7 * 5 = 280 coins on node 0.
// - Place 7 * 5 * 4 = 140 coins on node 1.
// - Place 8 * 2 * 2 = 32 coins on node 2.
// - All other nodes are leaves with subtree of size 1, place 1 coin on each of them.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012513.png" />
// Input: edges = [[0,1],[0,2]], cost = [1,2,-2]
// Output: [0,1,1]
// Explanation: Node 1 and 2 are leaves with subtree of size 1, place 1 coin on each of them. For node 0 the only possible product of cost is 2 * 1 * -2 = -4. Hence place 0 coins on node 0.

// Constraints:
//     2 <= n <= 2 * 10^4
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     cost.length == n
//     1 <= |cost[i]| <= 10^4
//     The input is generated such that edges represents a valid tree.

import "fmt"
import "slices"

func placedCoins(edges [][]int, cost []int) []int64 {
    n := len(cost)
    res, adj := make([]int64, n), make([][]int, n)
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    var dfs func(adj [][]int, cost []int, root, par int) []int64
    dfs = func(adj [][]int, cost []int, root, par int) []int64 {
        rootCosts := []int64{int64(cost[root])}
        // Traverse all the children of root
        // ans accumulate the costs
        for _, c := range adj[root] {
            if c == par { continue }
            currCost := dfs(adj, cost, c, root)
            for _, curC := range currCost {
                rootCosts = append(rootCosts, curC)
            }
        }
        // Now sort the root costs as we need top 3 and bottom 2
        slices.Sort(rootCosts)
        slices.Reverse(rootCosts)   // We want sorted to be decresing
        n := len(rootCosts)
        // If we don't have 3 costs then this must be a leaf node
        // just add 1 for thier result
        if n < 3 {
            res[root] = 1
            return rootCosts
        }
        // We can take 0th, 1st and 2nd or we can take 0th, n-1, n-2
        // whichever gives the greater result
        if rootCosts[1] * rootCosts[2] > rootCosts[n-1] * rootCosts[n-2] {
            res[root] = rootCosts[0] * rootCosts[1] * rootCosts[2]
        } else {
            res[root] = rootCosts[0] * rootCosts[n-1] * rootCosts[n-2]
        }
        // If after multiplying this gives negative number, then we have to set it 0
        res[root] = max(res[root], 0)
        // If we have 5 or more root costs, return top 3 and bottom 2
        if n <= 5 {
            return rootCosts
        }
        return []int64{ rootCosts[0], rootCosts[1], rootCosts[2], rootCosts[n-2], rootCosts[n-1] }
    }
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    dfs(adj, cost, 0, -1)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012641.png" />
    // Input: edges = [[0,1],[0,2],[0,3],[0,4],[0,5]], cost = [1,2,3,4,5,6]
    // Output: [120,1,1,1,1,1]
    // Explanation: For node 0 place 6 * 5 * 4 = 120 coins. All other nodes are leaves with subtree of size 1, place 1 coin on each of them.
    fmt.Println(placedCoins([][]int{{0,1},{0,2},{0,3},{0,4},{0,5}},[]int{1,2,3,4,5,6})) // [120,1,1,1,1,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012614.png" />
    // Input: edges = [[0,1],[0,2],[1,3],[1,4],[1,5],[2,6],[2,7],[2,8]], cost = [1,4,2,3,5,7,8,-4,2]
    // Output: [280,140,32,1,1,1,1,1,1]
    // Explanation: The coins placed on each node are:
    // - Place 8 * 7 * 5 = 280 coins on node 0.
    // - Place 7 * 5 * 4 = 140 coins on node 1.
    // - Place 8 * 2 * 2 = 32 coins on node 2.
    // - All other nodes are leaves with subtree of size 1, place 1 coin on each of them.
    fmt.Println(placedCoins([][]int{{0,1},{0,2},{1,3},{1,4},{1,5},{2,6},{2,7},{2,8}},[]int{1,4,2,3,5,7,8,-4,2})) // [280,140,32,1,1,1,1,1,1]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012513.png" />
    // Input: edges = [[0,1],[0,2]], cost = [1,2,-2]
    // Output: [0,1,1]
    // Explanation: Node 1 and 2 are leaves with subtree of size 1, place 1 coin on each of them. For node 0 the only possible product of cost is 2 * 1 * -2 = -4. Hence place 0 coins on node 0.
    fmt.Println(placedCoins([][]int{{0,1},{0,2}},[]int{1,2,-2})) // [0,1,1]
}