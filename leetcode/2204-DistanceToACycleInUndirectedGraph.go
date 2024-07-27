package main

// 2204. Distance to a Cycle in Undirected Graph
// You are given a positive integer n representing the number of nodes in a connected undirected graph containing exactly one cycle. 
// The nodes are numbered from 0 to n - 1 (inclusive).

// You are also given a 2D integer array edges, where edges[i] = [node1i, node2i] denotes that there is a bidirectional edge connecting node1i and node2i in the graph.
// The distance between two nodes a and b is defined to be the minimum number of edges that are needed to go from a to b.
// Return an integer array answer of size n, where answer[i] is the minimum distance between the ith node and any node in the cycle.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315154238-1.png" />
// Input: n = 7, edges = [[1,2],[2,4],[4,3],[3,1],[0,1],[5,2],[6,5]]
// Output: [1,0,0,0,0,1,2]
// Explanation:
// The nodes 1, 2, 3, and 4 form the cycle.
// The distance from 0 to 1 is 1.
// The distance from 1 to 1 is 0.
// The distance from 2 to 2 is 0.
// The distance from 3 to 3 is 0.
// The distance from 4 to 4 is 0.
// The distance from 5 to 2 is 1.
// The distance from 6 to 2 is 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315154634-1.png" />
// Input: n = 9, edges = [[0,1],[1,2],[0,2],[2,6],[6,7],[6,8],[0,3],[3,4],[3,5]]
// Output: [0,0,0,1,2,2,1,2,2]
// Explanation:
// The nodes 0, 1, and 2 form the cycle.
// The distance from 0 to 0 is 0.
// The distance from 1 to 1 is 0.
// The distance from 2 to 2 is 0.
// The distance from 3 to 1 is 1.
// The distance from 4 to 1 is 2.
// The distance from 5 to 1 is 2.
// The distance from 6 to 2 is 1.
// The distance from 7 to 2 is 2.
// The distance from 8 to 2 is 2.

// Constraints:
//     3 <= n <= 10^5
//     edges.length == n
//     edges[i].length == 2
//     0 <= node1i, node2i <= n - 1
//     node1i != node2i
//     The graph is connected.
//     The graph has exactly one cycle.
//     There is at most one edge between any pair of vertices.

import "fmt"

func distanceToCycle(n int, edges [][]int) []int {
    res, g, deg := make([]int, n), make([][]int, n), make([]int, n)
    for _, e := range edges {  // 建图
        v, w := e[0], e[1]
        g[v] = append(g[v], w)
        g[w] = append(g[w], v)
        deg[v]++
        deg[w]++
    }
    q := []int{} // 拓扑排序，剪掉所有树枝
    for i, d := range deg {
        if d == 1 {
            q = append(q, i)
        }
    }
    for len(q) > 0 {
        v := q[0]
        q = q[1:]
        for _, w := range g[v] {
            if deg[w]--; deg[w] == 1 {
                q = append(q, w)
            }
        }
    }
    var f func(int, int) 
    f = func(v, fa int) { // 从基环出发，求所有树枝上的点的深度
        for _, w := range g[v] {
            if w != fa && deg[w] < 2 {
                res[w] = res[v] + 1
                f(w, v)
            }
        }
    }
    for root, d := range deg {
        if d > 1 {
            f(root, -1)
        }
    }
    return res
}

// bfs
func distanceToCycle1(n int, edges [][]int) []int {
    g, inDegree := make([][]int, n), make([]int, n)
    for _, v := range edges {
        x, y := v[0], v[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
        inDegree[x]++
        inDegree[y]++
    }
    queue := []int{}
    for i, v := range inDegree {
        if v == 1 {
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 {
        x := queue[0]
        queue = queue[1:]
        for _, y := range g[x] {
            inDegree[y]--
            if inDegree[y] == 1 {
                queue = append(queue, y)
            }
        }
    }
    res, visited, stack := make([]int, n), make([]bool, n), []int{}
    for i, v := range inDegree {
        if v == 2 {
            visited[i] = true
            stack = append(stack, i)
        }
    }
    for i := range res {
        res[i] = 1 << 32 - 1
        if visited[i] {
            res[i] = 0
        }
    }
    for len(stack) > 0 {
        cur := stack[0]
        stack = stack[1:]
        for _, v := range g[cur] {
            if !visited[v] {
                visited[v] = true
                res[v] = res[cur] + 1
                stack = append(stack, v)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315154238-1.png" />
    // Input: n = 7, edges = [[1,2],[2,4],[4,3],[3,1],[0,1],[5,2],[6,5]]
    // Output: [1,0,0,0,0,1,2]
    // Explanation:
    // The nodes 1, 2, 3, and 4 form the cycle.
    // The distance from 0 to 1 is 1.
    // The distance from 1 to 1 is 0.
    // The distance from 2 to 2 is 0.
    // The distance from 3 to 3 is 0.
    // The distance from 4 to 4 is 0.
    // The distance from 5 to 2 is 1.
    // The distance from 6 to 2 is 2.
    fmt.Println(distanceToCycle(7,[][]int{{1,2},{2,4},{4,3},{3,1},{0,1},{5,2},{6,5}})) // [1,0,0,0,0,1,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315154634-1.png" />
    // Input: n = 9, edges = [[0,1],[1,2],[0,2],[2,6],[6,7],[6,8],[0,3],[3,4],[3,5]]
    // Output: [0,0,0,1,2,2,1,2,2]
    // Explanation:
    // The nodes 0, 1, and 2 form the cycle.
    // The distance from 0 to 0 is 0.
    // The distance from 1 to 1 is 0.
    // The distance from 2 to 2 is 0.
    // The distance from 3 to 1 is 1.
    // The distance from 4 to 1 is 2.
    // The distance from 5 to 1 is 2.
    // The distance from 6 to 2 is 1.
    // The distance from 7 to 2 is 2.
    // The distance from 8 to 2 is 2.
    fmt.Println(distanceToCycle(9,[][]int{{0,1},{1,2},{0,2},{2,6},{6,7},{6,8},{0,3},{3,4},{3,5}})) //  [0,0,0,1,2,2,1,2,2]

    fmt.Println(distanceToCycle1(7,[][]int{{1,2},{2,4},{4,3},{3,1},{0,1},{5,2},{6,5}})) // [1,0,0,0,0,1,2]
    fmt.Println(distanceToCycle1(9,[][]int{{0,1},{1,2},{0,2},{2,6},{6,7},{6,8},{0,3},{3,4},{3,5}})) //  [0,0,0,1,2,2,1,2,2]
}