package main

// 1135. Connecting Cities With Minimum Cost
// There are n cities labeled from 1 to n. 
// You are given the integer n and an array connections where connections[i] = [xi, yi, costi] indicates 
// that the cost of connecting city xi and city yi (bidirectional connection) is costi.

// Return the minimum cost to connect all the n cities such that there is at least one path between each pair of cities. 
// If it is impossible to connect all the n cities, return -1,

// The cost is the sum of the connections' costs used.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/04/20/1314_ex2.png" />
// Input: n = 3, connections = [[1,2,5],[1,3,6],[2,3,1]]
// Output: 6
// Explanation: Choosing any 2 edges will connect all cities so we choose the minimum 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/04/20/1314_ex1.png" />
// Input: n = 4, connections = [[1,2,3],[3,4,4]]
// Output: -1
// Explanation: There is no way to connect all cities even if all edges are used.
 
// Constraints:
//     1 <= n <= 10^4
//     1 <= connections.length <= 10^4
//     connections[i].length == 3
//     1 <= xi, yi <= n
//     xi != yi
//     0 <= costi <= 10^5

import "fmt"
import "sort"
import "container/heap"

// # Kruskal 算法
// 1 将所有的边按照权重从小到大排序。
// 2 取一条权重最小的边。
// 3 使用并查集（union-find）数据结构来判断加入这条边后是否会形成环。若不会构成环，则将这条边加入最小生成树中。
// 4 检查所有的结点是否已经全部联通，这一点可以通过目前已经加入的边的数量来判断。若全部联通，则结束算法；否则返回步骤 2.
func minimumCost(n int, connections [][]int) int {
    id := make([]int, n)
    for i := 0; i < n; i++ { id[i] = i } // 初始化并查集

    var find func (x int, id []int) int
    find = func (x int, id []int) int {
        if x == id[x] { return x }
        return find(id[x], id)
    }
    union := func (i, j int, id []int) {
        x, y := find(i, id), find(j, id)
        if x == y { return }
        id[x] = y
    }
    sort.Slice(connections, func(i, j int) bool { // 将所有的边按照权重从小到大排序
        return connections[i][2] < connections[j][2]
    })
    count, cost := 0, 0
    for _, connect := range connections {
        if count == n - 1 { // 如果已经有 n - 1 条边，说明说有点的点都已经联通
            break
        }
        if find(connect[0] - 1, id) == find(connect[1] - 1, id) { // 会形成环，不需要加入
            continue
        }
        union(connect[0] - 1, connect[1] - 1, id) // 关联两个点，并加入到最小生成树中
        count++
        cost += connect[2]
    }
    if count != n - 1 { // 无法联通所有
        return -1
    }
    return cost
}

// # Prim 算法
// 1. 根据 connections 记录每个顶点到其他顶点的权重，记为 edges 。
// 2. 使用 visited 记录所有被访问过的点。
// 3. 使用堆来根据权重比较所有的边。
// 4. 将任意一个点记为已访问，并将其所有连接的边放入堆中。
// 5. 从堆中拿出权重最小的边。
// 6. 如果已经访问过，直接丢弃。
// 7. 如果未访问过，标记为已访问，并且将其所有连接的边放入堆中，检查是否有 n 个点。
// 8. 重复操作 5
// 为了方便存储，将 1 - n 改成 0 - n-1
func minimumCost1(n int, connections [][]int) int {
    cost, count := 0, 1
    edges, visited, h := make([][]Edge, n), make([]bool, n), &EdgeHeap{}
    heap.Init(h)
    visited[0] = true // 从第 0 个开始
    for _, connect := range connections {
        edges[connect[0] - 1] = append(edges[connect[0] - 1], Edge{connect[1] - 1, connect[2]})
        edges[connect[1] - 1] = append(edges[connect[1] - 1], Edge{connect[0] - 1, connect[2]})
    }
    // 将与第一个城市连接的城市放入 heap 中
    for _, edge := range edges[0] {
        heap.Push(h,edge)
    }
    for h.Len() > 0 {    
        e := heap.Pop(h).(Edge)
        if visited[e.city] {
            continue
        }
        visited[e.city] = true
        for _, edge := range edges[e.city] {
            heap.Push(h, edge)
        }
        cost += e.cost
        count++
        if count == n {
            return cost
        }
    }
    return -1
}

type Edge struct {
    city int
    cost int
}

type EdgeHeap []Edge
func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *EdgeHeap) Push(x interface{}) {
    *h = append(*h, x.(Edge))
}
func (h *EdgeHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/04/20/1314_ex2.png" />
    // Input: n = 3, connections = [[1,2,5],[1,3,6],[2,3,1]]
    // Output: 6
    // Explanation: Choosing any 2 edges will connect all cities so we choose the minimum 2.
    fmt.Println(minimumCost(3,[][]int{{1,2,5},{1,3,6},{2,3,1}})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/04/20/1314_ex1.png" />
    // Input: n = 4, connections = [[1,2,3],[3,4,4]]
    // Output: -1
    // Explanation: There is no way to connect all cities even if all edges are used.
    fmt.Println(minimumCost(4,[][]int{{1,2,3},{3,4,4}})) // -1

    fmt.Println(minimumCost1(3,[][]int{{1,2,5},{1,3,6},{2,3,1}})) // 6
    fmt.Println(minimumCost1(4,[][]int{{1,2,3},{3,4,4}})) // -1
}