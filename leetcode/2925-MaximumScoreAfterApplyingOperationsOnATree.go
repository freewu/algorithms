package main

// 2925. Maximum Score After Applying Operations on a Tree
// There is an undirected tree with n nodes labeled from 0 to n - 1, and rooted at node 0. 
// You are given a 2D integer array edges of length n - 1, where edges[i] = [ai, bi] indicates 
// that there is an edge between nodes ai and bi in the tree.

// You are also given a 0-indexed integer array values of length n, 
// where values[i] is the value associated with the ith node.

// You start with a score of 0. In one operation, you can:
//     Pick any node i.
//     Add values[i] to your score.
//     Set values[i] to 0.

// A tree is healthy if the sum of values on the path from the root to any leaf node is different than zero.

// Return the maximum score you can obtain after performing these operations on the tree any number of times so that it remains healthy.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/10/11/graph-13-1.png" />
// Input: edges = [[0,1],[0,2],[0,3],[2,4],[4,5]], values = [5,2,5,2,1,1]
// Output: 11
// Explanation: We can choose nodes 1, 2, 3, 4, and 5. The value of the root is non-zero. Hence, the sum of values on the path from the root to any leaf is different than zero. Therefore, the tree is healthy and the score is values[1] + values[2] + values[3] + values[4] + values[5] = 11.
// It can be shown that 11 is the maximum score obtainable after any number of operations on the tree.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/10/11/graph-14-2.png" />
// Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], values = [20,10,9,7,4,3,5]
// Output: 40
// Explanation: We can choose nodes 0, 2, 3, and 4.
// - The sum of values on the path from 0 to 4 is equal to 10.
// - The sum of values on the path from 0 to 3 is equal to 10.
// - The sum of values on the path from 0 to 5 is equal to 3.
// - The sum of values on the path from 0 to 6 is equal to 5.
// Therefore, the tree is healthy and the score is values[0] + values[2] + values[3] + values[4] = 40.
// It can be shown that 40 is the maximum score obtainable after any number of operations on the tree.

// Constraints:
//     2 <= n <= 2 * 10^4
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     values.length == n
//     1 <= values[i] <= 10^9
//     The input is generated such that edges represents a valid tree.

import "fmt"

func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
    graph := make(map[int][]int)
    for _, v := range edges {
        graph[v[0]], graph[v[1]] = append(graph[v[0]], v[1]), append(graph[v[1]], v[0])
    }
    var dfs func(index int, values []int, parent int) (int64, int64)
    dfs = func(index int, values []int, parent int) (int64, int64) {
        score, cost := int64(0), int64(0)
        for i := range graph[index] {
            if graph[index][i] == parent { continue }
            pathScore, pathCost := dfs(graph[index][i], values, index)
            // sum all score and cost for each path
            score += pathScore
            cost += pathCost
        }
        // if cost is 0 that mean the node is leaf return score = 0 and the values as the cost
        if cost == 0 { return 0, int64(values[index]) }
        if int64(values[index]) < cost { // check if the current root is less than minCost 
            // if current value less than minCost then return root as minCost
            // and return maxScore = maxScore + minCost
            return score + cost, int64(values[index])
        }
        return score + int64(values[index]), cost // else add the current node value to score
    }
    res, _ := dfs(0, values, -1)
    return res
}

func maximumScoreAfterOperations1(edges [][]int, values []int) int64 {
    sum := 0
    for i := 0; i < len(values); i++{
        sum += values[i]
    }
    graph := make([][]int, len(values))
    for i := 0; i < len(values); i++{
        graph[i] = make([]int, 0)
    }
    for i := 0; i < len(edges); i++{
        graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
        graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var solve func(g [][]int, node int, values []int, fa int) int 
    solve = func(g [][]int, node int, values []int, fa int) int {
        if len(g[node]) == 1 && g[node][0] == fa {  return values[node] }
        sum := 0
        for i := 0; i < len(g[node]); i++{
            if g[node][i] == fa { continue }
            sum += solve(g, g[node][i], values, node)
        }
        return min(values[node], sum)
    }
    return int64(sum) - int64(solve(graph, 0, values, -1))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/10/11/graph-13-1.png" />
    // Input: edges = [[0,1],[0,2],[0,3],[2,4],[4,5]], values = [5,2,5,2,1,1]
    // Output: 11
    // Explanation: We can choose nodes 1, 2, 3, 4, and 5. The value of the root is non-zero. Hence, the sum of values on the path from the root to any leaf is different than zero. Therefore, the tree is healthy and the score is values[1] + values[2] + values[3] + values[4] + values[5] = 11.
    // It can be shown that 11 is the maximum score obtainable after any number of operations on the tree.
    fmt.Println(maximumScoreAfterOperations([][]int{{0,1},{0,2},{0,3},{2,4},{4,5}}, []int{5,2,5,2,1,1})) // 11
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/10/11/graph-14-2.png" />
    // Input: edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], values = [20,10,9,7,4,3,5]
    // Output: 40
    // Explanation: We can choose nodes 0, 2, 3, and 4.
    // - The sum of values on the path from 0 to 4 is equal to 10.
    // - The sum of values on the path from 0 to 3 is equal to 10.
    // - The sum of values on the path from 0 to 5 is equal to 3.
    // - The sum of values on the path from 0 to 6 is equal to 5.
    // Therefore, the tree is healthy and the score is values[0] + values[2] + values[3] + values[4] = 40.
    // It can be shown that 40 is the maximum score obtainable after any number of operations on the tree.
    fmt.Println(maximumScoreAfterOperations([][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}, []int{20,10,9,7,4,3,5})) // 40

    fmt.Println(maximumScoreAfterOperations1([][]int{{0,1},{0,2},{0,3},{2,4},{4,5}}, []int{5,2,5,2,1,1})) // 11
    fmt.Println(maximumScoreAfterOperations1([][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}, []int{20,10,9,7,4,3,5})) // 40
}