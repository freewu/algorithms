package main

// 3486. Longest Special Path II
// You are given an undirected tree rooted at node 0, with n nodes numbered from 0 to n - 1. 
// This is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi, lengthi] indicates an edge between nodes ui and vi with length lengthi. 
// You are also given an integer array nums, where nums[i] represents the value at node i.

// A special path is defined as a downward path from an ancestor node to a descendant node in which all node values are distinct, 
// except for at most one value that may appear twice.

// Return an array result of size 2, where result[0] is the length of the longest special path, 
// and result[1] is the minimum number of nodes in all possible longest special paths.

// Example 1:
// Input: edges = [[0,1,1],[1,2,3],[1,3,1],[2,4,6],[4,7,2],[3,5,2],[3,6,5],[6,8,3]], nums = [1,1,0,3,1,2,1,1,0]
// Output: [9,3]
// Explanation:
// In the image below, nodes are colored by their corresponding values in nums.
// <img src="https://assets.leetcode.com/uploads/2025/02/18/e1.png" />
// The longest special paths are 1 -> 2 -> 4 and 1 -> 3 -> 6 -> 8, both having a length of 9. 
// The minimum number of nodes across all longest special paths is 3.

// Example 2:
// Input: edges = [[1,0,3],[0,2,4],[0,3,5]], nums = [1,1,0,2]
// Output: [5,2]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/02/18/e2.png" />
// The longest path is 0 -> 3 consisting of 2 nodes with a length of 5.

// Constraints:
//     2 <= n <= 5 * 10^4
//     edges.length == n - 1
//     edges[i].length == 3
//     0 <= ui, vi < n
//     1 <= lengthi <= 10^3
//     nums.length == n
//     0 <= nums[i] <= 5 * 10^4
//     The input is generated such that edges represents a valid tree.

import "fmt"

func longestSpecialPath(edges [][]int, nums []int) []int {
    type Edge struct{ node, weight int }
    g := make([][]Edge, len(nums))
    for _, e := range edges {
        x, y, w := e[0], e[1], e[2]
        g[x] = append(g[x], Edge{ y, w })
        g[y] = append(g[y], Edge{ x, w })
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    maxLen, minNodes := -1, 0
    dis := []int{0}
    // 颜色 -> 该颜色最近一次出现的深度 +1，注意这里已经 +1 了，下面不需要再 +1
    lastDepth := map[int]int{}
    var dfs func(x, fa, topDepth, last1 int) 
    dfs = func(x, fa, topDepth, last1 int) {
        color := nums[x]
        last2 := lastDepth[color]
        topDepth = max(topDepth, min(last1, last2)) // 相较 3425 题，维护窗口左端点的逻辑变了
        length := dis[len(dis)-1] - dis[topDepth]
        nodes := len(dis) - topDepth
        if length > maxLen || length == maxLen && nodes < minNodes {
            maxLen, minNodes = length, nodes
        }
        lastDepth[color] = len(dis)
        for _, e := range g[x] {
            y := e.node
            if y != fa {
                dis = append(dis, dis[len(dis)-1] + e.weight)
                dfs(y, x, topDepth, max(last1, last2)) // 相较 3425 题，额外维护 last1
                dis = dis[:len(dis)-1]
            }
        }
        lastDepth[color] = last2
    }
    dfs(0, -1, 0, 0)
    return []int{ maxLen, minNodes }
}

func longestSpecialPath1(edges [][]int, nums []int) []int {
    n := len(edges) + 1
    adj := make([][][2]int, n)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        adj[u] = append(adj[u], [2]int{v, w})
        adj[v] = append(adj[v], [2]int{u, w})
    }
    var at [50505][]int
    var dfs func(adj [][][2]int, C []int, res *[]int, u, dep, par, tp1, tp2 int, ord *[]int)
    dfs = func(adj [][][2]int, C []int, res *[]int, u, dep, par, tp1, tp2 int, ord *[]int) {
        if len(at[C[u]]) > 0 {
            x := at[C[u]][len(at[C[u]])-1]
            if x > tp1 {
                x, tp1 = tp1, x
            }
            if x > tp2 {
                x, tp2 = tp2, x
            }
        }
        *ord = append(*ord, dep)
        at[C[u]] = append(at[C[u]], len(*ord))
        if dep-(*ord)[tp2] > (*res)[0] || (dep-(*ord)[tp2] == (*res)[0] && tp2-len(*ord) > (*res)[1]) {
            *res = []int{dep - (*ord)[tp2], tp2 - len(*ord)}
        }
        for _, vw := range adj[u] {
            v, w := vw[0], vw[1]
            if v == par {
                continue
            }
            dfs(adj, C, res, v, dep+w, u, tp1, tp2, ord)
        }
        *ord = (*ord)[:len(*ord)-1]
        at[C[u]] = at[C[u]][:len(at[C[u]])-1]
    }
    res, ord := []int{-1, -1}, []int{}
    dfs(adj, nums, &res, 0, 0, -1, 0, 0, &ord)
    res[1] = -res[1]
    return res
}

func main() {
    // Example 1:
    // Input: edges = [[0,1,1],[1,2,3],[1,3,1],[2,4,6],[4,7,2],[3,5,2],[3,6,5],[6,8,3]], nums = [1,1,0,3,1,2,1,1,0]
    // Output: [9,3]
    // Explanation:
    // In the image below, nodes are colored by their corresponding values in nums.
    // <img src="https://assets.leetcode.com/uploads/2025/02/18/e1.png" />
    // The longest special paths are 1 -> 2 -> 4 and 1 -> 3 -> 6 -> 8, both having a length of 9. 
    // The minimum number of nodes across all longest special paths is 3.
    fmt.Println(longestSpecialPath([][]int{{0,1,1},{1,2,3},{1,3,1},{2,4,6},{4,7,2},{3,5,2},{3,6,5},{6,8,3}}, []int{1,1,0,3,1,2,1,1,0})) // [9,3]
    // Example 2:
    // Input: edges = [[1,0,3],[0,2,4],[0,3,5]], nums = [1,1,0,2]
    // Output: [5,2]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/02/18/e2.png" />
    // The longest path is 0 -> 3 consisting of 2 nodes with a length of 5.
    fmt.Println(longestSpecialPath([][]int{{1,0,3},{0,2,4},{0,3,5}}, []int{1,1,0,2})) // [5,2]

    fmt.Println(longestSpecialPath1([][]int{{0,1,1},{1,2,3},{1,3,1},{2,4,6},{4,7,2},{3,5,2},{3,6,5},{6,8,3}}, []int{1,1,0,3,1,2,1,1,0})) // [9,3]
    fmt.Println(longestSpecialPath1([][]int{{1,0,3},{0,2,4},{0,3,5}}, []int{1,1,0,2})) // [5,2]
}