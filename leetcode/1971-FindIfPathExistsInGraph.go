package main

// 1971. Find if Path Exists in Graph
// There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive). 
// The edges in the graph are represented as a 2D integer array edges, 
// where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi. 
// Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.

// You want to determine if there is a valid path that exists from vertex source to vertex destination.

// Given edges and the integers n, source, and destination, 
// return true if there is a valid path from source to destination, or false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/14/validpath-ex1.png" />
// Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
// Output: true
// Explanation: There are two paths from vertex 0 to vertex 2:
// - 0 → 1 → 2
// - 0 → 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/14/validpath-ex2.png" />
// Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
// Output: false
// Explanation: There is no path from vertex 0 to vertex 5.

// Constraints:
//     1 <= n <= 2 * 10^5
//     0 <= edges.length <= 2 * 10^5
//     edges[i].length == 2
//     0 <= ui, vi <= n - 1
//     ui != vi
//     0 <= source, destination <= n - 1
//     There are no duplicate edges.
//     There are no self edges.

import "fmt"

// bfs
func validPath(n int, edges [][]int, source int, destination int) bool {
    marked := make([]bool, n)
    adj := make([][]int, n)
    for _, edge := range edges{
        v, w := edge[0], edge[1]
        adj[v] = append(adj[v], w)
        adj[w] = append(adj[w], v)
    }
    queue := []int{source}
    for len(queue) > 0{
        pop := queue[0]
        queue = queue[1:]
        if pop == destination {
            return true
        }
        marked[pop] = true
        for _, v := range adj[pop]{
            if !marked[v]{
                queue = append(queue, v)
            }
        }
    }
    return false
}

func validPath1(n int, edges [][]int, source int, destination int) bool {
    root := make([]int, n)
    for i := range root {
        root[i] = i
    }
    var findRoot func (i int) int
    findRoot = func (i int) int {
        if root[i] == i {
            return i
        } else {
            root[i] = findRoot(root[i])
            return root[i]
        }
    }
    join := func (p, q int) {
        p = findRoot(p)
        q = findRoot(q)
        if p != q {
            root[q] = p
        }
    }
    isConnect := func (p, q int) bool {
        p = findRoot(p)
        q = findRoot(q)
        return p == q
    }
    for _, e := range edges {
        join(e[0], e[1])
    }
    return isConnect(source, destination)
}

// 并查集
func validPath2(n int, edges [][]int, source int, destination int) bool {
    uf := make([]int, n)
    for i := range uf {
        uf[i] = i
    }
    find := func(x int) int {
        for x != uf[x] {
            x, uf[x] = uf[x], uf[uf[x]]
        }
        return x
    }
    union := func(x, y int) {
        x, y = find(x), find(y)
        uf[x] = y
    }
    for _, e := range edges {
        u, v := e[0], e[1]
        union(u, v)
    }
    return find(source) == find(destination)
}

func main() {
    // Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
    // Output: true
    // Explanation: There are two paths from vertex 0 to vertex 2:
    // - 0 → 1 → 2
    // - 0 → 2
    fmt.Println(validPath(3,[][]int{{0,1},{1,2},{2,0}},0, 2)) // true
    // Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
    // Output: false
    // Explanation: There is no path from vertex 0 to vertex 5.
    fmt.Println(validPath(6,[][]int{{0,1},{0,2},{3,5},{5,4},{4,3}},0, 5)) // false

    fmt.Println(validPath1(3,[][]int{{0,1},{1,2},{2,0}},0, 2)) // true
    fmt.Println(validPath1(6,[][]int{{0,1},{0,2},{3,5},{5,4},{4,3}},0, 5)) // false

    fmt.Println(validPath2(3,[][]int{{0,1},{1,2},{2,0}},0, 2)) // true
    fmt.Println(validPath2(6,[][]int{{0,1},{0,2},{3,5},{5,4},{4,3}},0, 5)) // false
}