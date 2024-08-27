package main

// 2242. Maximum Score of a Node Sequence
// There is an undirected graph with n nodes, numbered from 0 to n - 1.

// You are given a 0-indexed integer array scores of length n where scores[i] denotes the score of node i. 
// You are also given a 2D integer array edges where edges[i] = [ai, bi] denotes 
// that there exists an undirected edge connecting nodes ai and bi.

// A node sequence is valid if it meets the following conditions:
//     There is an edge connecting every pair of adjacent nodes in the sequence.
//     No node appears more than once in the sequence.

// The score of a node sequence is defined as the sum of the scores of the nodes in the sequence.
// Return the maximum score of a valid node sequence with a length of 4. If no such sequence exists, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/04/15/ex1new3.png" />
// Input: scores = [5,2,9,8,4], edges = [[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]]
// Output: 24
// Explanation: The figure above shows the graph and the chosen node sequence [0,1,2,3].
// The score of the node sequence is 5 + 2 + 9 + 8 = 24.
// It can be shown that no other node sequence has a score of more than 24.
// Note that the sequences [3,1,2,0] and [1,0,2,3] are also valid and have a score of 24.
// The sequence [0,3,2,4] is not valid since no edge connects nodes 0 and 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/17/ex2.png" />
// Input: scores = [9,20,6,4,11,12], edges = [[0,3],[5,3],[2,4],[1,3]]
// Output: -1
// Explanation: The figure above shows the graph.
// There are no valid node sequences of length 4, so we return -1.

// Constraints:
//     n == scores.length
//     4 <= n <= 5 * 10^4
//     1 <= scores[i] <= 10^8
//     0 <= edges.length <= 5 * 10^4
//     edges[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi
//     There are no duplicate edges.

import "fmt"
import "sort"
import "slices"

func maximumScore(scores []int, edges [][]int) int {
    res, neighbors := -1, map[int][]int{}
    for _, e := range edges {
        neighbors[e[0]] = append(neighbors[e[0]], e[1])
        neighbors[e[1]] = append(neighbors[e[1]], e[0])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for v := range neighbors {
        sort.Slice(neighbors[v], func(i, j int) bool {
            return scores[neighbors[v][i]] > scores[neighbors[v][j]]
        })
        l := min(3, len(neighbors[v])) // sort all neighbors by decreasing score and limit to top 3
        neighbors[v] = neighbors[v][:l]
    }
    // iterate through all sequences that look like [leftNeighbor, l, r, rightNeighbor]
    for _, e := range edges {
        l, r := e[0], e[1]
        currScore := scores[l] + scores[r]
        // iterate through all leftNeighbors of l which aren't r
        for _, leftNeighbor := range neighbors[l] {
            if leftNeighbor == r {
                continue
            }
            // iterate through all rightNeighbors of r which aren't l or leftNeighbor
            for _, rightNeighbor := range neighbors[r] {
                if rightNeighbor == l || rightNeighbor == leftNeighbor {
                    continue
                }
                // add scores of all 4 nodes and compare to current max
                res = max(res, currScore + scores[leftNeighbor] + scores[rightNeighbor])
            }
        }
    }
    return res
}

// 四个点,而必须要相邻,可以枚举中间2个点,而这2个点需要相连=>即一条边的两个端点b,c=>所以可以枚举边
// 然后再分别寻找b,c两点的临界点中的最大值,且要满足不是另一个顶点,而不能和另一个顶点选的一样(题目要求4个点)
// - 比如寻找b的临界点的最大值, 不能是c,也不能和c选的一样=>所以保留最大的3个即可
func maximumScore1(scores []int, edges [][]int) int {
    res, n := -1, len(scores)
    type Pair struct{ to, score int }
    g := make([][]Pair, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        g[u] = append(g[u], Pair{v, scores[v]})
        g[v] = append(g[v], Pair{u, scores[u]})
    }
    for i, pairs := range g {
        if len(pairs) > 3 {
            slices.SortFunc(pairs, func(a, b Pair) int {
                return b.score - a.score
            })
            g[i] = pairs[:3]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, e := range edges {
        u, v := e[0], e[1]
        sum := scores[u] + scores[v]
        for _, pu := range g[u] { // 一个点至多保留3个邻点,所以最多循环3*3次
            for _, pv := range g[v] {
                if pu.to != v && pv.to != u && pu.to != pv.to {
                    res = max(res, sum + scores[pu.to] + scores[pv.to])
                }
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/04/15/ex1new3.png" />
    // Input: scores = [5,2,9,8,4], edges = [[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]]
    // Output: 24
    // Explanation: The figure above shows the graph and the chosen node sequence [0,1,2,3].
    // The score of the node sequence is 5 + 2 + 9 + 8 = 24.
    // It can be shown that no other node sequence has a score of more than 24.
    // Note that the sequences [3,1,2,0] and [1,0,2,3] are also valid and have a score of 24.
    // The sequence [0,3,2,4] is not valid since no edge connects nodes 0 and 3.
    fmt.Println(maximumScore([]int{5,2,9,8,4},[][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}})) // 24
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/17/ex2.png" />
    // Input: scores = [9,20,6,4,11,12], edges = [[0,3],[5,3],[2,4],[1,3]]
    // Output: -1
    // Explanation: The figure above shows the graph.
    // There are no valid node sequences of length 4, so we return -1.
    fmt.Println(maximumScore([]int{9,20,6,4,11,12},[][]int{{0,3},{5,3},{2,4},{1,3}})) // -1

    fmt.Println(maximumScore1([]int{5,2,9,8,4},[][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}})) // 24
    fmt.Println(maximumScore1([]int{9,20,6,4,11,12},[][]int{{0,3},{5,3},{2,4},{1,3}})) // -1
}