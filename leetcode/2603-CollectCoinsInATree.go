package main

// 2603. Collect Coins in a Tree
// There exists an undirected and unrooted tree with n nodes indexed from 0 to n - 1. 
// You are given an integer n and a 2D integer array edges of length n - 1, where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree. 
// You are also given an array coins of size n where coins[i] can be either 0 or 1, where 1 indicates the presence of a coin in the vertex i.

// Initially, you choose to start at any vertex in the tree. 
// Then, you can perform the following operations any number of times: 
//     1. Collect all the coins that are at a distance of at most 2 from the current vertex, or
//     2. Move to any adjacent vertex in the tree.

// Find the minimum number of edges you need to go through to collect all the coins and go back to the initial vertex.

// Note that if you pass an edge several times, you need to count it into the answer several times.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/03/01/graph-2.png" />
// Input: coins = [1,0,0,0,0,1], edges = [[0,1],[1,2],[2,3],[3,4],[4,5]]
// Output: 2
// Explanation: Start at vertex 2, collect the coin at vertex 0, move to vertex 3, collect the coin at vertex 5 then move back to vertex 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/03/02/graph-4.png" />
// Input: coins = [0,0,0,1,1,0,0,1], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]]
// Output: 2
// Explanation: Start at vertex 0, collect the coins at vertices 4 and 3, move to vertex 2,  collect the coin at vertex 7, then move back to vertex 0.

// Constraints:
//     n == coins.length
//     1 <= n <= 3 * 10^4
//     0 <= coins[i] <= 1
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     edges represents a valid tree.

import "fmt"

func collectTheCoins(coins []int, edges [][]int) int {
    n := len(coins)
    graph := make([][]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    removeElement := func(arr []int, element int) []int {
        index := -1
        for i, e := range arr {
            if e == element {
                index = i
                break
            }
        }
        if index == -1 { return arr }
        return append(arr[:index], arr[index + 1:]...)
    }
    leaf := make([]int, 0)
    for i := 0; i < n; i++ {
        u := i
        for len(graph[u]) == 1 && coins[u] == 0 {
            v := graph[u][0]
            graph[u] = graph[u][1:]
            graph[v] = removeElement(graph[v], u)
            u = v
        }
        if len(graph[u]) == 1 {
            leaf = append(leaf, u)
        }
    }
    for k := 2; k > 0; k-- {
        for l := len(leaf); l > 0; l-- {
            u := leaf[0]
            leaf = leaf[1:]
            if len(graph[u]) != 0 {
                v := graph[u][0]
                graph[u] = graph[u][1:]
                graph[v] = removeElement(graph[v], u)
                if len(graph[v]) == 1 {
                    leaf = append(leaf, v)
                }
            }
        }
    }
    res := 0
    for _, row := range graph {
        res += len(row)
    }
    return res
}

func collectTheCoins1(coins []int, edges [][]int) int {
    n := len(coins)
    graph, ingree := make([][]int, n), make([]int, n)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
        ingree[v[0]]++
        ingree[v[1]]++ // 统计每个节点的度数（邻居个数）
    }
    left := n - 1 // 剩余边数
    queue := []int{}
    for i, v := range ingree {
        if v == 1 && coins[i] == 0 { // 没有金币的叶子
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 { // 拓扑排序，去掉没有金币的子树
        for i := len(queue); i > 0; i-- {
            top := queue[0]
            queue = queue[1:]
            ingree[top]--
            left-- // 删除节点 top 到其父节点的边
            for _, v := range graph[top] {
                ingree[v]--
                if ingree[v] == 1 && coins[v] == 0 {  // 没有金币的叶子
                    queue = append(queue, v)
                }
            }
        }
    }
    for i, v := range ingree { // 再次拓扑排序
        if v == 1 && coins[i] == 1 { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
            queue = append(queue, i)
        }
    }
    left -= len(queue) // 删除所有叶子（到其父节点的边）
    for _, i := range queue {
        for _, j := range graph[i] {
            ingree[j]--
            if ingree[j] == 1 { // j 现在是叶子了
                left-- // 删除 j（到其父节点的边）
            }
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return max(left * 2, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/03/01/graph-2.png" />
    // Input: coins = [1,0,0,0,0,1], edges = [[0,1],[1,2],[2,3],[3,4],[4,5]]
    // Output: 2
    // Explanation: Start at vertex 2, collect the coin at vertex 0, move to vertex 3, collect the coin at vertex 5 then move back to vertex 2.
    fmt.Println(collectTheCoins([]int{1,0,0,0,0,1}, [][]int{{0,1},{1,2},{2,3},{3,4},{4,5}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/03/02/graph-4.png" />
    // Input: coins = [0,0,0,1,1,0,0,1], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5,7]]
    // Output: 2
    // Explanation: Start at vertex 0, collect the coins at vertices 4 and 3, move to vertex 2,  collect the coin at vertex 7, then move back to vertex 0.
    fmt.Println(collectTheCoins([]int{0,0,0,1,1,0,0,1}, [][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{5,6},{5,7}})) // 2

    fmt.Println(collectTheCoins1([]int{1,0,0,0,0,1}, [][]int{{0,1},{1,2},{2,3},{3,4},{4,5}})) // 2
    fmt.Println(collectTheCoins1([]int{0,0,0,1,1,0,0,1}, [][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{5,6},{5,7}})) // 2
}