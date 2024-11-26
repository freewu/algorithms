package main

// 2737. Find the Closest Marked Node 
// You are given a positive integer n which is the number of nodes of a 0-indexed directed weighted graph 
// and a 0-indexed 2D array edges where edges[i] = [ui, vi, wi] indicates
// that there is an edge from node ui to node vi with weight wi.

// You are also given a node s and a node array marked; 
// your task is to find the minimum distance from s to any of the nodes in marked.

// Return an integer denoting the minimum distance from s to any node in marked 
// or -1 if there are no paths from s to any of the marked nodes.

// Example 1:
// Input: n = 4, edges = [[0,1,1],[1,2,3],[2,3,2],[0,3,4]], s = 0, marked = [2,3]
// Output: 4
// Explanation: There is one path from node 0 (the green node) to node 2 (a red node), which is 0->1->2, and has a distance of 1 + 3 = 4.
// There are two paths from node 0 to node 3 (a red node), which are 0->1->2->3 and 0->3, the first one has a distance of 1 + 3 + 2 = 6 and the second one has a distance of 4.
// The minimum of them is 4.
// <img src="https://assets.leetcode.com/uploads/2023/06/13/image_2023-06-13_16-34-38.png" />

// Example 2:
// Input: n = 5, edges = [[0,1,2],[0,2,4],[1,3,1],[2,3,3],[3,4,2]], s = 1, marked = [0,4]
// Output: 3
// Explanation: There are no paths from node 1 (the green node) to node 0 (a red node).
// There is one path from node 1 to node 4 (a red node), which is 1->3->4, and has a distance of 1 + 2 = 3.
// So the answer is 3.
// <img src="https://assets.leetcode.com/uploads/2023/06/13/image_2023-06-13_16-35-13.png" />

// Example 3:
// Input: n = 4, edges = [[0,1,1],[1,2,3],[2,3,2]], s = 3, marked = [0,1]
// Output: -1
// Explanation: There are no paths from node 3 (the green node) to any of the marked nodes (the red nodes), so the answer is -1.
// <img src="https://assets.leetcode.com/uploads/2023/06/13/image_2023-06-13_16-35-47.png" />

// Constraints:
//     2 <= n <= 500
//     1 <= edges.length <= 10^4
//     edges[i].length = 3
//     0 <= edges[i][0], edges[i][1] <= n - 1
//     1 <= edges[i][2] <= 10^6
//     1 <= marked.length <= n - 1
//     0 <= s, marked[i] <= n - 1
//     s != marked[i]
//     marked[i] != marked[j] for every i != j
//     The graph might have repeated edges.
//     The graph is generated such that it has no self-loops.

import "fmt"

func minimumDistance(n int, edges [][]int, s int, marked []int) int {
    inf := 1 << 31
    graph, dist := make([][]int, n), make([]int, n)
    for i := range graph {
        graph[i] = make([]int, n)
        for j := range graph[i] {
            graph[i][j] = inf
        }
        dist[i] = inf
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    dist[s] = 0
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        graph[u][v] = min(graph[u][v], w)
    }
    visited := make([]bool, n)
    for _ = range graph {
        t := -1
        for j := 0; j < n; j++ {
            if !visited[j] && (t == -1 || dist[j] < dist[t]) {
                t = j
            }
        }
        visited[t] = true
        for j := 0; j < n; j++ {
            dist[j] = min(dist[j], dist[t] + graph[t][j])
        }
    }
    res := inf
    for _, i := range marked {
        res = min(res, dist[i])
    }
    if res >= inf { return -1 }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, edges = [[0,1,1],[1,2,3],[2,3,2],[0,3,4]], s = 0, marked = [2,3]
    // Output: 4
    // Explanation: There is one path from node 0 (the green node) to node 2 (a red node), which is 0->1->2, and has a distance of 1 + 3 = 4.
    // There are two paths from node 0 to node 3 (a red node), which are 0->1->2->3 and 0->3, the first one has a distance of 1 + 3 + 2 = 6 and the second one has a distance of 4.
    // The minimum of them is 4.
    // <img src="https://assets.leetcode.com/uploads/2023/06/13/image_2023-06-13_16-34-38.png" />
    fmt.Println(minimumDistance(4,[][]int{{0,1,1},{1,2,3},{2,3,2},{0,3,4}}, 0, []int{2,3})) // 4
    // Example 2:
    // Input: n = 5, edges = [[0,1,2],[0,2,4],[1,3,1],[2,3,3],[3,4,2]], s = 1, marked = [0,4]
    // Output: 3
    // Explanation: There are no paths from node 1 (the green node) to node 0 (a red node).
    // There is one path from node 1 to node 4 (a red node), which is 1->3->4, and has a distance of 1 + 2 = 3.
    // So the answer is 3.
    // <img src="https://assets.leetcode.com/uploads/2023/06/13/image_2023-06-13_16-35-13.png" />
    fmt.Println(minimumDistance(5,[][]int{{0,1,2},{0,2,4},{1,3,1},{2,3,3},{3,4,2}}, 1, []int{0,4})) // 3
    // Example 3:
    // Input: n = 4, edges = [[0,1,1],[1,2,3],[2,3,2]], s = 3, marked = [0,1]
    // Output: -1
    // Explanation: There are no paths from node 3 (the green node) to any of the marked nodes (the red nodes), so the answer is -1.
    // <img src="https://assets.leetcode.com/uploads/2023/06/13/image_2023-06-13_16-35-47.png" />
    fmt.Println(minimumDistance(4,[][]int{{0,1,1},{1,2,3},{2,3,2}}, 3, []int{0,1})) // -1
}