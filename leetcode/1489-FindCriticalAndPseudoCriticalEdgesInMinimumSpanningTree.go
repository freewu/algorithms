package main

// 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree
// Given a weighted undirected connected graph with n vertices numbered from 0 to n - 1, 
// and an array edges where edges[i] = [ai, bi, weighti] represents a bidirectional 
// and weighted edge between nodes ai and bi. 
// A minimum spanning tree (MST) is a subset of the graph's edges 
// that connects all vertices without cycles and with the minimum possible total edge weight.

// Find all the critical and pseudo-critical edges in the given graph's minimum spanning tree (MST). 
// An MST edge whose deletion from the graph would cause the MST weight to increase is called a critical edge. 
// On the other hand, a pseudo-critical edge is that which can appear in some MSTs but not all.

// Note that you can return the indices of the edges in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/06/04/ex1.png" />
// Input: n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
// Output: [[0,1],[2,3,4,5]]
// Explanation: The figure above describes the graph.
// The following figure shows all the possible MSTs:
// <img src="https://assets.leetcode.com/uploads/2020/06/04/msts.png" />
// Notice that the two edges 0 and 1 appear in all MSTs, therefore they are critical edges, so we return them in the first list of the output.
// The edges 2, 3, 4, and 5 are only part of some MSTs, therefore they are considered pseudo-critical edges. We add them to the second list of the output.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/06/04/ex2.png" />
// Input: n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
// Output: [[],[0,1,2,3]]
// Explanation: We can observe that since all 4 edges have equal weight, choosing any 3 edges from the given 4 will yield an MST. Therefore all 4 edges are pseudo-critical.

// Constraints:
//     2 <= n <= 100
//     1 <= edges.length <= min(200, n * (n - 1) / 2)
//     edges[i].length == 3
//     0 <= ai < bi < n
//     1 <= weighti <= 1000
//     All pairs (ai, bi) are distinct.

import "fmt"
import "sort"

type UnionFind struct {
    parent []int
    size []int
    MaxSize int
}

func (f *UnionFind) Find(x int) int {
    if x != f.parent[x] {
        f.parent[x] = f.Find(f.parent[x])
    }
    return f.parent[x]
}

func (f *UnionFind) Unite(x, y int) bool {
    rootX, rootY := f.Find(x), f.Find(y)
    if rootX != rootY {
        if f.size[rootX] < f.size[rootY] {
            rootX, rootY = rootY, rootX
        }
        f.parent[rootY] = rootX
        f.size[rootX] += f.size[rootY]
        if f.MaxSize < f.size[rootX] {
            f.MaxSize = f.size[rootX]
        }
        return true
    }
    return false
}

func newUnionFind(n int) *UnionFind {
    parent, size := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        parent[i], size[i] = i, 1
    }

    return &UnionFind{parent, size, 1}
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
    m := len(edges)
    newEdges := make([][]int, m)
    for i := 0; i < m; i++ {
        newEdges[i] = make([]int, 4)
        for j := 0; j < 3; j++ {
            newEdges[i][j] = edges[i][j]
        }
        newEdges[i][3] = i
    }
    sort.Slice(newEdges, func(i, j int) bool {
        return newEdges[i][2] < newEdges[j][2]
    })
    ufStd, stdWeight := newUnionFind(n), 0
    for _, edge := range newEdges {
        if ufStd.Unite(edge[0], edge[1]) {
            stdWeight += edge[2]
        }
    }
    res := make([][]int, 2)
    for i := 0; i < m; i++ {
        ufIgnore, ignoreWeight := newUnionFind(n), 0
        for j := 0; j < m; j++ {
            if i != j && ufIgnore.Unite(newEdges[j][0], newEdges[j][1]) {
                ignoreWeight += newEdges[j][2]
            }
        }

        if ufIgnore.MaxSize < n || ignoreWeight > stdWeight {
            res[0] = append(res[0], newEdges[i][3])
        } else {
            ufForce := newUnionFind(n)
            ufForce.Unite(newEdges[i][0], newEdges[i][1])
            forceWeight := newEdges[i][2]
            for j := 0; j < m; j++ {
                if i != j && ufForce.Unite(newEdges[j][0], newEdges[j][1]) {
                    forceWeight += newEdges[j][2]
                }
            }
            if forceWeight == stdWeight {
                res[1] = append(res[1], newEdges[i][3])
            }
        }
    }
    return res
}

type UnionFind1 struct {
    parent, size []int
}

func NewUnionFind1(n int) *UnionFind1 {
    parent := make([]int, n)
    size := make([]int, n)
    for i := range parent {
        parent[i] = i
        size[i] = 1
    }
    return &UnionFind1{parent, size}
}

