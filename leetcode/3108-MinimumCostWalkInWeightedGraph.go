package main

// 3108. Minimum Cost Walk in Weighted Graph
// There is an undirected weighted graph with n vertices labeled from 0 to n - 1.

// You are given the integer n and an array edges, where edges[i] = [ui, vi, wi] indicates that there is an edge between vertices ui and vi with a weight of wi.

// A walk on a graph is a sequence of vertices and edges. 
// The walk starts and ends with a vertex, and each edge connects the vertex that comes before it and the vertex that comes after it. 
// It's important to note that a walk may visit the same edge or vertex more than once.

// The cost of a walk starting at node u and ending at node v is defined as the bitwise AND of the weights of the edges traversed during the walk. 
// In other words, if the sequence of edge weights encountered during the walk is w0, w1, w2, ..., wk, 
// then the cost is calculated as w0 & w1 & w2 & ... & wk, where & denotes the bitwise AND operator.

// You are also given a 2D array query, where query[i] = [si, ti]. 
// For each query, you need to find the minimum cost of the walk starting at vertex si and ending at vertex ti. 
// If there exists no such walk, the answer is -1.

// Return the array answer, where answer[i] denotes the minimum cost of a walk for query i.

// Example 1:
// Input: n = 5, edges = [[0,1,7],[1,3,7],[1,2,1]], query = [[0,3],[3,4]]
// Output: [1,-1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/01/31/q4_example1-1.png"/>
// To achieve the cost of 1 in the first query, we need to move on the following edges: 0->1 (weight 7), 1->2 (weight 1), 2->1 (weight 1), 1->3 (weight 7).
// In the second query, there is no walk between nodes 3 and 4, so the answer is -1.

// Example 2:
// Input: n = 3, edges = [[0,2,7],[0,1,15],[1,2,6],[1,2,1]], query = [[1,2]]
// Output: [0]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/01/31/q4_example2e.png" />
// To achieve the cost of 0 in the first query, we need to move on the following edges: 1->2 (weight 1), 2->1 (weight 6), 1->2 (weight 1).

// Constraints:
//     2 <= n <= 10^5
//     0 <= edges.length <= 10^5
//     edges[i].length == 3
//     0 <= ui, vi <= n - 1
//     ui != vi
//     0 <= wi <= 10^5
//     1 <= query.length <= 10^5
//     query[i].length == 2
//     0 <= si, ti <= n - 1
//     si != ti

import "fmt"

func minimumCost(n int, edges, query [][]int) []int {
    type Edge struct{ To, Weight int }
    g := make([][]Edge, n)
    for _, e := range edges {
        x, y, w := e[0], e[1], e[2]
        g[x] = append(g[x], Edge{y, w})
        g[y] = append(g[y], Edge{x, w})
    }
    ids := make([]int, n) 
    for i := range ids {
        ids[i] = -1
    }
    ccAnd := []int{} 
    var dfs func(int) int
    dfs = func(x int) int {
        ids[x] = len(ccAnd) 
        and := -1
        for _, e := range g[x] {
            and &= e.Weight
            if ids[e.To] < 0 {
                and &= dfs(e.To)
            }
        }
        return and
    }
    for i, id := range ids {
        if id < 0 { // 没有访问过
            ccAnd = append(ccAnd, dfs(i))
        }
    }
    res := make([]int, len(query))
    for i, q := range query {
        s, t := q[0], q[1]
        if s == t { continue }
        if ids[s] != ids[t] {
            res[i] = -1
        } else {
            res[i] = ccAnd[ids[s]]
        }
    }
    return res
}

type UnionFind struct {
	arr []int
}

func Constructor(n int) UnionFind {
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = -1
    }
    return UnionFind{
        arr: arr,
    }
}

func (uf *UnionFind) Find(x int) int {
    for uf.arr[x] >= 0 {
        x = uf.arr[x]
    }
    return x
}

func (uf *UnionFind) Union(x, y int) {
    px, py := uf.Find(x), uf.Find(y)
    if px == py {
        return
    }
    if uf.arr[px] < uf.arr[py] {
        uf.arr[px] += uf.arr[py]
        uf.arr[py] = px
    } else {
        uf.arr[py] += uf.arr[px]
        uf.arr[px] = py
    }
}

func minimumCost1(n int, edges [][]int, query [][]int) []int {
    uf := Constructor(n)
    for i := range edges {
        from, to := edges[i][0], edges[i][1]
        uf.Union(from, to)
    }
    memo := make(map[int]int)
    for i := range edges {
        from := edges[i][0]
        pfrom := uf.Find(from)
        if _, ok := memo[pfrom]; ok {
            memo[pfrom] &= edges[i][2]
        } else {
            memo[pfrom] = edges[i][2]
        }
    }
    m := len(query)
    res := make([]int, m)
    for i := range query {
        from, to := query[i][0], query[i][1]
        pfrom, pto := uf.Find(from), uf.Find(to)
        if pfrom == pto {
            res[i] = memo[pfrom]
        } else {
            res[i] = -1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5, edges = [[0,1,7],[1,3,7],[1,2,1]], query = [[0,3],[3,4]]
    // Output: [1,-1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/01/31/q4_example1-1.png"/>
    // To achieve the cost of 1 in the first query, we need to move on the following edges: 0->1 (weight 7), 1->2 (weight 1), 2->1 (weight 1), 1->3 (weight 7).
    // In the second query, there is no walk between nodes 3 and 4, so the answer is -1.
    fmt.Println(minimumCost(5,[][]int{{0,1,7},{1,3,7},{1,2,1}}, [][]int{{0,3},{3,4}})) // [1,-1]
    // Example 2:
    // Input: n = 3, edges = [[0,2,7],[0,1,15],[1,2,6],[1,2,1]], query = [[1,2]]
    // Output: [0]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/01/31/q4_example2e.png" />
    // To achieve the cost of 0 in the first query, we need to move on the following edges: 1->2 (weight 1), 2->1 (weight 6), 1->2 (weight 1).
    fmt.Println(minimumCost(5,[][]int{{0,2,7},{0,1,15},{1,2,6},{1,2,1}}, [][]int{{1,2}})) // [0]

    fmt.Println(minimumCost1(5,[][]int{{0,1,7},{1,3,7},{1,2,1}}, [][]int{{0,3},{3,4}})) // [1,-1]
    fmt.Println(minimumCost1(5,[][]int{{0,2,7},{0,1,15},{1,2,6},{1,2,1}}, [][]int{{1,2}})) // [0]
}