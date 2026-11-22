package main

// 4122. Earliest Signal Activation Times
// You are given an integer n and a directed graph with n nodes labeled from 0 to n - 1. 
// The graph is represented by a 2D array edges, where edges[i] = [u_{i}, v_{i}, delay_{i}] indicates an edge from node u_{i} to v_{i} along which a signal takes delay_{i} units of time to arrive.

// Node u is an incoming neighbor of node v if there is an edge from u to v.

// Node 0 is active at time 0, and all other nodes are initially inactive.

// When a node becomes active at time t, it immediately sends one signal along each of its outgoing edges. 
// Once active, a node remains active and does not send its signals again. 
// A signal sent along an edge with delay d arrives at time t + d.

// You are also given an integer array requiredSignals of length n, 
// where requiredSignals[i] is the number of distinct incoming neighbors of node i from which a signal must arrive before node i can activate.

// An inactive node i becomes active at the earliest time t at which signals have arrived from at least requiredSignals[i] distinct incoming neighbors.

// Return an array ans of length n where ans[i] is the earliest time node i becomes active. 
// If node i never becomes active, ans[i] should be -1.

// Note: The value of requiredSignals[0] is irrelevant, since node 0 is active at time 0 and sends its outgoing signals at that time.

// Example 1:
// Input: n = 3, edges = [[0,1,1],[0,2,1]], requiredSignals = [1,1,2]
// Output: [0,1,-1]
// Explanation:
//     Node 0 is active at time 0 and sends signals to nodes 1 and 2.
//     At time 1, both signals arrive. Node 1 has received its required 1 signal, so it activates.
//     Node 2 requires signals from 2 distinct incoming neighbors, but it has only one incoming neighbor. It never activates.
//     Therefore, ans = [0, 1, -1].

// Example 2:
// Input: n = 4, edges = [[0,1,1],[0,2,1],[1,3,2],[2,3,2]], requiredSignals = [3,1,1,2]
// Output: [0,1,1,3]
// Explanation:
//     Node 0 is active at time 0, regardless of requiredSignals[0] = 3, and sends signals to nodes 1 and 2.
//     At time 1, nodes 1 and 2 each receive their required 1 signal and activate. Each immediately sends a signal to node 3.
//     Both signals reach node 3 at time 1 + 2 = 3. They come from two distinct incoming neighbors, so node 3 activates at time 3.
//     Therefore, ans = [0, 1, 1, 3].

// Example 3:
// Input: n = 4, edges = [[0,1,2],[1,2,3],[0,2,5],[2,3,1]], requiredSignals = [3,1,2,1]
// Output: [0,2,5,6]
// Explanation:
//     Node 0 is active at time 0 and sends signals to nodes 1 and 2.
//     Node 1 receives its signal at time 2, activates, and sends a signal to node 2.
//     At time 5, node 2 receives the signal from node 0 and the signal from node 1, whose arrival time is 2 + 3 = 5. Its requirement of 2 distinct incoming neighbors is met, so it activates and sends a signal to node 3.
//     Node 3 receives that signal at time 5 + 1 = 6 and activates.
//     Therefore, ans = [0, 2, 5, 6].

// Example 4:
// Input: n = 4, edges = [[0,1,2],[1,0,1],[1,2,1],[2,3,1],[3,2,1]], requiredSignals = [3,1,2,1]
// Output: [0,2,-1,-1]
// Explanation:
//     Node 0 is active at time 0 and sends a signal to node 1.
//     Node 1 activates at time 2 and sends signals to nodes 0 and 2, which arrive at time 3.
//     Node 0 is already active, so receiving another signal does not cause it to send its signals again.
//     Node 2 requires signals from both nodes 1 and 3. It receives only the signal from node 1. Node 3 is waiting for a signal from node 2, so neither node activates.
//     Therefore, ans = [0, 2, -1, -1].

