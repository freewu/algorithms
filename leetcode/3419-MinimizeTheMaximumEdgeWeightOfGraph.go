package main

// 3419. Minimize the Maximum Edge Weight of Graph
// You are given two integers, n and threshold, as well as a directed weighted graph of n nodes numbered from 0 to n - 1. 
// The graph is represented by a 2D integer array edges, where edges[i] = [Ai, Bi, Wi] indicates that there is an edge going from node Ai to node Bi with weight Wi.

// You have to remove some edges from this graph (possibly none), so that it satisfies the following conditions:
//     Node 0 must be reachable from all other nodes.
//     The maximum edge weight in the resulting graph is minimized.
//     Each node has at most threshold outgoing edges.

// Return the minimum possible value of the maximum edge weight after removing the necessary edges. 
// If it is impossible for all conditions to be satisfied, return -1.

// Example 1:
// Input: n = 5, edges = [[1,0,1],[2,0,2],[3,0,1],[4,3,1],[2,1,1]], threshold = 2
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/09/s-1.png" />
// Remove the edge 2 -> 0. The maximum weight among the remaining edges is 1.

// Example 2:
// Input: n = 5, edges = [[0,1,1],[0,2,2],[0,3,1],[0,4,1],[1,2,1],[1,4,1]], threshold = 1
// Output: -1
// Explanation: 
// It is impossible to reach node 0 from node 2.

// Example 3:
// Input: n = 5, edges = [[1,2,1],[1,3,3],[1,4,5],[2,3,2],[3,4,2],[4,0,1]], threshold = 1
// Output: 2
// Explanation: 
// <img src="https://assets.leetcode.com/uploads/2024/12/09/s2-1.png" />
// Remove the edges 1 -> 3 and 1 -> 4. The maximum weight among the remaining edges is 2.

// Example 4:
// Input: n = 5, edges = [[1,2,1],[1,3,3],[1,4,5],[2,3,2],[4,0,1]], threshold = 1
// Output: -1

// Constraints:
//     2 <= n <= 10^5
//     1 <= threshold <= n - 1
//     1 <= edges.length <= min(10^5, n * (n - 1) / 2).
//     edges[i].length == 3
//     0 <= Ai, Bi < n
//     Ai != Bi
//     1 <= Wi <= 10^6
//     There may be multiple edges between a pair of nodes, but they must have unique weights.

import "fmt"
import "container/heap"
import "slices"

// type MinHeap [][2]int
// func (h MinHeap) Len() int            { return len(h) }
// func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
// func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
// func (h *MinHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[0 : n-1]
//     return x
// }

// func minMaxWeight(n int, edges [][]int, threshold int) int {
//     // Build adjacency list
//     graph := make([]map[int]int, n)
//     for i := range graph {
//         graph[i] = make(map[int]int)
//     }
//     for _, edge := range edges {
//         from, to, weight := edge[0], edge[1], edge[2]
//         if w, ok := graph[to][from]; !ok || weight < w {
//             graph[to][from] = weight
//         }
//     }
//     // Initialize Dijkstra's algorithm
//     h := &MinHeap{}
//     heap.Init(h)
//     heap.Push(h, [2]int{0, 0})
//     seen := make([]int, n)
//     for i := range seen {
//         seen[i] = 1 << 31
//     }
//     k := n
//     // Run Dijkstra's algorithm
//     for h.Len() > 0 && k > 0 {
//         item := heap.Pop(h).([2]int)
//         d, i := item[0], item[1]
//         if seen[i] < 1 << 31 { continue }
//         k--
//         seen[i] = d
//         for j, w := range graph[i] {
//             if seen[j] < 1 << 31 { continue }
//             heap.Push(h, [2]int{w, j})
//         }
//     }
//     // Check if all nodes are reachable
//     if k > 0 { return -1 }
//     // Find the maximum weight in the seen array
//     res := 0
//     for _, weight := range seen {
//         if weight > res {
//             res = weight
//         }
//     }
//     return res
// }

type Pair struct{ dis, x int } // 路径最大边权, 节点编号
type MinHeap []Pair
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minMaxWeight(n int, edges [][]int, _ int) int {
    if len(edges) < n - 1 { return -1 }
    type Edge struct{ to, w int }
    graph := make([][]Edge, n)
    for _, e := range edges {
        x, y, w := e[0], e[1], e[2]
        graph[y] = append(graph[y], Edge{ x, w })
    }
    dis := make([]int, n)
    for i := range dis {
        dis[i] = 1 << 31
    }
    dis[0] = 0
    h := MinHeap{{}}
    for len(h) > 0 {
        p := heap.Pop(&h).(Pair)
        x := p.x
        d := p.dis
        for _, e := range graph[x] {
            y := e.to
            nd := max(d, e.w)
            if nd < dis[y] {
                dis[y] = nd
                heap.Push(&h, Pair{nd, y})
            }
        }
    }
    res := slices.Max(dis)
    if res == 1 << 31 { return -1 }
    return res
}

