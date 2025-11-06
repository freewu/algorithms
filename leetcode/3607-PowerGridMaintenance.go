package main

// 3607. Power Grid Maintenance
// You are given an integer c representing c power stations, each with a unique identifier id from 1 to c (1‑based indexing).

// These stations are interconnected via n bidirectional cables, represented by a 2D array connections, where each element connections[i] = [ui, vi] indicates a connection between station ui and station vi. 
// Stations that are directly or indirectly connected form a power grid.

// Initially, all stations are online (operational).

// You are also given a 2D array queries, where each query is one of the following two types:

//     1. [1, x]: A maintenance check is requested for station x. 
//        If station x is online, it resolves the check by itself. 
//        If station x is offline, the check is resolved by the operational station with the smallest id in the same power grid as x. 
//        If no operational station exists in that grid, return -1.

//     2. [2, x]: Station x goes offline (i.e., it becomes non-operational).

// Return an array of integers representing the results of each query of type [1, x] in the order they appear.

// Note: The power grid preserves its structure; an offline (non‑operational) node remains part of its grid and taking it offline does not alter connectivity.

// Example 1:
// Input: c = 5, connections = [[1,2],[2,3],[3,4],[4,5]], queries = [[1,3],[2,1],[1,1],[2,2],[1,2]]
// Output: [3,2,3]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/04/15/powergrid.jpg" />
// Initially, all stations {1, 2, 3, 4, 5} are online and form a single power grid.
// Query [1,3]: Station 3 is online, so the maintenance check is resolved by station 3.
// Query [2,1]: Station 1 goes offline. The remaining online stations are {2, 3, 4, 5}.
// Query [1,1]: Station 1 is offline, so the check is resolved by the operational station with the smallest id among {2, 3, 4, 5}, which is station 2.
// Query [2,2]: Station 2 goes offline. The remaining online stations are {3, 4, 5}.
// Query [1,2]: Station 2 is offline, so the check is resolved by the operational station with the smallest id among {3, 4, 5}, which is station 3.

// Example 2:
// Input: c = 3, connections = [], queries = [[1,1],[2,1],[1,1]]
// Output: [1,-1]
// Explanation:
// There are no connections, so each station is its own isolated grid.
// Query [1,1]: Station 1 is online in its isolated grid, so the maintenance check is resolved by station 1.
// Query [2,1]: Station 1 goes offline.
// Query [1,1]: Station 1 is offline and there are no other stations in its grid, so the result is -1.

// Constraints:
//     1 <= c <= 10^5
//     0 <= n == connections.length <= min(10^5, c * (c - 1) / 2)
//     connections[i].length == 2
//     1 <= ui, vi <= c
//     ui != vi
//     1 <= queries.length <= 2 * 10^5
//     queries[i].length == 2
//     queries[i][0] is either 1 or 2.
//     1 <= queries[i][1] <= c

import "fmt"
import "sort"
//import "slices"

