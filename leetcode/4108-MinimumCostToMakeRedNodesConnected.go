package main

// 4108. Minimum Cost to Make Red Nodes Connected
// You are given an integer n representing a tree with n nodes, where the nodes are numbered from 0 to n - 1 and the tree is rooted at node 0. 
// The tree is represented by a 2D array edges of length n - 1, where edges[i] = [u_{i}, v_{i}] indicates an undirected edge between nodes u_{i} and v_{i}.

// You are also given a string colors of length n, where colors[i] represents the color of node i: 'R' if it is Red, or 'B' if it is Blue.

// You may flip the color of any node: from 'R' to 'B', or from 'B' to 'R'. 
// Flipping node u incurs a cost equal to the size of the subtree of node u.

// Your goal is to perform flips so the Red nodes form a single connected component in the tree. 
// Blue nodes may be disconnected. 
// A configuration with no Red nodes is also considered valid.

// Return the minimum total cost required to achieve this.

// In a rooted tree, the subtree of a node v is the set of all nodes whose path to the root contains v.

// Example 1:
// Input: n = 3, colors = "RBR", edges = [[0,1],[1,2]]
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33714am.png" style="height: 60px; width: 237px;">
//     There are two Red nodes (0 and 2) that are not connected to each other.
//     We can flip node 2 from 'R' (Red) to 'B' (Blue). The subtree size of node 2 is 1, so the cost is 1.
//     After this flip only node 0 remains Red, forming exactly one connected Red component.

// Example 2:
// Input: n = 3, colors = "RBR", edges = [[0,1],[0,2]]
// Output: 0
// Explanation:​​​​​​​
// <img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33754am.png" style="width: 214px; height: 150px;">
//     There are two Red nodes (0 and 2) that are already connected to each other.
//     No flips are needed since all Red nodes already form a single connected group. 
//     Thus, the minimum total cost is 0.

// Example 3:
// Input: n = 4, colors = "BRRR", edges = [[0,1],[0,2],[0,3]]
// Output: 2
// Explanation:​​​​​​​
// <img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33832am.png" style="width: 242px; height: 150px;">
//     There are three Red nodes (1, 2 and 3), which are not connected to each other.
//     All three Red nodes are leaves and are disconnected from one another.
//     To obtain a single Red group, we keep one Red leaf and flip the other two to Blue.
//     Each leaf subtree size is 1, so flipping any two costs 2.

// Example 4:
// Input: n = 3, colors = "BBB", edges = [[0,1],[1,2]]
// Output: 0
// Explanation:
// ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33900am.png" style="width: 277px; height: 55px;">
//     All nodes are Blue, so there are no Red nodes.
//     An empty Red group is considered a single connected group, so this state is already valid. 
//     Thus, the cost is 0.

// Constraints:​​​​​​​
//     1 <= n <= 10^{5}
//     n == colors.length
//     colors[i] is either 'R' or 'B'.
//     edges.length == n - 1
//     edges[i] = [u_{i}, v_{i}]
//     0 <= u_{i}, v_{i} <= n - 1
//     The input is generated such that edges represents a valid tree.

import "fmt"

func minCost(n int, colors string, edges [][]int) int64 {
    graph := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        graph[x], graph[y] = append(graph[x], y), append(graph[y], x)
    }
    base, mx := 0, 0
    var dfs func(x, fa int) (int, int)
    dfs = func(x, fa int) (int, int) {
        size, f := 1, 0
        // 遍历 x 的子树
        for _, y := range graph[x] {
            if y != fa {
                sz, fy := dfs(y, x)
                size += sz
                // 如果 fy < 0，那么子树 y 都变成蓝色是最优的，这已经算入 base 中
                f += max(fy, 0)
            }
        }
        // 考虑节点 x
        if colors[x] == 'R' {
            base += size // 把 x 变成蓝色的代价
            f += size // 把 x 恢复成红色可以让 base 减少 size
        } else {
            f -= size // 把 x 变成红色会让 base 增大 size
        }
        mx = max(mx, f) // 最多可以让 base 减少多少
        return size, f
    }
    dfs(0, -1)
    return int64(base - mx)
}

