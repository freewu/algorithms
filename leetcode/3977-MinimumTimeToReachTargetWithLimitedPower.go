package main

// 3977. Minimum Time to Reach Target With Limited Power
// You are given a directed weighted graph with n nodes labeled from 0 to n - 1.

// The graph is represented by a 2D integer array edges, 
// where edges[i] = [ui, vi, ti] indicates a directed edge from node ui to node vi that takes ti seconds to traverse.

// You are also given an integer power representing the initial available power, 
// and an integer array cost of length n, where cost[u] represents the power required to forward the signal from node u through any one of its outgoing edges.

// You are given two integers source and target.

// The signal starts at source at time 0 with power units of power and follows these rules:
//     1. The signal may traverse a directed edge from node u only if the remaining power is at least cost[u].
//     2. No power is consumed when the signal arrives at a node, unless it later leaves that node by traversing another edge.
//     3. When the signal is forwarded from node u, the remaining power is decreased by cost[u] units.
//     4. Traversing an edge edges[i] = [ui, vi, ti] increases the total time by ti seconds.

// Return an integer array answer of size 2, where:
//     1. answer[0] is the minimum time required for the signal to reach node target.
//     2. answer[1] is the maximum remaining power among all paths that achieve answer[0].

// If the signal cannot reach target, return [-1, -1].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2026/02/09/g1.png" />
// Input: n = 5, edges = [[0,1,1],[1,4,1],[0,2,1],[2,3,1],[3,4,1]], power = 4, cost = [2,3,1,1,1], source = 0, target = 4
// Output: [3,0]
// Explanation:
// The signal starts at node 0 with 4 units of power.
// The path 0 -> 1 -> 4 is not valid, because after leaving node 0, the signal has 2 units of power remaining, which is less than cost[1] = 3.
// The valid path 0 -> 2 -> 3 -> 4 takes a total time of 3.
// The total power consumed along this path is cost[0] + cost[2] + cost[3] = 4, leaving 0 remaining power.
// Hence, the answer is [3, 0].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2026/02/09/g22.png" />
// Input: n = 3, edges = [[0,1,2],[1,2,2],[2,0,2]], power = 3, cost = [1,1,1], source = 1, target = 1
// Output: [0,3]
// Explanation:
// Since the source and target are the same node, no traversal is required.
// Hence, the minimum total time taken is 0, and no power is consumed.
// Therefore, the answer is [0, 3].

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2026/02/09/g23.png" />
// ​​​​​Input: n = 4, edges = [[0,1,3],[2,3,4]], power = 3, cost = [1,1,1,1], source = 0, target = 3
// Output: [-1,-1]
// Explanation:
// There is no valid path from source to target, therefore return [-1, -1].

// Constraints:
//     1 <= n <= 1000
//     0 <= edges.length <= 1000
//     edges[i] = [ui, vi, ti]
//     0 <= ui, vi <= n - 1
//     1 <= ti <= 10^9
//     1 <= power <= 1000
//     cost.length == n
//     1 <= cost[i] <= 2000
//     0 <= source, target <= n - 1

import "fmt"
import "container/heap"

type State struct {
    time int64
    node int
    remp int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].time < pq[j].time }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(State)) }
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    m := len(old)
    item := old[m-1]
    *pq = old[:m-1]
    return item
}

func minTimeMaxPower(n int, edges [][]int, p int, cost []int, s int, t int) []int64 {
    adj := make([][][2]int, n)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        adj[u] = append(adj[u], [2]int{v, w})
    }
    res := []int64{-1, -1}
    dist := make([][]int64, n)
    for i := range dist {
        dist[i] = make([]int64, p+1)
        for j := range dist[i] {
            dist[i][j] = int64(1 << 61)
        }
    }
    dist[s][p] = 0
    // priority queue stores {time, node, remaining power}
    pq := &PriorityQueue{}
    *pq = append(*pq, State{0, s, p})
    heap.Init(pq)
    best, bestp := int64(-1), int64(-1)
    for pq.Len() > 0 {
        vec := heap.Pop(pq).(State)
        time, u, remp := vec.time, vec.node, vec.remp
        if time != dist[u][remp] {
            continue
        }
        if best != -1 && time > best {
            break
        }
        if u == t {
            if best == -1 {
                best = time
            }
            if int64(remp) > bestp {
                bestp = int64(remp)
            }
            continue
        }
        if remp < cost[u] {
            continue
        }
        nxtp := remp - cost[u]
        for _, nei := range adj[u] {
            v, w := nei[0], nei[1]
            ntime := time + int64(w)
            if ntime < dist[v][nxtp] {
                dist[v][nxtp] = ntime
                heap.Push(pq, State{ntime, v, nxtp})
            }
        }
    }
    res[0], res[1] = best, bestp
    return res  
}

type State1 struct {
    t, p, u int64
}

func (a State1) Less(b State1) bool {
    if a.t == b.t{
        return a.p > b.p
    }
    return a.t < b.t
}

type PriorityQueue1[T any] struct {
    data []T
    less func(a, b T) bool
}

func NewPriorityQueue1[T any](less func(a, b T) bool) *PriorityQueue1[T]{
    return &PriorityQueue1[T]{data: make([]T, 0), less: less}
}

