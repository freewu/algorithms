package main

// 3620. Network Recovery Pathways
// You are given a directed acyclic graph of n nodes numbered from 0 to n − 1. 
// This is represented by a 2D array edges of length m, 
// where edges[i] = [ui, vi, costi] indicates a one‑way communication from node ui to node vi with a recovery cost of costi.

// Some nodes may be offline. 
// You are given a boolean array online where online[i] = true means node i is online. 
// Nodes 0 and n − 1 are always online.

// A path from 0 to n − 1 is valid if:
//     1. All intermediate nodes on the path are online.
//     2. The total recovery cost of all edges on the path does not exceed k.

// For each valid path, define its score as the minimum edge‑cost along that path.

// Return the maximum path score (i.e., the largest minimum-edge cost) among all valid paths. 
// If no valid path exists, return -1.

// Example 1:
// Input: edges = [[0,1,5],[1,3,10],[0,2,3],[2,3,4]], online = [true,true,true,true], k = 10
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/06/06/graph-10.png" />
// The graph has two possible routes from node 0 to node 3:
// Path 0 → 1 → 3
// Total cost = 5 + 10 = 15, which exceeds k (15 > 10), so this path is invalid.
// Path 0 → 2 → 3
// Total cost = 3 + 4 = 7 <= k, so this path is valid.
// The minimum edge‐cost along this path is min(3, 4) = 3.
// There are no other valid paths. Hence, the maximum among all valid path‐scores is 3.

// Example 2:
// Input: edges = [[0,1,7],[1,4,5],[0,2,6],[2,3,6],[3,4,2],[2,4,6]], online = [true,true,true,false,true], k = 12
// Output: 6
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/06/06/graph-11.png" />
// Node 3 is offline, so any path passing through 3 is invalid.
// Consider the remaining routes from 0 to 4:
// Path 0 → 1 → 4
// Total cost = 7 + 5 = 12 <= k, so this path is valid.
// The minimum edge‐cost along this path is min(7, 5) = 5.
// Path 0 → 2 → 3 → 4
// Node 3 is offline, so this path is invalid regardless of cost.
// Path 0 → 2 → 4
// Total cost = 6 + 6 = 12 <= k, so this path is valid.
// The minimum edge‐cost along this path is min(6, 6) = 6.
// Among the two valid paths, their scores are 5 and 6. Therefore, the answer is 6.

// Constraints:
//     n == online.length
//     2 <= n <= 5 * 10^4
//     0 <= m == edges.length <= min(10^5, n * (n - 1) / 2)
//     edges[i] = [ui, vi, costi]
//     0 <= ui, vi < n
//     ui != vi
//     0 <= costi <= 10^9
//     0 <= k <= 5 * 10^13
//     online[i] is either true or false, and both online[0] and online[n − 1] are true.
//     The given graph is a directed acyclic graph.

import "fmt"
import "sort"
import "slices"
import "container/heap"

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
    type Edge struct{ to, weight int }
    mx, n := 0, len(online)
    g, deg := make([][]Edge, n), make([]int, n)
    for _, e := range edges {
        x, y, weight := e[0], e[1], e[2]
        if online[x] && online[y] {
            g[x] = append(g[x], Edge{y, weight})
            deg[y]++
            mx = max(mx, weight)
        }
    }
    // 先清理无法从 0 到达的边
    q := []int{}
    for i := 1; i < n; i++ {
        if deg[i] == 0 {
            q = append(q, i)
        }
    }
    for len(q) > 0 {
        v := q[0]
        q = q[1:]
        for _, e := range g[v] {
            y := e.to
            deg[y]--
            if deg[y] == 0 && y > 0 {
                q = append(q, y)
            }
        }
    }
    f := make([]int, n)
    return sort.Search(mx + 1, func(lower int) bool {
        deg := slices.Clone(deg)
        for i := 1; i < n; i++ {
            f[i] = 1 << 61
        }
        q := []int{0}
        for len(q) > 0 {
            x := q[0]
            if x == n-1 {
                return f[x] > int(k)
            }
            q = q[1:]
            for _, e := range g[x] {
                y := e.to
                weight := e.weight
                if weight >= lower {
                    f[y] = min(f[y], f[x] + weight)
                }
                deg[y]--
                if deg[y] == 0 {
                    q = append(q, y)
                }
            }
        }
        return true
    }) - 1
}

