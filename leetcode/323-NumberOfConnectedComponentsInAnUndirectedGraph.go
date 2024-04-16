package main

// 323. Number of Connected Components in an Undirected Graph
// You have a graph of n nodes. 
// You are given an integer n and an array edges where edges[i] = [ai, bi] indicates that there is an edge between ai and bi in the graph.

// Return the number of connected components in the graph.

// Example 1:
// (0) --- (1)    (3)
//          |      |
//          |      |
//         (2)    (4)
// <img src="https://assets.leetcode.com/uploads/2021/03/14/conn1-graph.jpg" />
// Input: n = 5, edges = [[0,1],[1,2],[3,4]]
// Output: 2

// Example 2:
// (0) --- (1)    (3)
//          |    / |
//          |   /  |
//         (2)/   (4)
// <img src="https://assets.leetcode.com/uploads/2021/03/14/conn2-graph.jpg" />
// Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
// Output: 1
 
// Constraints:
//     1 <= n <= 2000
//     1 <= edges.length <= 5000
//     edges[i].length == 2
//     0 <= ai <= bi < n
//     ai != bi
//     There are no repeated edges.

import "fmt"

// dfs
func countComponents(n int, edges [][]int) int {
    // 遍历所有边，记录已经链接的所有节点
    visit, res := make([]bool, n), 0
    var dfs func(edges [][]int, visit []bool, s int)
    dfs = func(edges [][]int, visit []bool, s int) {
        if visit[s] {
            return
        }
        for i := 0; i < len(edges); i++ { // 如果在所有的边里，有链接这个节点，就递归dfs
            e := edges[i]
            if e[0] == s {
                visit[s] = true
                dfs(edges, visit, e[1])
            } else if e[1] == s {
                visit[s] = true
                dfs(edges, visit, e[0])
            }
        }
    }
    for i := 0; i < n; i++ {
        if !visit[i] { // 存在没有关联上的点
            dfs(edges, visit, i)
            res++
        }
    }
    return res
}

func countComponents1(n int, edges [][]int) int {
    graph := NewGraph(n)
    for _, edge := range edges {
        graph.AddEdge(edge[0], edge[1])
    }
    //fmt.Println(graph)
    return graph.CountComponents()
}

type Graph struct {
    adjList [][]int
}

func NewGraph(n int) *Graph {
    return &Graph{ adjList: make([][]int, n),  }
}

func (g *Graph) AddEdge(u, v int) {
    g.adjList[u] = append(g.adjList[u], v)
    g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) dfs(visited []bool, src int) {
    visited[src] = true
    for _, neighbor := range g.adjList[src] {
        if !visited[neighbor] {
            g.dfs(visited, neighbor)
        }
    }
}

func (g *Graph) CountComponents() int {
    visited, cnt :=  make([]bool, len(g.adjList)), 0
    for i := 0; i < len(g.adjList); i++ {
        if !visited[i] {
            cnt++
            g.dfs(visited, i)
        }
    }
    return cnt
}

// 并查集
func countComponents2(n int, edges [][]int) int {
    visited, cnt:= make([]int, n), n
    for i := range visited {
        visited[i] = i 
    }
    var find func(u int) int
    find = func(u int) int {
        if visited[u] != u {
            visited[u] = find(visited[u])
        }
        return visited[u]
    }
    join := func (u, v int) {
        u, v = find(u), find(v)
        if u != v {
            cnt--
            visited[u] = v
        }
    }
    for _, e := range edges {
        join(e[0], e[1])
    }
    return cnt
}

func main() {
    fmt.Println(countComponents(5,[][]int{{0,1},{1,2},{3,4}})) // 2
    fmt.Println(countComponents(5,[][]int{{0,1},{1,2},{2,3},{3,4}})) // 1

    fmt.Println(countComponents1(5,[][]int{{0,1},{1,2},{3,4}})) // 2
    fmt.Println(countComponents1(5,[][]int{{0,1},{1,2},{2,3},{3,4}})) // 1

    fmt.Println(countComponents2(5,[][]int{{0,1},{1,2},{3,4}})) // 2
    fmt.Println(countComponents2(5,[][]int{{0,1},{1,2},{2,3},{3,4}})) // 1
}