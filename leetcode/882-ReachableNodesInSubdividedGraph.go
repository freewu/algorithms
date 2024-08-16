package main

// 882. Reachable Nodes In Subdivided Graph
// You are given an undirected graph (the "original graph") with n nodes labeled from 0 to n - 1. 
// You decide to subdivide each edge in the graph into a chain of nodes, 
// with the number of new nodes varying between each edge.

// The graph is given as a 2D array of edges where edges[i] = [ui, vi, cnti] indicates 
// that there is an edge between nodes ui and vi in the original graph, 
// and cnti is the total number of new nodes that you will subdivide the edge into. 
// Note that cnti == 0 means you will not subdivide the edge.

// To subdivide the edge [ui, vi], replace it with (cnti + 1) new edges and cnti new nodes. 
// The new nodes are x1, x2, ..., xcnti, and the new edges are [ui, x1], [x1, x2], [x2, x3], ..., [xcnti-1, xcnti], [xcnti, vi].

// In this new graph, you want to know how many nodes are reachable from the node 0, 
// where a node is reachable if the distance is maxMoves or less.

// Given the original graph and maxMoves, 
// return the number of nodes that are reachable from node 0 in the new graph.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/01/origfinal.png" />
// Input: edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
// Output: 13
// Explanation: The edge subdivisions are shown in the image above.
// The nodes that are reachable are highlighted in yellow.

// Example 2:
// Input: edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
// Output: 23

// Example 3:
// Input: edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
// Output: 1
// Explanation: Node 0 is disconnected from the rest of the graph, so only node 0 is reachable.

// Constraints:
//     0 <= edges.length <= min(n * (n - 1) / 2, 10^4)
//     edges[i].length == 3
//     0 <= ui < vi < n
//     There are no multiple edges in the graph.
//     0 <= cnti <= 10^4
//     0 <= maxMoves <= 10^9
//     1 <= n <= 3000

import "fmt"
import "container/heap"

func reachableNodes(edges [][]int, maxMoves int, n int) int {
    const  N = 3010
    const INF = 0x3f3f3f3f
    res, dist, graphs := 0, [N]int{}, [N]map[int]int{}
    for i := 0; i < n; i++ {
        dist[i] = INF
        graphs[i] = map[int]int{}
    }
    dist[0] = 0
    for i := 0; i < len(edges); i++ {
        e := edges[i]
        graphs[e[0]][e[1]] = e[2] + 1
        graphs[e[1]][e[0]] = e[2] + 1
    }
    queue, set := []int{0}, [N]bool{}
    set[0] = true
    for len(queue) > 0 {
        t := queue[0]
        queue = queue[1:]
        set[t] = false
        for j, w := range graphs[t] {
            if dist[j] > dist[t] + w {
                dist[j] = dist[t] + w
                if !set[j] {
                    queue= append(queue, j)
                    set[j] = true
                }
            }
        }
    }
    for i := 0; i < n; i++ {
        if dist[i] <= maxMoves {
            res++
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, e := range edges {
        x := max(maxMoves - dist[e[0]], 0)
        y := max(maxMoves - dist[e[1]], 0)
        res += min(e[2], x + y)
    }
    return res
}

func reachableNodes1(edges [][]int, maxMoves int, n int) int {
    res, g := 0, make([][]neigbour, n)
    for _, edge := range edges {
        u, v, cnt := edge[0], edge[1], edge[2]
        g[u] = append(g[u], neigbour{v, cnt + 1})
        g[v] = append(g[v], neigbour{u, cnt + 1})
    }
    dist := dijkstra(g, 0)
    for i := range dist {
        if dist[i] <= maxMoves {
            res++
        }
    }
    for _, e := range edges {
        u, v, cnt := e[0], e[1], e[2]
        a := max(maxMoves-dist[u], 0)
        b := max(maxMoves-dist[v], 0)
        res += min(a+b, cnt)
    }
    return res
}

type neigbour struct{ to, weight int }

func dijkstra(g [][]neigbour, start int) []int {
    dist, inf := make([]int, len(g)), 1 << 31
    for i := range dist {
        dist[i] = inf
    }
    dist[start] = 0
    var mh = &hp{pair{start, 0}}
    heap.Init(mh)
    for mh.Len() > 0 {
        head := heap.Pop(mh).(pair)
        x := head.x
        weight := head.dist
        if dist[x] < weight {
            continue
        }
        for _, y := range g[x] {
            to := y.to
            if d := dist[x] + y.weight; d < dist[to] {
                dist[to] = d
                heap.Push(mh, pair{to, d})
            }
        }
    }
    return dist
}

type pair struct{ x, dist int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }


func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/01/origfinal.png" />
    // Input: edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
    // Output: 13
    // Explanation: The edge subdivisions are shown in the image above.
    // The nodes that are reachable are highlighted in yellow.
    fmt.Println(reachableNodes([][]int{{0,1,10},{0,2,1},{1,2,2}}, 6, 3)) // 13
    // Example 2:
    // Input: edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
    // Output: 23
    fmt.Println(reachableNodes([][]int{{0,1,4},{1,2,6},{0,2,8},{1,3,1}}, 10, 4)) // 23
    // Example 3:
    // Input: edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
    // Output: 1
    // Explanation: Node 0 is disconnected from the rest of the graph, so only node 0 is reachable.
    fmt.Println(reachableNodes([][]int{{1,2,4},{1,4,5},{1,3,1},{2,3,4},{3,4,5}}, 12, 5)) // 1

    fmt.Println(reachableNodes1([][]int{{0,1,10},{0,2,1},{1,2,2}}, 6, 3)) // 13
    fmt.Println(reachableNodes1([][]int{{0,1,4},{1,2,6},{0,2,8},{1,3,1}}, 10, 4)) // 23
    fmt.Println(reachableNodes1([][]int{{1,2,4},{1,4,5},{1,3,1},{2,3,4},{3,4,5}}, 12, 5)) // 1
}