// Constraints:
//     2 <= n == requiredSignals.length <= 10^{5}
//     0 <= edges.length <= 10^{5}
//     edges[i] = [u_{i}, v_{i}, delay_{i}]
//     0 <= u_{i}, v_{i} <= n - 1
//     u_{i} != v_{i}
//     1 <= delay_{i} <= 10^{9}
//     There is at most one directed edge from any node u_{i} to any node v_{i}.
//     1 <= requiredSignals[i] <= n - 1

import "fmt"
import "container/heap"

type EdgeItem struct {
    to int
    d  int64
}

type SignalState struct {
    t int64
    u int
    v int
}

type MinSignalHeap []SignalState

func (h MinSignalHeap) Len() int           { return len(h) }
func (h MinSignalHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h MinSignalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinSignalHeap) Push(x any)        { *h = append(*h, x.(SignalState)) }
func (h *MinSignalHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}

func earliestActivationTimes(n int, edges [][]int, requiredSignals []int) []int64 {
    adj := make([][]EdgeItem, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], EdgeItem{to: e[1], d: int64(e[2])})
    }
    res, seen :=  make([]int64, n), make([]map[int]struct{}, n)
    for i := 0; i < n; i++ {
        res[i] = -1
        seen[i] = make(map[int]struct{})
    }
    pq := &MinSignalHeap{{t: 0, u: -1, v: 0}}
    heap.Init(pq)
    for pq.Len() > 0 {
        curr := heap.Pop(pq).(SignalState)
        t, u, v := curr.t, curr.u, curr.v

        if res[v] != -1 {
            continue
        }
        if u != -1 {
            seen[v][u] = struct{}{}
        }
        if v == 0 || len(seen[v]) >= requiredSignals[v] {
            res[v] = t
            for _, nxtEdge := range adj[v] {
                if res[nxtEdge.to] == -1 {
                    heap.Push(pq, SignalState{t: t + nxtEdge.d, u: v, v: nxtEdge.to})
                }
            }
        }
    }
    return res
}

