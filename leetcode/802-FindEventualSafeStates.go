package main

// 802. Find Eventual Safe States
// There is a directed graph of n nodes with each node labeled from 0 to n - 1. 
// The graph is represented by a 0-indexed 2D integer array graph where graph[i] is an integer array of nodes adjacent to node i, 
// meaning there is an edge from node i to each node in graph[i].

// A node is a terminal node if there are no outgoing edges. 
// A node is a safe node if every possible path starting from that node leads to a terminal node (or another safe node).

// Return an array containing all the safe nodes of the graph. The answer should be sorted in ascending order.

// Example 1:
// Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
// Output: [2,4,5,6]
// Explanation: The given graph is shown above.
// Nodes 5 and 6 are terminal nodes as there are no outgoing edges from either of them.
// Every path starting at nodes 2, 4, 5, and 6 all lead to either node 5 or 6.

// Example 2:
// Input: graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]
// Output: [4]
// Explanation:
// Only node 4 is a terminal node, and every path starting at node 4 leads to node 4.

// Constraints:
//     n == graph.length
//     1 <= n <= 10^4
//     0 <= graph[i].length <= n
//     0 <= graph[i][j] <= n - 1
//     graph[i] is sorted in a strictly increasing order.
//     The graph may contain self-loops.
//     The number of edges in the graph will be in the range [1, 4 * 10^4].

import "fmt"

// dfs
func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    res, status := []int{}, make([]int, n)
    var check func(int) bool
    check = func(x int) bool {
        if status[x] > 0 { return status[x] == 2; }
        status[x] = 1
        for _, y := range graph[x] {
            if check(y) == false { return false; }
        }
        status[x] = 2
        return true
    }
    for i := 0; i < n; i++ {
        if check(i) {
            res = append(res, i)
        }
    }
    return res
}

// 反向拓扑排序 bfs
func eventualSafeNodes1(graph [][]int) []int {
    // 存在环的时候，不能完成所有课程
    numCourses := len(graph)
    res, outdegree, indegree := []int{}, make([]int, numCourses), make([][]int, numCourses)
    for from, toList := range graph {
        outdegree[from] += len(toList)
        for _, to := range toList {
            indegree[to] = append(indegree[to], from) // 所有的入度
        }
    }
    queue := []int(nil)
    for i, out := range outdegree { // 不指向任何节点的， 可以入队列
        if out == 0 { // 终端节点
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, from := range indegree[cur] { // 指向我的节点
            outdegree[from]--
            if outdegree[from] == 0 {
                queue = append(queue, from) // 这个节点已经没有 往后面的边了
            }
        } 
    }
    for i, out := range outdegree {
        if out == 0 { // 终端节点
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
    // Output: [2,4,5,6]
    // Explanation: The given graph is shown above.
    // Nodes 5 and 6 are terminal nodes as there are no outgoing edges from either of them.
    // Every path starting at nodes 2, 4, 5, and 6 all lead to either node 5 or 6.
    graph1 := [][]int{{1,2},{2,3},{5},{0},{5},{},{}}
    fmt.Println(eventualSafeNodes(graph1)) // [2,4,5,6]
    // Example 2:
    // Input: graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]
    // Output: [4]
    // Explanation:
    // Only node 4 is a terminal node, and every path starting at node 4 leads to node 4.
    graph2 := [][]int{{1,2,3,4},{1,2},{3,4},{0,4},{}}
    fmt.Println(eventualSafeNodes(graph2)) // [4]

    fmt.Println(eventualSafeNodes1(graph1)) // [2,4,5,6]
    fmt.Println(eventualSafeNodes1(graph2)) // [4]
}