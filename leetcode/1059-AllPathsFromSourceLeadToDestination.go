package main

// 1059. All Paths from Source Lead to Destination
// Given the edges of a directed graph where edges[i] = [ai, bi] indicates there is an edge between nodes ai and bi, 
// and two nodes source and destination of this graph, determine whether or not all paths starting from source eventually, 
// end at destination, that is:
//     At least one path exists from the source node to the destination node
//     If a path exists from the source node to a node with no outgoing edges, then that node is equal to destination.
//     The number of possible paths from source to destination is a finite number.

// Return true if and only if all roads from source lead to destination.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/03/16/485_example_1.png" />
// Input: n = 3, edges = [[0,1],[0,2]], source = 0, destination = 2
// Output: false
// Explanation: It is possible to reach and get stuck on both node 1 and node 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/03/16/485_example_2.png" />
// Input: n = 4, edges = [[0,1],[0,3],[1,2],[2,1]], source = 0, destination = 3
// Output: false
// Explanation: We have two possibilities: to end at node 3, or to loop over node 1 and node 2 indefinitely.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/03/16/485_example_3.png" />
// Input: n = 4, edges = [[0,1],[0,2],[1,3],[2,3]], source = 0, destination = 3
// Output: true

// Constraints:
//     1 <= n <= 10^4
//     0 <= edges.length <= 10^4
//     edges.length == 2
//     0 <= ai, bi <= n - 1
//     0 <= source <= n - 1
//     0 <= destination <= n - 1
//     The given graph may have self-loops and parallel edges.

import "fmt"

// 拓扑排序 bfs
func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
    g, deg := make([][]int, n), make([]int, n) 
    for _, e := range edges { // 统计出度
        a, b := e[0], e[1]
        g[b] = append(g[b], a) // 反向建图, 列表为父节点
        deg[a]++
    }
    if deg[destination] > 0 {
        return false
    }
    q := []int{ destination } // 将 dest 入队列, 出度必须为 0, 因为 拓扑排序 是逐渐去掉叶子节点的算法
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        if source == cur {
            return true
        }
        for _, x := range g[cur] {
            deg[x]-- // 去掉当前叶子节点影响, 因为反向(入度)建图, 其父节点出度 -1
            if deg[x] == 0 {
                q = append(q, x)
            }
        }
    }
    return false // 队列为空, 也没找到 source
}

// dfs
func leadsToDestination1(n int, edges [][]int, source int, destination int) bool {
    a := make([]int, n)
    a[destination] = 1
    mp, visited := make([][]int, n), make([]bool, n)
    for _, e := range edges {
        mp[e[0]] = append(mp[e[0]], e[1])
    }
    if len(mp[destination]) > 0 {
        return false
    }
    var dfs func(cur int, mp [][]int, a []int, visited []bool)
    dfs = func(cur int, mp [][]int, a []int, visited []bool) {
        if a[cur] != 0 {
            return
        }
        if visited[cur] {
            a[cur] = -1
            return
        }
        visited[cur] = true
        tag := -1
        for _, next := range mp[cur] {
            dfs(next, mp, a, visited)
            if a[next] == 1 {
                tag = 1
            } else {
                tag = -1
                break
            }
        }
        a[cur] = tag
    }
    dfs(source, mp, a, visited)
    return a[source] == 1
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/03/16/485_example_1.png" />
    // Input: n = 3, edges = [[0,1],[0,2]], source = 0, destination = 2
    // Output: false
    // Explanation: It is possible to reach and get stuck on both node 1 and node 2.
    fmt.Println(leadsToDestination(3,[][]int{{0,1},{0,2}},0,2)) // false
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/03/16/485_example_2.png" />
    // Input: n = 4, edges = [[0,1],[0,3],[1,2],[2,1]], source = 0, destination = 3
    // Output: false
    // Explanation: We have two possibilities: to end at node 3, or to loop over node 1 and node 2 indefinitely.
    fmt.Println(leadsToDestination(4,[][]int{{0,1},{0,3},{1,2},{2,1}},0,3)) // false
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/03/16/485_example_3.png" />
    // Input: n = 4, edges = [[0,1],[0,2],[1,3],[2,3]], source = 0, destination = 3
    // Output: true
    fmt.Println(leadsToDestination(4,[][]int{{0,1},{0,2},{1,3},{2,3}},0,3)) // true

    fmt.Println(leadsToDestination1(3,[][]int{{0,1},{0,2}},0,2)) // false
    fmt.Println(leadsToDestination1(4,[][]int{{0,1},{0,3},{1,2},{2,1}},0,3)) // false
    fmt.Println(leadsToDestination1(4,[][]int{{0,1},{0,2},{1,3},{2,3}},0,3)) // true
}