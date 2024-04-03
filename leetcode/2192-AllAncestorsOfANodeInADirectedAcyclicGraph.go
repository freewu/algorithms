package main

// 2192. All Ancestors of a Node in a Directed Acyclic Graph
// You are given a positive integer n representing the number of nodes of a Directed Acyclic Graph (DAG). 
// The nodes are numbered from 0 to n - 1 (inclusive).

// You are also given a 2D integer array edges, where edges[i] = [fromi, toi] denotes that there is a unidirectional edge from fromi to toi in the graph.
// Return a list answer, where answer[i] is the list of ancestors of the ith node, sorted in ascending order.
// A node u is an ancestor of another node v if u can reach v via a set of edges.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/12/12/e1.png" />
// Input: n = 8, edgeList = [[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]
// Output: [[],[],[],[0,1],[0,2],[0,1,3],[0,1,2,3,4],[0,1,2,3]]
// Explanation:
// The above diagram represents the input graph.
// - Nodes 0, 1, and 2 do not have any ancestors.
// - Node 3 has two ancestors 0 and 1.
// - Node 4 has two ancestors 0 and 2.
// - Node 5 has three ancestors 0, 1, and 3.
// - Node 6 has five ancestors 0, 1, 2, 3, and 4.
// - Node 7 has four ancestors 0, 1, 2, and 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/12/12/e2.png" />
// Input: n = 5, edgeList = [[0,1],[0,2],[0,3],[0,4],[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
// Output: [[],[0],[0,1],[0,1,2],[0,1,2,3]]
// Explanation:
// The above diagram represents the input graph.
// - Node 0 does not have any ancestor.
// - Node 1 has one ancestor 0.
// - Node 2 has two ancestors 0 and 1.
// - Node 3 has three ancestors 0, 1, and 2.
// - Node 4 has four ancestors 0, 1, 2, and 3.
 
// Constraints:
//     1 <= n <= 1000
//     0 <= edges.length <= min(2000, n * (n - 1) / 2)
//     edges[i].length == 2
//     0 <= fromi, toi <= n - 1
//     fromi != toi
//     There are no duplicate edges.
//     The graph is directed and acyclic.

import "fmt"
import "container/list"

// bfs
func getAncestors(n int, edges [][]int) [][]int {
    graph := make(map[int][]int)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
    }
    bfs := func(res [][]int, graph map[int][]int, parent int) {
        seen := map[int]bool{ parent: true }
        queue := list.New()
        queue.PushBack(parent)
        for queue.Len() > 0 {
            current := queue.Front().Value.(int)
            queue.Remove(queue.Front())
            for _, children := range graph[current] {
                if !seen[children] {
                    seen[children] = true
                    queue.PushBack(children)
                    res[children] = append(res[children], parent)
                }
            }
        }
    }
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        bfs(res, graph, i)
    }
    return res
}

// dfs
func getAncestors1(n int, edges [][]int) [][]int {
    graph := make(map[int][]int)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
    }
    dfs := func(res [][]int, graph map[int][]int, parent int) {
        seen := map[int]bool{parent: true}
        stack := list.New()
        stack.PushBack(parent)
        for stack.Len() > 0 {
            current := stack.Back().Value.(int)
            stack.Remove(stack.Back())
            for _, children := range graph[current] {
                if !seen[children] {
                    seen[children] = true
                    stack.PushBack(children)
                    res[children] = append(res[children], parent)
                }
            }
        }
    }
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        dfs(res, graph, i)
    }
    return res
}

// dfs
func getAncestors2(n int, edges [][]int) [][]int {
    m := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        m[y] = append(m[y], x)
    }
    vis := make([]bool, n)
    var dfs func(int)
    dfs = func(x int) {
        vis[x] = true // 避免重复访问
        for _, y := range m[x] {
            if !vis[y] {
                dfs(y) // 只递归没有访问过的点
            }
        }
    }
    res := make([][]int, n)
    for i := range res {
        clear(vis)
        dfs(i) // 从 i 开始 DFS
        vis[i] = false // ans[i] 不含 i
        for j, b := range vis {
            if b {
                res[i] = append(res[i], j)
            }
        }
    }
    return res
}

func main() {
    // The above diagram represents the input graph.
    // - Nodes 0, 1, and 2 do not have any ancestors.
    // - Node 3 has two ancestors 0 and 1.
    // - Node 4 has two ancestors 0 and 2.
    // - Node 5 has three ancestors 0, 1, and 3.
    // - Node 6 has five ancestors 0, 1, 2, 3, and 4.
    // - Node 7 has four ancestors 0, 1, 2, and 3.
    fmt.Println(getAncestors(8,[][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}})) // [[],[],[],[0,1],[0,2],[0,1,3],[0,1,2,3,4],[0,1,2,3]]

    // The above diagram represents the input graph.
    // - Node 0 does not have any ancestor.
    // - Node 1 has one ancestor 0.
    // - Node 2 has two ancestors 0 and 1.
    // - Node 3 has three ancestors 0, 1, and 2.
    // - Node 4 has four ancestors 0, 1, 2, and 3.
    fmt.Println(getAncestors(5,[][]int{{0,1},{0,2},{0,3},{0,4},{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}})) // [[],[0],[0,1],[0,1,2],[0,1,2,3]]
    
    fmt.Println(getAncestors1(8,[][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}})) // [[],[],[],[0,1],[0,2],[0,1,3],[0,1,2,3,4],[0,1,2,3]]
    fmt.Println(getAncestors1(5,[][]int{{0,1},{0,2},{0,3},{0,4},{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}})) // [[],[0],[0,1],[0,1,2],[0,1,2,3]]

    fmt.Println(getAncestors2(8,[][]int{{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6}})) // [[],[],[],[0,1],[0,2],[0,1,3],[0,1,2,3,4],[0,1,2,3]]
    fmt.Println(getAncestors2(5,[][]int{{0,1},{0,2},{0,3},{0,4},{1,2},{1,3},{1,4},{2,3},{2,4},{3,4}})) // [[],[0],[0,1],[0,1,2],[0,1,2,3]]

}