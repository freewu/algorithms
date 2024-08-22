package main

// LCR 110. 所有可能的路径
// 给定一个有 n 个节点的有向无环图，用二维数组 graph 表示，请找到所有从 0 到 n-1 的路径并输出（不要求按顺序）。

// graph 的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些结点（译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a ），
// 若为空，就是没有下一个节点了。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg" />
// 输入：graph = [[1,2],[3],[3],[]]
// 输出：[[0,1,3],[0,2,3]]
// 解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2020/09/28/all_2.jpg" />
// 输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
// 输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

// 示例 3：
// 输入：graph = [[1],[]]
// 输出：[[0,1]]

// 示例 4：
// 输入：graph = [[1,2,3],[2],[3],[]]
// 输出：[[0,1,2,3],[0,2,3],[0,3]]

// 示例 5：
// 输入：graph = [[1,3],[2],[3],[]]
// 输出：[[0,1,2,3],[0,3]]

// 提示：
//     n == graph.length
//     2 <= n <= 15
//     0 <= graph[i][j] < n
//     graph[i][j] != i 
//     保证输入为有向无环图 (GAD)

import "fmt"

// dfs
func allPathsSourceTarget(graph [][]int) [][]int {
    var res [][]int
    var dfs func([]int, int)
    dfs = func(path []int, u int) {
        path = append(path, u)
        if u == len(graph)-1 {
            res = append(res, append([]int{}, path...))
        }
        for _, v := range graph[u] {
            dfs(path, v)
        }
    }
    dfs([]int{}, 0)
    return res
}

// bfs stack
func allPathsSourceTarget1(graph [][]int) [][]int {
    res, stack := [][]int{}, [][]int{{0}}
    for len(stack) > 0 {
        p := stack[len(stack)-1]
        u := p[len(p)-1]
        stack = stack[:len(stack)-1]
        if u == len(graph)-1 {
            res = append(res, p)
        }
        for _, v := range graph[u] {
            vp := append(p[:len(p):len(p)], v)
            stack = append(stack, vp)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg" />
    // Input: graph = [[1,2],[3],[3],[]]
    // Output: [[0,1,3],[0,2,3]]
    // Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
    fmt.Println(allPathsSourceTarget([][]int{{1,2},{3},{3},{}})) // [[0,1,3],[0,2,3]]
    // Example 2:
    // Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
    // Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
    fmt.Println(allPathsSourceTarget([][]int{{4,3,1},{3,2,4},{3},{4},{}})) // [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

    fmt.Println(allPathsSourceTarget1([][]int{{1,2},{3},{3},{}})) // [[0,1,3],[0,2,3]]
    fmt.Println(allPathsSourceTarget1([][]int{{4,3,1},{3,2,4},{3},{4},{}})) // [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
}