func (this *UnionFind1) find(x int) int {
    if this.parent[x] != x {
        this.parent[x] = this.find(this.parent[x])
    }
    return this.parent[x]
}

func (this *UnionFind1) union(x, y int) bool {
    fx, fy := this.find(x), this.find(y)
    if fx == fy {
        return false
    }
    if this.size[fx] < this.size[fy] {
        fx, fy = fy, fx
    }
    this.size[fx] += this.size[fy]
    this.parent[fy] = fx
    return true
}

func findCriticalAndPseudoCriticalEdges1(n int, edges [][]int) [][]int {
    m := len(edges)
    edgeType := make([]int, m) // -1：不在最小生成树中；0：伪关键边；1：关键边

    for i, e := range edges {
        edges[i] = append(e, i)
    }
    sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

    type neighbor struct{ to, edgeID int }
    graph := make([][]neighbor, n)
    dfn := make([]int, n) // 遍历到该顶点时的时间戳
    timestamp := 0
    var tarjan func(int, int) int
    tarjan = func(v, pid int) int {
        timestamp++
        dfn[v] = timestamp
        lowV := timestamp
        for _, e := range graph[v] {
            if w := e.to; dfn[w] == 0 {
                lowW := tarjan(w, e.edgeID)
                if lowW > dfn[v] {
                    edgeType[e.edgeID] = 1
                }
                lowV = min(lowV, lowW)
            } else if e.edgeID != pid {
                lowV = min(lowV, dfn[w])
            }
        }
        return lowV
    }
    uf := NewUnionFind1(n)
    for i := 0; i < m; {
        vs := []int{}
        // 将权值相同的边分为一组，建图，然后用 Tarjan 算法找桥边
        for weight := edges[i][2]; i < m && edges[i][2] == weight; i++ {
            e := edges[i]
            v, w, edgeID := uf.find(e[0]), uf.find(e[1]), e[3]
            if v != w {
                graph[v] = append(graph[v], neighbor{w, edgeID})
                graph[w] = append(graph[w], neighbor{v, edgeID})
                vs = append(vs, v, w) // 记录图中顶点
            } else {
                edgeType[edgeID] = -1
            }
        }
        for _, v := range vs {
            if dfn[v] == 0 {
                tarjan(v, -1)
            }
        }
        // 合并顶点、重置数据
        for j := 0; j < len(vs); j += 2 {
            v, w := vs[j], vs[j+1]
            uf.union(v, w)
            graph[v] = nil
            graph[w] = nil
            dfn[v] = 0
            dfn[w] = 0
        }
    }
    keyEdges, pseudokeyEdges := []int{},[]int{}
    for i, tp := range edgeType {
        if tp == 0 {
            pseudokeyEdges = append(pseudokeyEdges, i)
        } else if tp == 1 {
            keyEdges = append(keyEdges, i)
        }
    }
    return [][]int{keyEdges, pseudokeyEdges}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/06/04/ex1.png" />
    // Input: n = 5, edges = [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]]
    // Output: [[0,1],[2,3,4,5]]
    // Explanation: The figure above describes the graph.
    // The following figure shows all the possible MSTs:
    // <img src="https://assets.leetcode.com/uploads/2020/06/04/msts.png" />
    // Notice that the two edges 0 and 1 appear in all MSTs, therefore they are critical edges, so we return them in the first list of the output.
    // The edges 2, 3, 4, and 5 are only part of some MSTs, therefore they are considered pseudo-critical edges. We add them to the second list of the output.
    fmt.Println(findCriticalAndPseudoCriticalEdges(5,[][]int{{0,1,1},{1,2,1},{2,3,2},{0,3,2},{0,4,3},{3,4,3},{1,4,6}})) // [[0,1],[2,3,4,5]]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/06/04/ex2.png" />
    // Input: n = 4, edges = [[0,1,1],[1,2,1],[2,3,1],[0,3,1]]
    // Output: [[],[0,1,2,3]]
    // Explanation: We can observe that since all 4 edges have equal weight, choosing any 3 edges from the given 4 will yield an MST. Therefore all 4 edges are pseudo-critical.
    fmt.Println(findCriticalAndPseudoCriticalEdges(4,[][]int{{0,1,1},{1,2,1},{2,3,1},{0,3,1}})) // [[],[0,1,2,3]]

    fmt.Println(findCriticalAndPseudoCriticalEdges1(5,[][]int{{0,1,1},{1,2,1},{2,3,2},{0,3,2},{0,4,3},{3,4,3},{1,4,6}})) // [[0,1],[2,3,4,5]]
    fmt.Println(findCriticalAndPseudoCriticalEdges1(4,[][]int{{0,1,1},{1,2,1},{2,3,1},{0,3,1}})) // [[],[0,1,2,3]]
}