package main

// 3331. Find Subtree Sizes After Changes
// You are given a tree rooted at node 0 that consists of n nodes numbered from 0 to n - 1. 
// The tree is represented by an array parent of size n, where parent[i] is the parent of node i. 
// Since node 0 is the root, parent[0] == -1.

// You are also given a string s of length n, where s[i] is the character assigned to node i.

// We make the following changes on the tree one time simultaneously for all nodes x from 1 to n - 1:
//     1. Find the closest node y to node x such that y is an ancestor of x, and s[x] == s[y].
//     2. If node y does not exist, do nothing.
//     3. Otherwise, remove the edge between x and its current parent and make node y the new parent of x by adding an edge between them.

// Return an array answer of size n where answer[i] is the size of the subtree rooted at node i in the final tree.

// Example 1:
// Input: parent = [-1,0,0,1,1,1], s = "abaabc"
// Output: [6,3,1,1,1,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/15/graphex1drawio.png" />
// The parent of node 3 will change from node 1 to node 0.

// Example 2:
// Input: parent = [-1,0,4,0,1], s = "abbba"
// Output: [5,2,1,1,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/08/20/exgraph2drawio.png" />
// The following changes will happen at the same time:
// The parent of node 4 will change from node 1 to node 0.
// The parent of node 2 will change from node 4 to node 1.

// Constraints:
//     n == parent.length == s.length
//     1 <= n <= 10^5
//     0 <= parent[i] <= n - 1 for all i >= 1.
//     parent[0] == -1
//     parent represents a valid tree.
//     s consists only of lowercase English letters.

import "fmt"

func findSubtreeSizes(parent []int, s string) []int {
    n := len(s)
    graph := make([][]int, n)
    for i := 1; i < n; i++ {
        graph[parent[i]] = append(graph[parent[i]], i)
    }
    res, d := make([]int, n), make([][]int, 26)
    var dfs func(int, int)
    dfs = func(i, fa int) {
        res[i] = 1
        index := int(s[i] - 'a')
        d[index] = append(d[index], i)
        for _, j := range graph[i] {
            dfs(j, i)
        }
        k := fa
        if len(d[index]) > 1 {
            k = d[index][len(d[index]) - 2]
        }
        if k != -1 {
            res[k] += res[i]
        }
        d[index] = d[index][:len(d[index]) - 1]
    }
    dfs(0, -1)
    return res
}

func findSubtreeSizes1(parent []int, s string) []int {
    n := len(parent)
    graph := make([][]int, n)
    for i := 1; i < n; i++ {
        graph[parent[i]] = append(graph[parent[i]], i)
    }
    res, ancestor := make([]int, n), make([]int, 26)
    for i := range ancestor {
        ancestor[i] = -1
    }
    var dfs func(int)
    dfs = func(x int) {
        res[x] = 1
        sx := s[x] - 'a'
        old := ancestor[sx]
        ancestor[sx] = x
        for _, y := range graph[x] {
            dfs(y)
            anc := ancestor[s[y]-'a']
            if anc < 0 {
                anc = x
            }
            res[anc] += res[y]
        }
        ancestor[sx] = old // 恢复现场
    }
    dfs(0)
    return res
}

func main() {
    // Example 1:
    // Input: parent = [-1,0,0,1,1,1], s = "abaabc"
    // Output: [6,3,1,1,1,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/15/graphex1drawio.png" />
    // The parent of node 3 will change from node 1 to node 0.
    fmt.Println(findSubtreeSizes([]int{-1,0,0,1,1,1}, "abaabc")) // [6,3,1,1,1,1]
    // Example 2:
    // Input: parent = [-1,0,4,0,1], s = "abbba"
    // Output: [5,2,1,1,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/08/20/exgraph2drawio.png" />
    // The following changes will happen at the same time:
    // The parent of node 4 will change from node 1 to node 0.
    // The parent of node 2 will change from node 4 to node 1.
    fmt.Println(findSubtreeSizes([]int{-1,0,4,0,1}, "abbba")) //  [5,2,1,1,1]

    fmt.Println(findSubtreeSizes1([]int{-1,0,0,1,1,1}, "abaabc")) // [6,3,1,1,1,1]
    fmt.Println(findSubtreeSizes1([]int{-1,0,4,0,1}, "abbba")) //  [5,2,1,1,1]
}