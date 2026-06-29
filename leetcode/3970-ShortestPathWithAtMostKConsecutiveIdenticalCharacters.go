package main

// 3970. Shortest Path With At Most K Consecutive Identical Characters
// You are given an integer n representing the number of nodes in a directed weighted graph, numbered from 0 to n - 1. 
// This is represented by a 2D integer array edges, where edges[i] = [ui, vi, wi] represents a directed edge from node ui to node vi with weight wi.

// You are also given a string labels of length n, where labels[i] is the character assigned to node i, and an integer k.

// Return the minimum total edge weight of a path from node 0 to node n - 1 such that the concatenation of the labels of the nodes along the path contains at most k consecutive identical characters. 
// If no valid path exists, return -1.

// Example 1:
// Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,3]], labels = "aab", k = 1
// Output: 3
// Explanation:
// The optimal valid path from node 0 to node 2 is as follows:
// Use edges[2] = [0, 2, 3] to reach node 2 with a weight wi = 3.
// The corresponding concatenation of labels is "ab", which satisfies at most k = 1 consecutive identical characters. Thus, the answer is 3.

// Example 2:
// Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,3]], labels = "aab", k = 2
// Output: 2
// Explanation:
// The optimal valid path from node 0 to node 2 is as follows:
// Use edges[0] = [0, 1, 1] to reach node 1 with weight wi = 1.
// Use edges[1] = [1, 2, 1] to reach node 2 with weight wi = 1.
// The corresponding concatenation of labels is "aab", which satisfies at most k = 2 consecutive identical characters. Thus, the answer is 2.

// Example 3:
// Input: n = 3, edges = [[0,1,1],[1,2,1]], labels = "aaa", k = 2
// Output: -1
// Explanation:
// There is no valid path from node 0 to node 2 that satisfies at most k = 2 consecutive identical characters. Thus, the answer is -1.

// Constraints:
//     1 <= n == labels.length <= 5 * 10^4
//     0 <= edges.length <= 5 * 10^4
//     edges[i] == [ui, vi, wi]
//     0 <= ui, vi <= n - 1
//     ui != vi
//     1 <= wi <= 10^4
//     labels consists of lowercase English letters
//     1 <= k <= 50

import "fmt"
import "container/heap"

// 最短路长度, 节点, 最后连续相同字母个数
type Tuple struct{ distance, x, count int }
type MinHeap []Tuple
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Tuple)) }
func (h *MinHeap) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func shortestPath(n int, edges [][]int, labels string, k int) int {
    type Edge struct{ to, weight int }
    graph := make([][]Edge, n)
    for _, e := range edges {
        x, y, w := e[0], e[1], e[2]
        graph[x] = append(graph[x], Edge{y, w})
    }
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, k+1)
        for j := range dis[i] {
            dis[i][j] = 1 << 61
        }
    }
    h := MinHeap{}
    add := func(x, y, d int) {
        if d < dis[x][y] {
            dis[x][y] = d
            heap.Push(&h, Tuple{d, x, y})
        }
    }
    add(0, 1, 0)
    for len(h) > 0 {
        top := heap.Pop(&h).(Tuple)
        d := top.distance
        x, count := top.x, top.count
        if x == n-1 {
            return d
        }
        if d > dis[x][count] {
            continue
        }
        for _, e := range graph[x] {
            y := e.to
            if labels[y] != labels[x] {
                add(y, 1, d+e.weight)
            } else if count+1 <= k {
                add(y, count+1, d+e.weight)
            }
        }
    }
    return -1
}

