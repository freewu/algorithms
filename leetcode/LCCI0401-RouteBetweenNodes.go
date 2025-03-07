package main

// 面试题 04.01. Route Between Nodes LCCI
// Given a directed graph, design an algorithm to find out whether there is a route between two nodes.

// Example1:
// Input: n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
// Output: true

// Example2:
// Input: n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
// Output true

// Note:
//     0 <= n <= 100000
//     All node numbers are within the range [0, n].
//     There might be self cycles and duplicated edges.

import "fmt"

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
    edge, visited := make([][]int, n), make(map[int]bool)
    for i := 0; i < len(graph); i++ {
        u, v := graph[i][0], graph[i][1]
        edge[u] = append(edge[u], v)
    }
    var dfs func(start int) bool
    dfs = func(start int) bool {
        if start == target { return true }
        if visited[start] { return false }
        visited[start] = true
        for _, v := range edge[start] {
            if dfs(v) {
                return true
            }
        }
        return false
    }
    return dfs(start)
}

func findWhetherExistsPath1(n int, graph [][]int, start int, target int) bool {
    if start == target { return true }
    visited := make(map[int]bool)
    var dfs func(start, target int) bool
    dfs = func(start, target int) bool {
        for i := 0; i < len(graph); i ++ {
            if !visited[i] && graph[i][1] == target {
                if graph[i][0] == start { return true }
                visited[i] = true
                if dfs(start, graph[i][0]) { return true }
                visited[i] = false
            }
        }
        return false
    }
    return dfs(start, target)
}

func main() {
    // Example1:
    // Input: n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
    // Output: true
    fmt.Println(findWhetherExistsPath(3, [][]int{{0,1},{0,2},{1,2},{1,2}}, 0, 2)) // true
    // Example2:
    // Input: n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [1, 3], [2, 3], [3, 4]], start = 0, target = 4
    // Output true
    fmt.Println(findWhetherExistsPath(5, [][]int{{0,1},{0,2},{0,4},{0,4},{0,1},{1,3},{1,4},{1,3},{2,3},{3,4}}, 0, 4)) // true

    fmt.Println(findWhetherExistsPath1(3, [][]int{{0,1},{0,2},{1,2},{1,2}}, 0, 2)) // true
    fmt.Println(findWhetherExistsPath1(5, [][]int{{0,1},{0,2},{0,4},{0,4},{0,1},{1,3},{1,4},{1,3},{2,3},{3,4}}, 0, 4)) // true
}