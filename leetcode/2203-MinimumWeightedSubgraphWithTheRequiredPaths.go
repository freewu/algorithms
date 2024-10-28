package main

// 2203. Minimum Weighted Subgraph With the Required Paths
// You are given an integer n denoting the number of nodes of a weighted directed graph. 
// The nodes are numbered from 0 to n - 1.

// You are also given a 2D integer array edges where edges[i] = [fromi, toi, weighti] denotes 
// that there exists a directed edge from fromi to toi with weight weighti.

// Lastly, you are given three distinct integers src1, src2, 
// and dest denoting three distinct nodes of the graph.

// Return the minimum weight of a subgraph of the graph 
// such that it is possible to reach dest from both src1 and src2 via a set of edges of this subgraph. 
// In case such a subgraph does not exist, return -1.

// A subgraph is a graph whose vertices and edges are subsets of the original graph. 
// The weight of a subgraph is the sum of weights of its constituent edges.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/02/17/example1drawio.png" />
// Input: n = 6, edges = [[0,2,2],[0,5,6],[1,0,3],[1,4,5],[2,1,1],[2,3,3],[2,3,4],[3,4,2],[4,5,1]], src1 = 0, src2 = 1, dest = 5
// Output: 9
// Explanation:
// The above figure represents the input graph.
// The blue edges represent one of the subgraphs that yield the optimal answer.
// Note that the subgraph [[1,0,3],[0,5,6]] also yields the optimal answer. It is not possible to get a subgraph with less weight satisfying all the constraints.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/02/17/example2-1drawio.png" />
// Input: n = 3, edges = [[0,1,1],[2,1,1]], src1 = 0, src2 = 1, dest = 2
// Output: -1
// Explanation:
// The above figure represents the input graph.
// It can be seen that there does not exist any path from node 1 to node 2, hence there are no subgraphs satisfying all the constraints.

// Constraints:
//     3 <= n <= 10^5
//     0 <= edges.length <= 10^5
//     edges[i].length == 3
//     0 <= fromi, toi, src1, src2, dest <= n - 1
//     fromi != toi
//     src1, src2, and dest are pairwise distinct.
//     1 <= weight[i] <= 10^5

import "fmt"

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
    inf := int64(10000000000)
    type Pair struct {
        Key int64
        Value int
    }
    res := inf
    g, rg := make([][]Pair, n), make([][]Pair, n)
    l1, l2, ld := make([]int64, n), make([]int64, n), make([]int64, n)
    for i := 0; i < n; i++ {
        l1[i], l2[i], ld[i] = inf, inf, inf
    }
    l1[src1], l2[src2], ld[dest] = 0, 0, 0
    for _, edge := range edges {
        g[edge[0]] = append(g[edge[0]], Pair{int64(edge[1]), edge[2]})
        rg[edge[1]] = append(rg[edge[1]], Pair{int64(edge[0]), edge[2]})
    }
    min := func (x, y int64) int64 { if x < y { return x; }; return y; }
    dijkstras := func(node int, g [][]Pair, visited *[]int64) {
        queue := []Pair{ Pair{ 0, node } }
        for len(queue) > 0 {
            item := queue[0]
            queue = queue[1:]
            if (*visited)[item.Value] == item.Key {
                for _, c := range g[item.Value] {
                    if (*visited)[c.Key] > item.Key + int64(c.Value) {
                        (*visited)[c.Key] = item.Key + int64(c.Value)
                        queue = append(queue, Pair{(*visited)[int(c.Key)], int(c.Key)})
                    }
                }
            }
        }
    }
    dijkstras(src1, g, &l1)
    dijkstras(src2, g, &l2)
    dijkstras(dest, rg, &ld)
    if ld[src1] == inf || ld[src2] == inf { return -1 }
    for i := 0; i < n; i++ {
        res = min(res, ld[i] + l1[i] + l2[i])
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/17/example1drawio.png" />
    // Input: n = 6, edges = [[0,2,2],[0,5,6],[1,0,3],[1,4,5],[2,1,1],[2,3,3],[2,3,4],[3,4,2],[4,5,1]], src1 = 0, src2 = 1, dest = 5
    // Output: 9
    // Explanation:
    // The above figure represents the input graph.
    // The blue edges represent one of the subgraphs that yield the optimal answer.
    // Note that the subgraph [[1,0,3],[0,5,6]] also yields the optimal answer. It is not possible to get a subgraph with less weight satisfying all the constraints.
    edges1 := [][]int{{0,2,2},{0,5,6},{1,0,3},{1,4,5},{2,1,1},{2,3,3},{2,3,4},{3,4,2},{4,5,1}}
    fmt.Println(minimumWeight(6, edges1, 0, 1, 5)) // 9
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/17/example2-1drawio.png" />
    // Input: n = 3, edges = [[0,1,1],[2,1,1]], src1 = 0, src2 = 1, dest = 2
    // Output: -1
    // Explanation:
    // The above figure represents the input graph.
    // It can be seen that there does not exist any path from node 1 to node 2, hence there are no subgraphs satisfying all the constraints.
    edges2 := [][]int{{0,1,1},{2,1,1}}
    fmt.Println(minimumWeight(3, edges2, 0, 1, 2)) // -1
}