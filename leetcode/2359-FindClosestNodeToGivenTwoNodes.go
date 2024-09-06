package main

// 2359. Find Closest Node to Given Two Nodes
// You are given a directed graph of n nodes numbered from 0 to n - 1, 
// where each node has at most one outgoing edge.

// The graph is represented with a given 0-indexed array edges of size n, 
// indicating that there is a directed edge from node i to node edges[i]. 
// If there is no outgoing edge from i, then edges[i] == -1.

// You are also given two integers node1 and node2.

// Return the index of the node that can be reached from both node1 and node2, 
// such that the maximum between the distance from node1 to that node, and from node2 to that node is minimized. 
// If there are multiple answers, return the node with the smallest index, 
// and if no possible answer exists, return -1.

// Note that edges may contain cycles.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/06/07/graph4drawio-2.png" />
// Input: edges = [2,2,3,-1], node1 = 0, node2 = 1
// Output: 2
// Explanation: The distance from node 0 to node 2 is 1, and the distance from node 1 to node 2 is 1.
// The maximum of those two distances is 1. It can be proven that we cannot get a node with a smaller maximum distance than 1, so we return node 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/06/07/graph4drawio-4.png" />
// Input: edges = [1,2,-1], node1 = 0, node2 = 2
// Output: 2
// Explanation: The distance from node 0 to node 2 is 2, and the distance from node 2 to itself is 0.
// The maximum of those two distances is 2. It can be proven that we cannot get a node with a smaller maximum distance than 2, so we return node 2.

// Constraints:
//     n == edges.length
//     2 <= n <= 10^5
//     -1 <= edges[i] < n
//     edges[i] != i
//     0 <= node1, node2 < n

import "fmt"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
    n, res, minDist := len(edges), -1, 1 << 31
    dist1, dist2 := make([]int, n), make([]int, n)
    vis1, vis2 := make(map[int]bool), make(map[int]bool)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    hasVisited := func (node int, vis map[int]bool) bool {
        if _, ok := vis[node]; ok { return true }
        return false
    }
    var dfs func(edges []int, node int, dist *[]int, vis map[int]bool)
    dfs = func(edges []int, node int, dist *[]int, vis map[int]bool) {
        vis[node] = true
        neigh := edges[node]
        if neigh != -1 {
            if !hasVisited(neigh, vis) {
                (*dist)[neigh] = (*dist)[node] + 1
                dfs(edges, neigh, dist, vis)
            }
        }
    }
    dfs(edges, node1, &dist1, vis1)
    dfs(edges, node2, &dist2, vis2)
    for i := 0; i < n; i++ {
        if hasVisited(i, vis1) && hasVisited(i, vis2) {
            if minDist > max(dist1[i], dist2[i]) {
                minDist = max(dist1[i], dist2[i])
                res = i
            }
        }
    }
    return res
}

func closestMeetingNode1(edges []int, node1 int, node2 int) int {
    n := len(edges)
    if node1 > node2 { node1, node2 = node2, node1 }

    queue1, queue2 := make([]int, 0, n), make([]int, 0, n)
    visit1, visit2 := make([]bool, n), make([]bool, n)
    push := func(que *[]int, visit []bool, node int) {
        if node < 0 || visit[node] {
            return
        }
        visit[node] = true
        *que = append(*que, node)

    }
    pop := func(que *[]int) int {
        tmp := (*que)[0]
        *que = (*que)[1:]
        return tmp
    }
    push(&queue1, visit1, node1)
    push(&queue2, visit2, node2)
    res, minDist, d := n, n, 0
    for len(queue1) > 0 || len(queue2) > 0 {
        d++
        l1, l2 := len(queue1), len(queue2)
        for j := 0; j < l2; j++ { // 大的号先出队列
            node := pop(&queue2)
            if visit1[node] && (minDist > d || (minDist == d && node < res)) {
                res = node
                minDist = d
            }
            push(&queue2, visit2, edges[node])
        }
        for i := 0; i < l1; i++ {
            node := pop(&queue1)
            if visit2[node] && (minDist > d || (minDist == d && node < res)) {
                res = node
                minDist = d
            }
            push(&queue1, visit1, edges[node])
        }
        if minDist < n {
            break
        }
    }
    if res == len(edges) {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/06/07/graph4drawio-2.png" />
    // Input: edges = [2,2,3,-1], node1 = 0, node2 = 1
    // Output: 2
    // Explanation: The distance from node 0 to node 2 is 1, and the distance from node 1 to node 2 is 1.
    // The maximum of those two distances is 1. It can be proven that we cannot get a node with a smaller maximum distance than 1, so we return node 2.
    fmt.Println(closestMeetingNode([]int{2,2,3,-1}, 0, 1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/06/07/graph4drawio-4.png" />
    // Input: edges = [1,2,-1], node1 = 0, node2 = 2
    // Output: 2
    // Explanation: The distance from node 0 to node 2 is 2, and the distance from node 2 to itself is 0.
    // The maximum of those two distances is 2. It can be proven that we cannot get a node with a smaller maximum distance than 2, so we return node 2.
    fmt.Println(closestMeetingNode([]int{1,2,-1}, 0, 2)) // 2

    fmt.Println(closestMeetingNode1([]int{2,2,3,-1}, 0, 1)) // 2
    fmt.Println(closestMeetingNode1([]int{1,2,-1}, 0, 2)) // 2
}