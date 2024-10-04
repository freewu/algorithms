package main

// 2479. Maximum XOR of Two Non-Overlapping Subtrees
// There is an undirected tree with n nodes labeled from 0 to n - 1. 
// You are given the integer n and a 2D integer array edges of length n - 1, 
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree. 
// The root of the tree is the node labeled 0.

// Each node has an associated value. 
// You are given an array values of length n, where values[i] is the value of the ith node.

// Select any two non-overlapping subtrees. 
// Your score is the bitwise XOR of the sum of the values within those subtrees.

// Return the maximum possible score you can achieve. 
// If it is impossible to find two nonoverlapping subtrees, return 0.

// Note that:
//     The subtree of a node is the tree consisting of that node and all of its descendants.
//     Two subtrees are non-overlapping if they do not share any common node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/11/22/treemaxxor.png" />
// Input: n = 6, edges = [[0,1],[0,2],[1,3],[1,4],[2,5]], values = [2,8,3,6,2,5]
// Output: 24
// Explanation: Node 1's subtree has sum of values 16, while node 2's subtree has sum of values 8, so choosing these nodes will yield a score of 16 XOR 8 = 24. It can be proved that is the maximum possible score we can obtain.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/11/22/tree3drawio.png" />
// Input: n = 3, edges = [[0,1],[1,2]], values = [4,6,1]
// Output: 0
// Explanation: There is no possible way to select two non-overlapping subtrees, so we just return 0.

// Constraints:
//     2 <= n <= 5 * 10^4
//     edges.length == n - 1
//     0 <= ai, bi < n
//     values.length == n
//     1 <= values[i] <= 10^9
//     It is guaranteed that edges represents a valid tree.

import "fmt"

type Trie struct {
    Children [2]*Trie 
}

func (this *Trie) Find(x int) int {
    res := 0
    for i := 60 ;i >= 0; i-- {
        p := int((x >> i) & 1)
        if this.Children[1-p] == nil {
            if this.Children[p] == nil { return res } 
            this = this.Children[p]
        }else {
            res += 1 << i
            this = this.Children[1-p]
        } 
    }
    return res
}

func (this *Trie) Insert(x int)  {
    for i := 60 ; i >= 0; i-- {
        p := ( x >> i ) & 1
        if this.Children[p] == nil {
            this.Children[p] = &Trie{}
        }
        this = this.Children[p]
    }
}

func maxXor(n int, edges [][]int, values []int) int64 {
    res, n := 0, len(values)
    g, sum := make([][]int, n), make([]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    }
    tr := &Trie{}
    var dfs func(u int, fa int, val []int) int
    dfs = func(u int, fa int, val []int) int {
        sum[u] += val[u]
        for _, v := range g[u] {
            if v != fa {
                sum[u] += dfs(v, u, val)
            }
        }
        return sum[u]
    }
    var dfs2 func(u int, fa int)
    dfs2 = func(u int, fa int) {
        res = max(res, tr.Find(sum[u]))
        for _, v := range g[u] {
            if v != fa { dfs2( v, u) }
        }
        tr.Insert(sum[u])
    }
    dfs(0, -1, values)
    dfs2(0, -1)
    return int64(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/11/22/treemaxxor.png" />
    // Input: n = 6, edges = [[0,1],[0,2],[1,3],[1,4],[2,5]], values = [2,8,3,6,2,5]
    // Output: 24
    // Explanation: Node 1's subtree has sum of values 16, while node 2's subtree has sum of values 8, so choosing these nodes will yield a score of 16 XOR 8 = 24. It can be proved that is the maximum possible score we can obtain.
    fmt.Println(maxXor(6, [][]int{{0,1},{0,2},{1,3},{1,4},{2,5}}, []int{2,8,3,6,2,5})) // 24
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/11/22/tree3drawio.png" />
    // Input: n = 3, edges = [[0,1],[1,2]], values = [4,6,1]
    // Output: 0
    // Explanation: There is no possible way to select two non-overlapping subtrees, so we just return 0.
    fmt.Println(maxXor(3, [][]int{{0,1},{1,2}}, []int{4,6,1})) // 0
}