package main

// 1857. Largest Color Value in a Directed Graph
// There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.

// You are given a string colors where colors[i] is a lowercase English letter representing the color of the ith node in this graph (0-indexed). 
// You are also given a 2D array edges where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.

// A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that there is a directed edge 
// from xi to xi+1 for every 1 <= i < k. The color value of the path is the number of nodes that are colored the most frequently occurring color along that path.

// Return the largest color value of any valid path in the given graph, or -1 if the graph contains a cycle.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/21/leet1.png" />
// Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
// Output: 3
// Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/21/leet2.png" />
// Input: colors = "a", edges = [[0,0]]
// Output: -1
// Explanation: There is a cycle from 0 to 0.

// Constraints:
//     n == colors.length
//     m == edges.length
//     1 <= n <= 10^5
//     0 <= m <= 10^5
//     colors consists of lowercase English letters.
//     0 <= aj, bj < n

import "fmt"

// Topological Sort
func largestPathValue(colors string, edges [][]int) int {
    n := len(colors)
    adj, indeg := make([][]int, n), make([]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        indeg[v]++
    }
    stack := []int{}
    for u, d := range indeg {
        if d == 0 {
            stack = append(stack, u)
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, dp := 0, make([][26]int, n )
    for len(stack) > 0 {
        u := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        dp[u][colors[u]-'a'] += 1
        res = max(res, dp[u][colors[u]-'a'])
        for _, v := range adj[u] {
            for c := range dp[u] {
                dp[v][c] = max(dp[v][c], dp[u][c])
            }
            if indeg[v] -= 1; indeg[v] == 0 {
                stack = append(stack, v)
            }
        }
    }
    for _, d := range indeg {
        if d != 0 {
            return -1
        }
    }
    return res
}

func largestPathValue1(colors string, edges [][]int) int {
    res, n := 0, len(colors)
    g, deg := make([][]int, n), make([]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        if x == y { return -1 }// 自环
        g[x] = append(g[x], y)
        deg[y]++
    }
    q := make([]int, 0, n)
    for i, d := range deg {
        if d == 0 {
            q = append(q, i) // 入度为 0 的点入队
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    f := make([][26]int, n)
    for len(q) > 0 {
        x := q[0] // x 的所有转移来源都计算完毕，也都更新到 f[x] 中
        q = q[1:]
        ch := colors[x] - 'a'
        f[x][ch]++
        res = max(res, f[x][ch])
        for _, y := range g[x] {
            for i, cnt := range f[x] {
                f[y][i] = max(f[y][i], cnt) // 刷表法，更新邻居的最大值
            }
            deg[y]--
            if deg[y] == 0 {
                q = append(q, y)
            }
        }
    }
    if cap(q) > 0 { // 有节点没入队，说明有环
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/21/leet1.png" />
    // Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
    // Output: 3
    // Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).
    fmt.Println(largestPathValue("abaca",[][]int{{0,1},{0,2},{2,3},{3,4}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/21/leet2.png" />
    // Input: colors = "a", edges = [[0,0]]
    // Output: -1
    // Explanation: There is a cycle from 0 to 0.
    fmt.Println(largestPathValue("a",[][]int{{0,0}})) // -1

    fmt.Println(largestPathValue("abaca",[][]int{{0,1},{0,2},{2,3},{3,4}})) // 3
    fmt.Println(largestPathValue("a",[][]int{{0,0}})) // -1
}