package main

// 1938. Maximum Genetic Difference Query
// There is a rooted tree consisting of n nodes numbered 0 to n - 1. 
// Each node's number denotes its unique genetic value (i.e. the genetic value of node x is x). 
// The genetic difference between two genetic values is defined as the bitwise-XOR of their values. 
// You are given the integer array parents, where parents[i] is the parent for node i. 
// If node x is the root of the tree, then parents[x] == -1.

// You are also given the array queries where queries[i] = [nodei, vali]. 
// For each query i, find the maximum genetic difference between vali and pi, 
// where pi is the genetic value of any node that is on the path between nodei and the root (including nodei and the root). More formally, you want to maximize vali XOR pi.

// Return an array ans where ans[i] is the answer to the ith query.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/c1.png" />
// Input: parents = [-1,0,1,1], queries = [[0,2],[3,2],[2,5]]
// Output: [2,3,7]
// Explanation: The queries are processed as follows:
// - [0,2]: The node with the maximum genetic difference is 0, with a difference of 2 XOR 0 = 2.
// - [3,2]: The node with the maximum genetic difference is 1, with a difference of 2 XOR 1 = 3.
// - [2,5]: The node with the maximum genetic difference is 2, with a difference of 5 XOR 2 = 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/29/c2.png" />
// Input: parents = [3,7,-1,2,0,7,0,2], queries = [[4,6],[1,15],[0,5]]
// Output: [6,14,7]
// Explanation: The queries are processed as follows:
// - [4,6]: The node with the maximum genetic difference is 0, with a difference of 6 XOR 0 = 6.
// - [1,15]: The node with the maximum genetic difference is 1, with a difference of 15 XOR 1 = 14.
// - [0,5]: The node with the maximum genetic difference is 2, with a difference of 5 XOR 2 = 7.

// Constraints:
//     2 <= parents.length <= 10^5
//     0 <= parents[i] <= parents.length - 1 for every node i that is not the root.
//     parents[root] == -1
//     1 <= queries.length <= 3 * 10^4
//     0 <= nodei <= parents.length - 1
//     0 <= vali <= 2 * 10^5

import "fmt"

type node struct {
    son [2]*node
    cnt int
}
type Trie struct{ root *node }

func (this *Trie) Put(v int) *node {
    o := this.root
    for i := 17; i >= 0; i-- {
        b := v >> i & 1
        if o.son[b] == nil {
            o.son[b] = &node{}
        }
        o = o.son[b]
        o.cnt++
    }
    return o
}

func (this *Trie) Delete(v int) *node {
    o := this.root
    for i := 17; i >= 0; i-- {
        o = o.son[v>>i&1]
        o.cnt-- // 删除操作只需要减少 cnt 就行，cnt 为 0 就视作删掉了该节点
    }
    return o
}

func (this *Trie) MaxXor(v int) int {
    res, o := 0, this.root
    for i := 17; i >= 0; i-- {
        b := v >> i & 1
        if o.son[b^1] != nil && o.son[b^1].cnt > 0 {
            res |= 1 << i
            b ^= 1
        }
        o = o.son[b]
    }
    return res
}

// 离线+字典树
func maxGeneticDifference(parents []int, queries [][]int) []int {
    n, root := len(parents), 0
    g := make([][]int, n) // 建树
    for v, pa := range parents {
        if pa == -1 {
            root = v
        } else {
            g[pa] = append(g[pa], v)
        }
    }
    // 离线，将查询分组
    type query struct{ val, i int }
    qs := make([][]query, n)
    for i, q := range queries {
        qs[q[0]] = append(qs[q[0]], query{q[1], i})
    }
    res := make([]int, len(queries))
    t := &Trie{&node{}}
    var dfs func(int) // 遍历整棵树，每访问一个节点就将其插入 trie 树，访问结束时将其从 trie 中删去
    dfs = func(v int) {
        t.Put(v)
        for _, q := range qs[v] {
            res[q.i] = t.MaxXor(q.val)
        }
        for _, w := range g[v] {
            dfs(w)
        }
        t.Delete(v)
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/c1.png" />
    // Input: parents = [-1,0,1,1], queries = [[0,2],[3,2],[2,5]]
    // Output: [2,3,7]
    // Explanation: The queries are processed as follows:
    // - [0,2]: The node with the maximum genetic difference is 0, with a difference of 2 XOR 0 = 2.
    // - [3,2]: The node with the maximum genetic difference is 1, with a difference of 2 XOR 1 = 3.
    // - [2,5]: The node with the maximum genetic difference is 2, with a difference of 5 XOR 2 = 7.
    fmt.Println(maxGeneticDifference([]int{-1,0,1,1},[][]int{{0,2},{3,2},{2,5}})) // [2,3,7]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/29/c2.png" />
    // Input: parents = [3,7,-1,2,0,7,0,2], queries = [[4,6],[1,15],[0,5]]
    // Output: [6,14,7]
    // Explanation: The queries are processed as follows:
    // - [4,6]: The node with the maximum genetic difference is 0, with a difference of 6 XOR 0 = 6.
    // - [1,15]: The node with the maximum genetic difference is 1, with a difference of 15 XOR 1 = 14.
    // - [0,5]: The node with the maximum genetic difference is 2, with a difference of 5 XOR 2 = 7.
    fmt.Println(maxGeneticDifference([]int{3,7,-1,2,0,7,0,2},[][]int{{4,6},{1,15},{0,5}})) // [6,14,7]
}