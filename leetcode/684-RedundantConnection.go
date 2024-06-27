package main

// 684. Redundant Connection
// In this problem, a tree is an undirected graph that is connected and has no cycles.

// You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added. 
// The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed. 
// The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.

// Return an edge that can be removed so that the resulting graph is a tree of n nodes. 
// If there are multiple answers, return the answer that occurs last in the input.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/02/reduntant1-1-graph.jpg" />
// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/02/reduntant1-2-graph.jpg" />
// Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
// Output: [1,4]
 
// Constraints:
//     n == edges.length
//     3 <= n <= 1000
//     edges[i].length == 2
//     1 <= ai < bi <= edges.length
//     ai != bi
//     There are no repeated edges.
//     The given graph is connected.

import "fmt"

// 并查集
func findRedundantConnection(edges [][]int) []int {
    // union find algorithm
    n := len(edges)
    parent, rank := make([]int, n+1), make([]int, n+1)
    find := func(x int) int {
        par := parent[x]
        for par != parent[par] {
            parent[par] = parent[parent[par]]
            par = parent[par]
        }
        return par
    }
    union := func(x int, y int) {
        parentX, parentY := find(x), find(y)
        if parentX == parentY {
            return
        }
        if rank[parentX] >= rank[parentY] {
            parent[parentY] = x
            rank[parentX]++
        } else {
            parent[parentX] = parentY
            rank[parentY]++
        }
    }
    for i := 1; i <= n; i++ {
        parent[i] = i
        rank[i] = 1
    }
    for _, edge := range edges {
        x, y := edge[0], edge[1]
        if find(x) == find(y) {
            return edge
        }
        union(x, y)
    }
    return []int{}
}

func findRedundantConnection1(edges [][]int) []int {
    nums := make([]int , len(edges)+1)
    for i := range nums {
        nums[i] = i
    }
    for _ , v := range edges {
        a, b := v[0], v[1]
        for nums[a] != a { a = nums[a] }
        for nums[b] != b { b = nums[b] }
        if a == b { return v }
        nums[a] = b
    }
    return []int{}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/02/reduntant1-1-graph.jpg" />
    // Input: edges = [[1,2],[1,3],[2,3]]
    // Output: [2,3]
    fmt.Println(findRedundantConnection([][]int{{1,2},{1,3},{2,3}})) // [2,3]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/02/reduntant1-2-graph.jpg" />
    // Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
    // Output: [1,4]
    fmt.Println(findRedundantConnection([][]int{{1,2},{2,3},{3,4},{1,4},{1,5}})) // [1,4]

    fmt.Println(findRedundantConnection1([][]int{{1,2},{1,3},{2,3}})) // [2,3]
    fmt.Println(findRedundantConnection1([][]int{{1,2},{2,3},{3,4},{1,4},{1,5}})) // [1,4]
}