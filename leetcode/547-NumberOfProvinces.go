package main

// 547. Number of Provinces
// There are n cities. Some of them are connected, while some are not. 
// If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

// A province is a group of directly or indirectly connected cities and no other cities outside of the group.
// You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
// Return the total number of provinces.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/graph1.jpg" />
// Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
// Output: 2

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/graph2.jpg" />
// Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
// Output: 3
 
// Constraints:
//     1 <= n <= 200
//     n == isConnected.length
//     n == isConnected[i].length
//     isConnected[i][j] is 1 or 0.
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


func main() {
    fmt.Println(findCircleNum([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3

    fmt.Println(findCircleNum1([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum1([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3

    fmt.Println(findCircleNum2([][]int{{1,1,0},{1,1,0},{0,0,1}})) // 2
    fmt.Println(findCircleNum2([][]int{{1,0,0},{0,1,0},{0,0,1}})) // 3
}