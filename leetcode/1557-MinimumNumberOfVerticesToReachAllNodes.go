package main

// 1557. Minimum Number of Vertices to Reach All Nodes
// Given a directed acyclic graph, with n vertices numbered from 0 to n-1, 
// and an array edges where edges[i] = [fromi, toi] represents a directed edge from node fromi to node toi.

// Find the smallest set of vertices from which all nodes in the graph are reachable. 
// It's guaranteed that a unique solution exists.

// Notice that you can return the vertices in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/07/untitled22.png" />
// Input: n = 6, edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]
// Output: [0,3]
// Explanation: It's not possible to reach all the nodes from a single vertex. From 0 we can reach [0,1,2,5]. From 3 we can reach [3,4,2,5]. So we output [0,3].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/07/untitled.png" />
// Input: n = 5, edges = [[0,1],[2,1],[3,1],[1,4],[2,4]]
// Output: [0,2,3]
// Explanation: Notice that vertices 0, 3 and 2 are not reachable from any other node, so we must include them. Also any of these vertices can reach nodes 1 and 4.

// Constraints:
//     2 <= n <= 10^5
//     1 <= edges.length <= min(10^5, n * (n - 1) / 2)
//     edges[i].length == 2
//     0 <= fromi, toi < n
//     All pairs (fromi, toi) are distinct.

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
    isOutgoing := make([]bool, n)
    res := make([]int, 0, n)
    for _, edge := range edges {
        isOutgoing[edge[1]] = true
    }
    for i, outgoing := range isOutgoing {
        if !outgoing {
            res = append(res, i)
        }
    }
    return res
}

func findSmallestSetOfVertices1(n int, edges [][]int) []int {
    // 所有入度为0的点
    degree := make([]int, n)
    for _, e := range edges {
        degree[e[1]]++
    }
    res := []int{}
    for v, deg := range degree {
        if deg == 0 {
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/07/untitled22.png" />
    // Input: n = 6, edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]
    // Output: [0,3]
    // Explanation: It's not possible to reach all the nodes from a single vertex. From 0 we can reach [0,1,2,5]. From 3 we can reach [3,4,2,5]. So we output [0,3].
    fmt.Println(findSmallestSetOfVertices(6,[][]int{{0,1},{0,2},{2,5},{3,4},{4,2}})) // [0,3]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/07/untitled.png" />
    // Input: n = 5, edges = [[0,1],[2,1],[3,1],[1,4],[2,4]]
    // Output: [0,2,3]
    // Explanation: Notice that vertices 0, 3 and 2 are not reachable from any other node, so we must include them. Also any of these vertices can reach nodes 1 and 4.
    fmt.Println(findSmallestSetOfVertices(5,[][]int{{0,1},{2,1},{3,1},{1,4},{2,4}})) // [0,2,3]

    fmt.Println(findSmallestSetOfVertices1(6,[][]int{{0,1},{0,2},{2,5},{3,4},{4,2}})) // [0,3]
    fmt.Println(findSmallestSetOfVertices1(5,[][]int{{0,1},{2,1},{3,1},{1,4},{2,4}})) // [0,2,3]
}