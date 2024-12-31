package main

// 3313. Find the Last Marked Nodes in Tree
// There exists an undirected tree with n nodes numbered 0 to n - 1. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the tree.

// Initially, all nodes are unmarked. 
// After every second, you mark all unmarked nodes which have at least one marked node adjacent to them.

// Return an array nodes where nodes[i] is the last node to get marked in the tree, 
// if you mark node i at time t = 0. If nodes[i] has multiple answers for any node i, 
// you can choose any one answer.

// Example 1:
// Input: edges = [[0,1],[0,2]]
// Output: [2,2,1]
// Explanation:
//         0
//      /     \
//     1       2
// <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" />
// For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2]. Either 1 or 2 can be the answer.
// For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2]. Node 2 is marked last.
// For i = 2, the nodes are marked in the sequence: [2] -> [0,2] -> [0,1,2]. Node 1 is marked last.

// Example 2:
// Input: edges = [[0,1]]
// Output: [1,0]
// Explanation:
//         1
//       /
//      0
// <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png">
// For i = 0, the nodes are marked in the sequence: [0] -> [0,1].
// For i = 1, the nodes are marked in the sequence: [1] -> [0,1].

// Example 3:
// Input: edges = [[0,1],[0,2],[2,3],[2,4]]
// Output: [3,3,1,1,1]
// Explanation:
//         0
//      /     \
//     1       2
//            /  \
//           3    4
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" />
// For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2] -> [0,1,2,3,4].
// For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2] -> [0,1,2,3,4].
// For i = 2, the nodes are marked in the sequence: [2] -> [0,2,3,4] -> [0,1,2,3,4].
// For i = 3, the nodes are marked in the sequence: [3] -> [2,3] -> [0,2,3,4] -> [0,1,2,3,4].
// For i = 4, the nodes are marked in the sequence: [4] -> [2,4] -> [0,2,3,4] -> [0,1,2,3,4].

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= edges[i][0], edges[i][1] <= n - 1
//     The input is generated such that edges represents a valid tree.

import "fmt"

func lastMarkedNodes(edges [][]int) []int {
    n := len(edges) + 1
    res, graph := []int{}, make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    var dfs func(i, parent int, dist []int)
    dfs = func(i, parent int, dist []int) {
        for _, j := range graph[i] {
            if j != parent {
                dist[j] = dist[i] + 1
                dfs(j, i, dist)
            }
        }
    }
    maxNode := func(dist []int) int {
        mx := 0
        for i, d := range dist {
            if dist[mx] < d {
                mx = i
            }
        }
        return mx
    }
    dist1 := make([]int, n)
    dfs(0, -1, dist1)
    a := maxNode(dist1)
    dist2 := make([]int, n)
    dfs(a, -1, dist2)
    b := maxNode(dist2)
    dist3 := make([]int, n)
    dfs(b, -1, dist3)
    for i, x := range dist2 {
        if x > dist3[i] {
            res = append(res, a)
        } else {
            res = append(res, b)
        }
    }
    return res
}

func lastMarkedNodes1(edges [][]int) []int {
    n := len(edges) + 1
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    mx, u := -1, 0
    var findMaxDepth func(v, parent, d int)
    findMaxDepth = func(v, parent, d int) {
        if d > mx {
            mx, u = d, v
        }
        for _, w := range graph[v] {
            if w != parent {
                findMaxDepth(w, v, d + 1)
            }
        }
    }
    findMaxDepth(0, -1, 0)
    dv := u
    mx = -1
    findMaxDepth(u, -1, 0)
    dw := u
    res, farthest := make([]int, n), make([]int, n)
    for i := range farthest {
        farthest[i] = -1
    }
    var findFarthest func(v, parent, d, target int)
    findFarthest = func(v, parent, d, target int) {
        if d > farthest[v] {
            farthest[v] = d
            res[v] = target
        }
        for _, w := range graph[v] {
            if w != parent {
                findFarthest(w, v, d + 1, target)
            }
        }
    }
    findFarthest(dv, -1, 0, dv)
    findFarthest(dw, -1, 0, dw)
    return res
}

func main() {
    // Example 1:
    // Input: edges = [[0,1],[0,2]]
    // Output: [2,2,1]
    // Explanation:
    //         0
    //      /     \
    //     1       2
    // <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" />
    // For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2]. Either 1 or 2 can be the answer.
    // For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2]. Node 2 is marked last.
    // For i = 2, the nodes are marked in the sequence: [2] -> [0,2] -> [0,1,2]. Node 1 is marked last.
    fmt.Println(lastMarkedNodes([][]int{{0,1},{0,2}})) // [2,2,1]
    // Example 2:
    // Input: edges = [[0,1]]
    // Output: [1,0]
    // Explanation:
    //         1
    //       /
    //      0
    // <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png">
    // For i = 0, the nodes are marked in the sequence: [0] -> [0,1].
    // For i = 1, the nodes are marked in the sequence: [1] -> [0,1].
    fmt.Println(lastMarkedNodes([][]int{{0,1}})) // [1,0]
    // Example 3:
    // Input: edges = [[0,1],[0,2],[2,3],[2,4]]
    // Output: [3,3,1,1,1]
    // Explanation:
    //         0
    //      /     \
    //     1       2
    //            /  \
    //           3    4
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" />
    // For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2] -> [0,1,2,3,4].
    // For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2] -> [0,1,2,3,4].
    // For i = 2, the nodes are marked in the sequence: [2] -> [0,2,3,4] -> [0,1,2,3,4].
    // For i = 3, the nodes are marked in the sequence: [3] -> [2,3] -> [0,2,3,4] -> [0,1,2,3,4].
    // For i = 4, the nodes are marked in the sequence: [4] -> [2,4] -> [0,2,3,4] -> [0,1,2,3,4].
    fmt.Println(lastMarkedNodes([][]int{{0,1},{0,2},{2,3},{2,4}})) // [3,3,1,1,1]

    fmt.Println(lastMarkedNodes1([][]int{{0,1},{0,2}})) // [2,2,1]
    fmt.Println(lastMarkedNodes1([][]int{{0,1}})) // [1,0]
    fmt.Println(lastMarkedNodes1([][]int{{0,1},{0,2},{2,3},{2,4}})) // [3,3,1,1,1]
}