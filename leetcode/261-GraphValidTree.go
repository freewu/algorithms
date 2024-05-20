package main

// 261. Graph Valid Tree
// You have a graph of n nodes labeled from 0 to n - 1. 
// You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates 
// that there is an undirected edge between nodes ai and bi in the graph.

// Return true if the edges of the given graph make up a valid tree, and false otherwise.

// Example 1
// <img src="https://assets.leetcode.com/uploads/2021/03/12/tree1-graph.jpg" />
// Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/12/tree1-graph.jpg" />
// Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
// Output: false

// Constraints:
//     1 <= n <= 2000
//     0 <= edges.length <= 5000
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     There are no self-loops or repeated edges.

import "fmt"

// dfs 
func validTree(n int, edges [][]int) bool {
    if len(edges) != n - 1 {  // 树的变数等于节点数-1；没有指向根节点的边
        return false
    }
    g := make([][]int, n) // 构造邻接表
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    num, visited := 0,make([]bool, n)
    var dfs func(v int)
    dfs = func(v int) { // 统计遍历的节点的数量
        if visited[v] { return }
        visited[v] = true // 标记
        num++
        for _, a := range g[v] {
            if !visited[a] {
                dfs(a)
            }
        }
    }
    dfs(0) // 从一个节点开始遍历
    return num == n // num == n,说明图是联通的
}

// 并查集
func validTree1(n int, edges [][]int) bool {
    uf := NewUnionFind(n)
    for _, e := range edges {
        x, y := e[0], e[1]
        if uf.Connected(x, y) { // 不存在环
            return false
        }
        uf.Union(x, y) // 连通图
    }
    return uf.Count() == 1 // 联通图
}

type UnionFind struct {
    parent []int // 指向的父节点，一开始默认指向自己
    count int // 联通分量的个数
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    return &UnionFind{parent, n}
}

func (u *UnionFind) Union(p, q int) {
    rootP := u.parent[p]
    rootQ := u.parent[q]
    if rootP == rootQ { return; }
    u.parent[rootQ] = rootP
    u.count--
}

func (u *UnionFind) Find(x int) int {
    if x != u.parent[x] { // 路径压缩， 最终都指向跟节点
        u.parent[x] = u.Find(u.parent[x])
    }
    return u.parent[x]
}

func (u *UnionFind) Connected(p, q int) bool {
    return u.Find(p) == u.Find(q)
}

func (u *UnionFind) Count() int {
    return u.count
}

func validTree2(n int, edges [][]int) bool {
    if len(edges) < n-1{
        return false
    }
    g := make([][]int,n)
    for _,edge := range edges{
        x,y := edge[0],edge[1]
        g[x] = append(g[x],y)
        g[y] = append(g[y],x)
    }
    visited, count, bad := make([]bool,n), 1, false
    var dfs func(x,fa int)
    dfs = func(x,fa int){
        visited[x] = true
        for _,y := range g[x] {
            if !bad && y != fa {
                count++
                if !visited[y] {
                    dfs(y,x)
                } else {
                    bad = true
                    return
                }
            }
        }
    }
    dfs(0,-1)
    return !bad && count == n
}

func main() {
    // Example 1
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/tree1-graph.jpg" />
    // Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
    // Output: true
    fmt.Println(validTree(5,[][]int{{0,1}, {0,2}, {0,3}, {0,4}})) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/tree1-graph.jpg" />
    // Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
    // Output: false
    fmt.Println(validTree(5,[][]int{{0,1}, {1,2}, {2,3}, {1,3}, {1,4}})) // false

    fmt.Println(validTree1(5,[][]int{{0,1}, {0,2}, {0,3}, {0,4}})) // true
    fmt.Println(validTree1(5,[][]int{{0,1}, {1,2}, {2,3}, {1,3}, {1,4}})) // false

    fmt.Println(validTree2(5,[][]int{{0,1}, {0,2}, {0,3}, {0,4}})) // true
    fmt.Println(validTree2(5,[][]int{{0,1}, {1,2}, {2,3}, {1,3}, {1,4}})) // false
}