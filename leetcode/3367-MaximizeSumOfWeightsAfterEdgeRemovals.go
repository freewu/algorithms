package main

// 3367. Maximize Sum of Weights after Edge Removals
// There exists an undirected tree with n nodes numbered 0 to n - 1. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ui, vi, wi] indicates that there is an edge between nodes ui and vi with weight wi in the tree.

// Your task is to remove zero or more edges such that:
//     1. Each node has an edge with at most k other nodes, where k is given.
//     2. The sum of the weights of the remaining edges is maximized.

// Return the maximum possible sum of weights for the remaining edges after making the necessary removals.

// Example 1:
// Input: edges = [[0,1,4],[0,2,2],[2,3,12],[2,4,6]], k = 2
// Output: 22
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/30/test1drawio.png" />
// Node 2 has edges with 3 other nodes. We remove the edge [0, 2, 2], ensuring that no node has edges with more than k = 2 nodes.
// The sum of weights is 22, and we can't achieve a greater sum. Thus, the answer is 22.

// Example 2:
// Input: edges = [[0,1,5],[1,2,10],[0,3,15],[3,4,20],[3,5,5],[0,6,10]], k = 3
// Output: 65
// Explanation:
// Since no node has edges connecting it to more than k = 3 nodes, we don't remove any edges.
// The sum of weights is 65. Thus, the answer is 65.

// Constraints:
//     2 <= n <= 10^5
//     1 <= k <= n - 1
//     edges.length == n - 1
//     edges[i].length == 3
//     0 <= edges[i][0] <= n - 1
//     0 <= edges[i][1] <= n - 1
//     1 <= edges[i][2] <= 10^6
//     The input is generated such that edges form a valid tree.

import "fmt"
import "sort"

func maximizeSumOfWeights(edges [][]int, k int) int64 {
    type Edge struct{ vertix, weight int }
    type Tree [][]Edge
    t := make(Tree, len(edges) + 1)
    for _, e := range edges {
        x, y, w := e[0], e[1], e[2]
        t[x], t[y] = append(t[x], Edge{ y, w}), append(t[y], Edge{ x, w})
    }
    var dfs func(i, p, k int) (int64, int64) 
    dfs = func(i, p, k int) (int64, int64) {
        res, diffs := [2]int64{}, []int64{}
        for _, e := range t[i] {
            if e.vertix == p { continue }
            exclusive, inclusive := dfs(e.vertix, i, k)
            res[1] += exclusive
            diffs = append(diffs, inclusive - exclusive + int64(e.weight))
        }
        res[0] = res[1]
        sort.Slice(diffs, func(i, j int) bool {
            return diffs[i] > diffs[j]
        })
        for i, d := range diffs {
            if d < 0 || i == k { break }
            res[0] += d
            if i < k - 1 {
                res[1] += d
            }
        }
        return res[0], res[1]
    }
    res, _ := dfs(0, -1, k)
    return res
}

func maximizeSumOfWeights1(edges [][]int, k int) int64 {
    type Pair struct { vertix, weight int }
    sum, simple, n := 0, true, len(edges) + 1
    graph := make([][]Pair, n)
    for _, e := range edges {
        u, v, wt := e[0], e[1], e[2]
        graph[u] = append(graph[u], Pair{ v, wt })
        graph[v] = append(graph[v], Pair{ u, wt })
        sum += wt
    }
    for _, v := range graph {
        if len(v) > k {
            simple = false
            break
        }
    }
    if simple { return int64(sum) }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x, fa int) (int, int) 
    dfs = func(x, fa int) (int, int) {
        inc, notChoose := []int{}, 0
        for _, e := range graph[x] {
            y := e.vertix
            if y == fa { continue }
            nc, c := dfs(y, x)
            notChoose += nc
            if d := c + e.weight - nc; d > 0 {
                inc = append(inc, d)
            }
        }
        sort.Slice(inc, func(i, j int) bool {
            return inc[i] > inc[j]
        })
        for i := 0; i < min(len(inc), k - 1); i++ {
            notChoose += inc[i]
        }
        choose := notChoose
        if len(inc) >= k {
            notChoose += inc[k-1]
        }
        return notChoose, choose
    }
    return int64(max(dfs(0, -1)))
}

func main() {
    // Example 1:
    // Input: edges = [[0,1,4],[0,2,2],[2,3,12],[2,4,6]], k = 2
    // Output: 22
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/30/test1drawio.png" />
    // Node 2 has edges with 3 other nodes. We remove the edge [0, 2, 2], ensuring that no node has edges with more than k = 2 nodes.
    // The sum of weights is 22, and we can't achieve a greater sum. Thus, the answer is 22.
    fmt.Println(maximizeSumOfWeights([][]int{{0,1,4},{0,2,2},{2,3,12},{2,4,6}}, 2)) // 22
    // Example 2:
    // Input: edges = [[0,1,5],[1,2,10],[0,3,15],[3,4,20],[3,5,5],[0,6,10]], k = 3
    // Output: 65
    // Explanation:
    // Since no node has edges connecting it to more than k = 3 nodes, we don't remove any edges.
    // The sum of weights is 65. Thus, the answer is 65.
    fmt.Println(maximizeSumOfWeights([][]int{{0,1,5},{1,2,10},{0,3,15},{3,4,20},{3,5,5},{0,6,10}}, 3)) // 65

    fmt.Println(maximizeSumOfWeights1([][]int{{0,1,4},{0,2,2},{2,3,12},{2,4,6}}, 2)) // 22
    fmt.Println(maximizeSumOfWeights1([][]int{{0,1,5},{1,2,10},{0,3,15},{3,4,20},{3,5,5},{0,6,10}}, 3)) // 65
}