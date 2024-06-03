package main

// 847. Shortest Path Visiting All Nodes
// You have an undirected, connected graph of n nodes labeled from 0 to n - 1. 
// You are given an array graph where graph[i] is a list of all the nodes connected with node i by an edge.

// Return the length of the shortest path that visits every node. 
// You may start and stop at any node, you may revisit nodes multiple times, and you may reuse edges.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/12/shortest1-graph.jpg" />
// Input: graph = [[1,2,3],[0],[0],[0]]
// Output: 4
// Explanation: One possible path is [1,0,2,0,3]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/12/shortest2-graph.jpg" />
// Input: graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]
// Output: 4
// Explanation: One possible path is [0,1,4,2,3]

// Constraints:
//     n == graph.length
//     1 <= n <= 12
//     0 <= graph[i].length < n
//     graph[i] does not contain i.
//     If graph[a] contains b, then graph[b] contains a.
//     The input graph is always connected.

import "fmt"

func shortestPathLength(graph [][]int) int {
    type State struct {
        mask, node, dist int
    }
    n := len(graph)
    visited, queue, allVisited := make(map[int]bool), []State{}, (1 << n) - 1
    for i := 0; i < n; i++ {
        queue = append(queue, State{1 << i, i, 0})
        visited[(1<<i) * 16 + i] = true
    }
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        if cur.mask == allVisited {
            return cur.dist
        }
        for _, neighbor := range graph[cur.node] {
            newMask := cur.mask | (1 << neighbor)
            hashValue := newMask * 16 + neighbor
            if !visited[hashValue] {
                visited[hashValue] = true
                queue = append(queue, State{newMask, neighbor, cur.dist + 1})
            }
        }
    }
    return -1
}


// # 思路：
//      1. 求访问所有节点，应该想到BFS，DFS
//      2. 求最短路径，因为是等权，所以可以考虑BFS
//      3. BFS遍历过程中，需要当前遍历的节点、记录遍历过哪些节点、当前已走路径的长度，且n=12，可考虑状态压缩
//      使用{u, mask, step}三元组记录状态，u：当前节点；mask：遍历过的节点；step：已走过的路径长度
//      最先到达mask == 2^n - 1的step则为解
func shortestPathLength1(graph [][]int) int {
    n := len(graph)
    type tuple struct { u, mask, step int }
    seen, queue := make([][]bool, n), []tuple{}
    for i := range seen {
        seen[i] = make([]bool, 1<<n)
    }
    for i := 0; i < n; i ++ { // 初始化
        queue = append(queue , tuple{i, 1 << i, 0})
    }
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        u, mask, step := cur.u, cur.mask, cur.step
        if mask == 1 << n - 1 {
            return step
        }
        for _, v := range graph[u] {
            newMask := mask | (1<<v)
            if !seen[v][newMask] {
                seen[v][newMask] = true
                queue = append(queue, tuple{v, newMask, step + 1})
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/12/shortest1-graph.jpg" />
    // Input: graph = [[1,2,3],[0],[0],[0]]
    // Output: 4
    // Explanation: One possible path is [1,0,2,0,3]
    fmt.Println(shortestPathLength([][]int{{1,2,3},{0},{0},{0}})) // 4
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/12/shortest2-graph.jpg" />
    // Input: graph = [[1],[0,2,4],[1,3,4],[2],[1,2]]
    // Output: 4
    // Explanation: One possible path is [0,1,4,2,3]
    fmt.Println(shortestPathLength([][]int{{1},{0,2,4},{1,3,4},{2},{1,2}})) // 4

    fmt.Println(shortestPathLength1([][]int{{1,2,3},{0},{0},{0}})) // 4
    fmt.Println(shortestPathLength1([][]int{{1},{0,2,4},{1,3,4},{2},{1,2}})) // 4
}