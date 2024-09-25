package main

// 1319. Number of Operations to Make Network Connected
// There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a network where connections[i] = [ai, bi] represents a connection between computers ai and bi. 
// Any computer can reach any other computer directly or indirectly through the network.

// You are given an initial computer network connections. 
// You can extract certain cables between two directly connected computers, 
// and place them between any pair of disconnected computers to make them directly connected.

// Return the minimum number of times you need to do this in order to make all the computers connected. 
// If it is not possible, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/02/sample_1_1677.png" />
// Input: n = 4, connections = [[0,1],[0,2],[1,2]]
// Output: 1
// Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/01/02/sample_2_1677.png" />
// Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
// Output: 2

// Example 3:
// Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
// Output: -1
// Explanation: There are not enough cables.

// Constraints:
//     1 <= n <= 10^5
//     1 <= connections.length <= min(n * (n - 1) / 2, 10^5)
//     connections[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     There are no repeated connections.
//     No two computers are connected by more than one cable.

import "fmt"

// union find
func makeConnected(n int, connections [][]int) int {
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }
    var find func(a int) int
    find = func(a int) int {
        if parent[a] == a { return a }
        parent[a] = find(parent[a])
        return parent[a]
    }
    union := func(a, b int) bool {
        a1, b1 := find(a), find(b)
        if a1 == b1 { return false }
        parent[b1] = a1
        return true
    }
    dup := 0
    for i := range connections {
        a, b := connections[i][0], connections[i][1]
        if !union(a, b) { 
            dup++ 
        }
    }
    set := make(map[int]bool)
    for i:= range parent {
        parent[i] = find(i)
        set[parent[i]] = true
    }
    res := len(set) - 1
    if dup < res{  return -1 }
    return res
}

func makeConnected1(n int, connections [][]int) int {
    if len(connections) < n-1 {
        return -1
    }
    uf := newUnionFind(n)
    for _, connect := range connections {
        uf.union(connect[0], connect[1])
    }
    return uf.setCount - 1
}

type unionFind struct {
    parent []int
    rank []int
    setCount int
}

func newUnionFind(size int) *unionFind {
    parent := make([]int, size)
    rank := make([]int, size)
    for i := range parent {
        parent[i] = i
        rank[i] = 0
    }
    return &unionFind{parent, rank, size}
}

func (uf *unionFind) find(p int) int {
    // 状态压缩
    if uf.parent[p] != p {
        uf.parent[p] = uf.find(uf.parent[p])
    }
    return uf.parent[p]
}

func (uf *unionFind) union(p, q int) {
    pRoot, qRoot := uf.find(p), uf.find(q)
    // 小树挂在大树下
    if pRoot != qRoot {
        if uf.rank[pRoot] > uf.rank[qRoot] {
            uf.parent[qRoot] = pRoot
        } else if uf.rank[pRoot] < uf.rank[qRoot] {
            uf.parent[pRoot] = qRoot
        } else {
            uf.parent[qRoot] = pRoot
            uf.rank[pRoot]++
        }
        uf.setCount--
    }
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/02/sample_1_1677.png" />
    // Input: n = 4, connections = [[0,1],[0,2],[1,2]]
    // Output: 1
    // Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.
    fmt.Println(makeConnected(4, [][]int{{0,1},{0,2},{1,2}})) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/01/02/sample_2_1677.png" />
    // Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
    // Output: 2
    fmt.Println(makeConnected(6, [][]int{{0,1},{0,2},{0,3},{1,2},{1,3}})) // 2
    // Example 3:
    // Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
    // Output: -1
    // Explanation: There are not enough cables.
    fmt.Println(makeConnected(6, [][]int{{0,1},{0,2},{0,3},{1,2}})) // -1

    fmt.Println(makeConnected1(4, [][]int{{0,1},{0,2},{1,2}})) // 1
    fmt.Println(makeConnected1(6, [][]int{{0,1},{0,2},{0,3},{1,2},{1,3}})) // 2
    fmt.Println(makeConnected1(6, [][]int{{0,1},{0,2},{0,3},{1,2}})) // -1
}