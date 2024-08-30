package main

// 2699. Modify Graph Edge Weights
// You are given an undirected weighted connected graph containing n nodes labeled from 0 to n - 1, 
// and an integer array edges where edges[i] = [ai, bi, wi] indicates 
// that there is an edge between nodes ai and bi with weight wi.

// Some edges have a weight of -1 (wi = -1), while others have a positive weight (wi > 0).

// Your task is to modify all edges with a weight of -1 by assigning them positive integer values in the range [1, 2 * 10^9] so that the shortest distance between the nodes source and destination becomes equal to an integer target. 
// If there are multiple modifications that make the shortest distance between source and destination equal to target, any of them will be considered correct.

// Return an array containing all edges (even unmodified ones) in any order if it is possible to make the shortest distance from source to destination equal to target, or an empty array if it's impossible.

// Note: You are not allowed to modify the weights of edges with initial positive weights.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/04/18/graph.png" />
// Input: n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0, destination = 1, target = 5
// Output: [[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
// Explanation: The graph above shows a possible modification to the edges, making the distance from 0 to 1 equal to 5.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/04/18/graph-2.png" />
// Input: n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target = 6
// Output: []
// Explanation: The graph above contains the initial edges. It is not possible to make the distance from 0 to 2 equal to 6 by modifying the edge with weight -1. So, an empty array is returned.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2023/04/19/graph-3.png" />
// Input: n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0, destination = 2, target = 6
// Output: [[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
// Explanation: The graph above shows a modified graph having the shortest distance from 0 to 2 as 6.

// Constraints:
//     1 <= n <= 10^0
//     1 <= edges.length <= n * (n - 1) / 2
//     edges[i].length == 3
//     0 <= ai, bi < n
//     wi = -1 or 1 <= wi <= 10^7
//     ai != bi
//     0 <= source, destination < n
//     source != destination
//     1 <= target <= 10^9
//     The graph is connected, and there are no self-loops or repeated edges

import "fmt"

// dijkstra
func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
    adjMatrix := make([][]int, n)
    for i := 0; i < n; i++ {
        adjMatrix[i] = make([]int, n)
        for j := 0; j < n; j++ {
            adjMatrix[i][j] = -1
        }
    }
    // 邻接矩阵中存储边的下标
    for i, e := range edges {
        u, v := e[0], e[1]
        adjMatrix[u][v], adjMatrix[v][u] = i, i
    }
    min := func(a, b int64) int64 { if a > b { return b; }; return a; }
    dijkstra := func(op, source int, edges [][]int, adjMatrix [][]int, target int, fromDestination []int64) []int64 {
        // 朴素的 dijistra 算法
        // adjMatrix 是一个邻接矩阵
        n := len(adjMatrix)
        dist, used := make([]int64, n), make([]bool, n)
        for i := 0; i < n; i++ {
            dist[i] = 0x3f3f3f3f3f
        }
        dist[source] = 0
        for round := 0; round < n - 1; round++ {
            u := -1
            for i := 0; i < n; i++ {
                if !used[i] && (u == -1 || dist[i] < dist[u]) {
                    u = i
                }
            }
            used[u] = true
            for v := 0; v < n; v++ {
                if !used[v] && adjMatrix[u][v] != -1 {
                    i := adjMatrix[u][v]
                    if edges[i][2] != -1 {
                        dist[v] = min(dist[v], dist[u] + int64(edges[i][2]))
                    } else {
                        if op == 0 {
                            dist[v] = min(dist[v], dist[u] + 1)
                        } else {
                            modify := int64(target) - dist[u] - fromDestination[v]
                            if modify > 0 {
                                dist[v] = min(dist[v], dist[u] + modify)
                                edges[i][2] = int(modify)
                            } else {
                                edges[i][2] = target
                            }
                        }
                    }
                }
            }
        }
        return dist
    }    
    fromDestination := dijkstra(0, destination, edges, adjMatrix, target, nil)
    if fromDestination[source] > int64(target) {
        return nil
    }
    fromSource := dijkstra(1, source, edges, adjMatrix, target, fromDestination)
    if fromSource[destination] != int64(target) {
        return nil
    }
    return edges
}

