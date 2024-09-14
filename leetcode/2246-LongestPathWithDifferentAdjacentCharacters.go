package main

// 2246. Longest Path With Different Adjacent Characters
// You are given a tree (i.e. a connected, undirected graph that has no cycles) rooted at node 0 consisting of n nodes numbered from 0 to n - 1.
// The tree is represented by a 0-indexed array parent of size n, where parent[i] is the parent of node i. 
// Since node 0 is the root, parent[0] == -1.

// You are also given a string s of length n, where s[i] is the character assigned to node i.

// Return the length of the longest path in the tree such that no pair of adjacent nodes on the path have the same character assigned to them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/25/testingdrawio.png" />
// Input: parent = [-1,0,0,1,1,2], s = "abacbe"
// Output: 3
// Explanation: The longest path where each two adjacent nodes have different characters in the tree is the path: 0 -> 1 -> 3. The length of this path is 3, so 3 is returned.
// It can be proven that there is no longer path that satisfies the conditions. 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/03/25/graph2drawio.png" />
// Input: parent = [-1,0,0,0], s = "aabc"
// Output: 3
// Explanation: The longest path where each two adjacent nodes have different characters is the path: 2 -> 0 -> 3. The length of this path is 3, so 3 is returned.

// Constraints:
//     n == parent.length == s.length
//     1 <= n <= 10^5
//     0 <= parent[i] <= n - 1 for all i >= 1
//     parent[0] == -1
//     parent represents a valid tree.
//     s consists of only lowercase English letters.

import "fmt"

func longestPath(parent []int, s string) int {
    res, edges := 0, make(map[int][]int)
    for v, u := range parent {
        if u != -1 { // 除开 root
            edges[u] = append(edges[u], v)
        }
    }
    var dfs func(int) int
    dfs = func(u int) int {
        d1, d2 := 0, 0
        for _, v := range edges[u] {
            dv := dfs(v)
            if s[v] != s[u] {
                if dv > d1 {
                    d2, d1 = d1, dv
                } else if dv > d2 {
                    d2 = dv
                }
            }
        }
        if d1 + d2 + 1 > res {
            res = d1 + d2 + 1
        }
        return d1 + 1
    }
    dfs(0)
    return res
}

func longestPath1(parent []int, s string) int {
    res, n := 0, len(parent)
    graph := make([][]int, n)
    for i := 1; i < n; i++ {
        graph[parent[i]] = append(graph[parent[i]], i)
    }
    var dfs func(x int) int
    dfs = func(x int) int {
        mx := 0
        for _, v := range graph[x] {
            l := dfs(v) + 1
            if s[v] != s[x] {
                res = max(res, l + mx)
                mx = max(mx, l)
            }
        }
        return mx
    }
    dfs(0)
    return res + 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/25/testingdrawio.png" />
    // Input: parent = [-1,0,0,1,1,2], s = "abacbe"
    // Output: 3
    // Explanation: The longest path where each two adjacent nodes have different characters in the tree is the path: 0 -> 1 -> 3. The length of this path is 3, so 3 is returned.
    // It can be proven that there is no longer path that satisfies the conditions. 
    fmt.Println(longestPath([]int{-1,0,0,1,1,2}, "abacbe")) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/03/25/graph2drawio.png" />
    // Input: parent = [-1,0,0,0], s = "aabc"
    // Output: 3
    // Explanation: The longest path where each two adjacent nodes have different characters is the path: 2 -> 0 -> 3. The length of this path is 3, so 3 is returned.
    fmt.Println(longestPath([]int{-1,0,0,0}, "aabc")) // 3

    fmt.Println(longestPath1([]int{-1,0,0,1,1,2}, "abacbe")) // 3
    fmt.Println(longestPath1([]int{-1,0,0,0}, "aabc")) // 3
}