package main

// 2316. Count Unreachable Pairs of Nodes in an Undirected Graph
// You are given an integer n. 
// There is an undirected graph with n nodes, numbered from 0 to n - 1. 
// You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge connecting nodes ai and bi.

// Return the number of pairs of different nodes that are unreachable from each other.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/05/05/tc-3.png" />
// Input: n = 3, edges = [[0,1],[0,2],[1,2]]
// Output: 0
// Explanation: There are no pairs of nodes that are unreachable from each other. Therefore, we return 0.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/05/05/tc-2.png" />
// Input: n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]
// Output: 14
// Explanation: There are 14 pairs of nodes that are unreachable from each other:
// [[0,1],[0,3],[0,6],[1,2],[1,3],[1,4],[1,5],[2,3],[2,6],[3,4],[3,5],[3,6],[4,6],[5,6]].
// Therefore, we return 14.

// Constraints:
//     1 <= n <= 10^5
//     0 <= edges.length <= 2 * 10^5
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     There are no repeated edges.

import "fmt"

type unionFind struct {
    parent, rank []int
}

func (u *unionFind) Find(x int) int {
    if u.parent[x] != x {
        u.parent[x] = u.Find(u.parent[x])
    }
    return u.parent[x]
}

func (u *unionFind) Set(x, y int) {
    xSet, ySet := u.Find(x), u.Find(y)
    
    switch {
    case xSet == ySet:
        return
    case u.rank[xSet] < u.rank[ySet]:
        u.parent[xSet] = ySet
    case u.rank[xSet] > u.rank[ySet]:
        u.parent[ySet] = xSet
    default:
        u.parent[ySet] = xSet
        u.rank[xSet]++
    }
}

func newUnionFind(size int) *unionFind {
    parent, rank := make([]int, size), make([]int, size)
    for i := 0; i < size; i++ {
        parent[i] = i
    }
    return &unionFind{parent, rank}
}

func countPairs(n int, edges [][]int) int64 {
    dsu := newUnionFind(n)
    for _, edge := range edges {
        dsu.Set(edge[0], edge[1])
    }
    componentSize := make(map[int]int)
    for i := 0; i < n; i++ {
        componentSize[dsu.Find(i)]++
    }
    numberOfPaths, remainingNodes := int64(0), n
    for _, cSize := range componentSize {
        numberOfPaths += int64(cSize * (remainingNodes - cSize))
        remainingNodes -= cSize
    }
    return numberOfPaths
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/05/05/tc-3.png" />
    // Input: n = 3, edges = [[0,1],[0,2],[1,2]]
    // Output: 0
    // Explanation: There are no pairs of nodes that are unreachable from each other. Therefore, we return 0.
    fmt.Println(countPairs(3,[][]int{{0,1},{0,2},{1,2}})) // 0
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/05/05/tc-2.png" />
    // Input: n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]
    // Output: 14
    // Explanation: There are 14 pairs of nodes that are unreachable from each other:
    // [[0,1],[0,3],[0,6],[1,2],[1,3],[1,4],[1,5],[2,3],[2,6],[3,4],[3,5],[3,6],[4,6],[5,6]].
    // Therefore, we return 14.
    fmt.Println(countPairs(7,[][]int{{0,2},{0,5},{2,4},{1,6},{5,4}})) // 14
}