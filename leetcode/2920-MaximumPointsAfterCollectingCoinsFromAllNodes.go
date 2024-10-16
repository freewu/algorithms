package main

// 2920. Maximum Points After Collecting Coins From All Nodes
// There exists an undirected tree rooted at node 0 with n nodes labeled from 0 to n - 1. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree. 
// You are also given a 0-indexed array coins of size n where coins[i] indicates the number of coins in the vertex i, 
// and an integer k.

// Starting from the root, you have to collect all the coins 
// such that the coins at a node can only be collected if the coins of its ancestors have been already collected.

// Coins at nodei can be collected in one of the following ways:
//     1. Collect all the coins, but you will get coins[i] - k points. 
//        If coins[i] - k is negative then you will lose abs(coins[i] - k) points.
//     2. Collect all the coins, but you will get floor(coins[i] / 2) points. 
//        If this way is used, then for all the nodej present in the subtree of nodei, coins[j] will get reduced to floor(coins[j] / 2).

// Return the maximum points you can get after collecting the coins from all the tree nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/09/18/ex1-copy.png" />
// Input: edges = [[0,1],[1,2],[2,3]], coins = [10,10,3,3], k = 5
// Output: 11                        
// Explanation: 
// Collect all the coins from node 0 using the first way. Total points = 10 - 5 = 5.
// Collect all the coins from node 1 using the first way. Total points = 5 + (10 - 5) = 10.
// Collect all the coins from node 2 using the second way so coins left at node 3 will be floor(3 / 2) = 1. Total points = 10 + floor(3 / 2) = 11.
// Collect all the coins from node 3 using the second way. Total points = 11 + floor(1 / 2) = 11.
// It can be shown that the maximum points we can get after collecting coins from all the nodes is 11. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/09/18/ex2.png" />
// Input: edges = [[0,1],[0,2]], coins = [8,4,4], k = 0
// Output: 16
// Explanation: 
// Coins will be collected from all the nodes using the first way. Therefore, total points = (8 - 0) + (4 - 0) + (4 - 0) = 16.

// Constraints:
//     n == coins.length
//     2 <= n <= 10^5
//     0 <= coins[i] <= 10^4
//     edges.length == n - 1
//     0 <= edges[i][0], edges[i][1] < n
//     0 <= k <= 10^4

import "fmt"

// dfs
func maximumPoints(edges [][]int, coins []int, k int) int {
    n, inf := len(coins), 1_000_000_000
    adj, memo := make([][]int, n), make([][]int, n)
    for i := range adj {
        adj[i] = make([]int,0)
        memo[i] = make([]int,15)
        for j := range memo[i] { // init
            memo[i][j] = -inf
        }
    }
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]],e[1])
        adj[e[1]] = append(adj[e[1]],e[0])
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(n int, shift int, parent int) int
    dfs = func(i int, shift int, parent int) int { 
        if shift > 14 { shift = 14 } // shift max =14
        if memo[i][shift] > -inf { return memo[i][shift] }
        c1, c2 := (coins[i] >> shift) - k, (coins[i] >> shift) / 2
        for _, to := range adj[i] {
            if to != parent {
                c1 += dfs(to, shift, i)
                c2 += dfs(to, shift + 1, i)
            }
        }
        memo[i][shift] = max(c1,c2)
        return memo[i][shift]
    }
    return dfs(0,0,-1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/09/18/ex1-copy.png" />
    // Input: edges = [[0,1],[1,2],[2,3]], coins = [10,10,3,3], k = 5
    // Output: 11                        
    // Explanation: 
    // Collect all the coins from node 0 using the first way. Total points = 10 - 5 = 5.
    // Collect all the coins from node 1 using the first way. Total points = 5 + (10 - 5) = 10.
    // Collect all the coins from node 2 using the second way so coins left at node 3 will be floor(3 / 2) = 1. Total points = 10 + floor(3 / 2) = 11.
    // Collect all the coins from node 3 using the second way. Total points = 11 + floor(1 / 2) = 11.
    // It can be shown that the maximum points we can get after collecting coins from all the nodes is 11. 
    fmt.Println(maximumPoints([][]int{{0,1},{1,2},{2,3}}, []int{10,10,3,3}, 5)) // 11
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/09/18/ex2.png" />
    // Input: edges = [[0,1],[0,2]], coins = [8,4,4], k = 0
    // Output: 16
    // Explanation: 
    // Coins will be collected from all the nodes using the first way. Therefore, total points = (8 - 0) + (4 - 0) + (4 - 0) = 16.
    fmt.Println(maximumPoints([][]int{{0,1},{0,2}}, []int{8,4,4}, 0)) // 16
}