func (h *PriorityQueue1[T]) Push(x T) {
    h.data = append(h.data, x)
    h.up(len(h.data) - 1)
}

func (h *PriorityQueue1[T]) Pop() T {
    n := len(h.data) - 1
    h.data[0], h.data[n] = h.data[n], h.data[0]
    x := h.data[n]
    h.data = h.data[0:n]
    h.down(0, n)
    return x
}

func (h *PriorityQueue1[T]) Len() int {
    return len(h.data)
}

func (h *PriorityQueue1[T]) up(j int) {
    for {
        i := (j - 1) / 2 
        if i == j || !h.less(h.data[j], h.data[i]) { 
            break
        }
        h.data[i], h.data[j] = h.data[j], h.data[i]
        j = i
    }
}

func (h *PriorityQueue1[T]) down(i0, n int) bool {
    i := i0
    for {
        j1 := 2*i + 1
        if j1 >= n || j1 < 0 { 
            break
        }
        j := j1 
        if j2 := j1 + 1; j2 < n && h.less(h.data[j2], h.data[j1]) {
            j = j2 
        }
        if !h.less(h.data[j], h.data[i]) {
            break
        }
        h.data[i], h.data[j] = h.data[j], h.data[i]
        i = j
    }
    return i > i0
}

func minTimeMaxPower1(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
    adj := make([][][2]int64, n)
    for _, e := range edges{
        adj[e[0]] = append(adj[e[0]], [2]int64{int64(e[1]), int64(e[2])})
    }
    pt := make([]int64, n)
    for i := range pt{
        pt[i] = int64(-1)
    }
    g1 := func(a, b State1) bool{
        if a.t == b.t{
            return a.p > b.p
        }
        return a.t < b.t
    }
    pq := NewPriorityQueue1[State1](g1)
    pq.Push(State1{0, int64(power), int64(source)})
    for pq.Len() > 0{
        c := pq.Pop()
        if c.p <= pt[c.u]{
            continue
        }
        pt[c.u] = c.p
        if c.u == int64(target){
            return []int64{int64(c.t), int64(c.p)}
        }
        if c.p >= int64(cost[c.u]){
            p := c.p - int64(cost[c.u])
            for _, e := range adj[c.u]{
                v, w := e[0], e[1]
                if p > pt[v]{
                    pq.Push(State1{c.t + w, p, v})
                }
            }
        }
    }
    return []int64{-1, -1}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2026/02/09/g1.png" />
    // Input: n = 5, edges = [[0,1,1],[1,4,1],[0,2,1],[2,3,1],[3,4,1]], power = 4, cost = [2,3,1,1,1], source = 0, target = 4
    // Output: [3,0]
    // Explanation:
    // The signal starts at node 0 with 4 units of power.
    // The path 0 -> 1 -> 4 is not valid, because after leaving node 0, the signal has 2 units of power remaining, which is less than cost[1] = 3.
    // The valid path 0 -> 2 -> 3 -> 4 takes a total time of 3.
    // The total power consumed along this path is cost[0] + cost[2] + cost[3] = 4, leaving 0 remaining power.
    // Hence, the answer is [3, 0].
    fmt.Println(minTimeMaxPower(5, [][]int{{0,1,1},{1,4,1},{0,2,1},{2,3,1},{3,4,1}}, 4, []int{2,3,1,1,1}, 0, 4)) // [3,0]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2026/02/09/g22.png" />
    // Input: n = 3, edges = [[0,1,2],[1,2,2],[2,0,2]], power = 3, cost = [1,1,1], source = 1, target = 1
    // Output: [0,3]
    // Explanation:
    // Since the source and target are the same node, no traversal is required.
    // Hence, the minimum total time taken is 0, and no power is consumed.
    // Therefore, the answer is [0, 3].
    fmt.Println(minTimeMaxPower(3, [][]int{{0,1,2},{1,2,2},{2,0,2}}, 3, []int{1,1,1}, 1, 1)) // [0,3]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2026/02/09/g23.png" />
    // ​​​​​Input: n = 4, edges = [[0,1,3],[2,3,4]], power = 3, cost = [1,1,1,1], source = 0, target = 3
    // Output: [-1,-1]
    // Explanation:
    // There is no valid path from source to target, therefore return [-1, -1].
    fmt.Println(minTimeMaxPower(4, [][]int{{0,1,3},{2,3,4}}, 3, []int{1,1,1,1}, 0, 3)) // [-1,-1]
    
    fmt.Println(minTimeMaxPower1(5, [][]int{{0,1,1},{1,4,1},{0,2,1},{2,3,1},{3,4,1}}, 4, []int{2,3,1,1,1}, 0, 4)) // [3,0]
    fmt.Println(minTimeMaxPower1(3, [][]int{{0,1,2},{1,2,2},{2,0,2}}, 3, []int{1,1,1}, 1, 1)) // [0,3]
    fmt.Println(minTimeMaxPower1(4, [][]int{{0,1,3},{2,3,4}}, 3, []int{1,1,1,1}, 0, 3)) // [-1,-1]
}