func findMaxPathScore1(edges [][]int, online []bool, k int64) int {
    n, mx := len(online), 0
    type Edge struct{ to, weight int }
    g, memo := make([][]Edge, n), make([]int, n)
    for _, e := range edges {
        x, y, weight := e[0], e[1], e[2]
        if online[x] && online[y] {
            g[x] = append(g[x], Edge{y, weight})
            mx = max(mx, weight)
        }
    }
    // 二分无法到达 n-1 的最小 lower，那么减一后，就是可以到达 n-1 的最大 lower
    return sort.Search(mx + 1, func(lower int) bool {
        for i := range memo {
            memo[i] = -1 // -1 表示没有计算过
        }
        var dfs func(int) int
        dfs = func(x int) int {
            if x == n-1 { return 0 } // 到达终点
            p := &memo[x]
            if *p != -1 { return *p } // 之前计算过
            res := 1 << 61 // 防止加法溢出
            for _, e := range g[x] {
                y := e.to
                if e.weight >= lower {
                    res = min(res, dfs(y) + e.weight)
                }
            }
            *p = res // 记忆化
            return res
        }
        return dfs(0) > int(k)
    }) - 1
}

// 超出内存限制 631 / 636 
func findMaxPathScore2(edges [][]int, online []bool, k int64) int {
    n := len(online)
    g := make([][]Pair, n)
    for _, e := range edges {
        u, v, c := e[0], e[1], e[2]
        if online[u] && online[v] {
            g[u] = append(g[u], Pair{ v, c, 0})
        }
    }
    hp := new(MinHeap)
    heap.Push(hp, Pair{0, 1 << 31, 0})
    for hp.Len() > 0 {
        p := heap.Pop(hp).(Pair)
        if p.to == n - 1 { return p.c }
        for _, u := range g[p.to] {
            tall := p.all + u.c
            if int64(tall) > k { continue }
            mn := min(u.c, p.c)
            heap.Push(hp, Pair{ u.to, mn, tall})
        }
    }
    return -1
}

type Pair struct{ to, c, all int }

type MinHeap []Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].c > h[j].c }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() interface{} {
    o := (*h)[len(*h) - 1]
    *h = (*h)[:len(*h) - 1]
    return o
}

func main() {
    // Example 1:
    // Input: edges = [[0,1,5],[1,3,10],[0,2,3],[2,3,4]], online = [true,true,true,true], k = 10
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/06/06/graph-10.png" />
    // The graph has two possible routes from node 0 to node 3:
    // Path 0 → 1 → 3
    // Total cost = 5 + 10 = 15, which exceeds k (15 > 10), so this path is invalid.
    // Path 0 → 2 → 3
    // Total cost = 3 + 4 = 7 <= k, so this path is valid.
    // The minimum edge‐cost along this path is min(3, 4) = 3.
    // There are no other valid paths. Hence, the maximum among all valid path‐scores is 3.
    fmt.Println(findMaxPathScore([][]int{{0,1,5},{1,3,10},{0,2,3},{2,3,4}},[]bool{true,true,true,true}, 10)) // 3
    // Example 2:
    // Input: edges = [[0,1,7],[1,4,5],[0,2,6],[2,3,6],[3,4,2],[2,4,6]], online = [true,true,true,false,true], k = 12
    // Output: 6
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/06/06/graph-11.png" />
    // Node 3 is offline, so any path passing through 3 is invalid.
    // Consider the remaining routes from 0 to 4:
    // Path 0 → 1 → 4
    // Total cost = 7 + 5 = 12 <= k, so this path is valid.
    // The minimum edge‐cost along this path is min(7, 5) = 5.
    // Path 0 → 2 → 3 → 4
    // Node 3 is offline, so this path is invalid regardless of cost.
    // Path 0 → 2 → 4
    // Total cost = 6 + 6 = 12 <= k, so this path is valid.
    // The minimum edge‐cost along this path is min(6, 6) = 6.
    // Among the two valid paths, their scores are 5 and 6. Therefore, the answer is 6.
    fmt.Println(findMaxPathScore([][]int{{0,1,7},{1,4,5},{0,2,6},{2,3,6},{3,4,2},{2,4,6}}, []bool{true,true,true,false,true}, 12)) // 3

    fmt.Println(findMaxPathScore1([][]int{{0,1,5},{1,3,10},{0,2,3},{2,3,4}},[]bool{true,true,true,true}, 10)) // 3
    fmt.Println(findMaxPathScore1([][]int{{0,1,7},{1,4,5},{0,2,6},{2,3,6},{3,4,2},{2,4,6}}, []bool{true,true,true,false,true}, 12)) // 3
    
    fmt.Println(findMaxPathScore2([][]int{{0,1,5},{1,3,10},{0,2,3},{2,3,4}},[]bool{true,true,true,true}, 10)) // 3
    fmt.Println(findMaxPathScore2([][]int{{0,1,7},{1,4,5},{0,2,6},{2,3,6},{3,4,2},{2,4,6}}, []bool{true,true,true,false,true}, 12)) // 3
}