func earliestActivationTimes1(n int, edges [][]int, requiredSignals []int) []int64 {
    head := make([]int32, n+1)
    for _, e := range edges {
        head[e[0]+1]++
    }
    for i := 0; i < n; i++ {
        head[i+1] += head[i]
    }
    to := make([]int32, len(edges))
    w := make([]int64, len(edges))
    pos := make([]int32, n)
    copy(pos, head[:n])
    for _, e := range edges {
        p := pos[e[0]]
        to[p], w[p] = int32(e[1]), int64(e[2])
        pos[e[0]]++
    }

    const B = 17
    h := make([]uint64, 1, len(edges)+1) // seed: time 0, node 0
    push := func(x uint64) {
        h = append(h, x)
        for i := len(h) - 1; i > 0; {
            p := (i - 1) >> 1
            if h[p] <= h[i] {
                break
            }
            h[p], h[i] = h[i], h[p]
            i = p
        }
    }
    pop := func() uint64 {
        top, last := h[0], len(h)-1
        h[0] = h[last]
        h = h[:last]
        for i := 0; ; {
            m := i
            if l := 2*i + 1; l < last && h[l] < h[m] {
                m = l
            }
            if r := 2*i + 2; r < last && h[r] < h[m] {
                m = r
            }
            if m == i {
                return top
            }
            h[i], h[m] = h[m], h[i]
            i = m
        }
    }
    res, count := make([]int64, n), make([]int, n)
    for i := range res {
        res[i] = -1
    }
    for len(h) > 0 {
        x := pop()
        t, u := int64(x>>B), int(x&(1<<B-1))
        if res[u] >= 0 {
            continue
        }
        if u != 0 {
            count[u]++
            if count[u] < requiredSignals[u] {
                continue
            }
        }
        res[u] = t
        for j := head[u]; j < head[u+1]; j++ {
            push(uint64(t+w[j])<<B | uint64(to[j]))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1,1],[0,2,1]], requiredSignals = [1,1,2]
    // Output: [0,1,-1]
    // Explanation:
    //     Node 0 is active at time 0 and sends signals to nodes 1 and 2.
    //     At time 1, both signals arrive. Node 1 has received its required 1 signal, so it activates.
    //     Node 2 requires signals from 2 distinct incoming neighbors, but it has only one incoming neighbor. It never activates.
    //     Therefore, ans = [0, 1, -1].
    fmt.Println(earliestActivationTimes(3, [][]int{{0,1,1},{0,2,1}}, []int{1,1,2})) // [0,1,-1]
    // Example 2:
    // Input: n = 4, edges = [[0,1,1],[0,2,1],[1,3,2],[2,3,2]], requiredSignals = [3,1,1,2]
    // Output: [0,1,1,3]
    // Explanation:
    //     Node 0 is active at time 0, regardless of requiredSignals[0] = 3, and sends signals to nodes 1 and 2.
    //     At time 1, nodes 1 and 2 each receive their required 1 signal and activate. Each immediately sends a signal to node 3.
    //     Both signals reach node 3 at time 1 + 2 = 3. They come from two distinct incoming neighbors, so node 3 activates at time 3.
    //     Therefore, ans = [0, 1, 1, 3].
    fmt.Println(earliestActivationTimes(4, [][]int{{0,1,1},{0,2,1},{1,3,2},{2,3,2}}, []int{3,1,1,2})) // [0,1,1,3]
    // Example 3:
    // Input: n = 4, edges = [[0,1,2],[1,2,3],[0,2,5],[2,3,1]], requiredSignals = [3,1,2,1]
    // Output: [0,2,5,6]
    // Explanation:
    //     Node 0 is active at time 0 and sends signals to nodes 1 and 2.
    //     Node 1 receives its signal at time 2, activates, and sends a signal to node 2.
    //     At time 5, node 2 receives the signal from node 0 and the signal from node 1, whose arrival time is 2 + 3 = 5. Its requirement of 2 distinct incoming neighbors is met, so it activates and sends a signal to node 3.
    //     Node 3 receives that signal at time 5 + 1 = 6 and activates.
    //     Therefore, ans = [0, 2, 5, 6].
    fmt.Println(earliestActivationTimes(4, [][]int{{0,1,2},{1,2,5},{0,2,5},{2,3,1}}, []int{3,1,2,1})) // [0,2,5,6]
    // Example 4:
    // Input: n = 4, edges = [[0,1,2],[1,0,1],[1,2,1],[2,3,1],[3,2,1]], requiredSignals = [3,1,2,1]
    // Output: [0,2,-1,-1]
    // Explanation:
    //     Node 0 is active at time 0 and sends a signal to node 1.
    //     Node 1 activates at time 2 and sends signals to nodes 0 and 2, which arrive at time 3.
    //     Node 0 is already active, so receiving another signal does not cause it to send its signals again.
    //     Node 2 requires signals from both nodes 1 and 3. It receives only the signal from node 1. Node 3 is waiting for a signal from node 2, so neither node activates.
    //     Therefore, ans = [0, 2, -1, -1].
    fmt.Println(earliestActivationTimes(4, [][]int{{0,1,2},{1,0,1},{1,2,1},{2,3,1},{3,2,1}}, []int{3,1,2,1})) // [0,2,-1,-1]

    fmt.Println(earliestActivationTimes1(3, [][]int{{0,1,1},{0,2,1}}, []int{1,1,2})) // [0,1,-1]
    fmt.Println(earliestActivationTimes1(4, [][]int{{0,1,1},{0,2,1},{1,3,2},{2,3,2}}, []int{3,1,1,2})) // [0,1,1,3]
    fmt.Println(earliestActivationTimes1(4, [][]int{{0,1,2},{1,2,5},{0,2,5},{2,3,1}}, []int{3,1,2,1})) // [0,2,5,6]
    fmt.Println(earliestActivationTimes1(4, [][]int{{0,1,2},{1,0,1},{1,2,1},{2,3,1},{3,2,1}}, []int{3,1,2,1})) // [0,2,-1,-1]
}