func modifiedGraphEdges1(n int, edges [][]int, source int, destination int, target int) [][]int {
    k := 0
    for _, e := range edges {
        if e[2] == -1 {
            k++
        }
    }
    dijkstra := func(source, destination int, adjMatrix [][]int) int64 {
        // 朴素的 dijistra 算法
        // adjMatrix 是一个邻接矩阵
        n := len(adjMatrix)
        dist, used := make([]int64, n), make([]bool, n)
        for i := 0; i < n; i++ {
            dist[i] = 0x3f3f3f3f3f
        }
        dist[source] = 0
        for round := 0; round < n - 1; round++ {
            u := -1
            for i := 0; i < n; i++ {
                if !used[i] && (u == -1 || dist[i] < dist[u]) {
                    u = i
                }
            }
            used[u] = true
            for v := 0; v < n; v++ {
                if !used[v] && adjMatrix[u][v] != -1 && dist[v] > dist[u] + int64(adjMatrix[u][v]) {
                    dist[v] = dist[u] + int64(adjMatrix[u][v])
                }
            }
        }
        return dist[destination]
    }
    construct := func(n int, edges [][]int, idx int64, target int) [][]int {
        // 需要构造出第 idx 种不同的边权情况，返回一个邻接矩阵
        adjMatrix := make([][]int, n)
        for i := 0; i < n; i++ {
            adjMatrix[i] = make([]int, n)
            for j := 0; j < n; j++ {
                adjMatrix[i][j] = -1
            }
        }
        for _, e := range edges {
            u, v, w := e[0], e[1], e[2]
            if w != -1 {
                adjMatrix[u][v], adjMatrix[v][u] = w, w
            } else {
                if idx >= int64(target - 1) {
                    adjMatrix[u][v], adjMatrix[v][u] = target, target
                    idx -= int64(target - 1)
                } else {
                    adjMatrix[u][v], adjMatrix[v][u] = int(1 + idx), int(1 + idx)
                    idx = 0
                }
            }
        }
        return adjMatrix
    }
    if dijkstra(source, destination, construct(n, edges, 0, target)) > int64(target) {
        return nil
    }
    if dijkstra(source, destination, construct(n, edges, int64(k) * int64(target - 1), target)) < int64(target) {
        return nil
    }
    left, right, ans := int64(0), int64(k) * int64(target - 1), int64(0)
    for left <= right {
        mid := int64(left + right) / 2
        if dijkstra(source, destination, construct(n, edges, mid, target)) >= int64(target) {
            ans, right = mid, mid - 1
        } else {
            left = mid + 1
        }
    }
    for _, e := range edges {
        if e[2] == -1 {
            if ans >= int64(target - 1) {
                e[2] = target
                ans -= int64(target - 1)
            } else {
                e[2] = int(1 + ans)
                ans = 0
            }
        }
    }
    return edges
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/04/18/graph.png" />
    // Input: n = 5, edges = [[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]], source = 0, destination = 1, target = 5
    // Output: [[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
    // Explanation: The graph above shows a possible modification to the edges, making the distance from 0 to 1 equal to 5.
    fmt.Println(modifiedGraphEdges(5,[][]int{{4,1,-1},{2,0,-1},{0,3,-1},{4,3,-1}},0,1,5)) // [[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/04/18/graph-2.png" />
    // Input: n = 3, edges = [[0,1,-1],[0,2,5]], source = 0, destination = 2, target = 6
    // Output: []
    // Explanation: The graph above contains the initial edges. It is not possible to make the distance from 0 to 2 equal to 6 by modifying the edge with weight -1. So, an empty array is returned.
    fmt.Println(modifiedGraphEdges(3,[][]int{{0,1,-1},{0,2,5}},0,2,6)) // []
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2023/04/19/graph-3.png" />
    // Input: n = 4, edges = [[1,0,4],[1,2,3],[2,3,5],[0,3,-1]], source = 0, destination = 2, target = 6
    // Output: [[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
    // Explanation: The graph above shows a modified graph having the shortest distance from 0 to 2 as 6.
    fmt.Println(modifiedGraphEdges(4,[][]int{{1,0,4},{1,2,3},{2,3,5},{0,3,-1}},0,2,6)) // [[1,0,4],[1,2,3],[2,3,5],[0,3,1]]

    fmt.Println(modifiedGraphEdges1(5,[][]int{{4,1,-1},{2,0,-1},{0,3,-1},{4,3,-1}},0,1,5)) // [[4,1,1],[2,0,1],[0,3,3],[4,3,1]]
    fmt.Println(modifiedGraphEdges1(3,[][]int{{0,1,-1},{0,2,5}},0,2,6)) // []
    fmt.Println(modifiedGraphEdges1(4,[][]int{{1,0,4},{1,2,3},{2,3,5},{0,3,-1}},0,2,6)) // [[1,0,4],[1,2,3],[2,3,5],[0,3,1]]
}