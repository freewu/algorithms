package main

// 1129. Shortest Path with Alternating Colors
// You are given an integer n, the number of nodes in a directed graph where the nodes are labeled from 0 to n - 1. 
// Each edge is red or blue in this graph, and there could be self-edges and parallel edges.
// You are given two arrays redEdges and blueEdges where:
//     redEdges[i] = [ai, bi] indicates that there is a directed red edge from node ai to node bi in the graph, and
//     blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node uj to node vj in the graph.
    
// Return an array answer of length n, where each answer[x] is the length of the shortest path from node 0 to node x such 
// that the edge colors alternate along the path, or -1 if such a path does not exist.

// Example 1:
// Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
// Output: [0,1,-1]

// Example 2:
// Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
// Output: [0,1,-1]

// Constraints:
//     1 <= n <= 100
//     0 <= redEdges.length, blueEdges.length <= 400
//     redEdges[i].length == blueEdges[j].length == 2
//     0 <= ai, bi, uj, vj < n

import "fmt"

// Bellman-Ford Algorithm
func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
    flag, redDist, blueDist, inf := true, make([]int, n), make([]int, n), 1 << 32 - 1
    for i := 1; i < n; i++ {
        redDist[i], blueDist[i] = inf, inf
    }
    for flag {
        flag = false
        for _, e := range red_edges {
            if d := blueDist[e[0]]; d != inf {
                if d + 1 < redDist[e[1]] {
                    redDist[e[1]] = d+1
                    flag = true
                }
            }
        }
        for _, e := range blue_edges {
            if d := redDist[e[0]]; d != inf {
                if d+1 < blueDist[e[1]] {
                    blueDist[e[1]] = d+1
                    flag = true
                }
            }
        }
    }
    for i, d := range redDist {
        if d > blueDist[i] {
            redDist[i] = blueDist[i]
        }
        if redDist[i] == inf {
            redDist[i] = -1
        }
    }
    return redDist
}

func shortestAlternatingPaths1(n int, redEdges [][]int, blueEdges [][]int) []int {
    type Edge struct{ dest, color int }
    type Graph [][]Edge
    new_graph := func (n int, redEdges [][]int, blueEdges [][]int) Graph {
        res := make(Graph, n)
        for _, e := range redEdges {
            res[e[0]] = append(res[e[0]], Edge{ e[1], 0 })
        }
        for _, e := range blueEdges {
            res[e[0]] = append(res[e[0]], Edge{ e[1], 1 })
        }
        return res
    }
    g := new_graph(n, redEdges, blueEdges)
    res, visited, queue := make([]int, n), make([][2]bool, n), make([]Edge, 2)
    visited[0][0], visited[0][1] = true, true
    queue[0], queue[1] = Edge{ 0, 0 }, Edge{ 0, 1 }
    for i := range res {
        res[i] = -1
    }
    for level := 0; len(queue) > 0; level++ {
        nextLevel := make([]Edge, 0)
        for _, from := range queue {
            if res[from.dest] < 0 {
                res[from.dest] = level
            }
            for _, to := range g[from.dest] {
                if from.color == to.color || visited[to.dest][to.color] {
                    continue
                }
                visited[to.dest][to.color] = true
                nextLevel = append(nextLevel, to)
            }
        }
        queue = nextLevel
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, redEdges = [[0,1],[1,2]], blueEdges = []
    // Output: [0,1,-1]
    fmt.Println(shortestAlternatingPaths(3, [][]int{{0,1},{1,2}},[][]int{})) // [0,1,-1]
    // Example 2:
    // Input: n = 3, redEdges = [[0,1]], blueEdges = [[2,1]]
    // Output: [0,1,-1]
    fmt.Println(shortestAlternatingPaths(3, [][]int{{0,1}},[][]int{{2,1}})) // [0,1,-1]

    fmt.Println(shortestAlternatingPaths1(3, [][]int{{0,1},{1,2}},[][]int{})) // [0,1,-1]
    fmt.Println(shortestAlternatingPaths1(3, [][]int{{0,1}},[][]int{{2,1}})) // [0,1,-1]
}