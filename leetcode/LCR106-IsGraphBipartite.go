package main

// LCR 106. 判断二分图
// 存在一个 无向图 ，图中有 n 个节点。其中每个节点都有一个介于 0 到 n - 1 之间的唯一编号。

// 给定一个二维数组 graph ，表示图，其中 graph[u] 是一个节点数组，由节点 u 的邻接节点组成。
// 形式上，对于 graph[u] 中的每个 v ，都存在一条位于节点 u 和节点 v 之间的无向边。
// 该无向图同时具有以下属性：
//     不存在自环（graph[u] 不包含 u）。
//     不存在平行边（graph[u] 不包含重复值）。
//     如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
//     这个图可能不是连通图，也就是说两个节点 u 和 v 之间可能不存在一条连通彼此的路径。

// 二分图 定义：
//     如果能将一个图的节点集合分割成两个独立的子集 A 和 B ，
//     并使图中的每一条边的两个节点一个来自 A 集合，一个来自 B 集合，就将这个图称为 二分图 。

// 如果图是二分图，返回 true ；否则，返回 false 。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/10/21/bi2.jpg" />
// 输入：graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
// 输出：false
// 解释：不能将节点分割成两个独立的子集，以使每条边都连通一个子集中的一个节点与另一个子集中的一个节点。

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2020/10/21/bi1.jpg" />
// 输入：graph = [[1,3],[0,2],[1,3],[0,2]]
// 输出：true
// 解释：可以将节点分成两组: {0, 2} 和 {1, 3} 。

// 提示：
//     graph.length == n
//     1 <= n <= 100
//     0 <= graph[u].length < n
//     0 <= graph[u][i] <= n - 1
//     graph[u] 不会包含 u
//     graph[u] 的所有值 互不相同
//     如果 graph[u] 包含 v，那么 graph[v] 也会包含 u

import "fmt"

func isBipartite(graph [][]int) bool {
    dp := make([]int, len(graph))
    for i := range dp {
        dp[i] = -1
    }
    var dfs func(node int, cur int, dp []int, graph [][]int) bool 
    dfs = func(node int, cur int, dp []int, graph [][]int) bool {
        dp[node] = cur
        for _, neighbor := range graph[node] {
            // 如果 v 在 graph[u] 内，那么 u 也应该在 graph[v] 内（该图是无向图）
            if dp[neighbor] == -1 {
                if dfs(neighbor, 1-cur, dp, graph) == false {
                    return false
                }
            } else if dp[neighbor] == cur { // 不存在自环（graph[u] 不包含 u）
                return false
            }
        }
        return true
    }
    for i := range dp { 
        if dp[i] == -1 {
            if dfs(i, 0, dp, graph) == false {
                return false
            }
        }
    }
    return true
}

func isBipartite1(graph [][]int) bool {
    n, res := len(graph), true
    color, visited := make([]bool, n),make([]bool, n) // color 记录图中节点的颜色，两种; visited 记录是否已经被访问过了
    var traverse func(x int)
    traverse = func(x int) {
        if !res {
            return
        }
        visited[x] = true
        for _, y := range graph[x] {
            if !visited[y] {// 相邻的节点y没有被访问过，则给y标记成和x不同的颜色
                color[y] = !color[x]
                traverse(y)
            } else {// 如果y已经被访问过了，并且x和y的颜色一样，那就不是二分图
                if color[y] == color[x] {
                    res = false
                    break
                }
            }
        }
    }
    for x := 0; x < n; x++ {
        if !visited[x] {
            traverse(x)
        }
    }
    return res
}

func isBipartite2(graph [][]int) bool {
    // 初始值 0   红 1 蓝2
    color := make([]int, len(graph))
    // 给index位置染色preColor
    var dfs func(index, preColor int) bool
    dfs = func(index, preColor int) bool {
        if color[index] != 0 {
            return color[index] == preColor
        }
        color[index] = preColor
        for _, v := range graph[index] {
            if preColor == 1 {
                if !dfs(v, 2) {
                    return false
                }
            } else {
                if !dfs(v, 1) {
                    return false
                }
            }
        }
        return true
    }
    for index, c := range color {
        if c == 0 {
            if !dfs(index, 1) {
                return false
            }
        }
    }
    return true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/10/21/bi2.jpg" />
    // Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
    // Output: false
    // Explanation: There is no way to partition the nodes into two independent sets such that every edge connects a node in one and a node in the other.
    fmt.Println(isBipartite([][]int{{1,2,3},{0,2},{0,1,3},{0,2}})) // false
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/10/21/bi1.jpg" />
    // Input: graph = [[1,3],[0,2],[1,3],[0,2]]
    // Output: true
    // Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.
    fmt.Println(isBipartite([][]int{{1,3},{0,2},{1,3},{0,2}})) // true

    fmt.Println(isBipartite1([][]int{{1,2,3},{0,2},{0,1,3},{0,2}})) // false
    fmt.Println(isBipartite1([][]int{{1,3},{0,2},{1,3},{0,2}})) // true

    fmt.Println(isBipartite2([][]int{{1,2,3},{0,2},{0,1,3},{0,2}})) // false
    fmt.Println(isBipartite2([][]int{{1,3},{0,2},{1,3},{0,2}})) // true
}