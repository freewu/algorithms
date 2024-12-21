package main

// 2872. Maximum Number of K-Divisible Components
// There is an undirected tree with n nodes labeled from 0 to n - 1. 
// You are given the integer n and a 2D integer array edges of length n - 1, 
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// You are also given a 0-indexed integer array values of length n, 
// where values[i] is the value associated with the ith node, and an integer k.

// A valid split of the tree is obtained by removing any set of edges, possibly empty, 
// from the tree such that the resulting components all have values that are divisible by k, 
// where the value of a connected component is the sum of the values of its nodes.

// Return the maximum number of components in any valid split.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/07/example12-cropped2svg.jpg">
// Input: n = 5, edges = [[0,2],[1,2],[1,3],[2,4]], values = [1,8,1,4,4], k = 6
// Output: 2
// Explanation: We remove the edge connecting node 1 with 2. The resulting split is valid because:
// - The value of the component containing nodes 1 and 3 is values[1] + values[3] = 12.
// - The value of the component containing nodes 0, 2, and 4 is values[0] + values[2] + values[4] = 6.
// It can be shown that no other valid split has more than 2 connected components.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/07/example21svg-1.jpg">
// Input: n = 7, edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], values = [3,0,6,1,5,2,1], k = 3
// Output: 3
// Explanation: We remove the edge connecting node 0 with 2, and the edge connecting node 0 with 1. The resulting split is valid because:
// - The value of the component containing node 0 is values[0] = 3.
// - The value of the component containing nodes 2, 5, and 6 is values[2] + values[5] + values[6] = 9.
// - The value of the component containing nodes 1, 3, and 4 is values[1] + values[3] + values[4] = 6.
// It can be shown that no other valid split has more than 3 connected components.

// Constraints:
//     1 <= n <= 3 * 10^4
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     values.length == n
//     0 <= values[i] <= 10^9
//     1 <= k <= 10^9
//     Sum of values is divisible by k.
//     The input is generated such that edges represents a valid tree.

import "fmt"

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
    res, graph := 0, make([][]int, n)
    for _, v := range edges {
        graph[v[0]], graph[v[1]] = append(graph[v[0]], v[1]), append(graph[v[1]], v[0])
    }
    var dfs func(index, parent int) int
    dfs = func(index, parent int) int {
        sum := values[index]
        for _, i := range graph[index] {
            if i != parent {
                sum += dfs(i, index)
            }
        }
        if sum % k == 0 {
            res++
        }
        return sum
    }
    dfs(0, -1)
    return res
}

func maxKDivisibleComponents1(n int, edges [][]int, values []int, k int) int {
    // Step 1: Create adjacency list from edges
    adjList := make([][]int, n)
    for _, v := range edges {
        adjList[v[0]] = append(adjList[v[0]], v[1])
        adjList[v[1]] = append(adjList[v[1]], v[0])
    }
    // Step 2: Initialize component count
    res := 0 // Use array to pass by reference
    // Step 3: Start DFS traversal from node 0
    var dfs func(currentNode, parentNode int, adjList [][]int) int
    dfs = func(currentNode, parentNode int, adjList [][]int) int {
        // Step 1: Initialize sum for the current subtree
        sum := 0
        // Step 2: Traverse all neighbors
        for _, neighborNode := range adjList[currentNode] {
            if neighborNode != parentNode {
                // Recursive call to process the subtree rooted at the neighbor
                sum += dfs(neighborNode, currentNode, adjList)
                sum %= k // Ensure the sum stays within bounds
            }
        }
        // Step 3: Add the value of the current node to the sum
        sum += values[currentNode]
        sum %= k
        // Step 4: Check if the sum is divisible by k
        if sum == 0 {
            res++
        }
        // Step 5: Return the computed sum for the current subtree
        return sum
    }
    dfs(0, -1, adjList)
    // Step 4: Return the total number of components
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/08/07/example12-cropped2svg.jpg">
    // Input: n = 5, edges = [[0,2],[1,2],[1,3],[2,4]], values = [1,8,1,4,4], k = 6
    // Output: 2
    // Explanation: We remove the edge connecting node 1 with 2. The resulting split is valid because:
    // - The value of the component containing nodes 1 and 3 is values[1] + values[3] = 12.
    // - The value of the component containing nodes 0, 2, and 4 is values[0] + values[2] + values[4] = 6.
    // It can be shown that no other valid split has more than 2 connected components.
    fmt.Println(maxKDivisibleComponents(5, [][]int{{0,2},{1,2},{1,3},{2,4}}, []int{1,8,1,4,4}, 6)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/08/07/example21svg-1.jpg">
    // Input: n = 7, edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], values = [3,0,6,1,5,2,1], k = 3
    // Output: 3
    // Explanation: We remove the edge connecting node 0 with 2, and the edge connecting node 0 with 1. The resulting split is valid because:
    // - The value of the component containing node 0 is values[0] = 3.
    // - The value of the component containing nodes 2, 5, and 6 is values[2] + values[5] + values[6] = 9.
    // - The value of the component containing nodes 1, 3, and 4 is values[1] + values[3] + values[4] = 6.
    // It can be shown that no other valid split has more than 3 connected components.
    fmt.Println(maxKDivisibleComponents(7, [][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}, []int{3,0,6,1,5,2,1}, 3)) // 3

    fmt.Println(maxKDivisibleComponents1(5, [][]int{{0,2},{1,2},{1,3},{2,4}}, []int{1,8,1,4,4}, 6)) // 2
    fmt.Println(maxKDivisibleComponents1(7, [][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}, []int{3,0,6,1,5,2,1}, 3)) // 3
}