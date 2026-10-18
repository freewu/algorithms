package main

// 4085. Find Root of a Colored Tree
// You are given an undirected tree of n nodes labeled from 0 to n - 1, 
// represented by a 2D integer array edges of length n - 1, 
// where edges[i] = [u_{i}, v_{i}] indicates that there is an undirected edge between nodes u_{i} and v_{i}.

// You are also given a string colors of length n, where colors[i] is one of 'R', 'B', or 'W', and denotes the color of node i.

// A node chosen as root must satisfy both of the following conditions once the tree is oriented away from it:
//     1. Binary-tree property: Every node has at most two children.
//     2. Layered color pattern: All nodes at the same depth have the same color. 
//        Starting at depth 0, these colors follow one of the three repeating sequences:
//             2.1 R → B → W → R → B → W → ...
//             2.2 B → W → R → B → W → R → ...
//             2.3 W → R → B → W → R → B → ...

// Return the index of any root that satisfies both conditions, or -1 if no such root exists.

// A tree is a connected undirected graph with no cycles.

// The depth of a node is the number of edges on the path from the root to that node. 
// The root has depth 0. A child of a node is an adjacent node whose depth is one greater.

// The degree of a node is the number of edges incident to it.

// Example 1:
// Input: edges = [[3,1],[3,2],[1,4],[2,0]], colors = "BRRWB"
// Output: 3
// Explanation:
// <img alt="" src="https://assets.leetcode.com/uploads/2025/06/15/treee.jpg" style="width: 351px; height: 224px;">
// Rooting at node 3 gives each node at most two children and every edge advances to the next color in the cycle (W → R → B), so return 3.

// Example 2:
// Input: edges = [[0,1],[1,2],[2,3]], colors = "RBRB"
// Output: -1
// Explanation:
// <img alt="" src="https://assets.leetcode.com/uploads/2025/06/15/treee2.jpg" style="width: 282px; height: 41px;">
// No root satisfies the required color step on every edge, so return -1.

// Constraints:
//     - 2 <= n <= 10^{5}
//     - edges.length == n - 1
//     - edges[i] = [u_{i}, v_{i}]
//     - 0 <= u_{i}, v_{i} < n
//     - colors.length == n
//     - colors[i] is one of 'R', 'B', 'W'.
//     - The input graph is a tree and every node has degree at most three.

import "fmt"

func findColoredRoot(edges [][]int, colors string) int {
    mp := map[byte]int{'W': 0, 'R': 1, 'B': 2}
    res, n := -1, len(colors)
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    for x, to := range g {
        // 一个点至多三个邻居
        if len(to) > 3 {
            return -1
        }
        count := [3]int{}
        for _, y := range to {
            count[mp[colors[y]]]++
        }
        colorX := mp[colors[x]]
        // 不能有邻居和 x 颜色相同，或者 x 的父节点的颜色个数超过 1
        if count[colorX] > 0 || count[(colorX + 2) % 3] > 1 {
            return -1
        }
        if len(to) < 3 && count[(colorX + 2) % 3] == 0 {
            // 不能有多个点能是根，那样当一个点是根的时候，另一个点不满足要求
            if res >= 0 {
                return -1
            }
            res = x
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: edges = [[3,1],[3,2],[1,4],[2,0]], colors = "BRRWB"
    // Output: 3
    // Explanation:
    // <img alt="" src="https://assets.leetcode.com/uploads/2025/06/15/treee.jpg" style="width: 351px; height: 224px;">
    // Rooting at node 3 gives each node at most two children and every edge advances to the next color in the cycle (W → R → B), so return 3.
    fmt.Println(findColoredRoot([][]int{{3,1},{3,2},{1,4},{2,0}}, "BRRWB")) // 3 
    // Example 2:
    // Input: edges = [[0,1],[1,2],[2,3]], colors = "RBRB"
    // Output: -1
    // Explanation:
    // <img alt="" src="https://assets.leetcode.com/uploads/2025/06/15/treee2.jpg" style="width: 282px; height: 41px;">
    // No root satisfies the required color step on every edge, so return -1.
    fmt.Println(findColoredRoot([][]int{{0,1},{1,2},{2,3}}, "RBRB")) // -1
}