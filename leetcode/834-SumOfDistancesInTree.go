package main

// 834. Sum of Distances in Tree
// There is an undirected connected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.
// You are given the integer n and the array edges where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.
// Return an array answer of length n where answer[i] is the sum of the distances between the ith node in the tree and all other nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-sumdist1.jpg" / >
//               0
//             /   \
//            1      2
//                 / | \
//                3  4  5
// Input: n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
// Output: [8,12,6,10,10,10]
// Explanation: The tree is shown above.
// We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
// equals 1 + 1 + 2 + 2 + 2 = 8.
// Hence, answer[0] = 8, and so on.

// Example 2:
// Input: n = 1, edges = []
// Output: [0]

// Example 3:
//      0
//     /
//    1
// Input: n = 2, edges = [[1,0]]
// Output: [1,1]
 
// Constraints:
//     1 <= n <= 3 * 10^4
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     The given input represents a valid tree.

import "fmt"

func sumOfDistancesInTree(n int, edges [][]int) []int {
    dp, tree, adj := make([]int, n), make([]int, n), make([][]int, n)
    for _, e := range edges { // 初始化邻接表
        adj[e[0]], adj[e[1]] = append(adj[e[0]], e[1]), append(adj[e[1]], e[0])
    }
    var dfs func(int, int)
    dfs = func(p, c int) { // "归"值，所以计算的是 c
        tree[c] = 1 // 每个节点的子树，至少有一个节点，就是它本身
        for _, i := range adj[c] {
            if i != p {
                dfs(c, i)
                dp[c] += dp[i] + tree[i]
                tree[c] += tree[i]
            }
        }
    }
    dfs(-1, 0) //  // 初始化 dp 和 tree（以 0 为 root，计算各子树的距离和）
    var dfsDP func(int, int)
    dfsDP = func(p, c int) { // "递"值，所以计算的是 i
        for _, i := range adj[c] {
            if i != p { // 以 i 为 root，根据“父节点”逆向计算 i 的 距离和
                dp[i] += dp[c] - dp[i] + n - tree[i]<<1
                dfsDP(c, i)
            }
        }
    }
    dfsDP(-1, 0) // 计算每一个节点的距离和
    return dp
}

func sumOfDistancesInTree1(n int, edges [][]int) []int {
    g := make([][]int, n) // g[x] 表示 x 的所有邻居
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    res, size := make([]int, n), make([]int, n)
    var dfs func(int, int, int)
    dfs = func(x, fa, depth int) {
        res[0] += depth // depth 为 0 到 x 的距离
        size[x] = 1
        for _, y := range g[x] { // 遍历 x 的邻居 y
            if y != fa { // 避免访问父节点
                dfs(y, x, depth+1) // x 是 y 的父节点
                size[x] += size[y] // 累加 x 的儿子 y 的子树大小
            }
        }
    }
    dfs(0, -1, 0) // 0 没有父节点

    var reroot func(int, int)
    reroot = func(x, fa int) {
        for _, y := range g[x] { // 遍历 x 的邻居 y
            if y != fa { // 避免访问父节点
                res[y] = res[x] + n - 2*size[y]
                reroot(y, x) // x 是 y 的父节点
            }
        }
    }
    reroot(0, -1) // 0 没有父节点
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/23/lc-sumdist1.jpg" / >
    //               0
    //             /   \
    //            1      2
    //                 / | \
    //                3  4  5
    // Input: n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
    // Output: [8,12,6,10,10,10]
    // Explanation: The tree is shown above.
    // We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
    // equals 1 + 1 + 2 + 2 + 2 = 8.
    // Hence, answer[0] = 8, and so on.
    fmt.Println(sumOfDistancesInTree(6,[][]int{{0,1},{0,2},{2,3},{2,4},{2,5}})) // [8,12,6,10,10,10]
    // Example 2:
    // Input: n = 1, edges = []
    // Output: [0]
    fmt.Println(sumOfDistancesInTree(1,[][]int{})) // [0]
    // Example 3:
    //      0
    //     /
    //    1
    // Input: n = 2, edges = [[1,0]]
    // Output: [1,1]
    fmt.Println(sumOfDistancesInTree(2,[][]int{{1,0}})) // [1,1]

    fmt.Println(sumOfDistancesInTree1(6,[][]int{{0,1},{0,2},{2,3},{2,4},{2,5}})) // [8,12,6,10,10,10]
    fmt.Println(sumOfDistancesInTree1(1,[][]int{})) // [0]
    fmt.Println(sumOfDistancesInTree1(2,[][]int{{1,0}})) // [1,1]
}