func shortestPath1(n int, edges [][]int, labels string, k int) int {
    if n == 1 {
        return 0
    }
    lb := []byte(labels)
    m := len(edges)
    head := make([]int, n)
    for i := 0; i < n; i++ {
        head[i] = -1
    }
    to := make([]int, m)
    wt := make([]int, m)
    next := make([]int, m)
    for i, e := range edges {
        u, v, w := e[0], e[1], e[2]
        to[i] = v
        wt[i] = w
        next[i] = head[u]
        head[u] = i
    }
    const INF int64 = 1 << 62
    kk := k + 1 // run length in [1..k]
    dist := make([]int64, n*kk)
    for i := range dist {
        dist[i] = INF
    }
    type State struct {
        d    int64
        node int
        run  int
    }
    heap := make([]State, 0, 1024)
    push := func(s State) {
        heap = append(heap, s)
        i := len(heap) - 1
        for i > 0 {
            p := (i - 1) >> 1
            if heap[p].d <= heap[i].d {
                break
            }
            heap[p], heap[i] = heap[i], heap[p]
            i = p
        }
    }
    pop := func() State {
        res := heap[0]
        last := heap[len(heap)-1]
        heap = heap[:len(heap)-1]
        if len(heap) > 0 {
            heap[0] = last
            i := 0
            for {
                l := i*2 + 1
                if l >= len(heap) {
                    break
                }
                r := l + 1
                c := l
                if r < len(heap) && heap[r].d < heap[l].d {
                    c = r
                }
                if heap[i].d <= heap[c].d {
                    break
                }
                heap[i], heap[c] = heap[c], heap[i]
                i = c
            }
        }
        return res
    }
    startIdx := 0*kk + 1
    dist[startIdx] = 0
    push(State{d: 0, node: 0, run: 1})
    for len(heap) > 0 {
        cur := pop()
        u, run := cur.node, cur.run
        d := cur.d
        idx := u*kk + run
        if d != dist[idx] {
            continue
        }
        if u == n-1 {
            return int(d)
        }
        ul := lb[u]
        for ei := head[u]; ei != -1; ei = next[ei] {
            v := to[ei]
            nr := 1
            if lb[v] == ul {
                if run == k {
                    continue
                }
                nr = run + 1
            }
            nd := d + int64(wt[ei])
            vidx := v*kk + nr
            if nd < dist[vidx] {
                dist[vidx] = nd
                push(State{d: nd, node: v, run: nr})
            }
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,3]], labels = "aab", k = 1
    // Output: 3
    // Explanation:
    // The optimal valid path from node 0 to node 2 is as follows:
    // Use edges[2] = [0, 2, 3] to reach node 2 with a weight wi = 3.
    // The corresponding concatenation of labels is "ab", which satisfies at most k = 1 consecutive identical characters. Thus, the answer is 3.
    fmt.Println(shortestPath(3, [][]int{{0,1,1},{1,2,1},{0,2,3}}, "aab", 1)) // 3
    // Example 2:
    // Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,3]], labels = "aab", k = 2
    // Output: 2
    // Explanation:
    // The optimal valid path from node 0 to node 2 is as follows:
    // Use edges[0] = [0, 1, 1] to reach node 1 with weight wi = 1.
    // Use edges[1] = [1, 2, 1] to reach node 2 with weight wi = 1.
    // The corresponding concatenation of labels is "aab", which satisfies at most k = 2 consecutive identical characters. Thus, the answer is 2.
    fmt.Println(shortestPath(3, [][]int{{0,1,1},{1,2,1},{0,2,3}}, "aab", 2)) // 2
    // Example 3:
    // Input: n = 3, edges = [[0,1,1],[1,2,1]], labels = "aaa", k = 2
    // Output: -1
    // Explanation:
    // There is no valid path from node 0 to node 2 that satisfies at most k = 2 consecutive identical characters. Thus, the answer is -1.
    fmt.Println(shortestPath(3, [][]int{{0,1,1},{1,2,1}}, "aaa", 2)) // -1

    fmt.Println(shortestPath1(3, [][]int{{0,1,1},{1,2,1},{0,2,3}}, "aab", 1)) // 3
    fmt.Println(shortestPath1(3, [][]int{{0,1,1},{1,2,1},{0,2,3}}, "aab", 2)) // 2
    fmt.Println(shortestPath1(3, [][]int{{0,1,1},{1,2,1}}, "aaa", 2)) // -1
}