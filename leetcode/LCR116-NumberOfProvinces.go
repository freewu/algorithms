package main

// LCR 116. 省份数量
// 有 n 个城市，其中一些彼此相连，另一些没有相连。
// 如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

// 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

// 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

// 返回矩阵中 省份 的数量。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg" />
// 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
// 输出：2

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg" />
// 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
// 输出：3
 
// 提示：
//     1 <= n <= 200
//     n == isConnected.length
//     n == isConnected[i].length
//     isConnected[i][j] 为 1 或 0
//     isConnected[i][i] == 1
//     isConnected[i][j] == isConnected[j][i]

import "fmt"

// dfs
func findCircleNum(isConnected [][]int) int {
    res, visited := 0,make([]bool,len(isConnected))
    var dfs func(curr int, visited []bool, isConnected [][]int)
    dfs = func (curr int, visited []bool, isConnected [][]int) {
        visited[curr] = true
        for i := 0; i < len(isConnected); i++ {
            if isConnected[curr][i] == 0 { 
                continue 
            } 
            if !visited[i] {
                dfs(i, visited, isConnected)
            }
        } 
    }
    for i := 0; i < len(isConnected); i++ {
        if !visited[i] {
            res++
            dfs(i, visited, isConnected)
        }
    }
    return res
}

// bfs
func findCircleNum1(isConnected [][]int) int {
    res, visited := 0,make([]bool,len(isConnected))
    bfs := func(i int,visited[]bool,grid[][]int) {
        queue := []int{}
        queue = append(queue,i)
        visited[i] = true
        for len(queue) != 0 {
            pop := queue[0]
            queue = queue[1:]
            for j := 0; j < len(grid); j++ {
                if grid[pop][j] == 1 && !visited[j] {
                    queue = append(queue,j)
                    visited[j] = true
                }
            }
        }
    }
    for i:= 0; i < len(isConnected); i++ {
        if !visited[i] {
            bfs(i,visited,isConnected)
            res++
        }
    }
    return res
}

func findCircleNum2(isConnected [][]int) int {
    res, n := 0, len(isConnected)
    visited := make([]int, n)
    var dfs func(i int)
    dfs = func(i int) {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 && visited[j] == 0 {
                visited[j] = 1
                dfs(j)
            }
        }
    }
    for i := 0; i < n; i++ {
        if visited[i] == 0 {
            dfs(i)
            res++
        }
    }
    return res
}

// 并查集
func findCircleNum3(isConnected [][]int) int {
    n := len(isConnected)
    set := make([]int, n)
    for i := range set {
        set[i] = i
    }
    var find func(i int) int
    find = func(i int) int {
        if set[set[i]] != set[i] {
            set[i] = find(set[i])
        }
        return set[i]
    }
    union := func(i, j int) bool { 
        pi, pj := find(i), find(j)
        if pi == pj {
            return false
        }
        set[pi] = pj
        return true
    }
    for i := 0; i < n; i++ {
        for j := 0 ; j < len(isConnected[0]); j++ {
            if isConnected[i][j] == 1 {
                union(i, j)
            }
        }
    }
    mp := make(map[int]bool)
    for i := 0 ; i < n ; i++ {
        mp[find(i)] = true
    }
    return len(mp)
}


func main() {
    fmt.Println(findCircleNum([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3

    fmt.Println(findCircleNum1([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum1([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3

    fmt.Println(findCircleNum2([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum2([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3

    fmt.Println(findCircleNum3([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum3([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3
}