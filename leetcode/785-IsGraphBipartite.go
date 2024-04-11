package main

// 785. Is Graph Bipartite?
// There is an undirected graph with n nodes, where each node is numbered between 0 and n - 1. 
// You are given a 2D array graph, where graph[u] is an array of nodes that node u is adjacent to. 
// More formally, for each v in graph[u], there is an undirected edge between node u and node v. 
// The graph has the following properties:
//     There are no self-edges (graph[u] does not contain u).
//     There are no parallel edges (graph[u] does not contain duplicate values).
//     If v is in graph[u], then u is in graph[v] (the graph is undirected).
//     The graph may not be connected, meaning there may be two nodes u and v such that there is no path between them.

// A graph is bipartite if the nodes can be partitioned into two independent sets A and B such that every edge in the graph connects a node in set A and a node in set B.
// Return true if and only if it is bipartite.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/21/bi2.jpg" />
// Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
// Output: false
// Explanation: There is no way to partition the nodes into two independent sets such that every edge connects a node in one and a node in the other.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/21/bi1.jpg" />
// Input: graph = [[1,3],[0,2],[1,3],[0,2]]
// Output: true
// Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.

// Constraints:
//     graph.length == n
//     1 <= n <= 100
//     0 <= graph[u].length < n
//     0 <= graph[u][i] <= n - 1
//     graph[u] does not contain u.
//     All the values of graph[u] are unique.
//     If graph[u] contains v, then graph[v] contains u.

import "fmt"

func isBipartite(graph [][]int) bool {
    dp := make([]int, len(graph))
    for i := range dp {
        dp[i] = -1
    }
    var dfs func(node int, cur int, dp []int, graph [][]int) bool 
    dfs = func(node int, cur int, dp []int, graph [][]int) bool {
        dp[node] = cur
        for _, neighbor := range graph[node] {
            // 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
            if dp[neighbor] == -1 {
                if dfs(neighbor, 1-cur, dp, graph) == false {
                    return false
                }
            } else if dp[neighbor] == cur { // 不存在自环（graph[u] 不包含 u）
                return false
            }
        }
        return true
    }
    for i := range dp { 
        if dp[i] == -1 {
            if dfs(i, 0, dp, graph) == false {
                return false
            }
        }
    }
    return true
}

func isBipartite1(graph [][]int) bool {
    n, res := len(graph), true
    color, visited := make([]bool, n),make([]bool, n) // color 记录图中节点的颜色，两种; visited 记录是否已经被访问过了
    var traverse func(x int)
    traverse = func(x int) {
        if !res {
            return
        }
        visited[x] = true
        for _, y := range graph[x] {
            if !visited[y] {// 相邻的节点y没有被访问过，则给y标记成和x不同的颜色
                color[y] = !color[x]
                traverse(y)
            } else {// 如果y已经被访问过了，并且x和y的颜色一样，那就不是二分图
                if color[y] == color[x] {
                    res = false
                    break
                }
            }
        }
    }
    for x := 0; x < n; x++ {
        if !visited[x] {
            traverse(x)
        }
    }
    return res
}

func main() {
    // Explanation: There is no way to partition the nodes into two independent sets such that every edge connects a node in one and a node in the other.
    fmt.Println(isBipartite([][]int{{1,2,3},{0,2},{0,1,3},{0,2}})) // false
    // Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.
    fmt.Println(isBipartite([][]int{{1,3},{0,2},{1,3},{0,2}})) // true

    fmt.Println(isBipartite1([][]int{{1,2,3},{0,2},{0,1,3},{0,2}})) // false
    fmt.Println(isBipartite1([][]int{{1,3},{0,2},{1,3},{0,2}})) // true
}