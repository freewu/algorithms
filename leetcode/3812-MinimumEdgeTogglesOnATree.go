package main

// 3812. Minimum Edge Toggles on a Tree
// You are given an undirected tree with n nodes, numbered from 0 to n - 1. 
// It is represented by a 2D integer array edges​​​​​​​ of length n - 1, w
// here edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// You are also given two binary strings start and target of length n. 
// For each node x, start[x] is its initial color and target[x] is its desired color.

// In one operation, you may pick an edge with index i and toggle both of its endpoints. 
// That is, if the edge is [u, v], then the colors of nodes u and v each flip from '0' to '1' or from '1' to '0'.

// Return an array of edge indices whose operations transform start into target. 
// Among all valid sequences with minimum possible length, return the edge indices in increasing​​​​​​​ order.

// If it is impossible to transform start into target, return an array containing a single element equal to -1.

// Example 1:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2025/12/18/example1.png" />
// Input: n = 3, edges = [[0,1],[1,2]], start = "010", target = "100"
// Output: [0]
// Explanation:
// Toggle edge with index 0, which flips nodes 0 and 1.
// ​​​​​​​The string changes from "010" to "100", matching the target.

// Example 2:
// ​​​​​<img src="https://assets.leetcode.com/uploads/2025/12/18/example2.png" />
// Input: n = 7, edges = [[0,1],[1,2],[2,3],[3,4],[3,5],[1,6]], start = "0011000", target = "0010001"
// Output: [1,2,5]
// Explanation:
// Toggle edge with index 1, which flips nodes 1 and 2.
// Toggle edge with index 2, which flips nodes 2 and 3.
// Toggle edge with index 5, which flips nodes 1 and 6.
// After these operations, the resulting string becomes "0010001", which matches the target.

// Example 3:
// ​​​​​<img src="https://assets.leetcode.com/uploads/2025/12/18/example3.png" />
// Input: n = 2, edges = [[0,1]], start = "00", target = "01"
// Output: [-1]
// Explanation:
// There is no sequence of edge toggles that transforms "00" into "01". Therefore, we return [-1].

// Constraints:
//     2 <= n == start.length == target.length <= 10^5
//     edges.length == n - 1
//     edges[i] = [ai, bi]
//     0 <= ai, bi < n
//     start[i] is either '0' or '1'.
//     target[i] is either '0' or '1'.
//     The input is generated such that edges represents a valid tree.

import "fmt"
import "sort"

func minimumFlips(n int, edges [][]int, start, target string) []int {
    type Edge struct{ to, i int }
    res, graph := make([]int,0), make([][]Edge, n)
    for i, e := range edges {
        x, y := e[0], e[1]
        graph[x] = append(graph[x], Edge{y, i})
        graph[y] = append(graph[y], Edge{x, i})
    }
    // 返回是否需要翻转 x-fa 这条边
    var dfs func(int, int) bool
    dfs = func(x, fa int) bool {
        rev := start[x] != target[x] // x-fa 是否要翻转
        for _, e := range graph[x] {
            y := e.to
            if y != fa && dfs(y, x) {
                res = append(res, e.i) // 需要翻转 y-x
                rev = !rev // x 被翻转了
            }
        }
        return rev
    }
    if dfs(0, -1) { // 只剩下一个根节点需要翻转，无法操作
        return []int{-1}
    }
    sort.Ints(res) // 按顺序输出
    return res
}

func minimumFlips1(n int, edges [][]int, start string, target string) []int {
    c0s, c0t := 0, 0
    for i := 0; i < n; i++ { // 统计出现的 0 的次数
        if start[i] == '0'  { c0s++ }
        if target[i] == '0' { c0t++ }
    }
    if ((c0s + c0t) & 1) == 1 { return []int{ -1 } }
    type Edge struct {
        to  int
        index int
    }
    adj := make([][]Edge, n)
    for i := 0; i < len(edges); i++ {
        u, v := edges[i][0], edges[i][1]
        adj[u] = append(adj[u], Edge{v, i})
        adj[v] = append(adj[v], Edge{u, i})
    }
    res := make([]int, 0)
    var dfs func(u, par int) int
    dfs = func(u, par int) int {
        sub := 0
        if start[u] != target[u] {
            sub = 1
        }
        for _, e := range adj[u] {
            v := e.to
            if v == par { continue }
            child := dfs(v, u)
            if child != 0 {
                res = append(res, e.index)
            }
            sub ^= child
        }
        return sub
    }
    dfs(0, -1)
    sort.Ints(res)
    return res
}

func main() {
    // Example 1:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2025/12/18/example1.png" />
    // Input: n = 3, edges = [[0,1],[1,2]], start = "010", target = "100"
    // Output: [0]
    // Explanation:
    // Toggle edge with index 0, which flips nodes 0 and 1.
    // ​​​​​​​The string changes from "010" to "100", matching the target.
    fmt.Println(minimumFlips(3, [][]int{{0,1},{1,2}}, "010", "100")) // [0]
    // Example 2:
    // ​​​​​<img src="https://assets.leetcode.com/uploads/2025/12/18/example2.png" />
    // Input: n = 7, edges = [[0,1],[1,2],[2,3],[3,4],[3,5],[1,6]], start = "0011000", target = "0010001"
    // Output: [1,2,5]
    // Explanation:
    // Toggle edge with index 1, which flips nodes 1 and 2.
    // Toggle edge with index 2, which flips nodes 2 and 3.
    // Toggle edge with index 5, which flips nodes 1 and 6.
    // After these operations, the resulting string becomes "0010001", which matches the target.
    fmt.Println(minimumFlips(7, [][]int{{0,1},{1,2},{2,3},{3,4},{3,5},{1,6}}, "0011000", "0010001")) // [1,2,5]
    // Example 3:
    // ​​​​​<img src="https://assets.leetcode.com/uploads/2025/12/18/example3.png" />
    // Input: n = 2, edges = [[0,1]], start = "00", target = "01"
    // Output: [-1]
    // Explanation:
    // There is no sequence of edge toggles that transforms "00" into "01". Therefore, we return [-1].
    fmt.Println(minimumFlips(2, [][]int{{0,1}}, "00", "01")) // [-1]

    fmt.Println(minimumFlips1(3, [][]int{{0,1},{1,2}}, "010", "100")) // [0]
    fmt.Println(minimumFlips1(7, [][]int{{0,1},{1,2},{2,3},{3,4},{3,5},{1,6}}, "0011000", "0010001")) // [1,2,5]
    fmt.Println(minimumFlips1(2, [][]int{{0,1}}, "00", "01")) // [-1]
}