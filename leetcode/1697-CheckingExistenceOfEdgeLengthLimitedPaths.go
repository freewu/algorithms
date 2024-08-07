package main

// 1697. Checking Existence of Edge Length Limited Paths
// An undirected graph of n nodes is defined by edgeList, where edgeList[i] = [ui, vi, disi] denotes an edge between nodes ui and vi with distance disi.
// Note that there may be multiple edges between two nodes.

// Given an array queries, where queries[j] = [pj, qj, limitj], 
// your task is to determine for each queries[j] whether there is a path between pj and qj such that each edge on the path has a distance strictly less than limitj .

// Return a boolean array answer, where answer.length == queries.length and the jth value of answer is true if there is a path for queries[j] is true, and false otherwise.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/08/h.png" />
// Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
// Output: [false,true]
// Explanation: The above figure shows the given graph. Note that there are two overlapping edges between 0 and 1 with distances 2 and 16.
// For the first query, between 0 and 1 there is no path where each distance is less than 2, thus we return false for this query.
// For the second query, there is a path (0 -> 1 -> 2) of two edges with distances less than 5, thus we return true for this query.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/08/q.png" />
// Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
// Output: [true,false]
// Explanation: The above figure shows the given graph.

// Constraints:
//     2 <= n <= 10^5
//     1 <= edgeList.length, queries.length <= 10^5
//     edgeList[i].length == 3
//     queries[j].length == 3
//     0 <= ui, vi, pj, qj <= n - 1
//     ui != vi
//     pj != qj
//     1 <= disi, limitj <= 10^9
//     There may be multiple edges between two nodes.

import "fmt"
import "sort"

// 并查集
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    var find func(x int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    var union func(x, y int)
    union = func(x, y int) {
        x = find(x)
        y = find(y)
        if x != y {
            parent[x] = y
        }
    }
    sort.Slice(edgeList, func(i, j int) bool {
        return edgeList[i][2] < edgeList[j][2]
    })
    for i, query := range queries {
        queries[i] = append(query, i)
    }
    sort.Slice(queries, func(i, j int) bool {
        return queries[i][2] < queries[j][2]
    })
    res, edgeIdx := make([]bool, len(queries)), 0
    for _, query := range queries {
        p, q, limit, idx := query[0], query[1], query[2], query[3]
        for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < limit {
            union(edgeList[edgeIdx][0], edgeList[edgeIdx][1])
            edgeIdx++
        }
        res[idx] = find(p) == find(q)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/08/h.png" />
    // Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
    // Output: [false,true]
    // Explanation: The above figure shows the given graph. Note that there are two overlapping edges between 0 and 1 with distances 2 and 16.
    // For the first query, between 0 and 1 there is no path where each distance is less than 2, thus we return false for this query.
    // For the second query, there is a path (0 -> 1 -> 2) of two edges with distances less than 5, thus we return true for this query.
    fmt.Println(distanceLimitedPathsExist(3, [][]int{{0,1,2},{1,2,4},{2,0,8},{1,0,16}}, [][]int{{0,1,2},{0,2,5}})) // [false,true]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/08/q.png" />
    // Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
    // Output: [true,false]
    // Explanation: The above figure shows the given graph.
    fmt.Println(distanceLimitedPathsExist(5, [][]int{{0,1,10},{1,2,5},{2,3,9},{3,4,13}}, [][]int{{0,4,14},{1,4,13}})) // [true,false]
}