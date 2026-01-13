package main

// 3807. Minimum Cost to Repair Edges to Traverse a Graph
// You are given an undirected graph with n nodes labeled from 0 to n - 1. 
// The graph consists of m edges represented by a 2D integer array edges, where edges[i] = [ui, vi, wi] indicates that there is an edge between nodes ui and vi with a repair cost of wi.

// You are also given an integer k. Initially, all edges are damaged.

// You may choose a non-negative integer money and repair all edges whose repair cost is less than or equal to money. 
// All other edges remain damaged and cannot be used.

// You want to travel from node 0 to node n - 1 using at most k edges.

// Return an integer denoting the minimum amount of money required to make this possible, or return -1 if it is impossible.

// Example 1:
// ​<img src="https://assets.leetcode.com/uploads/2026/01/04/ex1drawio.png" />
// Input: n = 3, edges = [[0,1,10],[1,2,10],[0,2,100]], k = 1
// Output: 100
// Explanation:
// The only valid path using at most k = 1 edge is 0 -> 2, which requires repairing the edge with cost 100. 
// Therefore, the minimum required amount of money is 100.

// Example 2:
// ​<img src="https://assets.leetcode.com/uploads/2026/01/04/ex2drawio.png" />
// Input: n = 6, edges = [[0,2,5],[2,3,6],[3,4,7],[4,5,5],[0,1,10],[1,5,12],[0,3,9],[1,2,8],[2,4,11]], k = 2
// Output: 12
// Explanation:
// With money = 12, all edges with repair cost at most 12 become usable.
// This allows the path 0 -> 1 -> 5, which uses exactly 2 edges and reaches node 5.
// If money < 12, there is no available path of length at most k = 2 from node 0 to node 5.
// Therefore, the minimum required money is 12.

// Example 3:
// ​<img src="https://assets.leetcode.com/uploads/2026/01/04/ex3drawio.png" />
// Input: n = 3, edges = [[0,1,1]], k = 1
// Output: -1
// Explanation:
// It is impossible to reach node 2 from node 0 using any amount of money. Therefore, the answer is -1.

// Constraints:
//     2 <= n <= 5 * 10^4
//     1 <= edges.length == m <= 10^5
//     edges[i] = [ui, vi, wi]
//     0 <= ui, vi < n
//     1 <= wi <= 10^9
//     1 <= k <= n
//     There are no self-loops or duplicate edges in the graph.

import "fmt"
import "sort"

func minCost(n int, edges [][]int, k int) int {
    weights, weightSet := make([]int,0), make(map[int]bool) // // 提取所有边的权重，用于二分查找的上下界
    for _, e := range edges {
        w := e[2]
        if !weightSet[w] {
            weightSet[w] = true
            weights = append(weights, w)
        }
    }
    sort.Ints(weights) // 排序权重，用于二分
    if len(weights) == 0 { return -1 } // 处理特殊情况：没有边（但防止边界）
    canReach := func (n int, edges [][]int, k int, money int) bool { // 检查给定 money 时，是否能从0出发，经过最多 k 条边到达 n-1
        adj := make([][]int, n) // 构建邻接表：只包含费用 ≤ money的边
        for _, e := range edges {
            u, v, w := e[0], e[1], e[2]
            if w <= money {
                adj[u] = append(adj[u], v)
                adj[v] = append(adj[v], u)
            }
        }
        // BFS：记录每个节点的最短步数
        visited := make([]int, n)
        for i := range visited {
            visited[i] = -1 // -1表示未访问
        }
        queue := []int{0}
        visited[0] = 0 // 起点步数为0
        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]
            if curr == n-1 { return visited[curr] <= k } // 到达终点，且步数≤k
            if visited[curr] >= k { continue } // 步数已经到k，无需继续扩展
            for _, next := range adj[curr] { // 遍历邻接节点
                if visited[next] == -1 { // 未访问过，避免环
                    visited[next] = visited[curr] + 1
                    queue = append(queue, next)
                    // 提前判断：如果下一个节点是终点且步数≤k，直接返回
                    if next == n-1 && visited[next] <= k {
                        return true
                    }
                }
            }
        }
        // 遍历完都没找到终点，或步数超过k
        return false
    }
    // 二分查找的左右指针
    res, left, right := -1, 0, len(weights) - 1
    // 二分查找核心逻辑
    for left <= right {
        mid := left + (right-left) / 2
        curr:= weights[mid]
        if canReach(n, edges, k, curr) { // 检查当前 money是否满足条件
            res, right = curr, mid - 1 // 尝试找更小的 money
        } else {
            left = mid + 1 // 需要更大的 money
        }
    }
    // 额外检查：是否存在比最大权重更大的值才能满足（比如所有边都需要）
    // 处理二分中没覆盖到的情况（比如权重数组去重后，最大权重可能不够）
    if res == -1 {
        mx := weights[len(weights) - 1]
        if canReach(n, edges, k, mx) {
            res = mx
        }
    }
    return res
}

func main() {
    // Example 1:
    // ​<img src="https://assets.leetcode.com/uploads/2026/01/04/ex1drawio.png" />
    // Input: n = 3, edges = [[0,1,10],[1,2,10],[0,2,100]], k = 1
    // Output: 100
    // Explanation:
    // The only valid path using at most k = 1 edge is 0 -> 2, which requires repairing the edge with cost 100. 
    // Therefore, the minimum required amount of money is 100.
    fmt.Println(minCost(3, [][]int{{0,1,10},{1,2,10},{0,2,100}}, 1)) // 100 
    // Example 2:
    // ​<img src="https://assets.leetcode.com/uploads/2026/01/04/ex2drawio.png" />
    // Input: n = 6, edges = [[0,2,5],[2,3,6],[3,4,7],[4,5,5],[0,1,10],[1,5,12],[0,3,9],[1,2,8],[2,4,11]], k = 2
    // Output: 12
    // Explanation:
    // With money = 12, all edges with repair cost at most 12 become usable.
    // This allows the path 0 -> 1 -> 5, which uses exactly 2 edges and reaches node 5.
    // If money < 12, there is no available path of length at most k = 2 from node 0 to node 5.
    // Therefore, the minimum required money is 12.
    fmt.Println(minCost(6, [][]int{{0,2,5},{2,3,6},{3,4,7},{4,5,5},{0,1,10},{1,5,12},{0,3,9},{1,2,8},{2,4,11}}, 2)) // 12
    // Example 3:
    // ​<img src="https://assets.leetcode.com/uploads/2026/01/04/ex3drawio.png" />
    // Input: n = 3, edges = [[0,1,1]], k = 1
    // Output: -1
    // Explanation:
    // It is impossible to reach node 2 from node 0 using any amount of money. Therefore, the answer is -1.
    fmt.Println(minCost(3, [][]int{{0,1,1}}, 1)) // -1
}