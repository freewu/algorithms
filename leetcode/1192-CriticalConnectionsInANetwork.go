package main 

// 1192. Critical Connections in a Network
// There are n servers numbered from 0 to n - 1 connected by undirected server-to-server connections forming a network where connections[i] = [ai, bi] represents a connection between servers ai and bi. 
// Any server can reach other servers directly or indirectly through the network.
// A critical connection is a connection that, if removed, will make some servers unable to reach some other server.

// Return all critical connections in the network in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/09/03/1537_ex1_2.png" />
// Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
// Output: [[1,3]]
// Explanation: [[3,1]] is also accepted.

// Example 2:
// Input: n = 2, connections = [[0,1]]
// Output: [[0,1]]

// Constraints:
//     2 <= n <= 10^5
//     n - 1 <= connections.length <= 10^5
//     0 <= ai, bi <= n - 1
//     ai != bi
//     There are no repeated connections.

import "fmt"

// tarjan 
// https://blog.csdn.net/qq_30277239/article/details/118683637
// https://baike.baidu.com/item/tarjan%E7%AE%97%E6%B3%95/10687825
func criticalConnections(n int, connections [][]int) [][]int {
    graph_adj, crit := make([][]int, n), [][]int{}
    for _, conn := range(connections) {
        graph_adj[conn[0]] = append(graph_adj[conn[0]], conn[1])
        graph_adj[conn[1]] = append(graph_adj[conn[1]], conn[0])
    }
    var findPart func(curr, last int, depth int, parts *[]int, graph_adj *[][]int, crit *[][]int)
    findPart = func(curr, last int, depth int, parts *[]int, graph_adj *[][]int, crit *[][]int) {
        (*parts)[curr] = depth
        for _, next := range((*graph_adj)[curr]) {
            if next == last {
                continue
            }
            if (*parts)[next] != 0 {
                if (*parts)[next] < (*parts)[curr] {
                    (*parts)[curr] = (*parts)[next]
                }
            } else {
                findPart(next, curr, depth + 1, parts, graph_adj, crit)
                if (*parts)[next] < (*parts)[curr] {
                    (*parts)[curr] = (*parts)[next]
                } else if depth < (*parts)[next] {
                    *crit = append(*crit, []int{curr, next})
                }
            }
        }
    }
    var parts = make([]int, n)
    findPart(0, -1, 1, &parts, &graph_adj, &crit)
    return crit
}

func criticalConnections1(n int, connections [][]int) [][]int {
    graph := make([][]int, n)
    for _, conn := range connections {
        graph[conn[0]] = append(graph[conn[0]], conn[1])
        graph[conn[1]] = append(graph[conn[1]], conn[0])
    }
    dfn, low, visited := make([]int, n), make([]int, n),make([]bool, n)
    res, time := [][]int{}, 0 
    var tarjan func(node, parent int)
    tarjan = func(node, parent int) {
        time++
        dfn[node], low[node] = time, time
        visited[node] = true
        for _, e := range graph[node] {
            if e == parent {
                continue
            }
            if !visited[e] {
                tarjan(e,node)
                low[node] = min(low[node], low[e])
                if low[e] > dfn[node] {
                    res = append(res, []int{node, e})
                }
            } else {
                low[node] = min(low[node],low[e])
            }
        }
    }
    tarjan(0,-1)
    return res
}

func main() {
    fmt.Println(criticalConnections(4,[][]int{{0,1},{1,2},{2,0},{1,3}})) //  [[1,3]]
    fmt.Println(criticalConnections(2,[][]int{{0,1}})) // [[0,1]]

    fmt.Println(criticalConnections1(4,[][]int{{0,1},{1,2},{2,0},{1,3}})) //  [[1,3]]
    fmt.Println(criticalConnections1(2,[][]int{{0,1}})) // [[0,1]]
}