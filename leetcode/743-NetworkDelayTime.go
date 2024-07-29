package main

// 743. Network Delay Time
// You are given a network of n nodes, labeled from 1 to n. 
// You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), 
// where ui is the source node, vi is the target node, and wi is the time it takes for a signal to travel from source to target.

// We will send a signal from a given node k. 
// Return the minimum time it takes for all the n nodes to receive the signal. 
// If it is impossible for all the n nodes to receive the signal, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png" />
// Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
// Output: 2

// Example 2:
// Input: times = [[1,2,1]], n = 2, k = 1
// Output: 1

// Example 3:
// Input: times = [[1,2,1]], n = 2, k = 2
// Output: -1

// Constraints:
//     1 <= k <= n <= 100
//     1 <= times.length <= 6000
//     times[i].length == 3
//     1 <= ui, vi <= n
//     ui != vi
//     0 <= wi <= 100
//     All the pairs (ui, vi) are unique. (i.e., no multiple edges.)

import "fmt"

// dijkstra
func networkDelayTime(times [][]int, n int, k int) int {
    res, graph := 0, make([]map[int]int, n + 1) 
    for i := 0; i < len(graph); i++ {
        graph[i] = make(map[int]int)
    }
    for _, v := range times {
        graph[v[0]][v[1]] = v[2]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    find := func (graph []map[int]int, minpath []int, visited []bool) (int, int) {
        l, newnode := 1 << 32 - 1, 0
        for i, v := range graph {
            if visited[i] {
                for node, weight := range v {
                    if !visited[node] {
                        if minpath[i] + weight < l {
                            l = minpath[i] + weight
                            newnode = node
                        }
                    }
                }
            }
        }
        return newnode, l
    }
    minpath, visited := make([]int, n + 1), make([]bool, n + 1)
    visited[k] = true
    for {
        node, l := find(graph, minpath, visited)
        if node == 0 {
            break
        }  
        visited[node] = true
        minpath[node] += l
    }
    for i := 1; i < len(visited); i++ {
        if !visited[i] {
            return -1
        }
    }
    for i := 1; i < len(minpath); i++ {
        res = max(res, minpath[i])
    }
    return res
}

func networkDelayTime1(times [][]int, n int, k int) int {
    res, g, inf := 0, make([][]int, n), (1 << 32 - 1) / 10
    for i := range g {
        g[i] = make([]int, n)
        for j := range g[i] {
            g[i][j] = inf
        }
    }
    for _, ch := range times {
        x, y, z := ch[0]-1, ch[1]-1, ch[2]
        g[x][y] = z
    }
    dis := make([]int, n)
    for i := range dis {
        dis[i] = inf
    }
    dis[k-1] = 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    done := make([]bool, n)
    for {
        x := -1
        for i, ok := range done {
            if !ok && (x < 0 || dis[i] <= dis[x]) {
                x = i
            }
        }
        if x < 0 { // 说明所有 n 个元素都更新的完了
            break
        }
        // 这里最好是>= 因为下面更新 dis[y] 时没有做判断，是直接加的，可能会比 inf 大
        // 说是不可达了
        if dis[x] >= inf {
            return -1
        }
        done[x] = true
        for y, d := range g[x] {
            // 这里可以做一步判断，判断是否 >= inf,也可以不判断，因为上面 >= inf 就都认为不可达
            // 也可以不判断
            if d >= inf {
                continue
            }
            dis[y] = min(dis[y], dis[x]+d)
        }
    }
    for _, v := range dis {
        if v > res { res = v}
    }
    return res
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
    // Output: 2
    fmt.Println(networkDelayTime([][]int{{2,1,1},{2,3,1},{3,4,1}}, 4, 2)) // 2
    // Example 2:
    // Input: times = [[1,2,1]], n = 2, k = 1
    // Output: 1
    fmt.Println(networkDelayTime([][]int{{1,2,1}}, 2, 1)) // 1
    // Example 3:
    // Input: times = [[1,2,1]], n = 2, k = 2
    // Output: -1
    fmt.Println(networkDelayTime([][]int{{1,2,1}}, 2, 2)) // -1

    fmt.Println(networkDelayTime1([][]int{{2,1,1},{2,3,1},{3,4,1}}, 4, 2)) // 2
    fmt.Println(networkDelayTime1([][]int{{1,2,1}}, 2, 1)) // 1
    fmt.Println(networkDelayTime1([][]int{{1,2,1}}, 2, 2)) // -1
}