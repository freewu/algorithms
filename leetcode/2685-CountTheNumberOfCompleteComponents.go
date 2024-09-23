package main

// 2685. Count the Number of Complete Components
// You are given an integer n. 
// There is an undirected graph with n vertices, numbered from 0 to n - 1. 
// You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting vertices ai and bi.

// Return the number of complete connected components of the graph.

// A connected component is a subgraph of a graph in which there exists a path between any two vertices, 
// and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

// A connected component is said to be complete if there exists an edge between every pair of its vertices.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/04/11/screenshot-from-2023-04-11-23-31-23.png" />
// Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4]]
// Output: 3
// Explanation: From the picture above, one can see that all of the components of this graph are complete.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/04/11/screenshot-from-2023-04-11-23-32-00.png" />
// Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]
// Output: 1
// Explanation: The component containing vertices 0, 1, and 2 is complete since there is an edge between every pair of two vertices. On the other hand, the component containing vertices 3, 4, and 5 is not complete since there is no edge between vertices 4 and 5. Thus, the number of complete components in this graph is 1.

// Constraints:
//     1 <= n <= 50
//     0 <= edges.length <= n * (n - 1) / 2
//     edges[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi
//     There are no repeated edges.

import "fmt"

// dfs
func countCompleteComponents(n int, edges [][]int) int {
    res, visited, adj := 0, make([]bool, n), make([][]int, n)
    for i := 0; i < n; i++ {
        adj[i] = make([]int, 0)
    }
    for i := 0; i < len(edges); i++ { // 邻接表
        adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
        adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
    }
    type Pair struct {
        first  int
        second int
    }
    var dfs func (node int) Pair
    dfs = func (node int) Pair {
        visited[node] = true
        vertices, edges := 1, len(adj[node])
        for _, v := range adj[node] {
            if !visited[v] {
                temp := dfs(v)
                vertices += temp.first
                edges += temp.second
            }
        }
        return Pair{ vertices, edges }
    }
    for i := 0; i < len(adj); i++ {
        if !visited[i] {
            t := dfs(i)
            if t.first * ( t.first - 1) == t.second {
                res++
            }
        }
    }
    return res
}

func countCompleteComponents1(n int, edges [][]int) int {
    parent, ecount, ncount := make([]int, n), make([]int, n), make([]int, n)
    for i := range parent {
        parent[i], ecount[i], ncount[i] = i, 0, 1
    }
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(x, y int) {
        fx, fy := find(x), find(y)
        ecount[fy]++
        if fx != fy {
            parent[fx] = fy
            ncount[fy] += ncount[fx]
            ecount[fy] += ecount[fx]
        }
    }
    for _, e := range edges {
        union(e[0], e[1])
    }
    res := 0
    for i := range parent {
        if parent[i] == i && ncount[i] * (ncount[i]-1) / 2 == ecount[i] {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/04/11/screenshot-from-2023-04-11-23-31-23.png" />
    // Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4]]
    // Output: 3
    // Explanation: From the picture above, one can see that all of the components of this graph are complete.
    fmt.Println(countCompleteComponents(6, [][]int{{0,1},{0,2},{1,2},{3,4}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/04/11/screenshot-from-2023-04-11-23-32-00.png" />
    // Input: n = 6, edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]
    // Output: 1
    // Explanation: The component containing vertices 0, 1, and 2 is complete since there is an edge between every pair of two vertices. On the other hand, the component containing vertices 3, 4, and 5 is not complete since there is no edge between vertices 4 and 5. Thus, the number of complete components in this graph is 1.
    fmt.Println(countCompleteComponents(6, [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}})) // 1

    fmt.Println(countCompleteComponents1(6, [][]int{{0,1},{0,2},{1,2},{3,4}})) // 3
    fmt.Println(countCompleteComponents1(6, [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}})) // 1
}