package main

// 1443. Minimum Time to Collect All Apples in a Tree
// Given an undirected tree consisting of n vertices numbered from 0 to n-1, which has some apples in their vertices. 
// You spend 1 second to walk over one edge of the tree. 
// Return the minimum time in seconds you have to spend to collect all apples in the tree, starting at vertex 0 and coming back to this vertex.

// The edges of the undirected tree are given in the array edges, 
// where edges[i] = [ai, bi] means that exists an edge connecting the vertices ai and bi. 
// Additionally, there is a boolean array hasApple, where hasApple[i] = true means that vertex i has an apple; 
// otherwise, it does not have any apple.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/04/23/min_time_collect_apple_1.png" />
// Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,true,true,false]
// Output: 8 
// Explanation: The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/04/23/min_time_collect_apple_2.png" />
// Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,false,true,false]
// Output: 6
// Explanation: The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  

// Example 3:
// Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,false,false,false,false,false]
// Output: 0

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= ai < bi <= n - 1
//     hasApple.length == n

import "fmt"

// func minTime(n int, edges [][]int, hasApple []bool) int {
//     time, prev, step := -1, make([]int, n), make([]bool, n)
//     for _, e := range edges {
//         prev[e[1]] = e[0]
//     }
//     step[0] = true
//     for i, has := range hasApple {
//         if has {
//             step[i] = true
//             i = prev[i]
//             for step[i] != true {
//                 step[i] = true
//                 i = prev[i]
//             }
//         }
//     }
//     for _, v := range step {
//         if v { time++ }
//     }
//     return time * 2 // to and fro
// }

// dfs
func minTime(n int, edges [][]int, hasApple []bool) int {
    neighbors := make([][]int, len(hasApple))
    for _, e := range edges {
        neighbors[e[0]] = append(neighbors[e[0]], e[1])
        neighbors[e[1]] = append(neighbors[e[1]], e[0])
    }
    var dfs func(u int, p int) int
    dfs = func(u int, p int) int {
        time := 0
        for _, nei := range neighbors[u] {
            if nei == p { continue }
            neiTime := dfs(nei, u)
            if neiTime > 0 || hasApple[nei] {
                time += (2 + neiTime)
            }
        }
        return time
    }
    return dfs(0, -1)
}

func minTime1(n int, edges [][]int, hasApple []bool) int {
    g := make([][]int, n)
    for _, ch := range edges {
        x, y := ch[0], ch[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    // 标识有苹果的路
    var dfs1 func(i, fa int) bool
    dfs1 = func(i, fa int) bool {
        res := hasApple[i]
        for _, nx := range g[i] {
            if nx == fa { continue } 
            next := dfs1(nx, i)
            res = res || next
        }
        hasApple[i] = res
        return res
    }
    dfs1(0, -1)
    var dfs2 func(i, fa int) int
    dfs2 = func(i, fa int) int {
        res := 0
        if !hasApple[i] { return 0 }
        for _, nx := range g[i] {
            if nx == fa { continue }
            if !hasApple[nx] { continue }
            res += dfs2(nx, i) + 2 // 一来一回
        }
        return res
    }
    return  dfs2(0, n)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/04/23/min_time_collect_apple_1.png" />
    // Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,true,true,false]
    // Output: 8 
    // Explanation: The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  
    fmt.Println(minTime(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, []bool{false,false,true,false,true,true,false})) // 8
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/04/23/min_time_collect_apple_2.png" />
    // Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,true,false,false,true,false]
    // Output: 6
    // Explanation: The figure above represents the given tree where red vertices have an apple. One optimal path to collect all apples is shown by the green arrows.  
    fmt.Println(minTime(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, []bool{false,false,true,false,false,true,false})) // 6
    // Example 3:
    // Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,false,false,false,false,false,false]
    // Output: 0
    fmt.Println(minTime(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, []bool{false,false,false,false,false,false,false})) // 0

    fmt.Println(minTime(4, [][]int{{0,2},{0,3},{1,2}}, []bool{false,true,false,false})) // 4

    fmt.Println(minTime1(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, []bool{false,false,true,false,true,true,false})) // 8
    fmt.Println(minTime1(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, []bool{false,false,true,false,false,true,false})) // 6
    fmt.Println(minTime1(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}, []bool{false,false,false,false,false,false,false})) // 0
    fmt.Println(minTime1(4, [][]int{{0,2},{0,3},{1,2}}, []bool{false,true,false,false})) // 4
}