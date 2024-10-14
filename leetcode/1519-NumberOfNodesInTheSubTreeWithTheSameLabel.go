package main

// 1519. Number of Nodes in the Sub-Tree With the Same Label
// You are given a tree (i.e. a connected, undirected graph that has no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges. 
// The root of the tree is the node 0, and each node of the tree has a label which is a lower-case character given in the string labels (i.e. The node with the number i has the label labels[i]).

// The edges array is given on the form edges[i] = [ai, bi], 
// which means there is an edge between nodes ai and bi in the tree.

// Return an array of size n where ans[i] is the number of nodes in the subtree of the ith node which have the same label as node i.

// A subtree of a tree T is the tree consisting of a node in T and all of its descendant nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/07/01/q3e1.jpg" />
// Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], labels = "abaedcd"
// Output: [2,1,1,1,1,1,1]
// Explanation: Node 0 has label 'a' and its sub-tree has node 2 with label 'a' as well, thus the answer is 2. Notice that any node is part of its sub-tree.
// Node 1 has a label 'b'. The sub-tree of node 1 contains nodes 1,4 and 5, as nodes 4 and 5 have different labels than node 1, the answer is just 1 (the node itself).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/07/01/q3e2.jpg" />
// Input: n = 4, edges = [[0,1],[1,2],[0,3]], labels = "bbbb"
// Output: [4,2,1,1]
// Explanation: The sub-tree of node 2 contains only node 2, so the answer is 1.
// The sub-tree of node 3 contains only node 3, so the answer is 1.
// The sub-tree of node 1 contains nodes 1 and 2, both have label 'b', thus the answer is 2.
// The sub-tree of node 0 contains nodes 0, 1, 2 and 3, all with label 'b', thus the answer is 4.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/07/01/q3e3.jpg" />
// Input: n = 5, edges = [[0,1],[0,2],[1,3],[0,4]], labels = "aabab"
// Output: [3,2,1,1,1]

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     labels.length == n
//     labels is consisting of only of lowercase English letters.

import "fmt"

func countSubTrees(n int, edges [][]int, labels string) []int {
    res, mp := make([]int, n), make(map[int][]int)
    for _, e := range edges {
        mp[e[0]] = append(mp[e[0]], e[1])
        mp[e[1]] = append(mp[e[1]], e[0])
    }
    var dfs func(graph map[int][]int, labels string, node, parent int) []int 
    dfs = func(graph map[int][]int, labels string, node, parent int) []int {
        nodeCnt := make([]int, 26)
        nodeCnt[labels[node] - 'a']++
        for _, nei := range graph[node] {
            if nei != parent {
                childCnt := dfs(graph, labels, nei, node)
                for i := 0; i < 26; i++ {
                    nodeCnt[i] += childCnt[i]
                }
            }
        }
        res[node] = nodeCnt[labels[node] - 'a']
        return nodeCnt
    }
    dfs(mp, labels, 0, -1)
    return res
}

func countSubTrees1(n int, edges [][]int, labels string) []int {
    // 二叉 dp 问题，不断更新的数据是 cnts
    // 搜集下层的数据，不断累加到嘴上一层
    graph := make([][]int, n)
    for _, e := range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }
    res := make([]int, n)
    var dfs func(u, f int) []int
    dfs = func(u, f int) []int {
        cnt := make([]int, 26)
        cnt[labels[u] - 'a']++
        for _, c := range graph[u]{
            if c == f { continue }
            sub := dfs(c, u)
            for i := 0; i < 26; i++{
                cnt[i] += sub[i]
            }
        }
        res[u] = cnt[labels[u]-'a']
        return cnt
    }
    dfs(0, -1)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/07/01/q3e1.jpg" />
    // Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], labels = "abaedcd"
    // Output: [2,1,1,1,1,1,1]
    // Explanation: Node 0 has label 'a' and its sub-tree has node 2 with label 'a' as well, thus the answer is 2. Notice that any node is part of its sub-tree.
    // Node 1 has a label 'b'. The sub-tree of node 1 contains nodes 1,4 and 5, as nodes 4 and 5 have different labels than node 1, the answer is just 1 (the node itself).
    fmt.Println(countSubTrees(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, "abaedcd")) // [2,1,1,1,1,1,1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/07/01/q3e2.jpg" />
    // Input: n = 4, edges = [[0,1],[1,2],[0,3]], labels = "bbbb"
    // Output: [4,2,1,1]
    // Explanation: The sub-tree of node 2 contains only node 2, so the answer is 1.
    // The sub-tree of node 3 contains only node 3, so the answer is 1.
    // The sub-tree of node 1 contains nodes 1 and 2, both have label 'b', thus the answer is 2.
    // The sub-tree of node 0 contains nodes 0, 1, 2 and 3, all with label 'b', thus the answer is 4.
    fmt.Println(countSubTrees(4, [][]int{{0,1},{1,2},{0,3}}, "bbbb")) // [4,2,1,1]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/07/01/q3e3.jpg" />
    // Input: n = 5, edges = [[0,1],[0,2],[1,3],[0,4]], labels = "aabab"
    // Output: [3,2,1,1,1]
    fmt.Println(countSubTrees(5, [][]int{{0,1},{0,2},{1,3},{0,4}}, "aabab")) // [3,2,1,1,1]

    fmt.Println(countSubTrees1(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, "abaedcd")) // [2,1,1,1,1,1,1]
    fmt.Println(countSubTrees1(4, [][]int{{0,1},{1,2},{0,3}}, "bbbb")) // [4,2,1,1]
    fmt.Println(countSubTrees1(5, [][]int{{0,1},{0,2},{1,3},{0,4}}, "aabab")) // [3,2,1,1,1]
}