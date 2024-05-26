package main

// 305. Number of Islands II
// You are given an empty 2D binary grid grid of size m x n. 
// The grid represents a map where 0's represent water and 1's represent land. 
// Initially, all the cells of grid are water cells (i.e., all the cells are 0's).

// We may perform an add land operation which turns the water at position into a land. 
// You are given an array positions where positions[i] = [ri, ci] is the position (ri, ci) at which we should operate the ith operation.

// Return an array of integers answer where answer[i] is the number of islands after turning the cell (ri, ci) into a land.

// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. 
// You may assume all four edges of the grid are all surrounded by water.

// Example 1:
// <img src="" />
// Input: m = 3, n = 3, positions = [[0,0],[0,1],[1,2],[2,1]]
// Output: [1,1,2,3]
// Explanation:
// Initially, the 2d grid is filled with water.
// - Operation #1: addLand(0, 0) turns the water at grid[0][0] into a land. We have 1 island.
// - Operation #2: addLand(0, 1) turns the water at grid[0][1] into a land. We still have 1 island.
// - Operation #3: addLand(1, 2) turns the water at grid[1][2] into a land. We have 2 islands.
// - Operation #4: addLand(2, 1) turns the water at grid[2][1] into a land. We have 3 islands.

// Example 2:
// Input: m = 1, n = 1, positions = [[0,0]]
// Output: [1]
 
// Constraints:
//     1 <= m, n, positions.length <= 10^4
//     1 <= m * n <= 10^4
//     positions[i].length == 2
//     0 <= ri < m
//     0 <= ci < n

// Follow up: Could you solve it in time complexity O(k log(mn)), where k == positions.length?

import "fmt"

type UnionFind struct {
    count  int
    parent []int
    weight []int
}

// 并查集模板
func NewUnionFind(x int) *UnionFind {
    uf := UnionFind{
        // 与模板不一样的地方：初始没有岛屿(题目要求初始都是海)
        count:  0,
        parent: make([]int, x),
        weight: make([]int, x),
    }
    for i := 0; i < x; i++ {
        uf.parent[i] = i
        uf.weight[i] = 1
    }
    return &uf
}

func (this *UnionFind) Find(x int) int {
    for x != this.parent[x] {
        x = this.parent[x]
        this.parent[x] = this.parent[this.parent[x]]
    }
    return x
}

func (this *UnionFind) Union(a, b int) {
    rootA, rootB := this.Find(a), this.Find(b)
    if rootA == rootB {
        return
    }
    if this.weight[rootA] < this.weight[rootB] {
        this.parent[rootA] = rootB
        this.weight[rootB] += this.weight[rootA]
    } else {
        this.parent[rootB] = rootA
        this.weight[rootA] += this.weight[rootB]
    }
    this.count--
}

func (this *UnionFind) Connected(a, b int) bool {
    return this.Find(a) == this.Find(b)
}

func (this *UnionFind) Count() int {
    return this.count
}

func (this *UnionFind) Add() {
    this.count++
}

// 并查集
func numIslands2(m int, n int, positions [][]int) []int {
    u := NewUnionFind(m * n) // 构造函数
    res, visited := []int{}, make([]bool, m * n) // 二维变一维
    var inGrid = func(x, y int) bool { // 检查当前坐标是否越界
        return x >= 0 && y >= 0 && x < m && y < n
    }
    direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // 四个方向
    for _, p := range positions {
        x, y := p[0], p[1]
        index := x * n + y // 二维坐标转换成一维 用 index = x ∗ n + y 表示 (x,y) 
        if visited[index] {
            res = append(res, u.Count())
            continue
        }
        visited[index] = true // 未访问过
        u.Add() // 不相连岛屿数量 + 1
        // 尝试在当前坐标(x,y)与四周(上下左右)进行连接
        // 观察四个方向是否有岛屿，且是否连接
        // 如果其中一个方向有已经填过的岛屿，且没有连接在一起
        // 那么就可以将填海位置与已经填过的岛屿连接在一起，同时孤岛数量-1(在Union里面自动做掉)
        for _, dir := range direction {
            X := x + dir[0]
            Y := y + dir[1]
            Index := X * n + Y
            if inGrid(X, Y) && visited[Index] && !u.Connected(index, Index) {
                u.Union(index, Index)
            }
        }
        res = append(res, u.Count())
    }
    return res
}

func main() {
    // Example 1:
    // <img src="" />
    // Input: m = 3, n = 3, positions = [[0,0],[0,1],[1,2],[2,1]]
    // Output: [1,1,2,3]
    // Explanation:
    // Initially, the 2d grid is filled with water.
    // - Operation #1: addLand(0, 0) turns the water at grid[0][0] into a land. We have 1 island.
    // - Operation #2: addLand(0, 1) turns the water at grid[0][1] into a land. We still have 1 island.
    // - Operation #3: addLand(1, 2) turns the water at grid[1][2] into a land. We have 2 islands.
    // - Operation #4: addLand(2, 1) turns the water at grid[2][1] into a land. We have 3 islands.
    fmt.Println(numIslands2(3,3,[][]int{{0,0},{0,1},{1,2},{2,1}})) // [1,1,2,3]
    // Example 2:
    // Input: m = 1, n = 1, positions = [[0,0]]
    // Output: [1]
    fmt.Println(numIslands2(1,1,[][]int{{0,0}})) // [1]
}