package main

// 1761. Minimum Degree of a Connected Trio in a Graph
// You are given an undirected graph. 
// You are given an integer n which is the number of nodes in the graph and an array edges, 
// where each edges[i] = [ui, vi] indicates that there is an undirected edge between ui and vi.

// A connected trio is a set of three nodes where there is an edge between every pair of them.

// The degree of a connected trio is the number of edges where one endpoint is in the trio, and the other is not.

// Return the minimum degree of a connected trio in the graph, or -1 if the graph has no connected trios.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/26/trios1.png" />
// Input: n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
// Output: 3
// Explanation: There is exactly one trio, which is [1,2,3]. 
// The edges that form its degree are bolded in the figure above.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/26/trios2.png" />
// Input: n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
// Output: 0
// Explanation: There are exactly three trios:
// 1) [1,4,3] with degree 0.
// 2) [2,5,6] with degree 2.
// 3) [5,6,7] with degree 2.

// Constraints:
//     2 <= n <= 400
//     edges[i].length == 2
//     1 <= edges.length <= n * (n-1) / 2
//     1 <= ui, vi <= n
//     ui != vi
//     There are no repeated edges.

import "fmt"

type Set struct {
    items map[int]bool
}

func (s *Set) Add(item int) *Set {
    if s == nil { s = &Set{} }
    if s.items == nil {
        s.items = make(map[int]bool)
    }
    if _, ok := s.items[item]; !ok {
        s.items[item] = true
    }
    return s
}

func (s *Set) Contains(item int) bool {
    if s == nil { return false }
    _, ok := s.items[item]
    return ok
}

func (s *Set) Size() int {
    return len(s.items)
}

func minTrioDegree(n int, edges [][]int) int {
    res, mp := 1 << 31, make(map[int]*Set)
    for _, v := range edges {
        mp[v[0]] = mp[v[0]].Add(v[1])
        mp[v[1]] = mp[v[1]].Add(v[0])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            for k := j + 1; k <= n; k++ {
                if ok := mp[i].Contains(j); !ok { continue }
                if ok := mp[i].Contains(k); !ok { continue }
                if ok := mp[j].Contains(k); !ok { continue }
                a, b, c := mp[i].Size(), mp[j].Size(), mp[k].Size()
                res = min(res, a  + b + c - 6)
            }
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func minTrioDegree1(n int, edges [][]int) int {
    connected, degree := [512][512]bool{}, [512]int{}
    for i := 1; i <= n; i++ {
        degree[i] = 0
        for j := 1; j <= n; j++ { connected[i][j] = false }
    }
    for _, e := range edges {
        u, v := e[0], e[1]
        degree[u], degree[v], connected[u][v], connected[v][u] = degree[u] + 1, degree[v] + 1, true, true
    }
    res := 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := i + 1; j <= n; j++ {
            if !connected[i][j] { continue }
            for k := j + 1; k <= n; k++ {
                if connected[i][k] && connected[j][k] {
                    res = min(res, degree[i] + degree[j] + degree[k] - 6)
                }
            }
        }
    }
    if res == 1 << 31 { return -1 }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/01/26/trios1.png" />
    // Input: n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
    // Output: 3
    // Explanation: There is exactly one trio, which is [1,2,3]. 
    // The edges that form its degree are bolded in the figure above.
    fmt.Println(minTrioDegree(6, [][]int{{1,2},{1,3},{3,2},{4,1},{5,2},{3,6}})) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/01/26/trios2.png" />
    // Input: n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
    // Output: 0
    // Explanation: There are exactly three trios:
    // 1) [1,4,3] with degree 0.
    // 2) [2,5,6] with degree 2.
    // 3) [5,6,7] with degree 2.
    fmt.Println(minTrioDegree(7, [][]int{{1,3},{4,1},{4,3},{2,5},{5,6},{6,7},{7,5},{2,6}})) // 0

    fmt.Println(minTrioDegree1(6, [][]int{{1,2},{1,3},{3,2},{4,1},{5,2},{3,6}})) // 3
    fmt.Println(minTrioDegree1(7, [][]int{{1,3},{4,1},{4,3},{2,5},{5,6},{6,7},{7,5},{2,6}})) // 0
}