var deg, adj [100_000]uint32
var size, diff [100_000]int64

func minCost1(n int, colors string, edges [][]int) int64 {
    clear(deg[:n])
    clear(adj[:n])
    clear(size[:n])
    clear(diff[:n])
    for _, e := range edges {
        u, v := e[0], e[1]
        deg[u]++
        deg[v]++
        adj[u] ^= uint32(v)
        adj[v] ^= uint32(u)
    }
    allblue, best := int64(0), int64(0)
    done := func(v int) (sz, d int64) {
        sz = size[v] + 1
        if colors[v] == 'R' {
            d = diff[v] - sz
            allblue += sz
        } else {
            d = diff[v] + sz
        }
        best = min(best, d)
        return sz, d
    }
    for i := 1; i < n; i++ {
        for v := i; v != 0 && deg[v] == 1; {
            p := int(adj[v])
            sz, d := done(v)
            size[p] += sz
            diff[p] += min(d, 0)
            deg[v] = 0
            deg[p]--
            adj[p] ^= uint32(v)
            v = p
        }
    }
    done(0)
    return allblue + best
}

func main() {
    // Example 1:
    // Input: n = 3, colors = "RBR", edges = [[0,1],[1,2]]
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33714am.png" style="height: 60px; width: 237px;">
    //     There are two Red nodes (0 and 2) that are not connected to each other.
    //     We can flip node 2 from 'R' (Red) to 'B' (Blue). The subtree size of node 2 is 1, so the cost is 1.
    //     After this flip only node 0 remains Red, forming exactly one connected Red component.
    fmt.Println(minCost(3, "RBR", [][]int{{0,1},{1,2}})) // 1 
    // Example 2:
    // Input: n = 3, colors = "RBR", edges = [[0,1],[0,2]]
    // Output: 0
    // Explanation:​​​​​​​
    // <img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33754am.png" style="width: 214px; height: 150px;">
    //     There are two Red nodes (0 and 2) that are already connected to each other.
    //     No flips are needed since all Red nodes already form a single connected group. 
    //     Thus, the minimum total cost is 0.
    fmt.Println(minCost(3, "RBR", [][]int{{0,1},{0,2}})) // 0 
    // Example 3:
    // Input: n = 4, colors = "BRRR", edges = [[0,1],[0,2],[0,3]]
    // Output: 2
    // Explanation:​​​​​​​
    // <img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33832am.png" style="width: 242px; height: 150px;">
    //     There are three Red nodes (1, 2 and 3), which are not connected to each other.
    //     All three Red nodes are leaves and are disconnected from one another.
    //     To obtain a single Red group, we keep one Red leaf and flip the other two to Blue.
    //     Each leaf subtree size is 1, so flipping any two costs 2.
    fmt.Println(minCost(4, "BRRR", [][]int{{0,1},{0,2},{0,3}})) // 2 
    // Example 4:
    // Input: n = 3, colors = "BBB", edges = [[0,1],[1,2]]
    // Output: 0
    // Explanation:
    // ​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/07/10/screenshot-2026-07-11-at-33900am.png" style="width: 277px; height: 55px;">
    //     All nodes are Blue, so there are no Red nodes.
    //     An empty Red group is considered a single connected group, so this state is already valid. 
    //     Thus, the cost is 0.
    fmt.Println(minCost(3, "BBB", [][]int{{0,1},{1,2}})) // 0 

    fmt.Println(minCost1(3, "RBR", [][]int{{0,1},{1,2}})) // 1 
    fmt.Println(minCost1(3, "RBR", [][]int{{0,1},{0,2}})) // 0 
    fmt.Println(minCost1(4, "BRRR", [][]int{{0,1},{0,2},{0,3}})) // 2 
    fmt.Println(minCost1(3, "BBB", [][]int{{0,1},{1,2}})) // 0 
}