func minMaxWeight1(n int, edges [][]int, threshold int) int {
    // Since threshold >= 1, we can always choose at most one outgoing edge per node
    // (thus outdegree <= 1), if connectivity allows. Therefore the main question is:
    // "What is the smallest W such that for every node i, there is a path i->0 using edges of weight <= W?"
    //
    // Equivalently, in the reversed graph (reverse every edge), we want:
    // "What is the smallest W such that node 0 can reach all other nodes in the reversed graph,
    // using only edges of weight <= W?"
    //
    // Approach:
    // 1. Reverse all edges: for an original edge u->v, in the reversed graph it becomes v->u.
    // 2. Sort these reversed edges by weight ascending.
    // 3. Maintain a BFS/queue from node 0 in the reversed graph. Initially visited[0] = true.
    // 4. Iterate over each distinct weight in ascending order:
    //      - Add all reversed edges of this weight into the adjacency list.
    //      - For each newly added edge (f->t), if f is visited and t is not, visit t and push it into the queue, 
    //        then continue BFS expansions from t.
    //      - If at any point all n nodes become visited, return the current weight.
    // 5. If after processing all edges we do not visit all nodes, return -1.
    if threshold < 1 {
        // Problem constraints guarantee threshold >= 1, but just a safety check.
        return -1
    }
    // Build reversed edges: (weight, from, to) where from->to in reversed graph
    revEdges := make([][3]int, len(edges))
    for i, e := range edges {
        revEdges[i] = [3]int{e[2], e[1], e[0]} // (weight, B, A) if original was A->B
    }
    // Sort by weight ascending
    // revEdges[i] = [weight, from, to]
    // from->to in the reversed graph
    // so node 'from' has an edge to 'to'
    // We want ascending order by revEdges[i][0].
    quickSort(revEdges, 0, len(revEdges)-1)
    // Prepare adjacency (reversed graph), visited array, and a queue
    adj := make([][]int, n)
    visited := make([]bool, n)
    visited[0] = true
    visitedCount := 1
    queue := make([]int, 0, n)
    queue = append(queue, 0)
    // A helper function to run BFS expansion whenever we add new edges
    // If from is visited and to not visited, we mark it visited and enqueue.
    bfsExpand := func() {
        for len(queue) > 0 {
            front := queue[len(queue)-1]
            queue = queue[:len(queue)-1]
            for _, nxt := range adj[front] {
                if !visited[nxt] {
                    visited[nxt] = true
                    visitedCount++
                    queue = append(queue, nxt)
                }
            }
        }
    }
    // Process edges by their weight in ascending order. We'll group edges of the same weight.
    idx := 0
    for idx < len(revEdges) {
        w := revEdges[idx][0]
        // Gather all edges with weight = w
        start := idx
        for idx < len(revEdges) && revEdges[idx][0] == w {
            idx++
        }
        // Add these edges in adjacency
        for i := start; i < idx; i++ {
            from := revEdges[i][1]
            to := revEdges[i][2]
            adj[from] = append(adj[from], to)
            // If 'from' is already visited, this new edge might let us visit 'to'
            if visited[from] && !visited[to] {
                visited[to] = true
                visitedCount++
                queue = append(queue, to)
            }
        }
        // Do BFS expansions
        bfsExpand()
        // Check if all nodes visited
        if visitedCount == n {
            return w
        }
    }
    // If we exhaust all edges and not all are visited
    return -1
}

// A simple in-place quicksort for [3]int by the 0th element (weight)
func quickSort(arr [][3]int, left, right int) {
    if left >= right {
        return
    }
    pivot := arr[(left+right)>>1][0]
    i, j := left, right
    for i <= j {
        for arr[i][0] < pivot {
            i++
        }
        for arr[j][0] > pivot {
            j--
        }
        if i <= j {
            arr[i], arr[j] = arr[j], arr[i]
            i++
            j--
        }
    }
    quickSort(arr, left, j)
    quickSort(arr, i, right)
}

func main() {
    // Example 1:
    // Input: n = 5, edges = [[1,0,1],[2,0,2],[3,0,1],[4,3,1],[2,1,1]], threshold = 2
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/09/s-1.png" />
    // Remove the edge 2 -> 0. The maximum weight among the remaining edges is 1.
    fmt.Println(minMaxWeight(5, [][]int{{1,0,1},{2,0,2},{3,0,1},{4,3,1},{2,1,1}}, 2)) // 1
    // Example 2:
    // Input: n = 5, edges = [[0,1,1],[0,2,2],[0,3,1],[0,4,1],[1,2,1],[1,4,1]], threshold = 1
    // Output: -1
    // Explanation: 
    // It is impossible to reach node 0 from node 2.
    fmt.Println(minMaxWeight(5, [][]int{{0,1,1},{0,2,2},{0,3,1},{0,4,1},{1,2,1},{1,4,1}}, 1)) // -1
    // Example 3:
    // Input: n = 5, edges = [[1,2,1],[1,3,3],[1,4,5],[2,3,2],[3,4,2],[4,0,1]], threshold = 1
    // Output: 2
    // Explanation: 
    // <img src="https://assets.leetcode.com/uploads/2024/12/09/s2-1.png" />
    // Remove the edges 1 -> 3 and 1 -> 4. The maximum weight among the remaining edges is 2.
    fmt.Println(minMaxWeight(5, [][]int{{1,2,1},{1,3,3},{1,4,5},{2,3,2},{3,4,2},{4,0,1}}, 1)) // 2
    // Example 4:
    // Input: n = 5, edges = [[1,2,1],[1,3,3],[1,4,5],[2,3,2],[4,0,1]], threshold = 1
    // Output: -1
    fmt.Println(minMaxWeight(5, [][]int{{1,2,1},{1,3,3},{1,4,5},{2,3,2},{4,0,1}}, 1)) // -1

    fmt.Println(minMaxWeight1(5, [][]int{{1,0,1},{2,0,2},{3,0,1},{4,3,1},{2,1,1}}, 2)) // 1
    fmt.Println(minMaxWeight1(5, [][]int{{0,1,1},{0,2,2},{0,3,1},{0,4,1},{1,2,1},{1,4,1}}, 1)) // -1
    fmt.Println(minMaxWeight1(5, [][]int{{1,2,1},{1,3,3},{1,4,5},{2,3,2},{3,4,2},{4,0,1}}, 1)) // 2
    fmt.Println(minMaxWeight1(5, [][]int{{1,2,1},{1,3,3},{1,4,5},{2,3,2},{4,0,1}}, 1)) // -1
}

