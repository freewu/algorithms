package main

// 2473. Minimum Cost to Buy Apples
// You are given a positive integer n representing n cities numbered from 1 to n. 
// You are also given a 2D array roads, where roads[i] = [ai, bi, costi] indicates 
// that there is a bidirectional road between cities ai and bi with a cost of traveling equal to costi.

// You can buy apples in any city you want, but some cities have different costs to buy apples. 
// You are given the 1-based array appleCost where appleCost[i] is the cost of buying one apple from city i.

// You start at some city, traverse through various roads, and eventually buy exactly one apple from any city. 
// After you buy that apple, you have to return back to the city you started at, 
// but now the cost of all the roads will be multiplied by a given factor k.

// Given the integer k, return a 1-based array answer of size n where answer[i] is the minimum total cost to buy an apple if you start at city i.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/15/graph55.png" />
// Input: n = 4, roads = [[1,2,4],[2,3,2],[2,4,5],[3,4,1],[1,3,4]], appleCost = [56,42,102,301], k = 2
// Output: [54,42,48,51]
// Explanation: The minimum cost for each starting city is the following:
// - Starting at city 1: You take the path 1 -> 2, buy an apple at city 2, and finally take the path 2 -> 1. The total cost is 4 + 42 + 4 * 2 = 54.
// - Starting at city 2: You directly buy an apple at city 2. The total cost is 42.
// - Starting at city 3: You take the path 3 -> 2, buy an apple at city 2, and finally take the path 2 -> 3. The total cost is 2 + 42 + 2 * 2 = 48.
// - Starting at city 4: You take the path 4 -> 3 -> 2 then you buy at city 2, and finally take the path 2 -> 3 -> 4. The total cost is 1 + 2 + 42 + 1 * 2 + 2 * 2 = 51.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/15/graph4.png" />
// Input: n = 3, roads = [[1,2,5],[2,3,1],[3,1,2]], appleCost = [2,3,1], k = 3
// Output: [2,3,1]
// Explanation: It is always optimal to buy the apple in the starting city.

// Constraints:
//     2 <= n <= 1000
//     1 <= roads.length <= 2000
//     1 <= ai, bi <= n
//     ai != bi
//     1 <= costi <= 10^5
//     appleCost.length == n
//     1 <= appleCost[i] <= 105
//     1 <= k <= 100
//     There are no repeated edges.

import "fmt"
import "container/heap"

// type Edge struct{ to, value int }
// type hp []Edge

// func (h hp) Len() int           { return len(h) }
// func (h hp) Less(i, j int) bool { return h[i].value < h[j].value }
// func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(v any)        { *h = append(*h, v.(Edge)) }
// func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

// func minCost(n int, roads [][]int, appleCost []int, k int) []int64 {
//     res, graph, inf := make([]int64, n), make([][]Edge, n + 1), 1 << 31
//     for i := 1; i <= n; i++ {
//         graph[i] = []Edge{}
//     }
//     for i := range res { res[i] = int64(inf) } // fill res
//     for i := 0; i < len(roads); i++ {
//         p := roads[i][2] * (k + 1)
//         graph[roads[i][0]] = append(graph[roads[i][0]],Edge{roads[i][1], p})
//         graph[roads[i][1]] = append(graph[roads[i][0]],Edge{roads[i][0], p})
//     }
//     dist := make([]int64, n + 1)
//     dijkstra := func(dist []int64, n int, start int) {
//         for i := range dist { dist[i] = 1 << 31 } 
//         visited := make([]bool, n + 1)
//         queue := &hp{Edge{start, 0}}
//         heap.Init(queue)
//         dist[start] = 0
//         for queue.Len() > 0 {
//             edge := heap.Pop(queue).(Edge)
//             if visited[edge.to] { continue } // 已范访问过
//             visited[edge.to] = true
//             for i := range graph[edge.to] {
//                 end, value := graph[edge.to][i].to, graph[edge.to][i].value
//                 if dist[end] > dist[edge.to] + int64(value) {
//                     dist[end] = dist[edge.to] + int64(value)
//                     queue.Push(Edge{ end, int(dist[end]) })
//                 }
//             }
//         }
//     }
//     for i := 1; i <= n; i++ {
//         dijkstra(dist, n, i)
//         for to := 1; to <= n; to++ {
//             if res[i-1] > dist[to] + int64(appleCost[to-1]) {
//                 res[i-1] = dist[to] + int64(appleCost[to-1])
//             }
//         }
//     }
//     return res
// }

