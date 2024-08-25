package main

// 1245. Tree Diameter
// The diameter of a tree is the number of edges in the longest path in that tree.

// There is an undirected tree of n nodes labeled from 0 to n - 1. 
// You are given a 2D array edges where edges.length == n - 1 and edges[i] = [ai, bi] indicates 
// that there is an undirected edge between nodes ai and bi in the tree.

// Return the diameter of the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/01/19/tree1.jpg" />
// Input: edges = [[0,1],[0,2]]
// Output: 2
// Explanation: The longest path of the tree is the path 1 - 0 - 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/01/19/tree2.jpg" />
// Input: edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
// Output: 4
// Explanation: The longest path of the tree is the path 3 - 2 - 1 - 4 - 5.

// Constraints:
//     n == edges.length + 1
//     1 <= n <= 10^4
//     0 <= ai, bi < n
//     ai != bi

import "fmt"

func treeDiameter(edges [][]int) int {
    if edges == nil || len(edges) == 0 {
        return 0
    }
    tree := make(map[int][]int, len(edges))
    for _, v := range edges {
        v0, v1 := v[0], v[1]
        if v1 < v0 {
            v0, v1 = v[1], v[0]
        }
        nodes, ok := tree[v0]
        if !ok {
            nodes = []int{}
        }
        tree[v0] = append(nodes, v1)
    }
    var dfs func(tree map[int][]int, k int) (int, int)
    dfs = func(tree map[int][]int, k int) (int, int) {
        childs, ok := tree[k]
        if !ok || len(childs) == 0 { // 不存在子节点，返回 0
            return 0, 0
        }
        diameter, h1, h2 := 0, -1, -1
        for _, c := range childs {
            h, d := dfs(tree, c)
            if diameter < d {
                diameter = d
            }
            if h1 < h {
                h1, h2 = h, h1
            } else if h2 < h {
                h2 = h
            }
        }
        if diameter < h1 + h2 + 2 { // 存在两个子节点，设所有子树中最高为 h1，次高为 h2。则有 diameter = h1 + h2 + 2
            diameter = h1 + h2 + 2
        }
        return h1 + 1, diameter
    }
    _, path := dfs(tree, 0)
    return path
}

func treeDiameter1(edges [][]int) int {
    // 树形dp
    // 枚举各个顶点作为root, 那么最长路径可能经过当前顶点,那么也是child中最长/次长的和d1+d2,
    // 也可能不经过root,则是其它顶点作为root,等后续枚举其它顶点作为root即可. dp[x]=max( d1+d2 ,dp[x_child1], dp[x_child2]...)
    res, n := 0, len(edges) + 1 // 树有n-1条边
    g := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(sn int, fa int) int
    dfs = func(sn int, fa int) int {
        d1, d2 := 0, 0 // ,兼容只有一个子树/没有子树的情况. distance,不是树高,是和最长链尖端的节点的距离, 如果只有一个节点,距离最远的节点还是0
        for _, v := range g[sn] {
            if v == fa {
                continue
            }
            d := dfs(v, sn) + 1
            if d > d1 { // >=或者>都可以, ==时不会丢失,因为还会与h2比一次
                d1, d2 = d, d1
            } else if d > d2 {
                d2 = d
            }
        }
        res = max(res, d1 + d2)
        return max(d1, d2) // 兼容没有child(最远的节点为自己,dis=0),只有一个child的情况
    }
    dfs(0, -1)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/01/19/tree1.jpg" />
    // Input: edges = [[0,1],[0,2]]
    // Output: 2
    // Explanation: The longest path of the tree is the path 1 - 0 - 2.
    fmt.Println(treeDiameter([][]int{{0,1},{0,2}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/01/19/tree2.jpg" />
    // Input: edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
    // Output: 4
    // Explanation: The longest path of the tree is the path 3 - 2 - 1 - 4 - 5.
    fmt.Println(treeDiameter([][]int{{0,1},{1,2},{2,3},{1,4},{4,5}})) // 4

    fmt.Println(treeDiameter1([][]int{{0,1},{0,2}})) // 2
    fmt.Println(treeDiameter1([][]int{{0,1},{1,2},{2,3},{1,4},{4,5}})) // 4
}