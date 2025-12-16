package main

// 3778. Minimum Distance Excluding One Maximum Weighted Edge
// You are given a positive integer n and a 2D integer array edges, where edges[i] = [ui, vi, wi].

// There is a weighted connected simple undirected graph with n nodes labeled from 0 to n - 1. 
// Each [ui, vi, wi] in edges represents an edge between node ui and node vi with positive weight wi.

// The cost of a path is the sum of weights of the edges in the path, excluding the edge with the maximum weight. 
// If there are multiple edges in the path with the maximum weight, only the first such edge is excluded.

// Return an integer representing the minimum cost of a path going from node 0 to node n - 1.

// Example 1:
// Input: n = 5, edges = [[0,1,2],[1,2,7],[2,3,7],[3,4,4]]
// Output: 13
// Explanation:
// There is only one path going from node 0 to node 4: 0 -> 1 -> 2 -> 3 -> 4.
// The edge weights on this path are 2, 7, 7, and 4.
// Excluding the first edge with maximum weight, which is 1 -> 2, the cost of this path is 2 + 7 + 4 = 13.

// Example 2:
// Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,50000]]
// Output: 0
// Explanation:
// There are two paths going from node 0 to node 2:
// 0 -> 1 -> 2
// The edge weights on this path are 1 and 1.
// Excluding the first edge with maximum weight, which is 0 -> 1, the cost of this path is 1.
// 0 -> 2
// The only edge weight on this path is 1.
// Excluding the first edge with maximum weight, which is 0 -> 2, the cost of this path is 0.
// The minimum cost is min(1, 0) = 0.

// Constraints:
//     2 <= n <= 5 * 10^4
//     n - 1 <= edges.length <= 10^9
//     edges[i] = [ui, vi, wi]
//     0 <= ui < vi < n
//     [ui, vi] != [uj, vj]
//     1 <= wi <= 5 * 10^4
//     The graph is connected.

import "fmt"
import "container/heap"
import "math"
import "sort"

type Edge struct {
    to     int
    weight int
}

type State struct {
    node int
    dist int64
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func dijkstra(n int, graph [][]Edge, start int) []int64 {
    dist := make([]int64, n)
    for i := range dist {
        dist[i] = math.MaxInt64
    }
    dist[start] = 0
    pq := &PriorityQueue{}
    heap.Push(pq, State{node: start, dist: 0})
    for pq.Len() > 0 {
        curr := heap.Pop(pq).(State)
        if curr.dist != dist[curr.node] {
            continue
        }
        for _, edge := range graph[curr.node] {
            newDist := dist[curr.node] + int64(edge.weight)
            if newDist < dist[edge.to] {
                dist[edge.to] = newDist
                heap.Push(pq, State{node: edge.to, dist: newDist})
            }
        }
    }
    return dist
}

// 超出时间限制 911 / 921 
func minCostExcludingMax(n int, edges [][]int) int64 {
    // Build list of edges with their weights
    type WeightedEdge struct {
        u, v, w int
    }
    edgeList, weightSet := make([]WeightedEdge, len(edges)), make(map[int]bool)
    for i, e := range edges {
        edgeList[i] = WeightedEdge{u: e[0], v: e[1], w: e[2]}
        weightSet[e[2]] = true
    }
    // Get unique weights and sort them
    uniqueWeights := make([]int, 0, len(weightSet))
    for w := range weightSet {
        uniqueWeights = append(uniqueWeights, w)
    }
    // Sort edges by weight
    sortedEdges := make([]WeightedEdge, len(edgeList))
    copy(sortedEdges, edgeList)
    // Group edges by weight
    edgesByWeight := make(map[int][]WeightedEdge)
    for _, e := range edgeList {
        edgesByWeight[e.w] = append(edgesByWeight[e.w], e)
    }
    // Sort unique weights
    // Simple bubble sort or use sort package, but to avoid import, implement quick sort
    sort.Ints(uniqueWeights)
    //quickSort(uniqueWeights, 0, len(uniqueWeights)-1)
    // Initialize graph & Keep track of which edges we've added (by weight)
    res, graph, addedWeights := int64(math.MaxInt64), make([][]Edge, n), make(map[int]bool)
    // Process weights in increasing order
    for _, w := range uniqueWeights {
        // Add all edges with weight w to the graph
        for _, e := range edgesByWeight[w] {
            graph[e.u] = append(graph[e.u], Edge{to: e.v, weight: e.w})
            graph[e.v] = append(graph[e.v], Edge{to: e.u, weight: e.w})
        }
        addedWeights[w] = true
        // Run Dijkstra from node 0
        dist0 := dijkstra(n, graph, 0)
        // Run Dijkstra from node n-1
        distN := dijkstra(n, graph, n-1)
        // Check all edges with weight exactly w
        for _, e := range edgesByWeight[w] {
            // Path: 0 -> e.u -> e.v -> n-1
            if dist0[e.u] != math.MaxInt64 && distN[e.v] != math.MaxInt64 {
                cost := dist0[e.u] + distN[e.v]
                if cost < res {
                    res = cost
                }
            }
            // Path: 0 -> e.v -> e.u -> n-1
            if dist0[e.v] != math.MaxInt64 && distN[e.u] != math.MaxInt64 {
                cost := dist0[e.v] + distN[e.u]
                if cost < res {
                    res = cost
                }
            }
        }
    }
    return res
}

// func quickSort(arr []int, low, high int) {
//     if low < high {
//         pi := partition(arr, low, high)
//         quickSort(arr, low, pi-1)
//         quickSort(arr, pi+1, high)
//     }
// }

// func partition(arr []int, low, high int) int {
//     pivot := arr[high]
//     i := low - 1
//     for j := low; j < high; j++ {
//         if arr[j] < pivot {
//             i++
//             arr[i], arr[j] = arr[j], arr[i]
//         }
//     }
//     arr[i+1], arr[high] = arr[high], arr[i+1]
//     return i + 1
// }

func main() {
    // Example 1:
    // Input: n = 5, edges = [[0,1,2],[1,2,7],[2,3,7],[3,4,4]]
    // Output: 13
    // Explanation:
    // There is only one path going from node 0 to node 4: 0 -> 1 -> 2 -> 3 -> 4.
    // The edge weights on this path are 2, 7, 7, and 4.
    // Excluding the first edge with maximum weight, which is 1 -> 2, the cost of this path is 2 + 7 + 4 = 13.
    fmt.Println(minCostExcludingMax(5, [][]int{{0,1,2},{1,2,7},{2,3,7},{3,4,4}})) // 13
    // Example 2:
    // Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,50000]]
    // Output: 0
    // Explanation:
    // There are two paths going from node 0 to node 2:
    // 0 -> 1 -> 2
    // The edge weights on this path are 1 and 1.
    // Excluding the first edge with maximum weight, which is 0 -> 1, the cost of this path is 1.
    // 0 -> 2
    // The only edge weight on this path is 1.
    // Excluding the first edge with maximum weight, which is 0 -> 2, the cost of this path is 0.
    // The minimum cost is min(1, 0) = 0.
    fmt.Println(minCostExcludingMax(3, [][]int{{0,1,1},{1,2,1},{0,2,50000}})) // 0
}