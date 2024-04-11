package main

// 797. All Paths From Source to Target
// Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, 
// find all possible paths from node 0 to node n - 1 and return them in any order.

// The graph is given as follows: graph[i] is a list of all nodes you can visit from node i 
// (i.e., there is a directed edge from node i to node graph[i][j]).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg" />
// Input: graph = [[1,2],[3],[3],[]]
// Output: [[0,1,3],[0,2,3]]
// Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.

// Example 2:
// Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
// Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
 
// Constraints:
//     n == graph.length
//     2 <= n <= 15
//     0 <= graph[i][j] < n
//     graph[i][j] != i (i.e., there will be no self-loops).
//     All the elements of graph[i] are unique.
//     The input graph is guaranteed to be a DAG.

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
    // Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
    fmt.Println(allPathsSourceTarget([][]int{{1,2},{3},{3},{}})) // [[0,1,3],[0,2,3]]
    fmt.Println(allPathsSourceTarget([][]int{{4,3,1},{3,2,4},{3},{4},{}})) // [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

    fmt.Println(allPathsSourceTarget1([][]int{{1,2},{3},{3},{}})) // [[0,1,3],[0,2,3]]
    fmt.Println(allPathsSourceTarget1([][]int{{4,3,1},{3,2,4},{3},{4},{}})) // [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

}