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


// 定义优先队列元素结构
type Item struct {
    dis  int64   // 起点到当前节点的距离
    del  int     // 是否删除过边（0：未删除，1：已删除）
    node int     // 当前节点编号
    // 堆内部使用的索引（container/heap 要求）
    index int
}

// 定义优先队列类型（最小堆）
type PriorityQueue1 []*Item

// 实现 heap.Interface 接口的 Len 方法
func (pq PriorityQueue1) Len() int { return len(pq) }

// 实现 heap.Interface 接口的 Less 方法（最小堆：dis 小的优先）
func (pq PriorityQueue1) Less(i, j int) bool {
	return pq[i].dis < pq[j].dis
}

// 实现 heap.Interface 接口的 Swap 方法
func (pq PriorityQueue1) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

// 实现 heap.Interface 接口的 Push 方法
func (pq *PriorityQueue1) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

// 实现 heap.Interface 接口的 Pop 方法
func (pq *PriorityQueue1) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil    // 避免内存泄漏
    item.index = -1   // 标记为已弹出
    *pq = old[0 : n-1]
    return item
}

func minCostExcludingMax1(n int, edges [][]int) int64 {
    // 构建邻接表
    graph := make([][][2]int, n)
    for _, edge := range edges {
        x := edge[0]
        y := edge[1]
        wt := edge[2]
        graph[x] = append(graph[x], [2]int{y, wt})
        graph[y] = append(graph[y], [2]int{x, wt})
    }
    // 初始化距离数组：dis[node][del] 表示到node节点、del状态（0/1）的最小距离
    inf := int64(math.MaxInt64)
    dis := make([][]int64, n)
    for i := range dis {
        dis[i] = []int64{inf, inf}
    }
    dis[0][0] = 0 // 起点0，未删除边的状态初始距离为0
    // 初始化优先队列
    pq := make(PriorityQueue1, 0)
    heap.Init(&pq)
    heap.Push(&pq, &Item{dis: 0, del: 0, node: 0})
    // Dijkstra 核心逻辑
    for pq.Len() > 0 {
        // 弹出当前距离最小的元素
        item := heap.Pop(&pq).(*Item)
        curDis := item.dis
        curDel := item.del
        curNode := item.node

        // 跳过过时的记录（已找到更优路径）
        if curDis > dis[curNode][curDel] { continue }

        // 遍历邻接节点
        for _, edge := range graph[curNode] {
            nextNode := edge[0]
            weight := int64(edge[1])
            // 情况1：当前未删除边，尝试删除当前这条边（不累计权重）
            if curDel == 0 {
                if curDis < dis[nextNode][1] {
                    dis[nextNode][1] = curDis
                    heap.Push(&pq, &Item{dis: curDis, del: 1, node: nextNode})
                }
            }
            // 情况2：不删除边，累计权重
            newDis := curDis + weight
            if newDis < dis[nextNode][curDel] {
                dis[nextNode][curDel] = newDis
                heap.Push(&pq, &Item{dis: newDis, del: curDel, node: nextNode})
            }
        }
    }
    // 返回终点（n-1）、已删除一条边状态的最小距离
    return dis[n-1][1]
}

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

    fmt.Println(minCostExcludingMax1(5, [][]int{{0,1,2},{1,2,7},{2,3,7},{3,4,4}})) // 13
    fmt.Println(minCostExcludingMax1(3, [][]int{{0,1,1},{1,2,1},{0,2,50000}})) // 0
}