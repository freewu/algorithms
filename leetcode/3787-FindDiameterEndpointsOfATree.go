package main

// 3787. Find Diameter Endpoints of a Tree
// You are given an undirected tree with n nodes, numbered from 0 to n - 1. 
// It is represented by a 2D integer array edges​​​​​​​ of length n - 1, where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

// A node is called special if it is an endpoint of any diameter path of the tree.

// Return a binary string s of length n, where s[i] = '1' if node i is special, and s[i] = '0' otherwise.

// A diameter path of a tree is the longest simple path between any two nodes. A tree may have multiple diameter paths.

// An endpoint of a path is the first or last node on that path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2025/11/30/pic1.png" />
// Input: n = 3, edges = [[0,1],[1,2]]
// Output: "101"
// Explanation:
// The diameter of this tree consists of 2 edges.
// The only diameter path is the path from node 0 to node 2
// The endpoints of this path are nodes 0 and 2, so they are special.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2025/11/30/pic2.png" />
// Input: n = 7, edges = [[0,1],[1,2],[2,3],[3,4],[3,5],[1,6]]
// Output: "1000111"
// Explanation:
// The diameter of this tree consists of 4 edges. There are 4 diameter paths:
// The path from node 0 to node 4
// The path from node 0 to node 5
// The path from node 6 to node 4
// The path from node 6 to node 5
// The special nodes are nodes 0, 4, 5, 6, as they are endpoints in at least one diameter path.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2025/11/30/pic3.png" />
// Input: n = 2, edges = [[0,1]]
// Output: "11"
// Explanation:
// The diameter of this tree consists of 1 edge.
// The only diameter path is the path from node 0 to node 1
// The endpoints of this path are nodes 0 and 1, so they are special.

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ai, bi]
//     0 <= ai, bi < n
//     The input is generated such that edges represents a valid tree.

import "fmt"

func findSpecialNodes(n int, edges [][]int) string {
    // 构建邻接表
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    // BFS函数：返回（最远节点，各节点到start的距离）
    bfs := func(start int) (int, []int) {
        dist := make([]int, n)
        for i := range dist {
            dist[i] = -1
        }
        q := []int{start}
        dist[start] = 0
        farthest := start
        for len(q) > 0 {
            curr := q[0]
            q = q[1:]
            for _, next := range adj[curr] {
                if dist[next] == -1 {
                    dist[next] = dist[curr] + 1
                    q = append(q, next)
                    if dist[next] > dist[farthest] {
                        farthest = next
                    }
                }
            }
        }
        return farthest, dist
    }
    // 第一步：找第一个端点u
    u, _ := bfs(0)
    // 第二步：找第二个端点v，同时得到u到所有节点的距离du
    v, du := bfs(u)
    // 第三步：得到v到所有节点的距离dv
    _, dv := bfs(v)
    // 直径长度
    m := du[v]
    // 特殊节点：是某条直径的端点 → 满足 (du[i] == m 或 dv[i] == m)
    // （因为du[i]=m 说明i是u的最远节点，即i是直径端点；dv[i]=m同理）
    res := make([]byte, n)
    for i := 0; i < n; i++ {
        if du[i] == m || dv[i] == m {
            res[i] = '1'
        } else {
            res[i] = '0'
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2025/11/30/pic1.png" />
    // Input: n = 3, edges = [[0,1],[1,2]]
    // Output: "101"
    // Explanation:
    // The diameter of this tree consists of 2 edges.
    // The only diameter path is the path from node 0 to node 2
    // The endpoints of this path are nodes 0 and 2, so they are special.
    fmt.Println(findSpecialNodes(3, [][]int{{0,1},{1,2}})) // "101"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2025/11/30/pic2.png" />
    // Input: n = 7, edges = [[0,1],[1,2],[2,3],[3,4],[3,5],[1,6]]
    // Output: "1000111"
    // Explanation:
    // The diameter of this tree consists of 4 edges. There are 4 diameter paths:
    // The path from node 0 to node 4
    // The path from node 0 to node 5
    // The path from node 6 to node 4
    // The path from node 6 to node 5
    // The special nodes are nodes 0, 4, 5, 6, as they are endpoints in at least one diameter path.
    fmt.Println(findSpecialNodes(7, [][]int{{0,1},{1,2},{2,3},{3,4},{3,5},{1,6}})) // "1000111"
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2025/11/30/pic3.png" />
    // Input: n = 2, edges = [[0,1]]
    // Output: "11"
    // Explanation:
    // The diameter of this tree consists of 1 edge.
    // The only diameter path is the path from node 0 to node 1
    // The endpoints of this path are nodes 0 and 1, so they are special.
    fmt.Println(findSpecialNodes(2, [][]int{{0,1}})) // "11"
}