type pair struct{ cost, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minCost(n int, roads [][]int, appleCost []int, k int) []int64 {
    g := make([][][2]int, n + 1)
    for i := range g {
        g[i] = make([][2]int, 0)
    }
    for _, v := range roads {
        x, y, z := v[0], v[1], v[2]
        g[x] = append(g[x], [2]int{y, z * (k + 1)})
        g[y] = append(g[y], [2]int{x, z * (k + 1)})
    }
    mh := &hp{}
    heap.Init(mh)
    for i := 0; i < n; i++ {
        heap.Push(mh, pair{appleCost[i], i + 1})
    }
    res := make([]int64, n + 1)
    for mh.Len() > 0 {
        v := heap.Pop(mh).(pair)
        nowd, now := v.cost, v.i
        if res[now] > 0 { continue }
        res[now] = int64(nowd)
        for _, v1 := range g[now] {
            next, a := v1[0], v1[1]
            if res[next] == 0 {
                heap.Push(mh, pair{nowd + a, next})
            }
        }
    }
    return res[1:]
}

// class Solution(object):
//     def minCost(self, n, roads, appleCost, k):
//         """
//         :type n: int
//         :type roads: List[List[int]]
//         :type appleCost: List[int]
//         :type k: int
//         :rtype: List[int]
//         """

//         g = [[] for _ in range(n + 1)]
//         for x, y, v in roads:
//             g[x].append([y, v * (k + 1)])
//             g[y].append([x, v * (k + 1)])
        
//         heap = []
//         for index in range(n):
//             heapq.heappush(heap, [appleCost[index], index + 1])
        
//         dis = [None] * (n + 1)
//         while heap:
//             nowd, now = heapq.heappop(heap)
//             if dis[now] is not None:
//                 continue
            
//             dis[now] = nowd
//             for next, v in g[now]:
//                 if dis[next] is None:
//                     heapq.heappush(heap, [nowd + v, next])
        
//         return dis[1:]

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/15/graph55.png" />
    // Input: n = 4, roads = [[1,2,4],[2,3,2],[2,4,5],[3,4,1],[1,3,4]], appleCost = [56,42,102,301], k = 2
    // Output: [54,42,48,51]
    // Explanation: The minimum cost for each starting city is the following:
    // - Starting at city 1: You take the path 1 -> 2, buy an apple at city 2, and finally take the path 2 -> 1. The total cost is 4 + 42 + 4 * 2 = 54.
    // - Starting at city 2: You directly buy an apple at city 2. The total cost is 42.
    // - Starting at city 3: You take the path 3 -> 2, buy an apple at city 2, and finally take the path 2 -> 3. The total cost is 2 + 42 + 2 * 2 = 48.
    // - Starting at city 4: You take the path 4 -> 3 -> 2 then you buy at city 2, and finally take the path 2 -> 3 -> 4. The total cost is 1 + 2 + 42 + 1 * 2 + 2 * 2 = 51.
    fmt.Println(minCost(4,[][]int{{1,2,4},{2,3,2},{2,4,5},{3,4,1},{1,3,4}}, []int{56,42,102,301}, 2)) // [54,42,48,51]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/15/graph4.png" />
    // Input: n = 3, roads = [[1,2,5],[2,3,1],[3,1,2]], appleCost = [2,3,1], k = 3
    // Output: [2,3,1]
    // Explanation: It is always optimal to buy the apple in the starting city.
    fmt.Println(minCost(3,[][]int{{1,2,5},{2,3,1},{3,1,2}}, []int{2,3,1}, 3)) // [2,3,1]
}