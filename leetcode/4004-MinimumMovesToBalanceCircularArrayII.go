package main

// 4004. Minimum Moves to Balance Circular Array II
// You are given a circular array balance of length n, where balance[i] is the net balance of person i.

// In one move, a person can transfer exactly 1 unit of balance to either their left or right neighbor.

// Return the minimum number of moves required so that every person has a non-negative balance. 
// If it is impossible, return -1.

// Example 1:
// Input: balance = [-1,2,-1]
// Output: 2
// Explanation:
// One optimal sequence of moves is:
// Move 1 unit from i = 1 to i = 0, resulting in balance = [0, 1, -1]
// Move 1 unit from i = 1 to i = 2, resulting in balance = [0, 0, 0]
// Thus, the minimum number of moves required is 2.

// Example 2:
// Input: balance = [4,-1,-2]
// Output: 3
// Explanation:
// One optimal sequence of moves is:
// Move 1 unit from i = 0 to i = 1, resulting in balance = [3, 0, -2]
// Move 1 unit from i = 0 to i = 2, resulting in balance = [2, 0, -1]
// Move 1 unit from i = 0 to i = 2, resulting in balance = [1, 0, 0]
// Thus, the minimum number of moves required is 3.

// Example 3:
// Input: balance = [-3,-3,5]
// Output: -1
// Explanation:
// It is impossible to make all balances non-negative for balance = [-3, -3, 5], so the answer is -1.

// Constraints:
//     1 <= n == balance.length <= 1000
//     -10^5 <= balance[i] <= 10^5

import "fmt"
import "container/heap"

type Edge struct {
    to, rev int
    cap     int64
    cost    int64
}

type MinCostFlow struct {
    g    [][]Edge
    h    []int64 // potential
    dist []int64
    prev []int
    preE []int
}

func NewMinCostFlow(n int) *MinCostFlow {
    return &MinCostFlow{
        g:    make([][]Edge, n),
        h:    make([]int64, n),
        dist: make([]int64, n),
        prev: make([]int, n),
        preE: make([]int, n),
    }
}

func (mcf *MinCostFlow) AddEdge(from, to int, cap int64, cost int64) {
    mcf.g[from] = append(mcf.g[from], Edge{to, len(mcf.g[to]), cap, cost})
    mcf.g[to] = append(mcf.g[to], Edge{from, len(mcf.g[from]) - 1, 0, -cost})
}

type Item struct {
    v   int
    d   int64
    idx int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].d < pq[j].d }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*Item))
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[:n-1]
    return x
}

// Flow : return (flow, cost)
func (mcf *MinCostFlow) Flow(s, t int, mx int64) (int64, int64) {
    flow, cost, inf := int64(0), int64(0), int64(1 << 61)
    for flow < mx {
        for i := range mcf.dist {
            mcf.dist[i] = inf
        }
        mcf.dist[s] = 0
        pq := make(PriorityQueue, 0)
        heap.Init(&pq)
        heap.Push(&pq, &Item{v: s, d: 0})
        for pq.Len() > 0 {
            item := heap.Pop(&pq).(*Item)
            v := item.v
            if item.d > mcf.dist[v] {
                continue
            }
            for i, e := range mcf.g[v] {
                if e.cap > 0 && mcf.dist[e.to] > mcf.dist[v]+e.cost+mcf.h[v]-mcf.h[e.to] {
                    mcf.dist[e.to] = mcf.dist[v] + e.cost + mcf.h[v] - mcf.h[e.to]
                    mcf.prev[e.to] = v
                    mcf.preE[e.to] = i
                    heap.Push(&pq, &Item{v: e.to, d: mcf.dist[e.to]})
                }
            }
        }
        if mcf.dist[t] == inf {
            return flow, cost // no more augment
        }
        for v := range mcf.h {
            if mcf.dist[v] < inf {
                mcf.h[v] += mcf.dist[v]
            }
        }
        d := mx - flow
        for v := t; v != s; v = mcf.prev[v] {
            d = min(d, mcf.g[mcf.prev[v]][mcf.preE[v]].cap)
        }
        flow += d
        cost += d * mcf.h[t]
        for v := t; v != s; v = mcf.prev[v] {
            e := &mcf.g[mcf.prev[v]][mcf.preE[v]]
            e.cap -= d
            mcf.g[v][e.rev].cap += d
        }
    }
    return flow, cost
}

func minMoves(balance []int) int64 {
    n, inf := len(balance), int64(1e18)
    S, T := n, n + 1
    mcf := NewMinCostFlow(n + 2)
    totalNeed, sumTotal := int64(0), int64(0)
    for _, v := range balance {
        sumTotal += int64(v)
    }
    if sumTotal < 0 {
        return -1
    }
    for i := 0; i < n; i++ {
        val := int64(balance[i])
        if val > 0 {
            mcf.AddEdge(S, i, val, 0)
        } else if val < 0 {
            need := -val
            mcf.AddEdge(i, T, need, 0)
            totalNeed += need
        }
        // 环形相邻双向边
        j := (i + 1) % n
        mcf.AddEdge(i, j, inf, 1)
        mcf.AddEdge(j, i, inf, 1)
    }
    f, c := mcf.Flow(S, T, totalNeed)
    if f < totalNeed {
        return -1
    }
    return c
}

func main() {
    // Example 1:
    // Input: balance = [-1,2,-1]
    // Output: 2
    // Explanation:
    // One optimal sequence of moves is:
    // Move 1 unit from i = 1 to i = 0, resulting in balance = [0, 1, -1]
    // Move 1 unit from i = 1 to i = 2, resulting in balance = [0, 0, 0]
    // Thus, the minimum number of moves required is 2.
    fmt.Println(minMoves([]int{-1,2,-1})) // 2
    // Example 2:
    // Input: balance = [4,-1,-2]
    // Output: 3
    // Explanation:
    // One optimal sequence of moves is:
    // Move 1 unit from i = 0 to i = 1, resulting in balance = [3, 0, -2]
    // Move 1 unit from i = 0 to i = 2, resulting in balance = [2, 0, -1]
    // Move 1 unit from i = 0 to i = 2, resulting in balance = [1, 0, 0]
    // Thus, the minimum number of moves required is 3.
    fmt.Println(minMoves([]int{4,-1,-2})) // 3
    // Example 3:
    // Input: balance = [-3,-3,5]
    // Output: -1
    // Explanation:
    // It is impossible to make all balances non-negative for balance = [-3, -3, 5], so the answer is -1.
    fmt.Println(minMoves([]int{-3,-3,5})) // -1

    // Example 4:
    // Input: balance = [16,6,-6,-16]
    // Output: 22
    fmt.Println(minMoves([]int{16,6,-6,-16})) // 22

    fmt.Println(minMoves([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(minMoves([]int{9,8,7,6,5,4,3,2,1})) // 0 
}