func processQueries(c int, connections [][]int, queries [][]int) []int {
    g := make([][]int, c+1)
    for _, e := range connections {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    belong := make([]int, c+1)
    for i := range belong {
        belong[i] = -1
    }
    cc := 0 // 连通块编号
    var dfs func(int)
    dfs = func(x int) {
        belong[x] = cc // 记录节点 x 在哪个连通块
        for _, y := range g[x] {
            if belong[y] < 0 {
                dfs(y)
            }
        }
    }
    for i := 1; i <= c; i++ {
        if belong[i] < 0 {
            dfs(i)
            cc++
        }
    }
    offlineTime := make([]int, c+1)
    for i := range offlineTime {
        offlineTime[i] = 1 << 31
    }
    q1 := 0
    //for i, q := range slices.Backward(queries) {
    for i := len(queries) - 1; i >= 0; i-- {
        q := queries[i]
        if q[0] == 2 {
            offlineTime[q[1]] = i // 记录最早离线时间
        } else {
            q1++
        }
    }
    // 维护每个连通块的在线电站的最小编号
    mn := make([]int, cc)
    for i := range mn {
        mn[i] = 1 << 31
    }
    for i := 1; i <= c; i++ {
        if offlineTime[i] == 1 << 31 { // 最终仍然在线
            j := belong[i]
            mn[j] = min(mn[j], i)
        }
    }
    res := make([]int, q1)
    //for i, q := range slices.Backward(queries) {
    for i := len(queries) - 1; i >= 0; i-- {
        q := queries[i]
        x := q[1]
        j := belong[x]
        if q[0] == 2 {
            if offlineTime[x] == i { // 变回在线
                mn[j] = min(mn[j], x)
            }
        } else {
            q1--
            if i < offlineTime[x] { // 已经在线（写 < 或者 <= 都可以）
                res[q1] = x
            } else if mn[j] != 1 << 31 {
                res[q1] = mn[j]
            } else {
                res[q1] = -1
            }
        }
    }
    return res
}


func processQueries1(c int, connections [][]int, queries [][]int) []int {
    fa := make([]int, c + 1)
    for i := range fa {
        fa[i] = i
    }
    var find func(x int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }
    var merge func(a, b int)
    merge = func(a, b int) {
        a, b = find(a), find(b)
        if a == b { return }
        if a > b {
            a, b = b, a
        }
        fa[b] = a
    }
    for _, c := range connections {
        merge(c[0], c[1])
    }
    belong, belongList := make([]int, c+1), make([][]int, c+1)
    for i := range belong {
        belong[i] = find(i)
        belongList[belong[i]] = append(belongList[belong[i]], i)
    }
    for _, v := range belongList {
        sort.Ints(v)
    }
    res, flag := make([]int, 0, len(queries)), make([]bool, c + 1)
    for _, v := range queries {
        switch v[0] {
        case 1:
            if !flag[v[1]] {
                res = append(res, v[1])
                continue
            }
            v[1] = find(v[1])
            for len(belongList[v[1]]) > 0 && flag[belongList[v[1]][0]] {
                belongList[v[1]] = belongList[v[1]][1:]
            }
            if len(belongList[v[1]]) > 0 {
                res = append(res, belongList[v[1]][0])
            } else {
                res = append(res, -1)
            }
        case 2:
            flag[v[1]] = true
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: c = 5, connections = [[1,2],[2,3],[3,4],[4,5]], queries = [[1,3],[2,1],[1,1],[2,2],[1,2]]
    // Output: [3,2,3]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/04/15/powergrid.jpg" />
    // Initially, all stations {1, 2, 3, 4, 5} are online and form a single power grid.
    // Query [1,3]: Station 3 is online, so the maintenance check is resolved by station 3.
    // Query [2,1]: Station 1 goes offline. The remaining online stations are {2, 3, 4, 5}.
    // Query [1,1]: Station 1 is offline, so the check is resolved by the operational station with the smallest id among {2, 3, 4, 5}, which is station 2.
    // Query [2,2]: Station 2 goes offline. The remaining online stations are {3, 4, 5}.
    // Query [1,2]: Station 2 is offline, so the check is resolved by the operational station with the smallest id among {3, 4, 5}, which is station 3.
    fmt.Println(processQueries(5, [][]int{{1,2},{2,3},{3,4},{4,5}}, [][]int{{1,3},{2,1},{1,1},{2,2},{1,2}})) // [3,2,3]
    // Example 2:
    // Input: c = 3, connections = [], queries = [[1,1],[2,1],[1,1]]
    // Output: [1,-1]
    // Explanation:
    // There are no connections, so each station is its own isolated grid.
    // Query [1,1]: Station 1 is online in its isolated grid, so the maintenance check is resolved by station 1.
    // Query [2,1]: Station 1 goes offline.
    // Query [1,1]: Station 1 is offline and there are no other stations in its grid, so the result is -1.
    fmt.Println(processQueries(3, [][]int{}, [][]int{{1,1},{2,1},{1,1}})) // [1,-1]

    fmt.Println(processQueries1(5, [][]int{{1,2},{2,3},{3,4},{4,5}}, [][]int{{1,3},{2,1},{1,1},{2,2},{1,2}})) // [3,2,3]
    fmt.Println(processQueries1(3, [][]int{}, [][]int{{1,1},{2,1},{1,1}})) // [1,-1]
}