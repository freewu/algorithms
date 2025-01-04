package main

// 2493. Divide Nodes Into the Maximum Number of Groups
// You are given a positive integer n representing the number of nodes in an undirected graph. 
// The nodes are labeled from 1 to n.

// You are also given a 2D integer array edges, 
// where edges[i] = [ai, bi] indicates that there is a bidirectional edge between nodes ai and bi. 
// Notice that the given graph may be disconnected.

// Divide the nodes of the graph into m groups (1-indexed) such that:
//     1. Each node in the graph belongs to exactly one group.
//     2. For every pair of nodes in the graph that are connected by an edge [ai, bi], 
//        if ai belongs to the group with index x, 
//        and bi belongs to the group with index y, then |y - x| = 1.

// Return the maximum number of groups (i.e., maximum m) into which you can divide the nodes. 
// Return -1 if it is impossible to group the nodes with the given conditions.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/10/13/example1.png" />
// Input: n = 6, edges = [[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]
// Output: 4
// Explanation: As shown in the image we:
// - Add node 5 to the first group.
// - Add node 1 to the second group.
// - Add nodes 2 and 4 to the third group.
// - Add nodes 3 and 6 to the fourth group.
// We can see that every edge is satisfied.
// It can be shown that that if we create a fifth group and move any node from the third or fourth group to it, at least on of the edges will not be satisfied.

// Example 2:
// Input: n = 3, edges = [[1,2],[2,3],[3,1]]
// Output: -1
// Explanation: If we add node 1 to the first group, node 2 to the second group, and node 3 to the third group to satisfy the first two edges, we can see that the third edge will not be satisfied.
// It can be shown that no grouping is possible.

// Constraints:
//     1 <= n <= 500
//     1 <= edges.length <= 10^4
//     edges[i].length == 2
//     1 <= ai, bi <= n
//     ai != bi
//     There is at most one edge between any pair of vertices.

import "fmt"

func magnificentSets(n int, edges [][]int) int {
    res, inf := 0,  1 << 31
    graph := make([][]int, n + 1)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    arr, visited := []int{}, make([]bool, n + 1)
    var dfs func(index int)
    dfs = func(index int) {
        arr = append(arr, index)
        visited[index] = true
        for _, i := range graph[index] {
            if !visited[i] {
                dfs(i)
            }
        }
    }
    bfs := func(k int) int {
        res, dist := 1, make([]int, n + 1)
        for i := range dist {
            dist[i] = inf
        }
        queue := []int{ k }
        dist[k] = 1
        for len(queue) > 0 {
            v := queue[0]
            queue = queue[1:]
            for _, i := range graph[v] {
                if dist[i] == inf {
                    dist[i] = dist[v] + 1
                    res = dist[i]
                    queue = append(queue, i)
                }
            }
        }
        for _, i := range arr {
            if dist[i] == inf {
                res++
                dist[i] = res
            }
        }
        for _, i := range arr {
            for _, j := range graph[i] {
                if abs(dist[i] - dist[j]) != 1 {
                    return -1
                }
            }
        }
        return res
    }
    for i := 1; i <= n; i++ {
        if !visited[i] {
            dfs(i)
            t := -1
            for _, v := range arr {
                t = max(t, bfs(v))
            }
            if t == -1 { return -1 }
            res += t
            arr = []int{}
        }
    }
    return res
}

func magnificentSets1(n int, edges [][]int) int {
    res, graph := 0, make([][]int, n)
    for _, v := range edges {
        a, b := v[0] - 1, v[1] - 1
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp := make([]int, n)
    for i := range dp {
        queue := []int{ i }
        dist := make([]int, n)
        dist[i] = 1
        mx, root := 1, i
        for len(queue) > 0 {
            v := queue[0]
            queue = queue[1:]
            root = min(root, v)
            for _, b := range graph[v] {
                if dist[b] == 0 {
                    dist[b] = dist[v] + 1
                    mx = max(mx, dist[b])
                    queue = append(queue, b)
                } else if abs(dist[b] - dist[v]) != 1 {
                    return -1
                }
            }
        }
        dp[root] = max(dp[root], mx)
    }
    for _, v := range dp {
        res += v
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/10/13/example1.png" />
    // Input: n = 6, edges = [[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]
    // Output: 4
    // Explanation: As shown in the image we:
    // - Add node 5 to the first group.
    // - Add node 1 to the second group.
    // - Add nodes 2 and 4 to the third group.
    // - Add nodes 3 and 6 to the fourth group.
    // We can see that every edge is satisfied.
    // It can be shown that that if we create a fifth group and move any node from the third or fourth group to it, at least on of the edges will not be satisfied.
    fmt.Println(magnificentSets(6, [][]int{{1,2},{1,4},{1,5},{2,6},{2,3},{4,6}})) // 4
    // Example 2:
    // Input: n = 3, edges = [[1,2],[2,3],[3,1]]
    // Output: -1
    // Explanation: If we add node 1 to the first group, node 2 to the second group, and node 3 to the third group to satisfy the first two edges, we can see that the third edge will not be satisfied.
    // It can be shown that no grouping is possible.
    fmt.Println(magnificentSets(3, [][]int{{1,2},{2,3},{3,1}})) // -1

    fmt.Println(magnificentSets1(6, [][]int{{1,2},{1,4},{1,5},{2,6},{2,3},{4,6}})) // 4
    fmt.Println(magnificentSets1(3, [][]int{{1,2},{2,3},{3,1}})) // -1
}