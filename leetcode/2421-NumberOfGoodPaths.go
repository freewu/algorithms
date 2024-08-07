package main

// 2421. Number of Good Paths
// There is a tree (i.e. a connected, undirected graph with no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges.

// You are given a 0-indexed integer array vals of length n where vals[i] denotes the value of the ith node. 
// You are also given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting nodes ai and bi.

// A good path is a simple path that satisfies the following conditions:
//     The starting node and the ending node have the same value.
//     All nodes between the starting node and the ending node have values less than or equal to the starting node (i.e. the starting node's value should be the maximum value along the path).

// Return the number of distinct good paths.

// Note that a path and its reverse are counted as the same path. 
// For example, 0 -> 1 is considered to be the same as 1 -> 0. A single node is also considered as a valid path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/08/04/f9caaac15b383af9115c5586779dec5.png" />
// Input: vals = [1,3,2,1,3], edges = [[0,1],[0,2],[2,3],[2,4]]
// Output: 6
// Explanation: There are 5 good paths consisting of a single node.
// There is 1 additional good path: 1 -> 0 -> 2 -> 4.
// (The reverse path 4 -> 2 -> 0 -> 1 is treated as the same as 1 -> 0 -> 2 -> 4.)
// Note that 0 -> 2 -> 3 is not a good path because vals[2] > vals[0].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/08/04/149d3065ec165a71a1b9aec890776ff.png" />
// Input: vals = [1,1,2,2,3], edges = [[0,1],[1,2],[2,3],[2,4]]
// Output: 7
// Explanation: There are 5 good paths consisting of a single node.
// There are 2 additional good paths: 0 -> 1 and 2 -> 3.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/08/04/31705e22af3d9c0a557459bc7d1b62d.png" />
// Input: vals = [1], edges = []
// Output: 1
// Explanation: The tree consists of only one node, so there is one good path.

// Constraints:
//     n == vals.length
//     1 <= n <= 3 * 10^4
//     0 <= vals[i] <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     edges represents a valid tree.

import "fmt"
import "sort"

type UnionFind struct {
    parent []int
    rank []int
}

func (this *UnionFind) Find(x int) int {
    if this.parent[x] != x {
        this.parent[x] = this.Find(this.parent[x])
    }
    return this.parent[x]
}

func (this *UnionFind) Set(x, y int) {
    xSet, ySet := this.Find(x), this.Find(y)

    if xSet != ySet {
        switch {
        case this.rank[xSet] < this.rank[ySet]:
            this.parent[xSet] = ySet
        case this.rank[xSet] > this.rank[ySet]:
            this.parent[ySet] = xSet
        default:
            this.parent[ySet] = xSet
            this.rank[xSet]++
        }
    }
}

func NewUnionFind(size int) *UnionFind {
    parent := make([]int, size) 
    for i := range parent {
        parent[i] = i
    }
    return &UnionFind{ parent, make([]int, size) }
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
    n := len(vals)
    adj := make([][]int, n)
    for i := range adj {
        adj[i] = []int{}
    }
    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], edge[1])
        adj[edge[1]] = append(adj[edge[1]], edge[0])
    }
    valuesToNodes := make(map[int][]int)
    for node := 0; node < n; node++ {
        valuesToNodes[vals[node]] = append(valuesToNodes[vals[node]], node)
    }
    keys := []int{}
    for value := range valuesToNodes {
        keys = append(keys, value)
    }
    sort.Ints(keys)
    dsu, goodPaths := NewUnionFind(n), 0
    for _, value := range keys {
        nodes := valuesToNodes[value]
        for _, node := range nodes {
            for _, neighbor := range adj[node] {
                if vals[node] >= vals[neighbor] {
                    dsu.Set(node, neighbor)
                }
            }
        }
        group := make(map[int]int)
        for _, node := range nodes {
            group[dsu.Find(node)]++
        }
        for _, size := range group {
            goodPaths += size * (size + 1) / 2
        }
    }
    return goodPaths
}

func numberOfGoodPaths1(vals []int, edges [][]int) int {
    n := len(vals)
    res, g := n, make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    fa, sz, id := make([]int, n), make([]int, n), make([]int, n)
    for i := range fa {
        fa[i], sz[i], id[i] = i, 1, i
    }
    find := func(x int) int {
        res := x
        for fa[res] != res {
            res = fa[res]
        }
        for fa[x] != res {
            fa[x], x = res, fa[x]
        }
        return res
    }
    sort.Slice(id, func(i, j int) bool {
        return vals[id[i]] < vals[id[j]]
    })
    for _, x := range id {
        vx, fx := vals[x], find(x)
        for _, y := range g[x] {
            fy := find(y)
            if fy == fx || vals[fy] > vx {
                continue
            }
            if vals[fy] == vx {
                res += sz[fy] * sz[fx]
                sz[fx] += sz[fy]
            }
            fa[fy] = fx
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/08/04/f9caaac15b383af9115c5586779dec5.png" />
    // Input: vals = [1,3,2,1,3], edges = [[0,1],[0,2],[2,3],[2,4]]
    // Output: 6
    // Explanation: There are 5 good paths consisting of a single node.
    // There is 1 additional good path: 1 -> 0 -> 2 -> 4.
    // (The reverse path 4 -> 2 -> 0 -> 1 is treated as the same as 1 -> 0 -> 2 -> 4.)
    // Note that 0 -> 2 -> 3 is not a good path because vals[2] > vals[0].
    fmt.Println(numberOfGoodPaths([]int{1,3,2,1,3},[][]int{{0,1},{0,2},{2,3},{2,4}})) // 6
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/08/04/149d3065ec165a71a1b9aec890776ff.png" />
    // Input: vals = [1,1,2,2,3], edges = [[0,1],[1,2],[2,3],[2,4]]
    // Output: 7
    // Explanation: There are 5 good paths consisting of a single node.
    // There are 2 additional good paths: 0 -> 1 and 2 -> 3.
    fmt.Println(numberOfGoodPaths([]int{1,1,2,2,3},[][]int{{0,1},{1,2},{2,3},{2,4}})) // 7
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/08/04/31705e22af3d9c0a557459bc7d1b62d.png" />
    // Input: vals = [1], edges = []
    // Output: 1
    // Explanation: The tree consists of only one node, so there is one good path.
    fmt.Println(numberOfGoodPaths([]int{1},[][]int{})) // 1

    fmt.Println(numberOfGoodPaths1([]int{1,3,2,1,3},[][]int{{0,1},{0,2},{2,3},{2,4}})) // 6
    fmt.Println(numberOfGoodPaths1([]int{1,1,2,2,3},[][]int{{0,1},{1,2},{2,3},{2,4}})) // 7
    fmt.Println(numberOfGoodPaths1([]int{1},[][]int{})) // 1
}