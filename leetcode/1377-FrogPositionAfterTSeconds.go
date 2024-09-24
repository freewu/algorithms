package main

// 1377. Frog Position After T Seconds
// Given an undirected tree consisting of n vertices numbered from 1 to n. 
// A frog starts jumping from vertex 1. 
// In one second, the frog jumps from its current vertex to another unvisited vertex if they are directly connected. 
// The frog can not jump back to a visited vertex. 
// In case the frog can jump to several vertices, it jumps randomly to one of them with the same probability. 
// Otherwise, when the frog can not jump to any unvisited vertex, it jumps forever on the same vertex.

// The edges of the undirected tree are given in the array edges, where edges[i] = [ai, bi] means that exists an edge connecting the vertices ai and bi.

// Return the probability that after t seconds the frog is on the vertex target. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// <img src="" />
// Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 2, target = 4
// Output: 0.16666666666666666 
// Explanation: The figure above shows the given graph. The frog starts at vertex 1, jumping with 1/3 probability to the vertex 2 after second 1 and then jumping with 1/2 probability to vertex 4 after second 2. Thus the probability for the frog is on the vertex 4 after 2 seconds is 1/3 * 1/2 = 1/6 = 0.16666666666666666. 

// Example 2:
// <img src="" />
// Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 1, target = 7
// Output: 0.3333333333333333
// Explanation: The figure above shows the given graph. The frog starts at vertex 1, jumping with 1/3 = 0.3333333333333333 probability to the vertex 7 after second 1. 

// Constraints:
//     1 <= n <= 100
//     edges.length == n - 1
//     edges[i].length == 2
//     1 <= ai, bi <= n
//     1 <= t <= 50
//     1 <= target <= n

import "fmt"

func frogPosition(n int, edges [][]int, t int, target int) float64 {
    graph := make([][]int, n+1)
    for _, edge := range edges { // 邻接表
        x, y := edge[0], edge[1]
        graph[x] = append(graph[x], y)
        graph[y] = append(graph[y], x)
    }
    visited := make([]bool, n + 1)
    visited[0], visited[1] = true, true
    probabilities := make([]float64, n + 1)
    probabilities[1] = 1
    cur := []int{1}
    for time := 0; time < t; time++ {
        next := []int{}
        for _, x := range cur {
            count := 0
            for _, y := range graph[x] {
                if !visited[y] {
                    count++
                    next = append(next, y)
                }
            }
            if count != 0 {
                for _, y := range graph[x] {
                    if !visited[y] {
                        probabilities[y] = probabilities[x] / float64(count)
                        visited[y] = true
                    }
                }
                probabilities[x] = 0
            }
        }
        cur = next
    }
    return probabilities[target]
}

func frogPosition1(n int, edges [][]int, t int, target int) float64 {
    getGraph := func(edges [][]int, n int) [][]int {
        res := make([][]int, n+1)
        res[1] = append(res[0], 0)
        for _, edge := range edges {
            x, y := edge[0], edge[1]
            res[x] = append(res[x], y)
            res[y] = append(res[y], x)
        }
        return res
    }
    graph := getGraph(edges, n)
    var dfs func (index int, parent int, graph [][]int, t int, target int) float64
    dfs = func (index int, parent int, graph [][]int, t int, target int) float64 {
        cur := graph[index]
        if t == 0 {
            if index == target { return 1 }
            return 0
        }
        if index == target {
            if len(cur) == 1 { return 1 }
            return 0
        }
        for _, child := range cur {
            if child != parent {
                res := dfs(child, index, graph, t-1, target)
                if res != 0 {
                    return res * 1 / float64(len(cur) - 1)
                }
            }
        }
        return 0
    }
    return dfs(1, 0, graph, t, target)
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 2, target = 4
    // Output: 0.16666666666666666 
    // Explanation: The figure above shows the given graph. The frog starts at vertex 1, jumping with 1/3 probability to the vertex 2 after second 1 and then jumping with 1/2 probability to vertex 4 after second 2. Thus the probability for the frog is on the vertex 4 after 2 seconds is 1/3 * 1/2 = 1/6 = 0.16666666666666666. 
    fmt.Println(frogPosition(7,[][]int{{1,2},{1,3},{1,7},{2,4},{2,6},{3,5}}, 2, 4)) // 0.16666666666666666 
    // Example 2:
    // <img src="" />
    // Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 1, target = 7
    // Output: 0.3333333333333333
    // Explanation: The figure above shows the given graph. The frog starts at vertex 1, jumping with 1/3 = 0.3333333333333333 probability to the vertex 7 after second 1. 
    fmt.Println(frogPosition(7,[][]int{{1,2},{1,3},{1,7},{2,4},{2,6},{3,5}}, 1, 7)) // 0.3333333333333333 

    fmt.Println(frogPosition1(7,[][]int{{1,2},{1,3},{1,7},{2,4},{2,6},{3,5}}, 2, 4)) // 0.16666666666666666 
    fmt.Println(frogPosition1(7,[][]int{{1,2},{1,3},{1,7},{2,4},{2,6},{3,5}}, 1, 7)) // 0.3333333333333333 
}